package cache

import "errors"

var ErrKeyNotFound = errors.New("key not exists")

//Storage interface for cache.
type Storage interface {
	//Put insert value into cache.
	Put(key string, value string) error
	//Get get value from cache by key, if value for key not exists ErrKeyNotFound will be returned.
	Get(key string) (string, error)
	// Delete delete key and associated value from cache.
	Delete(key string)
}
