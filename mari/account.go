package mari

type Account struct {
	Email       string
	GameWorld   GameWorld
	HTTPHeaders HTTPHeaders
}

// @Account'un alt gereksinimleri:
type HTTPHeaders struct {
	Accept         string
	AcceptEncoding string
	AcceptLanguage string
	UserAgent      string
}
