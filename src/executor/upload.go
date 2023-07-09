package executor

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func uploadFile(url string, filePath string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a new form file field
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return err
	}

	// Copy the file content to the form field
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		return err
	}

	// Create a new HTTP request
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	// Set the content type header
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Make the HTTP request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("upload failed with status code %d", response.StatusCode)
	}

	fmt.Println("File uploaded successfully!")
	return nil
}

func Upload() {
	url := "https://example.com/upload"
	filePath := "/path/to/file.txt"

	err := uploadFile(url, filePath)
	if err != nil {
		fmt.Printf("Error uploading file: %s\n", err.Error())
	}
}
