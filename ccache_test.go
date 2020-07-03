package cache2

import (
	"fmt"
	"testing"
	"time"
)

var cache *CCache = NewCCache(100)

func Test_ccache(t *testing.T) {

	f := func() (interface{}, error) {
		return "xxxx", nil
	}

	res, err := cache.Use("key", f, time.Second*1)
	fmt.Println(res, err)
	time.Sleep(time.Second * 2)

	f2 := func() (interface{}, error) {
		return "bbbbb", nil
	}
	res2, err := cache.Use("key", f2, time.Second*10)
	fmt.Println(res2, err)

}
