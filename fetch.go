package cryptotime

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/getlantern/errors"
)

const (
	//VERSION is using semantic versioning
	VERSION = "1.0.0"
)

var (
	//Client is the HTTP client used to fetch the datetime
	Client http.Client

	//Endpoint is queried for the current timestamp
	Endpoint = "https://api.syrinsecurity.net/v1/common/cryptotime/"

	//UserAgent is used in the request to the API
	UserAgent = "CryptoTime lib (https://github.com/syrinsecurity/cryptotime, v" + VERSION + ", " + runtime.GOOS + " " + runtime.GOARCH + ", " + runtime.Version() + ")"
	//ErrStatusCodeNotOk is returned when the API returns a statuscode that was not expected
	ErrStatusCodeNotOk = errors.New("cryptotime: status code not ok")
)

type response struct {
	Unix     int64 `json:"unix"`
	UnixNano int64 `json:"unixNano"`
	Nano     int64 `json:"nano"`

	UnixDate string `json:"unixDate"`
	RFC850   string `json:"RFC850"`
	RFC1123  string `json:"RFC1123"`

	Stamp      string `json:"stamp"`
	StampMilli string `json:"stampMillisecond"`
	StampMicro string `json:"stampMicrosecond"`
	StampNano  string `json:"stampNanoSecond"`

	Challenge string `json:"challenge"`
	Hash      string `json:"hash"`
	Signature string `json:"signature"`
}

func fetchTimeNow(challenge string) (*response, error) {

	req, err := http.NewRequest("GET", Endpoint+challenge, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, ErrStatusCodeNotOk
	}

	var r response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}

	return &r, nil
}
