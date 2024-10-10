package service

import (
	"github.com/CyberneticsMan/GoBackendLearning/data/request"
	"github.com/CyberneticsMan/GoBackendLearning/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
