# Journal Management

Alluxio keeps the history of all metadata related changes, such as creating files or renaming directories,
in edit logs referred to as "journal".
Upon startup, the Alluxio master will replay all the steps recorded in the journal to recover its last saved state.
Also when the leading master falls back to a different master for
[high availability (HA)](../deploy/Running-Alluxio-On-a-HA-Cluster.md) mode,
the new leading master also replays the journal to recover the last state of the leading master.
The purpose of this documentation is to help Alluxio administrators understand and manage the Alluxio journal.

## Embedded Journal vs UFS Journal

There are two types of journals that Alluxio supports, `EMBEDDED` or `UFS`.
The embedded journal stores edit logs on each master's local file system and
coordinates multiple masters in HA mode to access the logs
based on a self-managed consensus protocol;
whereas UFS journal stores edit logs in an external shared UFS storage,
and relies on an external Zookeeper for coordination for HA mode.
Starting from 2.2, the default journal type is `EMBEDDED`.
This can be changed by setting the property "`alluxio.master.journal.type`" to "`UFS`"
instead of "`EMBEDDED`".

To choose between the default Embedded Journal and UFS journal,
here are some aspects to consider:

- **External Dependency:**
Embedded journal does not rely on extra services.
UFS journal requires an external Zookeeper cluster in HA mode to coordinate who is the leading
master writing the journal, and requires a UFS for persistent storage.
If the UFS and Zookeeper clusters are not readily available and stable,
it is recommended to use the embedded journal over the UFS journal.
- **Fault tolerance:**
With `n` masters, using the embedded journal can tolerate only `floor(n/2)` master failures,
compared to `n-1` for UFS journal.
For example, With `3` masters, UFS journal can tolerate `2` failures,
while embedded journal can only tolerate `1`.
However, UFS journal depends on Zookeeper,
which similarly only supports `floor(#zookeeper_nodes / 2)` failures.
- **Journal Storage Type:**
When using a single Alluxio master, UFS journal can be local storage;
when using multiple Alluxio masters for HA mode,
this UFS storage must be shared among masters with reading and writing access.
To get reasonable performance, the UFS journal requires a UFS that supports fast streaming writes,
such as HDFS or NFS. In contrast, S3 is not recommended for the UFS journal.

## Configuring Embedded Journal

### Required configuration

The following configuration must be configured to a local path on the masters. The default
value is local directory `${alluxio.work.dir}/journal`.
```
alluxio.master.journal.folder=/local/path/to/store/journal/files/
```

Set the addresses of all masters in the cluster. The default embedded journal port is `19200`.
This must be set on all Alluxio servers, as well as Alluxio clients.

```
alluxio.master.embedded.journal.addresses=master_hostname_1:19200,master_hostname_2:19200,master_hostname_3:19200
```

### Optional configuration

* `alluxio.master.embedded.journal.port`: The port masters use for embedded journal communication. Default: `19200`.
* `alluxio.master.rpc.port`: The port masters use for RPCs. Default: `19998`.
* `alluxio.master.rpc.addresses`: A list of comma-separated `host:port` RPC addresses where the client should look for masters
when using multiple masters without Zookeeper. This property is not used when Zookeeper is enabled, since Zookeeper already stores the master addresses.
If this is not set, clients will look for masters using the hostnames from `alluxio.master.embedded.journal.addresses`
and the master rpc port (Default:`19998`).

### Advanced configuration

