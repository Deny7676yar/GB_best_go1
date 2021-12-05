package queryfile

import (
	"encoding/csv"
	"fmt"
	"github.com/Deny7676yar/Go_level2/bookcsv/internal/controllers"
	"github.com/Deny7676yar/Go_level2/bookcsv/internal/readerfile"
	"os"
	"strconv"
	"time"
)



type CSVQueryre interface {
	Select(data []readerfile.Entry)
    SaveCSVFile(filepath string, data []readerfile.Entry) error
	Insert(pS *readerfile.Entry, data []readerfile.Entry) error
	DeleteEntry(key string, data []readerfile.Entry)error
	Search(key string, data []readerfile.Entry) *readerfile.Entry
}

func Select(data []readerfile.Entry) {
	for _, v := range data {
		fmt.Println(v)
	}
}

func SaveCSVFile(filepath string, data []readerfile.Entry) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	for _, row := range data {
		temp := []string{row.Name, row.Surname, row.Tel, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}

func Insert(pS *readerfile.Entry, data []readerfile.Entry) error {
	csvfile := readerfile.CSVFILEinput
	// If it already exists, do not add it
	_, ok := controllers.Index[(*pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}
	data = append(data, *pS)
	// Update the index
	_ = controllers.CreateIndex(data)

	err := SaveCSVFile(csvfile, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEntry(key string, data []readerfile.Entry) error {
	csvfile := readerfile.CSVFILEinput
	i, ok := controllers.Index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}
	data = append(data[:i], data[i+1:]...)
	// Update the index - key does not exist any more
	delete(controllers.Index, key)

	err := SaveCSVFile(csvfile, data)
	if err != nil {
		return err
	}
	return nil
}

func Search(key string, data []readerfile.Entry) *readerfile.Entry {
	i, ok := controllers.Index[key]
	if !ok {
		return nil
	}
	data[i].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
	return &data[i]
}