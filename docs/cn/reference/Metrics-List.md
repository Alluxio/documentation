# List of Metrics

在 Alluxio 中，有两种类型的指标，集群范围内的合计指标和每个进程的详细指标。

* 集群指标由 leading master 收集和计算的，并且在 web UI 下的指标标签下展示。
  这些指标旨在提供 Alluxio 服务的集群状态以及数据与元数据总量的快照。

* 进程指标由每个 Alluxio 进程收集，并通过任何配置的接收器以机器可读的格式暴露出来。
  进程指标高度详细，旨在被第三方监测工具使用。
  用户可以通过细粒度的数据面板查看每个指标的时间序列图。
  比如数据传输量或 RPC 调用次数。

Alluxio 的 master 节点指标具有以下格式：

```
Master.[metricName].[tag1].[tag2]...
```

Alluxio 的非 master 节点指标具有以下格式

```
[processType].[metricName].[tag1].[tag2]...[hostName] 
```

通常情况下，Alluxio 会为每一种 RPC 调用生成一个指标，无论是调用 Alluxio 还是调用下层存储。

标签是指标的附加元数据，如用户名或存储位置。
标签可用于进一步筛选或聚合各种特征。

## 集群指标

Worker 和 client 通过心跳包将指标数据发送到 Alluxio master。心跳间隔分别由 `alluxio.master.worker.heartbeat.interval` 和 `alluxio.user.metrics.heartbeat.interval` 属性定义。

字节指标是来自 worker 或 client 的聚合值。字节吞吐量指标是在 master 上计算的。
字节吞吐量的值等于字节指标计数器值除以指标记录时间，并以字节/分钟的形式呈现。

<table class="table table-striped">
<tbody>
  <tr><th>名称</th><th>类型</th><th>描述</th></tr>
  <tr><td>Cluster.ActiveRpcReadCount</td><td>COUNTER</td><td>worker 上进行中的 read-RPC 数量</td></tr>
  <tr><td>Cluster.ActiveRpcWriteCount</td><td>COUNTER</td><td>worker 上进行中的 write-RPC 数量</td></tr>
  <tr><td>Cluster.BytesReadDirect</td><td>COUNTER</td><td>汇总在所有 worker 上不通过 RPC 读取的字节数。这记录了 worker 内部调用（e.g. 嵌入在 worker 中的客户端）读取的数据，数据存在于 worker 缓存中或由 worker 从 UFS 获取</td></tr>
  <tr><td>Cluster.BytesReadDirectThroughput</td><td>GAUGE</td><td>汇总在所有 worker 上不通过 RPC 读取字节的吞吐量。这记录了 worker 内部调用（e.g. 嵌入在 worker 中的客户端）读取的数据，数据存在于 worker 缓存中或由 worker 从 UFS 获取</td></tr>
  <tr><td>Cluster.BytesReadDomain</td><td>COUNTER</td><td>从所有 worker 通过域套接字读取的总字节数</td></tr>
  <tr><td>Cluster.BytesReadDomainThroughput</td><td>GAUGE</td><td>通过域套接字从所有 worker 每分钟读取字节的吞吐量</td></tr>
  <tr><td>Cluster.BytesReadLocal</td><td>COUNTER</td><td>由所有客户端报告的短路读取的总字节数</td></tr>
  <tr><td>Cluster.BytesReadLocalThroughput</td><td>GAUGE</td><td>由所有客户端报告的每分钟短路读取字节的吞吐量</td></tr>
  <tr><td>Cluster.BytesReadPerUfs</td><td>COUNTER</td><td>所有 worker 从特定 UFS 读取的字节数总和</td></tr>
  <tr><td>Cluster.BytesReadRemote</td><td>COUNTER</td><td>从所有 worker 通过网络（RPC）读取的总字节数。数据存在于 worker 存储中，或者由 worker 从 UFS 获取。这不包括本地短路读和域套接字读</td></tr>
  <tr><td>Cluster.BytesReadRemoteThroughput</td><td>GAUGE</td><td>从所有 worker 通过网络（RPC 调用）每分钟读取的字节数吞吐量。数据存在于 worker 存储中，或者由 worker 从 UFS 获取。这不包括短路本地读取和域套接字读取</td></tr>
  <tr><td>Cluster.BytesReadUfsAll</td><td>COUNTER</td><td>所有 worker 从所有 UFS 读取的字节数总和</td></tr>
  <tr><td>Cluster.BytesReadUfsThroughput</td><td>GAUGE</td><td>所有 worker 从所有 UFS 每分钟读取的字节数吞吐量</td></tr>
  <tr><td>Cluster.BytesWrittenDomain</td><td>COUNTER</td><td>通过域套接字写入所有 worker 的字节数总和</td></tr>
  <tr><td>Cluster.BytesWrittenDomainThroughput</td><td>GAUGE</td><td>通过域套接字向所有 worker 每分钟写入字节的吞吐量</td></tr>
  <tr><td>Cluster.BytesWrittenLocal</td><td>COUNTER</td><td>所有客户端短路写入到本地 worker 数据存储的字节数总和</td></tr>
  <tr><td>Cluster.BytesWrittenLocalThroughput</td><td>GAUGE</td><td>所有客户端每分钟写入本地 worker 数据存储字节的吞吐量</td></tr>
  <tr><td>Cluster.BytesWrittenPerUfs</td><td>COUNTER</td><td>所有 worker 向特定的 Alluxio UFS 写入的字节数总和</td></tr>
  <tr><td>Cluster.BytesWrittenRemote</td><td>COUNTER</td><td>通过网络（RPC）写入 worker 的字节数总和。数据被写入 worker 存储，或者由 worker 写入底层 UFS。其中不包括短路本地写入和域套接字写入</td></tr>
  <tr><td>Cluster.BytesWrittenRemoteThroughput</td><td>GAUGE</td><td>通过网络（RPC）每分钟向 worker 写入字节的吞吐量。数据被写入 worker 存储，或者由 worker 写入底层 UFS。其中不包括短路本地写入和域套接字写入</td></tr>
  <tr><td>Cluster.BytesWrittenUfsAll</td><td>COUNTER</td><td>所有 worker 向所有 UFS 写入的字节数总和</td></tr>
  <tr><td>Cluster.BytesWrittenUfsThroughput</td><td>GAUGE</td><td>所有 worker 每分钟向所有 UFS 写入字节的吞吐量</td></tr>
  <tr><td>Cluster.CacheHitRate</td><td>GAUGE</td><td>缓存命中率：（#从缓存读取的字节数）/（#请求的字节数）</td></tr>
  <tr><td>Cluster.CapacityFree</td><td>GAUGE</td><td>Alluxio 所有 worker 上所有层的总空闲字节数</td></tr>
  <tr><td>Cluster.CapacityTotal</td><td>GAUGE</td><td>Alluxio 所有 worker 上所有层的总容量（以字节为单位）</td></tr>
  <tr><td>Cluster.CapacityUsed</td><td>GAUGE</td><td>Alluxio 所有 worker 上所有层的使用字节总数</td></tr>
  <tr><td>Cluster.LeaderId</td><td>GAUGE</td><td>展示当前 primary master id</td></tr>
  <tr><td>Cluster.LeaderIndex</td><td>GAUGE</td><td>当前 primary master 的序号</td></tr>
  <tr><td>Cluster.LostWorkers</td><td>GAUGE</td><td>集群内丢失的 worker 总数</td></tr>
  <tr><td>Cluster.RootUfsCapacityFree</td><td>GAUGE</td><td>Alluxio 根 UFS 的空闲容量（以字节为单位）</td></tr>
  <tr><td>Cluster.RootUfsCapacityTotal</td><td>GAUGE</td><td>Alluxio 根 UFS 的总容量（以字节为单位）</td></tr>
  <tr><td>Cluster.RootUfsCapacityUsed</td><td>GAUGE</td><td>Alluxio 根 UFS 的使用容量（以字节为单位）</td></tr>
  <tr><td>Cluster.Workers</td><td>GAUGE</td><td>集群内活跃的 worker 总数</td></tr>
</tbody>
</table>


## 进程指标

所有 Alluxio 服务器和客户端进程共享的指标。

<table class="table table-striped">
<tbody><tr><th>名称</th><th>类型</th><th>描述</th></tr>
  <tr>
    <td>Process.pool.direct.mem.used</td>
    <td>GAUGE</td>
    <td>NIO Direct buffer pool已使用的直接内存</td>
  </tr>
</tbody></table>

## 服务器指标

Alluxio 服务器共享的指标。

<table class="table table-striped">
<tbody>
  <tr><th>名称</th><th>类型</th><th>描述</th></tr>
  <tr><td>Server.JvmPauseMonitorInfoTimeExceeded</td><td>GAUGE</td><td>JVM 暂停时间长于 alluxio.jvm.monitor.info.threshold 阈值的总次数</td></tr>
  <tr><td>Server.JvmPauseMonitorTotalExtraTime</td><td>GAUGE</td><td>JVM 暂停的总时间，JVM暂停通常由GC或jstack等事件引发</td></tr>
  <tr><td>Server.JvmPauseMonitorWarnTimeExceeded</td><td>GAUGE</td><td>JVM 暂停时间长于 alluxio.jvm.monitor.warn.threshold 阈值的总次数</td></tr>
</tbody>
</table>


## Master 指标

默认 Master 指标:

