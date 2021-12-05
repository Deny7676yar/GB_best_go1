package controllers

import (
	"github.com/Deny7676yar/Go_level2/bookcsv/internal/readerfile"
	"regexp"
	"strconv"
	"time"
)

type Controller interface {
	CreateIndex(data []readerfile.Entry) error
	InitS(N, S, T string) *readerfile.Entry
	MatchTel(s string) bool
}

var Index map[string]int

func CreateIndex(data []readerfile.Entry) error {
	Index = make(map[string]int)
	for i, k := range data {
		key := k.Tel
		Index[key] = i
	}
	return nil
}

// Initialized by the user – returns a pointer
// If it returns nil, there was an error
func InitS(N, S, T string) *readerfile.Entry {
	// Both of them should have a value
	if T == "" || S == "" {
		return nil
	}
	// Give LastAccess a value
	LastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &readerfile.Entry{Name: N, Surname: S, Tel: T, LastAccess: LastAccess}
}

func MatchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}