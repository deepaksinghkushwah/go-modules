package helpers

import (
	"log"
	"module-mvc/domain"	
	"os"
)

//GetUserDetail return user details from db
func GetUserDetail(userID int) domain.User {
	db := GetDB()
	defer CloseDB(db)

	var user domain.User
	sql := "SELECT id, username, first_name, last_name, image FROM `user` WHERE id = ?"

	err := db.QueryRow(sql, userID).Scan(&user.ID, &user.Username,&user.FirstName,&user.LastName,&user.Image)
	if err != nil {
		log.Fatalln(err)
	}
	return user
}

//DeleteOldProfileImage delete profile image which is not required
func DeleteOldProfileImage(userID int) {
	user := GetUserDetail(userID)	
	
	if user.Image.String != "noimg.png" {
		err := os.Remove("static/images/" + user.Image.String)
		if err != nil {
			log.Println(err)			
		}		
	}	
}