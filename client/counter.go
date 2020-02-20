package main

type Counter struct {
	Succ int64  // 成功量
	Fail int64  // 失败量
	Total int64 // 总量
	Concurrency int64 // 并发量
	Cost int64  // 总耗时 ms
}
