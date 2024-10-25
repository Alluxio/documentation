# 配置项列表

所有Alluxio配置属性都属于以下六类之一：
[共有配置项](#共有配置项)（由Master和Worker共享），
[Master配置项](#master配置项)，[Worker配置项](#worker配置项)，
[用户配置项](#用户配置项)，[集群管理配置项](#集群管理配置项)（用于在诸如Mesos和YARN的集群管理器上运行Alluxio）
以及[安全性配置项](#安全性配置项)（由Master，Worker和用户共享）。

## 共有配置项

共有配置项包含了不同组件共享的常量。

<table class="table table-striped">
<tbody><tr><th>属性名</th><th>默认值</th><th>描述</th></tr>

  <tr>
    <td><a class="anchor" name="alluxio.conf.dynamic.update.enabled"></a> alluxio.conf.dynamic.update.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.debug"></a> alluxio.debug</td>
    <td>false</td>
    <td>设置为true后即可启用debug模式，会记录更多日志，且在Web UI中会显示更多信息。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.exit.collect.info"></a> alluxio.exit.collect.info</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.auth.policy.class"></a> alluxio.fuse.auth.policy.class</td>
    <td>alluxio.fuse.auth.LaunchUserGroupAuthPolicy</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.auth.policy.custom.group"></a> alluxio.fuse.auth.policy.custom.group</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.auth.policy.custom.user"></a> alluxio.fuse.auth.policy.custom.user</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.cached.paths.max"></a> alluxio.fuse.cached.paths.max</td>
    <td>500</td>
    <td>用于FUSE conversion的Alluxio缓存路径的最大数量。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.debug.enabled"></a> alluxio.fuse.debug.enabled</td>
    <td>false</td>
    <td>以debug模式运行FUSE，让fuse进程记录每个FS请求。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.fs.name"></a> alluxio.fuse.fs.name</td>
    <td>alluxio-fuse</td>
    <td>FUSE文件系统名。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.jnifuse.enabled"></a> alluxio.fuse.jnifuse.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.jnifuse.libfuse.version"></a> alluxio.fuse.jnifuse.libfuse.version</td>
    <td>2</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.logging.threshold"></a> alluxio.fuse.logging.threshold</td>
    <td>10s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.mount.alluxio.path"></a> alluxio.fuse.mount.alluxio.path</td>
    <td>/</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.mount.options"></a> alluxio.fuse.mount.options</td>
    <td>attr_timeout=600,entry_timeout=600</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.mount.point"></a> alluxio.fuse.mount.point</td>
    <td>/mnt/alluxio-fuse</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.shared.caching.reader.enabled"></a> alluxio.fuse.shared.caching.reader.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.special.command.enabled"></a> alluxio.fuse.special.command.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.stat.cache.refresh.interval"></a> alluxio.fuse.stat.cache.refresh.interval</td>
    <td>5min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.umount.timeout"></a> alluxio.fuse.umount.timeout</td>
    <td>0s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.user.group.translation.enabled"></a> alluxio.fuse.user.group.translation.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.web.bind.host"></a> alluxio.fuse.web.bind.host</td>
    <td>0.0.0.0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.web.enabled"></a> alluxio.fuse.web.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.web.hostname"></a> alluxio.fuse.web.hostname</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.fuse.web.port"></a> alluxio.fuse.web.port</td>
    <td>49999</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.grpc.reflection.enabled"></a> alluxio.grpc.reflection.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.hadoop.kerberos.keytab.login.autorenewal"></a> alluxio.hadoop.kerberos.keytab.login.autorenewal</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.hadoop.security.authentication"></a> alluxio.hadoop.security.authentication</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.hadoop.security.krb5.conf"></a> alluxio.hadoop.security.krb5.conf</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.home"></a> alluxio.home</td>
    <td>/opt/alluxio</td>
    <td>Alluxio安装目录。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.batch.size"></a> alluxio.job.batch.size</td>
    <td>20</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.bind.host"></a> alluxio.job.master.bind.host</td>
    <td>0.0.0.0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.client.threads"></a> alluxio.job.master.client.threads</td>
    <td>1024</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.embedded.journal.addresses"></a> alluxio.job.master.embedded.journal.addresses</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.embedded.journal.port"></a> alluxio.job.master.embedded.journal.port</td>
    <td>20003</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.finished.job.purge.count"></a> alluxio.job.master.finished.job.purge.count</td>
    <td>-1</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.finished.job.retention.time"></a> alluxio.job.master.finished.job.retention.time</td>
    <td>60sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.hostname"></a> alluxio.job.master.hostname</td>
    <td>${alluxio.master.hostname}</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.job.capacity"></a> alluxio.job.master.job.capacity</td>
    <td>100000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.lost.worker.interval"></a> alluxio.job.master.lost.worker.interval</td>
    <td>1sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.network.flowcontrol.window"></a> alluxio.job.master.network.flowcontrol.window</td>
    <td>2MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.network.keepalive.time"></a> alluxio.job.master.network.keepalive.time</td>
    <td>2h</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.network.keepalive.timeout"></a> alluxio.job.master.network.keepalive.timeout</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.network.max.inbound.message.size"></a> alluxio.job.master.network.max.inbound.message.size</td>
    <td>100MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.network.permit.keepalive.time"></a> alluxio.job.master.network.permit.keepalive.time</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.rpc.addresses"></a> alluxio.job.master.rpc.addresses</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.rpc.port"></a> alluxio.job.master.rpc.port</td>
    <td>20001</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.web.bind.host"></a> alluxio.job.master.web.bind.host</td>
    <td>0.0.0.0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.web.hostname"></a> alluxio.job.master.web.hostname</td>
    <td>${alluxio.job.master.hostname}</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.web.port"></a> alluxio.job.master.web.port</td>
    <td>20002</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.worker.heartbeat.interval"></a> alluxio.job.master.worker.heartbeat.interval</td>
    <td>1sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.master.worker.timeout"></a> alluxio.job.master.worker.timeout</td>
    <td>60sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.request.batch.size"></a> alluxio.job.request.batch.size</td>
    <td>1</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.retention.time"></a> alluxio.job.retention.time</td>
    <td>1d</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.worker.bind.host"></a> alluxio.job.worker.bind.host</td>
    <td>0.0.0.0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.worker.data.port"></a> alluxio.job.worker.data.port</td>
    <td>30002</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.worker.hostname"></a> alluxio.job.worker.hostname</td>
    <td>${alluxio.worker.hostname}</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.worker.rpc.port"></a> alluxio.job.worker.rpc.port</td>
    <td>30001</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.worker.threadpool.size"></a> alluxio.job.worker.threadpool.size</td>
    <td>10</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.worker.throttling"></a> alluxio.job.worker.throttling</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.worker.web.bind.host"></a> alluxio.job.worker.web.bind.host</td>
    <td>0.0.0.0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.job.worker.web.port"></a> alluxio.job.worker.web.port</td>
    <td>30003</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.jvm.monitor.info.threshold"></a> alluxio.jvm.monitor.info.threshold</td>
    <td>1sec</td>
    <td>额外的睡眠时间超过这个阈值，记录INFO。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.jvm.monitor.sleep.interval"></a> alluxio.jvm.monitor.sleep.interval</td>
    <td>1sec</td>
    <td>JVM monitor线程的睡眠时间。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.jvm.monitor.warn.threshold"></a> alluxio.jvm.monitor.warn.threshold</td>
    <td>10sec</td>
    <td>额外的睡眠时间超过这个阈值，记录WARN。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.leak.detector.exit.on.leak"></a> alluxio.leak.detector.exit.on.leak</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.leak.detector.level"></a> alluxio.leak.detector.level</td>
    <td>DISABLED</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.locality.compare.node.ip"></a> alluxio.locality.compare.node.ip</td>
    <td>false</td>
    <td>是否在本地性检查时尝试解析IP地址。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.logserver.hostname"></a> alluxio.logserver.hostname</td>
    <td></td>
    <td>Alluxio logserver主机名。注意:必须将此属性指定为JVM属性;它在alluxio-site.properties不被接受。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.logserver.logs.dir"></a> alluxio.logserver.logs.dir</td>
    <td>${alluxio.work.dir}/logs</td>
    <td>远程日志文件默认位置。注意:必须将此属性指定为JVM属性;它在alluxio-site.properties不被接受。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.logserver.port"></a> alluxio.logserver.port</td>
    <td>45600</td>
    <td>从Alluxio server接收日志的默认端口号。注意:必须将此属性指定为JVM属性;它在alluxio-site.properties不被接受。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.logserver.threads.max"></a> alluxio.logserver.threads.max</td>
    <td>2048</td>
    <td>logserver用于响应日志请求的最大线程数。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.logserver.threads.min"></a> alluxio.logserver.threads.min</td>
    <td>512</td>
    <td>logserver用于响应日志请求的最小线程数。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.metrics.conf.file"></a> alluxio.metrics.conf.file</td>
    <td>${alluxio.conf.dir}/metrics.properties</td>
    <td>度量系统配置文件路径，默认是`conf`文件夹下的`metrics.properties`文件。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.metrics.executor.task.warn.frequency"></a> alluxio.metrics.executor.task.warn.frequency</td>
    <td>5sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.metrics.executor.task.warn.size"></a> alluxio.metrics.executor.task.warn.size</td>
    <td>1000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.network.connection.auth.timeout"></a> alluxio.network.connection.auth.timeout</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.network.connection.health.check.timeout"></a> alluxio.network.connection.health.check.timeout</td>
    <td>5sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.network.connection.server.shutdown.timeout"></a> alluxio.network.connection.server.shutdown.timeout</td>
    <td>60sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.network.connection.shutdown.graceful.timeout"></a> alluxio.network.connection.shutdown.graceful.timeout</td>
    <td>45sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.network.connection.shutdown.timeout"></a> alluxio.network.connection.shutdown.timeout</td>
    <td>15sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.network.host.resolution.timeout"></a> alluxio.network.host.resolution.timeout</td>
    <td>5sec</td>
    <td>Alluxio在启动Master和Worker过程中，需要确保它们在监听外部可达且可解析的主机名。如果没有显式指定主机名，Alluxio会自动尝试选择一个合适的主机名。该配置项指定用于判断一个候选主机名在网络上是否可达的最长等待时间。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.network.ip.address.used"></a> alluxio.network.ip.address.used</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.audit.logging.enabled"></a> alluxio.proxy.audit.logging.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.bucket.naming.restrictions.enabled"></a> alluxio.proxy.s3.bucket.naming.restrictions.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.bucketpathcache.timeout"></a> alluxio.proxy.s3.bucketpathcache.timeout</td>
    <td>0min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.complete.multipart.upload.keepalive.enabled"></a> alluxio.proxy.s3.complete.multipart.upload.keepalive.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.complete.multipart.upload.keepalive.time.interval"></a> alluxio.proxy.s3.complete.multipart.upload.keepalive.time.interval</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.complete.multipart.upload.min.part.size"></a> alluxio.proxy.s3.complete.multipart.upload.min.part.size</td>
    <td>5MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.complete.multipart.upload.pool.size"></a> alluxio.proxy.s3.complete.multipart.upload.pool.size</td>
    <td>20</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.deletetype"></a> alluxio.proxy.s3.deletetype</td>
    <td>ALLUXIO_AND_UFS</td>
    <td>通过S3 API删除桶和对象的删除类型，可选择 `ALLUXIO_AND_UFS`（删除Alluxio和UFS）或`ALLUXIO_ONLY`（只删除Alluxio名空间下）</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.global.read.rate.limit.mb"></a> alluxio.proxy.s3.global.read.rate.limit.mb</td>
    <td>0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.header.metadata.max.size"></a> alluxio.proxy.s3.header.metadata.max.size</td>
    <td>2KB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.multipart.upload.cleaner.enabled"></a> alluxio.proxy.s3.multipart.upload.cleaner.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.multipart.upload.cleaner.pool.size"></a> alluxio.proxy.s3.multipart.upload.cleaner.pool.size</td>
    <td>1</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.multipart.upload.cleaner.retry.count"></a> alluxio.proxy.s3.multipart.upload.cleaner.retry.count</td>
    <td>3</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.multipart.upload.cleaner.retry.delay"></a> alluxio.proxy.s3.multipart.upload.cleaner.retry.delay</td>
    <td>10sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.multipart.upload.cleaner.timeout"></a> alluxio.proxy.s3.multipart.upload.cleaner.timeout</td>
    <td>10min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.single.connection.read.rate.limit.mb"></a> alluxio.proxy.s3.single.connection.read.rate.limit.mb</td>
    <td>0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.tagging.restrictions.enabled"></a> alluxio.proxy.s3.tagging.restrictions.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.v2.async.heavy.pool.core.thread.number"></a> alluxio.proxy.s3.v2.async.heavy.pool.core.thread.number</td>
    <td>8</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.v2.async.heavy.pool.maximum.thread.number"></a> alluxio.proxy.s3.v2.async.heavy.pool.maximum.thread.number</td>
    <td>64</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.v2.async.heavy.pool.queue.size"></a> alluxio.proxy.s3.v2.async.heavy.pool.queue.size</td>
    <td>65536</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.v2.async.light.pool.core.thread.number"></a> alluxio.proxy.s3.v2.async.light.pool.core.thread.number</td>
    <td>8</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.v2.async.light.pool.maximum.thread.number"></a> alluxio.proxy.s3.v2.async.light.pool.maximum.thread.number</td>
    <td>64</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.v2.async.light.pool.queue.size"></a> alluxio.proxy.s3.v2.async.light.pool.queue.size</td>
    <td>65536</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.v2.async.processing.enabled"></a> alluxio.proxy.s3.v2.async.processing.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.v2.version.enabled"></a> alluxio.proxy.s3.v2.version.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.s3.writetype"></a> alluxio.proxy.s3.writetype</td>
    <td>CACHE_THROUGH</td>
    <td>通过S3 API创建桶和对象的写类型，可选择"MUST_CACHE"(只写入Alluxio，必须存储在Alluxio中)、"CACHE_THROUGH"(尝试缓存，同步写入到UnderFS)、"THROUGH"(无缓存，同步写入到UnderFS)。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.stream.cache.timeout"></a> alluxio.proxy.stream.cache.timeout</td>
    <td>1hour</td>
    <td>在代理中输入和输出流缓存回收的时限。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.web.bind.host"></a> alluxio.proxy.web.bind.host</td>
    <td>0.0.0.0</td>
    <td>Alluxio代理网络服务器运行的主机名。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.web.hostname"></a> alluxio.proxy.web.hostname</td>
    <td></td>
    <td>Alluxio代理的web UI绑定的主机名。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.proxy.web.port"></a> alluxio.proxy.web.port</td>
    <td>39999</td>
    <td>Alluxio代理的web UI绑定的端口。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.s3.rest.authentication.enabled"></a> alluxio.s3.rest.authentication.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.s3.rest.authenticator.classname"></a> alluxio.s3.rest.authenticator.classname</td>
    <td>alluxio.proxy.s3.auth.PassAllAuthenticator</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.secondary.master.metastore.dir"></a> alluxio.secondary.master.metastore.dir</td>
    <td>${alluxio.work.dir}/secondary-metastore</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.site.conf.dir"></a> alluxio.site.conf.dir</td>
    <td>${alluxio.conf.dir}/,${user.home}/.alluxio/,/etc/alluxio/</td>
    <td>加载配置文件时默认的搜索路径</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.site.conf.rocks.block.file"></a> alluxio.site.conf.rocks.block.file</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.site.conf.rocks.inode.file"></a> alluxio.site.conf.rocks.inode.file</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.standalone.fuse.jvm.monitor.enabled"></a> alluxio.standalone.fuse.jvm.monitor.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.standby.master.metrics.sink.enabled"></a> alluxio.standby.master.metrics.sink.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.standby.master.web.enabled"></a> alluxio.standby.master.web.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.table.catalog.path"></a> alluxio.table.catalog.path</td>
    <td>/catalog</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.table.catalog.udb.sync.timeout"></a> alluxio.table.catalog.udb.sync.timeout</td>
    <td>1h</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.table.enabled"></a> alluxio.table.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.table.journal.partitions.chunk.size"></a> alluxio.table.journal.partitions.chunk.size</td>
    <td>500</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.table.load.default.replication"></a> alluxio.table.load.default.replication</td>
    <td>1</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.table.transform.manager.job.history.retention.time"></a> alluxio.table.transform.manager.job.history.retention.time</td>
    <td>300sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.table.transform.manager.job.monitor.interval"></a> alluxio.table.transform.manager.job.monitor.interval</td>
    <td>10s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.table.udb.hive.clientpool.MAX"></a> alluxio.table.udb.hive.clientpool.MAX</td>
    <td>256</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.table.udb.hive.clientpool.min"></a> alluxio.table.udb.hive.clientpool.min</td>
    <td>16</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.test.deprecated.key"></a> alluxio.test.deprecated.key</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.tmp.dirs"></a> alluxio.tmp.dirs</td>
    <td>/tmp</td>
    <td>存储Alluxio临时文件的路径，使用逗号作为分隔符。如果指定了多个路径，每个临时文件将随机选择一个路径。目前，只有上传到对象存储的文件存储在这些路径中。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.allow.set.owner.failure"></a> alluxio.underfs.allow.set.owner.failure</td>
    <td>false</td>
    <td>是否允许UFS中的设置所有者失败。当设置为true时，文件或目录所有者可能会在Alluxio和UFS之间产生分歧。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.cephfs.auth.id"></a> alluxio.underfs.cephfs.auth.id</td>
    <td>admin</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.cephfs.auth.key"></a> alluxio.underfs.cephfs.auth.key</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.cephfs.auth.keyfile"></a> alluxio.underfs.cephfs.auth.keyfile</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.cephfs.auth.keyring"></a> alluxio.underfs.cephfs.auth.keyring</td>
    <td>/etc/ceph/ceph.client.admin.keyring</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.cephfs.conf.file"></a> alluxio.underfs.cephfs.conf.file</td>
    <td>/etc/ceph/ceph.conf</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.cephfs.conf.options"></a> alluxio.underfs.cephfs.conf.options</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.cephfs.localize.reads"></a> alluxio.underfs.cephfs.localize.reads</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.cephfs.mds.namespace"></a> alluxio.underfs.cephfs.mds.namespace</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.cephfs.mon.host"></a> alluxio.underfs.cephfs.mon.host</td>
    <td>0.0.0.0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.cephfs.mount.gid"></a> alluxio.underfs.cephfs.mount.gid</td>
    <td>0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.cephfs.mount.point"></a> alluxio.underfs.cephfs.mount.point</td>
    <td>/</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.cephfs.mount.uid"></a> alluxio.underfs.cephfs.mount.uid</td>
    <td>0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.cleanup.enabled"></a> alluxio.underfs.cleanup.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.cleanup.interval"></a> alluxio.underfs.cleanup.interval</td>
    <td>1day</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.eventual.consistency.retry.base.sleep"></a> alluxio.underfs.eventual.consistency.retry.base.sleep</td>
    <td>50ms</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.eventual.consistency.retry.max.num"></a> alluxio.underfs.eventual.consistency.retry.max.num</td>
    <td>0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.eventual.consistency.retry.max.sleep"></a> alluxio.underfs.eventual.consistency.retry.max.sleep</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.gcs.default.mode"></a> alluxio.underfs.gcs.default.mode</td>
    <td>0700</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.gcs.directory.suffix"></a> alluxio.underfs.gcs.directory.suffix</td>
    <td>/</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.gcs.owner.id.to.username.mapping"></a> alluxio.underfs.gcs.owner.id.to.username.mapping</td>
    <td></td>
    <td>可选配置项，指定一个预设的gcs拥有者ID到Alluxio用户名的静态映射，格式为"id1=user1;id2=user2"。谷歌云存储的ID可以在控制台地址https://console.cloud.google.com/storage/settings找到。请使用"Owners"选项。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.gcs.retry.delay.multiplier"></a> alluxio.underfs.gcs.retry.delay.multiplier</td>
    <td>2</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.gcs.retry.initial.delay"></a> alluxio.underfs.gcs.retry.initial.delay</td>
    <td>1000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.gcs.retry.jitter"></a> alluxio.underfs.gcs.retry.jitter</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.gcs.retry.max"></a> alluxio.underfs.gcs.retry.max</td>
    <td>60</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.gcs.retry.max.delay"></a> alluxio.underfs.gcs.retry.max.delay</td>
    <td>1min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.gcs.retry.total.duration"></a> alluxio.underfs.gcs.retry.total.duration</td>
    <td>5min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.gcs.version"></a> alluxio.underfs.gcs.version</td>
    <td>2</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.hdfs.configuration"></a> alluxio.underfs.hdfs.configuration</td>
    <td>${alluxio.conf.dir}/core-site.xml:${alluxio.conf.dir}/hdfs-site.xml</td>
    <td>hdfs配置文件的位置。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.hdfs.impl"></a> alluxio.underfs.hdfs.impl</td>
    <td>org.apache.hadoop.hdfs.DistributedFileSystem</td>
    <td>作为底层存储系统的hdfs的实现类。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.hdfs.prefixes"></a> alluxio.underfs.hdfs.prefixes</td>
    <td>hdfs://,glusterfs:///</td>
    <td>可选配置项，指定以哪些前缀开头的文件应该存放在Apache Hadoop底层文件系统。分隔符为任何空白符或者','。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.hdfs.remote"></a> alluxio.underfs.hdfs.remote</td>
    <td>true</td>
    <td>底层存储系统的worker节点对于Alluxio worker节点来说是否是远程的。如果该值为true，那么Alluxio将不会尝试从底层存储系统获取locality相关信息，因为这种情况下不可能存在任何locality，这样做可以提高性能。默认值为false。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.io.threads"></a> alluxio.underfs.io.threads</td>
    <td>Use 3*{CPU core count} for UFS IO.</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.kodo.connect.timeout"></a> alluxio.underfs.kodo.connect.timeout</td>
    <td>50sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.kodo.downloadhost"></a> alluxio.underfs.kodo.downloadhost</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.kodo.endpoint"></a> alluxio.underfs.kodo.endpoint</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.kodo.requests.max"></a> alluxio.underfs.kodo.requests.max</td>
    <td>64</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.listing.length"></a> alluxio.underfs.listing.length</td>
    <td>1000</td>
    <td>底层文件系统在一次查询中可以列出的目录条目的最大数量。如果条目数量大于指定长度，则需要多次查询才能列出。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.local.skip.broken.symlinks"></a> alluxio.underfs.local.skip.broken.symlinks</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.logging.threshold"></a> alluxio.underfs.logging.threshold</td>
    <td>10s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.object.store.breadcrumbs.enabled"></a> alluxio.underfs.object.store.breadcrumbs.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.object.store.mount.shared.publicly"></a> alluxio.underfs.object.store.mount.shared.publicly</td>
    <td>false</td>
    <td>是否对所有Alluxio用户共享挂载的对象存储系统。 注意，该配置对HDFS或者本地文件系统没有任何影响。默认值是false。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.object.store.multi.range.chunk.size"></a> alluxio.underfs.object.store.multi.range.chunk.size</td>
    <td>${alluxio.user.block.size.bytes.default}</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.object.store.service.threads"></a> alluxio.underfs.object.store.service.threads</td>
    <td>20</td>
    <td>存储底层文件系统操作的并行对象在执行程序池中的线程数。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.object.store.skip.parent.directory.creation"></a> alluxio.underfs.object.store.skip.parent.directory.creation</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.object.store.streaming.upload.part.timeout"></a> alluxio.underfs.object.store.streaming.upload.part.timeout</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.obs.intermediate.upload.clean.age"></a> alluxio.underfs.obs.intermediate.upload.clean.age</td>
    <td>3day</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.obs.streaming.upload.enabled"></a> alluxio.underfs.obs.streaming.upload.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.obs.streaming.upload.partition.size"></a> alluxio.underfs.obs.streaming.upload.partition.size</td>
    <td>64MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.obs.streaming.upload.threads"></a> alluxio.underfs.obs.streaming.upload.threads</td>
    <td>20</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.oss.connection.max"></a> alluxio.underfs.oss.connection.max</td>
    <td>1024</td>
    <td>OSS连接最大数目。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.oss.connection.timeout"></a> alluxio.underfs.oss.connection.timeout</td>
    <td>50sec</td>
    <td>OSS连接时限。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.oss.connection.ttl"></a> alluxio.underfs.oss.connection.ttl</td>
    <td>-1</td>
    <td>OSS连接的TTL（ms）。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.oss.ecs.ram.role"></a> alluxio.underfs.oss.ecs.ram.role</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.oss.intermediate.upload.clean.age"></a> alluxio.underfs.oss.intermediate.upload.clean.age</td>
    <td>3day</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.oss.retry.max"></a> alluxio.underfs.oss.retry.max</td>
    <td>3</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.oss.socket.timeout"></a> alluxio.underfs.oss.socket.timeout</td>
    <td>50sec</td>
    <td>OSS套接字时限。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.oss.streaming.upload.enabled"></a> alluxio.underfs.oss.streaming.upload.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.oss.streaming.upload.partition.size"></a> alluxio.underfs.oss.streaming.upload.partition.size</td>
    <td>64MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.oss.streaming.upload.threads"></a> alluxio.underfs.oss.streaming.upload.threads</td>
    <td>20</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.oss.sts.ecs.metadata.service.endpoint"></a> alluxio.underfs.oss.sts.ecs.metadata.service.endpoint</td>
    <td>http://100.100.100.200/latest/meta-data/ram/security-credentials/</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.oss.sts.enabled"></a> alluxio.underfs.oss.sts.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.oss.sts.token.refresh.interval.ms"></a> alluxio.underfs.oss.sts.token.refresh.interval.ms</td>
    <td>30m</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.ozone.prefixes"></a> alluxio.underfs.ozone.prefixes</td>
    <td>o3fs://,ofs://</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.persistence.async.temp.dir"></a> alluxio.underfs.persistence.async.temp.dir</td>
    <td>.alluxio_ufs_persistence</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.admin.threads.max"></a> alluxio.underfs.s3.admin.threads.max</td>
    <td>20</td>
    <td>与S3通信时进行元数据操作所使用的最大线程数目，这些操作会并发且频繁执行，但不会花费太多时间。默认值为20。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.connection.ttl"></a> alluxio.underfs.s3.connection.ttl</td>
    <td>-1</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.default.mode"></a> alluxio.underfs.s3.default.mode</td>
    <td>0700</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.directory.suffix"></a> alluxio.underfs.s3.directory.suffix</td>
    <td>/</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.disable.dns.buckets"></a> alluxio.underfs.s3.disable.dns.buckets</td>
    <td>false</td>
    <td>可选配置项，指定所有S3请求路径样式。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.endpoint"></a> alluxio.underfs.s3.endpoint</td>
    <td></td>
    <td>可选配置项，在组织AWS服务请求的时候可以指定某个区域地址来降低数据延迟或者访问某些隔离在不同AWS Region的资源。 一个endpoint是某个服务的一个入口地址。举例，s3.cn-north-1.amazonaws.com.cn 就是一个北京区域的亚马逊S3服务的一个endpoint。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.endpoint.region"></a> alluxio.underfs.s3.endpoint.region</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.inherit.acl"></a> alluxio.underfs.s3.inherit.acl</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.intermediate.upload.clean.age"></a> alluxio.underfs.s3.intermediate.upload.clean.age</td>
    <td>3day</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.list.objects.v1"></a> alluxio.underfs.s3.list.objects.v1</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.max.error.retry"></a> alluxio.underfs.s3.max.error.retry</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.owner.id.to.username.mapping"></a> alluxio.underfs.s3.owner.id.to.username.mapping</td>
    <td></td>
    <td>可选配置项，指定一个预设的s3规范ID到Alluxio用户名的静态映射，格式为"id1=user1;id2=user2"。AWS的s3规范ID可以在控制台地址https://console.aws.amazon.com/iam/home?#security_credential找到。请展开"Account Identifiers"选项卡，并参考"Canonical User ID"。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.proxy.host"></a> alluxio.underfs.s3.proxy.host</td>
    <td></td>
    <td>可选配置项，指定与S3通信的代理主机。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.proxy.port"></a> alluxio.underfs.s3.proxy.port</td>
    <td></td>
    <td>可选配置项，指定与S3通信的代理端口。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.region"></a> alluxio.underfs.s3.region</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.request.timeout"></a> alluxio.underfs.s3.request.timeout</td>
    <td>1min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.secure.http.enabled"></a> alluxio.underfs.s3.secure.http.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.server.side.encryption.enabled"></a> alluxio.underfs.s3.server.side.encryption.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.signer.algorithm"></a> alluxio.underfs.s3.signer.algorithm</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.socket.timeout"></a> alluxio.underfs.s3.socket.timeout</td>
    <td>50sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.streaming.upload.enabled"></a> alluxio.underfs.s3.streaming.upload.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.streaming.upload.partition.size"></a> alluxio.underfs.s3.streaming.upload.partition.size</td>
    <td>64MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.threads.max"></a> alluxio.underfs.s3.threads.max</td>
    <td>40</td>
    <td>与S3通信时使用的最大线程数目和最大并发连接数。包括数据上载以及元数据操作线程。该数目至少要等于最大管理线程与最大上传线程数目之和。默认值为40,即默认管理线程与默认上传线程池大小之和。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.s3.upload.threads.max"></a> alluxio.underfs.s3.upload.threads.max</td>
    <td>20</td>
    <td>进行多部分上传时上传数据到S3所使用的最大线程数目，这些操作会相对耗时，然而由于某些线程会引起错误，过多线程会导致带宽饥饿。默认值为2。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.strict.version.match.enabled"></a> alluxio.underfs.strict.version.match.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.web.connnection.timeout"></a> alluxio.underfs.web.connnection.timeout</td>
    <td>60s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.web.header.last.modified"></a> alluxio.underfs.web.header.last.modified</td>
    <td>EEE, dd MMM yyyy HH:mm:ss zzz</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.web.parent.names"></a> alluxio.underfs.web.parent.names</td>
    <td>Parent Directory,..,../</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.underfs.web.titles"></a> alluxio.underfs.web.titles</td>
    <td>Index of,Directory listing for</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.web.cors.allow.credential"></a> alluxio.web.cors.allow.credential</td>
    <td>false</td>
    <td>是否可以将对请求的响应暴露给页面。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.web.cors.allow.headers"></a> alluxio.web.cors.allow.headers</td>
    <td>*</td>
    <td>哪些头部信息是允许的，使用 * 允许所有的任何头部信息。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.web.cors.allow.methods"></a> alluxio.web.cors.allow.methods</td>
    <td>*</td>
    <td>哪些方法是允许的，使用 * 允许所有的任何方法。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.web.cors.allow.origins"></a> alluxio.web.cors.allow.origins</td>
    <td>*</td>
    <td>哪些 Origin 是允许的，使用 * 允许所有的任何 Origin。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.web.cors.enabled"></a> alluxio.web.cors.enabled</td>
    <td>false</td>
    <td>启用 Web 用户界面的 CORS 配置</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.web.cors.exposed.headers"></a> alluxio.web.cors.exposed.headers</td>
    <td>*</td>
    <td>在访问跨源资源时，允许在响应中设置哪些头部信息。 使用 * 允许所有头部信息。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.web.cors.max.age"></a> alluxio.web.cors.max.age</td>
    <td>-1</td>
    <td>在访问跨源资源时，结果可以被缓存的最大秒数。"-1表示不缓存。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.web.file.info.enabled"></a> alluxio.web.file.info.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.web.refresh.interval"></a> alluxio.web.refresh.interval</td>
    <td>15s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.web.threaddump.log.enabled"></a> alluxio.web.threaddump.log.enabled</td>
    <td>false</td>
    <td>当访问 threaddump api 时是否也将线程信息打印到日志</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.web.threads"></a> alluxio.web.threads</td>
    <td>1</td>
    <td>web服务器的线程数目。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.web.ui.enabled"></a> alluxio.web.ui.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.work.dir"></a> alluxio.work.dir</td>
    <td>${alluxio.home}</td>
    <td>Alluxio的工作目录。默认情况下，journal、logs以及底层文件系统的数据（如果使用本地文件系统）都会写在该目录下。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.zookeeper.address"></a> alluxio.zookeeper.address</td>
    <td></td>
    <td>ZooKeeper地址。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.zookeeper.auth.enabled"></a> alluxio.zookeeper.auth.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.zookeeper.connection.timeout"></a> alluxio.zookeeper.connection.timeout</td>
    <td>15s</td>
    <td>连接到Zookeeper时的连接时限。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.zookeeper.election.path"></a> alluxio.zookeeper.election.path</td>
    <td>/alluxio/election</td>
    <td>ZooKeeper的选举文件夹。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.zookeeper.enabled"></a> alluxio.zookeeper.enabled</td>
    <td>false</td>
    <td>若为true，则使用zooKeeper启动master容错机制。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.zookeeper.job.election.path"></a> alluxio.zookeeper.job.election.path</td>
    <td>/alluxio/job_election</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.zookeeper.job.leader.path"></a> alluxio.zookeeper.job.leader.path</td>
    <td>/alluxio/job_leader</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.zookeeper.leader.connection.error.policy"></a> alluxio.zookeeper.leader.connection.error.policy</td>
    <td>SESSION</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.zookeeper.leader.inquiry.retry"></a> alluxio.zookeeper.leader.inquiry.retry</td>
    <td>10</td>
    <td>从ZooKeeper申请leader的最大请求次数。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.zookeeper.leader.path"></a> alluxio.zookeeper.leader.path</td>
    <td>/alluxio/leader</td>
    <td>ZooKeeper的leader文件夹。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.zookeeper.session.timeout"></a> alluxio.zookeeper.session.timeout</td>
    <td>60s</td>
    <td>连接到Zookeeper时的会话时限。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.azure.account.oauth2.client.endpoint"></a> fs.azure.account.oauth2.client.endpoint</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.azure.account.oauth2.client.id"></a> fs.azure.account.oauth2.client.id</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.azure.account.oauth2.client.secret"></a> fs.azure.account.oauth2.client.secret</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.azure.account.oauth2.msi.endpoint"></a> fs.azure.account.oauth2.msi.endpoint</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.azure.account.oauth2.msi.tenant"></a> fs.azure.account.oauth2.msi.tenant</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.cos.access.key"></a> fs.cos.access.key</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.cos.app.id"></a> fs.cos.app.id</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.cos.connection.max"></a> fs.cos.connection.max</td>
    <td>1024</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.cos.connection.timeout"></a> fs.cos.connection.timeout</td>
    <td>50sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.cos.region"></a> fs.cos.region</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.cos.secret.key"></a> fs.cos.secret.key</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.cos.socket.timeout"></a> fs.cos.socket.timeout</td>
    <td>50sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.gcs.accessKeyId"></a> fs.gcs.accessKeyId</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.gcs.credential.path"></a> fs.gcs.credential.path</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.gcs.secretAccessKey"></a> fs.gcs.secretAccessKey</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.kodo.accesskey"></a> fs.kodo.accesskey</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.kodo.secretkey"></a> fs.kodo.secretkey</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.obs.accessKey"></a> fs.obs.accessKey</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.obs.bucketType"></a> fs.obs.bucketType</td>
    <td>obs</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.obs.endpoint"></a> fs.obs.endpoint</td>
    <td>obs.myhwclouds.com</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.obs.secretKey"></a> fs.obs.secretKey</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.oss.accessKeyId"></a> fs.oss.accessKeyId</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.oss.accessKeySecret"></a> fs.oss.accessKeySecret</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.oss.endpoint"></a> fs.oss.endpoint</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.swift.auth.method"></a> fs.swift.auth.method</td>
    <td></td>
    <td>选择身份验证方法：[tempauth (default), swiftauth, keystone, keystonev3]</td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.swift.auth.url"></a> fs.swift.auth.url</td>
    <td></td>
    <td>REST服务器的认证URL，例如http://server:8090/auth/v1.0</td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.swift.password"></a> fs.swift.password</td>
    <td></td>
    <td>user:tenant认证的密码</td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.swift.region"></a> fs.swift.region</td>
    <td></td>
    <td>Keystone认证的服务地区</td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.swift.simulation"></a> fs.swift.simulation</td>
    <td></td>
    <td>是否模拟单节点Swift后台用于验证，默认为否</td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.swift.tenant"></a> fs.swift.tenant</td>
    <td></td>
    <td>Swift user用于认证</td>
  </tr>

  <tr>
    <td><a class="anchor" name="fs.swift.user"></a> fs.swift.user</td>
    <td></td>
    <td>Swift tenant用于认证</td>
  </tr>

  <tr>
    <td><a class="anchor" name="s3a.accessKeyId"></a> s3a.accessKeyId</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="s3a.secretKey"></a> s3a.secretKey</td>
    <td></td>
    <td></td>
  </tr>

</tbody></table>

## Master配置项

Master配置项指定master节点的信息，例如地址和端口号。

<table class="table table-striped">
<tbody><tr><th>属性名</th><th>默认值</th><th>描述</th></tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.audit.logging.enabled"></a> alluxio.master.audit.logging.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.audit.logging.queue.capacity"></a> alluxio.master.audit.logging.queue.capacity</td>
    <td>10000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.backup.abandon.timeout"></a> alluxio.master.backup.abandon.timeout</td>
    <td>1min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.backup.connect.interval.max"></a> alluxio.master.backup.connect.interval.max</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.backup.connect.interval.min"></a> alluxio.master.backup.connect.interval.min</td>
    <td>1sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.backup.delegation.enabled"></a> alluxio.master.backup.delegation.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.backup.directory"></a> alluxio.master.backup.directory</td>
    <td>/alluxio_backups</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.backup.entry.buffer.count"></a> alluxio.master.backup.entry.buffer.count</td>
    <td>10000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.backup.heartbeat.interval"></a> alluxio.master.backup.heartbeat.interval</td>
    <td>2sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.backup.state.lock.exclusive.duration"></a> alluxio.master.backup.state.lock.exclusive.duration</td>
    <td>0ms</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.backup.state.lock.forced.duration"></a> alluxio.master.backup.state.lock.forced.duration</td>
    <td>15min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.backup.state.lock.interrupt.cycle.enabled"></a> alluxio.master.backup.state.lock.interrupt.cycle.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.backup.state.lock.interrupt.cycle.interval"></a> alluxio.master.backup.state.lock.interrupt.cycle.interval</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.backup.suspend.timeout"></a> alluxio.master.backup.suspend.timeout</td>
    <td>3min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.backup.transport.timeout"></a> alluxio.master.backup.transport.timeout</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.bind.host"></a> alluxio.master.bind.host</td>
    <td>0.0.0.0</td>
    <td>Alluxio master绑定的主机名。参考<a href="#configure-multihomed-networks">多宿主网络</a></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.block.scan.invalid.batch.max.size"></a> alluxio.master.block.scan.invalid.batch.max.size</td>
    <td>10000000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.container.id.reservation.size"></a> alluxio.master.container.id.reservation.size</td>
    <td>1000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.daily.backup.enabled"></a> alluxio.master.daily.backup.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.daily.backup.files.retained"></a> alluxio.master.daily.backup.files.retained</td>
    <td>3</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.daily.backup.state.lock.grace.mode"></a> alluxio.master.daily.backup.state.lock.grace.mode</td>
    <td>TIMEOUT</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.daily.backup.state.lock.sleep.duration"></a> alluxio.master.daily.backup.state.lock.sleep.duration</td>
    <td>5m</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.daily.backup.state.lock.timeout"></a> alluxio.master.daily.backup.state.lock.timeout</td>
    <td>1h</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.daily.backup.state.lock.try.duration"></a> alluxio.master.daily.backup.state.lock.try.duration</td>
    <td>2m</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.daily.backup.time"></a> alluxio.master.daily.backup.time</td>
    <td>05:00</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.addresses"></a> alluxio.master.embedded.journal.addresses</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.catchup.retry.wait"></a> alluxio.master.embedded.journal.catchup.retry.wait</td>
    <td>1s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.election.timeout.max"></a> alluxio.master.embedded.journal.election.timeout.max</td>
    <td>20s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.election.timeout.min"></a> alluxio.master.embedded.journal.election.timeout.min</td>
    <td>10s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.entry.size.max"></a> alluxio.master.embedded.journal.entry.size.max</td>
    <td>10MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.flush.size.max"></a> alluxio.master.embedded.journal.flush.size.max</td>
    <td>160MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.port"></a> alluxio.master.embedded.journal.port</td>
    <td>19200</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.raft.client.request.interval"></a> alluxio.master.embedded.journal.raft.client.request.interval</td>
    <td>100ms</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.raft.client.request.timeout"></a> alluxio.master.embedded.journal.raft.client.request.timeout</td>
    <td>60sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.ratis.config"></a> alluxio.master.embedded.journal.ratis.config</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.retry.cache.expiry.time"></a> alluxio.master.embedded.journal.retry.cache.expiry.time</td>
    <td>60s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.snapshot.replication.chunk.size"></a> alluxio.master.embedded.journal.snapshot.replication.chunk.size</td>
    <td>4MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.snapshot.replication.compression.level"></a> alluxio.master.embedded.journal.snapshot.replication.compression.level</td>
    <td>1</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.snapshot.replication.compression.type"></a> alluxio.master.embedded.journal.snapshot.replication.compression.type</td>
    <td>NO_COMPRESSION</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.transport.max.inbound.message.size"></a> alluxio.master.embedded.journal.transport.max.inbound.message.size</td>
    <td>100MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.transport.request.timeout.ms"></a> alluxio.master.embedded.journal.transport.request.timeout.ms</td>
    <td>5sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.unsafe.flush.enabled"></a> alluxio.master.embedded.journal.unsafe.flush.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.embedded.journal.write.timeout"></a> alluxio.master.embedded.journal.write.timeout</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.failover.collect.info"></a> alluxio.master.failover.collect.info</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.file.access.time.journal.flush.interval"></a> alluxio.master.file.access.time.journal.flush.interval</td>
    <td>1h</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.file.access.time.update.precision"></a> alluxio.master.file.access.time.update.precision</td>
    <td>1d</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.file.access.time.updater.enabled"></a> alluxio.master.file.access.time.updater.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.file.access.time.updater.shutdown.timeout"></a> alluxio.master.file.access.time.updater.shutdown.timeout</td>
    <td>1sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.filesystem.liststatus.result.message.length"></a> alluxio.master.filesystem.liststatus.result.message.length</td>
    <td>10000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.filesystem.merge.inode.journals"></a> alluxio.master.filesystem.merge.inode.journals</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.filesystem.operation.retry.cache.enabled"></a> alluxio.master.filesystem.operation.retry.cache.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.filesystem.operation.retry.cache.size"></a> alluxio.master.filesystem.operation.retry.cache.size</td>
    <td>100000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.format.file.prefix"></a> alluxio.master.format.file.prefix</td>
    <td>_format_</td>
    <td>当journal被格式化时，在joural文件夹下生成的文件的文件名前缀。当判断journal是否被格式化时master会查找文件名以该前缀开头的文件。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.heartbeat.timeout"></a> alluxio.master.heartbeat.timeout</td>
    <td>10min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.hostname"></a> alluxio.master.hostname</td>
    <td></td>
    <td>Alluxio master主机名。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.backup.when.corrupted"></a> alluxio.master.journal.backup.when.corrupted</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.catchup.protect.enabled"></a> alluxio.master.journal.catchup.protect.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.checkpoint.period.entries"></a> alluxio.master.journal.checkpoint.period.entries</td>
    <td>2000000</td>
    <td>在创建一个新journal检查点之前写入的journal数。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.exit.on.demotion"></a> alluxio.master.journal.exit.on.demotion</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.flush.batch.time"></a> alluxio.master.journal.flush.batch.time</td>
    <td>100ms</td>
    <td>等待批处理日志写入的时间。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.flush.timeout"></a> alluxio.master.journal.flush.timeout</td>
    <td>5min</td>
    <td>在放弃和关闭master之前保持重试日志写入的时间量。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.folder"></a> alluxio.master.journal.folder</td>
    <td>${alluxio.work.dir}/journal</td>
    <td>存储master journal日志的路径。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.gc.period"></a> alluxio.master.journal.gc.period</td>
    <td>2min</td>
    <td>扫描和删除陈旧的journal检查点的频率。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.gc.threshold"></a> alluxio.master.journal.gc.threshold</td>
    <td>5min</td>
    <td>垃圾收集检查点的最小年龄。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.init.from.backup"></a> alluxio.master.journal.init.from.backup</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.local.log.compaction"></a> alluxio.master.journal.local.log.compaction</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.log.size.bytes.max"></a> alluxio.master.journal.log.size.bytes.max</td>
    <td>10MB</td>
    <td>如果一个日志文件大小超过该值，会产生下一个文件。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.request.data.timeout"></a> alluxio.master.journal.request.data.timeout</td>
    <td>20000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.request.info.timeout"></a> alluxio.master.journal.request.info.timeout</td>
    <td>10000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.retry.interval"></a> alluxio.master.journal.retry.interval</td>
    <td>1sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.space.monitor.interval"></a> alluxio.master.journal.space.monitor.interval</td>
    <td>10min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.space.monitor.percent.free.threshold"></a> alluxio.master.journal.space.monitor.percent.free.threshold</td>
    <td>10</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.tailer.shutdown.quiet.wait.time"></a> alluxio.master.journal.tailer.shutdown.quiet.wait.time</td>
    <td>5sec</td>
    <td>在备用master停止监听线程之前，在该配置项指定的时间内不应对leader master的journal作任何更新。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.tailer.sleep.time"></a> alluxio.master.journal.tailer.sleep.time</td>
    <td>1sec</td>
    <td>指定当备用master无法检测到leader master journal的更新时，其睡眠时间。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.temporary.file.gc.threshold"></a> alluxio.master.journal.temporary.file.gc.threshold</td>
    <td>30min</td>
    <td>临时文件垃圾收集检查点的最小年龄。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.type"></a> alluxio.master.journal.type</td>
    <td>EMBEDDED</td>
    <td>使用journal类型，UFS（存储journal在UFS中）和NOOP（不使用journal）。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.journal.ufs.option"></a> alluxio.master.journal.ufs.option</td>
    <td></td>
    <td>journal操作使用的配置。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.jvm.monitor.enabled"></a> alluxio.master.jvm.monitor.enabled</td>
    <td>true</td>
    <td>是否在master上启动JVM monitor线程。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.keytab.file"></a> alluxio.master.keytab.file</td>
    <td></td>
    <td>Alluxio master的Kerberos密钥表文件。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.lock.pool.concurrency.level"></a> alluxio.master.lock.pool.concurrency.level</td>
    <td>100</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.lock.pool.high.watermark"></a> alluxio.master.lock.pool.high.watermark</td>
    <td>1000000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.lock.pool.initsize"></a> alluxio.master.lock.pool.initsize</td>
    <td>1000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.lock.pool.low.watermark"></a> alluxio.master.lock.pool.low.watermark</td>
    <td>500000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.log.config.report.heartbeat.interval"></a> alluxio.master.log.config.report.heartbeat.interval</td>
    <td>1h</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.lost.worker.deletion.timeout"></a> alluxio.master.lost.worker.deletion.timeout</td>
    <td>30min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.lost.worker.detection.interval"></a> alluxio.master.lost.worker.detection.interval</td>
    <td>10sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.lost.worker.file.detection.interval"></a> alluxio.master.lost.worker.file.detection.interval</td>
    <td>5min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.merge.journal.context.num.entries.logging.threshold"></a> alluxio.master.merge.journal.context.num.entries.logging.threshold</td>
    <td>10000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metadata.concurrent.sync.dedup"></a> alluxio.master.metadata.concurrent.sync.dedup</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metadata.sync.concurrency.level"></a> alluxio.master.metadata.sync.concurrency.level</td>
    <td>6</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metadata.sync.executor.pool.size"></a> alluxio.master.metadata.sync.executor.pool.size</td>
    <td>The total number of threads which can concurrently execute metadata sync operations.</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metadata.sync.ignore.ttl"></a> alluxio.master.metadata.sync.ignore.ttl</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metadata.sync.instrument.executor"></a> alluxio.master.metadata.sync.instrument.executor</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metadata.sync.lock.pool.concurrency.level"></a> alluxio.master.metadata.sync.lock.pool.concurrency.level</td>
    <td>20</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metadata.sync.lock.pool.high.watermark"></a> alluxio.master.metadata.sync.lock.pool.high.watermark</td>
    <td>50000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metadata.sync.lock.pool.initsize"></a> alluxio.master.metadata.sync.lock.pool.initsize</td>
    <td>1000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metadata.sync.lock.pool.low.watermark"></a> alluxio.master.metadata.sync.lock.pool.low.watermark</td>
    <td>20000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metadata.sync.traversal.order"></a> alluxio.master.metadata.sync.traversal.order</td>
    <td>BFS</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metadata.sync.ufs.prefetch.pool.size"></a> alluxio.master.metadata.sync.ufs.prefetch.pool.size</td>
    <td>The number of threads which can concurrently fetch metadata from UFSes during a metadata sync operations.</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metadata.sync.ufs.prefetch.status"></a> alluxio.master.metadata.sync.ufs.prefetch.status</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metadata.sync.ufs.prefetch.timeout"></a> alluxio.master.metadata.sync.ufs.prefetch.timeout</td>
    <td>100ms</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore"></a> alluxio.master.metastore</td>
    <td>ROCKS</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.block"></a> alluxio.master.metastore.block</td>
    <td>ROCKS</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.dir"></a> alluxio.master.metastore.dir</td>
    <td>${alluxio.work.dir}/metastore</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.dir.block"></a> alluxio.master.metastore.dir.block</td>
    <td>${alluxio.master.metastore.dir}</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.dir.inode"></a> alluxio.master.metastore.dir.inode</td>
    <td>${alluxio.master.metastore.dir}</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.inode"></a> alluxio.master.metastore.inode</td>
    <td>ROCKS</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.inode.cache.evict.batch.size"></a> alluxio.master.metastore.inode.cache.evict.batch.size</td>
    <td>1000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.inode.cache.high.water.mark.ratio"></a> alluxio.master.metastore.inode.cache.high.water.mark.ratio</td>
    <td>0.85</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.inode.cache.low.water.mark.ratio"></a> alluxio.master.metastore.inode.cache.low.water.mark.ratio</td>
    <td>0.8</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.inode.cache.max.size"></a> alluxio.master.metastore.inode.cache.max.size</td>
    <td>{Max memory of master JVM} / 2 / 2 KB per inode</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.inode.enumerator.buffer.count"></a> alluxio.master.metastore.inode.enumerator.buffer.count</td>
    <td>10000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.inode.inherit.owner.and.group"></a> alluxio.master.metastore.inode.inherit.owner.and.group</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.inode.iteration.crawler.count"></a> alluxio.master.metastore.inode.iteration.crawler.count</td>
    <td>Use {CPU core count} for enumeration.</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.iterator.readahead.size"></a> alluxio.master.metastore.iterator.readahead.size</td>
    <td>64MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.metrics.refresh.interval"></a> alluxio.master.metastore.metrics.refresh.interval</td>
    <td>5s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.block.location.block.index"></a> alluxio.master.metastore.rocks.block.location.block.index</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.block.location.bloom.filter"></a> alluxio.master.metastore.rocks.block.location.bloom.filter</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.block.location.cache.size"></a> alluxio.master.metastore.rocks.block.location.cache.size</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.block.location.index"></a> alluxio.master.metastore.rocks.block.location.index</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.block.meta.block.index"></a> alluxio.master.metastore.rocks.block.meta.block.index</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.block.meta.bloom.filter"></a> alluxio.master.metastore.rocks.block.meta.bloom.filter</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.block.meta.cache.size"></a> alluxio.master.metastore.rocks.block.meta.cache.size</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.block.meta.index"></a> alluxio.master.metastore.rocks.block.meta.index</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.checkpoint.compression.type"></a> alluxio.master.metastore.rocks.checkpoint.compression.type</td>
    <td>LZ4_COMPRESSION</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.edge.block.index"></a> alluxio.master.metastore.rocks.edge.block.index</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.edge.bloom.filter"></a> alluxio.master.metastore.rocks.edge.bloom.filter</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.edge.cache.size"></a> alluxio.master.metastore.rocks.edge.cache.size</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.edge.index"></a> alluxio.master.metastore.rocks.edge.index</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.inode.block.index"></a> alluxio.master.metastore.rocks.inode.block.index</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.inode.bloom.filter"></a> alluxio.master.metastore.rocks.inode.bloom.filter</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.inode.cache.size"></a> alluxio.master.metastore.rocks.inode.cache.size</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.inode.index"></a> alluxio.master.metastore.rocks.inode.index</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.parallel.backup"></a> alluxio.master.metastore.rocks.parallel.backup</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metastore.rocks.parallel.backup.threads"></a> alluxio.master.metastore.rocks.parallel.backup.threads</td>
    <td>The default number of threads used by backing up rocksdb in parallel.</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metrics.file.size.distribution.buckets"></a> alluxio.master.metrics.file.size.distribution.buckets</td>
    <td>1KB,1MB,10MB,100MB,1GB,10GB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metrics.heap.enabled"></a> alluxio.master.metrics.heap.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metrics.service.threads"></a> alluxio.master.metrics.service.threads</td>
    <td>5</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.metrics.time.series.interval"></a> alluxio.master.metrics.time.series.interval</td>
    <td>5min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.mount.table.root.alluxio"></a> alluxio.master.mount.table.root.alluxio</td>
    <td>/</td>
    <td>Alluxio mount根节点。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.mount.table.root.option"></a> alluxio.master.mount.table.root.option</td>
    <td></td>
    <td>Alluxio mount根节点UFS配置。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.mount.table.root.readonly"></a> alluxio.master.mount.table.root.readonly</td>
    <td>false</td>
    <td>Alluxio mount根节点是否只读。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.mount.table.root.shared"></a> alluxio.master.mount.table.root.shared</td>
    <td>true</td>
    <td>Alluxio mount根节点是否共享。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.mount.table.root.ufs"></a> alluxio.master.mount.table.root.ufs</td>
    <td>${alluxio.work.dir}/underFSStorage</td>
    <td>挂载到Alluxio mount根节点的UFS。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.network.flowcontrol.window"></a> alluxio.master.network.flowcontrol.window</td>
    <td>2MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.network.keepalive.time"></a> alluxio.master.network.keepalive.time</td>
    <td>2h</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.network.keepalive.timeout"></a> alluxio.master.network.keepalive.timeout</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.network.max.inbound.message.size"></a> alluxio.master.network.max.inbound.message.size</td>
    <td>100MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.network.permit.keepalive.time"></a> alluxio.master.network.permit.keepalive.time</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.periodic.block.integrity.check.interval"></a> alluxio.master.periodic.block.integrity.check.interval</td>
    <td>1hr</td>
    <td>块完整性检查的间隔，如果小于0则不启用。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.periodic.block.integrity.check.repair"></a> alluxio.master.periodic.block.integrity.check.repair</td>
    <td>true</td>
    <td>完整性检查时是否要删除孤儿块。这是个实验性的属性。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.persistence.blacklist"></a> alluxio.master.persistence.blacklist</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.persistence.checker.interval"></a> alluxio.master.persistence.checker.interval</td>
    <td>1s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.persistence.initial.interval"></a> alluxio.master.persistence.initial.interval</td>
    <td>1s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.persistence.max.interval"></a> alluxio.master.persistence.max.interval</td>
    <td>1hr</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.persistence.max.total.wait.time"></a> alluxio.master.persistence.max.total.wait.time</td>
    <td>1day</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.persistence.scheduler.interval"></a> alluxio.master.persistence.scheduler.interval</td>
    <td>1s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.principal"></a> alluxio.master.principal</td>
    <td></td>
    <td>Alluxio master的Kerberos主体。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.recursive.operation.journal.force.flush.max.entries"></a> alluxio.master.recursive.operation.journal.force.flush.max.entries</td>
    <td>100</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.replication.check.interval"></a> alluxio.master.replication.check.interval</td>
    <td>1min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.rpc.addresses"></a> alluxio.master.rpc.addresses</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.rpc.executor.core.pool.size"></a> alluxio.master.rpc.executor.core.pool.size</td>
    <td>500</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.rpc.executor.fjp.async"></a> alluxio.master.rpc.executor.fjp.async</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.rpc.executor.fjp.min.runnable"></a> alluxio.master.rpc.executor.fjp.min.runnable</td>
    <td>1</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.rpc.executor.fjp.parallelism"></a> alluxio.master.rpc.executor.fjp.parallelism</td>
    <td>2 * {CPU core count}</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.rpc.executor.keepalive"></a> alluxio.master.rpc.executor.keepalive</td>
    <td>60sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.rpc.executor.max.pool.size"></a> alluxio.master.rpc.executor.max.pool.size</td>
    <td>500</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.rpc.executor.tpe.allow.core.threads.timeout"></a> alluxio.master.rpc.executor.tpe.allow.core.threads.timeout</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.rpc.executor.tpe.queue.type"></a> alluxio.master.rpc.executor.tpe.queue.type</td>
    <td>LINKED_BLOCKING_QUEUE</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.rpc.executor.type"></a> alluxio.master.rpc.executor.type</td>
    <td>TPE</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.rpc.port"></a> alluxio.master.rpc.port</td>
    <td>19998</td>
    <td>Alluxio master的运行端口。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.shell.backup.state.lock.grace.mode"></a> alluxio.master.shell.backup.state.lock.grace.mode</td>
    <td>FORCED</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.shell.backup.state.lock.sleep.duration"></a> alluxio.master.shell.backup.state.lock.sleep.duration</td>
    <td>0s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.shell.backup.state.lock.timeout"></a> alluxio.master.shell.backup.state.lock.timeout</td>
    <td>0s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.shell.backup.state.lock.try.duration"></a> alluxio.master.shell.backup.state.lock.try.duration</td>
    <td>0s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.standby.heartbeat.interval"></a> alluxio.master.standby.heartbeat.interval</td>
    <td>2min</td>
    <td>Alluxio master进程间的心跳间隔时间。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.startup.block.integrity.check.enabled"></a> alluxio.master.startup.block.integrity.check.enabled</td>
    <td>false</td>
    <td>是否应该在启动时检查系统孤立的块(由于各种系统故障而没有相应文件但仍然占用系统资源的块)。如果此属性为真，则在主启动期间将删除孤立的块。此属性自1.7.1开始可用。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.state.lock.error.threshold"></a> alluxio.master.state.lock.error.threshold</td>
    <td>20</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.active.cpu.load.ratio"></a> alluxio.master.throttle.active.cpu.load.ratio</td>
    <td>0.5</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.active.heap.gc.time"></a> alluxio.master.throttle.active.heap.gc.time</td>
    <td>1sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.active.heap.used.ratio"></a> alluxio.master.throttle.active.heap.used.ratio</td>
    <td>0.5</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.active.rpc.queue.size"></a> alluxio.master.throttle.active.rpc.queue.size</td>
    <td>50000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.background.enabled"></a> alluxio.master.throttle.background.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.enabled"></a> alluxio.master.throttle.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.filesystem.op.per.sec"></a> alluxio.master.throttle.filesystem.op.per.sec</td>
    <td>2000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.filesystem.rpc.queue.size.limit"></a> alluxio.master.throttle.filesystem.rpc.queue.size.limit</td>
    <td>1000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.foreground.enabled"></a> alluxio.master.throttle.foreground.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.heartbeat.interval"></a> alluxio.master.throttle.heartbeat.interval</td>
    <td>3sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.observed.pit.number"></a> alluxio.master.throttle.observed.pit.number</td>
    <td>3</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.overloaded.cpu.load.ratio"></a> alluxio.master.throttle.overloaded.cpu.load.ratio</td>
    <td>0.95</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.overloaded.heap.gc.time"></a> alluxio.master.throttle.overloaded.heap.gc.time</td>
    <td>10sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.overloaded.heap.used.ratio"></a> alluxio.master.throttle.overloaded.heap.used.ratio</td>
    <td>0.9</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.overloaded.rpc.queue.size"></a> alluxio.master.throttle.overloaded.rpc.queue.size</td>
    <td>150000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.stressed.cpu.load.ratio"></a> alluxio.master.throttle.stressed.cpu.load.ratio</td>
    <td>0.8</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.stressed.heap.gc.time"></a> alluxio.master.throttle.stressed.heap.gc.time</td>
    <td>5sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.stressed.heap.used.ratio"></a> alluxio.master.throttle.stressed.heap.used.ratio</td>
    <td>0.8</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.throttle.stressed.rpc.queue.size"></a> alluxio.master.throttle.stressed.rpc.queue.size</td>
    <td>100000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.tieredstore.global.level0.alias"></a> alluxio.master.tieredstore.global.level0.alias</td>
    <td>MEM</td>
    <td>整个系统中最高存储层的名称。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.tieredstore.global.level1.alias"></a> alluxio.master.tieredstore.global.level1.alias</td>
    <td>SSD</td>
    <td>整个系统中第二存储层的名称。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.tieredstore.global.level2.alias"></a> alluxio.master.tieredstore.global.level2.alias</td>
    <td>HDD</td>
    <td>整个系统中第三存储层的名称。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.tieredstore.global.levels"></a> alluxio.master.tieredstore.global.levels</td>
    <td>3</td>
    <td>系统中存储层的总数目。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.tieredstore.global.mediumtype"></a> alluxio.master.tieredstore.global.mediumtype</td>
    <td>MEM,SSD,HDD</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.ttl.checker.interval"></a> alluxio.master.ttl.checker.interval</td>
    <td>1hour</td>
    <td>清除过期ttl值的文件任务的时间间隔。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.ufs.active.sync.event.rate.interval"></a> alluxio.master.ufs.active.sync.event.rate.interval</td>
    <td>60sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.ufs.active.sync.interval"></a> alluxio.master.ufs.active.sync.interval</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.ufs.active.sync.max.activities"></a> alluxio.master.ufs.active.sync.max.activities</td>
    <td>10</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.ufs.active.sync.max.age"></a> alluxio.master.ufs.active.sync.max.age</td>
    <td>10</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.ufs.active.sync.poll.batch.size"></a> alluxio.master.ufs.active.sync.poll.batch.size</td>
    <td>1024</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.ufs.active.sync.poll.timeout"></a> alluxio.master.ufs.active.sync.poll.timeout</td>
    <td>10sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.ufs.active.sync.retry.timeout"></a> alluxio.master.ufs.active.sync.retry.timeout</td>
    <td>10sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.ufs.active.sync.thread.pool.size"></a> alluxio.master.ufs.active.sync.thread.pool.size</td>
    <td>The number of threads used by the active sync provider process active sync events. A higher number allow the master to use more CPU to process events from an event stream in parallel. If this value is too low, Alluxio may fall behind processing events. Defaults to # of processors / 2.</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.ufs.block.location.cache.capacity"></a> alluxio.master.ufs.block.location.cache.capacity</td>
    <td>1000000</td>
    <td>UFS块缓存的容量。这个cache缓存UFS块位置，适用于要保存但不在Alluxio空间中的文件，以便这些文件的列表状态不需要反复询问UFS的块位置。如果将此设置为0，则缓存将被禁用。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.ufs.journal.max.catchup.time"></a> alluxio.master.ufs.journal.max.catchup.time</td>
    <td>10min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.ufs.path.cache.capacity"></a> alluxio.master.ufs.path.cache.capacity</td>
    <td>100000</td>
    <td>UFS路径缓存的容量。此缓存用来近似`一次性`元数据加载行为。（查看 `alluxio.user.file.metadata.load.type`)。更大的缓存将耗费更大的内存，但是能够更好地近似`一次性`行为。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.ufs.path.cache.threads"></a> alluxio.master.ufs.path.cache.threads</td>
    <td>64</td>
    <td>线程池（可异步处理路径，用于缓存UFS路径）的最大容积。更多的线程数将减少异步缓存中的staleness数量，但可能会影响性能。 如果设置为0，缓存将被禁用，而alluxio.user.file.metadata.load.type = Once将表现为"Always"。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.unsafe.direct.persist.object.enabled"></a> alluxio.master.unsafe.direct.persist.object.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.update.check.enabled"></a> alluxio.master.update.check.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.update.check.interval"></a> alluxio.master.update.check.interval</td>
    <td>7day</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.web.bind.host"></a> alluxio.master.web.bind.host</td>
    <td>0.0.0.0</td>
    <td>Alluxio master web UI绑定的主机名。参考<a href="#configure-multihomed-networks">多宿主网络</a></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.web.hostname"></a> alluxio.master.web.hostname</td>
    <td></td>
    <td>提供Alluxio Master web UI的主机名。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.web.in.alluxio.data.page.count"></a> alluxio.master.web.in.alluxio.data.page.count</td>
    <td>1000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.web.port"></a> alluxio.master.web.port</td>
    <td>19999</td>
    <td>Alluxio web UI运行端口。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.whitelist"></a> alluxio.master.whitelist</td>
    <td>/</td>
    <td>以该配置中的前缀开头的路径是可缓存的，这些前缀用分号隔开。Alluxio在第一次读这些文件时会尝试缓存这些可缓存的文件。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.worker.connect.wait.time"></a> alluxio.master.worker.connect.wait.time</td>
    <td>5sec</td>
    <td>在开始接受client请求之前，Alluxio master会等待一段时间，让所有worker注册。此属性决定等待时间。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.worker.info.cache.refresh.time"></a> alluxio.master.worker.info.cache.refresh.time</td>
    <td>10sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.worker.register.lease.count"></a> alluxio.master.worker.register.lease.count</td>
    <td>25</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.worker.register.lease.enabled"></a> alluxio.master.worker.register.lease.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.worker.register.lease.respect.jvm.space"></a> alluxio.master.worker.register.lease.respect.jvm.space</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.worker.register.lease.ttl"></a> alluxio.master.worker.register.lease.ttl</td>
    <td>1min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.worker.register.stream.response.timeout"></a> alluxio.master.worker.register.stream.response.timeout</td>
    <td>10min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.master.worker.timeout"></a> alluxio.master.worker.timeout</td>
    <td>5min</td>
    <td>Alluxio master与worker之间响应的最大超时时间，超过该时间表明该worker失效。</td>
  </tr>

</tbody></table>

## Worker配置项

Worker配置项指定worker节点的信息，例如地址和端口号。

<table class="table table-striped">
<tbody><tr><th>属性名</th><th>默认值</th><th>描述</th></tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.allocator.class"></a> alluxio.worker.allocator.class</td>
    <td>alluxio.worker.block.allocator.MaxFreeAllocator</td>
    <td>worker在特定存储层上分配不同存储目录空间的策略，有效值包括：`alluxio.worker.block.allocator.MaxFreeAllocator`, `alluxio.worker.block.allocator.GreedyAllocator`, `alluxio.worker.block.allocator.RoundRobinAllocator`。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.bind.host"></a> alluxio.worker.bind.host</td>
    <td>0.0.0.0</td>
    <td>Alluxio worker节点绑定的主机名，参考<a href="#configure-multihomed-networks">多宿主网络</a></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.block.annotator.class"></a> alluxio.worker.block.annotator.class</td>
    <td>alluxio.worker.block.annotator.LRUAnnotator</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.block.annotator.lrfu.attenuation.factor"></a> alluxio.worker.block.annotator.lrfu.attenuation.factor</td>
    <td>2.0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.block.annotator.lrfu.step.factor"></a> alluxio.worker.block.annotator.lrfu.step.factor</td>
    <td>0.25</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.block.heartbeat.interval"></a> alluxio.worker.block.heartbeat.interval</td>
    <td>1sec</td>
    <td>worker心跳时间间隔。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.block.heartbeat.report.size.threshold"></a> alluxio.worker.block.heartbeat.report.size.threshold</td>
    <td>1000000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.block.heartbeat.timeout"></a> alluxio.worker.block.heartbeat.timeout</td>
    <td>${alluxio.worker.master.connect.retry.timeout}</td>
    <td>worker心跳超时时间。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.block.master.client.pool.size"></a> alluxio.worker.block.master.client.pool.size</td>
    <td>11</td>
    <td>block master在Alluxio worker上的client池容量。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.block.store.type"></a> alluxio.worker.block.store.type</td>
    <td>FILE</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.container.hostname"></a> alluxio.worker.container.hostname</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.data.folder"></a> alluxio.worker.data.folder</td>
    <td>/alluxioworker/</td>
    <td>每个存储目录中的一个相对路径，该路径被Alluxio worker用作层次化存储中存放数据的文件夹。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.data.folder.permissions"></a> alluxio.worker.data.folder.permissions</td>
    <td>rwxrwxrwx</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.data.folder.tmp"></a> alluxio.worker.data.folder.tmp</td>
    <td>.tmp_blocks</td>
    <td>相对于 alluxio.worker.data.folder 的路径, 用于存放临时数据.</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.data.server.domain.socket.address"></a> alluxio.worker.data.server.domain.socket.address</td>
    <td></td>
    <td>domain socket 路径。如果设置，Alluxio worker 通过这个路径读写数据。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.data.server.domain.socket.as.uuid"></a> alluxio.worker.data.server.domain.socket.as.uuid</td>
    <td>false</td>
    <td>如果为真，则属性worker.data.server.domain.socket是域套接字的主目录的路径，也是唯一标识符用作域套接字名称。此外，客户端忽略alluxio.user.hostname在检测本地工作人员进行短路操作时。如果为false，则该属性是UNIX域套接字的绝对路径。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.data.tmp.subdir.max"></a> alluxio.worker.data.tmp.subdir.max</td>
    <td>1024</td>
    <td>在 alluxio.worker.data.folder.tmp 中可以创建的文件夹的最大数目.</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.evictor.class"></a> alluxio.worker.evictor.class</td>
    <td></td>
    <td>当某个存储层空间不足时，worker剔除块文件的策略。可选值包括`alluxio.worker.block.evictor.LRFUEvictor`、 `alluxio.worker.block.evictor.GreedyEvictor`、 `alluxio.worker.block.evictor.LRUEvictor`。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.free.space.timeout"></a> alluxio.worker.free.space.timeout</td>
    <td>10sec</td>
    <td>worker等待驱逐来为客户端写请求提供空间的持续时间。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.fuse.enabled"></a> alluxio.worker.fuse.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.hostname"></a> alluxio.worker.hostname</td>
    <td></td>
    <td>Alluxio worker的主机名。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.jvm.monitor.enabled"></a> alluxio.worker.jvm.monitor.enabled</td>
    <td>true</td>
    <td>是否在worker上启用JVM monitor线程。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.keytab.file"></a> alluxio.worker.keytab.file</td>
    <td></td>
    <td>Alluxio worker的Kerberos密钥对文件。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.management.backoff.strategy"></a> alluxio.worker.management.backoff.strategy</td>
    <td>ANY</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.management.block.transfer.concurrency.limit"></a> alluxio.worker.management.block.transfer.concurrency.limit</td>
    <td>Use {CPU core count}/2 threads block transfer.</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.management.load.detection.cool.down.time"></a> alluxio.worker.management.load.detection.cool.down.time</td>
    <td>10sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.management.task.thread.count"></a> alluxio.worker.management.task.thread.count</td>
    <td>Use {CPU core count} threads for all management tasks.</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.management.tier.align.enabled"></a> alluxio.worker.management.tier.align.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.management.tier.align.range"></a> alluxio.worker.management.tier.align.range</td>
    <td>100</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.management.tier.align.reserved.bytes"></a> alluxio.worker.management.tier.align.reserved.bytes</td>
    <td>1GB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.management.tier.promote.enabled"></a> alluxio.worker.management.tier.promote.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.management.tier.promote.quota.percent"></a> alluxio.worker.management.tier.promote.quota.percent</td>
    <td>90</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.management.tier.promote.range"></a> alluxio.worker.management.tier.promote.range</td>
    <td>100</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.management.tier.swap.restore.enabled"></a> alluxio.worker.management.tier.swap.restore.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.master.connect.retry.timeout"></a> alluxio.worker.master.connect.retry.timeout</td>
    <td>1hour</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.master.periodical.rpc.timeout"></a> alluxio.worker.master.periodical.rpc.timeout</td>
    <td>5min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.async.cache.manager.queue.max"></a> alluxio.worker.network.async.cache.manager.queue.max</td>
    <td>512</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.async.cache.manager.threads.max"></a> alluxio.worker.network.async.cache.manager.threads.max</td>
    <td>2 * {CPU core count}</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.block.reader.threads.max"></a> alluxio.worker.network.block.reader.threads.max</td>
    <td>2048</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.block.writer.threads.max"></a> alluxio.worker.network.block.writer.threads.max</td>
    <td>1024</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.flowcontrol.window"></a> alluxio.worker.network.flowcontrol.window</td>
    <td>2MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.keepalive.time"></a> alluxio.worker.network.keepalive.time</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.keepalive.timeout"></a> alluxio.worker.network.keepalive.timeout</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.max.inbound.message.size"></a> alluxio.worker.network.max.inbound.message.size</td>
    <td>4MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.netty.boss.threads"></a> alluxio.worker.network.netty.boss.threads</td>
    <td>1</td>
    <td>收到新的请求时启用的线程数目。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.netty.channel"></a> alluxio.worker.network.netty.channel</td>
    <td>EPOLL</td>
    <td>netty通道类型：NIO或EPOLL。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.netty.shutdown.quiet.period"></a> alluxio.worker.network.netty.shutdown.quiet.period</td>
    <td>2sec</td>
    <td>沉默期时间长度。当netty服务器正终止时，要确保在该时间段内不会产生RPC调用。如果出现了RPC调用，那么在该netty服务器终止时会该沉默期会重新开始。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.netty.watermark.high"></a> alluxio.worker.network.netty.watermark.high</td>
    <td>32KB</td>
    <td>在切换到不可写状态之前，写队列中可存放的最大字节数。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.netty.watermark.low"></a> alluxio.worker.network.netty.watermark.low</td>
    <td>8KB</td>
    <td>一旦写队列中的high watermark达到了，该队列在切换到可写状态之前必须刷新到该配置项指定的low watermark。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.netty.worker.threads"></a> alluxio.worker.network.netty.worker.threads</td>
    <td>2 * {CPU core count}</td>
    <td>处理请求的线程数目，0表示#cpuCores * 2</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.permit.keepalive.time"></a> alluxio.worker.network.permit.keepalive.time</td>
    <td>30s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.reader.buffer.pooled"></a> alluxio.worker.network.reader.buffer.pooled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.reader.buffer.size"></a> alluxio.worker.network.reader.buffer.size</td>
    <td>4MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.reader.max.chunk.size.bytes"></a> alluxio.worker.network.reader.max.chunk.size.bytes</td>
    <td>2MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.shutdown.timeout"></a> alluxio.worker.network.shutdown.timeout</td>
    <td>15sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.writer.buffer.size.messages"></a> alluxio.worker.network.writer.buffer.size.messages</td>
    <td>8</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.network.zerocopy.enabled"></a> alluxio.worker.network.zerocopy.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.page.store.async.restore.enabled"></a> alluxio.worker.page.store.async.restore.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.page.store.async.write.enabled"></a> alluxio.worker.page.store.async.write.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.page.store.async.write.threads"></a> alluxio.worker.page.store.async.write.threads</td>
    <td>16</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.page.store.dirs"></a> alluxio.worker.page.store.dirs</td>
    <td>/tmp/alluxio_cache</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.page.store.eviction.retries"></a> alluxio.worker.page.store.eviction.retries</td>
    <td>10</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.page.store.evictor.class"></a> alluxio.worker.page.store.evictor.class</td>
    <td>alluxio.client.file.cache.evictor.LRUCacheEvictor</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.page.store.evictor.lfu.logbase"></a> alluxio.worker.page.store.evictor.lfu.logbase</td>
    <td>2.0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.page.store.evictor.nondeterministic.enabled"></a> alluxio.worker.page.store.evictor.nondeterministic.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.page.store.local.store.file.buckets"></a> alluxio.worker.page.store.local.store.file.buckets</td>
    <td>1000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.page.store.overhead"></a> alluxio.worker.page.store.overhead</td>
    <td>0.1</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.page.store.page.size"></a> alluxio.worker.page.store.page.size</td>
    <td>1MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.page.store.quota.enabled"></a> alluxio.worker.page.store.quota.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.page.store.sizes"></a> alluxio.worker.page.store.sizes</td>
    <td>512MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.page.store.timeout.duration"></a> alluxio.worker.page.store.timeout.duration</td>
    <td>-1</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.page.store.timeout.threads"></a> alluxio.worker.page.store.timeout.threads</td>
    <td>32</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.page.store.type"></a> alluxio.worker.page.store.type</td>
    <td>LOCAL</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.principal"></a> alluxio.worker.principal</td>
    <td></td>
    <td>Alluxio worker的Kerberos主体。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.ramdisk.size"></a> alluxio.worker.ramdisk.size</td>
    <td>2/3 of total system memory, or 1GB if system memory size cannot be determined</td>
    <td>每个worker节点的内存容量。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.register.lease.enabled"></a> alluxio.worker.register.lease.enabled</td>
    <td>${alluxio.master.worker.register.lease.enabled}</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.register.lease.retry.max.duration"></a> alluxio.worker.register.lease.retry.max.duration</td>
    <td>${alluxio.worker.master.connect.retry.timeout}</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.register.lease.retry.sleep.max"></a> alluxio.worker.register.lease.retry.sleep.max</td>
    <td>10sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.register.lease.retry.sleep.min"></a> alluxio.worker.register.lease.retry.sleep.min</td>
    <td>1sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.register.stream.batch.size"></a> alluxio.worker.register.stream.batch.size</td>
    <td>1000000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.register.stream.complete.timeout"></a> alluxio.worker.register.stream.complete.timeout</td>
    <td>5min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.register.stream.deadline"></a> alluxio.worker.register.stream.deadline</td>
    <td>15min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.register.stream.enabled"></a> alluxio.worker.register.stream.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.register.stream.response.timeout"></a> alluxio.worker.register.stream.response.timeout</td>
    <td>${alluxio.master.worker.register.stream.response.timeout}</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.register.to.all.masters"></a> alluxio.worker.register.to.all.masters</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.remote.io.slow.threshold"></a> alluxio.worker.remote.io.slow.threshold</td>
    <td>10s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.reviewer.class"></a> alluxio.worker.reviewer.class</td>
    <td>alluxio.worker.block.reviewer.ProbabilisticBufferReviewer</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.reviewer.probabilistic.hardlimit.bytes"></a> alluxio.worker.reviewer.probabilistic.hardlimit.bytes</td>
    <td>64MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.reviewer.probabilistic.softlimit.bytes"></a> alluxio.worker.reviewer.probabilistic.softlimit.bytes</td>
    <td>256MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.rpc.executor.core.pool.size"></a> alluxio.worker.rpc.executor.core.pool.size</td>
    <td>100</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.rpc.executor.fjp.async"></a> alluxio.worker.rpc.executor.fjp.async</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.rpc.executor.fjp.min.runnable"></a> alluxio.worker.rpc.executor.fjp.min.runnable</td>
    <td>1</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.rpc.executor.fjp.parallelism"></a> alluxio.worker.rpc.executor.fjp.parallelism</td>
    <td>2 * {CPU core count}</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.rpc.executor.keepalive"></a> alluxio.worker.rpc.executor.keepalive</td>
    <td>60sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.rpc.executor.max.pool.size"></a> alluxio.worker.rpc.executor.max.pool.size</td>
    <td>1000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.rpc.executor.tpe.allow.core.threads.timeout"></a> alluxio.worker.rpc.executor.tpe.allow.core.threads.timeout</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.rpc.executor.tpe.queue.type"></a> alluxio.worker.rpc.executor.tpe.queue.type</td>
    <td>LINKED_BLOCKING_QUEUE_WITH_CAP</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.rpc.executor.type"></a> alluxio.worker.rpc.executor.type</td>
    <td>TPE</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.rpc.port"></a> alluxio.worker.rpc.port</td>
    <td>29999</td>
    <td>Alluxio worker节点运行端口。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.session.timeout"></a> alluxio.worker.session.timeout</td>
    <td>1min</td>
    <td>worker和client连接的超时时间，超时后表明该会话失效。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.startup.timeout"></a> alluxio.worker.startup.timeout</td>
    <td>10min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.storage.checker.enabled"></a> alluxio.worker.storage.checker.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.block.lock.readers"></a> alluxio.worker.tieredstore.block.lock.readers</td>
    <td>1000</td>
    <td>一个Alluxio数据块锁最大允许的并行读数目。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.block.locks"></a> alluxio.worker.tieredstore.block.locks</td>
    <td>1000</td>
    <td>一个Alluxio数据块worker的数据块锁数目。较大值会达到更好的锁粒度，但会使用更多空间。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.free.ahead.bytes"></a> alluxio.worker.tieredstore.free.ahead.bytes</td>
    <td>0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level0.alias"></a> alluxio.worker.tieredstore.level0.alias</td>
    <td>MEM</td>
    <td>在worker上最高存储层的别名，该值一定要对应master配置项中全局存储层之一。禁止将全局继承结构中较低级别存储层的别名放在worker中较高级别，因此默认情况下，在任何worker上SSD都不能在MEM之前。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level0.dirs.mediumtype"></a> alluxio.worker.tieredstore.level0.dirs.mediumtype</td>
    <td>${alluxio.worker.tieredstore.level0.alias}</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level0.dirs.path"></a> alluxio.worker.tieredstore.level0.dirs.path</td>
    <td>/mnt/ramdisk on Linux, /Volumes/ramdisk on OSX</td>
    <td>顶层存储层在存储目录中的路径。注意对于MacoS该值应为`/Volumes/`。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level0.dirs.quota"></a> alluxio.worker.tieredstore.level0.dirs.quota</td>
    <td>${alluxio.worker.ramdisk.size}</td>
    <td>顶层存储层容量。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level0.watermark.high.ratio"></a> alluxio.worker.tieredstore.level0.watermark.high.ratio</td>
    <td>0.95</td>
    <td>在顶层存储层中的高水位比例 （取值为0到1之间）。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level0.watermark.low.ratio"></a> alluxio.worker.tieredstore.level0.watermark.low.ratio</td>
    <td>0.7</td>
    <td>在顶层存储层中的低水位比例 （取值为0到1之间）。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level1.alias"></a> alluxio.worker.tieredstore.level1.alias</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level1.dirs.mediumtype"></a> alluxio.worker.tieredstore.level1.dirs.mediumtype</td>
    <td>${alluxio.worker.tieredstore.level1.alias}</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level1.dirs.path"></a> alluxio.worker.tieredstore.level1.dirs.path</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level1.dirs.quota"></a> alluxio.worker.tieredstore.level1.dirs.quota</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level1.watermark.high.ratio"></a> alluxio.worker.tieredstore.level1.watermark.high.ratio</td>
    <td>0.95</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level1.watermark.low.ratio"></a> alluxio.worker.tieredstore.level1.watermark.low.ratio</td>
    <td>0.7</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level2.alias"></a> alluxio.worker.tieredstore.level2.alias</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level2.dirs.mediumtype"></a> alluxio.worker.tieredstore.level2.dirs.mediumtype</td>
    <td>${alluxio.worker.tieredstore.level2.alias}</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level2.dirs.path"></a> alluxio.worker.tieredstore.level2.dirs.path</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level2.dirs.quota"></a> alluxio.worker.tieredstore.level2.dirs.quota</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level2.watermark.high.ratio"></a> alluxio.worker.tieredstore.level2.watermark.high.ratio</td>
    <td>0.95</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.level2.watermark.low.ratio"></a> alluxio.worker.tieredstore.level2.watermark.low.ratio</td>
    <td>0.7</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.tieredstore.levels"></a> alluxio.worker.tieredstore.levels</td>
    <td>1</td>
    <td>worker上的存储层数目。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.ufs.instream.cache.enabled"></a> alluxio.worker.ufs.instream.cache.enabled</td>
    <td>true</td>
    <td>在存储输入流下启用缓存，以便以后在同一个文件上查找操作可重用缓存的输入流。这将提高位置读取性能，因为一些文件系统的打开操作是昂贵的。当UFS文件被修改时，缓存的输入流将过时，而不通知Alluxio。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.ufs.instream.cache.expiration.time"></a> alluxio.worker.ufs.instream.cache.expiration.time</td>
    <td>5min</td>
    <td>缓存的UFS输入流过期时间。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.ufs.instream.cache.max.size"></a> alluxio.worker.ufs.instream.cache.max.size</td>
    <td>5000</td>
    <td>UFS输入流缓存中最大输入数。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.web.bind.host"></a> alluxio.worker.web.bind.host</td>
    <td>0.0.0.0</td>
    <td>Alluxio worker web服务绑定的主机名，参考See <a href="#configure-multihomed-networks">多宿主网络</a></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.web.hostname"></a> alluxio.worker.web.hostname</td>
    <td></td>
    <td>Alluxio worker web UI绑定的主机名。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.web.port"></a> alluxio.worker.web.port</td>
    <td>30000</td>
    <td>Alluxio worker web UI运行的端口号。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.worker.whitelist"></a> alluxio.worker.whitelist</td>
    <td>/</td>
    <td></td>
  </tr>

</tbody></table>

## 用户配置项

用户配置项指定了文件系统访问的相关信息。

<table class="table table-striped">
<tbody><tr><th>属性名</th><th>默认值</th><th>描述</th></tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.app.id"></a> alluxio.user.app.id</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.block.avoid.eviction.policy.reserved.size.bytes"></a> alluxio.user.block.avoid.eviction.policy.reserved.size.bytes</td>
    <td>0MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.block.master.client.pool.gc.interval"></a> alluxio.user.block.master.client.pool.gc.interval</td>
    <td>120sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.block.master.client.pool.gc.threshold"></a> alluxio.user.block.master.client.pool.gc.threshold</td>
    <td>120sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.block.master.client.pool.size.max"></a> alluxio.user.block.master.client.pool.size.max</td>
    <td>500</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.block.master.client.pool.size.min"></a> alluxio.user.block.master.client.pool.size.min</td>
    <td>0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.block.read.metrics.enabled"></a> alluxio.user.block.read.metrics.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.block.read.retry.max.duration"></a> alluxio.user.block.read.retry.max.duration</td>
    <td>5min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.block.read.retry.sleep.base"></a> alluxio.user.block.read.retry.sleep.base</td>
    <td>250ms</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.block.read.retry.sleep.max"></a> alluxio.user.block.read.retry.sleep.max</td>
    <td>2sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.block.size.bytes.default"></a> alluxio.user.block.size.bytes.default</td>
    <td>64MB</td>
    <td>Alluxio文件的默认大小。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.block.worker.client.pool.gc.threshold"></a> alluxio.user.block.worker.client.pool.gc.threshold</td>
    <td>300sec</td>
    <td>数据块worker client如果闲置超过这个时间会被关闭。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.block.worker.client.pool.max"></a> alluxio.user.block.worker.client.pool.max</td>
    <td>1024</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.block.write.location.policy.class"></a> alluxio.user.block.write.location.policy.class</td>
    <td>alluxio.client.block.policy.LocalFirstPolicy</td>
    <td>选择worker进行写文件数据块时的默认定位机制。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.async.restore.enabled"></a> alluxio.user.client.cache.async.restore.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.async.write.enabled"></a> alluxio.user.client.cache.async.write.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.async.write.threads"></a> alluxio.user.client.cache.async.write.threads</td>
    <td>16</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.dirs"></a> alluxio.user.client.cache.dirs</td>
    <td>/tmp/alluxio_cache</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.enabled"></a> alluxio.user.client.cache.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.eviction.retries"></a> alluxio.user.client.cache.eviction.retries</td>
    <td>10</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.evictor.class"></a> alluxio.user.client.cache.evictor.class</td>
    <td>alluxio.client.file.cache.evictor.LRUCacheEvictor</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.evictor.lfu.logbase"></a> alluxio.user.client.cache.evictor.lfu.logbase</td>
    <td>2.0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.evictor.nondeterministic.enabled"></a> alluxio.user.client.cache.evictor.nondeterministic.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.filter.class"></a> alluxio.user.client.cache.filter.class</td>
    <td>alluxio.client.file.cache.filter.DefaultCacheFilter</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.filter.config-file"></a> alluxio.user.client.cache.filter.config-file</td>
    <td>${alluxio.conf.dir}/cache_filter.properties</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.instream_buffer_size"></a> alluxio.user.client.cache.instream_buffer_size</td>
    <td>0B</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.local.store.file.buckets"></a> alluxio.user.client.cache.local.store.file.buckets</td>
    <td>1000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.page.size"></a> alluxio.user.client.cache.page.size</td>
    <td>1MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.quota.enabled"></a> alluxio.user.client.cache.quota.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.shadow.bloomfilter.num"></a> alluxio.user.client.cache.shadow.bloomfilter.num</td>
    <td>4</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.shadow.cuckoo.clock.bits"></a> alluxio.user.client.cache.shadow.cuckoo.clock.bits</td>
    <td>6</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.shadow.cuckoo.scope.bits"></a> alluxio.user.client.cache.shadow.cuckoo.scope.bits</td>
    <td>8</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.shadow.cuckoo.size.bits"></a> alluxio.user.client.cache.shadow.cuckoo.size.bits</td>
    <td>20</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.shadow.cuckoo.size.encoder.enabled"></a> alluxio.user.client.cache.shadow.cuckoo.size.encoder.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.shadow.cuckoo.size.prefix.bits"></a> alluxio.user.client.cache.shadow.cuckoo.size.prefix.bits</td>
    <td>8</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.shadow.cuckoo.size.suffix.bits"></a> alluxio.user.client.cache.shadow.cuckoo.size.suffix.bits</td>
    <td>12</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.shadow.enabled"></a> alluxio.user.client.cache.shadow.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.shadow.memory.overhead"></a> alluxio.user.client.cache.shadow.memory.overhead</td>
    <td>125MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.shadow.type"></a> alluxio.user.client.cache.shadow.type</td>
    <td>CLOCK_CUCKOO_FILTER</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.shadow.window"></a> alluxio.user.client.cache.shadow.window</td>
    <td>24h</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.size"></a> alluxio.user.client.cache.size</td>
    <td>512MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.store.overhead"></a> alluxio.user.client.cache.store.overhead</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.store.type"></a> alluxio.user.client.cache.store.type</td>
    <td>LOCAL</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.timeout.duration"></a> alluxio.user.client.cache.timeout.duration</td>
    <td>-1</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.timeout.threads"></a> alluxio.user.client.cache.timeout.threads</td>
    <td>32</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.ttl.check.interval.seconds"></a> alluxio.user.client.cache.ttl.check.interval.seconds</td>
    <td>3600</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.ttl.enabled"></a> alluxio.user.client.cache.ttl.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.cache.ttl.threshold.seconds"></a> alluxio.user.client.cache.ttl.threshold.seconds</td>
    <td>10800</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.client.report.version.enabled"></a> alluxio.user.client.report.version.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.conf.cluster.default.enabled"></a> alluxio.user.conf.cluster.default.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.conf.sync.interval"></a> alluxio.user.conf.sync.interval</td>
    <td>1min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.date.format.pattern"></a> alluxio.user.date.format.pattern</td>
    <td>MM-dd-yyyy HH:mm:ss:SSS</td>
    <td>以指定的日期格式，在Cli命令和Web页面中显示日期。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.buffer.bytes"></a> alluxio.user.file.buffer.bytes</td>
    <td>8MB</td>
    <td>在文件系统中进行读写操作时使用的缓冲区大小。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.copyfromlocal.block.location.policy.class"></a> alluxio.user.file.copyfromlocal.block.location.policy.class</td>
    <td>alluxio.client.block.policy.RoundRobinPolicy</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.create.ttl"></a> alluxio.user.file.create.ttl</td>
    <td>-1</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.create.ttl.action"></a> alluxio.user.file.create.ttl.action</td>
    <td>FREE</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.delete.unchecked"></a> alluxio.user.file.delete.unchecked</td>
    <td>false</td>
    <td>在尝试以递归方式删除持久化目录之前，检查底层文件系统中的内容是否与Alluxio同步。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.include.operation.id"></a> alluxio.user.file.include.operation.id</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.master.client.pool.gc.interval"></a> alluxio.user.file.master.client.pool.gc.interval</td>
    <td>120sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.master.client.pool.gc.threshold"></a> alluxio.user.file.master.client.pool.gc.threshold</td>
    <td>120sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.master.client.pool.size.max"></a> alluxio.user.file.master.client.pool.size.max</td>
    <td>500</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.master.client.pool.size.min"></a> alluxio.user.file.master.client.pool.size.min</td>
    <td>0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.metadata.load.type"></a> alluxio.user.file.metadata.load.type</td>
    <td>ONCE</td>
    <td>从UFS中加载元数据的行为。当访问关于路径的信息，但该路径在Alluxio中不存在时，元数据能够从UFS中加载。合法的选项有`Always`，`Never`，`Once`。`Always`将总是访问UFS来看路径是否存在于UFS中。`Never`表示从来不会访问UFS。`Once`表示在"首次"的时候会访问UFS（根据缓存），但是以后都不会在访问。默认值为`Once`。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.metadata.sync.interval"></a> alluxio.user.file.metadata.sync.interval</td>
    <td>-1</td>
    <td>在调用路径上的操作之前同步UFS元数据的时间间隔。-1表示不会发生同步。0意味着在操作之前，代理总是会同步路径的元数据。如果指定了一个时间间隔，就可以在该时间间隔内(尽可能)不重新同步路径。同步路径的元数据必须与UFS交互，所以这是一个昂贵的操作。如果对一个操作执行同步，则配置为"alluxio.user.file.metadata.load.type"将被忽略。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.passive.cache.enabled"></a> alluxio.user.file.passive.cache.enabled</td>
    <td>true</td>
    <td>当从Alluxio远程worker读文件时，是否缓存文件到Alluxio的本地worker。当从UFS读文件时，是否缓存到本地worker与这个选项无关。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.persist.on.rename"></a> alluxio.user.file.persist.on.rename</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.persistence.initial.wait.time"></a> alluxio.user.file.persistence.initial.wait.time</td>
    <td>0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.readtype.default"></a> alluxio.user.file.readtype.default</td>
    <td>CACHE</td>
    <td>创建Alluxio文件时的默认读类型。可选值为`CACHE_PROMOTE` (如果数据已经在Alluxio存储内，将其移动到最高存储层，如果数据需要从底层存储进行读取，将其写到本地Alluxio的最高存储层)、`CACHE` (如果数据需要从底层存储进行读取，将其写到本地Alluxio的最高存储层), `NO_CACHE` (数据不与Alluxio交互，如果是从Alluxio中进行读取，将不会发生数据块迁移或者剔除)。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.replication.durable"></a> alluxio.user.file.replication.durable</td>
    <td>1</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.replication.max"></a> alluxio.user.file.replication.max</td>
    <td>-1</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.replication.min"></a> alluxio.user.file.replication.min</td>
    <td>0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.reserved.bytes"></a> alluxio.user.file.reserved.bytes</td>
    <td>${alluxio.user.block.size.bytes.default}</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.sequential.pread.threshold"></a> alluxio.user.file.sequential.pread.threshold</td>
    <td>2MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.target.media"></a> alluxio.user.file.target.media</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.ufs.tier.enabled"></a> alluxio.user.file.ufs.tier.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.waitcompleted.poll"></a> alluxio.user.file.waitcompleted.poll</td>
    <td>1sec</td>
    <td>当使用waitCompleted机制时，查询文件完成状态的时间间隔。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.write.init.max.duration"></a> alluxio.user.file.write.init.max.duration</td>
    <td>2min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.write.init.sleep.max"></a> alluxio.user.file.write.init.sleep.max</td>
    <td>5sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.write.init.sleep.min"></a> alluxio.user.file.write.init.sleep.min</td>
    <td>1sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.write.tier.default"></a> alluxio.user.file.write.tier.default</td>
    <td>0</td>
    <td>数据块写入的默认存储层。可选值为整型数值。非负值代表从高层到底层的存储层（0代表第一层存储层，１代表第二层存储层，以此类推）。如果给定值大于存储层数量,这个数字代表最底层的存储层。负值代表从底层到高层的存储层（-1代表最底层存储层，-2代表次底层存储层，以此类推）如果给定值的绝对值大于存储层数量，这个数字代表最高层存储层。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.file.writetype.default"></a> alluxio.user.file.writetype.default</td>
    <td>ASYNC_THROUGH</td>
    <td>创建Alluxio文件时的默认写类型。可选值为`MUST_CACHE` (数据仅仅存储在Alluxio中，并且必须存储在其中), `CACHE_THROUGH` (尽量缓冲数据，同时同步写入到底层文件系统), `THROUGH` (不缓冲数据，同步写入到底层文件系统), `ASYNC_THROUGH` (数据写入Alluxio, 异步写入底层文件系统, 配置alluxio.user.file.replication.durable, 在数据持久化之前在Alluxio中复制alluxio.user.file.replication.durable个副本)。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.hdfs.client.exclude.mount.info.on.list.status"></a> alluxio.user.hdfs.client.exclude.mount.info.on.list.status</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.hostname"></a> alluxio.user.hostname</td>
    <td></td>
    <td>给alluxio客户端使用的主机名。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.local.reader.chunk.size.bytes"></a> alluxio.user.local.reader.chunk.size.bytes</td>
    <td>8MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.local.writer.chunk.size.bytes"></a> alluxio.user.local.writer.chunk.size.bytes</td>
    <td>64KB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.logging.threshold"></a> alluxio.user.logging.threshold</td>
    <td>10s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.master.polling.timeout"></a> alluxio.user.master.polling.timeout</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.metadata.cache.enabled"></a> alluxio.user.metadata.cache.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.metadata.cache.expiration.time"></a> alluxio.user.metadata.cache.expiration.time</td>
    <td>10min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.metadata.cache.max.size"></a> alluxio.user.metadata.cache.max.size</td>
    <td>100000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.metrics.collection.enabled"></a> alluxio.user.metrics.collection.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.metrics.heartbeat.interval"></a> alluxio.user.metrics.heartbeat.interval</td>
    <td>10sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.data.timeout"></a> alluxio.user.network.data.timeout</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.flowcontrol.window"></a> alluxio.user.network.flowcontrol.window</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.keepalive.time"></a> alluxio.user.network.keepalive.time</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.keepalive.timeout"></a> alluxio.user.network.keepalive.timeout</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.max.inbound.message.size"></a> alluxio.user.network.max.inbound.message.size</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.netty.channel"></a> alluxio.user.network.netty.channel</td>
    <td></td>
    <td>netty网络通道类型。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.netty.worker.threads"></a> alluxio.user.network.netty.worker.threads</td>
    <td></td>
    <td>远程数据块worker client从远程数据块worker读取数据使用的线程数目。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.reader.buffer.size.messages"></a> alluxio.user.network.reader.buffer.size.messages</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.reader.chunk.size.bytes"></a> alluxio.user.network.reader.chunk.size.bytes</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.rpc.flowcontrol.window"></a> alluxio.user.network.rpc.flowcontrol.window</td>
    <td>2MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.rpc.keepalive.time"></a> alluxio.user.network.rpc.keepalive.time</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.rpc.keepalive.timeout"></a> alluxio.user.network.rpc.keepalive.timeout</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.rpc.max.connections"></a> alluxio.user.network.rpc.max.connections</td>
    <td>1</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.rpc.max.inbound.message.size"></a> alluxio.user.network.rpc.max.inbound.message.size</td>
    <td>100MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.rpc.netty.channel"></a> alluxio.user.network.rpc.netty.channel</td>
    <td>EPOLL</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.rpc.netty.worker.threads"></a> alluxio.user.network.rpc.netty.worker.threads</td>
    <td>0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.streaming.flowcontrol.window"></a> alluxio.user.network.streaming.flowcontrol.window</td>
    <td>2MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.streaming.keepalive.time"></a> alluxio.user.network.streaming.keepalive.time</td>
    <td>9223372036854775807</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.streaming.keepalive.timeout"></a> alluxio.user.network.streaming.keepalive.timeout</td>
    <td>30sec</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.streaming.max.connections"></a> alluxio.user.network.streaming.max.connections</td>
    <td>64</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.streaming.max.inbound.message.size"></a> alluxio.user.network.streaming.max.inbound.message.size</td>
    <td>100MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.streaming.netty.channel"></a> alluxio.user.network.streaming.netty.channel</td>
    <td>EPOLL</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.streaming.netty.worker.threads"></a> alluxio.user.network.streaming.netty.worker.threads</td>
    <td>0</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.writer.buffer.size.messages"></a> alluxio.user.network.writer.buffer.size.messages</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.writer.chunk.size.bytes"></a> alluxio.user.network.writer.chunk.size.bytes</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.writer.close.timeout"></a> alluxio.user.network.writer.close.timeout</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.writer.flush.timeout"></a> alluxio.user.network.writer.flush.timeout</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.network.zerocopy.enabled"></a> alluxio.user.network.zerocopy.enabled</td>
    <td></td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.rpc.retry.base.sleep"></a> alluxio.user.rpc.retry.base.sleep</td>
    <td>50ms</td>
    <td>在遇到一些错误的时候，Alluxio客户端的RPC会基于指数级的延迟进行重试。这个配置决定了这个指数级重试的基数。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.rpc.retry.max.duration"></a> alluxio.user.rpc.retry.max.duration</td>
    <td>2min</td>
    <td>在遇到一些错误的时候，Alluxio客户端的RPC会基于指数级的延迟进行重试。这个配置决定了放弃前重试的最大时延。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.rpc.retry.max.sleep"></a> alluxio.user.rpc.retry.max.sleep</td>
    <td>3sec</td>
    <td>在遇到一些错误的时候，Alluxio客户端的RPC会基于指数级的延迟进行重试。这个配置决定了这个重试延迟的最大值。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.rpc.shuffle.masters.enabled"></a> alluxio.user.rpc.shuffle.masters.enabled</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.short.circuit.enabled"></a> alluxio.user.short.circuit.enabled</td>
    <td>true</td>
    <td>是否允许用户绕过Alluxio读取数据。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.short.circuit.preferred"></a> alluxio.user.short.circuit.preferred</td>
    <td>false</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.streaming.data.read.timeout"></a> alluxio.user.streaming.data.read.timeout</td>
    <td>3m</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.streaming.data.write.timeout"></a> alluxio.user.streaming.data.write.timeout</td>
    <td>3m</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.streaming.reader.buffer.size.messages"></a> alluxio.user.streaming.reader.buffer.size.messages</td>
    <td>16</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.streaming.reader.chunk.size.bytes"></a> alluxio.user.streaming.reader.chunk.size.bytes</td>
    <td>1MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.streaming.reader.close.timeout"></a> alluxio.user.streaming.reader.close.timeout</td>
    <td>5s</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.streaming.writer.buffer.size.messages"></a> alluxio.user.streaming.writer.buffer.size.messages</td>
    <td>16</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.streaming.writer.chunk.size.bytes"></a> alluxio.user.streaming.writer.chunk.size.bytes</td>
    <td>1MB</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.streaming.writer.close.timeout"></a> alluxio.user.streaming.writer.close.timeout</td>
    <td>30min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.streaming.writer.flush.timeout"></a> alluxio.user.streaming.writer.flush.timeout</td>
    <td>30min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.streaming.zerocopy.enabled"></a> alluxio.user.streaming.zerocopy.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.ufs.block.location.all.fallback.enabled"></a> alluxio.user.ufs.block.location.all.fallback.enabled</td>
    <td>true</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.ufs.block.read.concurrency.max"></a> alluxio.user.ufs.block.read.concurrency.max</td>
    <td>2147483647</td>
    <td>一个Block Worker上的一个UFS块并发访问的最大个数。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.ufs.block.read.location.policy"></a> alluxio.user.ufs.block.read.location.policy</td>
    <td>alluxio.client.block.policy.LocalFirstPolicy</td>
    <td>当Alluxio client从UFS读取文件时，它将读取委托给Alluxio worker。client使用此策略来选择要阅读哪个worker。 内置选择有：[<a href="https://www.alluxio.org/javadoc/master/alluxio/client/block/policy/DeterministicHashPolicy.html">alluxio.client.block.policy.DeterministicHashPolicy</a>, <a href="https://www.alluxio.org/javadoc/master/alluxio/client/file/policy/LocalFirstAvoidEvictionPolicy.html">alluxio.client.block.policy.LocalFirstAvoidEvictionPolicy</a>, <a href="https://www.alluxio.org/javadoc/master/alluxio/client/file/policy/LocalFirstPolicy.html">alluxio.client.block.policy.LocalFirstPolicy</a>, <a href="https://www.alluxio.org/javadoc/master/alluxio/client/file/policy/MostAvailableFirstPolicy.html">alluxio.client.block.policy.MostAvailableFirstPolicy</a>, <a href="https://www.alluxio.org/javadoc/master/alluxio/client/file/policy/RoundRobinPolicy.html">alluxio.client.block.policy.RoundRobinPolicy</a>, <a href="https://www.alluxio.org/javadoc/master/alluxio/client/file/policy/SpecificHostPolicy.html">alluxio.client.block.policy.SpecificHostPolicy</a>]</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.ufs.block.read.location.policy.cache.expiration.time"></a> alluxio.user.ufs.block.read.location.policy.cache.expiration.time</td>
    <td>10min</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.ufs.block.read.location.policy.cache.size"></a> alluxio.user.ufs.block.read.location.policy.cache.size</td>
    <td>10000</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.ufs.block.read.location.policy.deterministic.hash.shards"></a> alluxio.user.ufs.block.read.location.policy.deterministic.hash.shards</td>
    <td>1</td>
    <td>当alluxio.user.ufs.block.read.location.policy设为alluxio.client.block.policy.DeterministicHashPolicy，这设定了hash shards的数量。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.user.worker.list.refresh.interval"></a> alluxio.user.worker.list.refresh.interval</td>
    <td>2min</td>
    <td></td>
  </tr>

</tbody></table>

## 集群管理配置项

<table class="table table-striped">
<tbody><tr><th>属性名</th><th>默认值</th><th>描述</th></tr>

  <tr>
    <td><a class="anchor" name="alluxio.integration.master.resource.cpu"></a> alluxio.integration.master.resource.cpu</td>
    <td>1</td>
    <td>运行Alluxio master所需要的cpu核数。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.integration.master.resource.mem"></a> alluxio.integration.master.resource.mem</td>
    <td>1024MB</td>
    <td>运行Alluxio master所需要的内存大小。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.integration.worker.resource.cpu"></a> alluxio.integration.worker.resource.cpu</td>
    <td>1</td>
    <td>运行Alluxio worker所需要的cpu核数。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.integration.worker.resource.mem"></a> alluxio.integration.worker.resource.mem</td>
    <td>1024MB</td>
    <td>运行Alluxio worker所需要的内存大小，该部分内存不包含为层次化存储配置的内容。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.integration.yarn.workers.per.host.max"></a> alluxio.integration.yarn.workers.per.host.max</td>
    <td>1</td>
    <td></td>
  </tr>

</tbody></table>

## 安全性配置项

安全性配置项指定了安全性相关的信息，如安全认证和文件权限。
安全认证相关的配置同时适用于master、worker和用户。
文件权限相关的配置只对master起作用。
更多安全性相关的信息详见[系统安全文档](../security/Security.md)。

<table class="table table-striped">
<tbody><tr><th>属性名</th><th>默认值</th><th>描述</th></tr>

  <tr>
    <td><a class="anchor" name="alluxio.security.authentication.custom.provider.class"></a> alluxio.security.authentication.custom.provider.class</td>
    <td></td>
    <td>当alluxio.security.authentication.type被设为CUSTOM时，实现安全认证功能的类。 该类必须实现'alluxio.security.authentication.AuthenticationProvider'接口。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.security.authentication.type"></a> alluxio.security.authentication.type</td>
    <td>SIMPLE</td>
    <td>安全认证模式。目前支持三种模式：NOSASL、SIMPLE和CUSTOM。 默认为SIMPLE，即不启用安全认证功能。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.security.authorization.permission.enabled"></a> alluxio.security.authorization.permission.enabled</td>
    <td>true</td>
    <td>是否启用文件权限访问控制功能。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.security.authorization.permission.supergroup"></a> alluxio.security.authorization.permission.supergroup</td>
    <td>supergroup</td>
    <td>Alluxio文件系统的超级用户组。所有该组下的用户拥有超级权限。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.security.authorization.permission.umask"></a> alluxio.security.authorization.permission.umask</td>
    <td>022</td>
    <td>创建文件和目录时的文件权限掩码。初始化的创建权限为777，目录和文件之间相差111。 因此，当默认文件权限掩码为022时，新建目录的权限为755，新建文件的权限为644。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.security.group.mapping.cache.timeout"></a> alluxio.security.group.mapping.cache.timeout</td>
    <td>1min</td>
    <td>缓存的组映射过期时间。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.security.group.mapping.class"></a> alluxio.security.group.mapping.class</td>
    <td>alluxio.security.group.provider.ShellBasedUnixGroupsMapping</td>
    <td>提供"用户-用户组"映射服务的实现类。Master可以通过该服务得到指定用户所在组的组成员。 该类必须实现'alluxio.security.group.GroupMappingService'接口。 默认的实现类通过执行'groups'命令得到指定用户所在组的组成员。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.security.login.impersonation.username"></a> alluxio.security.login.impersonation.username</td>
    <td>_HDFS_USER_</td>
    <td></td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.security.login.username"></a> alluxio.security.login.username</td>
    <td></td>
    <td>当alluxio.security.authentication.type被设为SIMPLE或CUSTOM时，用户应用使用此配置项 作为连接Alluxio的用户名。若未设置该项，则使用系统登录用户名。</td>
  </tr>

  <tr>
    <td><a class="anchor" name="alluxio.security.stale.channel.purge.interval"></a> alluxio.security.stale.channel.purge.interval</td>
    <td>3day</td>
    <td></td>
  </tr>

</tbody></table>
