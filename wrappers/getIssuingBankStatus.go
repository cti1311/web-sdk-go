package wrappers

import (
	"github.com/cti1311/web-sdk-go/http"
	"github.com/cti1311/web-sdk-go/utils"
	"net/url"
	"strconv"
)

// VerifyPayment is a wrapper for the verify_payment command
// It takes in the credentials and the var1 value and returns the response as a dictionary
func GetIssuingBankStatus(creds utils.Creds, apiEndPoint string , bin int) (map[string]interface{}, error) {

	command := "getIssuingBankStatus"
	var1 := strconv.Itoa(bin)
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