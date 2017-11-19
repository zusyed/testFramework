package testFramework

import (
	"testing"
)

func TestGetAllCountries(t *testing.T) {
	t.Logf("Making an HTTP request to get all countries...")
	resp, err := GetAllCountries()
	if err != nil {
		t.Fatalf("Encountered an error gettings all countries: %s", err)
	}

	t.Logf("Verifying HTTP status code...")
	if resp.StatusCode != 200 {
		t.Fatalf("Encountered an error gettings all countries: %s", err)
	}

	t.Logf("Verifying response body...")
	response, ok := resp.Body.(GetAllCountriesResponse)
	if !ok {
		t.Fatalf("Expected body to be of type GetAllCountriesResponse but was not")
	}

	t.Logf("Verifying messages...")
	if len(response.RestResponse.Messages) < 1 {
		t.Fatalf("Expected at least one message but got none")
	}

	total, err := GetTotal(response.RestResponse.Messages[0])
	if err != nil {
		t.Fatalf("Encountered an error getting total from message: %s", err)
	}

	if total != len(response.RestResponse.Result) {
		t.Fatalf("total in message and length of results do not match")
	}

	t.Logf("Verifying result set...")
	if len(response.RestResponse.Result) == 0 {
		t.Fatalf("Length of results is 0")
	}

	expectedFirstResult := Country{
		Name:       "Afghanistan",
		Alpha2Code: "AF",
		Alpha3Code: "AFG",
	}
	if response.RestResponse.Result[0] != expectedFirstResult {
		t.Fatalf("Expected the first result to be %+v but was %+v", expectedFirstResult, response.RestResponse.Result[0])
	}
}

func TestGetCountryByAlpha2Code(t *testing.T) {
	var tests = []struct {
		in       string
		expected Country
		message  string
	}{
		{
			"AF",
			Country{
				Name:       "Afghanistan",
				Alpha2Code: "AF",
				Alpha3Code: "AFG",
			},
			"Country found matching code [AF].",
		},
		{
			"AX",
			Country{
				Name:       "��land Islands",
				Alpha2Code: "AX",
				Alpha3Code: "ALA",
			},
			"Country found matching code [AX].",
		},
		{
			"CI",
			Country{
				Name:       "C��te d'Ivoire",
				Alpha2Code: "CI",
				Alpha3Code: "CIV",
			},
			"Country found matching code [CI].",
		},
		{
			"CW",
			Country{
				Name:       "Cura��ao",
				Alpha2Code: "CW",
				Alpha3Code: "CUW",
			},
			"Country found matching code [CW].",
		},
		{
			"IR",
			Country{
				Name:       "Iran (Islamic Republic of)",
				Alpha2Code: "IR",
				Alpha3Code: "IRN",
			},
			"Country found matching code [IR].",
		},
		{
			"BL",
			Country{
				Name:       "Saint Barth��lemy",
				Alpha2Code: "BL",
				Alpha3Code: "BLM",
			},
			"Country found matching code [BL].",
		},
		{
			in:      "AB",
			message: "No matching country found for requested code [AB].",
		},
		{
			in:      "af",
			message: "No matching country found for requested code [af].",
		},
		{
			in:      "ASDFGHJKL",
			message: "No matching country found for requested code [ASDFGHJKL].",
		},
	}

	for i, test := range tests {
		t.Logf("Running test %d", i)
		testGetCountryByAlpha2CodeHelper(t, test.in, test.expected, test.message)
	}
}

func testGetCountryByAlpha2CodeHelper(t *testing.T, alpha2Code string, expectedCountry Country, expectedMessage string) {
	t.Logf("Making an HTTP request to get country by alpha2 code '%s'...", alpha2Code)
	resp, err := GetCountryByAlpha2Code(alpha2Code)
	if err != nil {
		t.Fatalf("Encountered an error gettings country by alpha2 code: %s", err)
	}

	t.Logf("Verifying HTTP status code...")
	if resp.StatusCode != 200 {
		t.Fatalf("Encountered an error gettings all countries: %s", err)
	}

	t.Logf("Verifying response body...")
	response, ok := resp.Body.(GetCountryResponse)
	if !ok {
		t.Fatalf("Expected body to be of type GetCountryResponse but was not")
	}

	t.Logf("Verifying messages...")
	if len(response.RestResponse.Messages) < 1 {
		t.Fatalf("Expected at least one message but got none")
	}

	if response.RestResponse.Messages[0] != expectedMessage {
		t.Fatalf("Expected message to be '%s' but was '%s'", expectedMessage, response.RestResponse.Messages[0])
	}

	t.Logf("Verifying result...")
	if response.RestResponse.Result != expectedCountry {
		t.Fatalf("Expected result to be %+v but was %+v", expectedCountry, response.RestResponse.Result)
	}
}

