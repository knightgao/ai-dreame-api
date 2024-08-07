package doubao

import (
	"fmt"
	"github.com/knightgao/ai-dreame-api/relay/meta"
	"github.com/knightgao/ai-dreame-api/relay/relaymode"
)

func GetRequestURL(meta *meta.Meta) (string, error) {
	if meta.Mode == relaymode.ChatCompletions {
		return fmt.Sprintf("%s/api/v3/chat/completions", meta.BaseURL), nil
	}
	return "", fmt.Errorf("unsupported relay mode %d for doubao", meta.Mode)
}