<ul>
    <li><code>alluxio.master.embedded.journal.addresses</code>: A comma-separated list of journal addresses for all masters in the cluster. The format is 'hostname1:port1,hostname2:port2,...'. When left unset, Alluxio uses ${alluxio.master.hostname}:${alluxio.master.embedded.journal.port} by default Default: <code></code></li>
    <li><code>alluxio.master.embedded.journal.catchup.retry.wait</code>: Time for embedded journal leader to wait before retrying a catch up. This is added to avoid excessive retries when server is not ready. Default: <code>1s</code></li>
    <li><code>alluxio.master.embedded.journal.election.timeout.max</code>: The max election timeout for the embedded journal. When a random period between ${alluxio.master.embedded.journal.election.timeout.min} and ${alluxio.master.embedded.journal.election.timeout.max} elapses without a master receiving any messages, the master will attempt to become the primary Election timeout will be waited initially when the cluster is forming. So larger values for election timeout will cause longer start-up time. Smaller values might introduce instability to leadership. Default: <code>20s</code></li>
    <li><code>alluxio.master.embedded.journal.election.timeout.min</code>: The min election timeout for the embedded journal. Default: <code>10s</code></li>
    <li><code>alluxio.master.embedded.journal.entry.size.max</code>: The maximum single journal entry size allowed to be flushed. This value should be smaller than 30MB. Set to a larger value to allow larger journal entries when using the Alluxio Catalog service. Default: <code>10MB</code></li>
    <li><code>alluxio.master.embedded.journal.flush.size.max</code>: The maximum size in bytes of journal entries allowed in concurrent journal flushing (journal IO to standby masters and IO to local disks). Default: <code>160MB</code></li>
    <li><code>alluxio.master.embedded.journal.port</code>: The port to use for embedded journal communication with other masters. Default: <code>19200</code></li>
    <li><code>alluxio.master.embedded.journal.raft.client.request.interval</code>: Base interval for retrying Raft client calls. The retry policy is ExponentialBackoffRetry Default: <code>100ms</code></li>
    <li><code>alluxio.master.embedded.journal.raft.client.request.timeout</code>: Time after which calls made through the Raft client timeout. Default: <code>60sec</code></li>
    <li><code>alluxio.master.embedded.journal.ratis.config</code>: Prefix for Apache Ratis internal configuration options. For example, setting alluxio.master.embedded.journal.ratis.config.raft.server.rpc.request.timeout will set ratis.config.raft.server.rpc.request.timeout on the Ratis service in the Alluxio master. Default: <code></code></li>
    <li><code>alluxio.master.embedded.journal.retry.cache.expiry.time</code>: The time for embedded journal server retry cache to expire. Setting a bigger value allows embedded journal server to cache the responses for a longer time in case of journal writer retries, but will take up more memory in master. Default: <code>60s</code></li>
    <li><code>alluxio.master.embedded.journal.snapshot.replication.chunk.size</code>: The stream chunk size used by masters to replicate snapshots. Default: <code>4MB</code></li>
    <li><code>alluxio.master.embedded.journal.snapshot.replication.compression.level</code>: The zip compression level of sending a snapshot from one master to another. Only applicable when alluxio.master.embedded.journal.snapshot.replication.compression.type is not NO_COMPRESSION. The zip format defines ten levels of compression, ranging from 0 (no compression, but very fast) to 9 (best compression, but slow). Or -1 for the system default compression level. Default: <code>1</code></li>
    <li><code>alluxio.master.embedded.journal.snapshot.replication.compression.type</code>: The type of compression to use when transferring a snapshot from one master to another. Options are NO_COMPRESSION, GZIP, TAR_GZIP Default: <code>NO_COMPRESSION</code></li>
    <li><code>alluxio.master.embedded.journal.transport.max.inbound.message.size</code>: The maximum size of a message that can be sent to the embedded journal server node. Default: <code>100MB</code></li>
    <li><code>alluxio.master.embedded.journal.transport.request.timeout.ms</code>: The duration after which embedded journal masters will timeout messages sent between each other. Lower values might cause leadership instability when the network is slow. Default: <code>5sec</code></li>
    <li><code>alluxio.master.embedded.journal.unsafe.flush.enabled</code>: If true, embedded journal entries will be committed without waiting for the entry to be flushed to disk. This may improve performance of write operations on the Alluxio master if the journal is written to a slow or contested disk. WARNING: enabling this property may result in metadata loss if half or more of the master nodes fail. See Ratis property raft.server.log.unsafe-flush.enabled at https://github.com/apache/ratis/blob/master/ratis-docs/src/site/markdown/configuraions.md. Default: <code>false</code></li>
    <li><code>alluxio.master.embedded.journal.write.timeout</code>: Maximum time to wait for a write/flush on embedded journal. Default: <code>30sec</code></li> 
