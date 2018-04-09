package schema

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Necessary in order to use postgres dialect of GORM.
	"os"
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
	db.AutoMigrate(&Review{}, &User{})
}
