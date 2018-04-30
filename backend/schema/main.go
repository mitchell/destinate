package schema

import (
	"fmt"
	"os"
	"sort"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Necessary in order to use postgres dialect of GORM.
)

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%v port=5432 user=destinate dbname=destinate password=devPassword", os.Getenv("DBHOST")))
	if err != nil {
		panic(err)
	}

	return db
}

// Migrate is to be called from the migrate lambda function in order to migrate the DB
func Migrate() {
	db := connectDB()
	defer db.Close()
	db.AutoMigrate(&User{})
}

// CORSHeaders is used used to set the headers of each endpoint to allow CORS requests.
var CORSHeaders = map[string]string{
	"Access-Control-Allow-Origin":      "*",
	"Access-Control-Allow-Credentials": "true",
}

/*
 * Everything following was sourced from https://groups.google.com/forum/#!topic/golang-nuts/FT7cjmcL7gw and was written by Andrew Gerrand
 * Changes:
 *   - Exported RankByWordCount
 *   - Changed Value in Pair from int to float64
 */

// RankByWordCount is used to do what it says...
func RankByWordCount(wordFrequencies map[string]float64) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

// Pair is as key value pair in the form of a struct
type Pair struct {
	Key   string
	Value float64
}

// PairList is a slice of Pairs
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
