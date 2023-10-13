package lock

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter5 int
	wg5      sync.WaitGroup
	mutex    sync.Mutex
)

func main() {
	wg5.Add(2)
	go incCounter5(1)
	go incCounter5(2)
	wg5.Wait()
	fmt.Printf("Final Count:5s", counter)
}

// incCounter使用互斥锁来同步并保证安全访问，
// 增加包里counter变量的值
func incCounter5(id int) {
	// 在函数退出时调用Done来通知main函数工作已经完成
	defer wg5.Done()
	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个goroutine进入
		// 这个临界区
		mutex.Lock()
		{
			// 捕获counter的值
			value := counter
			// 当前goroutine从线程退出，并放回到队列
			runtime.Gosched()
			// 增加本地value变量的值
			value++
			// 将该值保存回counter
			counter = value
		}
		mutex.Unlock()
		// 释放锁，允许其他正在等待的goroutine
		// 进入临界区
	}
}
