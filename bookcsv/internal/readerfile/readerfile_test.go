package readerfile

import (
	"bytes"
	"context"
	"github.com/Deny7676yar/Go_level2/bookcsv/internal/queryfile"
	"github.com/Deny7676yar/Go_level2/bookcsv/internal/queryfile/mocks"
	"log"
	"reflect"
	"sync"
	"testing"
	"time"
)

type TestCaseType struct {
	filepath  string
	resultArray []queryfile.Entry
	ctx         context.Context
}

//Name,Surname,Tel,LastAccess
//Robeson,Modestine,Gorton,Budget
//Carillo,Nikaniki,Shasnan,Rathbourne

func TestNewReaderCSV(t *testing.T) {

	type args struct {
		data []TestCaseEntry
	}
	tests := []struct {
		name   string
		args   args
		filepath string
	}{
		{
			name: "testSelect1",
			args: args{
				data: []TestCaseEntry{
					{Name: "Name", Surname: "Surname", Tel: "Tel", LastAccess: "LastAccess"},
					{Name: "Robeson", Surname: "Modestine", Tel: "Gorton", LastAccess: "Budget"},
					{Name: "Carillo", Surname: "Nikaniki", Tel: "Shasnan", LastAccess: "Rathbourne"},
				},
			},
			filepath: "./test1.csv",
		},
	}
	cfg := config.Config{
		PathFile: "./test1.csv",
		Timeout:  10,
	}
	ctx := context.Background()
	var r readerfile.CSVReaderer

	q := NewCSVQuery(time.Duration(cfg.Timeout) * time.Second)
	r = readerfile.NewReaderCSV(q)

	gotResultArray, err := r.ReadCSVFile(ctx, cfg.PathFile)
	if err != nil {
		log.Printf("reade error: %v\n", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err, _ := r.ReadCSVFile(ctx, cfg.PathFile)(w, err !=nil); (err!=nil){

			}
			if gotW:= gotResultArray; gotW != tt.args.data{
				t.Errorf("ReadCSVFile() = %v, want %v", gotW, tt.args.data)
			}
		})
	}
}



func TestReaderCSV_ReadCSVFile(t *testing.T) {
	type fields struct {
		q           mocks.CSVQueryre
		chanPlaces  chan queryfile.Entry
		resultArray []queryfile.Entry
		mu          sync.RWMutex
	}
	type args struct {
		ctx      context.Context
		filepath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []queryfile.Entry
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ReaderCSV{
				q:           tt.fields.q,
				chanPlaces:  tt.fields.chanPlaces,
				resultArray: tt.fields.resultArray,
				mu:          tt.fields.mu,
			}
			got, err := r.ReadCSVFile(tt.args.ctx, tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCSVFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadCSVFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_worker(t *testing.T) {
	type args struct {
		filePath string
		resChan  chan queryfile.Entry
		wg       *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

