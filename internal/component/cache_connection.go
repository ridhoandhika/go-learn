package component

import (
	"context"
	"log"
	"ridhoandhika/backend-api/domain"
	"time"

	"github.com/allegro/bigcache/v3"
)

func GetCacheConnection() domain.CacheRepository {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(time.Hour*24))
	if err != nil {
		log.Fatalf("error connect cache %s", err.Error())
	}
	return cache
}
