package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Candidate struct {
	Content struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"content"`
}

type APIResponse struct {
	Candidates []Candidate `json:"candidates"`
}

func GenerateText(prompt string) (string, error) {
	Gem_Token := os.Getenv("GEMINI_API_TOKEN")
	fmt.Println("Preparing request...")
	data := map[string]interface{}{
		"contents": []interface{}{
			map[string]interface{}{
				"parts": []map[string]string{
					{
						"text": prompt,
					},
				},
			},
		},
	}

	postBody, _ := json.Marshal(data)
	responseBody := bytes.NewBuffer(postBody)
	fmt.Println("Sending request...")
	resp, err := http.Post(fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent?key=%s", Gem_Token), "application/json", responseBody)
	if err != nil {
		return "", fmt.Errorf("An Error Occurred: %v", err)
	}
	defer resp.Body.Close()
	fmt.Println("Reading response...")

	// Read and parse the response body

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body) // Read the response body for error details
		return "", fmt.Errorf("Request failed with status: %s, response: %s", resp.Status, string(bodyBytes))
	}
	body, _ := io.ReadAll(resp.Body)
	var apiResponse APIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return "", err
	}

	// Extract and return the text from the first candidate
	if len(apiResponse.Candidates) > 0 && len(apiResponse.Candidates[0].Content.Parts) > 0 {
		return apiResponse.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", fmt.Errorf("No valid response received")
}
