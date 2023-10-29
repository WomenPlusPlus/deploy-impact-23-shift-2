package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleCreateUser(t *testing.T) {
	server := &APIServer{} // Assuming your server instance initialization
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/your-endpoint", nil) // Adjust the HTTP method and endpoint as needed
	if err != nil {
		t.Error(err)
	}

	server.handleCreateUser(rr, req)

	result := rr.Result()

	if result.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", result.StatusCode)
	}
	defer result.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(result.Body).Decode(&response); err != nil {
		t.Error(err)
	}

	// Perform assertions on the response if needed
	// Example: Check a specific field in the response
	// if value, exists := response["field_name"]; exists {
	//     // Perform assertion
	// }

	// Example: Asserting an empty response (if expected)
	// if len(response) != 0 {
	//     t.Error("Expected an empty response")
	// }
}
