package cache

import (
	"errors"
	"sync"
	"time"

	"github.com/urfave/cli/v2"
)

// cache find no result
var (
	ErrNoResult       = errors.New("cache return no result")
	ErrObjectNotFound = errors.New("cache object not found")

	ExpireTime = 24 * time.Hour // 过期时间24小时
)

// CacheObjector save and load with all structure
type CacheObjector interface {
	GetStoreIndex() int64
}

type Cache interface {
	SaveObject(prefix string, k interface{}, x interface{}) error
	SaveFields(prefix string, k interface{}, fields map[string]interface{}) error
	LoadObject(prefix string, value interface{}, x interface{}) error
	LoadArray(prefix string, ownerId int64, pool *sync.Pool) ([]interface{}, error)
	DeleteObject(prefix string, k interface{}) error
	DeleteFields(prefix string, k interface{}, fieldsName []string) error
	Exit() error
}

func NewCache(ctx *cli.Context) Cache {
	return NewGoRedis(ctx)
}
