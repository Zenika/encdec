package executor

//func main() {
//	// Set the variables
//	username := os.Args[1]
//	password := os.Args[2]
//	portainerHost := os.Args[3]
//	environmentName := os.Args[4]
//	filePath := os.Args[5]
//
//	// Obtain an authentication token
//	token, err := getAuthToken(username, password, portainerHost)
//	if err != nil {
//		log.Fatalf("Failed to obtain authentication token: %v", err)
//	}
//
//	// Get the environment ID
//	environmentID, err := getEnvironmentID(token, portainerHost, environmentName)
//	if err != nil {
//		log.Fatalf("Failed to get environment ID: %v", err)
//	}
//
//	// Import the Docker image
//	err = importDockerImage(token, portainerHost, environmentID, filePath)
//	if err != nil {
//		log.Fatalf("Failed to import Docker image: %v", err)
//	}
//
//	fmt.Println("Image imported successfully!")
//}
//
//func getAuthToken(username, password, portainerHost string) (string, error) {
//	requestBody := fmt.Sprintf(`{"Username": "%s", "Password": "%s"}`, username, password)
//	resp, err := http.Post(portainerHost+"/api/auth", "application/json", bytes.NewBufferString(requestBody))
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK {
//		return "", fmt.Errorf("failed to obtain authentication token: %s", resp.Status)
//	}
//
//	var tokenResp struct {
//		JWT string `json:"jwt"`
//	}
//	err = decodeJSON(resp.Body, &tokenResp)
//	if err != nil {
//		return "", err
//	}
//
//	return tokenResp.JWT, nil
//}
//
//func getEnvironmentID(token, portainerHost, environmentName string) (string, error) {
//	req, err := http.NewRequest(http.MethodGet, portainerHost+"/api/endpoints", nil)
//	if err != nil {
//		return "", err
//	}
//	req.Header.Set("Authorization", "Bearer "+token)
//
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK {
//		return "", fmt.Errorf("failed to get environment ID: %s", resp.Status)
//	}
//
//	var endpoints []struct {
//		Name string `json:"Name"`
//		ID   string `json:"Id"`
//	}
//	err = decodeJSON(resp.Body, &endpoints)
//	if err != nil {
//		return "", err
//	}
//
//	for _, endpoint := range endpoints {
//		if endpoint.Name == environmentName {
//			return endpoint.ID, nil
//		}
//	}
//
//	return "", fmt.Errorf("environment not found: %s", environmentName)
//}
//
//func importDockerImage(token, portainerHost, environmentID, filePath string) error {
//	file, err := os.Open(filePath)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//
//	body := &bytes.Buffer{}
//	writer := multipart.NewWriter(body)
//	part, err := writer.CreateFormFile("file", filePath)
//	if err != nil {
//		return err
//	}
//
//	_, err = io.Copy(part, file)
//	if err != nil {
//		return err
//	}
//
//	err = writer.Close()
//	if err != nil {
//		return err
//	}
//
//	req, err := http.NewRequest(http.MethodPut, portainerHost+"/api/endpoints/"+environmentID+"/docker/images/load", body)
//	if err != nil {
//		return err
//	}
//	req.Header.Set("Authorization", "Bearer "+token)
//	req.Header.Set("Content-Type", writer.FormDataContentType())
//
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK {
//		return fmt.Errorf("failed to import Docker image: %s", resp.Status)
//	}
//
//	return nil
//}
//
//func decodeJSON(r io.Reader, v interface{}) error {
//	return json.NewDecoder(r).Decode(v)
//}
//
