# Deploy Alluxio on Kubernetes


Alluxio can be run on Kubernetes. This guide demonstrates how to run Alluxio
on Kubernetes using the specification included in the Alluxio Docker image or `helm`.


## Prerequisites

- A Kubernetes cluster (version 1.11+) with Beta feature gate APIs enabled
  - The [Alluxio Helm chart](https://github.com/Alluxio/alluxio/tree/master/integration/kubernetes/helm-chart/alluxio)
  which the Kubernetes resource specifications are built from supports
  Kubernetes version 1.11+.
  - Beta feature gates are [enabled by default](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/#feature-stages)
  for Kubernetes cluster installations
- Cluster access to an Alluxio Docker image [alluxio/alluxio](https://hub.docker.com/r/alluxio/alluxio/).
If using a private Docker registry, refer to the Kubernetes private image registry
[documentation](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/).
- Ensure the cluster's [Kubernetes Network Policy](https://kubernetes.io/docs/concepts/services-networking/network-policies/)
allows for connectivity between applications (Alluxio clients) and the Alluxio Pods on the defined
ports.

## Basic Setup

This tutorial walks through a basic Alluxio setup on Kubernetes. Alluxio supports two methods of
installation on Kubernetes: either using [helm](https://helm.sh/docs/) charts or using `kubectl`.
When available, `helm` is the preferred way to install Alluxio.
If `helm` is not available or if additional deployment customization is desired, `kubectl` can be
used directly using native Kubernetes resource specifications.

> Note: From Alluxio 2.3 on, Alluxio only supports helm 3.
> See how to migrate from helm 2 to 3 [here](https://helm.sh/docs/topics/v2_v3_migration/).


<details><summary>(Optional) Copy the Alluxio Helm chart to a private Helm repository</summary>

The Alluxio Helm chart source code is located [here](https://github.com/Alluxio/alluxio/tree/master/integration/kubernetes/helm-chart/alluxio).
Alternatively, you can extract the Helm chart directory from the Alluxio Docker image:

```console
$ id=$(docker create alluxio/alluxio:2.9.5)
$ docker cp $id:/opt/alluxio/integration/kubernetes/ - > kubernetes.tar
$ docker rm -v $id 1>/dev/null
$ tar -xvf kubernetes.tar
$ cd kubernetes/helm-chart/alluxio
```
</details>
<details><summary>(Optional) Provision a Persistent Volume</summary>

Depending on the configuration used to deploy Alluxio, you will likely
need to provision [Persistent Volumes](https://kubernetes.io/docs/concepts/storage/persistent-volumes/).
- [Embedded Journal](../operation/Journal.md#configuring-embedded-journal)
requires a Persistent Volume for each master Pod to be provisioned and is the preferred HA mechanism
for Alluxio on Kubernetes. The volume, once claimed, is persisted across restarts of the master process.
- When using the [UFS Journal](../operation/Journal.md#configuring-ufs-journal)
an Alluxio master can also be configured to use a persistent volume for storing the journal.
If Alluxio is configured to use a UFS journal and with an external journal location
like HDFS, the rest of this section can be skipped.
- When Alluxio workers have [short-circuit access](../kubernetes/Running-Alluxio-On-Kubernetes.md#enable-short-circuit-access),
you may need to use Volumes to mount the domain socket to the workers.

There are multiple ways to create a PersistentVolume.
This is an example which defines one with `hostPath` for the Alluxio Master journal:
```yaml
# Name the file alluxio-master-journal-pv.yaml
kind: PersistentVolume
apiVersion: v1
metadata:
  name: alluxio-journal-0
  labels:
    type: local
spec:
  storageClassName: standard
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: /tmp/alluxio-journal-0
```
Note:
- By default each journal volume should be at least 1Gi, because each Alluxio master Pod
will have one PersistentVolumeClaim that requests for 1Gi storage. You will see how to configure
the journal size in later sections.
- If this `hostPath` is not already present on the host, Kubernetes can be configured to create it. However
the assigned user:group permissions may prevent the Alluxio masters & workers from accessing it.
Please ensure the permissions are set to allow the pods to access the directory.
  - See the [Kubernetes volume docs](https://kubernetes.io/docs/concepts/storage/volumes/#hostpath) for more details
  - From Alluxio v2.1 on, Alluxio Docker containers will run as non-root user `alluxio`
with UID 1000 and GID 1000 by default.

Then create the persistent volume with `kubectl`:
```console
$ kubectl create -f alluxio-master-journal-pv.yaml
```
</details>

### Deploy

<details><summary>helm</summary>

#### Prerequisites

A. Install Helm

You should have helm 3.X installed.
You can install helm following instructions [here](https://helm.sh/docs/intro/install/).

B. A helm repo with the Alluxio helm chart must be available.

```console
$ helm repo add alluxio-charts https://alluxio-charts.storage.googleapis.com/openSource/2.9.5
```

#### Configuration

Once the helm repository is available, prepare the Alluxio configuration.
The minimal configuration must contain the under storage address:
```properties
properties:
  alluxio.master.mount.table.root.ufs: "<under_storage_address>"
```
> Note: The Alluxio under filesystem address MUST be modified. Any credentials MUST be modified.

To view the complete list of supported properties run the `helm inspect` command:
```console
$ helm inspect values alluxio-charts/alluxio
```

The remainder of this section describes various configuration options with examples.

##### Example: Amazon S3 as the under store
To [mount S3](../ufs/S3.md#root-mount-point) at the root of Alluxio
namespace specify all required properties as a key-value pair under `properties`.

```properties
properties:
  alluxio.master.mount.table.root.ufs: "s3a://<bucket>"
  alluxio.master.mount.table.root.option.s3a.accessKeyId: "<accessKey>"
  alluxio.master.mount.table.root.option.s3a.secretKey: "<secretKey>"
```
##### Example: Single Master and Journal in a Persistent Volume
The following configures [UFS Journal](../operation/Journal.md#configuring-ufs-journal)
with a persistent volume claim mounted locally to the master Pod at location `/journal`.

```properties
master:
  count: 1 # For multiMaster mode increase this to >1

journal:
  # [ Required values ]
  type: "UFS" # One of "UFS" or "EMBEDDED"
  folder: "/journal" # Master journal directory or equivalent storage path
  #
  # [ Conditionally required values ]
  #
  ## [ UFS-backed journal options ]
  ## - required when using a UFS-type journal (journal.type="UFS")
  ##
  ## ufsType is one of "local" or "HDFS"
  ## - "local" results in a PV being allocated to each Master Pod as the journal
  ## - "HDFS" results in no PV allocation, it is up to you to ensure you have
  ##   properly configured the required Alluxio properties for Alluxio to access
  ##   the HDFS URI designated as the journal folder
  ufsType: "local"
  #
  ## [ K8s volume options ]
  ## - required when using an EMBEDDED journal (journal.type="EMBEDDED")
  ## - required when using a local UFS journal (journal.type="UFS" and journal.ufsType="local")
  ##
  ## volumeType controls the type of journal volume.
  volumeType: persistentVolumeClaim # One of "persistentVolumeClaim" or "emptyDir"
  ## size sets the requested storage capacity for a persistentVolumeClaim,
  ## or the sizeLimit on an emptyDir PV.
  size: 1Gi
  ### Unique attributes to use when the journal is persistentVolumeClaim
  storageClass: "standard"
  accessModes:
    - ReadWriteOnce
```
##### Example: Single Master and Journal in an `emptyDir` Volume
The following configures [UFS Journal](../operation/Journal.md#configuring-ufs-journal)
with an `emptyDir` volume mounted locally to the master Pod at location `/journal`.

```properties
master:
  count: 1 # For multiMaster mode increase this to >1

journal:
  # [ Required values ]
  type: "UFS" # One of "UFS" or "EMBEDDED"
  folder: "/journal" # Master journal directory or equivalent storage path
  #
  # [ Conditionally required values ]
  #
  ## [ UFS-backed journal options ]
  ## - required when using a UFS-type journal (journal.type="UFS")
  ##
  ## ufsType is one of "local" or "HDFS"
  ## - "local" results in a PV being allocated to each Master Pod as the journal
  ## - "HDFS" results in no PV allocation, it is up to you to ensure you have
  ##   properly configured the required Alluxio properties for Alluxio to access
  ##   the HDFS URI designated as the journal folder
  ufsType: "local"
  #
  ## [ K8s volume options ]
  ## - required when using an EMBEDDED journal (journal.type="EMBEDDED")
  ## - required when using a local UFS journal (journal.type="UFS" and journal.ufsType="local")
  ##
  ## volumeType controls the type of journal volume.
  volumeType: emptyDir # One of "persistentVolumeClaim" or "emptyDir"
  ## size sets the requested storage capacity for a persistentVolumeClaim,
  ## or the sizeLimit on an emptyDir PV.
  size: 1Gi
  ### Unique attributes to use when the journal is emptyDir
  medium: ""
```

>Note: An `emptyDir` volume has the same lifetime as the Pod.
It is NOT a persistent storage.
The Alluxio journal will be LOST when the Pod is restarted or rescheduled.
Please only use this for experimental use cases.
Check [emptyDir](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir) for more details.

##### Example: HDFS as Journal
First create secrets for any configuration required by an HDFS client.
These are mounted under `/secrets`.
```console
$ kubectl create secret generic alluxio-hdfs-config --from-file=${HADOOP_CONF_DIR}/core-site.xml --from-file=${HADOOP_CONF_DIR}/hdfs-site.xml
```

```properties
journal:
  # [ Required values ]
  type: "UFS" # One of "UFS" or "EMBEDDED"
  folder: "hdfs://{$hostname}:{$hostport}/journal" # Master journal directory or equivalent storage path
  #
  # [ Conditionally required values ]
  #
  ## [ UFS-backed journal options ]
  ## - required when using a UFS-type journal (journal.type="UFS")
  ##
  ## ufsType is one of "local" or "HDFS"
  ## - "local" results in a PV being allocated to each Master Pod as the journal
  ## - "HDFS" results in no PV allocation, it is up to you to ensure you have
  ##   properly configured the required Alluxio properties for Alluxio to access
  ##   the HDFS URI designated as the journal folder
  ufsType: "HDFS"

properties:
  alluxio.master.mount.table.root.ufs: "hdfs://{$hostname}:{$hostport}/alluxio"
  alluxio.master.journal.ufs.option.alluxio.underfs.hdfs.configuration: "/secrets/hdfsConfig/core-site.xml:/secrets/hdfsConfig/hdfs-site.xml"

secrets:
  master:
    alluxio-hdfs-config: hdfsConfig
  worker:
    alluxio-hdfs-config: hdfsConfig
```
##### Example: Multi-master with Embedded Journal in Persistent Volumes

```properties
master:
  count: 3

  # [ Required values ]
  type: "EMBEDDED" # One of "UFS" or "EMBEDDED"
  folder: "/journal" # Master journal directory or equivalent storage path
  #
  # [ Conditionally required values ]
  #
  ## [ K8s volume options ]
  ## - required when using an EMBEDDED journal (journal.type="EMBEDDED")
  ## - required when using a local UFS journal (journal.type="UFS" and journal.ufsType="local")
  ##
  ## volumeType controls the type of journal volume.
  volumeType: persistentVolumeClaim # One of "persistentVolumeClaim" or "emptyDir"
  ## size sets the requested storage capacity for a persistentVolumeClaim,
  ## or the sizeLimit on an emptyDir PV.
  size: 1Gi
  ### Unique attributes to use when the journal is persistentVolumeClaim
  storageClass: "standard"
  accessModes:
    - ReadWriteOnce
```
##### Example: Multi-master with Embedded Journal in `emptyDir` Volumes

```properties
master:
  count: 3

journal:
  # [ Required values ]
  type: "EMBEDDED" # One of "UFS" or "EMBEDDED"
  folder: "/journal" # Master journal directory or equivalent storage path
  #
  # [ Conditionally required values ]
  #
  ## [ K8s volume options ]
  ## - required when using an EMBEDDED journal (journal.type="EMBEDDED")
  ## - required when using a local UFS journal (journal.type="UFS" and journal.ufsType="local")
  ##
  ## volumeType controls the type of journal volume.
  volumeType: emptyDir # One of "persistentVolumeClaim" or "emptyDir"
  ## size sets the requested storage capacity for a persistentVolumeClaim,
  ## or the sizeLimit on an emptyDir PV.
  size: 1Gi
  ### Unique attributes to use when the journal is emptyDir
  medium: ""
```

>Note: An `emptyDir` volume has the same lifetime as the Pod.
It is NOT a persistent storage.
The Alluxio journal will be LOST when the Pod is restarted or rescheduled.
Please only use this for experimental use cases.
Check [emptyDir](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir) for more details.

##### Example: HDFS as the under store

First create secrets for any configuration required by an HDFS client.
These are mounted under `/secrets`.
```console
$ kubectl create secret generic alluxio-hdfs-config --from-file=${HADOOP_CONF_DIR}/core-site.xml --from-file=${HADOOP_CONF_DIR}/hdfs-site.xml
```

```properties
properties:
  alluxio.master.mount.table.root.ufs: "hdfs://<ns>"
  alluxio.master.mount.table.root.option.alluxio.underfs.hdfs.configuration: "/secrets/hdfsConfig/core-site.xml:/secrets/hdfsConfig/hdfs-site.xml"
secrets:
  master:
    alluxio-hdfs-config: hdfsConfig
  worker:
    alluxio-hdfs-config: hdfsConfig
```
##### Example: Off-heap Metastore Management in Persistent Volumes

The following configuration creates a `PersistentVolumeClaim` for each Alluxio master Pod with the
specified configuration and configures the Pod to use the volume for an on-disk RocksDB-based
metastore.
```properties
properties:
  alluxio.master.metastore: ROCKS
  alluxio.master.metastore.dir: /metastore

metastore:
  volumeType: persistentVolumeClaim # Options: "persistentVolumeClaim" or "emptyDir"
  size: 1Gi
  mountPath: /metastore
  # Attributes to use when the metastore is persistentVolumeClaim
  storageClass: "standard"
  accessModes:
   - ReadWriteOnce
```
##### Example: Off-heap Metastore Management in `emptyDir` Volumes

The following configuration creates an `emptyDir` Volume for each Alluxio master Pod with the
specified configuration and configures the Pod to use the volume for an on-disk RocksDB-based
metastore.

```properties
properties:
  alluxio.master.metastore: ROCKS
  alluxio.master.metastore.dir: /metastore

metastore:
  volumeType: emptyDir # Options: "persistentVolumeClaim" or "emptyDir"
  size: 1Gi
  mountPath: /metastore
  # Attributes to use when the metastore is emptyDir
  medium: ""
```

>Note: An `emptyDir` volume has the same lifetime as the Pod.
It is NOT a persistent storage.
The Alluxio metadata will be LOST when the Pod is restarted or rescheduled.
Please only use this for experimental use cases.
Check [emptyDir](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir) for more details.

##### Example: Multiple Secrets

Multiple secrets can be mounted to both master and worker Pods.
The format for the section for each Pod is `<secretName>: <mountPath>`
```properties
secrets:
  master:
    alluxio-hdfs-config: hdfsConfig
    alluxio-ceph-config: cephConfig
  worker:
    alluxio-hdfs-config: hdfsConfig
    alluxio-ceph-config: cephConfig
```
##### Examples: Alluxio Storage Management

Alluxio manages local storage, including memory, on the worker Pods.
[Multiple-Tier Storage](../core-services/Caching.md#multiple-tier-storage)
can be configured using the following reference configurations.

There 3 supported volume `type`: [hostPath](https://kubernetes.io/docs/concepts/storage/volumes/#hostpath), [emptyDir](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir)
and [persistentVolumeClaim](https://kubernetes.io/docs/concepts/storage/volumes/#persistentvolumeclaim).

**Memory Tier Only**

```properties
tieredstore:
  levels:
  - level: 0
    mediumtype: MEM
    path: /dev/shm
    type: emptyDir
    high: 0.95
    low: 0.7
```

**Memory and SSD Storage in Multiple-Tiers**

```properties
tieredstore:
  levels:
  - level: 0
    mediumtype: MEM
    path: /dev/shm
    type: hostPath
    high: 0.95
    low: 0.7
  - level: 1
    mediumtype: SSD
    path: /ssd-disk
    type: hostPath
    high: 0.95
    low: 0.7
```

> Note: If a `hostPath` file or directory is created at runtime, it can only be used by the `root` user.
`hostPath` volumes do not have resource limits.
You can either run Alluxio containers with `root` or make sure the local paths exist and are accessible to
the user `alluxio` with UID and GID 1000.
You can find more details [here](https://kubernetes.io/docs/concepts/storage/volumes/#hostpath).

**Memory and SSD Storage in Multiple-Tiers, using PVC**

You can also use PVCs for each tier and provision [PersistentVolume](https://kubernetes.io/docs/concepts/storage/persistent-volumes/).
For worker tiered storage please use either `hostPath` or `local` volume so that the worker will
read and write locally to achieve the best performance.

```properties
tieredstore:
  levels:
  - level: 0
    mediumtype: MEM
    path: /dev/shm
    type: persistentVolumeClaim
    name: alluxio-mem
    quota: 1G
    high: 0.95
    low: 0.7
  - level: 1
    mediumtype: SSD
    path: /ssd-disk
    type: persistentVolumeClaim
    name: alluxio-ssd
    quota: 10G
    high: 0.95
    low: 0.7
```

> Note: There is one PVC per tier.
When the PVC is bound to a PV of type `hostPath` or `local`, each worker Pod will resolve to the
local path on the Node.
Please also note that a `local` volumes requires `nodeAffinity` and Pods using this volume can only
run on the Nodes specified in the `nodeAffinity` rule of this volume.
You can find more details [here](https://kubernetes.io/docs/concepts/storage/volumes/#local).

**Memory and SSD Storage in a Single-Tier**

You can also have multiple volumes on the same tier.
This configuration will create one `persistentVolumeClaim` for each volume.

```properties
tieredstore:
  levels:
  - level: 0
    mediumtype: MEM,SSD
    path: /dev/shm,/alluxio-ssd
    type: persistentVolumeClaim
    name: alluxio-mem,alluxio-ssd
    quota: 1GB,10GB
    high: 0.95
    low: 0.7
```

#### Install

Once the configuration is finalized in a file named `config.yaml`, install as follows:
```console
$ helm install alluxio -f config.yaml alluxio-charts/alluxio
```

In order to configure the Alluxio Master pod for use, you will need to format the Alluxio journal.

#### Format Journal

The master Pods in the StatefulSet use an `initContainer` to format the journal on startup.
This `initContainer` is switched on by `journal.format.runFormat=true`.
By default, the journal is not formatted when the master starts.

You can trigger the journal formatting by upgrading the existing helm deployment with
`journal.format.runFormat=true`.
```console
# Use the same config.yaml and switch on journal formatting
$ helm upgrade alluxio -f config.yaml --set journal.format.runFormat=true alluxio-charts/alluxio
```

> Note: `helm upgrade` will re-create the master Pods.

Or you can trigger the journal formatting at deployment.
```console
$ helm install alluxio -f config.yaml --set journal.format.runFormat=true alluxio-charts/alluxio
```

> Note: From Alluxio v2.1 on, Alluxio Docker containers will run as non-root user `alluxio`
with UID 1000 and GID 1000 by default.
You should make sure the journal is formatted using the same user that the Alluxio master Pod runs as.

#### Configure Worker Volumes

Additional configuration is required for the Alluxio Worker pod to be ready for use.
See the section for [enabling worker short-circuit access](../kubernetes/Running-Alluxio-On-Kubernetes.md#enable-short-circuit-access).

#### Uninstall

Uninstall Alluxio as follows:
```console
$ helm delete alluxio
```

</details>
<details><summary>kubectl</summary>

#### Choose the Sample YAML Template

First, extract the pre-templated Kubernetes specification YAMLs
from the Alluxio docker image:

```console
$ id=$(docker create alluxio/alluxio:2.9.5)
$ docker cp $id:/opt/alluxio/integration/kubernetes/ - > kubernetes.tar
$ docker rm -v $id 1>/dev/null
$ tar -xvf kubernetes.tar
$ cd kubernetes
```

The extracted directory contains a set of YAML templates generated from our Helm chart
for common deployment scenarios in the sub-directories:
*singleMaster-localJournal*, *singleMaster-hdfsJournal*, and
*multiMaster-embeddedJournal*.
> *singleMaster* means the templates generate 1 Alluxio master process, while *multiMaster* means 3.
*embedded* and *ufs* are the 2 [journal modes](../operation/Journal.md)
that Alluxio supports.

- *singleMaster-localJournal* directory gives you the necessary Kubernetes ConfigMap, 1 Alluxio
master process and a set of Alluxio workers.
The Alluxio master writes journal to the journal volume requested by `volumeClaimTemplates`.
- *multiMaster-EmbeddedJournal* directory gives you the Kubernetes ConfigMap, 3 Alluxio masters and
a set of Alluxio workers.
Each Alluxio master writes journal to its own journal volume requested by `volumeClaimTemplates`.
- *singleMaster-hdfsJournal* directory gives you the Kubernetes ConfigMap, 1 Alluxio master with a
set of workers.
The journal is in a shared UFS location. In this template we use HDFS as the UFS.

For customized templated YAMLs, see the [README](https://github.com/Alluxio/alluxio/tree/master/integration/kubernetes#generate-kubectl-yaml-templates-from-helm-chart)
for how to use `helm-generate.sh`. Otherwise you may manually write or modify YAML files as you see fit.

#### Configuration

Once the deployment option is chosen, copy the template from the desired sub-directory:
```console
$ cp alluxio-configmap.yaml.template alluxio-configmap.yaml
```

Modify or add any configuration properties as required.
The Alluxio under filesystem address MUST be modified. Any credentials MUST be modified.
Add to `ALLUXIO_JAVA_OPTS`:
```properties
-Dalluxio.master.mount.table.root.ufs=<under_storage_address>
```

Note:
- Replace `<under_storage_address>` with the appropriate URI, for example s3://my-bucket.
If using an under storage which requires credentials be sure to specify those as well.
- When running Alluxio with host networking, the ports assigned to Alluxio services must
not be occupied beforehand.

Create a ConfigMap.
```console
$ kubectl create -f alluxio-configmap.yaml
```

#### Install

***Prepare the Specification.*** Prepare the Alluxio deployment specs from the templates.
Modify any parameters required, such as location of the **Docker image**, and CPU and memory
requirements for Pods.

For the master(s), create the `Service` and `StatefulSet`:
```console
$ mv master/alluxio-master-service.yaml.template master/alluxio-master-service.yaml
$ mv master/alluxio-master-statefulset.yaml.template master/alluxio-master-statefulset.yaml
```
> Note: `alluxio-master-statefulset.yaml` uses `volumeClaimTemplates` to define the journal volume
for each master if it needs one.

For the workers, create the `DaemonSet`:
```console
$ mv worker/alluxio-worker-daemonset.yaml.template worker/alluxio-worker-daemonset.yaml
```

Note: Please make sure that the version of the Kubernetes specification matches the version of the
Alluxio Docker image being used.

##### (Optional) Remote Storage Access

Additional steps may be required when Alluxio is connecting to storage hosts outside the
Kubernetes cluster it is deployed on. The remainder of this section explains how to configure the
connection to a remote HDFS accessible but not managed by Kubernetes.

**Step 1: Add `hostAliases` for your HDFS connection.**  Kubernetes Pods don't recognize network
hostnames that are not managed by Kubernetes (not a Kubernetes Service), unless if specified by
[hostAliases](https://kubernetes.io/docs/concepts/services-networking/add-entries-to-pod-etc-hosts-with-host-aliases/#adding-additional-entries-with-hostaliases).

For example if your HDFS service can be reached at `hdfs://<namenode>:9000` where `<namenode>` is a
hostname, you will need to add `hostAliases` in the `spec` for all Alluxio Pods creating a map from
hostnames to IP addresses.

```yaml
spec:
  hostAliases:
  - ip: "<namenode_ip>"
    hostnames:
    - "<namenode>"
```

For the case of a StatefulSet or DaemonSet as used in `alluxio-master-statefulset.yaml.template` and
`alluxio-worker-daemonset.yaml.template`, `hostAliases` section should be added to each section of
`spec.template.spec` like below.

```yaml
kind: StatefulSet
metadata:
  name: alluxio-master
spec:
  ...
  serviceName: "alluxio-master"
  replicas: 1
  template:
    metadata:
      labels:
        app: alluxio-master
    spec:
      hostAliases:
      - ip: "ip for hdfs-host"
        hostnames:
        - "hdfs-host"
```

**Step 2: Create Kubernetes Secret for HDFS configuration files.** Run the following command to
create a Kubernetes Secret for the HDFS client configuration.

```console
kubectl create secret generic alluxio-hdfs-config --from-file=${HADOOP_CONF_DIR}/core-site.xml --from-file=${HADOOP_CONF_DIR}/hdfs-site.xml
```
These two configuration files are referred in `alluxio-master-statefulset.yaml` and `alluxio-worker-daemonset.yaml`.
Alluxio processes need the HDFS configuration files to connect, and the location of these files in
the container is controlled by property `alluxio.underfs.hdfs.configuration`.

**Step 3: Modify `alluxio-configmap.yaml.template`.** Now that your Pods know how to talk to your
HDFS service, update `alluxio.master.journal.folder` and `alluxio.master.mount.table.root.ufs` to
point to the desired HDFS destination.


Once all the pre-requisites and configuration have been setup, deploy Alluxio.
```console
$ kubectl create -f ./master/
$ kubectl create -f ./worker/
```

In order to configure the Alluxio Master pod for use, you will need to format the Alluxio journal.

#### Format Journal

You can manually add an `initContainer` to format the journal on Pod creation time.
This `initContainer` will run `alluxio formatJournal` when the Pod is created and formats the journal.

```yaml
- name: journal-format
  image: alluxio/alluxio:2.9.5
  imagePullPolicy: IfNotPresent
  securityContext:
    runAsUser: 1000
  command: ["alluxio","formatJournal"]
  envFrom:
    - configMapRef:
      name: alluxio-config
  volumeMounts:
    - name: alluxio-journal
      mountPath: /journal
```

> Note: From Alluxio v2.1 on, Alluxio Docker containers will run as non-root user `alluxio`
with UID 1000 and GID 1000 by default.
You should make sure the journal is formatted using the same user that the Alluxio master Pod runs as.

#### Configure Worker Volumes

Additional configuration is required for the Alluxio Worker pod to be ready for use.
See the section for [enabling worker short-circuit access](../kubernetes/Running-Alluxio-On-Kubernetes.md#enable-short-circuit-access).

#### Upgrade

This section will go over how to upgrade Alluxio in your Kubernetes cluster with `kubectl`.
##### Upgrading Alluxio

**Step 1: Upgrade the docker image version tag**

Each released Alluxio version will have the corresponding docker image released on
[dockerhub](https://hub.docker.com/r/alluxio/alluxio).

You should update the `image` field of all the Alluxio containers to use the target version tag.
Tag `latest` will point to the latest stable version.

For example, if you want to upgrade Alluxio to the latest stable version, update the containers as
below:

```yaml
containers:
- name: alluxio-master
  image: alluxio/alluxio:latest
  imagePullPolicy: IfNotPresent
  ...
- name: alluxio-job-master
  image: alluxio/alluxio:latest
  imagePullPolicy: IfNotPresent
  ...
```

**Step 2: Stop running Alluxio master and worker Pods**

Kill all running Alluxio worker Pods by deleting its DaemonSet.

```console
$ kubectl delete daemonset -l app=alluxio
```

Then kill all running Alluxio master Pods by killing each StatefulSet and each Service with label
`app=alluxio`.

```console
$ kubectl delete service -l app=alluxio
$ kubectl delete statefulset -l app=alluxio
```

Make sure all the Pods have been terminated before you move on to the next step.

**Step 3: Format journal and Alluxio storage if necessary**

Check the [Alluxio upgrade guide page](../administration/Upgrade.md)
for whether the Alluxio master journal has to be formatted. If no format is needed,
you are ready to skip the rest of this section and move on to restart all
Alluxio master and worker Pods.

You can follow [formatting journal with kubectl](../kubernetes/Running-Alluxio-On-Kubernetes.md#format-journal-1)
to format the Alluxio journals.

If you are running Alluxio workers with [tiered storage](../core-services/Caching.md#multiple-tier-storage),
and you have Persistent Volumes configured for Alluxio, the storage should be cleaned up too.
You should delete and recreate the Persistent Volumes.

Once all the journals and Alluxio storage have been formatted, you are ready to restart the Alluxio
master and worker Pods.

**Step 4: Restart Alluxio master and worker Pods**

Now that Alluxio masters and worker containers all use your desired version. Restart them to let it
take effect.

Now restart the Alluxio master and worker Pods from the YAML files.

```console
$ kubectl create -f ./master/
$ kubectl create -f ./worker/
```

**Step 5: Verify the Alluxio master and worker Pods are back up**

You should verify the Alluxio Pods are back up in Running status.

```console
# You should see all Alluxio master and worker Pods
$ kubectl get pods
```

You can do more comprehensive verification following [Verify Alluxio](../deploy/Running-Alluxio-Locally.md#verify-alluxio-is-running).

#### Uninstall

Uninstall Alluxio as follows:
```console
$ kubectl delete -f ./worker/
$ kubectl delete -f ./master/
$ kubectl delete configmap alluxio-config
```
> Note: This will delete all resources under `./master/` and `./worker/`.
Be careful if you have persistent volumes or other important resources you want to keep under those directories.

</details>

### Verify

If using persistent volumes, the status of the volume(s) should change
to `CLAIMED` and the status of the volume claims should be `BOUNDED`.
You can validate the status of your PersistentVolume and PersistentVolumeClaims
using the follow `kubectl` commands:
```console
$ kubectl get pv
$ kubectl get pvc
```
- If you have unbound PersistentVolumeClaims, please ensure you have provisioned
matching PersistentVolumes. See "(Optional) Provision a Persistent Volume" in
[Basic Setup](../kubernetes/Running-Alluxio-On-Kubernetes.md#basic-setup).

Once ready, access the Alluxio CLI from the master Pod and run basic I/O tests.
```console
$ kubectl exec -ti alluxio-master-0 -- /bin/bash
```

From the master Pod, execute the following:
```console
$ alluxio runTests
```

### Access the Web UI

The Alluxio UI can be accessed from outside the kubernetes cluster using port forwarding.
```console
$ kubectl port-forward alluxio-master-$i <local-port>:19999
```
The command above allocates a port on the local node `<local-port>` and forwards traffic
on `<local-port>` to port 19999 of pod `alluxio-master-$i`.
The pod `alluxio-master-$i` does NOT have to be on the node you are running this command.

> Note: `i=0` for the first master Pod. When running multiple masters, forward port for each
master. Only the primary master serves the Web UI.

For example, you are on a node with hostname `master-node-1` and you would like to serve
the Alluxio master web UI for `alluxio-master-0` on `master-node-1:8080`.
Here's the command you can run:
```console
[alice@master-node-1 ~]$ kubectl port-forward --address 0.0.0.0 pods/alluxio-master-0 8080:19999
``` 
This forwards the local port `master-node-1:8080` to the port on the Pod `alluxio-master-0:19999`.
The Pod `alluxio-master-0` does NOT need to be running on `master-node-1`.

You will see messages like below when there are incoming connections. 
```console
[alice@master-node-1 ~]$ kubectl port-forward --address 0.0.0.0 alluxio-master-0 8080:19999
Forwarding from 0.0.0.0:8080 -> 19999
Handling connection for 8080
Handling connection for 8080
Handling connection for 8080
Handling connection for 8080
```
You can terminate the process to stop the port forwarding,
with either `Ctrl + C` or `kill`.

For more information about K8s port-forward see the [K8s doc](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#port-forward).

## Advanced Setup

### Enable Short-circuit Access

Short-circuit access enables clients to perform read and write operations directly against the
worker bypassing the networking interface.
For performance-critical applications it is recommended to enable short-circuit operations
against Alluxio because it can increase a client's read and write throughput when co-located with
an Alluxio worker.

This feature is enabled by default (see next section to disable this feature), however requires extra configuration to work properly in
Kubernetes environments.

There are two modes for using short-circuit.

#### Option1: Use local mode

In this mode, the Alluxio client and local Alluxio worker recognize each other if the client hostname
matches the worker hostname.
This is called *Hostname Introspection*.
In this mode, the Alluxio client and local Alluxio worker share the tiered storage of Alluxio worker.

<details><summary>helm</summary>

You can use `local` policy by setting the properties as below:

```properties
shortCircuit:
  enabled: true
  policy: local
```

</details>
<details><summary>kubectl</summary>

In your `alluxio-configmap.yaml` you should add the following properties to `ALLUXIO_WORKER_JAVA_OPTS`:

```properties
-Dalluxio.user.short.circuit.enabled=true \
-Dalluxio.worker.data.server.domain.socket.as.uuid=false
```

Also you should remove the property `-Dalluxio.worker.data.server.domain.socket.address`.

</details>


#### Option2: Use uuid (default)

This is the **default** policy used for short-circuit in Kubernetes.

If the client or worker container is using virtual networking, their hostnames may not match.
In such a scenario, set the following property to use filesystem inspection to enable short-circuit
operations and **make sure the client container mounts the directory specified as the domain socket
path**.
Short-circuit writes are then enabled if the worker UUID is located on the client filesystem.

***Domain Socket Path.***
The domain socket is a volume which should be mounted on:

- All Alluxio workers
- All application containers which intend to read/write through Alluxio

This domain socket volume can be either a `PersistentVolumeClaim` or a `hostPath Volume`.

***Use PersistentVolumeClaim.***
By default, this domain socket volume is a `PersistentVolumeClaim`.
You need to provision a `PersistentVolume` to this `PersistentVolumeClaim`.
And this `PersistentVolume` should be either `local` or `hostPath`.

<details><summary>helm</summary>

You can use `uuid` policy by setting the properties as below:

```properties
# These are the default configurations
shortCircuit:
  enabled: true
  policy: uuid
  size: 1Mi
  # volumeType controls the type of shortCircuit volume.
  # It can be "persistentVolumeClaim" or "hostPath"
  volumeType: persistentVolumeClaim
  # Attributes to use if the domain socket volume is PVC
  pvcName: alluxio-worker-domain-socket
  accessModes:
    - ReadWriteOnce
  storageClass: standard
```

The field `shortCircuit.pvcName` defines the name of the `PersistentVolumeClaim` for domain socket.
This PVC will be created as part of `helm install`.

</details>
<details><summary>kubectl</summary>

You should verify the following properties in `ALLUXIO_WORKER_JAVA_OPTS`.
Actually they are set to these values by default:
```properties
-Dalluxio.worker.data.server.domain.socket.address=/opt/domain -Dalluxio.worker.data.server.domain.socket.as.uuid=true
```

Also you should make sure the worker Pods have domain socket defined in the `volumes`,
and all relevant containers have the domain socket volume mounted.
The domain socket volume is defined as below by default:
```properties
volumes:
  - name: alluxio-domain
    persistentVolumeClaim:
      claimName: "alluxio-worker-domain-socket"
```

> Note: Compute application containers **MUST** mount the domain socket volume to the same path
(`/opt/domain`) as configured for the Alluxio workers.

The `PersistenceVolumeClaim` is defined in `worker/alluxio-worker-pvc.yaml.template`.

</details>

***Use hostPath Volume.***
You can also directly define the workers to use a `hostPath Volume` for domain socket.


<details><summary>helm</summary>

You can switch to directly use a `hostPath` volume for the domain socket.
This is done by changing the `shortCircuit.volumeType` field to `hostPath`.
Note that you also need to define the path to use for the `hostPath` volume.

```properties
shortCircuit:
  enabled: true
  policy: uuid
  size: 1Mi
  # volumeType controls the type of shortCircuit volume.
  # It can be "persistentVolumeClaim" or "hostPath"
  volumeType: hostPath
  # Attributes to use if the domain socket volume is hostPath
  hostPath: "/tmp/alluxio-domain" # The hostPath directory to use
```
</details>
<details><summary>kubectl</summary>

You should verify the properties in `ALLUXIO_WORKER_JAVA_OPTS` in the same way as using `PersistentVolumeClaim`.

Also you should make sure the worker Pods have domain socket defined in the `volumes`,
and all relevant containers have the domain socket volume mounted.
The domain socket volume is defined as below by default:
```properties
volumes:
  - name: alluxio-domain
    hostPath:
      path: /tmp/alluxio-domain
      type: DirectoryOrCreate
```

> Note: Compute application containers **MUST** mount the domain socket volume to the same path
(`/opt/domain`) as configured for the Alluxio workers.

</details>

### Verify Short-circuit Operations

To verify short-circuit reads and writes monitor the metrics displayed under:
1. the metrics tab of the web UI as `Domain Socket Alluxio Read` and `Domain Socket Alluxio Write`
1. or, the [metrics json](../operation/Metrics-System.md) as
`cluster.BytesReadDomain` and `cluster.BytesWrittenDomain`
1. or, the [fsadmin metrics CLI](../operation/Admin-CLI.md) as
`Short-circuit Read (Domain Socket)` and `Alluxio Write (Domain Socket)`

### Disable Short-Circuit Operations
To disable short-circuit operations, the operation depends on how you deploy Alluxio.

> Note: As mentioned, disabling short-circuit access for Alluxio workers will result in
worse I/O throughput


<details><summary>helm</summary>

You can disable short circuit by setting the properties as below:

```properties
shortCircuit:
  enabled: false
```

</details>
<details><summary>kubectl</summary>

You should set the property `alluxio.user.short.circuit.enabled` to `false` in your
`ALLUXIO_WORKER_JAVA_OPTS`.
```properties
-Dalluxio.user.short.circuit.enabled=false
```

You should also manually remove the volume `alluxio-domain` from `volumes` of the Pod definition
and `volumeMounts` of each container if existing.

</details>


### Enable remote logging

Alluxio supports a centralized log server that collects logs for all Alluxio processes. 
You can find the specific section at [Remote logging](../administration/Remote-Logging.md).
This can be enabled on K8s too, so that all Alluxio pods will send logs to this log server.


<details><summary>helm</summary>

**Step 1: Configure the log server**

By default, the Alluxio remote log server is not started.
You can enable the log server by configuring the following properties:
```properties
logserver:
  enabled: true
```

If you are just testing and it is okay to discard logs, you can use an `emptyDir` to store the logs in the log server.
```properties
logserver:
  enabled: true
  # volumeType controls the type of log volume.
  # It can be "persistentVolumeClaim" or "hostPath" or "emptyDir"
  volumeType: emptyDir
  # Attributes to use when the log volume is emptyDir
  medium: ""
  size: 4Gi
```

For a production environment, you should always persist the logs with a Persistent Volume.
When you specify the `logserver.volumeType` to be `persistentVolumeClaim`, 
the Helm Chart will create a PVC.
If you are not using dynamic provisioning for PVs, you will need to manually create the PV.
Remember to make sure the selectors for PVC and PV match with each other.
```properties
logserver:
  enabled: true
  # volumeType controls the type of log volume.
  # It can be "persistentVolumeClaim" or "hostPath" or "emptyDir"
  volumeType: persistentVolumeClaim
  # Attributes to use if the log volume is PVC
  pvcName: alluxio-logserver-logs
  accessModes:
    - ReadWriteOnce
  storageClass: standard
  # If you are dynamically provisioning PVs, the selector on the PVC should be empty.
  # Ref: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#class-1
  selector: {}
  # If you are manually allocating PV for the logserver,
  # it is recommended to use selectors to make sure the PV and PVC match as expected.
  # You can specify selectors like below:
  # Example:
  # selector:
  #   matchLabels:
  #     role: alluxio-logserver
  #     app: alluxio
  #     chart: alluxio-<chart version>
  #     release: alluxio
  #     heritage: Helm
  #     dc: data-center-1
  #     region: us-east
```

**Step 2: Helm install with the updated configuration**

When you enable the remote log server, it will be managed by a K8s Deployment.
If you specify the volume type to be `persistentVolumeClaim`, a PVC will be created and mounted.
You will need to provision a PV for the PVC.
Then there will be a Service created for the Deployment, which all other Alluxio pods send logs to.

</details>
<details><summary>kubectl</summary>

**Step 1: Configure log server location with environment variables**

Add `ALLUXIO_LOGSERVER_HOSTNAME` and `ALLUXIO_LOGSERVER_PORT` properties to the configmap.
```properties
apiVersion: v1
kind: ConfigMap
metadata:
  ..omitted
data:
  ..omitted
  ALLUXIO_LOGSERVER_HOSTNAME: alluxio-logserver
  ALLUXIO_LOGSERVER_PORT: "45600"
```
> Note: The value for `ALLUXIO_LOGSERVER_PORT` must be a string or kubectl will fail to read it.

**Step 2: Configure and start log server**

In the sample YAML directory (e.g. `singleMaster-localJournal`), the `logserver/` directory
contains all resources for the log server, including a Deployment, a Service and a PVC if needed.

First you can prepare the YAML file and configure what volume to use for the Deployment.
```console
$ cp logserver/alluxio-logserver-deployment.yaml.template logserver/alluxio-logserver-deployment.yaml
```

If you are testing and it is okay to discard logs, you can use an `emptyDir` for the volume like below:
```properties
  volumes:      
  - name: alluxio-logs
    emptyDir:
      medium: 
      sizeLimit: "4Gi"
``` 

And the volume should be mounted to the log server container at `/opt/alluxio/logs`. 
```properties
  volumeMounts:
  - name: alluxio-logs
    mountPath: /opt/alluxio/logs
```

For a production environment, you should always persist the logs with a Persistent Volume.
```properties
  volumes:      
  - name: alluxio-logs
    persistentVolumeClaim:
      claimName: "alluxio-logserver-logs"
```

There is also a YAML template for PVC `alluxio-logserver-logs`.
```console
$ cp logserver/alluxio-logserver-pvc.yaml.template logserver/alluxio-logserver-pvc.yaml
```

You can further configure the resource and selector for the PVC, according to your environment.
```properties
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: alluxio-logserver-logs
  ..omitted
spec:
  volumeMode: Filesystem
  resources:
    requests:
      storage: 4Gi
  storageClassName: standard
  accessModes:
    - ReadWriteOnce
  # If you are using dynamic provisioning, leave the selector empty.
  selector: {}
  # If you are manually allocating PV for the logserver,
  # it is recommended to use selectors to make sure the PV and PVC match as expected.
  # You can specify selectors like below:
  # Example:
  # selector:
  #   matchLabels:
  #     role: alluxio-logserver
  #     app: alluxio
  #     chart: alluxio-<chart version>
  #     release: alluxio
  #     heritage: Helm
  #     dc: data-center-1
  #     region: us-east
```

Create the PVC when you are ready.
```console
$ kubectl create -f alluxio-logserver-pvc.yaml
```

(Optional) If you are not using dynamic provisioning, you need to prepare the PV yourself.
Remember to make sure the selectors on the PVC and PV match with each other.

After you configure the volume in the Deployment, you can go ahead to create it.
```console
$ kubectl create -f alluxio-logserver-deployment.yaml
```

There is also a Service associated to the Deployment.
```console
$ cp logserver/alluxio-logserver-service.yaml.template logserver/alluxio-logserver-service.yaml
$ kubectl create -f logserver/alluxio-logserver-service.yaml
```

**Step 3: Restart other Alluxio pods**

You need to restart your other Alluxio pods (masters, workers, FUSE etc) so they
capture the updated environment variables and send logs to the remote log server.

</details>

**Verify log server**

You can go into the log server pod and verify the logs exist.

```console
$ kubectl exec -it <logserver-pod-name> -- bash
# In the logserver pod
bash-4.4$ pwd
/opt/alluxio
# You should see logs collected from other Alluxio pods
bash-4.4$ ls -al logs
total 16
drwxrwsr-x    4 1001     bin           4096 Jan 12 03:14 .
drwxr-xr-x    1 alluxio  alluxio         18 Jan 12 02:38 ..
drwxr-sr-x    2 alluxio  bin           4096 Jan 12 03:14 job_master
-rw-r--r--    1 alluxio  bin            600 Jan 12 03:14 logserver.log
drwxr-sr-x    2 alluxio  bin           4096 Jan 12 03:14 master
drwxr-sr-x    2 alluxio  bin           4096 Jan 12 03:14 worker
drwxr-sr-x    2 alluxio  bin           4096 Jan 12 03:14 job_worker
```

### POSIX API

Once Alluxio is deployed on Kubernetes, there are multiple ways in which a client application can
connect to it. For applications using the [POSIX API](../api/POSIX-API.md),
application containers can simply mount the Alluxio FileSystem.

#### FUSE daemon

One way to use the POSIX API is to deploy the Alluxio FUSE daemon, creating pods running Alluxio Fuse processes 
at deployment time. The Fuse processes are long-running.

<details><summary>helm</summary>

You can deploy the FUSE daemon by configuring the following properties:
```properties
fuse:
  enabled: true
```

To modify the default Fuse mount configuration, one can set
- `mountPath`: The container path to be mounted. Default to `/mnt/alluxio-fuse`
- `alluxioPath`: The alluxio path to be mounted to container `mountPath`. Default to `/`
- `mountOptions`: The Fuse mount options. Default to `allow_other`.
See [Fuse mount options](../api/POSIX-API.md#fuse-mount-options) for more details.

```properties
fuse:
  enabled: true
  clientEnabled: true
  mountPath: /mnt/alluxio-fuse
  alluxioPath: /
  mountOptions: allow_other
```

Then follow the steps to install Alluxio with helm [here](#deploy).

If Alluxio has already been deployed with helm and now you want to enable FUSE, you use
`helm upgrade` to add the FUSE daemons.
```console
$ helm upgrade alluxio -f config.yaml \
  --set fuse.enabled=true \
  alluxio-charts/alluxio
```

##### Advanced POSIX API Configuration

- Alluxio fuse/client configuration:
```properties
properties:
  alluxio.user.metadata.cache.enabled: true
  alluxio.user.metadata.cache.expiration.time: 2day
  alluxio.user.metadata.cache.max.size: "1000000"
  alluxio.user.direct.memory.io.enabled: true
  alluxio.fuse.logging.threshold: 1000ms
```
- Alluxio fuse java opts
```properties
fuse:
  jvmOptions: " -Xmx16G -Xms16G -XX:MaxDirectMemorySize=32g"
```
- Alluxio fuse mount options
```properties
fuse:
  mountOptions: direct_io,ro,max_read=131072,attr_timeout=7200,entry_timeout=7200
```
- Alluxio fuse environment variables
```properties
fuse:
  env:
    MAX_IDLE_THREADS: "64"
```

[POSIX API docs](../api/POSIX-API.md) provides more details about how to configure Alluxio POSIX API.

</details>
<details><summary>kubectl</summary>

```console
$ cp alluxio-fuse.yaml.template alluxio-fuse.yaml
$ kubectl create -f alluxio-fuse.yaml
```
Note:
- The container running the Alluxio FUSE daemon must have the `securityContext.privileged=true` with
`SYS_ADMIN` capabilities.
Application containers that require Alluxio access do not need this privilege.

- Application containers can run on any Docker image.

Then data can then be accessed inside the application container under `/mnt/alluxio-fuse`.

##### Advanced POSIX API Configuration

- Alluxio fuse/client java opts can be set in `alluxio-configmap.yaml`:
```yaml
  ALLUXIO_FUSE_JAVA_OPTS: |-
    -Dalluxio.fuse.mount.point=/mnt/alluxio-fuse 
    -Dalluxio.fuse.mount.alluxio.path=/ 
    -Dalluxio.fuse.mount.options=direct_io,max_read=131072,entry_timeout=7200,attr_timeout=7200 
    -Dalluxio.user.hostname=${ALLUXIO_CLIENT_HOSTNAME} 
    -Dalluxio.user.metadata.cache.enabled=true 
    -Dalluxio.user.metadata.cache.expiration.time=40min 
    -Dalluxio.user.metadata.cache.max.size=10000000 
    -Dalluxio.user.logging.threshold=1000ms 
    -Dalluxio.fuse.logging.threshold=1000ms 
```
Note that if Alluxio Worker and Alluxio Fuse is co-located in the same node, Alluxio fuse
can read from the worker storage directly to improve read performance. 
In this case, Alluxio Fuse need to know about the worker storage information.
This is why worker storage configuration is set in `ALLUXIO_JAVA_OPTS` shared by all Alluxio containers.
- Alluxio fuse environment variables can be set in `fuse/alluxio-fuse.yaml`:
```yaml
containers:
  - name: alluxio-fuse
    env:
      - name: "MAX_IDLE_THREADS"
        value: "64"
```

[POSIX API docs](../api/POSIX-API.md) provides more details about how to configure Alluxio POSIX API.
</details>

To access data in Alluxio inside application containers, simply mount Alluxio with a `hostPath` mount of location `/mnt/alluxio-fuse`.

<details><summary>kubectl</summary>
{%Example %}

Below is a sample nginx pod that is able to access data from Alluxio under `/mnt/alluxio-fuse` inside the pod.
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  containers:
  - image: nginx
    name: nginx
    ports:
    - containerPort: 80
      protocol: TCP
    volumeMounts:
    - name: alluxio-fuse-mount
      mountPath: /mnt/alluxio-fuse 
  volumes:
  - name: alluxio-fuse-mount
    hostPath:
      path: /mnt/alluxio-fuse
      type: Directory
```
</details>

#### CSI
Other than using Alluxio FUSE daemon, you could also use CSI to mount the Alluxio FileSystem into application containers.
Unlike Fuse daemon which is a long-running process, the Fuse pod launched by CSI has the same life cycle as the 
application pods who mount Alluxio as a volume. Fuse pod is automatically launched when an application pod mounts Alluxio
inside itself, and automatically terminated when such application pods are terminated.

In order to use CSI, you need a Kubernetes cluster with version at least 1.17, 
with [RBAC](https://kubernetes.io/docs/reference/access-authn-authz/rbac/) enabled in API Server.

**Step 1: Customize configurations**

You can either use the default CSI configurations provided in
[here](https://github.com/Alluxio/alluxio/blob/master/integration/kubernetes/helm-chart/alluxio/values.yaml)
under the csi section, or you can customize them to make it suitable for your workload.
Here are some common properties that you can customize:
<table class="table table-striped">
  <tr>
    <th>property name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>alluxioPath</td>
    <td>The path in Alluxio which will be mounted</td>
  </tr>
  <tr>
    <td>mountInPod</td>
    <td>Set to true to launch Fuse process in an alluxio-fuse pod. Otherwise in the same container as nodeserver</td>
  </tr>
  <tr>
    <td>mountPath</td>
    <td>The path that Alluxio will be mounted to in the application container</td>
  </tr>
  <tr>
    <td>mountOptions</td>
    <td>Alluxio Fuse mount options</td>
  </tr>
</table>

**Step 2: Deploy CSI services**
You can use Helm to start the Alluxio CSI components with Alluxio cluster, or `kubectl` to create the resources manually,
or parts from Helm and parts manually.


<details><summary>helm</summary>

To start the CSI components via helm chart, set the following values in your helm configuration file:
```properties
csi:
  enabled: true
  clientEnabled: true
```

Related Alluxio CSI-related services will be started along with Alluxio cluster.

</details>
<details><summary>kubectl</summary>

Modify or add any configuration properties inside `values.yaml`, 
then please use `helm-generate.sh` (see [here](https://github.com/Alluxio/alluxio/tree/master/integration/kubernetes#generate-kubectl-yaml-templates-from-helm-chart) for usage)
to generate related templates. All CSI related templates will be under `${ALLUXIO_HOME}/integration/kubernetes/csi`.

```console
$ mv alluxio-csi-controller-rbac.yaml.template alluxio-csi-controller-rbac.yaml
$ mv alluxio-csi-controller.yaml.template alluxio-csi-controller.yaml
$ mv alluxio-csi-driver.yaml.template alluxio-csi-driver.yaml
$ mv alluxio-csi-fuse-configmap.yaml.template alluxio-csi-fuse-configmap.yaml
$ mv alluxio-csi-nodeplugin.yaml.template alluxio-csi-nodeplugin.yaml
```
Then run
```console
$ kubectl apply -f alluxio-csi-controller-rbac.yaml -f alluxio-csi-controller.yaml -f alluxio-csi-driver.yaml -f alluxio-csi-fuse-configmap.yaml -f alluxio-csi-nodeplugin.yaml
```
to deploy CSI-related services.

</details>

**Step 3: Provisioning**

We provide both templates for k8s dynamic provisioning and static provisioning.
Please choose the suitable provisioning methods according to your use case.
You can refer to [Persistent Volumes | Kubernetes](https://www.google.com/url?q=https://kubernetes.io/docs/concepts/storage/persistent-volumes/)
and [Dynamic Volume Provisioning | Kubernetes](https://kubernetes.io/docs/concepts/storage/dynamic-provisioning/) to get more details.

<details><summary>helm</summary>

Both dynamic provisioning and static provisioning resources are created via Helm chart. If you need additional resources, you need to create them
manaully through `kubectl`.

</details>
<details><summary>kubectl</summary>

##### Persistent Volumes

For static provisioning, we generate two template files: `alluxio-pv.yaml.template` and `alluxio-pvc-static.yaml.template`.
You can modify these two files based on your needs, then create the respective yaml files.
```console
$ mv alluxio-pv.yaml.template alluxio-pv.yaml
$ mv alluxio-pvc-static.yaml.template alluxio-pvc-static.yaml
```
Then run
```console
$ kubectl apply -f alluxio-pv.yaml
$ kubectl apply -f alluxio-pvc-static.yaml
```
to deploy the resources.

Note: If `mountInPod` is set to `true`, in `alluxio-pv.yaml`, the value of `spec.csi.volumeHandle`
needs to be unique for CSI to identify different volumes. If the values of `volumeHundle` of two
PVs are the same, CSI would regard them as the same volume, and thus may not launch Fuse pod,
affecting the business pods.

##### Dynamic Volume Provisioning

For dynamic provisioning, we generate two template files: `alluxio-storage-class.yaml.template` and `alluxio-pvc.yaml.template`.
You can modify these two files based on your needs, then create the respective yaml files.
```console
$ mv alluxio-storage-class.yaml.template alluxio-storage-class.yaml
$ mv alluxio-pvc.yaml.template alluxio-pvc.yaml
```
Then run
```console
$ kubectl apply -f alluxio-storage-class.yaml
$ kubectl apply -f alluxio-pvc.yaml
```
to deploy the resources.

</details>

**Step 4: Deploy applications**

Now you can put the PVC name in your application pod spec to use the Alluxio FileSystem.
<details><summary>Example</summary>

Below is a sample nginx pod that is able to access data from Alluxio under `/data` inside the pod.
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  containers:
    - image: nginx
      imagePullPolicy: Always
      name: nginx
      ports:
        - containerPort: 80
          protocol: TCP
      volumeMounts:
        - mountPath: /data
          name: alluxio-pvc
  volumes:
    - name: alluxio-pvc
      persistentVolumeClaim:
        claimName: alluxio-pvc

```
</details>

For more information on how to configure a pod to use a persistent volume for storage in Kubernetes,
please refer to [here](https://kubernetes.io/docs/tasks/configure-pod-container/configure-persistent-volume-storage/).

### Start Alluxio Proxy server

One can use either `helm` or `kubectl`  to set up Alluxio proxy servers inside a kubernetes cluster.

<details><summary>kubectl</summary>
{% navtab helm %}

By default, proxy uses daemonset, so every node would spawn a pod running proxy server. 
To start proxy server when deploying Alluxio, set the following property in the helm chart configuration file:

```properties
proxy:
  enabled: true
```

</details>
<details><summary>kubectl</summary>

#### Configuration
In the sample YAML directory (e.g. `singleMaster-localJournal`), the `proxy/` directory
contains the daemonset configuration file for the proxy. Users can modify the configurations
according to the needs.

#### Deploy proxy server
Run the following commands to deploy proxy daemonset:

```console
$ cp alluxio-proxy-daemonset.yaml.template alluxio-proxy-daemonset.yaml
$ kubectl create -f alluxio-proxy-daemonset.yaml
```

#### Stop proxy server
Run the following command to stop proxy daemonset:
```console
$ kubectl delete daemonset alluxio-proxy
```

</details>

### Toggle Master or Worker in Helm chart
In use cases where you wish to install Alluxio masters and workers separately
with the Helm chart, use the following respective toggles:

```properties
master:
  enabled: false

worker:
  enabled: false
```

### Kubernetes Configuration Options

The following options are provided in our Helm chart as additional
parameters for experienced Kubernetes users.

#### ServiceAccounts

[By default](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-the-default-service-account-to-access-the-api-server)
Kubernetes will assign the namespace's `default` ServiceAccount
to new pods in a namespace. You may specify for Alluxio pods to use
any existing ServiceAccounts you may have in your cluster through
the following:

<details><summary>helm</summary>

You may specify a top-level Helm value `serviceAccount` which will
apply to the Master, Worker, and FUSE pods in the chart.
```properties
serviceAccount: sa-alluxio
```

You can override the top-level Helm value by specifying a value
for the specific component's `serviceAccount` like below:
```properties
master:
  serviceAccount: sa-alluxio-master

worker:
  serviceAccount: sa-alluxio-worker
```

</details>
<details><summary>kubectl</summary>

You may add a `serviceAccountName` field to any of the Alluxio Pod template
specs to have the Pod run using the matching ServiceAccount. For example:
```properties
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: alluxio-master
spec:
  template:
    spec:
      serviceAccountName: sa-alluxio
```

</details>

#### Node Selectors & Tolerations

Kubernetes provides many options to control the scheduling of pods
onto nodes in the cluster. The most direct of which is a
[node selector](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector).

However, Kubernetes will avoid scheduling pods on any tainted nodes.
To allow certain pods to schedule on such nodes, Kubernetes allows
you to specify tolerations for those taints. See
[the Kubernetes documentation on taints and tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/)
for more details.

<details><summary>helm</summary>

You may specify a node selector in JSON as a top-level Helm value,
`nodeSelector`, which will apply to all pods in the chart. Similarly,
you may specify a list of tolerations in JSON as a top-level Helm value,
`tolerations`, which will also apply to all pods in the chart.
```properties
nodeSelector: {"app": "alluxio"}

tolerations: [ {"key": "env", "operator": "Equal", "value": "prod", "effect": "NoSchedule"} ]
```

You can **override** the top-level `nodeSelector` by specifying a value
for the specific component's `nodeSelector`.
```properties
master:
  nodeSelector: {"app": "alluxio-master"}

worker:
  nodeSelector: {"app": "alluxio-worker"}
```

You can **append** to the top-level `tolerations` by specifying a value
for the specific component's `tolerations`.
```properties
logserver:
  tolerations: [ {"key": "app", "operator": "Equal", "value": "logging", "effect": "NoSchedule"} ]
```

</details>
<details><summary>kubectl</summary>

You may add `nodeSelector` and `tolerations` fields to any of the Alluxio Pod template
specs. For example:
```properties
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: alluxio-master
spec:
  template:
    spec:
      nodeSelector:
        app: alluxio
      tolerations:
        - effect: NoSchedule
          key: env
          operator: Equal
          value: prod
```

</details>

#### Host Aliases

If you wish to add or override hostname resolution in the pods,
Kubernetes exposes the containers' `/etc/hosts` file via
[host aliases](https://kubernetes.io/docs/concepts/services-networking/add-entries-to-pod-etc-hosts-with-host-aliases/).
This can be particularly useful for providing hostname addresses
for services not managed by Kubernetes, like HDFS.

<details><summary>helm</summary>

You may specify a top-level Helm value `hostAliases` which will
apply to the Master and Worker pods in the chart.
```properties
hostAliases:
- ip: "127.0.0.1"
  hostnames:
    - "foo.local"
    - "bar.local"
- ip: "10.1.2.3"
  hostnames:
    - "foo.remote"
    - "bar.remote"
```

</details>
<details><summary>kubectl</summary>

You may add the `hostAliases` field to any of the Alluxio Pod template
specs. For example:
```properties
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: alluxio-master
spec:
  template:
    spec:
      hostAliases:
      - ip: "127.0.0.1"
        hostnames:
          - "foo.local"
          - "bar.local"
      - ip: "10.1.2.3"
        hostnames:
          - "foo.remote"
          - "bar.local"
```

</details>

#### Deployment Strategy

[By default](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#strategy)
Kubernetes will use the 'RollingUpdate' deployment strategy to progressively
upgrade Pods when changes are detected.

<details><summary>helm</summary>

The Helm chart currently only supports `strategy` for the logging server deployment:
```properties
logserver:
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 25%
      maxSurge: 1
```

</details>
<details><summary>kubectl</summary>

You may add a `strategy` field to any of the Alluxio Pod template
specs to have the Pod run using the matching ServiceAccount. For example:
```properties
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: alluxio-master
spec:
  template:
    spec:
      strategy:
        type: Recreate
```

</details>


#### ImagePullSecrets

Kubernetes supports [accessing images from a Private Registry](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/).
After creating the registry credentials `Secret` in Kubernetes, you pass the secret
to your Pods via `imagePullSecrets`.

<details><summary>helm</summary>

The following value applies the specified `imagePullSecrets` to all
Pods in the Helm chart.

```properties
imagePullSecrets:
  - ecr
  - dev
```

</details>
<details><summary>kubectl</summary>

Add `imagePullSecrets` to your Pod specs. Eg:

```properties
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: alluxio-master
spec:
  template:
    spec:
      containers:
      - name: alluxio-master
        image: private-registry/alluxio:2.9.5
      imagePullSecrets:
      - name: ecr
      - name: dev
```

</details>

## Troubleshooting

<details><summary>Worker Host Unreachable</summary>

Alluxio workers use host networking with the physical host IP as the hostname. Check the cluster
firewall if an error such as the following is encountered:
```
Caused by: io.netty.channel.AbstractChannel$AnnotatedConnectException: finishConnect(..) failed: Host is unreachable: <host>/<IP>:29999
```

- Check that `<host>` matches the physical host address and is not a virtual container hostname.
Ping from a remote client to check the address is resolvable.
```console
$ ping <host>
```
- Verify that a client can connect to the workers on the ports specified in the worker
deployment specification. The default ports are `[29998, 29999, 29996, 30001, 30002, 30003]`.
Check access to the given port from a remote client using a network utility such as `ncat`:
```console
$ nc -zv <IP> 29999
```
</details>
<details><summary>Permission Denied</summary>

From Alluxio v2.1 on, Alluxio Docker containers will run as non-root user `alluxio` with
UID 1000 and GID 1000 by default.
Kubernetes [`hostPath`](https://kubernetes.io/docs/concepts/storage/volumes/#hostpath) volumes
are only writable by root so you need to update the permission accordingly.

</details>
<details><summary>Enable Debug Logging</summary>

To change the log level for Alluxio servers (master and workers), use the CLI command `logLevel` as
follows:

Access the Alluxio CLI from the master Pod.
```console
$ kubectl exec -ti alluxio-master-0 -- /bin/bash
```

From the master Pod, execute the following:
```console
$ alluxio logLevel --level DEBUG --logName alluxio
```
</details>
<details><summary>Accessing Logs</summary>

The Alluxio master and job master run as separate containers of the master Pod. Similarly, the
Alluxio worker and job worker run as separate containers of a worker Pod. Logs can be accessed for
the individual containers as follows.

Master:
```console
$ kubectl logs -f alluxio-master-0 -c alluxio-master
```

Worker:
```console
$ kubectl logs -f alluxio-worker-<id> -c alluxio-worker
```

Job Master:
```console
$ kubectl logs -f alluxio-master-0 -c alluxio-job-master
```

Job Worker:
```console
$ kubectl logs -f alluxio-worker-<id> -c alluxio-job-worker
```
</details>
<details><summary>POSIX API</summary>

In order for an application container to mount the `hostPath` volume, the node running the container
must have the Alluxio FUSE daemon running. The default spec `alluxio-fuse.yaml` runs as a DaemonSet,
launching an Alluxio FUSE daemon on each node of the cluster.

If there are issues accessing Alluxio using the POSIX API:
1. Identify the node the application container ran on using the command
`kubectl describe pods` or the dashboard.
1. Use the command `kubectl describe nodes <node>` to identify the `alluxio-fuse` Pod running on
that node.
1. Tail logs for the identified Pod to view any errors encountered:
`kubectl logs -f alluxio-fuse-<id>`.
</details>
<details><summary>Filename too long</summary>

Alluxio workers create a domain socket used for short-circuit access by default.
On Mac OS X, Alluxio workers may fail to start if the location for this domain socket is a path
which is longer than what the filesystem accepts.
```
2020-07-27 21:39:06,030 ERROR GrpcDataServer - Alluxio worker gRPC server failed to start on /opt/domain/1d6d7c85-dee0-4ac5-bbd1-86eb496a2a50
java.io.IOException: Failed to bind
	at io.grpc.netty.NettyServer.start(NettyServer.java:252)
	at io.grpc.internal.ServerImpl.start(ServerImpl.java:184)
	at io.grpc.internal.ServerImpl.start(ServerImpl.java:90)
	at alluxio.grpc.GrpcServer.lambda$start$0(GrpcServer.java:77)
	at alluxio.retry.RetryUtils.retry(RetryUtils.java:39)
	at alluxio.grpc.GrpcServer.start(GrpcServer.java:77)
	at alluxio.worker.grpc.GrpcDataServer.<init>(GrpcDataServer.java:107)
	at sun.reflect.NativeConstructorAccessorImpl.newInstance0(Native Method)
	at sun.reflect.NativeConstructorAccessorImpl.newInstance(NativeConstructorAccessorImpl.java:62)
	at sun.reflect.DelegatingConstructorAccessorImpl.newInstance(DelegatingConstructorAccessorImpl.java:45)
	at java.lang.reflect.Constructor.newInstance(Constructor.java:423)
	at alluxio.util.CommonUtils.createNewClassInstance(CommonUtils.java:273)
	at alluxio.worker.DataServer$Factory.create(DataServer.java:47)
	at alluxio.worker.AlluxioWorkerProcess.<init>(AlluxioWorkerProcess.java:162)
	at alluxio.worker.WorkerProcess$Factory.create(WorkerProcess.java:46)
	at alluxio.worker.WorkerProcess$Factory.create(WorkerProcess.java:38)
	at alluxio.worker.AlluxioWorker.main(AlluxioWorker.java:72)
Caused by: io.netty.channel.unix.Errors$NativeIoException: bind(..) failed: Filename too long
```

If this is the case, set the following properties to limit the path length:
- `alluxio.worker.data.server.domain.socket.as.uuid=false`
- `alluxio.worker.data.server.domain.socket.address=/opt/domain/d`

> Note: You may see performance degradation due to lack of node locality.

</details>
<details><summary>Worker Pods get OOMKilled by the Kubernetes scheduler</summary>

This is most likely caused due to the Kubernetes configured
[Pod resource limits](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#requests-and-limits)
having the `limits.memory` set too low.

Firstly, double check the configured values for your Alluxio worker Pod `limits.memory`.
**Note that the Pod consists of two containers, each with their own resource limits.**

Check the configured resource requests and limits using `kubectl describe pod`,
`kubectl get pod`, or equivalent Kube API requests. eg.,

```
$ kubectl get po -o json alluxio-worker-xxxxx | jq '.spec.containers[].resources'
{
  "limits": {
    "cpu": "4",
    "memory": "4G"
  },
  "requests": {
    "cpu": "1",
    "memory": "2G"
  }
}
{
  "limits": {
    "cpu": "4",
    "memory": "4G"
  },
  "requests": {
    "cpu": "1",
    "memory": "1G"
  }
}
```

If you used the Helm chart,
[the default values](https://github.com/Alluxio/alluxio/blob/master/integration/kubernetes/helm-chart/alluxio/values.yaml)
are:

```
worker:
  resources:
    limits:
      cpu: "4"
      memory: "4G"
    requests:
      cpu: "1"
      memory: "2G"

jobWorker:
  resources:
    limits:
      cpu: "4"
      memory: "4G"
    requests:
      cpu: "1"
      memory: "1G"
```

- Even if you did not configure any values with Helm, you may still have resource limits in
  place due to a [LimitRange](https://kubernetes.io/docs/concepts/policy/limit-range/)
  applied to your namespace

Next, ensure that the nodes that the Alluxio worker pods are running on have
sufficient resources matching your configured values. You can check that the nodes
you intend to schedule Alluxio worker Pods on have sufficient resources to meet
your requests using `kubectl describe node`, `kubectl get node`, or equivalent Kube API requests. eg.,
```
$ kubectl get no -o json k8sworkernode-0 | jq '.status.allocatable'
{
  "cpu": "8",
  "ephemeral-storage": "123684658586",
  "hugepages-1Gi": "0",
  "hugepages-2Mi": "0",
  "memory": "64886128Ki",
  "pods": "110"
}
```

Isolating Alluxio worker Pods from other Pods in your Kubernetes cluster can be accomplished
with the help of [node selectors](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector)
and [node taints + tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/).
- Keep in mind that the Alluxio worker Pod definition uses a
  [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/),
  so there will be worker Pods assigned to all eligible nodes

Next, verify the Alluxio workers' configured ramdisk sizes (if any).
See [the list of Alluxio configuration properties](../reference/Properties-List.md)
for additional details.
- If you used the Helm chart, the Alluxio site properties are configured using `properties`. eg.,

```
properties:
  alluxio.worker.ramdisk.size: 2G
  alluxio.worker.tieredstore.levels: 1
  alluxio.worker.tieredstore.level0.alias: MEM
  alluxio.worker.tieredstore.level0.dirs.mediumtype: MEM
  alluxio.worker.tieredstore.level0.dirs.path: /dev/shm
  alluxio.worker.tieredstore.level0.dirs.quota: 2G
```

- Otherwise, you can view and modify the site properties in the `alluxio-config` ConfigMap. eg.,

```
$ kubectl get cm -o json alluxio-config | jq '.data.ALLUXIO_WORKER_JAVA_OPTS'
"-Dalluxio.worker.ramdisk.size=2G
-Dalluxio.worker.tieredstore.levels=1
-Dalluxio.worker.tieredstore.level0.alias=MEM
-Dalluxio.worker.tieredstore.level0.dirs.mediumtype=MEM
-Dalluxio.worker.tieredstore.level0.dirs.path=/dev/shm
-Dalluxio.worker.tieredstore.level0.dirs.quota=2G "
```

**NOTE: Our `DaemonSet` uses `emptyDir` volumes as the Alluxio worker Pod's ramdisk in Kubernetes.**
```
spec:
  template:
    spec:
      volumes:
        - name: mem
          emptyDir:
            medium: "Memory"
            sizeLimit: 1G
```

This results in the following nuances:
- `sizeLimit` has no effect on the size of the allocated ramdisk unless
the `SizeMemoryBackedVolumes` feature gate is enabled (enabled by default
as of Kubernetes 1.22).
- As stated in [the Kubernetes emptyDir documentation](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir),
if no size is specified then memory-backed `emptyDir` volumes will have capacity
allocated equal to **half the available memory on the host node**. This capacity
is reflected inside of your containers (for example when running `df -u`). However
if the combined size of your ramdisk and container memory usage exceeds the pod's
`limits.memory` then the Kubernetes scheduler will trigger an `OOMKilled` on that pod.
**This is a very likely overlooked source of memory consumption in Alluxio worker Pods.**

Lastly, verify the Alluxio worker JVM heap and off-heap maximum capacities. These are
configured with the JVM flags `-Xmx`/`-XX:MaxHeapSize` and `-XX:MaxDirectMemorySize` respectively.
- See [the Oracle Java documentation](https://docs.oracle.com/javase/8/docs/technotes/tools/windows/java.html)
for more details.

To adjust those values, you would have to manually update the
`(...)_JAVA_OPTS` environment variables in the `alluxio-config` ConfigMap.
For example:
```
apiVersion: v1
kind: ConfigMap
metadata:
  name: alluxio-config
data:
  ALLUXIO_JAVA_OPTS: |-
    -Xmx2g -Dalluxio.master.hostname=alluxio-master-0 ...
  ALLUXIO_MASTER_JAVA_OPTS: |-
    -Dalluxio.master.hostname=${ALLUXIO_MASTER_HOSTNAME}
  ALLUXIO_JOB_MASTER_JAVA_OPTS: |-
    -Dalluxio.master.hostname=${ALLUXIO_MASTER_HOSTNAME}
  ALLUXIO_WORKER_JAVA_OPTS: |-
    -XX:MaxDirectMemorySize=2g -Dalluxio.worker.hostname=${ALLUXIO_WORKER_HOSTNAME} ...
  ALLUXIO_JOB_WORKER_JAVA_OPTS: |-
    -XX:MaxDirectMemorySize=1g -Dalluxio.worker.hostname=${ALLUXIO_WORKER_HOSTNAME} ...
  ALLUXIO_FUSE_JAVA_OPTS: |-
    -Dalluxio.user.hostname=${ALLUXIO_CLIENT_HOSTNAME} -XX:MaxDirectMemorySize=2g
  ALLUXIO_WORKER_TIEREDSTORE_LEVEL0_DIRS_PATH: /dev/shm
```

Thus to avoid worker Pods running into `OOMKilled` errors,
1. Verify that the nodes your Alluxio worker Pods are scheduled on have
sufficient memory to satisfy all the `limits.memory` specifications assigned.
2. Ensure you have configured `alluxio.worker.ramdisk.size` and
`alluxio.worker.tieredstore.level0.dirs.quota` low enough such that
the memory consumed by the ramdisk combined with the JVM memory options
(`-Xmx`, `-XX:MaxDirectMemorySize`) do not exceed the Pod's `limits.memory`.
It is recommended to allow for some overhead as memory may be consumed
by other processes as well.

**Aside:** There is currently an [open issue](https://github.com/Alluxio/alluxio/issues/12277)
in Alluxio where Alluxio's interpretation of byte sizes differs from
Kubernetes (due to Kubernetes distinguishing between "-bibytes").
This is unlikely to cause `OOMKilled` errors unless you are operating on
very tight memory margins.

</details>
<details><summary>JVM not seeing correct memory limit from cgroup</summary>

It is a known issue that in some early versions of Java 8, the JVM running in a container
will determine its heap size(if not specified with `-Xmx` and `-Xms`) based on 
the memory of the physical host instead of the container.
In that case, the JVM may attempt to use more memory than the container
resource limit and gets killed. You can find more detailed explanations 
[here](https://developers.redhat.com/blog/2017/03/14/java-inside-docker).

Since Java 8u131, some JVM flags can be turned on in order to correctly read the memory from cgroup.
You can refer to our `values.yaml` from our Helm chart template, and uncomment the below options.
These options will be added to the JVM options of all Alluxio containers, including the
masters and workers etc. You can find more detailed explanations
[here](https://www.atamanroman.dev/articles/jvm-memory-settings-container-environment/).
```yaml
# Recommended JVM Heap options for running in Docker
# Ref: https://developers.redhat.com/blog/2017/03/14/java-inside-docker/
# These JVM options are common to all Alluxio services
jvmOptions:
  - "-XX:+UnlockExperimentalVMOptions"
  - "-XX:+UseCGroupMemoryLimitForHeap"
  - "-XX:MaxRAMFraction=2"
```

From Java git 8u191 on, the container support works out-of-the-box.
So you don't need to turn on the flags mentioned above any more.

You should check the Java version in the container you are using to ensure the
correct memory limits are respected. Also it is recommended to go to the 
running container and double check the JVM process is running with the correct memory consumption.
</details>
<details><summary>tmpfs is smaller than the configured size</summary>

In Kubernetes context, g or GB means 1000^3 and gi or GiB means 1024^3. However, in Alluxio context, g or GB means 1024^3.
So when we use g and pass the quota to Alluxio and K8s, K8s grants 1000^3 but Alluxio tries to utilize 1024^3.
For example if it is an emptyDir, then the pod using the emptyDir will be killed for overusing resources.

Therefore, we recommend using Gi whenever possible in helm chart or yaml files to avoid such issue.
</details>