<table class="table table-striped">
<tbody><tr><th>名称</th><th>类型</th><th>描述</th></tr>
<tr>  <td><a class="anchor" name="Master.AbsentCacheHits"></a> Master.AbsentCacheHits</td>  <td>GAUGE</td>  <td>Absent cache（记录不存在的路径）的缓存命中次数</td></tr>
<tr>  <td><a class="anchor" name="Master.AbsentCacheMisses"></a> Master.AbsentCacheMisses</td>  <td>GAUGE</td>  <td>Absent cache（记录不存在的路径）的缓存未命中次数</td></tr>
<tr>  <td><a class="anchor" name="Master.AbsentCacheSize"></a> Master.AbsentCacheSize</td>  <td>GAUGE</td>  <td>Absent cache（记录不存在的路径）的大小</td></tr>
<tr>  <td><a class="anchor" name="Master.AbsentPathCacheQueueSize"></a> Master.AbsentPathCacheQueueSize</td>  <td>GAUGE</td>  <td>Alluxio 维护了一个Absent cache（记录不存在的路径），这是正在处理的 UFS 路径数量。</td></tr>
<tr>  <td><a class="anchor" name="Master.AsyncPersistCancel"></a> Master.AsyncPersistCancel</td>  <td>COUNTER</td>  <td>已取消的 AsyncPersist 操作数量</td></tr>
<tr>  <td><a class="anchor" name="Master.AsyncPersistFail"></a> Master.AsyncPersistFail</td>  <td>COUNTER</td>  <td>失败的 AsyncPersist 操作数量</td></tr>
<tr>  <td><a class="anchor" name="Master.AsyncPersistFileCount"></a> Master.AsyncPersistFileCount</td>  <td>COUNTER</td>  <td>AsyncPersist 操作创建的文件数量</td></tr>
<tr>  <td><a class="anchor" name="Master.AsyncPersistFileSize"></a> Master.AsyncPersistFileSize</td>  <td>COUNTER</td>  <td>AsyncPersist 操作创建的文件总大小</td></tr>
<tr>  <td><a class="anchor" name="Master.AsyncPersistSuccess"></a> Master.AsyncPersistSuccess</td>  <td>COUNTER</td>  <td>成功的 AsyncPersist 操作数量</td></tr>
<tr>  <td><a class="anchor" name="Master.AuditLogEntriesSize"></a> Master.AuditLogEntriesSize</td>  <td>GAUGE</td>  <td>审核日志条目队列的大小</td></tr>
<tr>  <td><a class="anchor" name="Master.BlockHeapSize"></a> Master.BlockHeapSize</td>  <td>GAUGE</td>  <td>数据块元数据占 JVM 堆大小的估计值</td></tr>
<tr>  <td><a class="anchor" name="Master.BlockReplicaCount"></a> Master.BlockReplicaCount</td>  <td>GAUGE</td>  <td>Alluxio 中块副本的总数</td></tr>
<tr>  <td><a class="anchor" name="Master.CachedBlockLocations"></a> Master.CachedBlockLocations</td>  <td>GAUGE</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.CompleteFileOps"></a> Master.CompleteFileOps</td>  <td>COUNTER</td>  <td>CompleteFile 操作的总数</td></tr>
<tr>  <td><a class="anchor" name="Master.CompletedOperationRetryCount"></a> Master.CompletedOperationRetryCount</td>  <td>COUNTER</td>  <td>已由客户端重试的完成操作总数</td></tr>
<tr>  <td><a class="anchor" name="Master.CreateDirectoryOps"></a> Master.CreateDirectoryOps</td>  <td>COUNTER</td>  <td>CreateDirectory 操作的总数</td></tr>
<tr>  <td><a class="anchor" name="Master.CreateFileOps"></a> Master.CreateFileOps</td>  <td>COUNTER</td>  <td>CreateFile 操作的总数</td></tr>
<tr>  <td><a class="anchor" name="Master.DeletePathOps"></a> Master.DeletePathOps</td>  <td>COUNTER</td>  <td>Delete 操作的总数</td></tr>
<tr>  <td><a class="anchor" name="Master.DirectoriesCreated"></a> Master.DirectoriesCreated</td>  <td>COUNTER</td>  <td>CreateDirectory 操作的总数</td></tr>
<tr>  <td><a class="anchor" name="Master.EdgeCacheEvictions"></a> Master.EdgeCacheEvictions</td>  <td>GAUGE</td>  <td>从缓存中删除的边（inode 元数据）总数。边缓存负责管理从（parentId，childName）到 childId 的映射</td></tr>
<tr>  <td><a class="anchor" name="Master.EdgeCacheHits"></a> Master.EdgeCacheHits</td>  <td>GAUGE</td>  <td>边（inode 元数据）缓存的命中总数。边缓存负责管理从（parentId，childName）到 childId 的映射</td></tr>
<tr>  <td><a class="anchor" name="Master.EdgeCacheLoadTimes"></a> Master.EdgeCacheLoadTimes</td>  <td>GAUGE</td>  <td>导致缓存未命中的边（inode 元数据）缓存的总加载时间。边缓存负责管理从（parentId，childName）到 childId 的映射</td></tr>
<tr>  <td><a class="anchor" name="Master.EdgeCacheMisses"></a> Master.EdgeCacheMisses</td>  <td>GAUGE</td>  <td>边（inode 元数据）缓存的未命中总数。边缓存负责管理从（parentId，childName）到 childId 的映射</td></tr>
<tr>  <td><a class="anchor" name="Master.EdgeCacheSize"></a> Master.EdgeCacheSize</td>  <td>GAUGE</td>  <td>缓存的边（inode 元数据）总数。边缓存负责管理从（parentId，childName）到 childId 的映射</td></tr>
<tr>  <td><a class="anchor" name="Master.EdgeLockPoolSize"></a> Master.EdgeLockPoolSize</td>  <td>GAUGE</td>  <td>Edge 锁池的大小</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotDownloadDiskSize"></a> Master.EmbeddedJournalLastSnapshotDownloadDiskSize</td>  <td>GAUGE</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotDownloadDurationMs"></a> Master.EmbeddedJournalLastSnapshotDownloadDurationMs</td>  <td>GAUGE</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotDownloadSize"></a> Master.EmbeddedJournalLastSnapshotDownloadSize</td>  <td>GAUGE</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotDurationMs"></a> Master.EmbeddedJournalLastSnapshotDurationMs</td>  <td>GAUGE</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotEntriesCount"></a> Master.EmbeddedJournalLastSnapshotEntriesCount</td>  <td>GAUGE</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotReplayDurationMs"></a> Master.EmbeddedJournalLastSnapshotReplayDurationMs</td>  <td>GAUGE</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotReplayEntriesCount"></a> Master.EmbeddedJournalLastSnapshotReplayEntriesCount</td>  <td>GAUGE</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotUploadDiskSize"></a> Master.EmbeddedJournalLastSnapshotUploadDiskSize</td>  <td>GAUGE</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotUploadDurationMs"></a> Master.EmbeddedJournalLastSnapshotUploadDurationMs</td>  <td>GAUGE</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalLastSnapshotUploadSize"></a> Master.EmbeddedJournalLastSnapshotUploadSize</td>  <td>GAUGE</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotDownloadDiskHistogram"></a> Master.EmbeddedJournalSnapshotDownloadDiskHistogram</td>  <td>HISTOGRAM</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotDownloadGenerate"></a> Master.EmbeddedJournalSnapshotDownloadGenerate</td>  <td>TIMER</td>  <td>描述从集群中的其他主机下载日志快照所需的时间。只有在使用嵌入式日志时有效。使用此指标可以确定 Alluxio 主机之间是否存在潜在的通信瓶颈</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotDownloadHistogram"></a> Master.EmbeddedJournalSnapshotDownloadHistogram</td>  <td>HISTOGRAM</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotGenerateTimer"></a> Master.EmbeddedJournalSnapshotGenerateTimer</td>  <td>TIMER</td>  <td>描述在此主机上生成本地日志快照所需的时间。只有在使用嵌入式日志时有效。使用此指标可以测量 Alluxio 快照生成的性能</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotInstallTimer"></a> Master.EmbeddedJournalSnapshotInstallTimer</td>  <td>TIMER</td>  <td>描述从另一个主机安装下载的日志快照所需的时间。只有在使用嵌入式日志时有效。使用此指标可以确定 Alluxio 在从 leader 安装快照时的性能。较高的数字可能表示磁盘性能低或 CPU 竞争大</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotLastIndex"></a> Master.EmbeddedJournalSnapshotLastIndex</td>  <td>GAUGE</td>  <td>表示此主机在最近的本地快照或从集群中另一个主机下载的快照中记录的最新日志索引。只有在使用嵌入式日志时才有效</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotReplayTimer"></a> Master.EmbeddedJournalSnapshotReplayTimer</td>  <td>TIMER</td>  <td>描述将日志快照重放到主机状态机所需的时间。只有在使用嵌入式日志时才有效。使用此指标确定 Alluxio 重放日志快照文件时的性能。较高的数字可能表示磁盘性能低或 CPU 竞争大</td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotUploadDiskHistogram"></a> Master.EmbeddedJournalSnapshotUploadDiskHistogram</td>  <td>HISTOGRAM</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotUploadHistogram"></a> Master.EmbeddedJournalSnapshotUploadHistogram</td>  <td>HISTOGRAM</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.EmbeddedJournalSnapshotUploadTimer"></a> Master.EmbeddedJournalSnapshotUploadTimer</td>  <td>TIMER</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.FileBlockInfosGot"></a> Master.FileBlockInfosGot</td>  <td>COUNTER</td>  <td>成功的 GetFileBlockInfo 操作总数</td></tr>
<tr>  <td><a class="anchor" name="Master.FileInfosGot"></a> Master.FileInfosGot</td>  <td>COUNTER</td>  <td>成功的 GetFileInfo 操作总数</td></tr>
<tr>  <td><a class="anchor" name="Master.FileSize"></a> Master.FileSize</td>  <td>GAUGE</td>  <td>文件大小分布</td></tr>
<tr>  <td><a class="anchor" name="Master.FilesCompleted"></a> Master.FilesCompleted</td>  <td>COUNTER</td>  <td>成功的 CompleteFile 操作总数</td></tr>
<tr>  <td><a class="anchor" name="Master.FilesCreated"></a> Master.FilesCreated</td>  <td>COUNTER</td>  <td>成功的 CreateFile 操作总数</td></tr>
<tr>  <td><a class="anchor" name="Master.FilesFreed"></a> Master.FilesFreed</td>  <td>COUNTER</td>  <td>成功的 FreeFile 操作总数</td></tr>
<tr>  <td><a class="anchor" name="Master.FilesPersisted"></a> Master.FilesPersisted</td>  <td>COUNTER</td>  <td>成功持久化的文件总数</td></tr>
<tr>  <td><a class="anchor" name="Master.FilesPinned"></a> Master.FilesPinned</td>  <td>GAUGE</td>  <td>当前固定的文件总数</td></tr>
<tr>  <td><a class="anchor" name="Master.FilesToBePersisted"></a> Master.FilesToBePersisted</td>  <td>GAUGE</td>  <td>当前待持久化的文件总数</td></tr>
<tr>  <td><a class="anchor" name="Master.FreeFileOps"></a> Master.FreeFileOps</td>  <td>COUNTER</td>  <td>FreeFile 操作总数</td></tr>
<tr>  <td><a class="anchor" name="Master.GetFileBlockInfoOps"></a> Master.GetFileBlockInfoOps</td>  <td>COUNTER</td>  <td>GetFileBlockInfo 操作总数</td></tr>
<tr>  <td><a class="anchor" name="Master.GetFileInfoOps"></a> Master.GetFileInfoOps</td>  <td>COUNTER</td>  <td>GetFileInfo 操作总数</td></tr>
<tr>  <td><a class="anchor" name="Master.GetNewBlockOps"></a> Master.GetNewBlockOps</td>  <td>COUNTER</td>  <td>GetNewBlock 操作总数</td></tr>
<tr>  <td><a class="anchor" name="Master.InodeCacheEvictions"></a> Master.InodeCacheEvictions</td>  <td>GAUGE</td>  <td>缓存逐出的 inode 总数</td></tr>
<tr>  <td><a class="anchor" name="Master.InodeCacheHitRatio"></a> Master.InodeCacheHitRatio</td>  <td>GAUGE</td>  <td>Inode 缓存命中率</td></tr>
<tr>  <td><a class="anchor" name="Master.InodeCacheHits"></a> Master.InodeCacheHits</td>  <td>GAUGE</td>  <td>inode（inode 元数据）缓存的命中总数</td></tr>
<tr>  <td><a class="anchor" name="Master.InodeCacheLoadTimes"></a> Master.InodeCacheLoadTimes</td>  <td>GAUGE</td>  <td>缓存未命中导致的 inode（inode 元数据）加载次数总数</td></tr>
<tr>  <td><a class="anchor" name="Master.InodeCacheMisses"></a> Master.InodeCacheMisses</td>  <td>GAUGE</td>  <td>inode 缓存未命中总数</td></tr>
<tr>  <td><a class="anchor" name="Master.InodeCacheSize"></a> Master.InodeCacheSize</td>  <td>GAUGE</td>  <td>inode（inode 元数据）缓存的总数</td></tr>
<tr>  <td><a class="anchor" name="Master.InodeHeapSize"></a> Master.InodeHeapSize</td>  <td>GAUGE</td>  <td>inode 堆大小的估计值</td></tr>
<tr>  <td><a class="anchor" name="Master.InodeLockPoolSize"></a> Master.InodeLockPoolSize</td>  <td>GAUGE</td>  <td>master inode lock pool 大小</td></tr>
<tr>  <td><a class="anchor" name="Master.JobCanceled"></a> Master.JobCanceled</td>  <td>COUNTER</td>  <td>取消状态异步任务数</td></tr>
<tr>  <td><a class="anchor" name="Master.JobCompleted"></a> Master.JobCompleted</td>  <td>COUNTER</td>  <td>完成状态异步任务数</td></tr>
<tr>  <td><a class="anchor" name="Master.JobCount"></a> Master.JobCount</td>  <td>GAUGE</td>  <td>所有状态任务数</td></tr>
<tr>  <td><a class="anchor" name="Master.JobCreated"></a> Master.JobCreated</td>  <td>COUNTER</td>  <td>创建状态任务数</td></tr>
<tr>  <td><a class="anchor" name="Master.JobDistributedLoadBlockSizes"></a> Master.JobDistributedLoadBlockSizes</td>  <td>COUNTER</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.JobDistributedLoadCancel"></a> Master.JobDistributedLoadCancel</td>  <td>COUNTER</td>  <td>取消的 DistributedLoad 操作数</td></tr>
<tr>  <td><a class="anchor" name="Master.JobDistributedLoadFail"></a> Master.JobDistributedLoadFail</td>  <td>COUNTER</td>  <td>失败的 DistributedLoad 操作数</td></tr>
<tr>  <td><a class="anchor" name="Master.JobDistributedLoadFileCount"></a> Master.JobDistributedLoadFileCount</td>  <td>COUNTER</td>  <td>DistributedLoad 操作的文件数</td></tr>
<tr>  <td><a class="anchor" name="Master.JobDistributedLoadFileSizes"></a> Master.JobDistributedLoadFileSizes</td>  <td>COUNTER</td>  <td>DistributedLoad 操作的文件大小</td></tr>
<tr>  <td><a class="anchor" name="Master.JobDistributedLoadRate"></a> Master.JobDistributedLoadRate</td>  <td>METER</td>  <td>平均 DistributedLoad 加载率</td></tr>
<tr>  <td><a class="anchor" name="Master.JobDistributedLoadSuccess"></a> Master.JobDistributedLoadSuccess</td>  <td>COUNTER</td>  <td>DistributedLoad 操作成功数</td></tr>
<tr>  <td><a class="anchor" name="Master.JobFailed"></a> Master.JobFailed</td>  <td>COUNTER</td>  <td>失败状态异步任务数</td></tr>
<tr>  <td><a class="anchor" name="Master.JobLoadBlockCount"></a> Master.JobLoadBlockCount</td>  <td>COUNTER</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.JobLoadBlockFail"></a> Master.JobLoadBlockFail</td>  <td>COUNTER</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.JobLoadFail"></a> Master.JobLoadFail</td>  <td>COUNTER</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.JobLoadRate"></a> Master.JobLoadRate</td>  <td>METER</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.JobLoadSuccess"></a> Master.JobLoadSuccess</td>  <td>COUNTER</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.JobRunning"></a> Master.JobRunning</td>  <td>COUNTER</td>  <td>运行中状态异步任务数</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalCheckpointWarn"></a> Master.JournalCheckpointWarn</td>  <td>GAUGE</td>  <td>alluxio.master.journal.checkpoint.period.entries，并且最后一个检查点超过了 alluxio.master.journal.checkpoint.warning.threshold.time，则返回 1 以指示需要警告，否则返回 0</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalEntriesSinceCheckPoint"></a> Master.JournalEntriesSinceCheckPoint</td>  <td>GAUGE</td>  <td>自上次检查点以来的日志条目数</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalFlushFailure"></a> Master.JournalFlushFailure</td>  <td>COUNTER</td>  <td>日志刷新失败的总数</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalFlushTimer"></a> Master.JournalFlushTimer</td>  <td>TIMER</td>  <td>日志刷新计时器统计</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalFreeBytes"></a> Master.JournalFreeBytes</td>  <td>GAUGE</td>  <td>Alluxio 主机的日志磁盘上剩余的字节。此指标仅在 Linux 上使用内置日志时有效。使用此指标监视日志是否耗尽磁盘空间</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalFreePercent"></a> Master.JournalFreePercent</td>  <td>GAUGE</td>  <td>Alluxio 主机日志磁盘剩余字节。此指标仅在 Linux 上且使用内置日志时有效。使用此指标监控日志是否有剩余磁盘空间</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalGainPrimacyTimer"></a> Master.JournalGainPrimacyTimer</td>  <td>TIMER</td>  <td>日志获得优先权的计时器统计信息</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalLastAppliedCommitIndex"></a> Master.JournalLastAppliedCommitIndex</td>  <td>GAUGE</td>  <td>最后一个被应用到状态机的 raft 日志索引</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalLastCheckPointTime"></a> Master.JournalLastCheckPointTime</td>  <td>GAUGE</td>  <td>上一个日志检查点时间</td></tr>
<tr>  <td><a class="anchor" name="Master.JournalSequenceNumber"></a> Master.JournalSequenceNumber</td>  <td>GAUGE</td>  <td>当前日志序列号</td></tr>
<tr>  <td><a class="anchor" name="Master.LastBackupEntriesCount"></a> Master.LastBackupEntriesCount</td>  <td>GAUGE</td>  <td>上次主元数据备份中写入的条目总数</td></tr>
<tr>  <td><a class="anchor" name="Master.LastBackupRestoreCount"></a> Master.LastBackupRestoreCount</td>  <td>GAUGE</td>  <td>当 primary master 初始化元数据时，从备份还原的条目总数</td></tr>
<tr>  <td><a class="anchor" name="Master.LastBackupRestoreTimeMs"></a> Master.LastBackupRestoreTimeMs</td>  <td>GAUGE</td>  <td>最后一次从备份恢复的过程时间</td></tr>
<tr>  <td><a class="anchor" name="Master.LastBackupTimeMs"></a> Master.LastBackupTimeMs</td>  <td>GAUGE</td>  <td>上一次备份的时间</td></tr>
<tr>  <td><a class="anchor" name="Master.LastGainPrimacyTime"></a> Master.LastGainPrimacyTime</td>  <td>GAUGE</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.LastLosePrimacyTime"></a> Master.LastLosePrimacyTime</td>  <td>GAUGE</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.ListingCacheEvictions"></a> Master.ListingCacheEvictions</td>  <td>COUNTER</td>  <td>master 节点列表缓存中的总淘汰次数</td></tr>
<tr>  <td><a class="anchor" name="Master.ListingCacheHits"></a> Master.ListingCacheHits</td>  <td>COUNTER</td>  <td>master 列表缓存中的命中总数</td></tr>
<tr>  <td><a class="anchor" name="Master.ListingCacheLoadTimes"></a> Master.ListingCacheLoadTimes</td>  <td>COUNTER</td>  <td>master 列表缓存的总加载时间（以纳秒为单位），这是由缓存未命中所导致的</td></tr>
<tr>  <td><a class="anchor" name="Master.ListingCacheMisses"></a> Master.ListingCacheMisses</td>  <td>COUNTER</td>  <td>master 列表缓存中的未命中总数</td></tr>
<tr>  <td><a class="anchor" name="Master.ListingCacheSize"></a> Master.ListingCacheSize</td>  <td>GAUGE</td>  <td>master 列表缓存大小</td></tr>
<tr>  <td><a class="anchor" name="Master.LostBlockCount"></a> Master.LostBlockCount</td>  <td>GAUGE</td>  <td>丢失数据块计数</td></tr>
<tr>  <td><a class="anchor" name="Master.LostFileCount"></a> Master.LostFileCount</td>  <td>GAUGE</td>  <td>丢失文件的数量。这个数字是被缓存的，可能与 Master.LostBlockCount 不同步</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncActivePaths"></a> Master.MetadataSyncActivePaths</td>  <td>COUNTER</td>  <td>所有 InodeSyncStream 实例中正在进行的路径数量</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncExecutor"></a> Master.MetadataSyncExecutor</td>  <td>EXECUTOR_SERVICE</td>  <td>master 元数据同步执行器线程的指标。Master.MetadataSyncExecutor.submitted 是提交给执行器的任务的计数。Master.MetadataSyncExecutor.completed 是执行器完成的任务的计数。Master.MetadataSyncExecutor.activeTaskQueue 是在执行器中每次添加新任务时计算的活动任务（运行或提交）的幂指数衰减随机容器的数量。最大值是执行过程中任何时候的活动任务的最大数量。Master.MetadataSyncExecutor.running 是执行器正在运行的任务数量。Master.MetadataSyncExecutor.idle 是提交的任务（即在执行前等待队列）闲置的时间。Master.MetadataSyncExecutor.duration 是运行提交的任务的时间。如果执行器是线程池执行器，则 Master.MetadataSyncExecutor.queueSize 是任务队列的大小</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncExecutorQueueSize"></a> Master.MetadataSyncExecutorQueueSize</td>  <td>GAUGE</td>  <td>元数据同步线程池中排队的同步任务数，由 alluxio.master.metadata.sync.executor.pool.size 控制</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncFail"></a> Master.MetadataSyncFail</td>  <td>COUNTER</td>  <td>InodeSyncStream 失败的次数，无论是部分失败还是完全失败</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncNoChange"></a> Master.MetadataSyncNoChange</td>  <td>COUNTER</td>  <td>未更改 inodes 的 InodeSyncStream 完成数量</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncOpsCount"></a> Master.MetadataSyncOpsCount</td>  <td>COUNTER</td>  <td>元数据同步操作的数量。每个同步操作对应于一个 InodeSyncStream 实例</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPathsCancel"></a> Master.MetadataSyncPathsCancel</td>  <td>COUNTER</td>  <td>所有最终被忽略而没被处理的 InodeSyncStream 实例中未决路径的数量</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPathsFail"></a> Master.MetadataSyncPathsFail</td>  <td>COUNTER</td>  <td>在元数据同步所有 InodeSyncStream 实例期间失败的路径数量。</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPathsSuccess"></a> Master.MetadataSyncPathsSuccess</td>  <td>COUNTER</td>  <td>从所有 InodeSyncStream 实例同步的路径数量</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPendingPaths"></a> Master.MetadataSyncPendingPaths</td>  <td>COUNTER</td>  <td>所有活跃 InodeSyncStream 实例中等待元数据同步的的待处理路径数量</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPrefetchCancel"></a> Master.MetadataSyncPrefetchCancel</td>  <td>COUNTER</td>  <td>从元数据同步取消的预取任务数量（由于重复的预取请求）</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPrefetchExecutor"></a> Master.MetadataSyncPrefetchExecutor</td>  <td>EXECUTOR_SERVICE</td>  <td>关于主元数据同步预取执行线程的指标。Master.MetadataSyncPrefetchExecutor.submitted 是提交给执行器的任务的计数器。Master.MetadataSyncPrefetchExecutor.completed 是由执行器完成的任务的计数器。Master.MetadataSyncPrefetchExecutor.activeTaskQueue 是在执行器上运行或提交的活动任务的指数衰减随机容器，每次向执行器添加新任务时计算。最大值是执行期间任意时间内的最大活动任务数。Master.MetadataSyncPrefetchExecutor.running 是执行器正在运行的任务数。Master.MetadataSyncPrefetchExecutor.idle 是提交的任务（即等待队列中之前执行的时间）的空闲时间。Master.MetadataSyncPrefetchExecutor.duration 是运行提交的任务的时间。如果执行器是线程池执行器，则 Master.MetadataSyncPrefetchExecutor.queueSize 是任务队列的大小。</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPrefetchExecutorQueueSize"></a> Master.MetadataSyncPrefetchExecutorQueueSize</td>  <td>GAUGE</td>  <td>元数据同步线程池中排队的预取任务数，由 alluxio.master.metadata.sync.ufs.prefetch.pool.size 控制</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPrefetchFail"></a> Master.MetadataSyncPrefetchFail</td>  <td>COUNTER</td>  <td>元数据同步中失败的预取工作数量</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPrefetchOpsCount"></a> Master.MetadataSyncPrefetchOpsCount</td>  <td>COUNTER</td>  <td>由预取线程池处理的预取操作数量</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPrefetchPaths"></a> Master.MetadataSyncPrefetchPaths</td>  <td>COUNTER</td>  <td>元数据同步预取任务获取的 UFS 路径总数</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPrefetchRetries"></a> Master.MetadataSyncPrefetchRetries</td>  <td>COUNTER</td>  <td>元数据同步预取任务的重试获取次数</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncPrefetchSuccess"></a> Master.MetadataSyncPrefetchSuccess</td>  <td>COUNTER</td>  <td>元数据同步预取任务的成功获取次数</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncSkipped"></a> Master.MetadataSyncSkipped</td>  <td>COUNTER</td>  <td>由于 Alluxio 元数据比 alluxio.user.file.metadata.sync.interval 更新而跳过的 InodeSyncStream 数量</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncSuccess"></a> Master.MetadataSyncSuccess</td>  <td>COUNTER</td>  <td>InodeSyncStream 成功次数</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncTimeMs"></a> Master.MetadataSyncTimeMs</td>  <td>COUNTER</td>  <td>所有 InodeSyncStream 实例存在总时间</td></tr>
<tr>  <td><a class="anchor" name="Master.MetadataSyncUfsMount."></a> Master.MetadataSyncUfsMount.</td>  <td>COUNTER</td>  <td>对给定 UFS 挂载点进行 UFS sync 操作的次数</td></tr>
<tr>  <td><a class="anchor" name="Master.MigrateJobCancel"></a> Master.MigrateJobCancel</td>  <td>COUNTER</td>  <td>MigrateJob 取消次数</td></tr>
<tr>  <td><a class="anchor" name="Master.MigrateJobFail"></a> Master.MigrateJobFail</td>  <td>COUNTER</td>  <td>MigrateJob 失败次数</td></tr>
<tr>  <td><a class="anchor" name="Master.MigrateJobFileCount"></a> Master.MigrateJobFileCount</td>  <td>COUNTER</td>  <td>MigrateJob 文件数</td></tr>
<tr>  <td><a class="anchor" name="Master.MigrateJobFileSize"></a> Master.MigrateJobFileSize</td>  <td>COUNTER</td>  <td>MigrateJob 文件体积之和</td></tr>
<tr>  <td><a class="anchor" name="Master.MigrateJobSuccess"></a> Master.MigrateJobSuccess</td>  <td>COUNTER</td>  <td>MigrateJob 操作成功次数</td></tr>
<tr>  <td><a class="anchor" name="Master.MountOps"></a> Master.MountOps</td>  <td>COUNTER</td>  <td>Mount 操作次数</td></tr>
<tr>  <td><a class="anchor" name="Master.NewBlocksGot"></a> Master.NewBlocksGot</td>  <td>COUNTER</td>  <td>GetNewBlock 操作成功次数</td></tr>
<tr>  <td><a class="anchor" name="Master.PathsDeleted"></a> Master.PathsDeleted</td>  <td>COUNTER</td>  <td>Delete 操作成功次数</td></tr>
<tr>  <td><a class="anchor" name="Master.PathsMounted"></a> Master.PathsMounted</td>  <td>COUNTER</td>  <td>Mount 操作成功次数</td></tr>
<tr>  <td><a class="anchor" name="Master.PathsRenamed"></a> Master.PathsRenamed</td>  <td>COUNTER</td>  <td>Rename 操作成功次数</td></tr>
<tr>  <td><a class="anchor" name="Master.PathsUnmounted"></a> Master.PathsUnmounted</td>  <td>COUNTER</td>  <td>Unmount 操作成功次数</td></tr>
<tr>  <td><a class="anchor" name="Master.RenamePathOps"></a> Master.RenamePathOps</td>  <td>COUNTER</td>  <td>Rename 操作次数</td></tr>
<tr>  <td><a class="anchor" name="Master.ReplicaMgmtActiveJobSize"></a> Master.ReplicaMgmtActiveJobSize</td>  <td>GAUGE</td>  <td>活跃块复制/逐出任务的数。这些任务由 master 创建，以维护块副本因子。该值带有一定延迟，是估计值</td></tr>
<tr>  <td><a class="anchor" name="Master.ReplicationLimitedFiles"></a> Master.ReplicationLimitedFiles</td>  <td>COUNTER</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockBackgroundErrors"></a> Master.RocksBlockBackgroundErrors</td>  <td>GAUGE</td>  <td>RocksDB 块表。背景错误累积数</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockBlockCacheCapacity"></a> Master.RocksBlockBlockCacheCapacity</td>  <td>GAUGE</td>  <td>RocksDB 块表。块缓存容量</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockBlockCachePinnedUsage"></a> Master.RocksBlockBlockCachePinnedUsage</td>  <td>GAUGE</td>  <td>RocksDB 块表。固定条目内存体积</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockBlockCacheUsage"></a> Master.RocksBlockBlockCacheUsage</td>  <td>GAUGE</td>  <td>RocksDB 块表。存储在块缓存中的条目的内存大小</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockCompactionPending"></a> Master.RocksBlockCompactionPending</td>  <td>GAUGE</td>  <td>RocksDB 块表。如果有至少一个压缩操作正在等待，则此指标为 1；否则，此指标为 0</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockCurSizeActiveMemTable"></a> Master.RocksBlockCurSizeActiveMemTable</td>  <td>GAUGE</td>  <td>RocksDB 块表。活跃 MemTable 的近似字节大小</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockCurSizeAllMemTables"></a> Master.RocksBlockCurSizeAllMemTables</td>  <td>GAUGE</td>  <td>RocksDB 块表。活跃的、未刷新且不可变的，和固定住不可变的 MemTable 的以字节为单位的近似大小。固定不可变内存表是被保留在内存中的刷新内存表，用于在内存中保留写入历史记录</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockEstimateNumKeys"></a> Master.RocksBlockEstimateNumKeys</td>  <td>GAUGE</td>  <td>RocksDB 块表。活跃和未刷新地不可变 MemTable 以及存储中总键数的估计值</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockEstimatePendingCompactionBytes"></a> Master.RocksBlockEstimatePendingCompactionBytes</td>  <td>GAUGE</td>  <td>RocksDB 块表。估计一次压缩需要在磁盘上重写的总字节数，以使所有层降到目标大小之下。换句话说，这个指标与层压缩中的写入放大率有关。因此，这个指标对层压缩以外的压缩是无效的。</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockEstimateTableReadersMem"></a> Master.RocksBlockEstimateTableReadersMem</td>  <td>GAUGE</td>  <td>RocksDB inode 表。以字节为单位估计用于读取 SST 表的内存，不包括块缓存中使用的内存（e.g. 过滤器和索引块）。如果过滤器和索引不在块缓存中维护，此指标记录迭代器使用的内存以及过滤器和索引。此指标基本上反映了读取数据时块缓存外使用的内存</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockEstimatedMemUsage"></a> Master.RocksBlockEstimatedMemUsage</td>  <td>GAUGE</td>  <td>RocksDB块表。这个指标通过聚合 Master.RocksBlockBlockCacheUsage、Master.RocksBlockEstimateTableReadersMem、Master.RocksBlockCurSizeAllMemTables 和 Master.RocksBlockBlockCachePinnedUsage 的值来估计 RockDB 块表的内存使用情况。</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockLiveSstFilesSize"></a> Master.RocksBlockLiveSstFilesSize</td>  <td>GAUGE</td>  <td>RocksDB块表。属于最新 LSM 树的所有 SST 文件以字节为单位的总大小</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockMemTableFlushPending"></a> Master.RocksBlockMemTableFlushPending</td>  <td>GAUGE</td>  <td>RocksDB 块表。如果 Memtable 刷新操作正在等待，则此指标为 1；否则为 0</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockNumDeletesActiveMemTable"></a> Master.RocksBlockNumDeletesActiveMemTable</td>  <td>GAUGE</td>  <td>RocksDB 块表。活跃 Memtable 中的删除条目总数</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockNumDeletesImmMemTables"></a> Master.RocksBlockNumDeletesImmMemTables</td>  <td>GAUGE</td>  <td>RocksDB 块表。未刷新不可变 MemTable 中删除条目的总数</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockNumEntriesActiveMemTable"></a> Master.RocksBlockNumEntriesActiveMemTable</td>  <td>GAUGE</td>  <td>RocksDB 块表。活跃 MemTable 中的条目总数</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockNumEntriesImmMemTables"></a> Master.RocksBlockNumEntriesImmMemTables</td>  <td>GAUGE</td>  <td>RocksDB 块表。未刷新不可变 MemTable 中的条目总数</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockNumImmutableMemTable"></a> Master.RocksBlockNumImmutableMemTable</td>  <td>GAUGE</td>  <td>RocksDB 块表。尚未刷新的不可变 MemTable 的数量</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockNumLiveVersions"></a> Master.RocksBlockNumLiveVersions</td>  <td>GAUGE</td>  <td>RocksDB inode 表。存活版本数。存活版本较多时，通常意味着更多 SST 文件被迭代器或未完成的压缩保留而未被删除</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockNumRunningCompactions"></a> Master.RocksBlockNumRunningCompactions</td>  <td>GAUGE</td>  <td>RocksDB 块表。当前正在运行的压缩数量</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockNumRunningFlushes"></a> Master.RocksBlockNumRunningFlushes</td>  <td>GAUGE</td>  <td>RocksDB 块表。当前正在运行的刷新数量。</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockSizeAllMemTables"></a> Master.RocksBlockSizeAllMemTables</td>  <td>GAUGE</td>  <td>RocksDB 块表。所有 MemTable 的大小</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksBlockTotalSstFilesSize"></a> Master.RocksBlockTotalSstFilesSize</td>  <td>GAUGE</td>  <td>RocksDB 块表。所有 SST 文件以字节为单位的总大小</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeBackgroundErrors"></a> Master.RocksInodeBackgroundErrors</td>  <td>GAUGE</td>  <td>RocksDB inode 表。后台错误累积数</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeBlockCacheCapacity"></a> Master.RocksInodeBlockCacheCapacity</td>  <td>GAUGE</td>  <td>RocksDB inode 表。 块缓存容量</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeBlockCachePinnedUsage"></a> Master.RocksInodeBlockCachePinnedUsage</td>  <td>GAUGE</td>  <td>RocksDB inode 表。固定键内存体积</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeBlockCacheUsage"></a> Master.RocksInodeBlockCacheUsage</td>  <td>GAUGE</td>  <td>RocksDB inode 表。这是用来描述存储在块缓存中的条目内存大小的指标</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeCompactionPending"></a> Master.RocksInodeCompactionPending</td>  <td>GAUGE</td>  <td>RocksDB inode 表。 如果至少有一个压缩操作正在等待则该指标为 1；否则，该指标为 0</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeCurSizeActiveMemTable"></a> Master.RocksInodeCurSizeActiveMemTable</td>  <td>GAUGE</td>  <td>RocksDB inode 表。活跃 MemTable 以字节为单位的近似大小</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeCurSizeAllMemTables"></a> Master.RocksInodeCurSizeAllMemTables</td>  <td>GAUGE</td>  <td>RocksDB inode 表。活跃和未刷新不可变 MemTable 以字节为单位的近似大小</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeEstimateNumKeys"></a> Master.RocksInodeEstimateNumKeys</td>  <td>GAUGE</td>  <td>RocksDB inode 表。活跃和未刷新不可变 MemTable 以及存储中所有键的估计数量</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeEstimatePendingCompactionBytes"></a> Master.RocksInodeEstimatePendingCompactionBytes</td>  <td>GAUGE</td>  <td>RocksDB 块表。估计一次压缩为了将所有层降到目标大小以下需要在磁盘上重写的总字节数。换句话说，这个指标与层压缩中的写入放大率有关。因此，这个指标对层压缩以外的压缩是无效的</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeEstimateTableReadersMem"></a> Master.RocksInodeEstimateTableReadersMem</td>  <td>GAUGE</td>  <td>RocksDB inode 表。估计用于读取SST表的字节数，不包括用于块缓存的内存（e.g. 过滤器和索引块）。如果过滤器和索引不在块缓存中维护，则这个指标记录了迭代器以及过滤器和索引所使用的内存。这个指标基本上反应了在块缓存之外用于读取数据的内存。</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeEstimatedMemUsage"></a> Master.RocksInodeEstimatedMemUsage</td>  <td>GAUGE</td>  <td>RocksDB 块表。这个指标通过聚合 Master.RocksInodeBlockCacheUsage、Master.RocksInodeEstimateTableReadersMem、Master.RocksInodeCurSizeAllMemTables 和 Master.RocksInodeBlockCachePinnedUsage 的值，估计了 RockDB Inode 表的内存使用情况</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeLiveSstFilesSize"></a> Master.RocksInodeLiveSstFilesSize</td>  <td>GAUGE</td>  <td>RocksDB inode 表。属于最新 LSM 树的所有 SST 文件以字节为单位的总大小</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeMemTableFlushPending"></a> Master.RocksInodeMemTableFlushPending</td>  <td>GAUGE</td>  <td>RocksDB inode 表。如果 MemTable 刷新正在等待，则该指标为 1；否则，该指标为 0</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeNumDeletesActiveMemTable"></a> Master.RocksInodeNumDeletesActiveMemTable</td>  <td>GAUGE</td>  <td>RocksDB inode 表。活跃 MemTable 中删除条目的总数</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeNumDeletesImmMemTables"></a> Master.RocksInodeNumDeletesImmMemTables</td>  <td>GAUGE</td>  <td>RocksDB inode table. 未刷新不可变 MemTable 中删除条目的总数</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeNumEntriesActiveMemTable"></a> Master.RocksInodeNumEntriesActiveMemTable</td>  <td>GAUGE</td>  <td>RocksDB inode 表。活跃 MemTable 中的总条目数</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeNumEntriesImmMemTables"></a> Master.RocksInodeNumEntriesImmMemTables</td>  <td>GAUGE</td>  <td>RocksDB inode 表。未刷新不可变 MemTable 中的总条目数</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeNumImmutableMemTable"></a> Master.RocksInodeNumImmutableMemTable</td>  <td>GAUGE</td>  <td>RocksDB inode 表。尚未刷新的不可变 MemTable 的数量</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeNumLiveVersions"></a> Master.RocksInodeNumLiveVersions</td>  <td>GAUGE</td>  <td>RocksDB inode 表。活跃版本的数量。更多的活跃版本通常意味着被迭代器或未完成的压缩保留的不被删除的 SST 文件更多</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeNumRunningCompactions"></a> Master.RocksInodeNumRunningCompactions</td>  <td>GAUGE</td>  <td>RocksDB inode 表。当前正在进行的压缩数量</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeNumRunningFlushes"></a> Master.RocksInodeNumRunningFlushes</td>  <td>GAUGE</td>  <td>RocksDB inode 表。当前正在进行的刷新数量</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeSizeAllMemTables"></a> Master.RocksInodeSizeAllMemTables</td>  <td>GAUGE</td>  <td>RocksDB inode 表。当前活跃的，为刷新不变的，以及固定不变的 MemTable 以字节为单位的近似体积。固定不变的 MemTable 是保留在内存中用于维护内存写入历史更新过的 MemTable</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksInodeTotalSstFilesSize"></a> Master.RocksInodeTotalSstFilesSize</td>  <td>GAUGE</td>  <td>RocksDB inode 表。所有 SST 文件以字节为单位的总体积。</td></tr>
<tr>  <td><a class="anchor" name="Master.RocksTotalEstimatedMemUsage"></a> Master.RocksTotalEstimatedMemUsage</td>  <td>GAUGE</td>  <td>这个指标通过汇总 Master.RocksBlockEstimatedMemUs age 和 Master.RocksInodeEstimatedMemUsage 的值，给出了 RocksDB 内存使用总量</td></tr>
<tr>  <td><a class="anchor" name="Master.RoleId"></a> Master.RoleId</td>  <td>GAUGE</td>  <td>展示 master role id</td></tr>
<tr>  <td><a class="anchor" name="Master.RpcQueueLength"></a> Master.RpcQueueLength</td>  <td>GAUGE</td>  <td>master RPC 队列的长度。使用这个指标来监控 master 上的 RPC 压力</td></tr>
<tr>  <td><a class="anchor" name="Master.RpcThreadActiveCount"></a> Master.RpcThreadActiveCount</td>  <td>GAUGE</td>  <td>在 master RPC 执行器线程池中正在积极执行任务的线程数量。使用这个指标来监控 master 上的 RPC 压力</td></tr>
<tr>  <td><a class="anchor" name="Master.RpcThreadCurrentCount"></a> Master.RpcThreadCurrentCount</td>  <td>GAUGE</td>  <td>当前 master RPC 执行器线程池中的线程数。使用这个指标来监控主服务器上的 RPC 压力</td></tr>
<tr>  <td><a class="anchor" name="Master.SetAclOps"></a> Master.SetAclOps</td>  <td>COUNTER</td>  <td>SetAcl 操作总次数</td></tr>
<tr>  <td><a class="anchor" name="Master.SetAttributeOps"></a> Master.SetAttributeOps</td>  <td>COUNTER</td>  <td>SetAttribute 操作总次数</td></tr>
<tr>  <td><a class="anchor" name="Master.StartTime"></a> Master.StartTime</td>  <td>GAUGE</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.TTLBuckets"></a> Master.TTLBuckets</td>  <td>GAUGE</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.TTLInodes"></a> Master.TTLInodes</td>  <td>GAUGE</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Master.ToRemoveBlockCount"></a> Master.ToRemoveBlockCount</td>  <td>GAUGE</td>  <td>要从 worker 中移除的块副本数量。如果 1 个块要从 2 个 worker 中移除，会被记为 2 个</td></tr>
<tr>  <td><a class="anchor" name="Master.TotalPaths"></a> Master.TotalPaths</td>  <td>GAUGE</td>  <td>Alluxio 命名空间中的文件和目录总数</td></tr>
<tr>  <td><a class="anchor" name="Master.TotalRpcs"></a> Master.TotalRpcs</td>  <td>TIMER</td>  <td>master RPC 调用的吞吐量。这个指标表明 master 服务客户端请求的繁忙程度</td></tr>
<tr>  <td><a class="anchor" name="Master.UfsJournalCatchupTimer"></a> Master.UfsJournalCatchupTimer</td>  <td>TIMER</td>  <td>日志追赶的定时器统计只在使用 Ufs 日志时有效。它提供了一个 standby master 赶上 master 所需时间的概要，如果 master 转换时间过长则应进行监控</td></tr>
<tr>  <td><a class="anchor" name="Master.UfsJournalFailureRecoverTimer"></a> Master.UfsJournalFailureRecoverTimer</td>  <td>TIMER</td>  <td>UFS 日志故障恢复的定时器统计数据</td></tr>
<tr>  <td><a class="anchor" name="Master.UfsJournalInitialReplayTimeMs"></a> Master.UfsJournalInitialReplayTimeMs</td>  <td>GAUGE</td>  <td>启动时 UFS 日志初始回放过程的持续时间。只在使用 UFS 日志时有效。它记录了第一次日志回放的持续时间。使用这个指标来监测你的 master 启动时间是否过长</td></tr>
<tr>  <td><a class="anchor" name="Master.UfsStatusCacheChildrenSize"></a> Master.UfsStatusCacheChildrenSize</td>  <td>COUNTER</td>  <td>UFS 文件元数据缓存总量。该缓存在元数据同步期间使用</td></tr>
<tr>  <td><a class="anchor" name="Master.UfsStatusCacheSize"></a> Master.UfsStatusCacheSize</td>  <td>COUNTER</td>  <td>正在由元数据同步预取线程池处理的 Alluxio 路径总数</td></tr>
<tr>  <td><a class="anchor" name="Master.UniqueBlocks"></a> Master.UniqueBlocks</td>  <td>GAUGE</td>  <td>Alluxio 中数据块总数（不算副本）</td></tr>
<tr>  <td><a class="anchor" name="Master.UnmountOps"></a> Master.UnmountOps</td>  <td>COUNTER</td>  <td>Unmount 操作总次数</td></tr>

