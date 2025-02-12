package health

import (
    "context"
    "database/sql"
    "github.com/go-redis/redis/v8"
    "time"
)

type HealthChecker struct {
    db          *sql.DB
    redisClient *redis.Client
}

type Health struct {
    Status    string            `json:"status"`
    Details   map[string]string `json:"details"`
    Timestamp time.Time         `json:"timestamp"`
}

type User struct {
    gorm.Model
    Username  string `gorm:"uniqueIndex;not null;index:idx_username"`
    Password  string `gorm:"not null"`
    Role      string `gorm:"not null;index:idx_role"`
}

func NewHealthChecker(db *sql.DB, redisClient *redis.Client) *HealthChecker {
    return &HealthChecker{
        db:          db,
        redisClient: redisClient,
    }
}

func (h *HealthChecker) Check(ctx context.Context) Health {
    health := Health{
        Status:    "healthy",
        Details:   make(map[string]string),
        Timestamp: time.Now(),
    }
    
    // Check database
    if err := h.db.PingContext(ctx); err != nil {
        health.Status = "unhealthy"
        health.Details["database"] = err.Error()
    } else {
        health.Details["database"] = "connected"
    }
    
    // Check Redis
    if err := h.redisClient.Ping(ctx).Err(); err != nil {
        health.Status = "unhealthy"
        health.Details["redis"] = err.Error()
    } else {
        health.Details["redis"] = "connected"
    }
    
    return health
}
