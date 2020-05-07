package domain

import "database/sql"

var magicTable = map[string]string{
	"\xff\xd8\xff":      "image/jpeg",
	"\x89PNG\r\n\x1a\n": "image/png",
	"GIF87a":            "image/gif",
	"GIF89a":            "image/gif",
}

//User struct for users
type User struct {
	ID        int            `json:"id"`
	Email     string         `json:"email"`
	Username  string         `json:"username"`
	Password  string         `json:"-"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Image     sql.NullString `json:"image"`
}
