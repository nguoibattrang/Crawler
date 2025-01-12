package scheduler

import (
	"time"
)

func StartScheduler(interval time.Duration, task func()) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		task()
	}
}
