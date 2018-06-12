package proxy

import (
	"testing"
	"fmt"
)

func TestGetUsefulProxy(t *testing.T) {
	xx,err := GetUsefulProxy(100)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	for _,v := range xx {
		fmt.Println(v)
	}
}
