package chain_management

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/entity"
)

// GetBcResource get bc resource
type GetBcResource struct{}

// LoginVerify login verify
func (getBcResource *GetBcResource) LoginVerify() bool {
	return true
}

// Handle deal
func (getBcResource *GetBcResource) Handle(user *entity.User, ctx *gin.Context) {
	resourceInfos := make([]ResourceInfo, 0)
	for resourceType, resourceName := range ResourceName {
		resourceInfos = append(resourceInfos, ResourceInfo{
			ResourceName: resourceName,
			ResourceType: resourceType,
		})
	}
	common.ConvergeDataResponse(ctx, resourceInfos, nil)
}

// ResourceInfo resource info
type ResourceInfo struct {
	ResourceName string
	ResourceType int
}
