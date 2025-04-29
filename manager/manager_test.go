package manager

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

type Manager1 struct {
	batchId uint64
}

func Test1(t *testing.T) {
	var m Manager1

	var wg sync.WaitGroup

	// 启动多个 goroutines 来同时增加 batchId
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint64(&m.batchId, 1)
		}()
	}

	wg.Wait()

	// 打印最终的 batchId
	fmt.Println("Final batchId:", m.batchId)
}
