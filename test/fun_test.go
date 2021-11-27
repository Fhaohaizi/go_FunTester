package test

import (
	"fmt"
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
