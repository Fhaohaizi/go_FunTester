package test

import (
	"fmt"
	"testing"
	"time"
)

func TestFun(t *testing.T) {
	ticker := time.NewTicker(10 * time.Second)
	fmt.Println(time.Now().Unix())
	<-ticker.C
	fmt.Println(time.Now().Unix())

}
