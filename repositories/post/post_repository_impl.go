package repositories

import (
	"context"
	"errors"
	"fmt"
	"go-rest-prisma/helpers"
	"go-rest-prisma/models"
	"go-rest-prisma/prisma/db"
)

type PostRepositoryImpl struct {
	Db *db.PrismaClient
}

// Delete implements PostRepository.
func (p *PostRepositoryImpl) Delete(ctx context.Context, postId string) {
	result, err := p.Db.Post.FindUnique(db.Post.ID.Equals(postId)).Delete().Exec(ctx)
	helpers.ErrorPanic(err)
	fmt.Println("Rows affected: ", result)
}

// FindAll implements PostRepository.
func (p *PostRepositoryImpl) FindAll(ctx context.Context) []models.Post {
	allPost, err := p.Db.Post.FindMany().Exec(ctx)
	helpers.ErrorPanic(err)

	var posts []models.Post

	for _, post := range allPost {
		published, _ := post.Published()
		// description, _ := post.Decription()

		postData := models.Post{
			Id:          post.ID,
			Title:       post.Title,
			Published:   published,
			Description: post.Decription,
		}
		posts = append(posts, postData)
	}
	return posts
}

// FindById implements PostRepository.
func (p *PostRepositoryImpl) FindById(ctx context.Context, postId string) (models.Post, error) {
	post, err := p.Db.Post.FindUnique(db.Post.ID.Equals(postId)).Delete().Exec(ctx)
	helpers.ErrorPanic(err)

	published, _ := post.Published()
	postData := models.Post{
		Id:        post.ID,
		Title:     post.Title,
		Published: published,
	}

	if post != nil {
		return postData, nil
	} else {
		return postData, errors.New("Post id not found")
	}
}

// Save implements PostRepository.
func (p *PostRepositoryImpl) Save(ctx context.Context, post models.Post) {
	// panic("unimplemented")
	result, err := p.Db.Post.CreateOne(db.Post.Title.Set(post.Title), db.Post.Decription.Set(post.Description), db.Post.Published.Set(post.Published)).Exec(ctx)
	helpers.ErrorPanic(err)
	fmt.Println("Rows affected: ", result)
}

// Update implements PostRepository.
func (p *PostRepositoryImpl) Update(ctx context.Context, post models.Post) {
	// panic("unimplemented")
	result, err := p.Db.Post.FindMany(db.Post.ID.Equals(post.Id)).Update(db.Post.Decription.Set(post.Description), db.Post.Published.Set(post.Published)).Exec(ctx)
	helpers.ErrorPanic(err)
	fmt.Println("Rows affected: ", result)
}

func NewPostRepository(Db *db.PrismaClient) PostRepository {
	return &PostRepositoryImpl{Db: Db}
}
