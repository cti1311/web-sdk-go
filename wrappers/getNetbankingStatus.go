package wrappers

import (
	"github.com/cti1311/web-sdk-go/http"
	"github.com/cti1311/web-sdk-go/utils"
	"net/url"
)

// VerifyPayment is a wrapper for the verify_payment command
// It takes in the credentials and the var1 value and returns the response as a dictionary
func GetNetbankingStatus(creds utils.Creds, apiEndPoint string, var1 string) (map[string]interface{}, error) {
	command := "getNetbankingStatus"
	// Create the payload
	payload := url.Values{
		"key": {creds.Key},
		"command": {command},
		"var1": {var1},
		"hash": {utils.ApiHasher(creds, utils.ApiStruct{Command: command, Var1: var1})},
	}

	// Send the request and get the response
	response, err := http.PostForm(apiEndPoint, payload)
	if err != nil {
		return nil, err
	}

	// Return the response
	return response, nil
}