</ul>

### Configuring the Job service

It is usually best not to set any of these - by default the job master will use the same hostnames as the Alluxio master,
so it is enough to set only `alluxio.master.embedded.journal.addresses`. These properties only need to be set
when the job service runs independently of the rest of the system or using a non-standard port.

* `alluxio.job.master.embedded.journal.port`: the port job masters use for embedded journal communications. Default: `20003`.
* `alluxio.job.master.embedded.journal.addresses`: a comma-separated list of journal addresses for all job masters in the cluster.
The format is `hostname1:port1,hostname2:port2,...`.
* `alluxio.job.master.rpc.addresses`: A list of comma-separated host:port RPC addresses where the client should look for job masters
when using multiple job masters without Zookeeper. This property is not used when Zookeeper is enabled,
since Zookeeper already stores the job master addresses. If this property is not defined, clients will look for job masters using
`[alluxio.master.rpc.addresses]:alluxio.job.master.rpc.port` addresses first, then for
`[alluxio.job.master.embedded.journal.addresses]:alluxio.job.master.rpc.port`.

## Configuring UFS Journal

The most important configuration value to set for the journal is
`alluxio.master.journal.folder`. This must be set to a filesystem folder that is
available to all masters. In single-master mode, use a local filesystem path for simplicity. 
With multiple masters distributed across different machines, the folder must
be in a distributed system where all masters can access it. The journal folder
should be in a filesystem that supports flush such as HDFS or NFS. It is not
recommended to put the journal in an object store like S3. With an object store, every
metadata operation requires a new object to be created, which is
prohibitively slow for most serious use cases.

UFS journal options can be configured using the configuration prefix:

`alluxio.master.journal.ufs.option.<some alluxio property>`

**Configuration examples:**

Use HDFS to store the journal:
```
alluxio.master.journal.folder=hdfs://[namenodeserver]:[namenodeport]/alluxio_journal
alluxio.master.journal.ufs.option.alluxio.underfs.version=2.6
```

Use the local file system to store the journal:
```
alluxio.master.journal.folder=/opt/alluxio/journal
```

## Formatting the journal

Formatting the journal deletes all of its content and restores it to a fresh state.
Before starting Alluxio for the first time, the journal must be formatted.

**Warning**: the following command permanently deletes all Alluxio metadata,
so be careful with this command or backup your journal first (see next section).

```console
$ ./bin/alluxio formatMasters
```

## Backing up the journal

### Manually backing up the journal

