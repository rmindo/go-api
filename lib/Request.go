package lib

/**
* Import packages
 */
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	// "github.com/joho/godotenv"
)

/**
* API endpoint
 */
var endpoint = "http://52.187.132.233/api/v2/"

/**
* Get request
 */
func Get(name string, data map[string]string, token bool) string {
	return Send("GET", name, data, token)
}

/**
* Post request
 */
func Post(name string, data map[string]string, token bool) string {
	return Send("POST", name, data, token)
}

/**
* Create query string
 */
func Query(data map[string]string) string {

	param := url.Values{}
	for name, value := range data {
		param.Add(name, value)
	}
	return param.Encode()
}

/**
* Send request
 */
func Send(method string, name string, data map[string]string, token bool) string {

	client := &http.Client{}
	req, err := http.NewRequest(method, endpoint+name+"?"+Query(data), nil)

	if token {
		req.Header.Set("Authorization", "Bearer "+os.Getenv("API_TOKEN"))
	}
	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}
