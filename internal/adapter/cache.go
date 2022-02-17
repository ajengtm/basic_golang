package adapter

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
)

var myCache CacheItf

type AppCache struct {
	client *cache.Cache
}

type CacheItf interface {
	Set(key string, data interface{}, expiration time.Duration) error
	Get(key string) ([]byte, error)
}

func InitCache() CacheItf {
	myCache = &AppCache{
		client: cache.New(5*time.Minute, 10*time.Minute),
	}
	return myCache
}

func (r *AppCache) Get(key string) ([]byte, error) {
	res, exist := r.client.Get(key)
	if !exist {
		return nil, nil
	}

	resByte, ok := res.([]byte)
	if !ok {
		return nil, errors.New("Format is not arr of bytes")
	}

	return resByte, nil
}

func (r *AppCache) Set(key string, data interface{}, expiration time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	r.client.Set(key, b, expiration)
	return nil
}
