package lock

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg4     sync.WaitGroup
)

func main555() {
	wg4.Add(2)
	go incCounter(1)
	go incCounter(2)

	wg4.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg4.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		//当前goroutine从线程退出，并放回到队列
		runtime.Gosched()

	}
}
