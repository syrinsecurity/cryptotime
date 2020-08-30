# Crypto Time

![Go report card](https://goreportcard.com/badge/github.com/syrinsecurity/cryptotime)
[![Build Status](https://travis-ci.org/syrinsecurity/cryptotime.svg?branch=master)](https://travis-ci.org/syrinsecurity/cryptotime)
[![GoDoc](https://godoc.org/github.com/syrinsecurity/cryptotime?status.svg)](https://godoc.org/github.com/syrinsecurity/cryptotime)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/syrinsecurity/cryptotime/graphs/commit-activity)
[![License](https://img.shields.io/github/license/syrinsecurity/cryptotime.svg)](https://github.com/syrinsecurity/cryptotime/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/syrinsecurity/cryptotime.svg)](https://GitHub.com/syrinsecurity/cryptotime/releases/)
[![GitHub issues](https://img.shields.io/github/issues/syrinsecurity/cryptotime.svg)](https://GitHub.com/syrinsecurity/cryptotime/issues/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

CryptoTime puts a end to client side time spoofing attacks. Time is checked from our servers which is signed using our PGP key. This ensures the traffic has not been tampered with. To prevent reply attacks we also implement a 40 character nonce/challenge which the client sends with each request. This will prevent a attack for just sending the response form a previous reply.

- No longer trust client side time
- Trust the time is correct and hasn't been tampered with
- Ensure reply attacks can not occur
- Works just like the time package.

```go
cryptotime.Now()
```

Cryptotime.Now() will return everything you would expect time.Now(), a time.Time.
Its really simple, just a drop in replacement.

## REST API

Rest API supports both JSON and protobufers. To use protobufers just append `?protobuf=1`. https://api.syrinsecurity.net

```
GET /v1/common/cryptotime/:challenge
Host: api.syrinsecurity.net
```

```json
{
        "unix": 1598819088,
        "unixNano": 1598819088784807000,
        "nano": 784806859,
        "UnixDate": "Sun Aug 30 21:24:48 BST 2020",
        "RFC850": "Sunday, 30-Aug-20 21:24:48 BST",
        "RFC1123": "Sun, 30 Aug 2020 21:24:48 BST",
        "stamp": "Aug 30 21:24:48",
        "stampMillisecond": "Aug 30 21:24:48.784",
        "stampMicrosecond": "Aug 30 21:24:48.784806",
        "stampNanoSecond": "Aug 30 21:24:48.784806859",
        "challenge": "asdgadfhnw5ysdrhgssgnsfgnhdfgndfgnfgnsfg",
        "hash": "86AA971ED2DB1A44BEF92BBA580CF93B1BB3DCC3CEF8ED5EBF4A77F67A09FB51",
        "signature": "-----BEGIN PGP SIGNATURE-----\n\nws8dA5xo-----END PGP SIGNATURE-----",
        "success": true
}
```
