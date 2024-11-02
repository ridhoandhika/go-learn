package component

import (
	"context"
	"log"
	"santrikoding/backend-api/domain"
	"time"

	"github.com/allegro/bigcache/v3"
)

func GetCacheConnection() domain.CacheRepository {
	// bigcache := bigcache.BigCache{}
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		log.Fatalf("error connect cache %s", err.Error())
	}
	return cache
}
