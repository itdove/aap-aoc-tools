package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Retrieve the project_id from the credentials file specified by
// the GOOGLE_APPLICATION_CREDENTIALS environment variable
func GetProjectID() (string, error) {
	credsFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if len(credsFile) == 0 {
		return "", fmt.Errorf("GOOGLE_APPLICATION_CREDENTIALS not set")
	}

	b, err := ioutil.ReadFile(credsFile)
	if err != nil {
		return "", err
	}

	creds := make(map[string]string, 0)

	err = json.Unmarshal(b, &creds)
	if err != nil {
		return "", err
	}

	if project, ok := creds["project_id"]; ok {
		return project, nil
	}

	return "", fmt.Errorf("project_id not found in file %s", credsFile)
}
