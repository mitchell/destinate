package schema

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// User is the definition of the User model in the db and in the service
type User struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Password string
	Location string
	Radius   uint
	Likes    pq.StringArray `gorm:"type:text[]"`
	Dislikes pq.StringArray `gorm:"type:text[]"`
	Scores   postgres.Jsonb
}

// All finds all records of this entity type.
func (u User) All() string {
	var us []User
	db := connectDB()
	defer db.Close()

	db.Find(&us)
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
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 15)
	if err != nil {
		return "", err
	}
	u.Password = string(hash)

	db.Create(u)
	db.First(u)
	jsonba, _ := json.Marshal(u)

	return string(jsonba), nil
}

// Auth compares sent password with pass in db
func (u *User) Auth(jsons string) error {
	db := connectDB()
	defer db.Close()

	if err := json.Unmarshal([]byte(jsons), u); err != nil {
		return err
	}

	pass := u.Password
	db.Where("name = ?", u.Name).First(u)

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	if err != nil {
		return err
	}

	return nil
}

// Find finds a User record by the provided string.
func (u *User) Find(sid string) error {
	db := connectDB()
	defer db.Close()

	id, err := strconv.ParseUint(sid, 10, 32)
	if err != nil {
		return err
	}
	db.First(u, uint(id))

	return nil
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
	db.First(u)
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

// LikeDestination adds a destination to the user specified by id and destination specified by did.
func (u *User) LikeDestination(ctx context.Context, sid string, did string) error {
	db := connectDB()
	defer db.Close()

	resp, err := FindDestination(ctx, did)

	id, err := strconv.ParseUint(sid, 10, 32)
	if err != nil {
		return err
	}
	db.First(u, uint(id))
	u.Likes = append(u.Likes, did)

	val, err := u.Scores.Value()
	jmap := map[string]float64{}

	if val != nil {
		if err := json.Unmarshal(val.([]byte), &jmap); err != nil {
			return err
		}
	}

	for _, t := range resp.Types {
		jmap[t]++
	}

	jsonba, err := json.Marshal(&jmap)
	if err != nil {
		return err
	}
	u.Scores = postgres.Jsonb{
		json.RawMessage(string(jsonba)),
	}

	db.Save(u)
	db.First(u)

	return nil
}

// DislikeDestination adds a destination to the user specified by id and destination specified by did.
func (u *User) DislikeDestination(ctx context.Context, sid string, did string) error {
	db := connectDB()
	defer db.Close()

	resp, err := FindDestination(ctx, did)

	id, err := strconv.ParseUint(sid, 10, 32)
	if err != nil {
		return err
	}
	db.First(u, uint(id))
	u.Dislikes = append(u.Dislikes, did)

	val, err := u.Scores.Value()
	jmap := map[string]float64{}

	if val != nil {
		if err := json.Unmarshal(val.([]byte), &jmap); err != nil {
			return err
		}
	}

	for _, t := range resp.Types {
		jmap[t]--
	}

	jsonba, err := json.Marshal(&jmap)
	if err != nil {
		return err
	}
	u.Scores = postgres.Jsonb{
		json.RawMessage(string(jsonba)),
	}

	db.Save(u)
	db.First(u)

	return nil
}
