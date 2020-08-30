package cryptotime

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

var (
	//DefaultTime is used as a fallback if the time can not be obtained from the server
	DefaultTime time.Time = time.Unix(500, 500)
)

// Now returns the current local time.
func Now() time.Time {

	challengeBytes, err := cryptoRand(20)
	if err != nil {
		return DefaultTime
	}

	challenge := hex.EncodeToString(challengeBytes)

	r, err := fetchTimeNow(challenge)
	if err != nil {
		return DefaultTime
	}

	h := sha256.New()
	fmt.Fprint(h, r.Unix,
		r.UnixNano,
		r.Nano,

		r.UnixDate,
		r.RFC850,
		r.RFC1123,
		r.Stamp,
		r.StampMilli,
		r.StampMicro,
		r.StampNano,
		challenge,
	)

	if r.Hash != fmt.Sprintf("%X", h.Sum(nil)) {
		return DefaultTime
	}

	if verify([]byte(r.Hash), []byte(r.Signature)) == false {
		return DefaultTime
	}

	return time.Unix(r.Unix, r.Nano)
}

//cryptoRand returns a crypto random result
func cryptoRand(length int) ([]byte, error) {

	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
