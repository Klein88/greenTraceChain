package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
	"log"
	"net/http"
)

var sysMessage = `你是一个Golang程序员，只输出代码，不要输出除代码以外的任何其他话。以下是已知的接口可被调用：
type SDKInterface interface {
	// GetArgs get arg from transaction parameters
	// @return: 参数map
	GetArgs() map[string][]byte
	// GetStateByte get [key, field] from chain
	// @param key: 获取的参数名
	// @param field: 获取的参数名
	// @return1: 获取结果，格式为[]byte, nil表示不存在
	// @return2: 获取错误信息
	GetStateByte(key, field string) ([]byte, error)
	// PutStateByte put [key, field, value] to chain
	// @param1 key: 参数名
	// @param1 field: 参数名
	// @param2 value: 参数值，类型为[]byte
	// @return1: 上传参数错误信息
	PutStateByte(key, field string, value []byte) error
	// DelState delete [key, field] to chain
	// @param1 key: 删除的参数名
	// @param1 field: 删除的参数名
	// @return1：删除参数的错误信息
	DelState(key, field string) error
	// EmitEvent emit event, you can subscribe to the event using the SDK
	// @param1 topic: 合约事件的主题
	// @param2 data: 合约事件的数据
	EmitEvent(topic string, data []string)
	// GetTxTimeStamp 获得交易的时间戳
	GetTxTimeStamp() (string, error)
	// CallContract 跨合约调用另一个合约
	CallContract(contractName, method string, args map[string][]byte) protogo.Response
	// Sender 获得交易发送者的地址
	Sender() (string, error)
}
通过sdk.Instance可以使用上述接口的函数，另外还有sdk.Success(payload []byte) protogo.Response表示成功的返回，sdk.Error(msg string) protogo.Response表示失败的返回，sdk包下没有其他方法了。
我们就可以编写类似Fabric ChainCode的合约，只是合约格式有所不同，另外一定要注意：在代码中绝对不能出现time.Now()，随机数等不确定性的函数，获取时间用sdk.Instance.GetTxTimeStamp()函数。以下是一个最简单的存证数据的合约代码：

package main

import (
	"crypto/sha256"
	"encoding/hex"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sandbox"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// EvidenceContract 存证合约实现
type EvidenceContract struct {
}

// InitContract 初始化合约，必须实现的接口
func (e *EvidenceContract) InitContract() protogo.Response {
	return sdk.SuccessResponse
}

// UpgradeContract 升级合约，必须实现的接口
func (e *EvidenceContract) UpgradeContract() protogo.Response {
	return sdk.SuccessResponse
}

// InvokeContract 调用合约，必须实现的接口
func (e *EvidenceContract) InvokeContract(method string) (result protogo.Response) {
	// 记录异常结果日志
	defer func() {
		if result.Status != 0 {
			sdk.Instance.Warnf(result.Message)
		}
	}()

	switch method {
	case "PutEvidence":
		return e.PutEvidence()
	//其他case ...
	
	default:
		return sdk.Error("invalid method")
	}
}

func (e *EvidenceContract) PutEvidence() protogo.Response {
	params := sdk.Instance.GetArgs()
	// 获取参数
	data := params["data"]
	hash:=sha256.Sum256(data)
	err := sdk.Instance.PutStateByte(hex.EncodeToString( hash[:]), "data",data)
	if err != nil {
		return sdk.Error(err.Error())
	}
	// 返回OK
	return sdk.SuccessResponse
}

func main() {
	err := sandbox.Start(new(EvidenceContract))
	if err != nil {
		sdk.Instance.Errorf(err.Error())
	}
}
根据以上内容，请仅提供Go代码，不要包含任何Markdown标记，也不要有任何其他文字，只要代码，满足用户的要求：`

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})
	router.GET("/api", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Failed to upgrade connection:", err)
			return
		}
		defer conn.Close()

		// Get the "text" parameter from the query string
		inputText := c.Query("text")
		fmt.Println("Received text:", inputText)

		// 将这里的OPENAI_API_BASE和OPENAI_API_KEY替换为你的实际值
		llm, err := openai.NewChat(
			openai.WithEmbeddingModel("text-embedding-ada-002"),
			openai.WithAPIVersion("2023-07-01-preview"),
			openai.WithAPIType(openai.APITypeAzure),
			openai.WithBaseURL("url"), // 你的API基础URL
			openai.WithToken("api"),      // 你的API密钥
			openai.WithModel("ChainGPT"))
		if err != nil {
			log.Fatal(err)
		}
		var humanMessage = c.Query("text")

		ctx := context.Background()

		_, err = llm.Call(ctx, []schema.ChatMessage{
			schema.SystemChatMessage{Content: sysMessage},
			schema.HumanChatMessage{Content: humanMessage},
		}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			err = conn.WriteMessage(websocket.TextMessage, chunk)
			if err != nil {
				fmt.Println("Failed to write message:", err)
				return err
			}
			return nil
		}),
		)
		if err != nil {
			log.Fatal(err)
		}
	})

	router.Run(":8085")
}
