
package model

import "time"

type User struct {
    Id         int64     `json:"id"`
    Username   string    `json:"username"`
    Email      string    `json:"email"`
    Password   string    `json:"-"`
    Created_at time.Time `json:"created_at"`
    Updated_at time.Time `json:"updated_at"`
}

type EmailFetch struct {
    Id       int64  `json:"id"`
    Email    string `json:"email"`
    Password string `json:"-"`
}
