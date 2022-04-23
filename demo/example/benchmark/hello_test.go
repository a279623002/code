package benchmark

import (
	"bytes"
	"html/template"
	"testing"
)

//type BenchmarkResult struct {
//	N         int           // 迭代次数
//	T         time.Duration // 基准测试花费的时间
//	Bytes     int64         // 一次迭代处理的字节数
//	MemAllocs uint64        // 总的分配内存的次数
//	MemBytes  uint64        // 总的分配内存的字节数
//}

//go test -benchmem -bench .
func BenchmarkHello(b *testing.B) {
	//... // 耗时操作
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		hello()
	}
}

//使用 RunParallel 测试并发性能
func BenchmarkParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			// 所有 goroutine 一起，循环一共执行 b.N 次
			buf.Reset()
			templ.Execute(&buf, "world")
		}
	})
}