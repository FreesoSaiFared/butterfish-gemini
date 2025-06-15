// Simple GPT wrapper calling local Python script
package butterfish

import (
    "bytes"
    "context"
    "fmt"
    "io"
    "os/exec"
    "strings"

    "github.com/bakks/butterfish/util"
)

// GPT implements the LLM interface by running a Python script.
type GPT struct{}

// NewGPT returns a new GPT instance. Token and baseURL are ignored.
func NewGPT(token, baseURL string) *GPT {
    return &GPT{}
}

// runScript executes gpt4-script.py with the given prompt and returns stdout.
func runScript(prompt string) (string, error) {
    args := append([]string{"gpt4-script.py"}, strings.Fields(prompt)...)
    cmd := exec.Command("python3", args...)
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
        return "", fmt.Errorf("%v: %s", err, stderr.String())
    }
    return strings.TrimSpace(out.String()), nil
}

func (g *GPT) Completion(request *util.CompletionRequest) (*util.CompletionResponse, error) {
    result, err := runScript(request.Prompt)
    if err != nil {
        return nil, err
    }
    return &util.CompletionResponse{Completion: result}, nil
}

func (g *GPT) CompletionStream(request *util.CompletionRequest, writer io.Writer) (*util.CompletionResponse, error) {
    resp, err := g.Completion(request)
    if err != nil {
        return nil, err
    }
    if writer != nil {
        writer.Write([]byte(resp.Completion))
    }
    return resp, nil
}

func (g *GPT) Embeddings(ctx context.Context, input []string, verbose bool) ([][]float32, error) {
    return nil, fmt.Errorf("embeddings not supported")
}


func ShellHistoryTypeToRole(t int) string {
    switch t {
    case historyTypeLLMOutput:
        return "assistant"
    case historyTypeFunctionOutput:
        return "function"
    case historyTypeToolOutput:
        return "tool"
    default:
        return "user"
    }
}

