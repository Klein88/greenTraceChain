package EP

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InformationDecoding struct {
}

func (InformationDecoding *InformationDecoding) Handle(ctx *gin.Context) {
	var UserInfo struct {
		Message   string `json:"Message"`
		PublicKey string `json:"PublicKey"`
	}

	if err := ctx.ShouldBindJSON(&UserInfo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("\n 需要解密的信息为 : %s :", UserInfo.Message)
	fmt.Printf("\n 解密使用的公钥 : %s :", UserInfo.PublicKey)
	// 使用base64.StdEncoding进行解码
	decodedBytes, err := base64.StdEncoding.DecodeString(UserInfo.Message)
	if err != nil {
		fmt.Println("解码错误:", err)
		return
	}

	// 打印解码后的字节
	fmt.Println("解码后的字节:", decodedBytes)
	// 如果你想要解码后的字符串（假设解码后的字节是UTF-8编码的文本）
	decodedString := string(decodedBytes)
	fmt.Println("解码后的字符串:", decodedString)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    decodedString,
	})
}
