# List of Metrics

There are two types of metrics in Alluxio, cluster-wide aggregated metrics, and per-process detailed metrics.

* Cluster metrics are collected and calculated by the leading master and displayed in the metrics tab of the web UI. 
  These metrics are designed to provide a snapshot of the cluster state and the overall amount of data and metadata served by Alluxio.

* Process metrics are collected by each Alluxio process and exposed in a machine-readable format through any configured sinks.
  Process metrics are highly detailed and are intended to be consumed by third-party monitoring tools.
  Users can then view fine-grained dashboards with time-series graphs of each metric,
  such as data transferred or the number of RPC invocations.

Metrics in Alluxio have the following format for master node metrics:

```
Master.[metricName].[tag1].[tag2]...
```

Metrics in Alluxio have the following format for non-master node metrics:

```
[processType].[metricName].[tag1].[tag2]...[hostName]
```

There is generally an Alluxio metric for every RPC invocation, to Alluxio or to the under store.

Tags are additional pieces of metadata for the metric such as user name or under storage location.
Tags can be used to further filter or aggregate on various characteristics.

## Cluster Metrics

Workers and clients send metrics data to the Alluxio master through heartbeats.
The interval is defined by property `alluxio.master.worker.heartbeat.interval` and `alluxio.user.metrics.heartbeat.interval` respectively.

Bytes metrics are aggregated value from workers or clients. Bytes throughput metrics are calculated on the leading master.
The values of bytes throughput metrics equal to bytes metrics counter value divided by the metrics record time and shown as bytes per minute.

