package cache2

//  内存lru cache

import (
	"errors"
	gccache "github.com/karlseguin/ccache"
	"time"
)

const MAXSIZE = 1000000
const DEFAULTSIZE = 1000

type CCache struct {
	c *gccache.Cache
}

func NewCCache(maxSize int) *CCache {
	if maxSize > MAXSIZE {
		maxSize = MAXSIZE
	}
	if maxSize <= 0 {
		maxSize = DEFAULTSIZE
	}

	cache := new(CCache)
	count := uint32(maxSize/10 + 1)
	cache.c = gccache.New(gccache.Configure().MaxSize(int64(maxSize)).ItemsToPrune(count))
	return cache
}
func (c *CCache) Del(key string) {
	c.c.Delete(key)
}

//用key去找，找到返回，找不到则运行f，成功后把结果放入cache，再返回
func (c *CCache) Use(key string, f func() (interface{}, error), duration time.Duration) (interface{}, error) {

	if key == "" {
		return nil, errors.New("no key")
	}

	item := c.c.Get(key)
	if item != nil {
		if item.TTL().Seconds() > 0 {
			return item.Value(), nil
		}
	}
	//运行f
	res, err := f()
	if err != nil {
		return res, err
	}
	c.c.Set(key, res, duration)
	return res, nil
}
