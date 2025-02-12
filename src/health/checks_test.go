package health

import (
    "context"
    "database/sql"
    "testing"

    "github.com/go-redis/redis/v8"
    "github.com/stretchr/testify/assert"
)

func TestHealthChecker_Check(t *testing.T) {
    // Setup
    db, _ := sql.Open("postgres", "user=postgres password=postgres dbname=testdb sslmode=disable")
    redisClient := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
    healthChecker := NewHealthChecker(db, redisClient)

    // Run health check
    health := healthChecker.Check(context.Background())

    assert.Equal(t, "healthy", health.Status)
    assert.NotEmpty(t, health.Details)
}