func TestGetCountryByAlpha3Code(t *testing.T) {
	var tests = []struct {
		in       string
		expected Country
		message  string
	}{
		{
			"AFG",
			Country{
				Name:       "Afghanistan",
				Alpha2Code: "AF",
				Alpha3Code: "AFG",
			},
			"Country found matching code [AFG].",
		},
		{
			"ALA",
			Country{
				Name:       "��land Islands",
				Alpha2Code: "AX",
				Alpha3Code: "ALA",
			},
			"Country found matching code [ALA].",
		},
		{
			"CIV",
			Country{
				Name:       "C��te d'Ivoire",
				Alpha2Code: "CI",
				Alpha3Code: "CIV",
			},
			"Country found matching code [CIV].",
		},
		{
			"CUW",
			Country{
				Name:       "Cura��ao",
				Alpha2Code: "CW",
				Alpha3Code: "CUW",
			},
			"Country found matching code [CUW].",
		},
		{
			"HMD",
			Country{
				Name:       "Heard Island and McDonald Islands",
				Alpha2Code: "HM",
				Alpha3Code: "HMD",
			},
			"Country found matching code [HMD].",
		},
		{
			"IRN",
			Country{
				Name:       "Iran (Islamic Republic of)",
				Alpha2Code: "IR",
				Alpha3Code: "IRN",
			},
			"Country found matching code [IRN].",
		},
		{
			"BLM",
			Country{
				Name:       "Saint Barth��lemy",
				Alpha2Code: "BL",
				Alpha3Code: "BLM",
			},
			"Country found matching code [BLM].",
		},
		{
			"GBR",
			Country{
				Name:       "United Kingdom of Great Britain and Northern Ireland",
				Alpha2Code: "GB",
				Alpha3Code: "GBR",
			},
			"Country found matching code [GBR].",
		},
		{
			in:      "ABC",
			message: "No matching country found for requested code [ABC].",
		},
		{
			in:      "AB",
			message: "No matching country found for requested code [AB].",
		},
		{
			in:      "afg",
			message: "No matching country found for requested code [afg].",
		},
		{
			in:      "ASDFGHJKL",
			message: "No matching country found for requested code [ASDFGHJKL].",
		},
	}

	for i, test := range tests {
		t.Logf("Running test %d", i)
		testGetCountryByAlpha3CodeHelper(t, test.in, test.expected, test.message)
	}
}

func testGetCountryByAlpha3CodeHelper(t *testing.T, alpha3Code string, expectedCountry Country, expectedMessage string) {
	t.Logf("Making an HTTP request to get country by alpha3 code '%s'...", alpha3Code)
	resp, err := GetCountryByAlpha3Code(alpha3Code)
	if err != nil {
		t.Fatalf("Encountered an error gettings country by alpha2 code: %s", err)
	}

	t.Logf("Verifying HTTP status code...")
	if resp.StatusCode != 200 {
		t.Fatalf("Encountered an error gettings all countries: %s", err)
	}

	t.Logf("Verifying response body...")
	response, ok := resp.Body.(GetCountryResponse)
	if !ok {
		t.Fatalf("Expected body to be of type GetCountryResponse but was not")
	}

	t.Logf("Verifying messages...")
	if len(response.RestResponse.Messages) < 1 {
		t.Fatalf("Expected at least one message but got none")
	}

	if response.RestResponse.Messages[0] != expectedMessage {
		t.Fatalf("Expected message to be '%s' but was '%s'", expectedMessage, response.RestResponse.Messages[0])
	}

	t.Logf("Verifying result...")
	if response.RestResponse.Result != expectedCountry {
		t.Fatalf("Expected result to be %+v but was %+v", expectedCountry, response.RestResponse.Result)
	}
}

func TestGetTotal(t *testing.T) {
	var tests = []struct {
		in       string
		expected int
		err      error
	}{
		{"Total [249] records found.", 249, nil},
		{"Total [296] records found.", 296, nil},
		{"Total [30] records found.", 30, nil},
		{"InvalidMessage", 0, InvalidMessageErr},
		{"Invalid message", 0, InvalidMessageErr},
	}

	for i, test := range tests {
		t.Logf("Running test %d", i)

		actual, err := GetTotal(test.in)
		if err != test.err {
			t.Fatalf("Expected error to be %s but was %s", test.err, err)
		}

		if actual != test.expected {
			t.Errorf("Expected total to be %d but was %d", test.expected, actual)
		}
	}
}
