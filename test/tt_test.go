package test

import (
	"fmt"
	"funtester/ft"
	"sync"
	"testing"
	"time"
)

func TestStageJSON(t *testing.T) {
	fmt.Println("FunTester")
	time.Sleep(1 * time.Second)
	//t.Fail()
}

/*

#include



void sayHi() {

    printf("Hi");

}

*/
func TestCs(t *testing.T) {
	//C.sayHi()
	ft.Decorator(ft.Hello)("tester")

	intss := []int{1, 2, 3, 4}
	echo := ft.Echo(intss)
	sq := ft.Sq(echo)
	n, ok := <-sq
	if !ok {
		close(sq)
	}
	fmt.Println(n)

	var wait sync.WaitGroup
	wait.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().Format("2006年01月02日 15时 04分 05秒"))
			wait.Done()
		}()
	}
	wait.Add(1)
	time.Sleep(1 * time.Second)
	wait.Done()
	wait.Wait()
	fmt.Println("FunTester")
}
