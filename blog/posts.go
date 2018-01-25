package blog

import (
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetPosts(db *mgo.Database, skip int, limit int) ([]*Post, error) {
	query := db.C("posts").Find(bson.M{}).Skip(skip).Limit(limit)
	result := make([]*Post, limit)
	if err := query.All(&result); err != nil {
		return nil, errors.Wrap(err, "failed to find the posts in the database")
	}
	return result, nil
}

func CreatePost(db *mgo.Database, post *Post) (*Post, error) {
	if err := db.C("posts").Insert(post); err != nil {
		return nil, errors.Wrap(err, "failed to insert the post in the database")
	}
	return post, nil
}
