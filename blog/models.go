package blog

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string
	Password string
	Email    string
	IsAdmin  bool `bson:"isAdmin"`
}

type Post struct {
	ID          bson.ObjectId  `bson:"_id,omitempty" json:"id"`
	Title       string         `json:"title"`
	Content     string         `json:"content"`
	AuthorID    bson.ObjectId  `bson:"authorId" json:"authorId"`
	CreatedAt   time.Time      `bson:"createdAt" json:"createdAt"`
	UpdateAt    *time.Time     `bson:"updatedAt" json:"updatedAt"`
	UpdatedByID *bson.ObjectId `bson:"updatedById" json:"updatedById"`
}

type Comment struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Content  string
	Email    string
	Approved bool
}
