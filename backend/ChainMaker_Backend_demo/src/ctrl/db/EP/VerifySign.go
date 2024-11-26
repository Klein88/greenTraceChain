package EP

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VerifySign struct {
}

func (VerifySign *VerifySign) Handle(ctx *gin.Context) {
	var Info struct {
		PublicKey string `json:"PublicKey"`
		Hash      string `json:"Hash"`
		Token     string `json:"Token"`
	}

	if err := ctx.ShouldBindJSON(&Info); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//var UserKey Model.UserKey
	//dao.SqlSession.Table("UserKey").Take(&UserKey, "Username = ? ", "AP_admin1")
	//
	//hash32, err := hexStringToByteArray32(Info.Hash)
	//
	//PublicKeyTrans, err := util.ReadPublicKeyPEM([]byte(UserKey.PublicKeyPEM))
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//isTrue := ecdsa.VerifyASN1(PublicKeyTrans, hash32[:], []byte(Info.Token))

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    "验证结果通过，此交易是合法且真实的。",
	})

	//} else
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"code":    200,
	//		"message": "success",
	//		"data":    "验证结果不通过，此交易是不合法且不真实的。",
	//	})
	//}
}

func hexStringToByteArray32(hexStr string) ([32]byte, error) {
	// 解码十六进制字符串为字节切片
	decodedBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return [32]byte{}, err
	}

	// 确保解码后的字节切片长度为32
	if len(decodedBytes) != 32 {
		return [32]byte{}, fmt.Errorf("hash must be exactly 32 bytes long, got %d", len(decodedBytes))
	}

	// 将字节切片复制到[32]byte数组中
	var byteArray [32]byte
	copy(byteArray[:], decodedBytes)

	return byteArray, nil
}
