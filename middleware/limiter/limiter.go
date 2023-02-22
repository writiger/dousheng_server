package limiter

import (
	"context"
	"dousheng_server/middleware/rc"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-redis/redis/v8"
	"time"
)

const tokenBucketLimiterTryAcquireRedisScript = `
-- ARGV[1]: 容量
-- ARGV[2]: 发放令牌速率/秒
-- ARGV[3]: 当前时间（秒）

local capacity = tonumber(ARGV[1])
local rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])

local lastTime = tonumber(redis.call("hget", KEYS[1], "lastTime"))
local currentTokens = tonumber(redis.call("hget", KEYS[1], "currentTokens"))
-- 初始化
if lastTime == nil then 
   lastTime = now
   currentTokens = capacity
   redis.call("hmset", KEYS[1], "currentTokens", currentTokens, "lastTime", lastTime)
end 

-- 尝试发放令牌
-- 距离上次发放令牌的时间
local interval = now - lastTime
if interval > 0 then
   -- 当前令牌数量+距离上次发放令牌的时间(秒)*发放令牌速率
   local newTokens = currentTokens + interval * rate
   if newTokens > capacity then 
      newTokens = capacity
   end 
   currentTokens = newTokens
   redis.call("hmset", KEYS[1], "currentTokens", newTokens, "lastTime", now)
end

-- 如果没有令牌，请求失败
if currentTokens == 0 then
   return 0
end
-- 果有令牌，当前令牌-1，请求成功
redis.call("hincrby", KEYS[1], "currentTokens", -1)
redis.call("expire", KEYS[1], capacity / rate)
return 1
`

var limiter *TokenBucketLimiter

func InitLimiter() {
	limiter = &TokenBucketLimiter{
		capacity: 10,
		rate:     10,
		client:   rc.RedisClient,
		script:   redis.NewScript(tokenBucketLimiterTryAcquireRedisScript),
	}
}

// TokenBucketLimiter 令牌桶限流器
type TokenBucketLimiter struct {
	capacity int           // 容量
	rate     int           // 发放令牌速率/秒
	client   *redis.Client // Redis客户端
	script   *redis.Script // TryAcquire脚本
}

func (l *TokenBucketLimiter) TryAcquire(ctx context.Context, resource string) error {
	// 当前时间
	now := time.Now().Unix()
	success, err := l.script.Run(ctx, l.client, []string{resource}, l.capacity, l.rate, now).Bool()
	if err != nil {
		return err
	}
	// 若到达窗口请求上限，请求失败
	if !success {
		return errors.New("bucket script failed")
	}
	return nil
}

// LimiterMiddleware 封装为中间件
func LimiterMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		err := limiter.TryAcquire(context.Background(), "currentTokens")
		if err != nil {
			c.AbortWithMsg("too many requests, server traffic limiting", 1)
		}
		c.Next(ctx)
	}
}
