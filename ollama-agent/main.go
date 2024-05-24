package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ModelfileDir   = "/work/ollama_model/%s_modelfile"
	Ollama         = "ollama"
	OllamaEndpoint = "http://0.0.0.0:11434"

	Initialize = "/initialize"
	Path       = "path"
)

func initOllama() (*exec.Cmd, error) {
	cmd := exec.Command(Ollama, "serve")
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	if err := waitForOllama(10); err != nil {
		cmd.Process.Kill()
		return nil, err
	}
	return cmd, nil
}

// waitForOllama waits for the ollama service to be ready.
func waitForOllama(maxAttempts int) error {
	for i := 0; i < maxAttempts; i++ {
		log.Printf("Attempt %d to check if ollama is alive...\n", i+1)
		resp, err := http.Get(OllamaEndpoint)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			return nil
		}
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("failed to start ollama")
}

func startOllamaService(model string) error {
	// initialize ollama
	ollamaCmd, err := initOllama()
	if err != nil {
		return err
	}

	// create model from an existing modelfile
	log.Printf("Creating model '%s'...\n", model)
	createArgs := []string{"create", model, "-f", fmt.Sprintf(ModelfileDir, model)}
	if err = executeCommand(Ollama, createArgs...); err != nil {
		ollamaCmd.Process.Kill()
		return err
	}

	// run model to provide LLM service
	log.Printf("Running model '%s'...\n", model)
	runArgs := []string{"run", model}
	if err = executeCommand(Ollama, runArgs...); err != nil {
		ollamaCmd.Process.Kill()
		return err
	}
	return nil
}

func executeCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing command: %v\nOutput: %s\n", err, output)
		return err
	}
	return nil
}

func reverseProxy(c *gin.Context) {
	if c.Param(Path) == Initialize && c.Request.Method == http.MethodPost {
		modelName := os.Getenv("MODEL")
		const defaultModel = "qwen_0_5b"
		if modelName == "" {
			modelName = defaultModel
		}
		log.Printf("Starting ollama service with %s\n", modelName)
		err := startOllamaService(modelName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to start ollama service with %s due to %v", modelName, err)})
		}
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Successfully started ollama service with %s", modelName)})
		return
	}

	target, err := url.Parse(OllamaEndpoint)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = target.Host
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = c.Param(Path)
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {
	r := gin.Default()
	r.Any("/*"+Path, reverseProxy)

	r.Run("0.0.0.0:8000")
}
