package vault

import (
	"os"
	"bytes"
	"testing"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"net/http/httptest"
)

func TestCreateKeyfile(t *testing.T) {
	file, err := create_keyfile("test_key");
	defer os.Remove(file.Name())

	if err != nil { t.Error("Error creating key file") }

	bytes, err := ioutil.ReadFile(file.Name());

	if err != nil { t.Errorf("Error %v", err) }

	if string(bytes) != "test_key" {
		t.Fail()
	}
}

// NOTE: What this will really test is whether the stdinpipe is used correctly
// as ansible is not installed in the development environment
func TestEncryptConfig(t *testing.T) {
	test_string := "{'key':'key','another_key':'another_key','array':[1,2,3],'nested':{'a':1.2,'b':false}}"
	result, err := encrypt_config(&Payload { Key: "test_key", Body: test_string });
	if err != nil { t.Error(err) }

	if string(result) != test_string { t.Fail() }
}

func TestEncryptHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(encrypt_handler))
	defer server.Close()

	endpoint := server.URL + "/v0/api/vault"

	test_body := "{'key':'key','another_key':'another_key','array':[1,2,3]}"

	t.Run("200", func(t *testing.T) {
		payload, err := json.Marshal(Payload {
			Key: "test_key",
			Body: test_body,
		})

		if err != nil { t.Error(err) }

		res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(payload))
		if err != nil { t.Error(err) }

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil { t.Error(err) }

		if res.StatusCode != 200 { t.Fail() }
		if string(body) != test_body { t.Fail() }
	})

	t.Run("400 - Missing Key", func(t *testing.T) {
		payload, err := json.Marshal(Payload {
			Body: test_body,
		})
		if err != nil { t.Error(err) }

		res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(payload))
		if err != nil { t.Error(err) }

		if res.StatusCode != 400 { t.Fail() }
	})

	t.Run("400 - Missing Body", func(t *testing.T) {

		payload, err := json.Marshal(Payload {
			Key: "test_key",
		})
		if err != nil { t.Error(err) }

		res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(payload))
		if err != nil { t.Error(err) }

		if res.StatusCode != 400 { t.Fail() }
	})

	t.Run("400 - Payload too large", func(t *testing.T) {
		payload := make([]byte, 1048577, 1048577)

		res, err := http.Post(endpoint, "application/json", bytes.NewReader(payload))
		if err != nil { t.Error(err) }

		if res.StatusCode != 400 { t.Fail() }
	})

	t.Run("400 - Empty Body", func(t *testing.T) {
		res, err := http.Post(endpoint, "application/json", bytes.NewReader(make([]byte, 0)))
		if err != nil { t.Error(err) }

		if res.StatusCode != 400 { t.Fail() }
	})
}