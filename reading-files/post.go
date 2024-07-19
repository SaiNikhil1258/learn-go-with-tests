package readingfiles

import (
	"io/fs"
	"testing/fstest"
)

type Post struct {
}

func NewPostsFromFS(filesystem fstest.MapFS) ([]Post, error) {
	dir, _ := fs.ReadDir(filesystem, ".")
	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}
	return posts, nil
}
