package testFramework

import (
	"encoding/json"
	"errors"
	"fmt"
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

	GetCountriesRestResponse struct {
		Messages []string  `json:"messages"`
		Result   []Country `json:"result"`
	}

	GetCountriesResponse struct {
		RestResponse GetCountriesRestResponse `json:"RestResponse"`
	}

	GetCountryRestResponse struct {
		Messages []string `json:"messages"`
		Result   Country  `json:"result"`
	}

	GetCountryResponse struct {
		RestResponse GetCountryRestResponse `json:"RestResponse"`
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

	var response GetCountriesResponse
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

func GetCountryByAlpha2Code(alpha2Code string) (HTTPResponse, error) {
	var httpResponse HTTPResponse
	url := Host + fmt.Sprintf("/country/get/iso2code/%s", alpha2Code)
	resp, err := Get(url)
	if err != nil {
		return httpResponse, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return httpResponse, err
	}

	var response GetCountryResponse
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

func GetCountryByAlpha3Code(alpha3Code string) (HTTPResponse, error) {
	var httpResponse HTTPResponse
	url := Host + fmt.Sprintf("/country/get/iso3code/%s", alpha3Code)
	resp, err := Get(url)
	if err != nil {
		return httpResponse, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return httpResponse, err
	}

	var response GetCountryResponse
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

func GetCountriesBySearch(search string) (HTTPResponse, error) {
	var httpResponse HTTPResponse
	url := Host + fmt.Sprintf("/country/search?text=%s", search)
	resp, err := Get(url)
	if err != nil {
		return httpResponse, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return httpResponse, err
	}

	var response GetCountriesResponse
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
