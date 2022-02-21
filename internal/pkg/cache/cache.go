package cache

import (
	"errors"
	"time"
)

type object struct {
	Value []byte
	ExpAt int64
}

var cache map[string]*object

func Init() {
	cache = make(map[string]*object)
}

func Set(key string, value []byte, expAt int64) {
	cache[key] = &object{
		Value: value,
		ExpAt: expAt,
	}
}

func Get(key string) (result []byte, err error) {
	if cache[key] == nil || cache[key].ExpAt < time.Now().Unix() {
		err = errors.New("empty cache")
	} else if cache[key].ExpAt < time.Now().Unix() {
		cache[key] = nil
		err = errors.New("reset cache")
	} else {
		result = cache[key].Value
	}
	return
}
