package mari

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var DefaultHttpClient http.Client

// http için
const (
	methodGET             = "GET"
	methodPOST            = "POST"
	headerAccept          = "Accept"
	headerAcceptEncoding  = "Accept-Encoding"
	headerAcceptLanguage  = "Accept-Language"
	headerUserAgent       = "User-Agent"
	headerReferer         = "Referer"
	headerContentEncoding = "Content-Encoding"
	defaultUserAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	defaultAccept         = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
	defaultAcceptEncoding = "gzip, deflate, br"
	defaultAcceptLanguage = "tr-TR,tr;q=0.9,en-US;q=0.8,en;q=0.7"
)

func newRequest(method, url, ref string, body io.Reader, a *Account) (req *http.Request) {
	req, _ = http.NewRequest(method, url, body)
	if a == nil {
		req.Header.Add(headerAccept, defaultAccept)
		req.Header.Add(headerAcceptEncoding, defaultAcceptEncoding)
		req.Header.Add(headerAcceptLanguage, defaultAcceptLanguage)
		req.Header.Add(headerUserAgent, defaultUserAgent)
	} else {
		//todo: accounts'headers
	}
	if ref != "" {
		req.Header.Add(headerReferer, ref)
	}
	return
}

func execRequest(req *http.Request, a *Account) (res string, err error) {
	var resp *http.Response
	if a == nil {
		resp, err = DefaultHttpClient.Do(req)
	} else {
		//todo: accounts'client
	}
	if err != nil {
		return
	}
	//status codes
	if resp.StatusCode != http.StatusOK {
		return res, errors.New(fmt.Sprintf("http %d: %s", resp.StatusCode, resp.Status))
	}

	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get(headerContentEncoding) {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	//todo: daha fazla encoding eklenecek!
	default:
		reader = resp.Body
	}

	respBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	res = string(respBytes)
	return
}
