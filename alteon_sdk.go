package do_request

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func AlteonLogin(host string, username string, password string) (*http.Request, error) {
	// client := &http.Client{
	// 	Transport: &http.Transport{
	// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// 	},
	// }

	url := "https://10.170.99.1"

	authHeader := basicAuthHeader(username, password)
	//postData := []byte(`{"key": "value"}`)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating login request:", err)
		return nil, err
	}

	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")
	fmt.Println(req)

	// resp, err := client.Do(req)
	// //fmt.Println(resp)
	// if err != nil {
	// 	fmt.Println("Error making login request:", err)
	// 	return nil, err
	// }
	// defer resp.Body.Close()

	// // Handle the response here
	// fmt.Println("Login successful", resp.StatusCode)
	// fmt.Println("CODE REview")
	// fmt.Println(resp)
	return req, nil

}

// Request is a function in the do_request package
func Request(host string, method string, API string, Data map[string]string, request *http.Request) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	var APIRequest *http.Request
	var new_err error
	URL := "https://" + host + API
	APIBytes, err := json.Marshal(Data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	New_Request, err := http.NewRequest(strings.ToUpper(method), URL, bytes.NewBuffer(APIBytes))
	fmt.Println(New_Request)
	APIRequest = New_Request
	new_err = err
	if new_err != nil {
		fmt.Println("Error creating API request:", new_err)
		return
	}
	// Add headers to the request
	//APIRequest.Header.Set("Content-Type", "application/json")
	//authHeader := basicAuthHeader("admin", "admin")
	//APIRequest.Header.Set("Authorization", fmt.Sprintf("%s", authHeader))
	APIRequest.Header = request.Header.Clone()
	APIResponse, err := client.Do(APIRequest)
	if err != nil {
		fmt.Println("Error making API request:", err)
		return
	}
	defer APIResponse.Body.Close()
	fmt.Println("NEW")
	fmt.Println(APIResponse.StatusCode)

	// Read the response body of the API call
	APIResponseBody, err := ioutil.ReadAll(APIResponse.Body)
	if err != nil {
		fmt.Println("Error reading API response body:", err)
		return
	}

	// Handle the response data of the API call as needed
	// ...

	// Print the response body of the API call
	fmt.Println(string(APIResponseBody))

}

func basicAuthHeader(username, password string) string {
	auth := username + ":" + password
	b64 := base64.StdEncoding.EncodeToString([]byte(auth))
	return "Basic " + b64
}
