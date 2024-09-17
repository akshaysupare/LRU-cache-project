package config

import "backend-config.Cache/cache"

const (
	Port = "9000"
)

var (
	Lru *cache.LRUCache
	Err error
)
