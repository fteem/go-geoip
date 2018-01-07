package api

import (
	"github.com/fteem/go-geoip/util"
	"io/ioutil"
	"net/http"
)

func Fetch(target string) []byte {
	endpoint := "http://freegeoip.net/json/" + target

	resp, responseError := http.Get(endpoint)
	util.Check(responseError)

	body, readError := ioutil.ReadAll(resp.Body)
	util.Check(readError)

	return body
}
