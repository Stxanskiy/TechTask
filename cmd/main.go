package main

import (
	"PostgreBenchmark/activity/repo"
	"PostgreBenchmark/activity/uc"
	"PostgreBenchmark/config"
	"gitlab.com/nevasik7/lg"
	"log"
)

func main() {
	lg.Init()

	cfg, err := config.MustLoad()
	if err != nil {
		lg.Panicf("Конфигурация не настроена:%v", err)
	} else {
		lg.Info("Загрузка конфигурации прошла успешно!")
	}

	db, err := repo.ConnectDB(cfg.DSN)
	if err != nil {
		lg.Fatalf("Не удалоь подключиться к базе данных %v", err)
	} else {
		lg.Info("Покдлючение к базе днных  прошло успешно")
	}
	defer db.Close()

	result, err := uc.RunBenchmark(cfg, db)
	if err != nil {
		lg.Fatalf("Не удалось запустить Бенчмарк! %v ", err)
	} else {
		lg.Info("Бенчмарк Запущен!")
	}
	lg.Printf("Количество запросов: %d", result.TotalRequests)
	log.Fatalf("RPS: %.2f", result.RPS)

}
