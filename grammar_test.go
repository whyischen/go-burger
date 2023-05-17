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

// TestChannel
/*
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
	slice := make([]int, 3, 3)
	slice[0] = 0
	slice[1] = 1
	// slice 会根据 len 读取底层数组
	log.Printf("len test, slice: %+v\n---\n", slice)
	slice[2] = 2

	log.Printf("before ap1, slice: %+v", slice)
	// before ap1, slice: [0 1 2]
	// 外层函数的 len 不变，ap1 函数内增加的 111 对外不可见，打印的还是 [0 1 2]
	ap1(slice)
	log.Printf("after ap1, slice: %+v", slice)
	// 避免这种现象，可以使用指针类型传递
	// 主函数读取到了 ap1Point 内因为扩容改变了的最新地址
	log.Printf("before ap1Point, addr:%p, slice: %+v", slice, &slice)
	ap1Point(&slice)
	log.Printf("after  ap1Point, addr:%p, slice: %+v", slice, &slice)

}
func ap1(arr []int) {
	arr = append(arr, 111)
}
func ap1Point(arr *[]int) {
	*arr = append(*arr, 000, 111, 222)
}

func TestMap(t *testing.T) {
	mp := make(map[string]int, 0)
	mp["0"] = 0
	mp["1"] = 1
	mp["2"] = 2
	log.Printf("mp: %v", mp)

	/*
		测试用 new 创建 map
		语法检测通不过，会报错
		mp1 := new(map[string]int)
		mp1["0"] = 0
	*/
}
