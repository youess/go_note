/*
Golang的Timer在源码中，实现的方式是以一个小顶堆来维护所有的Timer集合。接着启动一个独立的goroutine，
循环从小顶堆中的检测最近一个到期的Timer的到期时间，接着它睡眠到最近一个定时器到期的时间。
最后会执行开始时设定的回调函数。Timer到期之后，会被Golang的runtime从小项堆中删除，并等待GC回收资源。
*/
package main

import (
	"fmt"
	"time"
)

// 这个例子中相当于用time.Sleep
func testTimer1() {
	timer := time.NewTimer(3 * time.Second)

	fmt.Printf("还不会阻塞的操作 - %s\n", time.Now())

	expire := <-timer.C
	fmt.Printf("timer 1 expired at - %s\n", expire)
}

// 这个可以用来设定一个时间，不管后来程序有没有执行完成，都能结束它
func testTimer2() {
	timer := time.NewTimer(5 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("timer 2 expired") // not ran
	}()

	stop := timer.Stop()
	if stop {
		fmt.Println("timer 2 stopped")
	}

}

func testTimer3() {
	timer := time.NewTimer(3 * time.Second)
	go func() {
		<-timer.C                      // 阻塞全部进程
		fmt.Println("timer 3 expired") // will print out after timer expired
	}()

	fmt.Printf("echo first\n")
	time.Sleep(5 * time.Second) // 如果注释掉，不会答应协程中的打印行
	fmt.Printf("echo last\n")
}

func testTimer4() {
	timer := time.NewTimer(3 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("timer 4 expired") // will not print out
	}()

	fmt.Printf("echo first\n")
	timer.Stop() // 提前结束
	time.Sleep(5 * time.Second)
	fmt.Printf("echo last\n")
}

func testTimer5() {
	timer := time.NewTimer(3 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("timer 5 expired") // will print out immediatantly
	}()

	fmt.Printf("echo first\n")
	timer.Reset(0 * time.Second) // 重置0秒，马上过期
	time.Sleep(5 * time.Second)
	fmt.Printf("echo last\n")
}

func main() {
	// testTimer1()
	// testTimer2()
	testTimer3()
	// testTimer4()
	// testTimer5()
}
