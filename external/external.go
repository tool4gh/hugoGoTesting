package external

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var baseUrl = "https://api.openpath.com/"

func Login() string {

	payload := strings.NewReader(`{"email":"<EMAIL>","password":"<PASSWORD>"}`)

	req, err := http.NewRequest("POST", baseUrl, payload)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()
	fmt.Println(err)
	body, _ := io.ReadAll(res.Body)
	return string(body)
}
