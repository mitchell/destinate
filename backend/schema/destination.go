package schema

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"strconv"
)

// Destination is the definition of the Destination model in the db and in the service
type Destination struct {
	gorm.Model
	Name        string
	Address     string
	Description string
}

// All finds all records of this entity type.
func (d Destination) All() string {
	var ds []Destination
	db := connectDB()
	defer db.Close()

	db.Find(&ds)
	jsonba, _ := json.Marshal(ds)

	return string(jsonba)
}

// Create creates a Destination record according to the json passed in by string.
func (d *Destination) Create(jsons string) (string, error) {
	db := connectDB()
	defer db.Close()

	if err := json.Unmarshal([]byte(jsons), d); err != nil {
		return "", err
	}
	db.Create(d)
	jsonba, _ := json.Marshal(d)

	return string(jsonba), nil
}

// Find finds a Destination record by the provided string.
func (d *Destination) Find(sid string) (string, error) {
	db := connectDB()
	defer db.Close()

	id, err := strconv.ParseUint(sid, 10, 32)
	if err != nil {
		return "", err
	}
	db.First(d, uint(id))
	jsonba, _ := json.Marshal(d)

	return string(jsonba), nil
}

// Update updates a Destination record according to the provided json.
func (d *Destination) Update(jsons string) (string, error) {
	db := connectDB()
	defer db.Close()

	if err := json.Unmarshal([]byte(jsons), d); err != nil {
		return "", err
	}
	db.First(d)
	if err := json.Unmarshal([]byte(jsons), d); err != nil {
		return "", err
	}
	db.Save(d)
	jsonba, _ := json.Marshal(d)

	return string(jsonba), nil
}

// Destroy soft deletes a record from the database.
func (d *Destination) Destroy(sid string) error {
	db := connectDB()
	defer db.Close()

	id, err := strconv.ParseUint(sid, 10, 32)
	if err != nil {
		return err
	}
	db.First(d, uint(id))
	db.Delete(d)

	return nil
}
