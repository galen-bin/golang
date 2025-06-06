package task02

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var value int32

func task02() {
	//指针
	//var a int = 2
	//var b = []int{1, 2, 3, 4}
	//zl(&a)
	//ps(&b)
	//Goroutine
	/*go gort(1)
	go gort(2)
	time.Sleep(time.Second)*/
	/*	for i := 1; i < 10; i++ {
			go gort_sleep(i)
			go gort_sleep(i)
			go gort_sleep(i)
		}
		time.Sleep(time.Second * 10)*/
	//面向对象
	/*re := Rectangle{width: 10, height: 5}
	s1 := re.Area()
	ci := Circle{radius: 50}
	s2 := ci.Perimeter()
	pr := Person{Name: "jack", Age: 23}
	em := Employee{EmployeeID: 10, per: pr}
	em.PrintInfo()
	fmt.Println(s1, s2)*/
	//Channel 通道
	/*chan1 := make(chan int)
	go chantest(chan1)
	go changet(chan1)*/
	/*chan2 := make(chan int, 100)
	go chant_100(chan2, 100)
	go changet(chan2)
	time.Sleep(time.Second * 2)*/
	//锁机制
	/*for i := 0; i < 10; i++ {
		go sync_lock(i + 1)
	}
	time.Sleep(time.Second * 2)*/

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go sync_atom()
	}

	wg.Wait() // 当所有goroutine调用Done()后，Wait()会返回
	fmt.Println("所有goroutine已完成")
	fmt.Println(atomic.LoadInt32(&value))
	time.Sleep(time.Second * 2)

}

func sync_atom() {

	//var numb int
	for i := 0; i < 100; i++ {
		//numb++
		atomic.AddInt32(&value, int32(i))
	}

	wg.Done()
}

var sy sync.Mutex

func sync_lock(cc int) {
	sy.Lock()
	var numb int
	for i := 0; i < 1000; i++ {
		numb++
	}
	sy.Unlock()
	fmt.Println(cc, "轮完成", numb)
}

func chant_100(chan1 chan int, f int) {
	for i := 0; i < f; i++ {
		chan1 <- i
	}
	close(chan1)
}

func chantest(chan1 chan int) {
	for i := 0; i < 10; i++ {
		chan1 <- i
	}
	close(chan1)
}

func changet(chan1 chan int) {

	for v := range chan1 {
		fmt.Println("取出", v)
	}

}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID int
	per        Person
}

type PrintInfo interface{}

func (em Employee) PrintInfo() {
	fmt.Println("年龄", em.per.Age)
	fmt.Println("Id", em.EmployeeID)
	fmt.Println("姓名", em.per.Name)
}

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	width  float32
	height float32
}

type Circle struct {
	radius float32
}

func (r Rectangle) Area() float32 {
	return r.width * r.height
}

func (r Circle) Perimeter() float32 {
	return r.radius * 3.14 * 2
}

func gort_sleep(s int) {
	start := time.Now()
	for i := 1; i < s; i++ {
		time.Sleep(time.Second)
	}

	duration := time.Now().Sub(start) // 计算经过的时间差
	fmt.Println("执行时长", duration)

}

func gort(parse int) {
	for i := 1; i <= 10; i++ {

		if parse == 1 {
			if i%2 == 1 {
				fmt.Println("奇数", i)
			}
		} else {
			if i%2 == 0 {
				fmt.Println("偶数", i)
			}
		}

	}
}

func zl(c *int) {

	fmt.Println(*c + 10)
}

func ps(c *[]int) {
	var s []int
	for _, v := range *c {
		s = append(s, v*2)
	}

	fmt.Println(s)
}
