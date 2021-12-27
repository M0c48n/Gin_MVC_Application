package user

import (
	"Gin_MVC/model/database"
	"Gin_MVC/model/notify"
	"Gin_MVC/model/priority"
	"log"
	"testing"

	"gorm.io/gorm"
)

func TestUser(t *testing.T) {
	_ = database.DBConnection()

	database.Migrator([]interface{}{&User{}, &notify.Notify{}, &priority.Priority{}})
	// log.Println(GetUser("saitou").Name)

	var u = User{
		//UserId:       0,
		Name:     "斉藤",
		Ruby:     "サイトウ",
		Username: "saito",
		Password: "asdfgdf",
		Tel:      "03000000000",
		Location: 0,
		Publish:  false,
	}
	err := database.DB.Transaction(
		func(tx *gorm.DB) error {
			err := UserCreate(&u)
			if err != nil {
				return err
			}
			err = notify.InitNotify(u.Id)
			if err != nil {
				return err
			}
			priority.CreatePriority(u.Id)
			if err != nil {
				return err
			}
			return nil
		})
	if err != nil {
		log.Panicln(err)
	}
}