package test

import (
	"fmt"
	"funtester/base"
	"funtester/ftool"
	"github.com/tealeg/xlsx"
	"log"
	"regexp"
	"strings"
	"testing"
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

func TestPrint(t *testing.T) {
	go func() {
		ftool.PrintTime(func() {
			ftool.Sleep(500)
		}, 1, base.FunTester)
	}()
}
