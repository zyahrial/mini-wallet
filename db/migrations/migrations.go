package migrations

import (
	"khaerus/mini-wallet/db/database"
	"khaerus/mini-wallet/models"
	"log"
)

var users = []models.Account{
	models.Account{
		// ID : "119b6e43-be95-4982-98d2-9b5b9b1dcede",
		Name: "khaerus",
		Email: "khaerus@gmail.com",
		Phone: "081111111",
		Password: "password",
	},
	models.Account{
		Name: "zyahrial",
		Email:    "zyahrial@gmail.com",
		Phone: "081111110",
		Password: "password",
	},
}

func Migrate() {

	// err := database.DBCon.DropTableIfExists(&models.Account{}, &models.Wallet{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot drop table: %v", err)
	// }

	err := database.DBCon.AutoMigrate(&models.Account{}, &models.Wallet{}, &models.Deposit{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	// err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	// if err != nil {
	// 	log.Fatalf("attaching foreign key error: %v", err)
	// }

	for i, _ := range users {
		err = database.DBCon.Model(&models.Account{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		// posts[i].AuthorID = users[i].ID

		// err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		// if err != nil {
		// 	log.Fatalf("cannot seed posts table: %v", err)
		// }
	}
}