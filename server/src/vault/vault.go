package vault

import (
	"io"
	"os"
	"os/exec"
	"net/http"
	"io/ioutil"

	"github.com/go-chi/chi"
)

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

func _handle_encrypt(res http.ResponseWriter, req *http.Request) {

}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", _handle_encrypt)
	return router
}
