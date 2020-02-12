package helpers

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"shop/domain"
	"time"

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
	token := c.Query("token")
	db := GetDB()
	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "You must provide token",
			"status": "0",
		})
		c.Abort()
		return
	}
	var u domain.User

	currentDate := time.Now().Local()
	//log.Println(currentDate.Format("2006-01-02"))
	err := db.QueryRow("SELECT id, username, email, first_name, last_name FROM `user` WHERE access_token = ? AND access_token_expire_date >= ?", token, currentDate.Format("2006-01-02")).Scan(&u.ID, &u.Username, &u.Email, &u.FirstName, &u.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{
				"msg":    "Invalid token, user not found",
				"status": "0",
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":    err.Error(),
			"status": "0",
		})
		c.Abort()
		return

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
