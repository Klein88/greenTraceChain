package service

import (
	"chainmaker_web/src/dao/dbhandle"
	"chainmaker_web/src/entity"
	"chainmaker_web/src/sync"

	"encoding/json"
	"errors"
	"strings"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"
)

// GetNFTDetailHandler deal nft
type GetNFTDetailHandler struct {
}

// Handle deal
func (getNFTDetailHandler *GetNFTDetailHandler) Handle(ctx *gin.Context) {
	params := BindGetNFTDetailHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	tranfer, err := dbhandle.GetTransferDetail(params.TokenId, params.ChainId)
	if err != nil && !strings.Contains(err.Error(), "record not found") {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}
	// 获取metedata
	data := sync.GetTokenMetaData(params.ChainId, params.ContractName, params.TokenId)
	metedata := Metedata{}
	err = json.Unmarshal(data, &metedata)
	if err != nil {
		ConvergeHandleFailureResponse(ctx, errors.New("合约参数未遵循长安链合约标准协议，解析失败。"))
		return
	}
	owner := sync.GetOwnerByTokenId(params.ChainId, params.ContractName, params.TokenId)

	view := &entity.NFTDetailView{
		CreateTime: tranfer.BlockTime,
		ChainId:    params.ChainId,
		TokenId:    params.TokenId,
		Name:       metedata.Name,
		Author:     metedata.Author,
		Url:        metedata.Image,
		Hash:       metedata.SeriesHash,
		Describe:   metedata.Description,
		Owner:      owner,
	}

	ConvergeDataResponse(ctx, view, nil)
}

// Metedata json解析出来的对象
type Metedata struct {
	Name        string `json:"name,omitempty"`
	Author      string `json:"author,omitempty"`
	OrgName     string `json:"orgName"`
	Image       string `json:"image,omitempty"`
	Description string `json:"description,omitempty"`
	TokenId     string `json:"tokenId,omitempty"`
	SeriesHash  string `json:"seriesHash"`
}

// GetTransferList get
type GetTransferList struct {
}

// Handle deal
func (getTransferList *GetTransferList) Handle(ctx *gin.Context) {
	params := BindGetTransferListHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	tranfers, totalCount, err := dbhandle.GetTransferList(params.Offset, params.Limit, params.ChainId,
		params.ContractName, params.TokenId)
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}

	views := arraylist.New()
	for _, transfer := range tranfers {
		transferView := &entity.TransferListView{
			Id:             transfer.Id,
			TxId:           transfer.TxId,
			BlockTime:      transfer.BlockTime,
			ContractMethod: transfer.ContractMethod,
			From:           transfer.From,
			To:             transfer.To,
			TokenId:        transfer.TokenId,
			Status:         transfer.TxStatusCode,
			Amount:         transfer.Amount,
		}
		views.Add(transferView)
	}

	ConvergeListResponse(ctx, views.Values(), totalCount, nil)
}
