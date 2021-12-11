package readerfile

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"sync"

	"github.com/Deny7676yar/Go_level2/bookcsv/internal/queryfile"
)

type CSVReaderer interface {
	ReadCSVFile(ctx context.Context, filepath string) ([]queryfile.Entry, error)
}

type ReaderCSV struct {
	q           queryfile.CSVQueryre
	chanPlaces  chan queryfile.Entry
	resultArray []queryfile.Entry
	mu          sync.RWMutex
}

func NewReaderCSV(q queryfile.CSVQueryre) *ReaderCSV {
	return &ReaderCSV{
		q:           q,
		chanPlaces:  make(chan queryfile.Entry),
		resultArray: []queryfile.Entry{},
		mu:          sync.RWMutex{},
	}
}

func (r *ReaderCSV) ReadCSVFile(ctx context.Context, filepath string) ([]queryfile.Entry, error) {
	var wg = &sync.WaitGroup{}

	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, err
	default:
		wg.Add(1)
		go worker(filepath, r.chanPlaces, wg)
	}
	go func() {
		for data := range r.chanPlaces {
			r.resultArray = append(r.resultArray, data)
		}
	}()

	wg.Wait()

	return r.resultArray, nil
}

func worker(filePath string, resChan chan queryfile.Entry, wg *sync.WaitGroup) {
	defer wg.Done()

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("open file error: %v\n", err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalf("read csv error: %v\n", err)
	}
	for _, line := range lines {
		data := queryfile.Entry{
			Name:       line[0],
			Surname:    line[1],
			Tel:        line[2],
			LastAccess: line[3],
		}
		resChan <- data
	}
}
