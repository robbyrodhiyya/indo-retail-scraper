package utils

import "net/http"

func SetBrowserHeaders(req *http.Request) {
	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36",
	)

	req.Header.Set(
		"Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
	)

	req.Header.Set(
		"Accept-Language",
		"en-US,en;q=0.9,id;q=0.8",
	)

	req.Header.Set("Connection", "keep-alive")

	req.Header.Set("deviceType", "Web")
	req.Header.Set("deviceModel", "firefox")
	req.Header.Set("latitude", "0")
	req.Header.Set("longitude", "0")

	req.Header.Set("trxId", "3600243464")

	req.Header.Set(
		"fingerprint",
		"mTHIgpcxKA/sk4kjzVXd84DXywT3h7Sh/zSotyFY7sq8SMfdsxvppl7y9ZHQ/oIO",
	)
}
