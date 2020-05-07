package controllers

import (
	"fmt"
	"log"
	"module-mvc/domain"
	"module-mvc/helpers"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Me function display after logged in
func Me(c *gin.Context) {
	session := sessions.Default(c)
	flash := session.Flashes()
	session.Save()
	currentPath := c.Request.URL.Path

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"Title":       "My Status",
		"Content":     nil,
		"IsLoggedIn":  session.Get(loggedInKey),
		"flashes":     flash,
		"currentPath": currentPath,
	})

}

// UpdateProfile show profile update section
func UpdateProfile(c *gin.Context) {
	session := sessions.Default(c)
	flash := session.Flashes()
	session.Save()
	currentPath := c.Request.URL.Path

	db := helpers.GetDB()
	defer helpers.CloseDB(db)

	userID := session.Get("userID")
	row := db.QueryRow("SELECT id, first_name, last_name,  image FROM `user` WHERE id = ?", userID)

	var user domain.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Image)
	if err != nil {
		panic(err)
	}

	c.HTML(http.StatusOK, "profile-form.html", gin.H{
		"Title":       "Update Profile",
		"Content":     nil,
		"IsLoggedIn":  session.Get(loggedInKey),
		"flashes":     flash,
		"currentPath": currentPath,
		"user":        user,
	})
}

// UpdateProfileHandler update profile image and fields
func UpdateProfileHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID").(int)
	db := helpers.GetDB()
	defer helpers.CloseDB(db)

	firstName := c.PostForm("firstname")
	lastName := c.PostForm("lastname")

	if firstName == "" || lastName == "" {
		session.AddFlash("Firstname and lastname are required fields")
		session.Save()
		c.Redirect(http.StatusSeeOther, "/member/update-profile")
	}

	sql := "UPDATE `user` SET first_name = ?, last_name = ?"
	file, err := c.FormFile("image")
	var filename string
	if err != nil {
		filename = helpers.GetUserDetail(userID).Image.String
	} else {
		fmt.Println("--------------------------------")
		fmt.Println(file.Header.Get("Content-Type"))
		fmt.Println("--------------------------------")
		isFileOk := ValidateImageFile(file.Header.Get("Content-Type"))
		if isFileOk == true {
			uuid := uuid.NewV4()
			filename = uuid.String() + filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, "static/images/"+filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
			helpers.DeleteOldProfileImage(userID)
		} else {
			session.AddFlash("Image file is not valid filetype")
			session.Save()
			filename = helpers.GetUserDetail(userID).Image.String
		}
	}
	sql += ", image = ? WHERE id = ?"

	_, err = db.Exec(sql, firstName, lastName, filename, userID)
	if err != nil {
		log.Fatalln(err)
	}

	session.AddFlash("Profile updated")
	session.Save()

	c.Redirect(http.StatusSeeOther, "/member/update-profile")

}

//ValidateImageFile check if fileis valid
func ValidateImageFile(fileType string) bool {
	var returnVal bool
	switch fileType {
	case "image/jpeg":
		returnVal = true
	case "image/jpg":
		returnVal = true
	case "image/png":
		returnVal = true
	case "image/gif":
		returnVal = true
	default:
		returnVal = false
	}

	return returnVal

}
