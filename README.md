## gorpc-benchmark
a benchmark tool for gorpc


## Quick Start

to test gorpc 

**gorpc** :
```
git clone https://github.com/lubanproj/gorpc-benchmark.git
cd gorpc-benchmark
# start gorpc server
go run server.go
# start gorpc-benchmark client，start another terminal and execute
go run client.go -concurrency=100 -total=1000000
```
The performance test results are as follows : 
```
> go run client.go -concurrency=100 -total=1000000
2020/02/29 15:56:57 client.go:71: [INFO] took 5214 ms for 1000000 requests
2020/02/29 15:56:57 client.go:72: [INFO] sent     requests      : 1000000
2020/02/29 15:56:57 client.go:73: [INFO] received requests      : 1000000
2020/02/29 15:56:57 client.go:74: [INFO] received requests succ : 1000000
2020/02/29 15:56:57 client.go:75: [INFO] received requests fail : 0
2020/02/29 15:56:57 client.go:76: [INFO] throughput  (TPS)      : 191791
```

to test other framework, such as grpc

**grpc** : 
Test grpc with the same machine :
```
git clone https://github.com/lubanproj/gorpc-benchmark.git
cd gorpc-benchmark/grpc
# 运行 gorpc server
go run server.go
# start gorpc-benchmark client, start another terminal and execute 
go run client.go -concurrency=100 -total=1000000
```
The performance test results are as follows : 
```
> go run client.go -concurrency=100 -total=1000000
2020/02/29 15:46:14 client.go:77: [INFO] took 17169 ms for 1000000 requests
2020/02/29 15:46:14 client.go:78: [INFO] sent     requests      : 1000000
2020/02/29 15:46:14 client.go:79: [INFO] received requests      : 1000000
2020/02/29 15:46:14 client.go:80: [INFO] received requests succ : 1000000
2020/02/29 15:46:14 client.go:81: [INFO] received requests fail : 0
2020/02/29 15:46:14 client.go:82: [INFO] throughput  (TPS)      : 58244
```
