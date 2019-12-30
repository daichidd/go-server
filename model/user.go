package model

import "time"

type Users []*User

type User struct {
	ID        uint64     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
