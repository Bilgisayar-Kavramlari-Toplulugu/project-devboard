// middleware/rate_limit_no_redis.go
package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RateLimitNoRedisMiddleware - Redis olmadan, veritabanı tabanlı rate limiter
// APIKeyCalls tablosundan son X dakikadaki çağrı sayısını kontrol eder
func RateLimitMiddleware(db *gorm.DB) gin.HandlerFunc {
	//return func(c *gin.Context) {
	// 	// API key yoksa rate limit kontrol etme
	// 	apiKeyIDValue, exists := c.Get("api_key_id")
	// 	if !exists {
	// 		c.Next()
	// 		return
	// 	}

	// 	apiKeyID, ok := apiKeyIDValue.(uuid.UUID)
	// 	if !ok {
	// 		c.Next()
	// 		return
	// 	}

	// 	// APIKey tablosundan rate limit bilgilerini çek
	// 	var apiKey entities.APIKey
	// 	err := db.Where("id = ?", apiKeyID).First(&apiKey).Error
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"error": "Failed to retrieve API key information",
	// 		})
	// 		c.Abort()
	// 		return
	// 	}

	// 	// Rate limit aktif değilse (admin, internal key'ler) kontrolü atla
	// 	if !apiKey.APIKeyType.HasRateLimit {
	// 		c.Next()
	// 		return
	// 	}

	// 	rateLimit := apiKey.APIKeyType.RateLimit
	// 	rateLimitWindow := apiKey.APIKeyType.RateLimitWindowMinutes // dakika cinsinden

	// 	// Zaman penceresini hesapla
	// 	windowStart := time.Now().Add(-time.Duration(rateLimitWindow) * time.Minute)

	// 	// Son X dakikadaki çağrı sayısını say
	// 	var callCount int64
	// 	err = db.Model(&entities.APIKeyCall{}).
	// 		Where("api_key_id = ? AND created_at >= ?", apiKeyID, windowStart).
	// 		Count(&callCount).Error

	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"error": "Rate limit check failed",
	// 		})
	// 		c.Abort()
	// 		return
	// 	}

	// 	// Kalan istek sayısı
	// 	remaining := rateLimit - int(callCount)
	// 	if remaining < 0 {
	// 		remaining = 0
	// 	}

	// 	// Header'ları ekle
	// 	c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", rateLimit))
	// 	c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))
	// 	c.Header("X-RateLimit-Window", fmt.Sprintf("%d minutes", rateLimitWindow))

	// 	// Limit aşıldı mı?
	// 	if callCount >= int64(rateLimit) {
	// 		c.Header("Retry-After", fmt.Sprintf("%d", rateLimitWindow*60)) // saniye cinsinden

	// 		c.JSON(http.StatusTooManyRequests, gin.H{
	// 			"error":       "Rate limit exceeded",
	// 			"limit":       rateLimit,
	// 			"window":      fmt.Sprintf("%d minutes", rateLimitWindow),
	// 			"retry_after": fmt.Sprintf("%d seconds", rateLimitWindow*60),
	// 		})
	// 		c.Abort()
	// 		return
	// 	}

	// 	c.Next()
	// }
	return func(c *gin.Context) {
		c.Next()
	}
}
