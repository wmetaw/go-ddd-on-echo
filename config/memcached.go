package config

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var (
	MemcacheCon *memcache.Client
)

func NewMemcacheConnection() *memcache.Client {
	return memcache.New(getEnv(MEMCACHE_HOST, "memcache-server") + ":" + getEnv(MEMCACHE_PORT, "11211"))
}
