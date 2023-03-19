package main

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

/***
Golang提供了指针用于操作数据内存，并通过引用来修改变量。
只声明未赋值的变量，golang都会自动为其初始化为零值，
基础数据类型的零值比较简单，引用类型和指针的零值都为nil，
nil类型不能直接赋值，因此需要通过new开辟一个内存，
或者通过make初始化数据类型，或者两者配合，然后才能赋值
*/

// WaitGroup
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	groupNum := 10
	wg.Add(groupNum)
	for i := 0; i < groupNum; i++ {
		go func(num int, wg *sync.WaitGroup) {
			fmt.Printf("print num: %v", num)
		}(i, &wg)
	}
	wg.Wait()
}

/***
// 只读 channel
var readOnlyChan <-chan int // channel 的类型为 int

// 只写 channel
var writeOnlyChan chan<- int

// 可读可写
var ch chan int

// 或者使用 make 直接初始化
readOnlyChan1 := make(<-chan int, 2) // 只读且带缓存区的 channel
readOnlyChan2 := make(<-chan int)    // 只读且不带缓存区 channel

writeOnlyChan3 := make(chan<- int, 4) // 只写且带缓存区 channel
writeOnlyChan4 := make(chan<- int)    // 只写且不带缓存区 channel

ch := make(chan int, 10) // 可读可写且带缓存区

ch <- 20      // 写数据
i := <-ch     // 读数据
i, ok := <-ch // 还可以判断读取的数据
*/
func TestChannel(t *testing.T) {

	ch := make(chan int)
	// 没有缓存的 channel，不先取值的话，线程会阻塞在向 channel 存值的地方
	go func(ch chan int) {
		val := <-ch
		log.Printf("val := <- ch, %v\n", val)
	}(ch)
	ch <- 0
	log.Printf("ch <- 0 ...\n")
	time.Sleep(time.Second * 3)

	log.Printf("done ...")
}

func TestSlice(t *testing.T) {
	slice := make([]int, 10)
	log.Printf("1. slice: %v", slice)
	slice[1] = 1
	log.Printf("2. slice: %v", slice)
}
