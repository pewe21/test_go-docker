package schema

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
