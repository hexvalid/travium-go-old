package mari

import "net/http"

type Account struct {
	Username      string
	Email         string
	GameWorldUUID string
	HTTPClient    http.Client
	HTTPHeaders   HTTPHeaders
	CookieJar     http.CookieJar
}

// @Account'un alt gereksinimleri:
type HTTPHeaders struct {
	Accept         string
	AcceptEncoding string
	AcceptLanguage string
	UserAgent      string
}