</tbody></table>

动态生成的 Master 指标:

| 名称 | 描述 |
|-------------------------|-----------------------------------------------------|
| Master.CapacityTotalTier{TIER_NAME} | Alluxio 文件系统中层 {TIER_NAME} 以字节为单位的总容量 |
| Master.CapacityUsedTier{TIER_NAME}  | Alluxio 文件系统中层 {TIER_NAME} 以字节为单位已使用的容量 |
| Master.CapacityFreeTier{TIER_NAME}  | Alluxio 文件系统中层 {TIER_NAME} 以字节为单位未使用的容量 |
| Master.UfsSessionCount-Ufs:{UFS_ADDRESS} | 当前打开并连接到给定 {UFS_ADDRESS} 的 UFS 会话数 |
| Master.{UFS_RPC_NAME}.UFS:{UFS_ADDRESS}.UFS_TYPE:{UFS_TYPE}.User:{USER} | 当前 master 完成的 UFS RPC 操作细节 |
| Master.PerUfsOp{UFS_RPC_NAME}.UFS:{UFS_ADDRESS} | 当前主 master 在 UFS {UFS_ADDRESS} 上运行的 UFS 操作 {UFS_RPC_NAME} 的总数 | 
| Master.{LEADING_MASTER_RPC_NAME} | 主 master 上暴露的 RPC 调用的持续时间统计信息 |

