package model

import "time"

type Payment struct {
	Id               int       `db:"id" uri:"id" example:"1"`
	Name             string    `db:"name" json:"name" example:"John Doe"`
	Password         string    `db:"password" json:"password" swaggerignore:"true"`
	Email            string    `db:"email" json:"email" example:"johndoe@mail.com"`
	CreatedAt        time.Time `db:"created_at" json:"created_at" example:"2009-11-10 23:00:00 +0000 UTC m=+0.000000001"`
}