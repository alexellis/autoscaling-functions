package function

import (
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
	if r.URL.Query().Get("action") == "decode" {

		if len(body) == 0 {
			http.Error(w, "No body", http.StatusBadRequest)
			return
		}
		decoded := decode(body)
		w.Write([]byte(decoded))
		return
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

	plain := data[:index]
	hash := data[index+1:]
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)); err != nil {
		return false
	}

	return true
}
