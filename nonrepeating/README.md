1、go test -coverprofile=c.out　　生成覆盖率数据
2、go tool cover -html c.out　　生成html页面展示

#单元测试代码覆盖率
go test -coverprofile=c.out

#生成 报告文件（报错，执行失败）
go tool cover -html=c.out




go test -bench .

go test -bench . -cpuprofile cpu.out

go tool pprof cpu.out

web