## Worker 指标

默认 worker 指标:

<table class="table table-striped">
<tbody><tr><th>名称</th><th>类型</th><th>描述</th></tr>
<tr>  <td><a class="anchor" name="Worker.ActiveClients"></a> Worker.ActiveClients</td>  <td>COUNTER</td>  <td>正在活跃地读取或写入此 worker 的客户端数量</td></tr>
<tr>  <td><a class="anchor" name="Worker.ActiveRpcReadCount"></a> Worker.ActiveRpcReadCount</td>  <td>COUNTER</td>  <td>此 worker 管理的读 RPC 数量</td></tr>
<tr>  <td><a class="anchor" name="Worker.ActiveRpcWriteCount"></a> Worker.ActiveRpcWriteCount</td>  <td>COUNTER</td>  <td>此 worker 管理的写 RPC 数量</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockReaderCompleteTaskCount"></a> Worker.BlockReaderCompleteTaskCount</td>  <td>GAUGE</td>  <td>已经完成执行的读任务的近似值</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockReaderThreadActiveCount"></a> Worker.BlockReaderThreadActiveCount</td>  <td>GAUGE</td>  <td>reader 线程池中正在活跃执行任务的读线程数量的近似值</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockReaderThreadCurrentCount"></a> Worker.BlockReaderThreadCurrentCount</td>  <td>GAUGE</td>  <td>此 reader 线程池中的读线程数</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockReaderThreadMaxCount"></a> Worker.BlockReaderThreadMaxCount</td>  <td>GAUGE</td>  <td>reader 线程池中读线程允许的最大数量</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockRemoverBlocksRemovedCount"></a> Worker.BlockRemoverBlocksRemovedCount</td>  <td>COUNTER</td>  <td>此 worker 中被 asynchronous block remover 成功移除的块的总数量</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockRemoverRemovingBlocksSize"></a> Worker.BlockRemoverRemovingBlocksSize</td>  <td>GAUGE</td>  <td>asynchronous block remover 正在从此 worker 移除的块大小</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockRemoverTryRemoveBlocksSize"></a> Worker.BlockRemoverTryRemoveBlocksSize</td>  <td>GAUGE</td>  <td>asynchronous block remover 正要从此 worker 移除的块大小</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockRemoverTryRemoveCount"></a> Worker.BlockRemoverTryRemoveCount</td>  <td>COUNTER</td>  <td>asynchronous block remover 尝试从此 worker 移除的块大小</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockSerializedCompleteTaskCount"></a> Worker.BlockSerializedCompleteTaskCount</td>  <td>GAUGE</td>  <td>完成执行的块序列化任务完成总量近似值</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockSerializedThreadActiveCount"></a> Worker.BlockSerializedThreadActiveCount</td>  <td>GAUGE</td>  <td>serialized 线程池中正在活跃执行任务的 block serialized 线程近似数量</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockSerializedThreadCurrentCount"></a> Worker.BlockSerializedThreadCurrentCount</td>  <td>GAUGE</td>  <td>此 serialized 线程池中 block serialized 线程数量</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockSerializedThreadMaxCount"></a> Worker.BlockSerializedThreadMaxCount</td>  <td>GAUGE</td>  <td>serialized 线程池中 block serialized 线程允许的最大数量</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockWriterCompleteTaskCount"></a> Worker.BlockWriterCompleteTaskCount</td>  <td>GAUGE</td>  <td>已经完成执行的 block serialized 任务的近似值</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockWriterThreadActiveCount"></a> Worker.BlockWriterThreadActiveCount</td>  <td>GAUGE</td>  <td>writer 线程池中正在活跃执行任务的写线程数量的近似值</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockWriterThreadCurrentCount"></a> Worker.BlockWriterThreadCurrentCount</td>  <td>GAUGE</td>  <td>此 writer 线程池中的写线程数</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlockWriterThreadMaxCount"></a> Worker.BlockWriterThreadMaxCount</td>  <td>GAUGE</td>  <td>writer 线程池中写线程允许的最大数量</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksAccessed"></a> Worker.BlocksAccessed</td>  <td>COUNTER</td>  <td>此 worker 中数据块被访问的总次数</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksCached"></a> Worker.BlocksCached</td>  <td>GAUGE</td>  <td>一个 Alluxio worker 中被用于缓存数据的块总数</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksCancelled"></a> Worker.BlocksCancelled</td>  <td>COUNTER</td>  <td>此 worker 中废弃的临时块总量</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksDeleted"></a> Worker.BlocksDeleted</td>  <td>COUNTER</td>  <td>此 worker 中被外部请求删除的块总量</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksEvicted"></a> Worker.BlocksEvicted</td>  <td>COUNTER</td>  <td>此 worker 中被驱逐的块总量</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksEvictionRate"></a> Worker.BlocksEvictionRate</td>  <td>METER</td>  <td>此 worker 的块驱逐率</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksLost"></a> Worker.BlocksLost</td>  <td>COUNTER</td>  <td>此 worker 丢失块总量</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksPromoted"></a> Worker.BlocksPromoted</td>  <td>COUNTER</td>  <td>此 worker 中，任何一个块被移到新层的总次数</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksReadLocal"></a> Worker.BlocksReadLocal</td>  <td>COUNTER</td>  <td>通过此 worker 本地读的数据块总数</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksReadRemote"></a> Worker.BlocksReadRemote</td>  <td>COUNTER</td>  <td>通过此 worker 远程读的数据块总数</td></tr>
<tr>  <td><a class="anchor" name="Worker.BlocksReadUfs"></a> Worker.BlocksReadUfs</td>  <td>COUNTER</td>  <td>通过此 worker 从 UFS 读取的数据块总数</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesReadDirect"></a> Worker.BytesReadDirect</td>  <td>COUNTER</td>  <td>此 worker 中没有外部 RPC 参与的总字节数。数据存在于 worker 存储中或者由此 worker 从底层 UFS 获取。此指标记录了 worker 内部调用读取的数据（e.g. 嵌入在此 worker 中的客户端）</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesReadDirectThroughput"></a> Worker.BytesReadDirectThroughput</td>  <td>METER</td>  <td>此 worker 中没有涉及外部 RPC 的字节读取吞吐量。数据存在于 worker 存储中或由该 worker 从底层 UFS 中获取。这记录了 worker 内部调用（e.g. 位于此 worker 中的客户端）读取的数据</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesReadDomain"></a> Worker.BytesReadDomain</td>  <td>COUNTER</td>  <td>此 worker 通过域套接字读取的总字节数</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesReadDomainThroughput"></a> Worker.BytesReadDomainThroughput</td>  <td>METER</td>  <td>此 worker 通过域套接字读取字节的吞吐量</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesReadPerUfs"></a> Worker.BytesReadPerUfs</td>  <td>COUNTER</td>  <td>此 worker 从特定 UFS 读取的总字节数</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesReadRemote"></a> Worker.BytesReadRemote</td>  <td>COUNTER</td>  <td>通过网络（RPC）远程读取此 worker 的字节总数。数据存在于 worker 存储中或由该 worker 从底层 UFS 获取。这不包括短路本地读取和域套接字读取</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesReadRemoteThroughput"></a> Worker.BytesReadRemoteThroughput</td>  <td>METER</td>  <td>这是一项衡量通过网络（RPC）从此 worker 读取的字节数的吞吐量的指标。数据存在于 worker 存储中，或者由该 worker 从底层 UFS 中获取。这不包括短路本地读取和域套接字读取</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesReadUfsThroughput"></a> Worker.BytesReadUfsThroughput</td>  <td>METER</td>  <td>由此 worker 从 UFS 读取字节的吞吐量</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesWrittenDirect"></a> Worker.BytesWrittenDirect</td>  <td>COUNTER</td>  <td>不涉及外部 RPC写入此 worker 的总字节数。数据写入 worker 存储或由此 worker 写入下层 UFS。这记录了 worker 内部调用（e.g. 嵌入在 此 worker 中的客户端）写入的数据</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesWrittenDirectThroughput"></a> Worker.BytesWrittenDirectThroughput</td>  <td>METER</td>  <td>不涉及外部 RPC 写入此 worker 的字节吞吐量。数据写入 worker 存储或由此 worker 写入下层 UFS。这记录了 worker 内部调用（e.g. 嵌入在此 worker 中的客户端）写入的数据</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesWrittenDomain"></a> Worker.BytesWrittenDomain</td>  <td>COUNTER</td>  <td>通过域套接字写入此 worker 的总字节数</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesWrittenDomainThroughput"></a> Worker.BytesWrittenDomainThroughput</td>  <td>METER</td>  <td>通过域套接字写入此 worker 的吞吐量</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesWrittenPerUfs"></a> Worker.BytesWrittenPerUfs</td>  <td>COUNTER</td>  <td>此 worker 向特定 UFS 写入的总字节数</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesWrittenRemote"></a> Worker.BytesWrittenRemote</td>  <td>COUNTER</td>  <td>通过网络（RPC）写入此 worker 的总字节数。数据写入 worker 存储或由此 worker 写入下层 UFS。这不包括短路本地写入和域套接字写入</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesWrittenRemoteThroughput"></a> Worker.BytesWrittenRemoteThroughput</td>  <td>METER</td>  <td>通过网络（RPC）写入此 worker 的字节写入吞吐量。数据写入 worker 存储或由此 worker 写入下层 UFS。这不包括短路本地写入和域套接字写入</td></tr>
<tr>  <td><a class="anchor" name="Worker.BytesWrittenUfsThroughput"></a> Worker.BytesWrittenUfsThroughput</td>  <td>METER</td>  <td>此 worker 向所有 Alluxio UFS 写入字节的吞吐量</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheBlocksSize"></a> Worker.CacheBlocksSize</td>  <td>COUNTER</td>  <td>通过缓存请求缓存的字节量</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheFailedBlocks"></a> Worker.CacheFailedBlocks</td>  <td>COUNTER</td>  <td>此 worker 缓存块失败数量</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheManagerCompleteTaskCount"></a> Worker.CacheManagerCompleteTaskCount</td>  <td>GAUGE</td>  <td>已经完成执行的块缓存任务的近似量</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheManagerThreadActiveCount"></a> Worker.CacheManagerThreadActiveCount</td>  <td>GAUGE</td>  <td>cache manager 线程池中正在活跃执行任务的块缓存线程数量的近似值</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheManagerThreadCurrentCount"></a> Worker.CacheManagerThreadCurrentCount</td>  <td>GAUGE</td>  <td>此 cache manager 线程池中的块缓存线程数</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheManagerThreadMaxCount"></a> Worker.CacheManagerThreadMaxCount</td>  <td>GAUGE</td>  <td>cache manager 线程池中块缓存线程允许的最大数量</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheManagerThreadQueueWaitingTaskCount"></a> Worker.CacheManagerThreadQueueWaitingTaskCount</td>  <td>GAUGE</td>  <td>此 worker 中 cache manager 线程池中工作队列中等待的任务数，受 alluxio.worker.network.async.cache.manager.queue.max 的限制。</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheRemoteBlocks"></a> Worker.CacheRemoteBlocks</td>  <td>COUNTER</td>  <td>此 worker 需要从远程源缓存的块的总数</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheRequests"></a> Worker.CacheRequests</td>  <td>COUNTER</td>  <td>此 worker 收到的缓存请求总数</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheRequestsAsync"></a> Worker.CacheRequestsAsync</td>  <td>COUNTER</td>  <td>此 worker 收到的异步缓存请求的总数</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheRequestsSync"></a> Worker.CacheRequestsSync</td>  <td>COUNTER</td>  <td>此 worker 收到的同步缓存请求的总数</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheSucceededBlocks"></a> Worker.CacheSucceededBlocks</td>  <td>COUNTER</td>  <td>此 worker 中的缓存成功的块的总数</td></tr>
<tr>  <td><a class="anchor" name="Worker.CacheUfsBlocks"></a> Worker.CacheUfsBlocks</td>  <td>COUNTER</td>  <td>此 worker 中需要从本地源缓存的块的总数</td></tr>
<tr>  <td><a class="anchor" name="Worker.CapacityFree"></a> Worker.CapacityFree</td>  <td>GAUGE</td>  <td>此 Alluxio worker 的所有层级上的总空闲字节</td></tr>
<tr>  <td><a class="anchor" name="Worker.CapacityTotal"></a> Worker.CapacityTotal</td>  <td>GAUGE</td>  <td>此 Alluxio worker 在所有层次上以字节为单位的总容量</td></tr>
<tr>  <td><a class="anchor" name="Worker.CapacityUsed"></a> Worker.CapacityUsed</td>  <td>GAUGE</td>  <td>此 Alluxio worker 所有层级上使用的总字节数</td></tr>
<tr>  <td><a class="anchor" name="Worker.MasterRegistrationSuccessCount"></a> Worker.MasterRegistrationSuccessCount</td>  <td>COUNTER</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Worker.RpcQueueLength"></a> Worker.RpcQueueLength</td>  <td>GAUGE</td>  <td>worker RPC 队列的长度。用此指标监视 worker 的 RPC 压力</td></tr>
<tr>  <td><a class="anchor" name="Worker.RpcThreadActiveCount"></a> Worker.RpcThreadActiveCount</td>  <td>GAUGE</td>  <td>此 worker RPC 执行程序线程池中正在执行任务的线程数。用此指标监视 worker 的 RPC 压力</td></tr>
<tr>  <td><a class="anchor" name="Worker.RpcThreadCurrentCount"></a> Worker.RpcThreadCurrentCount</td>  <td>GAUGE</td>  <td>此 worker RPC 执行器线程池中的线程数。用此指标监视 worker 的 RPC 压力</td></tr>

