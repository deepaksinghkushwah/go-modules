package general

import (
	"log"
	"mvc-gorm/models"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	// below package is for mysql integration
	//_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var userkey, dns, dbType string

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalln("Error at loading env file")
	}
	userkey = os.Getenv("userkey") // session userkey ie: user
	dns = os.Getenv("dns")         // mysql connection string
	dbType = os.Getenv("dbType")   // mysql
}

// GetDB return db object
func GetDB() *gorm.DB {
	db, err := gorm.Open(dbType, dns)
	checkError(err)
	db.AutoMigrate(&models.User{})
	return db
}

// CloseDB this will close db
func CloseDB(db *gorm.DB) {
	db.Close()
}

// Authorized middleware
func Authorized(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		session.AddFlash("Unauthorized Access")
		session.Save()
		c.Redirect(http.StatusSeeOther, "/site/login")
	}
	c.Next()
}

//AuthAdmin authorized admin user
func AuthAdmin(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	roleID := session.Get("roleID").(uint)
	if user == nil || roleID != 1 {
		session.AddFlash("Unauthorized Access")
		session.Save()
		c.Redirect(http.StatusSeeOther, "/")
	}
	c.Next()
}

//HashPassword generate hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash check hashed password with string password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//RunMigration will run first migrations
func RunMigration() {
	db := GetDB()

	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.User{})

	var role models.Role
	db.First(&role)
	if role.ID <= 0 {
		role1 := models.Role{Title: "Admin"}
		db.Create(&role1)

		role2 := models.Role{Title: "Registered"}
		db.Create(&role2)
	}
}
