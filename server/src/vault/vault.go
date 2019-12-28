package vault

import (
	"io"
	"os"
	"os/exec"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/go-chi/chi"
)

type Payload struct {
	Key string `json:key`
	Body string `json:body`
}

func create_keyfile(key string) (*os.File, error) {
	file, err := ioutil.TempFile("", "key")

	defer file.Close()
	defer file.Sync()

	if err != nil { return nil, err }

	file.WriteString(key)

	return file, nil
}

func encrypt_config(key string, config string) ([]byte, error) {
	// Create keyfile with provided password
	// Delete keyfile after encryption done
	file, err := create_keyfile(key);
	if err != nil { return nil, err }

	defer os.Remove(file.Name())

	// Prepare ansible-vault command
	cmd := exec.Command("ansible-vault encrypt", "--vault-password-file", file.Name())

	if os.Getenv("environment") == "dev" {
		// Run /bin/cat in dev mode
		cmd = exec.Command("/bin/cat")
	}

	stdin, err := cmd.StdinPipe()
	if err != nil { return nil, err }

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, config)
	}()

	return cmd.CombinedOutput()
}

func _encrypt_handler(res http.ResponseWriter, req *http.Request) {
	var payload Payload
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	if (payload.Key == "") || (payload.Body == "") {
		http.Error(res, "Missing key or Body", http.StatusBadRequest)
	}
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", _encrypt_handler)
	return router
}