</tbody></table>

动态的 worker 指标:

| 名称 | 描述 |
|-------------------------|-----------------------------------------------------|
| Worker.UfsSessionCount-Ufs:{UFS_ADDRESS} | 当前打开并连接到给定 {UFS_ADDRESS} 的 UFS 会话数 |
| Worker.{RPC_NAME}                        | worker 上暴露的 RPC 调用的持续时间统计信息 |

## Client 指标

每个客户端度量将使用其本地主机名或配置的 `alluxio.user.app.id` 进行记录。
如果配置了 `alluxio.user.app.id`，多个客户端可以组合成一个逻辑应用。

<table class="table table-striped">
<tbody><tr><th>名称</th><th>类型</th><th>描述</th></tr>
<tr>  <td><a class="anchor" name="Client.BlockMasterClientCount"></a> Client.BlockMasterClientCount</td>  <td>COUNTER</td>  <td>BlockMasterClientPool 中实例数量</td></tr>
<tr>  <td><a class="anchor" name="Client.BlockReadChunkRemote"></a> Client.BlockReadChunkRemote</td>  <td>TIMER</td>  <td>该客户端从远程 Alluxio worker 读取数据chunk数量。当 alluxio.user.block.read.metrics.enabled 设置为 true 时，才会记录此指标</td></tr>
<tr>  <td><a class="anchor" name="Client.BlockWorkerClientCount"></a> Client.BlockWorkerClientCount</td>  <td>COUNTER</td>  <td>BlockWorkerClientPool 中实例数量</td></tr>
<tr>  <td><a class="anchor" name="Client.BusyExceptionCount"></a> Client.BusyExceptionCount</td>  <td>COUNTER</td>  <td>观察到的 BusyException 数量</td></tr>
<tr>  <td><a class="anchor" name="Client.BytesReadLocal"></a> Client.BytesReadLocal</td>  <td>COUNTER</td>  <td>该客户端短路读取的总字节数</td></tr>
<tr>  <td><a class="anchor" name="Client.BytesReadLocalThroughput"></a> Client.BytesReadLocalThroughput</td>  <td>METER</td>  <td>该客户端短路读取的字节吞吐量</td></tr>
<tr>  <td><a class="anchor" name="Client.BytesWrittenLocal"></a> Client.BytesWrittenLocal</td>  <td>COUNTER</td>  <td>该客户端短路写入 Alluxio 缓存的字节总数</td></tr>
<tr>  <td><a class="anchor" name="Client.BytesWrittenLocalThroughput"></a> Client.BytesWrittenLocalThroughput</td>  <td>METER</td>  <td>该客户端短路写入 Alluxio 缓存的字节吞吐量</td></tr>
<tr>  <td><a class="anchor" name="Client.BytesWrittenUfs"></a> Client.BytesWrittenUfs</td>  <td>COUNTER</td>  <td>该客户端写入 UFS 的字节数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheBytesDiscarded"></a> Client.CacheBytesDiscarded</td>  <td>METER</td>  <td>该客户端缓存丢弃的总字节数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheBytesEvicted"></a> Client.CacheBytesEvicted</td>  <td>METER</td>  <td>该客户端缓存驱逐的总字节数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheBytesReadCache"></a> Client.CacheBytesReadCache</td>  <td>METER</td>  <td>从该客户端缓存读的总字节数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheBytesReadExternal"></a> Client.CacheBytesReadExternal</td>  <td>METER</td>  <td>由于该客户端缓存未命中从 Alluxio 集群读取的总字节数。chunk read 可能导致这个数字小于 Client.CacheBytesReadExternal</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheBytesReadInStreamBuffer"></a> Client.CacheBytesReadInStreamBuffer</td>  <td>METER</td>  <td>从该客户端缓存的输入流缓冲区中读取的总字节数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheBytesRequestedExternal"></a> Client.CacheBytesRequestedExternal</td>  <td>METER</td>  <td>引起缓存未命中的用户读请求总字节数。这个数字可能会比 Client.CacheBytesReadExternal 小，因为它可能被分成多个块读取</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheBytesWrittenCache"></a> Client.CacheBytesWrittenCache</td>  <td>METER</td>  <td>向该客户端缓存写入的总字节数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheCleanErrors"></a> Client.CacheCleanErrors</td>  <td>COUNTER</td>  <td>该客户端为了初始化新缓存时清理已存在缓存路径的失败总数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheCleanupGetErrors"></a> Client.CacheCleanupGetErrors</td>  <td>COUNTER</td>  <td>该客户端清理失败内存读取失败总数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheCleanupPutErrors"></a> Client.CacheCleanupPutErrors</td>  <td>COUNTER</td>  <td>该客户端清理失败内存写入失败总数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheCreateErrors"></a> Client.CacheCreateErrors</td>  <td>COUNTER</td>  <td>在该客户端缓存中创建缓存的失败总数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheDeleteErrors"></a> Client.CacheDeleteErrors</td>  <td>COUNTER</td>  <td>在该客户端缓存中删除缓存数据的失败总数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheDeleteFromStoreErrors"></a> Client.CacheDeleteFromStoreErrors</td>  <td>COUNTER</td>  <td>该客户端删除页的失败总数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheDeleteNonExistingPageErrors"></a> Client.CacheDeleteNonExistingPageErrors</td>  <td>COUNTER</td>  <td>该客户端由于页缺失导致删除页失败的总数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheDeleteNotReadyErrors"></a> Client.CacheDeleteNotReadyErrors</td>  <td>COUNTER</td>  <td>该客户端由于缓存未就绪删除页失败的总数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheGetErrors"></a> Client.CacheGetErrors</td>  <td>COUNTER</td>  <td>从该客户端缓存中获取缓存数据失败总数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheGetNotReadyErrors"></a> Client.CacheGetNotReadyErrors</td>  <td>COUNTER</td>  <td>该客户端由于缓存未就绪获取页失败的总数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheGetStoreReadErrors"></a> Client.CacheGetStoreReadErrors</td>  <td>COUNTER</td>  <td>该客户端由于从页存储读取失败导致客户端缓存中获取缓存数据失败的次数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheHitRate"></a> Client.CacheHitRate</td>  <td>GAUGE</td>  <td>缓存命中率：（# 从缓存读取的字节数）/（# 请求的字节数）</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePageReadCacheTimeNanos"></a> Client.CachePageReadCacheTimeNanos</td>  <td>METER</td>  <td>该客户端缓存命中时读取页面时间（ns）</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePageReadExternalTimeNanos"></a> Client.CachePageReadExternalTimeNanos</td>  <td>METER</td>  <td>该客户端当缓存未命中时，从外部源读取数据所花费时间（ns）</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePages"></a> Client.CachePages</td>  <td>COUNTER</td>  <td>该客户端缓存中的总页数</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePagesDiscarded"></a> Client.CachePagesDiscarded</td>  <td>METER</td>  <td>该客户端恢复页存储时丢失页的总数</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePagesEvicted"></a> Client.CachePagesEvicted</td>  <td>METER</td>  <td>从该客户端缓存中驱逐页的总数</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutAsyncRejectionErrors"></a> Client.CachePutAsyncRejectionErrors</td>  <td>COUNTER</td>  <td>该客户端缓存中放置缓存数据时，由于异步写队列注入失败而导致的失败次数</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutBenignRacingErrors"></a> Client.CachePutBenignRacingErrors</td>  <td>COUNTER</td>  <td>该客户端由于驱逐竞争而导致的缓存页添加失败的次数。这个错误是良性的</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutErrors"></a> Client.CachePutErrors</td>  <td>COUNTER</td>  <td>向该客户端缓存中放置缓存数据的失败次数</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutEvictionErrors"></a> Client.CachePutEvictionErrors</td>  <td>COUNTER</td>  <td>该客户端由于驱逐失败而导致的缓存页添加失败的次数。这个错误是良性的</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutInsufficientSpaceErrors"></a> Client.CachePutInsufficientSpaceErrors</td>  <td>COUNTER</td>  <td>该客户端由于在驱逐后空间不足导致的将缓存数据放入客户端缓存时的失败次数</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutNotReadyErrors"></a> Client.CachePutNotReadyErrors</td>  <td>COUNTER</td>  <td>该客户端由于缓存不能准备好添加页，添加页失败的次数</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutStoreDeleteErrors"></a> Client.CachePutStoreDeleteErrors</td>  <td>COUNTER</td>  <td>该客户端在页存储中删除失败导致的缓存数据放置失败的次数</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutStoreWriteErrors"></a> Client.CachePutStoreWriteErrors</td>  <td>COUNTER</td>  <td>该客户端由于向页面存储写入失败而导致的将缓存数据放入客户端缓存中失败的次数</td></tr>
<tr>  <td><a class="anchor" name="Client.CachePutStoreWriteNoSpaceErrors"></a> Client.CachePutStoreWriteNoSpaceErrors</td>  <td>COUNTER</td>  <td>该客户端未达到缓存容量上限但磁盘已满时将缓存数据放入客户端缓存时失败的次数。如果低估写入数据的存储开销比例，这种情况就可能会发生</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheShadowCacheBytes"></a> Client.CacheShadowCacheBytes</td>  <td>COUNTER</td>  <td>该客户端 shadow cache 的字节数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheShadowCacheBytesHit"></a> Client.CacheShadowCacheBytesHit</td>  <td>COUNTER</td>  <td>该客户端 shadow cache 命中的字节数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheShadowCacheBytesRead"></a> Client.CacheShadowCacheBytesRead</td>  <td>COUNTER</td>  <td>这个从客户端 shadow cache 读取的字节数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheShadowCacheFalsePositiveRatio"></a> Client.CacheShadowCacheFalsePositiveRatio</td>  <td>COUNTER</td>  <td>该客户端正在使用的工作集布隆过滤器犯错的概率。该值为 0-100。如果太高，则需要分配更多空间</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheShadowCachePages"></a> Client.CacheShadowCachePages</td>  <td>COUNTER</td>  <td>该客户端 shadow cache 中页的数量</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheShadowCachePagesHit"></a> Client.CacheShadowCachePagesHit</td>  <td>COUNTER</td>  <td>该客户端 shadow cache 中页的命中次数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheShadowCachePagesRead"></a> Client.CacheShadowCachePagesRead</td>  <td>COUNTER</td>  <td>从该客户端 shadow cache 中读取页的数量</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheSpaceAvailable"></a> Client.CacheSpaceAvailable</td>  <td>GAUGE</td>  <td>该客户端缓存中可用字节数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheSpaceUsed"></a> Client.CacheSpaceUsed</td>  <td>GAUGE</td>  <td>该客户端缓存使用字节数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheSpaceUsedCount"></a> Client.CacheSpaceUsedCount</td>  <td>COUNTER</td>  <td>该客户端缓存用作计数器的字节数量</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheState"></a> Client.CacheState</td>  <td>COUNTER</td>  <td>缓存状态：0（不在使用中），1（只读），2（读写）</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheStoreDeleteTimeout"></a> Client.CacheStoreDeleteTimeout</td>  <td>COUNTER</td>  <td>该客户端从页存储中删除页超时次数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheStoreGetTimeout"></a> Client.CacheStoreGetTimeout</td>  <td>COUNTER</td>  <td>该客户端从页存储中读取页超时次数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheStorePutTimeout"></a> Client.CacheStorePutTimeout</td>  <td>COUNTER</td>  <td>该客户端向页存储中写入新页超时次数</td></tr>
<tr>  <td><a class="anchor" name="Client.CacheStoreThreadsRejected"></a> Client.CacheStoreThreadsRejected</td>  <td>COUNTER</td>  <td>该客户端向线程池提交任务时拒绝 I/O 线程的次数，可能是由于本地文件系统无响应。</td></tr>
<tr>  <td><a class="anchor" name="Client.CloseAlluxioOutStreamLatency"></a> Client.CloseAlluxioOutStreamLatency</td>  <td>TIMER</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Client.CloseUFSOutStreamLatency"></a> Client.CloseUFSOutStreamLatency</td>  <td>TIMER</td>  <td></td></tr>
<tr>  <td><a class="anchor" name="Client.DefaultHiveClientCount"></a> Client.DefaultHiveClientCount</td>  <td>COUNTER</td>  <td>DefaultHiveClientPool 中实例数量</td></tr>
<tr>  <td><a class="anchor" name="Client.FileSystemMasterClientCount"></a> Client.FileSystemMasterClientCount</td>  <td>COUNTER</td>  <td>FileSystemMasterClientPool 中实例数量</td></tr>
<tr>  <td><a class="anchor" name="Client.MetadataCacheSize"></a> Client.MetadataCacheSize</td>  <td>GAUGE</td>  <td>该客户端被缓存的文件和目录的元数据总数。只在文件系统为 alluxio.client.file.MetadataCachingBaseFileSystem 时有效</td></tr>

