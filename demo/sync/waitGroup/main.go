package main

import (
	"fmt"
	"sync"
)

func main() {
	// 声明等待组
	var wg sync.WaitGroup
	// 设置，需要等待3个协程执行完成
	wg.Add(3)
	// 创建通道
	intChan := make(chan int)
	// 计算1-50的和，并引入对应的waitGroup
	go func(intChan chan int, wg *sync.WaitGroup) {
		sum := 0
		for i := 1; i <= 50; i++ {
			sum += i
		}
		intChan <- sum
		// 计数器减一
		wg.Done()
	}(intChan, &wg)
	// 计算51-100的和，并引入对应的waitGroup
	go func(intChan chan int, wg *sync.WaitGroup) {
		sum := 0
		for i := 51; i <= 100; i++ {
			sum += i
		}
		intChan <- sum
		// 计数器减一
		wg.Done()
	}(intChan, &wg)
	// 另外创建个channle聚合结果
	go func(intChan chan int, wg *sync.WaitGroup) {
		sum1 := <-intChan
		sum2 := <-intChan
		fmt.Printf("sum1 = %d sum2 = %d  \nsum1 + sum2 = %d \n", sum1, sum2, sum1+sum2)
		// 计数器减一
		wg.Done()
	}(intChan, &wg)
	// 阻塞，直到等待组的计数器等于0，主线程才能够结束，否则随着主线程的提前结束，各个goroutine也会被关闭
	wg.Wait()
	fmt.Println("运行结束")
}

/**输出
  sum1 = 3775 sum2 = 1275
  sum1 + sum2 = 5050
  运行结束
*/
