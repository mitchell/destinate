package schema

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"strconv"
)

// User is the definition of the User model in the db and in the service
type User struct {
	gorm.Model
	Name        string
	Password    string
	Reviews     []Review
	Location    string
	Preferences string
}

// All finds all records of this entity type.
func (u User) All() string {
	var us []User
	db := connectDB()
	defer db.Close()

	db.Preload("Reviews").Find(&us)
	jsonba, _ := json.Marshal(us)

	return string(jsonba)
}

// Create creates a User record according to the json passed in by string.
func (u *User) Create(jsons string) (string, error) {
	db := connectDB()
	defer db.Close()

	if err := json.Unmarshal([]byte(jsons), u); err != nil {
		return "", err
	}
	db.Create(u)
	db.Preload("Reviews").First(u)
	jsonba, _ := json.Marshal(u)

	return string(jsonba), nil
}

// Find finds a User record by the provided string.
func (u *User) Find(sid string) (string, error) {
	db := connectDB()
	defer db.Close()

	id, err := strconv.ParseUint(sid, 10, 32)
	if err != nil {
		return "", err
	}
	db.Preload("Reviews").First(u, uint(id))
	jsonba, _ := json.Marshal(u)

	return string(jsonba), nil
}

// Update updates a User record according to the provided json.
func (u *User) Update(jsons string) (string, error) {
	db := connectDB()
	defer db.Close()

	if err := json.Unmarshal([]byte(jsons), u); err != nil {
		return "", err
	}
	db.First(u)
	if err := json.Unmarshal([]byte(jsons), u); err != nil {
		return "", err
	}
	db.Save(u)
	db.Preload("Reviews").First(u)
	jsonba, _ := json.Marshal(u)

	return string(jsonba), nil
}

// Destroy soft deletes a record from the database.
func (u *User) Destroy(sid string) error {
	db := connectDB()
	defer db.Close()

	id, err := strconv.ParseUint(sid, 10, 32)
	if err != nil {
		return err
	}
	db.First(u, uint(id))
	db.Delete(u)

	return nil
}
