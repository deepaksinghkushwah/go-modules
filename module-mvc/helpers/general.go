package helpers

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	// below package is for mysql integration
	_ "github.com/go-sql-driver/mysql"

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
func GetDB() *sql.DB {
	db, err := sql.Open(dbType, dns)
	if err != nil {
		panic(err)
	}

	return db
}

// CloseDB this will close db
func CloseDB(db *sql.DB) {
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
