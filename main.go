package main

import (
	"github.com/gin-gonic/gin"

	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

const ollamaURL = "http://0.0.0.0:11434/api/chat"

var MODEL = "qwen_7b"

//var MODEL = "qwen_0_5b"

type inputPrompt struct {
	Prompt string `json:"prompt"`
}

type prompt struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func main() {
	fmt.Println("start gin server")
	r := gin.Default()

	r.POST("/initialize", func(c *gin.Context) {
		startOllamaModel()
		c.String(200, "start model finished")
	})

	r.GET("/ping", func(c *gin.Context) {
		input := inputPrompt{
			Prompt: "who are you",
		}
		err := doChat(input)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		c.JSON(200, "inference finished")
	})

	fmt.Println("start ollama init")
	go initOllama()
	time.Sleep(10 * time.Second)
	fmt.Println("finish ollama init")

	r.Run("0.0.0.0:8000")
}

func initOllama() {
	// execute cmd
	cmd := exec.Command("ollama", "serve")
	// get cmd response and print
	out, err := cmd.Output()
	fmt.Println("output get finished")
	if err != nil {
		panic(err)
	}
	println(string(out))
}

func startOllamaModel() {
	// execute cmd and get output and print
	args := []string{"create", MODEL, "-f", fmt.Sprintf("./ollama_model/%s_modelfile", MODEL)}
	cmd := exec.Command("ollama", args...)
	// get cmd response and print
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("ollama", "run", MODEL)
	out, err = cmd.Output()
	if err != nil {
		fmt.Println("err is", err)
		panic(err)
	}

	println(string(out))
}

func doChat(input inputPrompt) error {
	fmt.Println("input: ", input.Prompt)

	ollamaInput := prompt{
		Model: MODEL,
		Messages: []message{
			{
				Role:    "user",
				Content: input.Prompt,
			},
		},
		Stream: false,
	}

	jsonData, err := json.Marshal(ollamaInput)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return err
	}

	resp, err := http.Post(ollamaURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error making POST request: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return err
	}

	fmt.Printf("Status code: %d\n", resp.StatusCode)
	fmt.Println("Body:", string(body))
	return nil
}