Alluxio supports taking journal backups so that Alluxio metadata can be restored
to a previous point in time. Generating a backup causes temporary service
unavailability while the leading master is taking the backup.
(See [Backup delegation on HA cluster](#backup-delegation-on-ha-cluster) section for overcoming this limitation.)

To generate a backup, use the `fsadmin backup` CLI command.
```console
$ ./bin/alluxio fsadmin backup
```

By default, this will write a backup named
`alluxio-backup-YYYY-MM-DD-timestamp.gz` to the `/alluxio_backups` directory of
the root under storage system, e.g. `hdfs://cluster/alluxio_backups`. This default
backup directory can be configured by setting `alluxio.master.backup.directory`

```
alluxio.master.backup.directory=/alluxio/backups
```

See the [backup command documentation](../operation/Admin-CLI.md#backup)
for additional backup options.

### Automatically backing up the journal

Alluxio supports automatically taking leading master metadata snapshots every day at a fixed time
so that Alluxio metadata can be restored to at most one day before.
This functionality is enabled by setting the following property in `${ALLUXIO_HOME}/conf/alluxio-site.properties`:

```
alluxio.master.daily.backup.enabled=true
```

The time to take daily snapshots is defined by `alluxio.master.daily.backup.time`. For example, if
a user specified `alluxio.master.daily.backup.time=05:30`, the Alluxio leading master will back up its metadata
to the `alluxio.master.backup.directory` of the root UFS every day at 5:30am UTC.
We recommend setting the backup time to an off-peak time to avoid interfering with other users of the system.

In the daily backup, the backup directory needs to be an absolute path within the root UFS.
For example, if `alluxio.master.backup.directory=/alluxio_backups`
and `alluxio.master.mount.table.root.ufs=hdfs://192.168.1.1:9000/alluxio/underfs`,
the default backup directory would be `hdfs://192.168.1.1:9000/alluxio_backups`.

The files to retain in the backup directory is limited by `alluxio.master.daily.backup.files.retained`.
Users can set this property to the number of backup files they want to keep in the backup directory.

In addition, upon encountering journal corruption, the master will take a backup of its current state
automatically. This can be disabled by setting `alluxio.master.journal.backup.when.corrupted=false`.

### Backup delegation on HA cluster

Alluxio supports taking backup without causing service unavailability on a HA cluster configuration.
When enabled, Alluxio leading master delegates backups to standby masters in the cluster.
After configuring backup delegation, both manual and scheduled backups will run in delegated mode.
From Alluxio 2.9, backup delegation is by default enabled.

Backup delegation can be configured with the below properties:
- `alluxio.master.backup.delegation.enabled`: Whether to delegate backups to standby masters. Default: `false`.
- `alluxio.master.backup.heartbeat.interval`: Interval at which standby master that is taking the backup will update the leading master with current backup status. Default: `2sec`.

Some advanced properties control the communication between Alluxio masters for coordinating the backup:
- `alluxio.master.backup.transport.timeout`: Communication timeout for messaging between masters for coordinating backup. Default: `30sec`.
- `alluxio.master.backup.connect.interval.min`: Minimum delay between each connection attempt to backup-leader. Default: `1sec`.
- `alluxio.master.backup.connect.interval.max`: Maximum delay between each connection attempt to backup-leader. Default: `30sec`.
- `alluxio.master.backup.abandon.timeout`: Duration after which leader will abandon the backup if it has not received heartbeat from backup-worker. Default: `1min`.

Since it is uncertain which host will take the backup, it is suggested to use shared paths for taking backups with backup delegation.

A backup attempt will fail if delegation fails to find a standby master, thus favoring service availability.
For manual backups, you can pass `--allow-leader` option to allow the leading master to take a backup when there are no standby masters to delegate the backup.

You can also pass `--bypass-delegation` flag to disable delegation altogether.
Disabling backup delegation will cause temporary service unavailability while the leading master is writing a backup.

### Restoring from a backup

To restore the Alluxio system from a journal backup, stop the system, format the
journal, then restart the system, passing the URI of the backup with the `-i`
(import) flag.

```console
$ ./bin/alluxio-stop.sh masters
$ ./bin/alluxio formatMasters
$ ./bin/alluxio-start.sh -i <backup_uri> masters
```

The `<backup_uri>` should be a full URI path that is available to all masters, e.g.
`hdfs://[namenodeserver]:[namenodeport]/alluxio_backups/alluxio-journal-YYYY-MM-DD-timestamp.gz`
If backups to the local disk of the leader master, copy the backup file to the same location in each master
and pass in the local backup file path.

If starting up masters individually, pass the `-i` argument to each one. The master which
becomes leader first will import the journal backup, and the rest will ignore the `-i`.

If the restore succeeds, you should see a log message along the lines of
```
INFO AlluxioMasterProcess - Restored 57 entries from backup
```
in the leading master logs.

### Changing masters

#### Embedded Journal Cluster

When internal leader election is used, Alluxio masters are determined with a quorum. Adding or removing
masters requires keeping this quorum in a consistent state.

##### Adding a new master

To prevent inconsistencies in the cluster configuration across masters, only a single master
should be added to an existing embedded journal cluster at a time.

Below are the steps to add a new master to a live cluster:
* Prepare the new master.
    * New master should contain all existing masters in its embedded journal configuration, complete with its address.
* Start new master.
    * This will introduce the new master to the existing cluster.
    * New master will catch up with cluster's state in the background.
* Update existing masters' configuration with the new master address.
    * This is to make sure existing members will connect the new member directly upon a restart.

Note: Adding to an already shut down cluster still requires adding only a single master at a time.

Note: When adding a master to a single master cluster, you should shut down and update the configuration for the existing master.
Then both masters could be started together.

##### Removing a master

The embedded journal cluster will take notice when a member is not available anymore. Such masters will count
against the failure tolerance of the cluster based on the initial member count. To resize the cluster
after a node failure an explicit action is required.

Please note `-domain` parameter in below commands. This is because embedded journal based leader election is supported
for both regular masters and job service masters. You should supply correct value based on which cluster you intend to work on.

1. Check current quorum state:
```console
$ ./bin/alluxio fsadmin journal quorum info -domain <MASTER | JOB_MASTER>
```
This will print out node status for all currently participating members of the embedded journal cluster. You should verify
that the removed master is shown as `UNAVAILABLE`.

2. Remove member from the quorum:
`-address` option below should reflect the exact address that is returned by the `.. quorum info` command provided above.
```console
$ ./bin/alluxio fsadmin journal quorum remove -domain <MASTER | JOB_MASTER> -address <HOSTNAME:PORT>
```

3. Verify that the removed member is no longer shown in the `quorum info`.

##### Electing a specific master as leader
To aid in debugging and to add flexibility, it is possible to manually change the leader of an embedded journal cluster.

1. Check current quorum state:
```console
$ ./bin/alluxio fsadmin journal quorum info -domain MASTER
```
This will print out node status for all currently participating members of the embedded journal cluster. You should select one 
of the `AVAILABLE` masters. The current leader is also displayed by this command. 

2. Elect an available master as leader:
```console
$ ./bin/alluxio fsadmin journal quorum elect -address <HOSTNAME:PORT>
```
The `elect` command makes sure that the leadership has transferred to the designated master before returning and displaying
a success message. If the transfer is not successful, it will print a failure message.

#### UFS Journal Cluster

To add a master to an HA Alluxio cluster, you can simply start a new Alluxio master process, with
the appropriate configuration. The configuration for the new master should be the same as other masters,
except that the parameter `alluxio.master.hostname=<MASTER_HOSTNAME>` should reflect the new hostname.
Once the new master starts, it will start interacting with ZooKeeper to participate in leader election.

Removing a master is as simple as stopping the master process. If the cluster is a single master cluster,
stopping the master will essentially shutdown the cluster since the single master is down. If the
Alluxio cluster is an HA cluster, stopping the leading master will force ZooKeeper to elect a new leading master
and failover to that new leader. If a standby master stops, the operation of the cluster is
unaffected. Keep in mind, Alluxio masters' high availability depends on the availability of standby
masters. If there are not enough standby masters, the availability of the leading master will be affected.
It is recommended to have at least 3 masters for an HA Alluxio cluster.

## Advanced

### Managing the journal size

When running with a single master, the journal folder size will grow indefinitely
as metadata operations are written to journal log files. To address this, production
deployments should run in HA mode with multiple Alluxio masters. The standby
masters will create checkpoints of the master state and clean up the logs that
were written before the checkpoints. For example, if 3 million Alluxio files were
created and then 2 million were deleted, the journal logs would contain 5 million
total entries. Then if a checkpoint is taken, the checkpoint will contain only the
metadata for the 1 million remaining files, and the original 5 million entries will be deleted.

By default, checkpoints are automatically taken every 2 million entries. This can be configured by
setting `alluxio.master.journal.checkpoint.period.entries` on the masters. Setting
the value lower will reduce the amount of disk space needed by the journal at the
cost of additional work for the standby masters.

When the metadata are stored in RocksDB, Alluxio 2.9 added support to checkpointing with multiple threads.
`alluxio.master.metastore.rocks.parallel.backup=true` will turn on multi-threaded checkpointing and
make the checkpointing a few times faster(depending how many threads are used). 
`alluxio.master.metastore.rocks.parallel.backup.threads` controls how many threads to use.
`alluxio.master.metastore.rocks.parallel.backup.compression.level` specifies the compression level, 
where smaller means bigger file and less CPU consumption, and larger means smaller file and more CPU consumption. 

#### Checkpointing on secondary master

If HA mode is not an option, it is possible to run a master on the same node as a
dedicated standby master. This second master exists only to write checkpoints, and
will not serve client requests if the leading master dies. In this setup, both
masters have similar memory requirements since they both need to hold all Alluxio
metadata in memory. To start a dedicated standby master for writing periodic checkpoints,
run

```console
$ ./bin/alluxio-start.sh secondary_master
```

#### Checkpointing on the leading master

Checkpointing requires a pause in master metadata changes and causes temporary service
unavailability while the leading master is writing a checkpoint. 
This operation may take hours depending on Alluxio's namespace size.
Therefore, Alluxio's leading master will not create checkpoints by default.

Restarting the current leading master to transfer the leadership to another running master periodically
can help avoid leading master journal logs from growing unbounded when Alluxio is running in HA mode.

Starting from version 2.4, Alluxio embedded journal HA mode supports automatically transferring checkpoints from standby masters to the leading master. 
The leading master can use those checkpoints as taken locally to truncate its journal size without causing temporary service unavailability. 
No need to manually transfer leadership anymore.

If HA mode is not an option, the following command can be used to manually trigger the checkpoint:

```console
$ ./bin/alluxio fsadmin journal checkpoint
```

Similar to the `backup` command, `checkpoint` command should
be used on an off-peak time to avoid interfering with other users of the system.

### Recovering from journal issues

The journal is integral to Alluxio's health. If the filesystem storing the journal
loses availability, no metadata operations can be performed on Alluxio. Similarly,
if the journal is accidentally deleted or its storage system becomes corrupted,
Alluxio must be reformatted to recover. To avoid the need for full reformatting,
we recommend [taking regular journal backups](#automatically-backing-up-the-journal)
at a time when the cluster is under low load.
Then if something happens to the journal, you can recover from one of the backups.

By default, if a master encounters corruption when replaying a journal it will automatically
take a backup of the state up to the corrupted entry in the configured backup directory. The master will notice the
corruption when elected leader. The backup directory is configured by `alluxio.master.backup.directory`.
This feature can be disabled by setting `alluxio.master.journal.backup.when.corrupted` to `false`.

### Get a human-readable journal

Alluxio journal is serialized and not human-readable. The following command
read the Alluxio journal and write it to a directory in a human-readable format.

```console
$ ./bin/alluxio readJournal
```

See [here](../operation/User-CLI.md#readjournal) for more detailed usage.


### Exiting upon Demotion

By default, Alluxio will transition masters from primaries to standbys.
During this process the JVM is _not_ shut down at any point.
This occasionally leaves behind resources and may lead to a bloated memory footprint.
To avoid taking up too much memory this, there is a flag which forces a master JVM to exit once it
has been demoted from a primary to a standby.
This moves the responsibility of restarting the process to join the quorum as a standby to a
process supervisor such as a Kubernetes cluster manager or systemd.

To configure this behavior for an Alluxio master, set the following configuration inside of 
`alluxio-site.properties`

```properties
alluxio.master.journal.exit.on.demotion=true
```
