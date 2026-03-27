package DAP

import (
	contractUtil "ChainMaker_Backend_demo/src/ctrl/contract/Utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type GetPNGImageById struct {
}

func (GetPNGImageById *GetPNGImageById) Handle(ctx *gin.Context) {
	var UserInfo struct {
		Id        int    `json:"Id"`
		CompanyId string `json:"CompanyId"`
	}

	if err := ctx.ShouldBindJSON(&UserInfo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	path := contractUtil.FileURI + "/" + UserInfo.CompanyId + "/" + strconv.Itoa(UserInfo.Id) + "/PDFFile" + "/" + "PGE_" + strconv.Itoa(UserInfo.Id) + ".pdf"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// 获取文件名
	fileName := filepath.Base(path)
	// 设置响应头
	ctx.Header("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	ctx.Header("Content-Type", "application/octet-stream")

	// 将文件内容发送给客户端
	ctx.Status(http.StatusOK)
	ctx.Stream(func(w io.Writer) bool {
		buffer := make([]byte, 1024)
		for {
			n, err := file.Read(buffer)
			if err != nil && err == io.EOF {
				return false
			}
			w.Write(buffer[:n])
		}
	})
}
