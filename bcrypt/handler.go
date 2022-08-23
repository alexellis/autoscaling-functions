package function

import (
	"crypto/rand"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Handle a serverless request
func Handle(w http.ResponseWriter, r *http.Request) {

	var body []byte
	if r.Body != nil {
		body, _ = ioutil.ReadAll(r.Body)
		defer r.Body.Close()
	}

	if r.URL.Path == "/decode" {
		if len(body) == 0 {
			http.Error(w, "No body", http.StatusBadRequest)
			return
		}
		decoded := decode(body)
		w.Write([]byte(decoded))
		return
	}

	if len(body) == 0 {
		body := []byte{32}
		rand.Read(body)
	}

	res, _ := bcrypt.GenerateFromPassword(body, bcrypt.DefaultCost)
	w.Write(res)

}

func decode(req []byte) string {
	res := decodeValue(req)
	if res {
		return "true"
	}
	return "false"
}

func decodeValue(req []byte) bool {
	data := string(req)
	index := strings.Index(data, " ")
	if index == -1 {
		return false
	}

	hash := data[index+1:]
	plain := data[:index]

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)); err != nil {
		return false
	}

	return true
}