<table class="table table-striped">
  <tr><th>Name</th><th>Type</th><th>Description</th></tr>
  <tr><td><a class="anchor" name="Cluster.ActiveRpcReadCount"></a> Cluster.ActiveRpcReadCount</td><td>COUNTER</td><td>The number of active read-RPCs managed by workers</td></tr>
  <tr><td><a class="anchor" name="Cluster.ActiveRpcWriteCount"></a> Cluster.ActiveRpcWriteCount</td><td>COUNTER</td><td>The number of active write-RPCs managed by workers</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesReadDirect"></a> Cluster.BytesReadDirect</td><td>COUNTER</td><td>Total number of bytes read from all workers without external RPC involved. Data exists in worker storage or is fetched by workers from UFSes. This records data read by worker internal calls (e.g. clients embedded in workers).</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesReadDirectThroughput"></a> Cluster.BytesReadDirectThroughput</td><td>GAUGE</td><td>Total number of bytes read from all workers without external RPC involved. Data exists in worker storage or is fetched by workers from UFSes. This records data read by worker internal calls (e.g. clients embedded in workers).</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesReadDomain"></a> Cluster.BytesReadDomain</td><td>COUNTER</td><td>Total number of bytes read from all works via domain socket</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesReadDomainThroughput"></a> Cluster.BytesReadDomainThroughput</td><td>GAUGE</td><td>Bytes read per minute throughput from all workers via domain socket</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesReadLocal"></a> Cluster.BytesReadLocal</td><td>COUNTER</td><td>Total number of bytes short-circuit read reported by all clients. Each client reads data from the collocated worker data storage directly.</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesReadLocalThroughput"></a> Cluster.BytesReadLocalThroughput</td><td>GAUGE</td><td>Bytes per minute throughput short-circuit read reported by all clients</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesReadPerUfs"></a> Cluster.BytesReadPerUfs</td><td>COUNTER</td><td>Total number of bytes read from a specific UFS by all workers</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesReadRemote"></a> Cluster.BytesReadRemote</td><td>COUNTER</td><td>Total number of bytes read from all workers via network (RPC). Data exists in worker storage or is fetched by workers from UFSes. This does not include short-circuit local reads and domain socket reads</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesReadRemoteThroughput"></a> Cluster.BytesReadRemoteThroughput</td><td>GAUGE</td><td>Bytes read per minute throughput from all workers via network (RPC calls). Data exists in worker storage or is fetched by workers from UFSes. This does not include short-circuit local reads and domain socket reads</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesReadUfsAll"></a> Cluster.BytesReadUfsAll</td><td>COUNTER</td><td>Total number of bytes read from all Alluxio UFSes by all workers</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesReadUfsThroughput"></a> Cluster.BytesReadUfsThroughput</td><td>GAUGE</td><td>Bytes read per minute throughput from all Alluxio UFSes by all workers</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesWrittenDomain"></a> Cluster.BytesWrittenDomain</td><td>COUNTER</td><td>Total number of bytes written to all workers via domain socket</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesWrittenDomainThroughput"></a> Cluster.BytesWrittenDomainThroughput</td><td>GAUGE</td><td>Throughput of bytes written per minute to all workers via domain socket</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesWrittenLocal"></a> Cluster.BytesWrittenLocal</td><td>COUNTER</td><td>Total number of bytes short-circuit written to local worker data storage by all clients</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesWrittenLocalThroughput"></a> Cluster.BytesWrittenLocalThroughput</td><td>GAUGE</td><td>Bytes per minute throughput written to local worker data storage by all clients</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesWrittenPerUfs"></a> Cluster.BytesWrittenPerUfs</td><td>COUNTER</td><td>Total number of bytes written to a specific Alluxio UFS by all workers</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesWrittenRemote"></a> Cluster.BytesWrittenRemote</td><td>COUNTER</td><td>Total number of bytes written to workers via network (RPC). Data is written to worker storage or is written by workers to underlying UFSes. This does not include short-circuit local writes and domain socket writes.</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesWrittenRemoteThroughput"></a> Cluster.BytesWrittenRemoteThroughput</td><td>GAUGE</td><td>Bytes write per minute throughput to workers via network (RPC). Data is written to worker storage or is written by workers to underlying UFSes. This does not include short-circuit local writes and domain socket writes.</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesWrittenUfsAll"></a> Cluster.BytesWrittenUfsAll</td><td>COUNTER</td><td>Total number of bytes written to all Alluxio UFSes by all workers</td></tr>
  <tr><td><a class="anchor" name="Cluster.BytesWrittenUfsThroughput"></a> Cluster.BytesWrittenUfsThroughput</td><td>GAUGE</td><td>Bytes write per minute throughput to all Alluxio UFSes by all workers</td></tr>
  <tr><td><a class="anchor" name="Cluster.CacheHitRate"></a> Cluster.CacheHitRate</td><td>GAUGE</td><td>Cache hit rate: (# bytes read from cache) / (# bytes requested)</td></tr>
  <tr><td><a class="anchor" name="Cluster.CapacityFree"></a> Cluster.CapacityFree</td><td>GAUGE</td><td>Total free bytes on all tiers, on all workers of Alluxio</td></tr>
  <tr><td><a class="anchor" name="Cluster.CapacityTotal"></a> Cluster.CapacityTotal</td><td>GAUGE</td><td>Total capacity (in bytes) on all tiers, on all workers of Alluxio</td></tr>
  <tr><td><a class="anchor" name="Cluster.CapacityUsed"></a> Cluster.CapacityUsed</td><td>GAUGE</td><td>Total used bytes on all tiers, on all workers of Alluxio</td></tr>
  <tr><td><a class="anchor" name="Cluster.LeaderId"></a> Cluster.LeaderId</td><td>GAUGE</td><td>Display current leader id</td></tr>
  <tr><td><a class="anchor" name="Cluster.LeaderIndex"></a> Cluster.LeaderIndex</td><td>GAUGE</td><td>Index of current leader</td></tr>
  <tr><td><a class="anchor" name="Cluster.LostWorkers"></a> Cluster.LostWorkers</td><td>GAUGE</td><td>Total number of lost workers inside the cluster</td></tr>
  <tr><td><a class="anchor" name="Cluster.RootUfsCapacityFree"></a> Cluster.RootUfsCapacityFree</td><td>GAUGE</td><td>Free capacity of the Alluxio root UFS in bytes</td></tr>
  <tr><td><a class="anchor" name="Cluster.RootUfsCapacityTotal"></a> Cluster.RootUfsCapacityTotal</td><td>GAUGE</td><td>Total capacity of the Alluxio root UFS in bytes</td></tr>
  <tr><td><a class="anchor" name="Cluster.RootUfsCapacityUsed"></a> Cluster.RootUfsCapacityUsed</td><td>GAUGE</td><td>Used capacity of the Alluxio root UFS in bytes</td></tr>
  <tr><td><a class="anchor" name="Cluster.Workers"></a> Cluster.Workers</td><td>GAUGE</td><td>Total number of active workers inside the cluster</td></tr>
</table>


## Process Metrics

Metrics shared by the all Alluxio server and client processes.

<table class="table table-striped">
<tr><th>Name</th><th>Type</th><th>Description</th></tr>
  <tr>
    <td><a class="anchor" name="Process.pool.direct.mem.used"></a> Process.pool.direct.mem.used</td>
    <td>GAUGE</td>
    <td>The used direct memory by NIO direct buffer pool</td>
  </tr>
</table>

## Server Metrics

Metrics shared by the Alluxio server processes.
<table class="table table-striped">
<tr><th>Name</th><th>Type</th><th>Description</th></tr>

  <tr>
    <td><a class="anchor" name="Server.JvmPauseMonitorInfoTimeExceeded"></a> Server.JvmPauseMonitorInfoTimeExceeded</td>
    <td>GAUGE</td>
    <td>The total number of times that JVM slept and the sleep period is larger than the info level threshold defined by alluxio.jvm.monitor.info.threshold</td>
  </tr>

  <tr>
    <td><a class="anchor" name="Server.JvmPauseMonitorTotalExtraTime"></a> Server.JvmPauseMonitorTotalExtraTime</td>
    <td>GAUGE</td>
    <td>The total time that JVM slept and didn't do GC</td>
  </tr>

  <tr>
    <td><a class="anchor" name="Server.JvmPauseMonitorWarnTimeExceeded"></a> Server.JvmPauseMonitorWarnTimeExceeded</td>
    <td>GAUGE</td>
    <td>The total number of times that JVM slept and the sleep period is larger than the warn level threshold defined by alluxio.jvm.monitor.warn.threshold</td>
  </tr>

</table>

## Master Metrics

Default master metrics:

<table class="table table-striped">
<tr><th>Name</th><th>Type</th><th>Description</th></tr>
<tr>  <td><a class="anchor" name="Master.AbsentCacheHits"></a> Master.AbsentCacheHits</td>  <td>GAUGE</td>  <td>Number of cache hits on the absent cache</td></tr>
<tr>  <td><a class="anchor" name="Master.AbsentCacheMisses"></a> Master.AbsentCacheMisses</td>  <td>GAUGE</td>  <td>Number of cache misses on the absent cache</td></tr>
<tr>  <td><a class="anchor" name="Master.AbsentCacheSize"></a> Master.AbsentCacheSize</td>  <td>GAUGE</td>  <td>Size of the absent cache</td></tr>
<tr>  <td><a class="anchor" name="Master.AbsentPathCacheQueueSize"></a> Master.AbsentPathCacheQueueSize</td>  <td>GAUGE</td>  <td>Alluxio maintains a cache of absent UFS paths. This is the number of UFS paths being processed.</td></tr>
<tr>  <td><a class="anchor" name="Master.AsyncPersistCancel"></a> Master.AsyncPersistCancel</td>  <td>COUNTER</td>  <td>The number of cancelled AsyncPersist operations</td></tr>
<tr>  <td><a class="anchor" name="Master.AsyncPersistFail"></a> Master.AsyncPersistFail</td>  <td>COUNTER</td>  <td>The number of failed AsyncPersist operations</td></tr>
<tr>  <td><a class="anchor" name="Master.AsyncPersistFileCount"></a> Master.AsyncPersistFileCount</td>  <td>COUNTER</td>  <td>The number of files created by AsyncPersist operations</td></tr>
<tr>  <td><a class="anchor" name="Master.AsyncPersistFileSize"></a> Master.AsyncPersistFileSize</td>  <td>COUNTER</td>  <td>The total size of files created by AsyncPersist operations</td></tr>
<tr>  <td><a class="anchor" name="Master.AsyncPersistSuccess"></a> Master.AsyncPersistSuccess</td>  <td>COUNTER</td>  <td>The number of successful AsyncPersist operations</td></tr>
<tr>  <td><a class="anchor" name="Master.AuditLogEntriesSize"></a> Master.AuditLogEntriesSize</td>  <td>GAUGE</td>  <td>The size of the audit log entries blocking queue</td></tr>
<tr>  <td><a class="anchor" name="Master.BlockHeapSize"></a> Master.BlockHeapSize</td>  <td>GAUGE</td>  <td>An estimate of the blocks heap size</td></tr>
<tr>  <td><a class="anchor" name="Master.BlockReplicaCount"></a> Master.BlockReplicaCount</td>  <td>GAUGE</td>  <td>Total number of block replicas in Alluxio</td></tr>
<tr>  <td><a class="anchor" name="Master.CachedBlockLocations"></a> Master.CachedBlockLocations</td>  <td>GAUGE</td>  <td>Total number of cached block locations</td></tr>
<tr>  <td><a class="anchor" name="Master.CompleteFileOps"></a> Master.CompleteFileOps</td>  <td>COUNTER</td>  <td>Total number of the CompleteFile operations</td></tr>
<tr>  <td><a class="anchor" name="Master.CompletedOperationRetryCount"></a> Master.CompletedOperationRetryCount</td>  <td>COUNTER</td>  <td>Total number of completed operations that has been retried by client.</td></tr>
<tr>  <td><a class="anchor" name="Master.CreateDirectoryOps"></a> Master.CreateDirectoryOps</td>  <td>COUNTER</td>  <td>Total number of the CreateDirectory operations</td></tr>
<tr>  <td><a class="anchor" name="Master.CreateFileOps"></a> Master.CreateFileOps</td>  <td>COUNTER</td>  <td>Total number of the CreateFile operations</td></tr>
<tr>  <td><a class="anchor" name="Master.DeletePathOps"></a> Master.DeletePathOps</td>  <td>COUNTER</td>  <td>Total number of the Delete operations</td></tr>
<tr>  <td><a class="anchor" name="Master.DirectoriesCreated"></a> Master.DirectoriesCreated</td>  <td>COUNTER</td>  <td>Total number of the succeed CreateDirectory operations</td></tr>
<tr>  <td><a class="anchor" name="Master.EdgeCacheEvictions"></a> Master.EdgeCacheEvictions</td>  <td>GAUGE</td>  <td>Total number of edges (inode metadata) that was evicted from cache. The edge cache is responsible for managing the mapping from (parentId, childName) to childId.</td></tr>
<tr>  <td><a class="anchor" name="Master.EdgeCacheHits"></a> Master.EdgeCacheHits</td>  <td>GAUGE</td>  <td>Total number of hits in the edge (inode metadata) cache. The edge cache is responsible for managing the mapping from (parentId, childName) to childId.</td></tr>
<tr>  <td><a class="anchor" name="Master.EdgeCacheLoadTimes"></a> Master.EdgeCacheLoadTimes</td>  <td>GAUGE</td>  <td>Total load times in the edge (inode metadata) cache that resulted from a cache miss. The edge cache is responsible for managing the mapping from (parentId, childName) to childId.</td></tr>
<tr>  <td><a class="anchor" name="Master.EdgeCacheMisses"></a> Master.EdgeCacheMisses</td>  <td>GAUGE</td>  <td>Total number of misses in the edge (inode metadata) cache. The edge cache is responsible for managing the mapping from (parentId, childName) to childId.</td></tr>
<tr>  <td><a class="anchor" name="Master.EdgeCacheSize"></a> Master.EdgeCacheSize</td>  <td>GAUGE</td>  <td>Total number of edges (inode metadata) cached. The edge cache is responsible for managing the mapping from (parentId, childName) to childId.</td></tr>
<tr>  <td><a class="anchor" name="Master.EdgeLockPoolSize"></a> Master.EdgeLockPoolSize</td>  <td>GAUGE</td>  <td>The size of master edge lock pool</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotDownloadDiskSize"></a> Master.EmbeddedJournalLastSnapshotDownloadDiskSize</td>  <td>GAUGE</td>  <td>Describes the size on disk of the snapshot downloaded from other masters in the cluster the previous time the download occurred. Only valid when using the embedded journal.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotDownloadDurationMs"></a> Master.EmbeddedJournalLastSnapshotDownloadDurationMs</td>  <td>GAUGE</td>  <td>Describes the amount of time taken to download journal snapshots from other masters in the cluster the previous time the download occurred. Only valid when using the embedded journal.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotDownloadSize"></a> Master.EmbeddedJournalLastSnapshotDownloadSize</td>  <td>GAUGE</td>  <td>Describes the size of the snapshot downloaded from other masters in the cluster the previous time the download occurred. Only valid when using the embedded journal.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotDurationMs"></a> Master.EmbeddedJournalLastSnapshotDurationMs</td>  <td>GAUGE</td>  <td>Describes the amount of time taken to generate the last local journal snapshots on this master. Only valid when using the embedded journal.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotEntriesCount"></a> Master.EmbeddedJournalLastSnapshotEntriesCount</td>  <td>GAUGE</td>  <td>Describes the number of entries in the last local journal snapshots on this master. Only valid when using the embedded journal.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotReplayDurationMs"></a> Master.EmbeddedJournalLastSnapshotReplayDurationMs</td>  <td>GAUGE</td>  <td>Represents the time the last restore from checkpoint operation took in milliseconds.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotReplayEntriesCount"></a> Master.EmbeddedJournalLastSnapshotReplayEntriesCount</td>  <td>GAUGE</td>  <td>Represents the time the last restore from checkpoint operation took in milliseconds.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotUploadDiskSize"></a> Master.EmbeddedJournalLastSnapshotUploadDiskSize</td>  <td>GAUGE</td>  <td>Describes the size on disk of the snapshot uploaded to other masters in the cluster the previous time the download occurred. Only valid when using the embedded journal.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotUploadDurationMs"></a> Master.EmbeddedJournalLastSnapshotUploadDurationMs</td>  <td>GAUGE</td>  <td>Describes the amount of time taken to upload journal snapshots to another master in the cluster the previous time the upload occurred. Only valid when using the embedded journal.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotUploadSize"></a> Master.EmbeddedJournalLastSnapshotUploadSize</td>  <td>GAUGE</td>  <td>Describes the size of the snapshot uploaded to other masters in the cluster the previous time the download occurred. Only valid when using the embedded journal.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotDownloadDiskHistogram"></a> Master.EmbeddedJournalSnapshotDownloadDiskHistogram</td>  <td>HISTOGRAM</td>  <td>Describes the size on disk of the snapshot downloaded from another master in the cluster. Only valid when using the embedded journal. Long running average.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotDownloadGenerate"></a> Master.EmbeddedJournalSnapshotDownloadGenerate</td>  <td>TIMER</td>  <td>Describes the amount of time taken to download journal snapshots from other masters in the cluster. Only valid when using the embedded journal. Long running average.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotDownloadHistogram"></a> Master.EmbeddedJournalSnapshotDownloadHistogram</td>  <td>HISTOGRAM</td>  <td>Describes the size of the snapshot downloaded from another master in the cluster. Only valid when using the embedded journal. Long running average.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotGenerateTimer"></a> Master.EmbeddedJournalSnapshotGenerateTimer</td>  <td>TIMER</td>  <td>Describes the amount of time taken to generate local journal snapshots on this master. Only valid when using the embedded journal. Use this metric to measure the performance of Alluxio's snapshot generation.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotInstallTimer"></a> Master.EmbeddedJournalSnapshotInstallTimer</td>  <td>TIMER</td>  <td>Describes the amount of time taken to install a downloaded journal snapshot from another master. Only valid only when using the embedded journal. Use this metric to determine the performance of Alluxio when installing snapshots from the leader. Higher numbers may indicate a slow disk or CPU contention.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotLastIndex"></a> Master.EmbeddedJournalSnapshotLastIndex</td>  <td>GAUGE</td>  <td>Represents the latest journal index that was recorded by this master in the most recent local snapshot or from a snapshot downloaded from another master in the cluster. Only valid when using the embedded journal.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotReplayTimer"></a> Master.EmbeddedJournalSnapshotReplayTimer</td>  <td>TIMER</td>  <td>Describes the amount of time taken to replay a journal snapshot onto the master's state machine. Only valid only when using the embedded journal. Use this metric to determine the performance of Alluxio when replaying journal snapshot file. Higher numbers may indicate a slow disk or CPU contention</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotUploadDiskHistogram"></a> Master.EmbeddedJournalSnapshotUploadDiskHistogram</td>  <td>HISTOGRAM</td>  <td>Describes the size on disk of the snapshot uploaded to another master in the cluster. Only valid when using the embedded journal. Long running average.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotUploadHistogram"></a> Master.EmbeddedJournalSnapshotUploadHistogram</td>  <td>HISTOGRAM</td>  <td>Describes the size of the snapshot uploaded to another master in the cluster. Only valid when using the embedded journal. Long running average.</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotUploadTimer"></a> Master.EmbeddedJournalSnapshotUploadTimer</td>  <td>TIMER</td>  <td>Describes the amount of time taken to upload journal snapshots to another master in the cluster. Only valid when using the embedded journal. long running average</td></tr>
<tr>  <td><a class="anchor" name="Master.FileBlockInfosGot"></a> Master.FileBlockInfosGot</td>  <td>COUNTER</td>  <td>Total number of succeed GetFileBlockInfo operations</td></tr>
<tr>  <td><a class="anchor" name="Master.FileInfosGot"></a> Master.FileInfosGot</td>  <td>COUNTER</td>  <td>Total number of the succeed GetFileInfo operations</td></tr>
<tr>  <td><a class="anchor" name="Master.FileSize"></a> Master.FileSize</td>  <td>GAUGE</td>  <td>File size distribution</td></tr>
<tr>  <td><a class="anchor" name="Master.FilesCompleted"></a> Master.FilesCompleted</td>  <td>COUNTER</td>  <td>Total number of the succeed CompleteFile operations</td></tr>
<tr>  <td><a class="anchor" name="Master.FilesCreated"></a> Master.FilesCreated</td>  <td>COUNTER</td>  <td>Total number of the succeed CreateFile operations</td></tr>
<tr>  <td><a class="anchor" name="Master.FilesFreed"></a> Master.FilesFreed</td>  <td>COUNTER</td>  <td>Total number of succeed FreeFile operations</td></tr>
<tr>  <td><a class="anchor" name="Master.FilesPersisted"></a> Master.FilesPersisted</td>  <td>COUNTER</td>  <td>Total number of successfully persisted files</td></tr>
<tr>  <td><a class="anchor" name="Master.FilesPinned"></a> Master.FilesPinned</td>  <td>GAUGE</td>  <td>Total number of currently pinned files. Note that IDs for these files are stored in memory.</td></tr>
<tr>  <td><a class="anchor" name="Master.FilesToBePersisted"></a> Master.FilesToBePersisted</td>  <td>GAUGE</td>  <td>Total number of currently to be persisted files. Note that the IDs for these files are stored in memory.</td></tr>
<tr>  <td><a class="anchor" name="Master.FreeFileOps"></a> Master.FreeFileOps</td>  <td>COUNTER</td>  <td>Total number of FreeFile operations</td></tr>
<tr>  <td><a class="anchor" name="Master.GetFileBlockInfoOps"></a> Master.GetFileBlockInfoOps</td>  <td>COUNTER</td>  <td>Total number of GetFileBlockInfo operations</td></tr>
<tr>  <td><a class="anchor" name="Master.GetFileInfoOps"></a> Master.GetFileInfoOps</td>  <td>COUNTER</td>  <td>Total number of the GetFileInfo operations</td></tr>
<tr>  <td><a class="anchor" name="Master.GetNewBlockOps"></a> Master.GetNewBlockOps</td>  <td>COUNTER</td>  <td>Total number of the GetNewBlock operations</td></tr>
<tr>  <td><a class="anchor" name="Master.InodeCacheEvictions"></a> Master.InodeCacheEvictions</td>  <td>GAUGE</td>  <td>Total number of inodes that was evicted from the cache.</td></tr>
<tr>  <td><a class="anchor" name="Master.InodeCacheHitRatio"></a> Master.InodeCacheHitRatio</td>  <td>GAUGE</td>  <td>Inode Cache hit ratio</td></tr>
<tr>  <td><a class="anchor" name="Master.InodeCacheHits"></a> Master.InodeCacheHits</td>  <td>GAUGE</td>  <td>Total number of hits in the inodes (inode metadata) cache.</td></tr>
<tr>  <td><a class="anchor" name="Master.InodeCacheLoadTimes"></a> Master.InodeCacheLoadTimes</td>  <td>GAUGE</td>  <td>Total load times in the inodes (inode metadata) cache that resulted from a cache miss.</td></tr>
<tr>  <td><a class="anchor" name="Master.InodeCacheMisses"></a> Master.InodeCacheMisses</td>  <td>GAUGE</td>  <td>Total number of misses in the inodes (inode metadata) cache.</td></tr>
<tr>  <td><a class="anchor" name="Master.InodeCacheSize"></a> Master.InodeCacheSize</td>  <td>GAUGE</td>  <td>Total number of inodes (inode metadata) cached.</td></tr>
<tr>  <td><a class="anchor" name="Master.InodeHeapSize"></a> Master.InodeHeapSize</td>  <td>GAUGE</td>  <td>An estimate of the inode heap size</td></tr>
<tr>  <td><a class="anchor" name="Master.InodeLockPoolSize"></a> Master.InodeLockPoolSize</td>  <td>GAUGE</td>  <td>The size of master inode lock pool</td></tr>
<tr>  <td><a class="anchor" name="Master.JobCanceled"></a> Master.JobCanceled</td>  <td>COUNTER</td>  <td>The number of canceled status job</td></tr>
<tr>  <td><a class="anchor" name="Master.JobCompleted"></a> Master.JobCompleted</td>  <td>COUNTER</td>  <td>The number of completed status job</td></tr>
<tr>  <td><a class="anchor" name="Master.JobCount"></a> Master.JobCount</td>  <td>GAUGE</td>  <td>The number of all status job</td></tr>
<tr>  <td><a class="anchor" name="Master.JobCreated"></a> Master.JobCreated</td>  <td>COUNTER</td>  <td>The number of created status job</td></tr>
<tr>  <td><a class="anchor" name="Master.JobDistributedLoadBlockSizes"></a> Master.JobDistributedLoadBlockSizes</td>  <td>COUNTER</td>  <td>The total block size loaded by load commands</td></tr>
<tr>  <td><a class="anchor" name="Master.JobDistributedLoadCancel"></a> Master.JobDistributedLoadCancel</td>  <td>COUNTER</td>  <td>The number of cancelled DistributedLoad operations</td></tr>
<tr>  <td><a class="anchor" name="Master.JobDistributedLoadFail"></a> Master.JobDistributedLoadFail</td>  <td>COUNTER</td>  <td>The number of failed DistributedLoad operations</td></tr>
<tr>  <td><a class="anchor" name="Master.JobDistributedLoadFileCount"></a> Master.JobDistributedLoadFileCount</td>  <td>COUNTER</td>  <td>The number of files by DistributedLoad operations</td></tr>
<tr>  <td><a class="anchor" name="Master.JobDistributedLoadFileSizes"></a> Master.JobDistributedLoadFileSizes</td>  <td>COUNTER</td>  <td>The total file size by DistributedLoad operations</td></tr>
<tr>  <td><a class="anchor" name="Master.JobDistributedLoadRate"></a> Master.JobDistributedLoadRate</td>  <td>METER</td>  <td>The average DistributedLoad loading rate</td></tr>
<tr>  <td><a class="anchor" name="Master.JobDistributedLoadSuccess"></a> Master.JobDistributedLoadSuccess</td>  <td>COUNTER</td>  <td>The number of successful DistributedLoad operations</td></tr>
<tr>  <td><a class="anchor" name="Master.JobFailed"></a> Master.JobFailed</td>  <td>COUNTER</td>  <td>The number of failed status job</td></tr>
<tr>  <td><a class="anchor" name="Master.JobLoadBlockCount"></a> Master.JobLoadBlockCount</td>  <td>COUNTER</td>  <td>The number of blocks loaded by load commands</td></tr>
<tr>  <td><a class="anchor" name="Master.JobLoadBlockFail"></a> Master.JobLoadBlockFail</td>  <td>COUNTER</td>  <td>The number of blocks failed to be loaded by load commands</td></tr>
<tr>  <td><a class="anchor" name="Master.JobLoadFail"></a> Master.JobLoadFail</td>  <td>COUNTER</td>  <td>The number of failed Load commands</td></tr>
<tr>  <td><a class="anchor" name="Master.JobLoadRate"></a> Master.JobLoadRate</td>  <td>METER</td>  <td>The average loading rate of Load commands</td></tr>
<tr>  <td><a class="anchor" name="Master.JobLoadSuccess"></a> Master.JobLoadSuccess</td>  <td>COUNTER</td>  <td>The number of successful Load commands</td></tr>
<tr>  <td><a class="anchor" name="Master.JobRunning"></a> Master.JobRunning</td>  <td>COUNTER</td>  <td>The number of running status job</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalCheckpointWarn"></a> Master.JournalCheckpointWarn</td>  <td>GAUGE</td>  <td>If the raft log index exceeds alluxio.master.journal.checkpoint.period.entries, and the last checkpoint exceeds alluxio.master.journal.checkpoint.warning.threshold.time, it returns 1 to indicate that a warning is required, otherwise it returns 0</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalEntriesSinceCheckPoint"></a> Master.JournalEntriesSinceCheckPoint</td>  <td>GAUGE</td>  <td>Journal entries since last checkpoint</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalFlushFailure"></a> Master.JournalFlushFailure</td>  <td>COUNTER</td>  <td>Total number of failed journal flush</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalFlushTimer"></a> Master.JournalFlushTimer</td>  <td>TIMER</td>  <td>The timer statistics of journal flush</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalFreeBytes"></a> Master.JournalFreeBytes</td>  <td>GAUGE</td>  <td>Bytes left on the journal disk(s) for an Alluxio master. This metric is only valid on Linux and when embedded journal is used. Use this metric to monitor whether your journal is running out of disk space.</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalFreePercent"></a> Master.JournalFreePercent</td>  <td>GAUGE</td>  <td>Percentage of free space left on the journal disk(s) for an Alluxio master.This metric is only valid on Linux and when embedded journal is used. Use this metric to monitor whether your journal is running out of disk space.</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalGainPrimacyTimer"></a> Master.JournalGainPrimacyTimer</td>  <td>TIMER</td>  <td>The timer statistics of journal gain primacy</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalLastAppliedCommitIndex"></a> Master.JournalLastAppliedCommitIndex</td>  <td>GAUGE</td>  <td>The last raft log index which was applied to the state machine</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalLastCheckPointTime"></a> Master.JournalLastCheckPointTime</td>  <td>GAUGE</td>  <td>Last Journal Checkpoint Time</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalSequenceNumber"></a> Master.JournalSequenceNumber</td>  <td>GAUGE</td>  <td>Current journal sequence number</td></tr>
<tr>  <td><a class="anchor" name="Master.LastBackupEntriesCount"></a> Master.LastBackupEntriesCount</td>  <td>GAUGE</td>  <td>The total number of entries written in the last leading master metadata backup</td></tr>
<tr>  <td><a class="anchor" name="Master.LastBackupRestoreCount"></a> Master.LastBackupRestoreCount</td>  <td>GAUGE</td>  <td>The total number of entries restored from backup when a leading master initializes its metadata</td></tr>
<tr>  <td><a class="anchor" name="Master.LastBackupRestoreTimeMs"></a> Master.LastBackupRestoreTimeMs</td>  <td>GAUGE</td>  <td>The process time of the last restore from backup</td></tr>
<tr>  <td><a class="anchor" name="Master.LastBackupTimeMs"></a> Master.LastBackupTimeMs</td>  <td>GAUGE</td>  <td>The process time of the last backup</td></tr>
<tr>  <td><a class="anchor" name="Master.LastGainPrimacyTime"></a> Master.LastGainPrimacyTime</td>  <td>GAUGE</td>  <td>Last time the master gains primacy</td></tr>
<tr>  <td><a class="anchor" name="Master.LastLosePrimacyTime"></a> Master.LastLosePrimacyTime</td>  <td>GAUGE</td>  <td>Last time the master loses primacy</td></tr>
<tr>  <td><a class="anchor" name="Master.ListingCacheEvictions"></a> Master.ListingCacheEvictions</td>  <td>COUNTER</td>  <td>The total number of evictions in master listing cache</td></tr>
<tr>  <td><a class="anchor" name="Master.ListingCacheHits"></a> Master.ListingCacheHits</td>  <td>COUNTER</td>  <td>The total number of hits in master listing cache</td></tr>
<tr>  <td><a class="anchor" name="Master.ListingCacheLoadTimes"></a> Master.ListingCacheLoadTimes</td>  <td>COUNTER</td>  <td>The total load time (in nanoseconds) in master listing cache that resulted from a cache miss.</td></tr>
<tr>  <td><a class="anchor" name="Master.ListingCacheMisses"></a> Master.ListingCacheMisses</td>  <td>COUNTER</td>  <td>The total number of misses in master listing cache</td></tr>
<tr>  <td><a class="anchor" name="Master.ListingCacheSize"></a> Master.ListingCacheSize</td>  <td>GAUGE</td>  <td>The size of master listing cache</td></tr>
<tr>  <td><a class="anchor" name="Master.LostBlockCount"></a> Master.LostBlockCount</td>  <td>GAUGE</td>  <td>Count of lost unique blocks</td></tr>
<tr>  <td><a class="anchor" name="Master.LostFileCount"></a> Master.LostFileCount</td>  <td>GAUGE</td>  <td>Count of lost files. This number is cached and may not be in sync with Master.LostBlockCount</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncActivePaths"></a> Master.MetadataSyncActivePaths</td>  <td>COUNTER</td>  <td>The number of in-progress paths from all InodeSyncStream instances</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncExecutor"></a> Master.MetadataSyncExecutor</td>  <td>EXECUTOR_SERVICE</td>  <td>Metrics concerning the master metadata sync executor threads. Master.MetadataSyncExecutor.submitted is a meter of the tasks submitted to the executor. Master.MetadataSyncExecutor.completed is a meter of the tasks completed by the executor. Master.MetadataSyncExecutor.activeTaskQueue is exponentially-decaying random reservoir of the number of active tasks (running or submitted) at the executor calculated each time a new task is added to the executor. The max value is the maximum number of active tasks at any time during execution. Master.MetadataSyncExecutor.running is the number of tasks actively being run by the executor. Master.MetadataSyncExecutor.idle is the time spent idling by the submitted tasks (i.e. waiting the the queue before being executed). Master.MetadataSyncExecutor.duration is the time spent running the submitted tasks. If the executor is a thread pool executor then Master.MetadataSyncExecutor.queueSize is the size of the task queue.</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncExecutorQueueSize"></a> Master.MetadataSyncExecutorQueueSize</td>  <td>GAUGE</td>  <td>The number of queuing sync tasks in the metadata sync thread pool controlled by alluxio.master.metadata.sync.executor.pool.size</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncFail"></a> Master.MetadataSyncFail</td>  <td>COUNTER</td>  <td>The number of InodeSyncStream that failed, either partially or fully</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncNoChange"></a> Master.MetadataSyncNoChange</td>  <td>COUNTER</td>  <td>The number of InodeSyncStream that finished with no change to inodes.</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncOpsCount"></a> Master.MetadataSyncOpsCount</td>  <td>COUNTER</td>  <td>The number of metadata sync operations. Each sync operation corresponds to one InodeSyncStream instance.</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPathsCancel"></a> Master.MetadataSyncPathsCancel</td>  <td>COUNTER</td>  <td>The number of pending paths from all InodeSyncStream instances that are ignored in the end instead of processed</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPathsFail"></a> Master.MetadataSyncPathsFail</td>  <td>COUNTER</td>  <td>The number of paths that failed during metadata sync from all InodeSyncStream instances</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPathsSuccess"></a> Master.MetadataSyncPathsSuccess</td>  <td>COUNTER</td>  <td>The number of paths sync-ed from all InodeSyncStream instances</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPendingPaths"></a> Master.MetadataSyncPendingPaths</td>  <td>COUNTER</td>  <td>The number of pending paths from all active InodeSyncStream instances,waiting for metadata sync</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPrefetchCancel"></a> Master.MetadataSyncPrefetchCancel</td>  <td>COUNTER</td>  <td>Number of cancelled prefetch jobs from metadata sync</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPrefetchExecutor"></a> Master.MetadataSyncPrefetchExecutor</td>  <td>EXECUTOR_SERVICE</td>  <td>Metrics concerning the master metadata sync prefetchexecutor threads. Master.MetadataSyncPrefetchExecutor.submitted is a meter of the tasks submitted to the executor. Master.MetadataSyncPrefetchExecutor.completed is a meter of the tasks completed by the executor. Master.MetadataSyncPrefetchExecutor.activeTaskQueue is exponentially-decaying random reservoir of the number of active tasks (running or submitted) at the executor calculated each time a new task is added to the executor. The max value is the maximum number of active tasks at any time during execution. Master.MetadataSyncPrefetchExecutor.running is the number of tasks actively being run by the executor. Master.MetadataSyncPrefetchExecutor.idle is the time spent idling by the submitted tasks (i.e. waiting the the queue before being executed). Master.MetadataSyncPrefetchExecutor.duration is the time spent running the submitted tasks. If the executor is a thread pool executor then Master.MetadataSyncPrefetchExecutor.queueSize is the size of the task queue.</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPrefetchExecutorQueueSize"></a> Master.MetadataSyncPrefetchExecutorQueueSize</td>  <td>GAUGE</td>  <td>The number of queuing prefetch tasks in the metadata sync thread pool controlled by alluxio.master.metadata.sync.ufs.prefetch.pool.size</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPrefetchFail"></a> Master.MetadataSyncPrefetchFail</td>  <td>COUNTER</td>  <td>Number of failed prefetch jobs from metadata sync</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPrefetchOpsCount"></a> Master.MetadataSyncPrefetchOpsCount</td>  <td>COUNTER</td>  <td>The number of prefetch operations handled by the prefetch thread pool</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPrefetchPaths"></a> Master.MetadataSyncPrefetchPaths</td>  <td>COUNTER</td>  <td>Total number of UFS paths fetched by prefetch jobs from metadata sync</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPrefetchRetries"></a> Master.MetadataSyncPrefetchRetries</td>  <td>COUNTER</td>  <td>Number of retries to get from prefetch jobs from metadata sync</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPrefetchSuccess"></a> Master.MetadataSyncPrefetchSuccess</td>  <td>COUNTER</td>  <td>Number of successful prefetch jobs from metadata sync</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncSkipped"></a> Master.MetadataSyncSkipped</td>  <td>COUNTER</td>  <td>The number of InodeSyncStream that are skipped because the Alluxio metadata is fresher than alluxio.user.file.metadata.sync.interval</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncSuccess"></a> Master.MetadataSyncSuccess</td>  <td>COUNTER</td>  <td>The number of InodeSyncStream that succeeded</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncTimeMs"></a> Master.MetadataSyncTimeMs</td>  <td>COUNTER</td>  <td>The total time elapsed in all InodeSyncStream instances</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncUfsMount."></a> Master.MetadataSyncUfsMount.</td>  <td>COUNTER</td>  <td>The number of UFS sync operations for a given mount point</td></tr>
<tr>  <td><a class="anchor" name="Master.MigrateJobCancel"></a> Master.MigrateJobCancel</td>  <td>COUNTER</td>  <td>The number of cancelled MigrateJob operations</td></tr>
<tr>  <td><a class="anchor" name="Master.MigrateJobFail"></a> Master.MigrateJobFail</td>  <td>COUNTER</td>  <td>The number of failed MigrateJob operations</td></tr>
<tr>  <td><a class="anchor" name="Master.MigrateJobFileCount"></a> Master.MigrateJobFileCount</td>  <td>COUNTER</td>  <td>The number of MigrateJob files</td></tr>
<tr>  <td><a class="anchor" name="Master.MigrateJobFileSize"></a> Master.MigrateJobFileSize</td>  <td>COUNTER</td>  <td>The total size of MigrateJob files</td></tr>
<tr>  <td><a class="anchor" name="Master.MigrateJobSuccess"></a> Master.MigrateJobSuccess</td>  <td>COUNTER</td>  <td>The number of successful MigrateJob operations</td></tr>
<tr>  <td><a class="anchor" name="Master.MountOps"></a> Master.MountOps</td>  <td>COUNTER</td>  <td>Total number of Mount operations</td></tr>
<tr>  <td><a class="anchor" name="Master.NewBlocksGot"></a> Master.NewBlocksGot</td>  <td>COUNTER</td>  <td>Total number of the succeed GetNewBlock operations</td></tr>
<tr>  <td><a class="anchor" name="Master.PathsDeleted"></a> Master.PathsDeleted</td>  <td>COUNTER</td>  <td>Total number of the succeed Delete operations</td></tr>
<tr>  <td><a class="anchor" name="Master.PathsMounted"></a> Master.PathsMounted</td>  <td>COUNTER</td>  <td>Total number of succeed Mount operations</td></tr>
<tr>  <td><a class="anchor" name="Master.PathsRenamed"></a> Master.PathsRenamed</td>  <td>COUNTER</td>  <td>Total number of succeed Rename operations</td></tr>
<tr>  <td><a class="anchor" name="Master.PathsUnmounted"></a> Master.PathsUnmounted</td>  <td>COUNTER</td>  <td>Total number of succeed Unmount operations</td></tr>
<tr>  <td><a class="anchor" name="Master.RenamePathOps"></a> Master.RenamePathOps</td>  <td>COUNTER</td>  <td>Total number of Rename operations</td></tr>
<tr>  <td><a class="anchor" name="Master.ReplicaMgmtActiveJobSize"></a> Master.ReplicaMgmtActiveJobSize</td>  <td>GAUGE</td>  <td>Number of active block replication/eviction jobs. These jobs are created by the master to maintain the block replica factor. The value is an estimate with lag. </td></tr>
<tr>  <td><a class="anchor" name="Master.ReplicationLimitedFiles"></a> Master.ReplicationLimitedFiles</td>  <td>COUNTER</td>  <td>Number of files that have a replication count set to a non-default value. Note that these files have IDs that are stored in memory.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockBackgroundErrors"></a> Master.RocksBlockBackgroundErrors</td>  <td>GAUGE</td>  <td>RocksDB block table. Accumulated number of background errors.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockBlockCacheCapacity"></a> Master.RocksBlockBlockCacheCapacity</td>  <td>GAUGE</td>  <td>RocksDB block table. Block cache capacity.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockBlockCachePinnedUsage"></a> Master.RocksBlockBlockCachePinnedUsage</td>  <td>GAUGE</td>  <td>RocksDB block table. Memory size for the entries being pinned.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockBlockCacheUsage"></a> Master.RocksBlockBlockCacheUsage</td>  <td>GAUGE</td>  <td>RocksDB block table. Memory size for the entries residing in block cache.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockCompactionPending"></a> Master.RocksBlockCompactionPending</td>  <td>GAUGE</td>  <td>RocksDB block table. This metric 1 if at least one compaction is pending; otherwise, the metric reports 0.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockCurSizeActiveMemTable"></a> Master.RocksBlockCurSizeActiveMemTable</td>  <td>GAUGE</td>  <td>RocksDB block table. Approximate size of active memtable in bytes.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockCurSizeAllMemTables"></a> Master.RocksBlockCurSizeAllMemTables</td>  <td>GAUGE</td>  <td>RocksDB block table. Approximate size of active, unflushed immutable, and pinned immutable memtables in bytes. Pinned immutable memtables are flushed memtables that are kept in memory to maintain write history in memory.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockEstimateNumKeys"></a> Master.RocksBlockEstimateNumKeys</td>  <td>GAUGE</td>  <td>RocksDB block table. Estimated number of total keys in the active and unflushed immutable memtables and storage.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockEstimatePendingCompactionBytes"></a> Master.RocksBlockEstimatePendingCompactionBytes</td>  <td>GAUGE</td>  <td>RocksDB block table. Estimated total number of bytes a compaction needs to rewrite on disk to get all levels down to under target size. In other words, this metrics relates to the write amplification in level compaction. Thus, this metric is not valid for compactions other than level-based.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockEstimateTableReadersMem"></a> Master.RocksBlockEstimateTableReadersMem</td>  <td>GAUGE</td>  <td>RocksDB inode table. Estimated memory in bytes used for reading SST tables, excluding memory used in block cache (e.g., filter and index blocks). This metric records the memory used by iterators as well as filters and indices if the filters and indices are not maintained in the block cache. Basically this metric reports the memory used outside the block cache to read data.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockEstimatedMemUsage"></a> Master.RocksBlockEstimatedMemUsage</td>  <td>GAUGE</td>  <td>RocksDB block table. This metric estimates the memory usage of the RockDB Block table by aggregating the values of Master.RocksBlockBlockCacheUsage, Master.RocksBlockEstimateTableReadersMem, Master.RocksBlockCurSizeAllMemTables, and Master.RocksBlockBlockCachePinnedUsage</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockLiveSstFilesSize"></a> Master.RocksBlockLiveSstFilesSize</td>  <td>GAUGE</td>  <td>RocksDB block table. Total size in bytes of all SST files that belong to the latest LSM tree.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockMemTableFlushPending"></a> Master.RocksBlockMemTableFlushPending</td>  <td>GAUGE</td>  <td>RocksDB block table. This metric returns 1 if a memtable flush is pending; otherwhise it returns 0.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockNumDeletesActiveMemTable"></a> Master.RocksBlockNumDeletesActiveMemTable</td>  <td>GAUGE</td>  <td>RocksDB block table. Total number of delete entries in the active memtable.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockNumDeletesImmMemTables"></a> Master.RocksBlockNumDeletesImmMemTables</td>  <td>GAUGE</td>  <td>RocksDB block table. Total number of delete entries in the unflushed immutable memtables.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockNumEntriesActiveMemTable"></a> Master.RocksBlockNumEntriesActiveMemTable</td>  <td>GAUGE</td>  <td>RocksDB block table. Total number of entries in the active memtable.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockNumEntriesImmMemTables"></a> Master.RocksBlockNumEntriesImmMemTables</td>  <td>GAUGE</td>  <td>RocksDB block table. Total number of entries in the unflushed immutable memtables.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockNumImmutableMemTable"></a> Master.RocksBlockNumImmutableMemTable</td>  <td>GAUGE</td>  <td>RocksDB block table. Number of immutable memtables that have not yet been flushed.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockNumLiveVersions"></a> Master.RocksBlockNumLiveVersions</td>  <td>GAUGE</td>  <td>RocksDB inode table. Number of live versions. More live versions often mean more SST files are held from being deleted, by iterators or unfinished compactions.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockNumRunningCompactions"></a> Master.RocksBlockNumRunningCompactions</td>  <td>GAUGE</td>  <td>RocksDB block table. Number of currently running compactions.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockNumRunningFlushes"></a> Master.RocksBlockNumRunningFlushes</td>  <td>GAUGE</td>  <td>RocksDB block table. Number of currently running flushes.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockSizeAllMemTables"></a> Master.RocksBlockSizeAllMemTables</td>  <td>GAUGE</td>  <td>RocksDB block table. Size all mem tables.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockTotalSstFilesSize"></a> Master.RocksBlockTotalSstFilesSize</td>  <td>GAUGE</td>  <td>RocksDB block table. Total size in bytes of all SST files.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeBackgroundErrors"></a> Master.RocksInodeBackgroundErrors</td>  <td>GAUGE</td>  <td>RocksDB inode table. Accumulated number of background errors.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeBlockCacheCapacity"></a> Master.RocksInodeBlockCacheCapacity</td>  <td>GAUGE</td>  <td>RocksDB inode table. Block cache capacity.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeBlockCachePinnedUsage"></a> Master.RocksInodeBlockCachePinnedUsage</td>  <td>GAUGE</td>  <td>RocksDB inode table. Memory size for the entries being pinned.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeBlockCacheUsage"></a> Master.RocksInodeBlockCacheUsage</td>  <td>GAUGE</td>  <td>RocksDB inode table. Memory size for the entries residing in block cache.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeCompactionPending"></a> Master.RocksInodeCompactionPending</td>  <td>GAUGE</td>  <td>RocksDB inode table. This metric 1 if at least one compaction is pending; otherwise, the metric reports 0.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeCurSizeActiveMemTable"></a> Master.RocksInodeCurSizeActiveMemTable</td>  <td>GAUGE</td>  <td>RocksDB inode table. Approximate size of active memtable in bytes.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeCurSizeAllMemTables"></a> Master.RocksInodeCurSizeAllMemTables</td>  <td>GAUGE</td>  <td>RocksDB inode table. Approximate size of active and unflushed immutable memtable in bytes.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeEstimateNumKeys"></a> Master.RocksInodeEstimateNumKeys</td>  <td>GAUGE</td>  <td>RocksDB inode table. Estimated number of total keys in the active and unflushed immutable memtables and storage.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeEstimatePendingCompactionBytes"></a> Master.RocksInodeEstimatePendingCompactionBytes</td>  <td>GAUGE</td>  <td>RocksDB block table. Estimated total number of bytes a compaction needs to rewrite on disk to get all levels down to under target size. In other words, this metrics relates to the write amplification in level compaction. Thus, this metric is not valid for compactions other than level-based.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeEstimateTableReadersMem"></a> Master.RocksInodeEstimateTableReadersMem</td>  <td>GAUGE</td>  <td>RocksDB inode table. Estimated memory in bytes used for reading SST tables, excluding memory used in block cache (e.g., filter and index blocks). This metric records the memory used by iterators as well as filters and indices if the filters and indices are not maintained in the block cache. Basically this metric reports the memory used outside the block cache to read data.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeEstimatedMemUsage"></a> Master.RocksInodeEstimatedMemUsage</td>  <td>GAUGE</td>  <td>RocksDB block table. This metric estimates the memory usage of the RockDB Inode table by aggregating the values of Master.RocksInodeBlockCacheUsage, Master.RocksInodeEstimateTableReadersMem, Master.RocksInodeCurSizeAllMemTables, and Master.RocksInodeBlockCachePinnedUsage</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeLiveSstFilesSize"></a> Master.RocksInodeLiveSstFilesSize</td>  <td>GAUGE</td>  <td>RocksDB inode table. Total size in bytes of all SST files that belong to the latest LSM tree.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeMemTableFlushPending"></a> Master.RocksInodeMemTableFlushPending</td>  <td>GAUGE</td>  <td>RocksDB inode table. This metric returns 1 if a memtable flush is pending; otherwhise it returns 0.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeNumDeletesActiveMemTable"></a> Master.RocksInodeNumDeletesActiveMemTable</td>  <td>GAUGE</td>  <td>RocksDB inode table. Total number of delete entries in the active memtable.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeNumDeletesImmMemTables"></a> Master.RocksInodeNumDeletesImmMemTables</td>  <td>GAUGE</td>  <td>RocksDB inode table. Total number of delete entries in the unflushed immutable memtables.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeNumEntriesActiveMemTable"></a> Master.RocksInodeNumEntriesActiveMemTable</td>  <td>GAUGE</td>  <td>RocksDB inode table. Total number of entries in the active memtable.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeNumEntriesImmMemTables"></a> Master.RocksInodeNumEntriesImmMemTables</td>  <td>GAUGE</td>  <td>RocksDB inode table. Total number of entries in the unflushed immutable memtables.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeNumImmutableMemTable"></a> Master.RocksInodeNumImmutableMemTable</td>  <td>GAUGE</td>  <td>RocksDB inode table. Number of immutable memtables that have not yet been flushed.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeNumLiveVersions"></a> Master.RocksInodeNumLiveVersions</td>  <td>GAUGE</td>  <td>RocksDB inode table. Number of live versions. More live versions often mean more SST files are held from being deleted, by iterators or unfinished compactions.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeNumRunningCompactions"></a> Master.RocksInodeNumRunningCompactions</td>  <td>GAUGE</td>  <td>RocksDB inode table. Number of currently running compactions.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeNumRunningFlushes"></a> Master.RocksInodeNumRunningFlushes</td>  <td>GAUGE</td>  <td>RocksDB inode table. Number of currently running flushes.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeSizeAllMemTables"></a> Master.RocksInodeSizeAllMemTables</td>  <td>GAUGE</td>  <td>RocksDB inode table. Approximate size of active, unflushed immutable, and pinned immutable memtables in bytes. Pinned immutable memtables are flushed memtables that are kept in memory to maintain write history in memory.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeTotalSstFilesSize"></a> Master.RocksInodeTotalSstFilesSize</td>  <td>GAUGE</td>  <td>RocksDB inode table. Total size in bytes of all SST files.</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksTotalEstimatedMemUsage"></a> Master.RocksTotalEstimatedMemUsage</td>  <td>GAUGE</td>  <td>This metric gives an estimate of the total memory used by RocksDB by aggregating the values of Master.RocksBlockEstimatedMemUsage and Master.RocksInodeEstimatedMemUsage</td></tr>
<tr>  <td><a class="anchor" name="Master.RoleId"></a> Master.RoleId</td>  <td>GAUGE</td>  <td>Display master role id</td></tr>
<tr>  <td><a class="anchor" name="Master.RpcQueueLength"></a> Master.RpcQueueLength</td>  <td>GAUGE</td>  <td>Length of the master rpc queue. Use this metric to monitor the RPC pressure on master.</td></tr>
<tr>  <td><a class="anchor" name="Master.RpcThreadActiveCount"></a> Master.RpcThreadActiveCount</td>  <td>GAUGE</td>  <td>The number of threads that are actively executing tasks in the master RPC executor thread pool. Use this metric to monitor the RPC pressure on master.</td></tr>
<tr>  <td><a class="anchor" name="Master.RpcThreadCurrentCount"></a> Master.RpcThreadCurrentCount</td>  <td>GAUGE</td>  <td>Current count of threads in the master RPC executor thread pool. Use this metric to monitor the RPC pressure on master.</td></tr>
<tr>  <td><a class="anchor" name="Master.SetAclOps"></a> Master.SetAclOps</td>  <td>COUNTER</td>  <td>Total number of SetAcl operations</td></tr>
<tr>  <td><a class="anchor" name="Master.SetAttributeOps"></a> Master.SetAttributeOps</td>  <td>COUNTER</td>  <td>Total number of SetAttribute operations</td></tr>
<tr>  <td><a class="anchor" name="Master.StartTime"></a> Master.StartTime</td>  <td>GAUGE</td>  <td>The start time of the master process</td></tr>
<tr>  <td><a class="anchor" name="Master.TTLBuckets"></a> Master.TTLBuckets</td>  <td>GAUGE</td>  <td>The number of TTL buckets at the master. Note that these buckets are stored in memory.</td></tr>
<tr>  <td><a class="anchor" name="Master.TTLInodes"></a> Master.TTLInodes</td>  <td>GAUGE</td>  <td>The total number of inodes contained in TTL buckets at the mater. Note that these inodes are stored in memory.</td></tr>
<tr>  <td><a class="anchor" name="Master.ToRemoveBlockCount"></a> Master.ToRemoveBlockCount</td>  <td>GAUGE</td>  <td>Count of block replicas to be removed from the workers. If 1 block is to be removed from 2 workers, 2 will be counted here.</td></tr>
<tr>  <td><a class="anchor" name="Master.TotalPaths"></a> Master.TotalPaths</td>  <td>GAUGE</td>  <td>Total number of files and directory in Alluxio namespace</td></tr>
<tr>  <td><a class="anchor" name="Master.TotalRpcs"></a> Master.TotalRpcs</td>  <td>TIMER</td>  <td>Throughput of master RPC calls. This metrics indicates how busy the master is serving client and worker requests</td></tr>
<tr>  <td><a class="anchor" name="Master.UfsJournalCatchupTimer"></a> Master.UfsJournalCatchupTimer</td>  <td>TIMER</td>  <td>The timer statistics of journal catchupOnly valid when ufs journal is used. This provides a summary of how long a standby master takes to catch up with primary master, and should be monitored if master transition takes too long</td></tr>
<tr>  <td><a class="anchor" name="Master.UfsJournalFailureRecoverTimer"></a> Master.UfsJournalFailureRecoverTimer</td>  <td>TIMER</td>  <td>The timer statistics of ufs journal failure recover</td></tr>
<tr>  <td><a class="anchor" name="Master.UfsJournalInitialReplayTimeMs"></a> Master.UfsJournalInitialReplayTimeMs</td>  <td>GAUGE</td>  <td>The process time of the ufs journal initial replay.Only valid when ufs journal is used. It records the time it took for the very first journal replay. Use this metric to monitor when your master boot-up time is high。</td></tr>
<tr>  <td><a class="anchor" name="Master.UfsStatusCacheChildrenSize"></a> Master.UfsStatusCacheChildrenSize</td>  <td>COUNTER</td>  <td>Total number of UFS file metadata cached. The cache is used during metadata sync.</td></tr>
<tr>  <td><a class="anchor" name="Master.UfsStatusCacheSize"></a> Master.UfsStatusCacheSize</td>  <td>COUNTER</td>  <td>Total number of Alluxio paths being processed by the metadata sync prefetch thread pool.</td></tr>
<tr>  <td><a class="anchor" name="Master.UniqueBlocks"></a> Master.UniqueBlocks</td>  <td>GAUGE</td>  <td>Total number of unique blocks in Alluxio</td></tr>
<tr>  <td><a class="anchor" name="Master.UnmountOps"></a> Master.UnmountOps</td>  <td>COUNTER</td>  <td>Total number of Unmount operations</td></tr>
</table>

Dynamically generated master metrics:

| Metric Name | Description |
|-------------------------|-----------------------------------------------------|
| Master.CapacityTotalTier{TIER_NAME} | Total capacity in tier {TIER_NAME} of the Alluxio file system in bytes |
| Master.CapacityUsedTier{TIER_NAME}  | Used capacity in tier {TIER_NAME} of the Alluxio file system in bytes |
| Master.CapacityFreeTier{TIER_NAME}  | Free capacity in tier {TIER_NAME} of the Alluxio file system in bytes |
| Master.UfsSessionCount-Ufs:{UFS_ADDRESS} | The total number of currently opened UFS sessions to connect to the given {UFS_ADDRESS} |
| Master.{UFS_RPC_NAME}.UFS:{UFS_ADDRESS}.UFS_TYPE:{UFS_TYPE}.User:{USER} | The details UFS rpc operation done by the current master |
| Master.PerUfsOp{UFS_RPC_NAME}.UFS:{UFS_ADDRESS} | The aggregated number of UFS operation {UFS_RPC_NAME} ran on UFS {UFS_ADDRESS} by leading master |  
| Master.{LEADING_MASTER_RPC_NAME} | The duration statistics of RPC calls exposed on leading master |

## Worker Metrics

Default worker metrics:

<table class="table table-striped">
<tbody><tr><th>Name</th><th>Type</th><th>Description</th></tr>
<tr>  <td><a class="anchor" name="Worker.ActiveClients"></a> Worker.ActiveClients</td>  <td>COUNTER</td>  <td>The number of clients actively reading from or writing to this worker</td></tr>
<tr>  <td><a class="anchor" name="Worker.ActiveRpcReadCount"></a> Worker.ActiveRpcReadCount</td>  <td>COUNTER</td>  <td>The number of active read-RPCs managed by this worker</td></tr>
<tr>  <td><a class="anchor" name="Worker.ActiveRpcWriteCount"></a> Worker.ActiveRpcWriteCount</td>  <td>COUNTER</td>  <td>The number of active write-RPCs managed by this worker</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockReaderCompleteTaskCount"></a> Worker.BlockReaderCompleteTaskCount</td>  <td>GAUGE</td>  <td>The approximate total number of block read tasks that have completed execution</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockReaderThreadActiveCount"></a> Worker.BlockReaderThreadActiveCount</td>  <td>GAUGE</td>  <td>The approximate number of block read threads that are actively executing tasks in reader thread pool</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockReaderThreadCurrentCount"></a> Worker.BlockReaderThreadCurrentCount</td>  <td>GAUGE</td>  <td>The current number of read threads in the reader thread pool</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockReaderThreadMaxCount"></a> Worker.BlockReaderThreadMaxCount</td>  <td>GAUGE</td>  <td>The maximum allowed number of block read thread in the reader thread pool</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockRemoverBlocksRemovedCount"></a> Worker.BlockRemoverBlocksRemovedCount</td>  <td>COUNTER</td>  <td>The total number of blocks successfully removed from this worker by asynchronous block remover.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockRemoverRemovingBlocksSize"></a> Worker.BlockRemoverRemovingBlocksSize</td>  <td>GAUGE</td>  <td>The size of blocks is being removed from this worker at a moment by asynchronous block remover.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockRemoverTryRemoveBlocksSize"></a> Worker.BlockRemoverTryRemoveBlocksSize</td>  <td>GAUGE</td>  <td>The number of blocks to be removed from this worker at a moment by asynchronous block remover.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockRemoverTryRemoveCount"></a> Worker.BlockRemoverTryRemoveCount</td>  <td>COUNTER</td>  <td>The total number of blocks this worker attempted to remove with asynchronous block remover.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockSerializedCompleteTaskCount"></a> Worker.BlockSerializedCompleteTaskCount</td>  <td>GAUGE</td>  <td>The approximate total number of block serialized tasks that have completed execution</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockSerializedThreadActiveCount"></a> Worker.BlockSerializedThreadActiveCount</td>  <td>GAUGE</td>  <td>The approximate number of block serialized threads that are actively executing tasks in serialized thread pool</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockSerializedThreadCurrentCount"></a> Worker.BlockSerializedThreadCurrentCount</td>  <td>GAUGE</td>  <td>The current number of serialized threads in the serialized thread pool</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockSerializedThreadMaxCount"></a> Worker.BlockSerializedThreadMaxCount</td>  <td>GAUGE</td>  <td>The maximum allowed number of block serialized thread in the serialized thread pool</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockWriterCompleteTaskCount"></a> Worker.BlockWriterCompleteTaskCount</td>  <td>GAUGE</td>  <td>The approximate total number of block write tasks that have completed execution</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockWriterThreadActiveCount"></a> Worker.BlockWriterThreadActiveCount</td>  <td>GAUGE</td>  <td>The approximate number of block write threads that are actively executing tasks in writer thread pool</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockWriterThreadCurrentCount"></a> Worker.BlockWriterThreadCurrentCount</td>  <td>GAUGE</td>  <td>The current number of write threads in the writer thread pool</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockWriterThreadMaxCount"></a> Worker.BlockWriterThreadMaxCount</td>  <td>GAUGE</td>  <td>The maximum allowed number of block write thread in the writer thread pool</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksAccessed"></a> Worker.BlocksAccessed</td>  <td>COUNTER</td>  <td>Total number of times any one of the blocks in this worker is accessed.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksCached"></a> Worker.BlocksCached</td>  <td>GAUGE</td>  <td>Total number of blocks used for caching data in an Alluxio worker</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksCancelled"></a> Worker.BlocksCancelled</td>  <td>COUNTER</td>  <td>Total number of aborted temporary blocks in this worker.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksDeleted"></a> Worker.BlocksDeleted</td>  <td>COUNTER</td>  <td>Total number of deleted blocks in this worker by external requests.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksEvicted"></a> Worker.BlocksEvicted</td>  <td>COUNTER</td>  <td>Total number of evicted blocks in this worker.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksEvictionRate"></a> Worker.BlocksEvictionRate</td>  <td>METER</td>  <td>Block eviction rate in this worker.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksLost"></a> Worker.BlocksLost</td>  <td>COUNTER</td>  <td>Total number of lost blocks in this worker.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksPromoted"></a> Worker.BlocksPromoted</td>  <td>COUNTER</td>  <td>Total number of times any one of the blocks in this worker moved to a new tier.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksReadLocal"></a> Worker.BlocksReadLocal</td>  <td>COUNTER</td>  <td>Total number of local blocks read by this worker.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksReadRemote"></a> Worker.BlocksReadRemote</td>  <td>COUNTER</td>  <td>Total number of a remote blocks read by this worker.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksReadUfs"></a> Worker.BlocksReadUfs</td>  <td>COUNTER</td>  <td>Total number of a UFS blocks read by this worker.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesReadDirect"></a> Worker.BytesReadDirect</td>  <td>COUNTER</td>  <td>Total number of bytes read from the this worker without external RPC involved. Data exists in worker storage or is fetched by this worker from underlying UFSes. This records data read by worker internal calls (e.g. a client embedded in this worker).</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesReadDirectThroughput"></a> Worker.BytesReadDirectThroughput</td>  <td>METER</td>  <td>Throughput of bytes read from the this worker without external RPC involved. Data exists in worker storage or is fetched by this worker from underlying UFSes. This records data read by worker internal calls (e.g. a client embedded in this worker).</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesReadDomain"></a> Worker.BytesReadDomain</td>  <td>COUNTER</td>  <td>Total number of bytes read from the this worker via domain socket</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesReadDomainThroughput"></a> Worker.BytesReadDomainThroughput</td>  <td>METER</td>  <td>Bytes read throughput from the this worker via domain socket</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesReadPerUfs"></a> Worker.BytesReadPerUfs</td>  <td>COUNTER</td>  <td>Total number of bytes read from a specific Alluxio UFS by this worker</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesReadRemote"></a> Worker.BytesReadRemote</td>  <td>COUNTER</td>  <td>Total number of bytes read from the this worker via network (RPC). Data exists in worker storage or is fetched by this worker from underlying UFSes. This does not include short-circuit local reads and domain socket reads.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesReadRemoteThroughput"></a> Worker.BytesReadRemoteThroughput</td>  <td>METER</td>  <td>Throughput of bytes read from the this worker via network (RPC). Data exists in worker storage or is fetched by this worker from underlying UFSes. This does not include short-circuit local reads and domain socket reads</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesReadUfsThroughput"></a> Worker.BytesReadUfsThroughput</td>  <td>METER</td>  <td>Bytes read throughput from all Alluxio UFSes by this worker</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesWrittenDirect"></a> Worker.BytesWrittenDirect</td>  <td>COUNTER</td>  <td>Total number of bytes written to this worker without external RPC involved. Data is written to worker storage or is written by this worker to underlying UFSes. This records data written by worker internal calls (e.g. a client embedded in this worker).</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesWrittenDirectThroughput"></a> Worker.BytesWrittenDirectThroughput</td>  <td>METER</td>  <td>Total number of bytes written to this worker without external RPC involved. Data is written to worker storage or is written by this worker to underlying UFSes. This records data written by worker internal calls (e.g. a client embedded in this worker).</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesWrittenDomain"></a> Worker.BytesWrittenDomain</td>  <td>COUNTER</td>  <td>Total number of bytes written to this worker via domain socket</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesWrittenDomainThroughput"></a> Worker.BytesWrittenDomainThroughput</td>  <td>METER</td>  <td>Throughput of bytes written to this worker via domain socket</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesWrittenPerUfs"></a> Worker.BytesWrittenPerUfs</td>  <td>COUNTER</td>  <td>Total number of bytes written to a specific Alluxio UFS by this worker</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesWrittenRemote"></a> Worker.BytesWrittenRemote</td>  <td>COUNTER</td>  <td>Total number of bytes written to this worker via network (RPC). Data is written to worker storage or is written by this worker to underlying UFSes. This does not include short-circuit local writes and domain socket writes.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesWrittenRemoteThroughput"></a> Worker.BytesWrittenRemoteThroughput</td>  <td>METER</td>  <td>Bytes write throughput to this worker via network (RPC). Data is written to worker storage or is written by this worker to underlying UFSes. This does not include short-circuit local writes and domain socket writes.</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesWrittenUfsThroughput"></a> Worker.BytesWrittenUfsThroughput</td>  <td>METER</td>  <td>Bytes write throughput to all Alluxio UFSes by this worker</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheBlocksSize"></a> Worker.CacheBlocksSize</td>  <td>COUNTER</td>  <td>Total number of bytes that being cached through cache requests</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheFailedBlocks"></a> Worker.CacheFailedBlocks</td>  <td>COUNTER</td>  <td>Total number of failed cache blocks in this worker</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheManagerCompleteTaskCount"></a> Worker.CacheManagerCompleteTaskCount</td>  <td>GAUGE</td>  <td>The approximate total number of block cache tasks that have completed execution</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheManagerThreadActiveCount"></a> Worker.CacheManagerThreadActiveCount</td>  <td>GAUGE</td>  <td>The approximate number of block cache threads that are actively executing tasks in the cache manager thread pool</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheManagerThreadCurrentCount"></a> Worker.CacheManagerThreadCurrentCount</td>  <td>GAUGE</td>  <td>The current number of cache threads in the cache manager thread pool</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheManagerThreadMaxCount"></a> Worker.CacheManagerThreadMaxCount</td>  <td>GAUGE</td>  <td>The maximum allowed number of block cache thread in the cache manager thread pool</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheManagerThreadQueueWaitingTaskCount"></a> Worker.CacheManagerThreadQueueWaitingTaskCount</td>  <td>GAUGE</td>  <td>The current number of tasks waiting in the work queue in the cache manager thread pool, bounded by alluxio.worker.network.async.cache.manager.queue.max</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheRemoteBlocks"></a> Worker.CacheRemoteBlocks</td>  <td>COUNTER</td>  <td>Total number of blocks that need to be cached from remote source</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheRequests"></a> Worker.CacheRequests</td>  <td>COUNTER</td>  <td>Total number of cache request received by this worker</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheRequestsAsync"></a> Worker.CacheRequestsAsync</td>  <td>COUNTER</td>  <td>Total number of async cache request received by this worker</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheRequestsSync"></a> Worker.CacheRequestsSync</td>  <td>COUNTER</td>  <td>Total number of sync cache request received by this worker</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheSucceededBlocks"></a> Worker.CacheSucceededBlocks</td>  <td>COUNTER</td>  <td>Total number of cache succeeded blocks in this worker</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheUfsBlocks"></a> Worker.CacheUfsBlocks</td>  <td>COUNTER</td>  <td>Total number of blocks that need to be cached from local source</td></tr>
<tr>  <td><a class="anchor" name="Worker.CapacityFree"></a> Worker.CapacityFree</td>  <td>GAUGE</td>  <td>Total free bytes on all tiers of a specific Alluxio worker</td></tr>
<tr>  <td><a class="anchor" name="Worker.CapacityTotal"></a> Worker.CapacityTotal</td>  <td>GAUGE</td>  <td>Total capacity (in bytes) on all tiers of a specific Alluxio worker</td></tr>
<tr>  <td><a class="anchor" name="Worker.CapacityUsed"></a> Worker.CapacityUsed</td>  <td>GAUGE</td>  <td>Total used bytes on all tiers of a specific Alluxio worker</td></tr>
<tr>  <td><a class="anchor" name="Worker.MasterRegistrationSuccessCount"></a> Worker.MasterRegistrationSuccessCount</td>  <td>COUNTER</td>  <td>Total number of the succeed master registration.</td></tr>
<tr>  <td><a class="anchor" name="Worker.RpcQueueLength"></a> Worker.RpcQueueLength</td>  <td>GAUGE</td>  <td>Length of the worker rpc queue. Use this metric to monitor the RPC pressure on worker.</td></tr>
<tr>  <td><a class="anchor" name="Worker.RpcThreadActiveCount"></a> Worker.RpcThreadActiveCount</td>  <td>GAUGE</td>  <td>The number of threads that are actively executing tasks in the worker RPC executor thread pool. Use this metric to monitor the RPC pressure on worker.</td></tr>
<tr>  <td><a class="anchor" name="Worker.RpcThreadCurrentCount"></a> Worker.RpcThreadCurrentCount</td>  <td>GAUGE</td>  <td>Current count of threads in the worker RPC executor thread pool. Use this metric to monitor the RPC pressure on worker.</td></tr>

</tbody></table>

Dynamically generated worker metrics:

| Metric Name | Description |
|-------------------------|-----------------------------------------------------|
| Worker.UfsSessionCount-Ufs:{UFS_ADDRESS} | The total number of currently opened UFS sessions to connect to the given {UFS_ADDRESS} |
| Worker.{RPC_NAME}                        | The duration statistics of RPC calls exposed on workers |

## Client Metrics

Each client metric will be recorded with its local hostname or `alluxio.user.app.id` is configured.
If `alluxio.user.app.id` is configured, multiple clients can be combined into a logical application.

<table class="table table-striped">
<tbody><tr><th>Name</th><th>Type</th><th>Description</th></tr>
<tr>  <td><a class="anchor" name="Client.BlockMasterClientCount"></a> Client.BlockMasterClientCount</td>  <td>COUNTER</td>  <td>Number of instances in the BlockMasterClientPool.</td></tr>
<tr>  <td><a class="anchor" name="Client.BlockReadChunkRemote"></a> Client.BlockReadChunkRemote</td>  <td>TIMER</td>  <td>The timer statistics of reading block data in chunks from remote Alluxio workers via RPC framework. This metrics will only be recorded when alluxio.user.block.read.metrics.enabled is set to true</td></tr>
<tr>  <td><a class="anchor" name="Client.BlockWorkerClientCount"></a> Client.BlockWorkerClientCount</td>  <td>COUNTER</td>  <td>Number of instances in the BlockWorkerClientPool.</td></tr>
<tr>  <td><a class="anchor" name="Client.BusyExceptionCount"></a> Client.BusyExceptionCount</td>  <td>COUNTER</td>  <td>Total number of BusyException observed</td></tr>
<tr>  <td><a class="anchor" name="Client.BytesReadLocal"></a> Client.BytesReadLocal</td>  <td>COUNTER</td>  <td>Total number of bytes short-circuit read from worker data storage that collocates with the client</td></tr>
<tr>  <td><a class="anchor" name="Client.BytesReadLocalThroughput"></a> Client.BytesReadLocalThroughput</td>  <td>METER</td>  <td>Bytes throughput short-circuit read from worker data storage that collocated with this client</td></tr>
<tr>  <td><a class="anchor" name="Client.BytesWrittenLocal"></a> Client.BytesWrittenLocal</td>  <td>COUNTER</td>  <td>Total number of bytes short-circuit written to local storage by this client</td></tr>
<tr>  <td><a class="anchor" name="Client.BytesWrittenLocalThroughput"></a> Client.BytesWrittenLocalThroughput</td>  <td>METER</td>  <td>Bytes throughput short-circuit written to local storage by this client</td></tr>
<tr>  <td><a class="anchor" name="Client.BytesWrittenUfs"></a> Client.BytesWrittenUfs</td>  <td>COUNTER</td>  <td>Total number of bytes write to Alluxio UFS by this client</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheBytesDiscarded"></a> Client.CacheBytesDiscarded</td>  <td>METER</td>  <td>Total number of bytes discarded when restoring the page store.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheBytesEvicted"></a> Client.CacheBytesEvicted</td>  <td>METER</td>  <td>Total number of bytes evicted from the client cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheBytesReadCache"></a> Client.CacheBytesReadCache</td>  <td>METER</td>  <td>Total number of bytes read from the client cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheBytesReadExternal"></a> Client.CacheBytesReadExternal</td>  <td>METER</td>  <td>Total number of bytes read from external storage due to a cache miss on the client cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheBytesReadInStreamBuffer"></a> Client.CacheBytesReadInStreamBuffer</td>  <td>METER</td>  <td>Total number of bytes read from the client cache's in stream buffer.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheBytesRequestedExternal"></a> Client.CacheBytesRequestedExternal</td>  <td>METER</td>  <td>Total number of bytes the user requested to read which resulted in a cache miss. This number may be smaller than Client.CacheBytesReadExternal due to chunk reads.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheBytesWrittenCache"></a> Client.CacheBytesWrittenCache</td>  <td>METER</td>  <td>Total number of bytes written to the client cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheCleanErrors"></a> Client.CacheCleanErrors</td>  <td>COUNTER</td>  <td>Number of failures when cleaning out the existing cache directory to initialize a new cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheCleanupGetErrors"></a> Client.CacheCleanupGetErrors</td>  <td>COUNTER</td>  <td>Number of failures when cleaning up a failed cache read.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheCleanupPutErrors"></a> Client.CacheCleanupPutErrors</td>  <td>COUNTER</td>  <td>Number of failures when cleaning up a failed cache write.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheCreateErrors"></a> Client.CacheCreateErrors</td>  <td>COUNTER</td>  <td>Number of failures when creating a cache in the client cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheDeleteErrors"></a> Client.CacheDeleteErrors</td>  <td>COUNTER</td>  <td>Number of failures when deleting cached data in the client cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheDeleteFromStoreErrors"></a> Client.CacheDeleteFromStoreErrors</td>  <td>COUNTER</td>  <td>Number of failures when deleting pages from page stores.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheDeleteNonExistingPageErrors"></a> Client.CacheDeleteNonExistingPageErrors</td>  <td>COUNTER</td>  <td>Number of failures when deleting pages due to absence.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheDeleteNotReadyErrors"></a> Client.CacheDeleteNotReadyErrors</td>  <td>COUNTER</td>  <td>Number of failures when cache is not ready to delete pages.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheGetErrors"></a> Client.CacheGetErrors</td>  <td>COUNTER</td>  <td>Number of failures when getting cached data in the client cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheGetNotReadyErrors"></a> Client.CacheGetNotReadyErrors</td>  <td>COUNTER</td>  <td>Number of failures when cache is not ready to get pages.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheGetStoreReadErrors"></a> Client.CacheGetStoreReadErrors</td>  <td>COUNTER</td>  <td>Number of failures when getting cached data in the client cache due to failed read from page stores.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheHitRate"></a> Client.CacheHitRate</td>  <td>GAUGE</td>  <td>Cache hit rate: (# bytes read from cache) / (# bytes requested).</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePageReadCacheTimeNanos"></a> Client.CachePageReadCacheTimeNanos</td>  <td>METER</td>  <td>Time in nanoseconds taken to read a page from the client cache when the cache hits.</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePageReadExternalTimeNanos"></a> Client.CachePageReadExternalTimeNanos</td>  <td>METER</td>  <td>Time in nanoseconds taken to read a page from external source when the cache misses.</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePages"></a> Client.CachePages</td>  <td>COUNTER</td>  <td>Total number of pages in the client cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePagesDiscarded"></a> Client.CachePagesDiscarded</td>  <td>METER</td>  <td>Total number of pages discarded when restoring the page store.</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePagesEvicted"></a> Client.CachePagesEvicted</td>  <td>METER</td>  <td>Total number of pages evicted from the client cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutAsyncRejectionErrors"></a> Client.CachePutAsyncRejectionErrors</td>  <td>COUNTER</td>  <td>Number of failures when putting cached data in the client cache due to failed injection to async write queue.</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutBenignRacingErrors"></a> Client.CachePutBenignRacingErrors</td>  <td>COUNTER</td>  <td>Number of failures when adding pages due to racing eviction. This error is benign.</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutErrors"></a> Client.CachePutErrors</td>  <td>COUNTER</td>  <td>Number of failures when putting cached data in the client cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutEvictionErrors"></a> Client.CachePutEvictionErrors</td>  <td>COUNTER</td>  <td>Number of failures when putting cached data in the client cache due to failed eviction.</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutInsufficientSpaceErrors"></a> Client.CachePutInsufficientSpaceErrors</td>  <td>COUNTER</td>  <td>Number of failures when putting cached data in the client cache due to insufficient space made after eviction.</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutNotReadyErrors"></a> Client.CachePutNotReadyErrors</td>  <td>COUNTER</td>  <td>Number of failures when cache is not ready to add pages.</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutStoreDeleteErrors"></a> Client.CachePutStoreDeleteErrors</td>  <td>COUNTER</td>  <td>Number of failures when putting cached data in the client cache due to failed deletes in page store.</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutStoreWriteErrors"></a> Client.CachePutStoreWriteErrors</td>  <td>COUNTER</td>  <td>Number of failures when putting cached data in the client cache due to failed writes to page store.</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutStoreWriteNoSpaceErrors"></a> Client.CachePutStoreWriteNoSpaceErrors</td>  <td>COUNTER</td>  <td>Number of failures when putting cached data in the client cache but getting disk is full while cache capacity is not achieved. This can happen if the storage overhead ratio to write data is underestimated.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheShadowCacheBytes"></a> Client.CacheShadowCacheBytes</td>  <td>COUNTER</td>  <td>Amount of bytes in the client shadow cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheShadowCacheBytesHit"></a> Client.CacheShadowCacheBytesHit</td>  <td>COUNTER</td>  <td>Total number of bytes hit the client shadow cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheShadowCacheBytesRead"></a> Client.CacheShadowCacheBytesRead</td>  <td>COUNTER</td>  <td>Total number of bytes read from the client shadow cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheShadowCacheFalsePositiveRatio"></a> Client.CacheShadowCacheFalsePositiveRatio</td>  <td>COUNTER</td>  <td>Probability that the working set bloom filter makes an error. The value is 0-100. If too high, need to allocate more space</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheShadowCachePages"></a> Client.CacheShadowCachePages</td>  <td>COUNTER</td>  <td>Amount of pages in the client shadow cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheShadowCachePagesHit"></a> Client.CacheShadowCachePagesHit</td>  <td>COUNTER</td>  <td>Total number of pages hit the client shadow cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheShadowCachePagesRead"></a> Client.CacheShadowCachePagesRead</td>  <td>COUNTER</td>  <td>Total number of pages read from the client shadow cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheSpaceAvailable"></a> Client.CacheSpaceAvailable</td>  <td>GAUGE</td>  <td>Amount of bytes available in the client cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheSpaceUsed"></a> Client.CacheSpaceUsed</td>  <td>GAUGE</td>  <td>Amount of bytes used by the client cache.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheSpaceUsedCount"></a> Client.CacheSpaceUsedCount</td>  <td>COUNTER</td>  <td>Amount of bytes used by the client cache as a counter.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheState"></a> Client.CacheState</td>  <td>COUNTER</td>  <td>State of the cache: 0 (NOT_IN_USE), 1 (READ_ONLY) and 2 (READ_WRITE)</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheStoreDeleteTimeout"></a> Client.CacheStoreDeleteTimeout</td>  <td>COUNTER</td>  <td>Number of timeouts when deleting pages from page store.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheStoreGetTimeout"></a> Client.CacheStoreGetTimeout</td>  <td>COUNTER</td>  <td>Number of timeouts when reading pages from page store.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheStorePutTimeout"></a> Client.CacheStorePutTimeout</td>  <td>COUNTER</td>  <td>Number of timeouts when writing new pages to page store.</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheStoreThreadsRejected"></a> Client.CacheStoreThreadsRejected</td>  <td>COUNTER</td>  <td>Number of rejection of I/O threads on submitting tasks to thread pool, likely due to unresponsive local file system.</td></tr>
<tr>  <td><a class="anchor" name="Client.CloseAlluxioOutStreamLatency"></a> Client.CloseAlluxioOutStreamLatency</td>  <td>TIMER</td>  <td>Latency of close Alluxio outstream latency</td></tr>
<tr>  <td><a class="anchor" name="Client.CloseUFSOutStreamLatency"></a> Client.CloseUFSOutStreamLatency</td>  <td>TIMER</td>  <td>Latency of close UFS outstream latency</td></tr>
<tr>  <td><a class="anchor" name="Client.DefaultHiveClientCount"></a> Client.DefaultHiveClientCount</td>  <td>COUNTER</td>  <td>Number of instances in the DefaultHiveClientPool.</td></tr>
<tr>  <td><a class="anchor" name="Client.FileSystemMasterClientCount"></a> Client.FileSystemMasterClientCount</td>  <td>COUNTER</td>  <td>Number of instances in the FileSystemMasterClientPool.</td></tr>
<tr>  <td><a class="anchor" name="Client.MetadataCacheSize"></a> Client.MetadataCacheSize</td>  <td>GAUGE</td>  <td>The total number of files and directories whose metadata is cached on the client-side. Only valid if the filesystem is alluxio.client.file.MetadataCachingBaseFileSystem.</td></tr>

</tbody></table>

## Fuse Metrics

Fuse is a long-running Alluxio client. 
Depending on the launching ways, Fuse metrics show as
* client metrics when Fuse client is launching in a standalone AlluxioFuse process.
* worker metrics when Fuse client is embedded in the AlluxioWorker process.

Fuse metrics includes:

<table class="table table-striped">
<tbody><tr><th>Name</th><th>Type</th><th>Description</th></tr>
<tr>  <td><a class="anchor" name="Fuse.CachedPathCount"></a> Fuse.CachedPathCount</td>  <td>GAUGE</td>  <td>Total number of FUSE-to-Alluxio path mappings being cached. This value will be smaller or equal to alluxio.fuse.cached.paths.max</td></tr>
<tr>  <td><a class="anchor" name="Fuse.ReadWriteFileCount"></a> Fuse.ReadWriteFileCount</td>  <td>GAUGE</td>  <td>Total number of files being opened for reading or writing concurrently.</td></tr>
<tr>  <td><a class="anchor" name="Fuse.TotalCalls"></a> Fuse.TotalCalls</td>  <td>TIMER</td>  <td>Throughput of JNI FUSE operation calls. This metrics indicates how busy the Alluxio Fuse application is serving requests</td></tr>

</tbody></table>

Fuse reading/writing file count can be used as the indicators for Fuse application pressure.
If a large amount of concurrent read/write occur in a short period of time, each of the read/write operations may take longer time to finish.

When a user or an application runs a filesystem command under Fuse mount point, 
this command will be processed and translated by operating system which will trigger the related Fuse operations
exposed in [AlluxioFuse](https://github.com/Alluxio/alluxio/blob/db01aae966849e88d342a71609ab3d910215afeb/integration/fuse/src/main/java/alluxio/fuse/AlluxioJniFuseFileSystem.java).
The count of how many times each operation is called, and the duration of each call will be recorded with metrics name `Fuse.<FUSE_OPERATION_NAME>` dynamically.

The important Fuse metrics include:

| Metric Name | Description |
|-------------------------|-----------------------------------------------------|
| Fuse.readdir | The duration metrics of listing a directory |
| Fuse.getattr | The duration metrics of getting the metadata of a file |
| Fuse.open | The duration metrics of opening a file for read or overwrite |
| Fuse.read | The duration metrics of reading a part of a file |
| Fuse.create | The duration metrics of creating a file for write |
| Fuse.write | The duration metrics of writing a file |
| Fuse.release | The duration metrics of closing a file after read or write. Note that release is async so fuse threads will not wait for release to finish |
| Fuse.mkdir | The duration metrics of creating a directory |
| Fuse.unlink | The duration metrics of removing a file or a directory |
| Fuse.rename | The duration metrics of renaming a file or a directory |
| Fuse.chmod | The duration metrics of modifying the mode of a file or a directory |
| Fuse.chown | The duration metrics of modifying the user and/or group ownership of a file or a directory |

Fuse related metrics include:
* `Client.TotalRPCClients`shows the total number of RPC clients exist that is using to or can be used to connect to master or worker for operations.
* Worker metrics with `Direct` keyword. When Fuse is embedded in worker process, it can go through worker internal API to read from / write to this worker.
The related metrics are ended with `Direct`. For example, `Worker.BytesReadDirect` shows how many bytes are served by this worker to its embedded Fuse client for read.
* If `alluxio.user.block.read.metrics.enabled=true` is configured, `Client.BlockReadChunkRemote` will be recorded. 
This metric shows the duration statistics of reading data from remote workers via gRPC.

`Client.TotalRPCClients` and `Fuse.TotalCalls` metrics are good indicator of the current load of the Fuse applications.
If applications (e.g. Tensorflow) are running on top of Alluxio Fuse but these two metrics show a much lower value than before,
the training job may be stuck with Alluxio.

## Process Common Metrics

The following metrics are collected on each instance (Master, Worker or Client).

### JVM Attributes

| Metric Name | Description |
|-------------------------|-----------------------------------------------------|
| name | The name of the JVM |
| uptime | The uptime of the JVM |
| vendor | The current JVM vendor |

### Garbage Collector Statistics

| Metric Name | Description |
|-------------------------|-----------------------------------------------------|
| PS-MarkSweep.count | Total number of mark and sweep |
| PS-MarkSweep.time | The time used to mark and sweep |
| PS-Scavenge.count | Total number of scavenge |
| PS-Scavenge.time | The time used to scavenge |

### Memory Usage

Alluxio provides overall and detailed memory usage information.
Detailed memory usage information of code cache, compressed class space, metaspace, PS Eden space, PS old gen, and PS survivor space
is collected in each process.

A subset of the memory usage metrics are listed as following:

| Metric Name | Description |
|------------------------------|-----------------------------------------------------|
| total.committed | The amount of memory in bytes that is guaranteed to be available for use by the JVM |
| total.init | The amount of the memory in bytes that is available for use by the JVM |
| total.max | The maximum amount of memory in bytes that is available for use by the JVM |
| total.used | The amount of memory currently used in bytes |
| heap.committed | The amount of memory from heap area guaranteed to be available |
| heap.init | The amount of memory from heap area available at initialization |
| heap.max | The maximum amount of memory from heap area that is available |
| heap.usage | The amount of memory from heap area currently used in GB|
| heap.used | The amount of memory from heap area that has been used |
| pools.Code-Cache.used | Used memory of collection usage from the pool from which memory is used for compilation and storage of native code |
| pools.Compressed-Class-Space.used | Used memory of collection usage from the pool from which memory is use for class metadata |
| pools.PS-Eden-Space.used | Used memory of collection usage from the pool from which memory is initially allocated for most objects |
| pools.PS-Survivor-Space.used | Used memory of collection usage from the pool containing objects that have survived the garbage collection of the Eden space |

### ClassLoading Statistics

| Metric Name | Description |
|-------------------------|-----------------------------------------------------|
| loaded | The total number of classes loaded |
| unloaded | The total number of unloaded classes |

### Thread Statistics

| Metric Name | Description |
|-------------------------|-----------------------------------------------------|
| count | The current number of live threads |
| daemon.count | The current number of live daemon threads |
| peak.count | The peak live thread count |
| total_started.count | The total number of threads started |
| deadlock.count | The number of deadlocked threads |
| deadlock | The call stack of each thread related deadlock |
| new.count | The number of threads with new state |
| blocked.count | The number of threads with blocked state |
| runnable.count | The number of threads with runnable state |
| terminated.count | The number of threads with terminated state |
| timed_waiting.count | The number of threads with timed_waiting state |
