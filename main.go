package main

import (
	"backend-config.Cache/cache"
	"backend-config.Cache/config"
	"backend-config.Cache/router"
)

var err error

func NewLRUCache(c int) *cache.LRUCache {
	return &cache.LRUCache{
		Cap:      c,
		CacheMap: make(map[string]cache.Cache, c),
	}
}

func main() {

	config.Lru = NewLRUCache(1024) // default size

	//Initalising router
	router.InitRoutes()

}
