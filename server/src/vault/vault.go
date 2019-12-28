package vault

import (
	"io"
	"os"
	"errors"
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

func encrypt_handler(res http.ResponseWriter, req *http.Request) {
	// 2MB max payload
	const MAX_BYTES = 1048576

	// Enforce application/json
	if req.Header.Get("Content-Type") != "application/json" {
		http.Error(res, "Content Type must be 'application/json'", http.StatusUnsupportedMediaType)
		return
	}

	body := http.MaxBytesReader(res, req.Body, MAX_BYTES)

	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	var payload Payload
	err := decoder.Decode(&payload)

	// TODO :: Convert this disaster into a middleware or something
	if err != nil {
		var syntax *json.SyntaxError
		var decoding *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntax):
			http.Error(res, err.Error(), http.StatusBadRequest)

		case errors.As(err, &decoding):
			http.Error(res, err.Error(), http.StatusBadRequest)

		case errors.Is(err, io.EOF):
			http.Error(res, "Request body empty", http.StatusBadRequest)

		default:
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if (payload.Key == "") || (payload.Body == "") {
		http.Error(res, "Missing key or Body", http.StatusBadRequest)
	}
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", encrypt_handler)
	return router
}
