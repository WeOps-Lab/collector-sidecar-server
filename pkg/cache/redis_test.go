package cache

import (
	"testing"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func TestDefaultRedisCache(t *testing.T) {
	db, mock := redismock.NewClientMock()
	cache := &defaultRedisCache{client: db}

	// Test Set method
	mock.ExpectSet("key", "value", 0).SetVal("OK")
	err := cache.Set("key", "value")
	assert.NoError(t, err)

	// Test Get method
	mock.ExpectGet("key").SetVal("value")
	val, err := cache.Get("key")
	assert.NoError(t, err)
	assert.Equal(t, "value", val)

}
