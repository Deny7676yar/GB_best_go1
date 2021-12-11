package queryfile

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

var CSVFILEinput = "/home/den/Go_level2/bookcsv/MOCK_DATA.csv"

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

type CSVQueryre interface {
	Select(ctx context.Context, data []Entry)
	SaveCSVFile(ctx context.Context, filepath string, data []Entry) error
	Insert(ctx context.Context, pS *Entry, data []Entry) error
	DeleteEntry(ctx context.Context, key string, data []Entry) error
	Search(ctx context.Context, key string, data []Entry) *Entry
}

type csvQuery struct {
	timeout time.Duration
}

func NewCSVQuery(timeout time.Duration) *csvQuery {
	return &csvQuery{timeout: timeout}
}

func (c *csvQuery) Select(ctx context.Context, data []Entry) {
	select {
	case <-ctx.Done():
		return
	default:
		for _, v := range data {
			fmt.Println(v)
		}
	}
}

func (c *csvQuery) SaveCSVFile(ctx context.Context, filepath string, data []Entry) error {
	select {
	case <-ctx.Done():
		return nil
	default:
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
}

func (c *csvQuery) Insert(ctx context.Context, pS *Entry, data []Entry) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		csvfile := CSVFILEinput
		_, ok := Index[(*pS).Tel]
		if ok {
			return fmt.Errorf("%s already exists", pS.Tel)
		}
		data = append(data, *pS)
		_ = CreateIndex(data)

		err := c.SaveCSVFile(ctx, csvfile, data)
		if err != nil {
			return err
		}
		return nil
	}
}

func (c *csvQuery) DeleteEntry(ctx context.Context, key string, data []Entry) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		csvfile := CSVFILEinput
		i, ok := Index[key]
		if !ok {
			return fmt.Errorf("%s cannot be found!", key)//nolint
		}
		data = append(data[:i], data[i+1:]...)
		delete(Index, key)

		err := c.SaveCSVFile(ctx, csvfile, data)
		if err != nil {
			return err
		}
		return nil
	}
}

func (c *csvQuery) Search(ctx context.Context, key string, data []Entry) *Entry {
	select {
	case <-ctx.Done():
		return nil
	default:
		i, ok := Index[key]
		if !ok {
			return nil
		}
		data[i].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
		return &data[i]
	}
}

type Controller interface {
	CreateIndex(data []Entry) error
	InitS(N, S, T string) *Entry
	MatchTel(s string) bool
}

var Index map[string]int

func CreateIndex(data []Entry) error {
	Index = make(map[string]int)
	for i, k := range data {
		key := k.Tel
		Index[key] = i
	}
	return nil
}

func InitS(N, S, T string) *Entry {
	if T == "" || S == "" {
		return nil
	}
	LastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &Entry{Name: N, Surname: S, Tel: T, LastAccess: LastAccess}
}

func MatchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}
