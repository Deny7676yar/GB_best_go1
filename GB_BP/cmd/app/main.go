package main

import (
	"context"
	crawler2 "github.com/Deny7676yar/Go_level2/GB_BP/internal/crawler"
	"github.com/Deny7676yar/Go_level2/GB_BP/internal/services"
	log "github.com/sirupsen/logrus"
	cfg "github.com/Deny7676yar/Go_level2/GB_BP/internal/config"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	//Обьявляем поле с информациеей о стартее
	log.WithFields(log.Fields{
		"Start crawler": time.Now(),
	}).Info()

	cfg := cfg.Config{
		MaxDepth:   3,
		MaxResults: 10,
		MaxErrors:  5,
		Url:        "https://telegram.org",
		Timeout:    10,
	}
	var cr crawler2.Crawler
	var r services.Requester
	wg := &sync.WaitGroup{}
	wg.Add(1)

	//server := &UrlGetServer{NewInMemoryGet()}
	//log.Fatal(http.ListenAndServe(":5000", server))

	r = services.NewRequester(time.Duration(cfg.Timeout) * time.Second)
	log.WithFields(log.Fields{
		"New request": r,
	}).Debug()

	cr = crawler2.NewCrawler(r)
	log.WithFields(log.Fields{
		"New Crawler": cr,
	}).Debug()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(cfg.Timeout))//общий таймаут в секундах

	go cr.Scan(ctx, wg, cfg.Url, cfg.MaxDepth) //Запускаем краулер в отдельной рутине

	go crawler2.ProcessResult(ctx, cancel, cr, cfg) //Обрабатываем результаты в отдельной рутине

	crawler := crawler2.SearchDepthCrawler(cfg.MaxDepth)
	go crawler2.SigDepth(ctx, crawler, 2)

	sigCh := make(chan os.Signal)        //Создаем канал для приема сигналов
	signal.Notify(sigCh, syscall.SIGINT) //Подписываемся на сигнал SIGINT

	for {
		select {
		case <-ctx.Done(): //Если всё завершили - выходим
			return
		case <-sigCh:
			log.WithFields(log.Fields{
				"SIGINT": <-sigCh,
			}).Info("cencel context")
			cancel() //Если пришёл сигнал SigInt - завершаем контекст
		}
	}

	wg.Wait()

	log.WithFields(log.Fields{
		"Wait ": "cancel crawler",
	}).Info()

}