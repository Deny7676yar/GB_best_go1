package main

import (
	"context"
	"fmt"
	"github.com/Deny7676yar/Go_level2/bookcsv/internal/config"
	"github.com/Deny7676yar/Go_level2/bookcsv/internal/controllers"
	"github.com/Deny7676yar/Go_level2/bookcsv/internal/queryfile"
	"github.com/Deny7676yar/Go_level2/bookcsv/internal/readerfile"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)


func main() {


	cfg := config.Config{
		PathFile: readerfile.CSVFILEinput,
		Timeout:  100,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(cfg.Timeout))

	var cr readerfile.CSVReaderer
	resultArray, err := cr.ReadCSVFile(ctx, readerfile.CSVFILEinput)
	if err != nil {
		log.Printf("reade error: %v\n", err)
	}
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: insert|delete|search|select <arguments>")
		return
	}

	fileInfo, err := os.Stat(cfg.PathFile)
	// Is it a regular file?
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(cfg.PathFile, "not a regular file!")
		return
	}

	var control controllers.Controller
	err = control.CreateIndex(resultArray)
	if err != nil {
		fmt.Println("Cannot create index.")
		return
	}
	var cq queryfile.CSVQueryre
	// Differentiating between the commands
	switch arguments[1] {
	case "insert":
		if len(arguments) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(arguments[4], "-", "")
		if !control.MatchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		temp := control.InitS(arguments[2], arguments[3], t)
		if temp != nil {
			err := cq.Insert(temp, resultArray)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	case "delete":
		if len(arguments) != 3 {
			fmt.Println("Usage: delete Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !control.MatchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		err := cq.DeleteEntry(t, resultArray)
		if err != nil {
			fmt.Println(err)
		}
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !control.MatchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		temp := cq.Search(t, resultArray)
		if temp == nil {
			fmt.Println("Number not found:", t)
			return
		}
		fmt.Println(*temp)
	case "select":
		cq.Select(resultArray)
	default:
		fmt.Println("Not a valid option")
	}

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT)
	for {
		select {
		case <-ctx.Done():
			return
		case <-sigCh:
			cancel()
		}
	}
}
