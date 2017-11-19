package testFramework

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	Host = "http://services.groupkt.com"
)

var (
	client *http.Client

	InvalidMessageErr = errors.New("Invalid message")
)

type (
	HTTPResponse struct {
		StatusCode int
		Body       interface{}
	}

	Country struct {
		Name       string `json:"name"`
		Alpha2Code string `json:"alpha2_code"`
		Alpha3Code string `json:"alpha3_code"`
	}

	RestResponse struct {
		Messages []string  `json:"messages"`
		Result   []Country `json:"result"`
	}

	Response struct {
		RestResponse RestResponse `json:"RestResponse"`
	}
)

func init() {
	client = &http.Client{}
}

func GetAllCountries() (HTTPResponse, error) {
	var httpResponse HTTPResponse
	url := Host + "/country/get/all"
	resp, err := Get(url)
	if err != nil {
		return httpResponse, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return httpResponse, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return httpResponse, err
	}

	httpResponse = HTTPResponse{
		StatusCode: resp.StatusCode,
		Body:       response,
	}

	return httpResponse, nil
}

func Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}

func GetTotal(message string) (int, error) {
	splits := strings.Split(message, " ")
	if len(splits) < 2 {
		return 0, InvalidMessageErr
	}

	str := splits[1]
	totalStr := str[1 : len(str)-1]

	total, err := strconv.Atoi(totalStr)
	if err != nil {
		return 0, InvalidMessageErr
	}

	return total, nil
}
