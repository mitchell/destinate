package schema

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"strconv"
)

// Review is the definition of the Review model in the db and in the service
type Review struct {
	gorm.Model
	DestinationID string
	User          User
	UserID        uint
	Comment       string
	Rating        uint
}

// All finds all records of this entity type.
func (r Review) All() string {
	var rs []Review
	db := connectDB()
	defer db.Close()

	db.Preload("Destination").Preload("User").Find(&rs)
	jsonba, _ := json.Marshal(rs)

	return string(jsonba)
}

// Create creates a Review record according to the json passed in by string.
func (r *Review) Create(jsons string) (string, error) {
	db := connectDB()
	defer db.Close()

	if err := json.Unmarshal([]byte(jsons), r); err != nil {
		return "", err
	}
	db.Create(r)
	db.Preload("Destination").Preload("User").First(r)
	jsonba, _ := json.Marshal(r)

	return string(jsonba), nil
}

// Find finds a Review record by the provided string.
func (r *Review) Find(sid string) (string, error) {
	db := connectDB()
	defer db.Close()

	id, err := strconv.ParseUint(sid, 10, 32)
	if err != nil {
		return "", err
	}
	db.Preload("Destination").Preload("User").First(r, uint(id))
	jsonba, _ := json.Marshal(r)

	return string(jsonba), nil
}

// Update updates a Review record according to the provided json.
func (r *Review) Update(jsons string) (string, error) {
	db := connectDB()
	defer db.Close()

	if err := json.Unmarshal([]byte(jsons), r); err != nil {
		return "", err
	}
	db.First(r)
	if err := json.Unmarshal([]byte(jsons), r); err != nil {
		return "", err
	}
	db.Save(r)
	db.Preload("Destination").Preload("User").First(r)
	jsonba, _ := json.Marshal(r)

	return string(jsonba), nil
}

// Destroy soft deletes a record from the database.
func (r *Review) Destroy(sid string) error {
	db := connectDB()
	defer db.Close()

	id, err := strconv.ParseUint(sid, 10, 32)
	if err != nil {
		return err
	}
	db.First(r, uint(id))
	db.Delete(r)

	return nil
}