</tbody></table>

## Fuse 指标

Fuse 是长期运行的 Alluxio 客户端。 
根据启动方式，Fuse 指标将显示为：
* 当文件系统客户端在独立的 AlluxioFuse 进程中启动时，显示为客户端指标。
* 当 Fuse 客户端嵌入在 AlluxioWorker 进程中时，显示为 worker 指标。

Fuse metrics includes:

<table class="table table-striped">
<tbody><tr><th>描述</th><th>类型</th><th>描述</th></tr>

  <tr>
    <td><a class="anchor" name="Fuse.CachedPathCount"></a> Fuse.CachedPathCount</td>
    <td>GAUGE</td>
    <td>缓存的 Alluxio 路径映射的总数。这个值小于或等于 alluxio.fuse.cached.paths.max</td>
  </tr>

  <tr>
    <td><a class="anchor" name="Fuse.ReadWriteFileCount"></a> Fuse.ReadWriteFileCount</td>
    <td>GAUGE</td>
    <td>当前被打开的读写文件数量</td>
  </tr>

  <tr>
    <td><a class="anchor" name="Fuse.TotalCalls"></a> Fuse.TotalCalls</td>
    <td>TIMER</td>
    <td>JNI FUSE 操作调用的吞吐量。此指标表明 Alluxio Fuse 应用处理请求的繁忙程度</td>
  </tr>

