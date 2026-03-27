package overview

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/entity"
	"management_backend/src/sync"
)

// GetResourceListHandler getResourceListHandler
type GetResourceListHandler struct {
}

// LoginVerify login verify
func (handler *GetResourceListHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *GetResourceListHandler) Handle(user *entity.User, ctx *gin.Context) {
	resourceViews := arraylist.New()
	for resourceName, typeValue := range sync.ResourceNameMap {
		resourceView := &ResourceView{
			ResourceName: resourceName,
			Type:         typeValue,
		}
		resourceViews.Add(resourceView)
	}

	common.ConvergeListResponse(ctx, resourceViews.Values(), int64(len(resourceViews.Values())), nil)
}
