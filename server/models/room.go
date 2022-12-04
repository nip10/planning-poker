package models

import (
	"fmt"
	"log"
	"planning-poker/server/db"

	"github.com/alexedwards/argon2id"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Password string `json:"password"`
}

type RoomModel struct{}

// Create ...
func (roomModel RoomModel) Create(room Room) {
	db.GetDB().Create(&room)
}

// One ...
func (roomModel RoomModel) One(room Room) {
	db.GetDB().Find(&room, room.ID)
}

// All ...
func (roomModel RoomModel) All(room []Room) {
	db.GetDB().Find(&room)
}

// add beforesave hook to hash password
func (room *Room) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Println("before save")
	fmt.Println(room.Password)
	if room.Password != "" {
		hash, err := argon2id.CreateHash("pa$$word", argon2id.DefaultParams)
		if err != nil {
			log.Fatal(err)
		}

		room.Password = hash
	}

	fmt.Println("after save", room.Password)
	return
}
