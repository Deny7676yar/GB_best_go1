package main

//Приложение запускается с аргументами командной строки
//Введите при запуске go run main: insert|delete|search|select <arguments>

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Deny7676yar/Go_level2/bookcsv/internal/config"
	"github.com/Deny7676yar/Go_level2/bookcsv/internal/queryfile"
	"github.com/Deny7676yar/Go_level2/bookcsv/internal/readerfile"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	log.WithFields(log.Fields{
		"Start program": time.Now(),
	}).Info()

	cfg := config.Config{
		PathFile: queryfile.CSVFILEinput,
		Timeout:  2,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(cfg.Timeout))
	var r readerfile.CSVReaderer
	var q queryfile.CSVQueryre

	q = queryfile.NewCSVQuery(time.Duration(cfg.Timeout) * time.Second)
	log.WithFields(log.Fields{
		"New CSV query": q,
	}).Debug()

	r = readerfile.NewReaderCSV(q)
	log.WithFields(log.Fields{
		"New Reader": r,
	}).Debug()

	resultArray, err := r.ReadCSVFile(ctx, cfg.PathFile)
	if err != nil {
		log.WithFields(log.Fields{
			"reader error:": err,
		}).Errorf("Do not read file")
	}

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: insert|delete|search|select <arguments>")
		return
	}

	err = queryfile.CreateIndex(resultArray)
	if err != nil {
		log.WithFields(log.Fields{
			"Create index error:": err,
		}).Errorf("Cannot create index.")
		return
	}

	switch arguments[1] {
	case "insert":
		if len(arguments) != 5 {
			log.WithFields(log.Fields{
				"Insert error:": err,
			}).Errorf("Usage: insert Name Surname Telephone.")
			return
		}
		t := strings.ReplaceAll(arguments[4], "-", "")
		if !queryfile.MatchTel(t) {
			log.WithFields(log.Fields{
				"Insert error:": err,
			}).Errorf("Error: %v, Not a valid telephone number: %s", err, t)
			return
		}
		temp := queryfile.InitS(arguments[2], arguments[3], t)
		if temp != nil {
			err := q.Insert(ctx, temp, resultArray)
			if err != nil {
				log.WithFields(log.Fields{
					"Inits error:": err,
				}).Errorf("Error: %v", err)
				return
			}
		}
	case "delete":
		if len(arguments) != 3 {
			log.WithFields(log.Fields{
				"Inits error:": err,
			}).Errorf("Usage: delete Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !queryfile.MatchTel(t) {
			log.WithFields(log.Fields{
				"Delete error:": err,
			}).Errorf("Error: %v, Not a valid telephone number: %s", err, t)
			return
		}
		err := q.DeleteEntry(ctx, t, resultArray)
		if err != nil {
			fmt.Println(err)
		}
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !queryfile.MatchTel(t) {
			log.WithFields(log.Fields{
				"Search error:": err,
			}).Errorf("Error: %v, Not a valid telephone number: %s", err, t)
			return
		}
		temp := q.Search(ctx, t, resultArray)
		if temp == nil {
			log.WithFields(log.Fields{
			}).Errorf("Number not found: %s", t)
			return
		}
		fmt.Println(*temp)
	case "select":
		q.Select(ctx, resultArray)
	default:
		log.WithFields(log.Fields{
			"Not a valid option": "err",
		}).Debug()
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	for {
		select {
		case <-sigCh:
			cancel()
			log.WithFields(log.Fields{
				"SIGINT": <-sigCh,
			}).Info("cencel context")
		case <-ctx.Done():
			return
		}
	}
}
