package vault

import (
	"os"
	"testing"
	// "net/http"
	"io/ioutil"
	// "net/http/httptest"
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
	result, err := encrypt_config("test_key", test_string);
	if err != nil { t.Error(err) }

	if string(result) != test_string {
		t.Fail()
	}
}

func TestEncryptHandler(t *testing.T) {

}
