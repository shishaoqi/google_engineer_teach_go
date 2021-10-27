# PProf 笔记

https://mp.weixin.qq.com/s?__biz=MzUxMDI4MDc1NA==&amp;mid=2247488702&amp;idx=1&amp;sn=b941ddb5473e8f6b85cd970e81225347&amp;chksm=f90401e3ce7388f50f390eb4dfd887481a7866cb50011802d1916ec644c3ba5485ea0e423036&amp;scene=21#wechat_redirect

### 查看可视化界面

```
wget http://localhost:6060/debug/pprof/profile
```
等待 30 秒，执行完毕后可在当前目录下发现采集的文件 profile，针对可视化界面我们有两种方式可进行下一步分析：

1. 方法一（推荐）：该命令将在所指定的端口号运行一个 PProf 的分析用的站点。
```
go tool pprof -http=:6001 profile
```

2. 方法二：通过 web 命令将以 svg 的文件格式写入图形，然后在 Web 浏览器中将其打开。
```
go tool pprof profile
```