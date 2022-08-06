package test

import (
	"fmt"
	"funtester/base"
	"funtester/ftool"
	"github.com/tealeg/xlsx"
	"log"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestFun(t *testing.T) {
	tt := "2006年01月02日 15时04分05秒"
	ss := strings.Replace(tt, "0", "1", -1)
	ss1 := strings.ReplaceAll(tt, "0", "1")
	fmt.Println(ss)
	fmt.Println(ss1)
	compile := regexp.MustCompile(`\d+`)
	submatch := compile.FindAllStringSubmatch(tt, -1)
	fmt.Println(submatch)
	find := compile.Find([]byte(tt))
	matchString := compile.MatchString(tt)
	fmt.Println(string(find))
	fmt.Println(matchString)
}
func TestExce1l(t *testing.T) {
	str := "0001a111222a22a"
	//fmt.Println(ftool.Match(str,"\\d+"))
	//fmt.Println(ftool.Find(str,"\\d+"))
	fmt.Println(ftool.FindAll(str, "\\d+"))
}

func TestExcel(t *testing.T) {

	output, err := xlsx.FileToSlice("/Users/oker/Desktop/a.xlsx")
	if err != nil {
		panic(err.Error())
	}
	log.Println(output[0][1][1])
	for rowIndex, row := range output[0] {
		for cellIndex, cell := range row {
			log.Println(fmt.Sprintf("第%d行，第%d个单元格：%s", rowIndex+1, cellIndex+1, cell))
		}
	}

}

func TestChannel(t *testing.T) {
	c := make(chan int)
	go func() {
		//for {
		ftool.Sleep(100)
		c <- ftool.RandomInt(100)
		close(c)
		//}
	}()
	for {

		num, ok := <-c
		log.Println(num)
		log.Println(ok)
		if !ok {
			break
		}
	}
}

func TestPrint(t *testing.T) {
	go func() {
		ftool.PrintTime(func() {
			ftool.Sleep(500)
		}, 1, base.FunTester)
	}()
}

func TestOnce(t *testing.T) {
	var group sync.WaitGroup
	var once sync.Once
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			log.Printf("异步执行%d 次", i)
			once.Do(func() {
				log.Println("执行一次")
			})
			group.Done()
		}()
	}
	group.Wait()
}

// TestOnceSimple once对象简单测试
//  @Description:
//  @param t
//
func TestOnceSimple(t *testing.T) {
	var once sync.Once
	for i := 0; i < 10; i++ {
		go once.Do(func() {
			log.Println("执行一次")
		})
	}
	time.Sleep(time.Second)
}
