package cache

import (
	"gitlab.com/convee/system/redis"
)

const (
	//TokenExpire token有效期
	TokenExpire = 7 * 24 * 60 * 60
)

// Cache cache结构体
type Cache struct {
	redis *redis.Redis
}

//NewCache cache句柄
func NewCache() *Cache {
	name := "chess"
	return &Cache{redis: redis.New(name)}
}

//SetTokenInfo 设置token
func (c *Cache) SetTokenInfo(token string, tokenInfo string) (bool, error) {
	return c.redis.Setex(token, tokenInfo, TokenExpire)
}

//GetTokenInfo 获取token
func (c *Cache) GetTokenInfo(token string) (string, error) {
	return c.redis.Get(token)
}

//DelToken 删除token
func (c *Cache) DelToken(token string) error {
	_, err := c.redis.Del(token)
	return err
}
