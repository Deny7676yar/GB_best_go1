package readerfile

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"sync"
)

var CSVFILEinput = "/home/den/Go_level2/bookcsv/internal/readfile/MOCK_DATA.csv"

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

type CSVReaderer interface {
	ReadCSVFile(ctx context.Context,  filepath string)([]Entry, error)
}

type ReaderCSV struct {
	r CSVReaderer
	resChan chan Entry
	pathCSV string
	resultArray []Entry
	mu sync.RWMutex
	wg *sync.WaitGroup
}



func (r *ReaderCSV)ReadCSVFile(ctx context.Context,  filepath string)([]Entry, error){

	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, err
	default:
		r.wg.Add(1)
		go r.worker(filepath, r.resChan)
	}
	go func() {
		for data := range r.resChan{
			r.resultArray = append(r.resultArray, data)
		}
	}()

	r.wg.Wait()

	return r.resultArray, nil
}

func (r *ReaderCSV)worker(filePath string, resChan chan Entry) {
	defer r.wg.Done()

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("open file error: %v\n", err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalf("read csv error: %v\n", err)
	}
	for _, line := range lines{
		data := Entry{
			Name:       line[0],
			Surname:    line[1],
			Tel:        line[2],
			LastAccess: line[3],
		}
		resChan <- data
	}
}