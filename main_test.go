package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"testing"
)

func Test_authRequired(t *testing.T) {
	var tests []struct {
		name string
		want gin.HandlerFunc
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := authRequired(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authRequired() = %v, want %v", got, tt.want)
			}
		})
	}
	url := "https://dev-c-559zpw.auth0.com/oauth/token"

	payload := strings.NewReader("{\"client_id\":\"u2ZbAXZKz4kM0MM27o6R7mmYQ8pteoFw\",\"client_secret\":\"MSCR94_grYm42kems9ng7jPvAVOlvpHLV-c7xg4UXm190LBbgUtgwkgBujgPHzHq\",\"audience\":\"https://dev-c-559zpw.auth0.com/api/v2/\",\"grant_type\":\"client_credentials\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func Test_setAuth0Variables(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}

	url := "https://dev-c-559zpw.auth0.com/oauth/token"

	payload := strings.NewReader("{\"client_id\":\"u2ZbAXZKz4kM0MM27o6R7mmYQ8pteoFw\",\"client_secret\":\"MSCR94_grYm42kems9ng7jPvAVOlvpHLV-c7xg4UXm190LBbgUtgwkgBujgPHzHq\",\"audience\":\"https://dev-c-559zpw.auth0.com/api/v2/\",\"grant_type\":\"client_credentials\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func Test_sendingToken(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}

	url := "https://localhost:3000/todo"

	payload := strings.NewReader("{\"client_id\":\"u2ZbAXZKz4kM0MM27o6R7mmYQ8pteoFw\",\"client_secret\":\"MSCR94_grYm42kems9ng7jPvAVOlvpHLV-c7xg4UXm190LBbgUtgwkgBujgPHzHq\",\"audience\":\"https://dev-c-559zpw.auth0.com/api/v2/\",\"grant_type\":\"client_credentials\"}")

	req, _ := http.NewRequest("GET", url, payload)

	req.Header.Add("authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ik1rTkJSVEpFTnpVMFJrUTFSRVE1TnpZeU1rSTJNVVJDTmpjNU5USTFORGc0UTBORk16RTFPUSJ9.eyJpc3MiOiJodHRwczovL2Rldi1jLTU1OXpwdy5hdXRoMC5jb20vIiwic3ViIjoidTJaYkFYWkt6NGtNME1NMjdvNlI3bW1ZUThwdGVvRndAY2xpZW50cyIsImF1ZCI6Imh0dHBzOi8vZ29sYW5nLWFuZ3VsYXItYXBpIiwiaWF0IjoxNTgxNjc5NzcyLCJleHAiOjE1ODE3NjYxNzIsImF6cCI6InUyWmJBWFpLejRrTTBNTTI3bzZSN21tWVE4cHRlb0Z3IiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIiwicGVybWlzc2lvbnMiOltdfQ.hMK3cEOQi5AmX9jzoalUWkhcDQ-e-cWZ5FtsOzZhzOQwrR6d6meNa3ZNKcoxpT-lUjS5xm3o86j-Ld6VXASAr4Iws-7rddD3i0aa0oyhREJYNNnOUwF-etJEshCd6pjiyAC7c9U-UMFdSMu0C4ENIHfceGUL_lVt65_bBJkI1JnFAsUi5W9bTpodfcqJdkzlEAzEBZHRFwuC0zybOiesSx8ZHKd5cAJOOpZNVcoksYwAS4Dd7rBG0pED2YoXepyQh-3oNxtogII9EupY6P8dF3yfugXbD0Sr1Xa0zhiffwc__SPvPRe4ousV-8I65lRwkkyP366R9Afkvp0YuEZ1ug")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func Test_authRequired2(t *testing.T) {
	var tests []struct {
		name string
		want gin.HandlerFunc
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := authRequired(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authRequired() = %v, want %v", got, tt.want)
			}
		})
	}
	domain := "dev-c-559zpw.auth0.com"

	fmt.Println(string(domain))
}

func Test_terminateWithError(t *testing.T) {
	type args struct {
		statusCode int
		message    string
		c          *gin.Context
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSetupRouter(t *testing.T) {
	fmt.Println("start router")
	testRouter := SetupRouter()

	postValues := url.Values{}
	postValues.Set("song", rndStr(4))
	postDataStr := postValues.Encode()
	postDataBytes := []byte(postDataStr)
	postBytesReader := bytes.NewReader(postDataBytes)

	req, err := http.NewRequest("GET", "/song", nil)
	post, err := http.NewRequest("POST", "/song", postBytesReader)
	if err != nil {
		fmt.Println(err)
	}

	wg := &sync.WaitGroup{}
	for count := 0; count < 2; count++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			response := httptest.NewRecorder()
			testRouter.ServeHTTP(response, req)
			testRouter.ServeHTTP(response, post)
			fmt.Println(response.Body)
		}()
	}
	wg.Wait()
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func rndStr(n int) string {
	rnd_str := make([]rune, n)
	for i := range rnd_str {
		rnd_str[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(rnd_str)
}
