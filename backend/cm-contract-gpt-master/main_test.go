package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/schema"
)

func TestLLM(t *testing.T) {
	llm, err := openai.New(
		openai.WithEmbeddingModel("text-embedding-ada-002"),
		openai.WithAPIVersion("2023-07-01-preview"),
		openai.WithAPIType(openai.APITypeAzure),
		openai.WithBaseURL(os.Getenv("OPENAI_API_BASE")),
		openai.WithToken(os.Getenv("OPENAI_API_KEY")),
		openai.WithModel("gpt-35-turbo"))
	if err != nil {
		log.Fatal(err)
	}
	var humanMessage = "一个最简单的存证合约"

	ctx := context.Background()

	c := chains.NewConversation(llm, memory.NewConversationTokenBuffer(llm, 2000))

	result, err := chains.Run(ctx, c, []schema.ChatMessage{
		schema.SystemChatMessage{Content: sysMessage},
		schema.HumanChatMessage{Content: humanMessage},
	}, chains.WithStreamingFunc(
		func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n----------------对话结束！")
	fmt.Println(result)
}
