package config

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var (
	MCCon *memcache.Client
)

func NewMemcacheConnection() *memcache.Client {
	return memcache.New("memcached-server:11211")
}
