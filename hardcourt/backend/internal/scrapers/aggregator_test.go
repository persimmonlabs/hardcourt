package scrapers

import (
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestAggregator_Creation(t *testing.T) {
	// Test with nil repos (just checking initialization logic)
	agg := NewAggregator(nil, nil, nil)

	if agg == nil {
		t.Fatal("Expected aggregator to be created, got nil")
	}

	if agg.sofascore == nil {
		t.Error("Expected Sofascore client to be initialized")
	}

	if agg.limiter == nil {
		t.Error("Expected rate limiter to be initialized")
	}

	if agg.cache == nil {
		t.Error("Expected cache to be initialized")
	}

	if agg.cacheExpiry != 30*time.Second {
		t.Errorf("Expected cache expiry 30s, got %v", agg.cacheExpiry)
	}
}

func TestAggregator_CacheOperations(t *testing.T) {
	agg := NewAggregator(nil, nil, nil)

	// Test cache miss
	_, exists := agg.GetCachedMatch("nonexistent")
	if exists {
		t.Error("Expected cache miss for nonexistent match")
	}

	// Clear cache should not panic
	agg.ClearExpiredCache()

	// Verify cache is still usable after clear
	if agg.cache == nil {
		t.Error("Cache should still be initialized after clear")
	}
}

func TestAggregator_RateLimiting(t *testing.T) {
	agg := NewAggregator(nil, nil, nil)

	// Limiter should be configured for 1 request per 2 seconds
	if agg.limiter.Limit() != rate.Every(2*time.Second) {
		t.Errorf("Expected rate limit of 1 per 2s, got %v", agg.limiter.Limit())
	}

	if agg.limiter.Burst() != 1 {
		t.Errorf("Expected burst of 1, got %d", agg.limiter.Burst())
	}
}
