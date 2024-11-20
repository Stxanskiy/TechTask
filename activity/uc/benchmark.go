package uc

import (
	"PostgreBenchmark/activity/model"
	"PostgreBenchmark/config"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/nevasik7/lg"
	"sync"
	"time"
)

func RunBenchmark(config config.PostgresConfig, db *pgxpool.Pool) (model.BenchmarkResult, error) {
	var (
		wg           sync.WaitGroup
		requestCount int
		mu           sync.Mutex
	)
	stopChannel := make(chan struct{})
	startTime := time.Now()
	duration := time.Duration(config.DurationMs) * time.Millisecond

	worker := func() {
		defer wg.Done()
		for {
			select {
			case <-stopChannel:
				return
			default:
				_, err := db.Exec(context.Background(), config.SQLQuery)
				if err != nil {
					lg.Printf("Ошибка запроса: %v", err)
					continue
				}
				mu.Lock()
				requestCount++
				mu.Unlock()
			}

		}
	}

	wg.Add(config.Concurrency)
	for i := 0; i < config.Concurrency; i++ {
		go worker()
	}

	time.Sleep(duration)
	close(stopChannel)
	wg.Wait()

	elapsedTime := time.Since(startTime).Seconds()
	return model.BenchmarkResult{
		TotalRequests: requestCount,
		RPS:           float64(requestCount) / elapsedTime,
	}, nil

}
