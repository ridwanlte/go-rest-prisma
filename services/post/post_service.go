package service

import (
	"context"
	request "go-rest-prisma/data/request/post"
	response "go-rest-prisma/data/response/post"
)

type PostService interface {
	Create(ctx context.Context, request request.PostCreateRequest)
	Update(ctx context.Context, request request.PostUpdateRequest)
	Delete(ctx context.Context, postId string)
	FindAll(ctx context.Context) []response.PostResponse
	FindById(ctx context.Context, postId string) response.PostResponse
}