</tbody></table>

Fuse 读/写文件数量可用作 Fuse 应用程序压力的指标。
如果在短时间内发生大量并发读/写操作，则每个读/写操作可能需要更长的时间来完成。

当用户或应用程序在 Fuse 挂载点下运行文件系统命令时，该命令将由操作系统处理和转换，并触发在 [AlluxioFuse](https://github.com/Alluxio/alluxio/blob/db01aae966849e88d342a71609ab3d910215afeb/integration/fuse/src/main/java/alluxio/fuse/AlluxioJniFuseFileSystem.java) 中暴露的相关 Fuse 操作。每个操作被调用的次数以及每次调用的持续时间将使用动态指标名称 `Fuse.<FUSE_OPERATION_NAME>` 记录。


重要的 Fuse 指标包括：

| 名称 | 描述 |
|-------------------------|-----------------------------------------------------|
| Fuse.readdir | 列出目录的持续时间指标 |
| Fuse.getattr | 获取文件元数据的持续时间指标 |
| Fuse.open | 打开文件进行读或覆写的持续时间指标 |
| Fuse.read | 读取文件的一部分的持续时间指标 |
| Fuse.create | 为了写入创建文件的持续时间指标 |
| Fuse.write | 写入文件的持续时间指标 |
| Fuse.release | 在读取或写入后关闭文件的持续时间指标。请注意，释放是异步的，因此 FUSE 线程不会等待释放完成 |
| Fuse.mkdir | 创建目录的持续时间指标 |
| Fuse.unlink | 删除文件或目录的持续时间指标 |
| Fuse.rename | 重命名文件或目录的持续时间指标 |
| Fuse.chmod | 更改文件或目录模式的持续时间指标 |
| Fuse.chown | 修改文件或目录的用户和/或组所有权的持续时间指标 |

Fuse相关的指标包括:
* `Client.TotalRPCClients` 显示用于连接到或可连接到 master 或 worker 进行操作的 RPC 客户端的总数。
* 带有 `Direct` 关键字的 worker 指标。当 Fuse 嵌入到 worker 进程中时，它可以通过 worker 内部 API 从该 worker 读取/写入。
相关指标以 `Direct` 结尾。例如，`Worker.BytesReadDirect` 显示该 worker 为其嵌入的 Fuse 客户端提供读取的字节数。
* 如果配置了 `alluxio.user.block.read.metrics.enabled=true`，则会记录 `Client.BlockReadChunkRemote`。 该指标显示通过 gRPC 从远程 worker 读取数据的持续时间统计。

`Client.TotalRPCClients` 和 `Fuse.TotalCalls` 指标是 Fuse 应用程序当前负载的优秀指标。
如果在 Alluxio Fuse 上运行应用程序（e.g. Tensorflow），但这两个指标值比之前低得多，则训练作业可能会卡在 Alluxio 上。

## 普通进程指标

在每个实例（Master、Worker 或 Client）上收集的指标。

### JVM Attributes

| 名称 | 描述 |
|-------------------------|-----------------------------------------------------|
| name | JVM 名称 |
| uptime | JVM 的运行时间 |
| vendor | 当前的 JVM 供应商 |

### GC 统计

| 名称 | 描述 |
|-------------------------|-----------------------------------------------------|
| PS-MarkSweep.count | 标记和清除 old gen 的总数 |
| PS-MarkSweep.time | 标记和清除 old gen 的总时间 |
| PS-Scavenge.count | 清除 young gen 总数 |
| PS-Scavenge.time | 清除 young gen 总时间 |

### 内存使用情况

Alluxio 提供整体和详细的内存使用信息。
每个进程中代码缓存、压缩类空间、元数据空间、PS Eden 空间、PS old gen 以及 PS survivor 空间的详细内存使用信息都会被收集。

以下是内存使用指标的子集：

| 名称 | 描述 |
|------------------------------|-----------------------------------------------------|
| total.committed | 保证可供 JVM 使用的以字节为单位的内存数量 |
| total.init | 可供 JVM 使用的以字节为单位的内存数量 |
| total.max | 以字节为单位的 JVM 可用的最大内存量 |
| total.used | 以字节为单位当前使用的内存大小 |
| heap.committed | 在堆上保证可用的内存大小 |
| heap.init | 初始化时堆上可用的内存量 |
| heap.max | 在堆上可用的最大内存量 |
| heap.usage | 堆上当前正在使用的以 GB 为单位的内存量 |
| heap.used | 堆上当前已经使用过的以 GB 为单位的内存量 |
| pools.Code-Cache.used | 内存池中用于编译和存储本地代码的内存总量 |
| pools.Compressed-Class-Space.used | 内存池中用于类元数据的内存总量 |
| pools.PS-Eden-Space.used | 内存池中用于大多数对象初始分配的内存总量 |
| pools.PS-Survivor-Space.used | 从包含在 Eden space 的垃圾回收中幸存下来的对象的池中使用的内存总量 |

### 类加载统计

| 名称 | 描述 |
|-------------------------|-----------------------------------------------------|
| loaded | 加载的类总数 |
| unloaded | 未加载的类总量 |

### 线程统计

| 名称 | 描述 |
|-------------------------|-----------------------------------------------------|
| count | 当前存活线程数 |
| daemon.count | 当前守护线程的数量 |
| peak.count | 存活线程数峰值 |
| total_started.count | 启动线程总数 |
| deadlock.count | 死锁线程总数 |
| deadlock | 与每个线程有关的死锁的调用栈 |
| new.count | 有新状态的线程数 |
| blocked.count | 阻塞态线程数 |
| runnable.count | 可运行状态线程数 |
| terminated.count | 终结态线程数 |
| timed_waiting.count | 定时等待状态的线程数量 |
