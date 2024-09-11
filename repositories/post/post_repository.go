package repositories

import (
	"context"
	"go-rest-prisma/models"
)

type PostRepository interface {
	Save(ctx context.Context, post models.Post)
	Update(ctx context.Context, post models.Post)
	Delete(ctx context.Context, postId string)
	FindById(ctx context.Context, postId string) (models.Post, error)
	FindAll(ctx context.Context) []models.Post
}