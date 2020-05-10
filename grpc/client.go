package main

import (
	"context"
	"flag"
	"sync"
	"sync/atomic"
	"time"

	"github.com/lubanproj/gorpc/log"
	"google.golang.org/grpc"

	pb "github.com/lubanproj/gorpc-benchmark/grpc/helloworld"
)

var concurrency = flag.Int64("concurrency", 500, "concurrency")
var total = flag.Int64("total", 1000000, "total requests")
func main() {
	flag.Parse()
	request(*total, *concurrency)
}


func request(totalReqs int64, concurrency int64) {

	perClientReqs := totalReqs / concurrency

	counter := &Counter{
		Total: perClientReqs * concurrency ,
		Concurrency: concurrency,
	}


	req := &pb.HelloRequest{Name: "hello"}


	var wg sync.WaitGroup
	wg.Add(int(concurrency))

	startTime := time.Now().UnixNano()

	for i:=int64(0); i<counter.Concurrency; i++ {

		go func(i int64) {
			// Set up a connection to the server.
			conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
			if err != nil {
				log.Info("did not connect: %v", err)
			}
			defer conn.Close()

			for j:=int64(0); j< perClientReqs; j++ {

				c := pb.NewGreeterClient(conn)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				rsp, err := c.SayHello(ctx, req)
				if err != nil {
					log.Info("could not greet: %v", err)
				}

				if err == nil && rsp.Message == "world" {
					atomic.AddInt64(&counter.Succ, 1)
				} else {
					log.Info("rsp fail : %v", err)
					atomic.AddInt64(&counter.Fail, 1)
				}
			}

			wg.Done()
		}(i)
	}

	wg.Wait()

	counter.Cost = (time.Now().UnixNano() - startTime) / 1000000

	log.Info("took %d ms for %d requests", counter.Cost, counter.Total)
	log.Info("sent     requests      : %d\n", counter.Total)
	log.Info("received requests      : %d\n", atomic.LoadInt64(&counter.Succ) + atomic.LoadInt64(&counter.Fail))
	log.Info("received requests succ : %d\n", atomic.LoadInt64(&counter.Succ))
	log.Info("received requests fail : %d\n", atomic.LoadInt64(&counter.Fail))
	log.Info("throughput  (TPS)      : %d\n", totalReqs*1000/counter.Cost)

}

type Counter struct {
	Succ int64  // 成功量
	Fail int64  // 失败量
	Total int64 // 总量
	Concurrency int64 // 并发量
	Cost int64  // 总耗时 ms
}
