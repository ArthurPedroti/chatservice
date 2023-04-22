package chatcompletionstream

import (
	"github.com/ArthurPedroti/chatservice/internal/domain/gateway"
	openai "github.com/sashabaranov/go-openai"
)

type ChatCompletionUseCase struct {
	ChatGateway  gateway.ChatGateway
	OpenAiClient *openai.Client
}

func NewChatCompletionUseCase(chatGateway gateway.ChatGateway, openaiClient *openai.Client) *ChatCompletionUseCase {
	return &ChatCompletionUseCase{
		ChatGateway:  chatGateway,
		OpenAiClient: openaiClient,
	}
}
