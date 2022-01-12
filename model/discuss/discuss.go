package discuss

import (
	"Gin_MVC/model/user"
	"encoding/json"
	"hash"
)

type Discuss struct {
	Id           int `gorm:"primaryKey;autoIncrement"`
	Ref_Id       int `gorm:"primaryKey;autoIncremant"`
	Create_User  user.User
	Discuss_Type int
	Opened       int
	Content      json.RawMessage `json:"content"`
}

type ContentJSON []struct {
	Title       string    `json:"title"`
	Hash        hash.Hash `json:"hash"`
	Create_User int       `json:"createUser"`
	Body        string    `json:"body"`
	MentionTo   []int     `json:"mentionTo"`
}

// func CreateDiscuss(discuss Discuss) error{
// 	var s string
// 	_ = json.NewEncoder().Encode(Discuss{})

// 	return database.DB.Create(&Discuss{
// 		Content: s,
// 	}).Error
// }