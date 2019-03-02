package main

import (
	"crypto/tls"
	"github.com/hexvalid/travium-go/mari"
	"github.com/kataras/golog"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"time"
)

const proxy = true

func main() {
	golog.SetTimeFormat("[15:04:05.00]")
	golog.SetLevel("debug")

	var baseURL = "https://www.travian.com/tr"

	gameWorlds, languageGroup, _ := mari.GetGameWorlds(baseURL)
	rand.Seed(time.Now().UTC().UnixNano())
	number := strconv.Itoa(rand.Intn(10000))
	account := mari.Account{
		Username: "prasfax" + number,
		Email:    "prasfax" + number + "@mail-2-you.com",
		HTTPHeaders: mari.HTTPHeaders{
			Accept:         mari.DefaultAccept,
			AcceptEncoding: mari.DefaultAcceptEncoding,
			AcceptLanguage: mari.DefaultAcceptLanguage,
			UserAgent:      mari.DefaultUserAgent,
		},
	}
	account.CookieJar, _ = cookiejar.New(nil)
	account.HTTPClient = http.Client{
		Jar: account.CookieJar,
	}
	golog.Debug("Hesap adı: ", account.Username)

	if proxy {
		proxyUrl, _ := url.Parse("http://localhost:8080")
		account.HTTPClient.Transport = &http.Transport{Proxy: http.ProxyURL(proxyUrl),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	account.Register(baseURL, gameWorlds[1], languageGroup, "", false)

}
