package users

import (
	"log"
	"mvc-gorm/helpers/general"
	"mvc-gorm/models"
	"os"
)

//GetUserDetail return user details from db
func GetUserDetail(userID uint) models.User {
	db := general.GetDB()
	defer general.CloseDB(db)

	var user models.User
	sql := "SELECT id, username, first_name, last_name, image FROM `users` WHERE id = ?"

	err := db.Raw(sql, userID).Row().Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Image)
	if err != nil {
		log.Fatalln(err)
	}
	return user
}

//DeleteOldProfileImage delete profile image which is not required
func DeleteOldProfileImage(userID uint) {
	user := GetUserDetail(userID)

	if user.Image.String != "noimg.png" {
		err := os.Remove("static/images/" + user.Image.String)
		if err != nil {
			log.Println(err)
		}
	}
}
