package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type LimiterIface interface {
	Key(c *gin.Context) string                          // generate key for each bucket
	GetBucket(key string) (*ratelimit.Bucket, bool)     // get token bucket by key
	AddBuckets(rules ...LimiterBucketRule) LimiterIface // add token buckets
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string        // key, usually is the api path
	FillInterval time.Duration // token fill interval
	Capacity     int64         // token bucket capacity
	Quantum      int64         // token fill quantum
}
