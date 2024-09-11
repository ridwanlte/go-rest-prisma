package service

import (
	"context"
	request "go-rest-prisma/data/request/post"
	response "go-rest-prisma/data/response/post"
	"go-rest-prisma/helpers"
	"go-rest-prisma/models"
	repositories "go-rest-prisma/repositories/post"
)

// Define the PostServiceImpl type here
type PostServiceImpl struct {
	PostRepository repositories.PostRepository
}

// Create implements PostService.
func (p *PostServiceImpl) Create(ctx context.Context, request request.PostCreateRequest) {
	// panic("unimplemented")
	postData := models.Post{
		Title:       request.Title,
		Published:   request.Published,
		Description: request.Description,
	}

	p.PostRepository.Save(ctx, postData)
}

// Delete implements PostService.
func (p *PostServiceImpl) Delete(ctx context.Context, postId string) {
	// panic("unimplemented")
	post, err := p.PostRepository.FindById(ctx, postId)
	helpers.ErrorPanic(err)
	p.PostRepository.Delete(ctx, post.Id)
}

// FindAll implements PostService.
func (p *PostServiceImpl) FindAll(ctx context.Context) []response.PostResponse {
	// panic("unimplemented")
	posts := p.PostRepository.FindAll(ctx)

	var rsp []response.PostResponse

	for _, value := range posts {
		post := response.PostResponse{
			Id:          value.Id,
			Title:       value.Title,
			Published:   value.Published,
			Description: value.Description,
		}
		rsp = append(rsp, post)
	}

	return rsp
}

// FindById implements PostService.
func (p *PostServiceImpl) FindById(ctx context.Context, postId string) response.PostResponse {
	// panic("unimplemented")
	post, err := p.PostRepository.FindById(ctx, postId)
	helpers.ErrorPanic(err)

	postResponse := response.PostResponse{
		Id:          post.Id,
		Title:       post.Title,
		Published:   post.Published,
		Description: post.Description,
	}

	return postResponse
}

// Update implements PostService.
func (p *PostServiceImpl) Update(ctx context.Context, request request.PostUpdateRequest) {
	// panic("unimplemented")

	postData := models.Post{
		Id:          request.Id,
		Title:       request.Title,
		Published:   request.Published,
		Description: request.Description,
	}
	p.PostRepository.Update(ctx, postData)

}

func NewPostService(postRepository repositories.PostRepository) PostService {
	return &PostServiceImpl{PostRepository: postRepository}
}
