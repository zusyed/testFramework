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
	response, ok := resp.Body.(Response)
	if !ok {
		t.Fatalf("Expected body to be of type Response but was not")
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
