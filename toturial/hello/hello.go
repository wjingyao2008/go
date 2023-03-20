package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
)

// https://go.dev/tour/list
func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func pointerMain() {
	fmt.Println("pointerMain start")
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	//& is generate pointer
	//* is denote pointer's underlying value

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
	fmt.Println("pointerMain end")
}

type Vertex struct {
	X int
	Y int
}

func structMain() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)
	p := &v
	p.X = 1e9 //no need to dereference, automatically done
	fmt.Println(v.X)
}

func arrayMain() {
	var a [2]string
	a[0] = "a"
	a[1] = "b"
	fmt.Println(a)

	arr1 := [6]int{0, 1, 2, 3, 4}
	fmt.Println(arr1)

}

func slicesMain() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	slice := primes[2:5]
	fmt.Printf("value:%v,type:%T\n", primes, primes)
	fmt.Printf("value:%v,len: %v type:%T\n", slice, len(slice), slice)
}
func sliceExtendMain() {
	fmt.Println("slice extend:")

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)
	b := []int{2, 3, 5, 7, 11, 13}
	b1 := b[:4]
	printSlice(b1)
	b2 := b[4:]
	printSlice(b2)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func rangMain(array []int) {
	for i, v := range array {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for v := range array {
		fmt.Printf("%d\n", v)
	}
}

func funcCompute(fn func(int, int) int) int {
	return fn(3, 4)
}

func printFuncNameStartEnd(fn func()) {
	name := GetFunctionName(fn)
	fmt.Printf("start execute [%v] =======\n", name)
	fn()
	fmt.Printf("execute [%v] done =======\n", name)
	fmt.Println("")

}
func GetFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func (v Vertex) abs() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}
func VertexMathMain() {
	v := Vertex{3, 4}
	fmt.Println(v.abs())

	fmt.Println(*(v.scale(10)))

}

// Pointer receivers
func (v *Vertex) scale(s int) *Vertex {
	v.X *= s
	v.Y *= s
	return v
}

// if a function take a pointer argument,it must take pointer,not value
// but if function is a pointer receiver,this function can apply to a pointer or a value.
// var v Vertex
// v.Scale(5)  // OK
// p := &v
// p.Scale(10) // OK

// Goroutines
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// channels
func sum(arr []int, c chan int) {
	sum := 0
	fmt.Printf("i got: %\n", arr)
	for v := range arr {
		sum += v
	}
	c <- sum
}
func goRoutineMain() {
	go say("world")
	say("hello")
}

// channel can be buffer
func goChannelMain() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

// channel can be close, to end the loop
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// lock,no need channel
type LockCounter struct {
	mus     sync.Mutex
	counter int
}

func (lock *LockCounter) increase() {
	lock.mus.Lock()
	lock.counter++
	lock.mus.Unlock()
}

func mutexMain() {
	mylock := LockCounter{}
	for i := 0; i < 1000; i++ {
		go mylock.increase()
	}
	time.Sleep(time.Second)
	fmt.Println(mylock.counter)
}

func main() {
	printFuncNameStartEnd(pointerMain)
	printFuncNameStartEnd(structMain)
	printFuncNameStartEnd(arrayMain)
	printFuncNameStartEnd(slicesMain)
	printFuncNameStartEnd(sliceExtendMain)
	printFuncNameStartEnd(VertexMathMain)
	printFuncNameStartEnd(goRoutineMain)
	printFuncNameStartEnd(goChannelMain)
	printFuncNameStartEnd(mutexMain)
}
