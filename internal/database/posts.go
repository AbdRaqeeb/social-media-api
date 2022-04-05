package database

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Post
type Post struct {
	ID string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string `json:"userEmail"`
	Text string `json:"text"`
}

func (c Client) CreatePost(userEmail, text string) (Post, error) {
	post := Post{}

	data, err := c.readDB()
	if err != nil {
		return post, err
	}

	_, ok := data.Users[userEmail]
	if !ok {
		return post, errors.New("user not found")
	}

	id := uuid.New().String()

	post = Post{
		ID: id,
		CreatedAt: time.Now().UTC(),
		UserEmail: userEmail,
		Text: text,
	}

	data.Posts[id] = post

	err = c.updateDB(data)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func (c Client) GetPosts(userEmail string) ([]Post, error) {
	data, err := c.readDB()
	if err != nil {
		return []Post{}, err
	}

	var postSlices []Post

	for _, value := range data.Posts {
		if userEmail == value.UserEmail {
			postSlices = append(postSlices, value)
			continue
		}
	}

	return postSlices, nil
}

func (c Client) DeletePost(id string) error {
	data, err := c.readDB()
	if err != nil {
		return err
	}

	delete(data.Posts, id)

	err = c.updateDB(data)
	if err != nil {
		return err
	}

	return nil
}
