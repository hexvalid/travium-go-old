package mari

type GameWorld struct {
	UUID                    string `msgpack:"uuid" json:"uuid"`
	Shortcut                string `msgpack:"shortcut" json:"shortcut"`
	Name                    string `msgpack:"name" json:"name"`
	URL                     string `msgpack:"url" json:"url"`
	Status                  int    `msgpack:"status" json:"status"`
	RegistrationKeyRequired bool   `msgpack:"registrationKeyRequired" json:"registrationKeyRequired"`
	Start                   int64  `msgpack:"start" json:"start"`
}

type LanguageGroup struct {
	DetectActive   bool   `msgpack:"detect_active" json:"detect_active"`
	Name           string `msgpack:"name" json:"name"`
	Flag           string `msgpack:"flag" json:"flag"`
	Language       string `msgpack:"language" json:"language"`
	LangNative     string `msgpack:"langNative" json:"langNative"`
	LangEnglish    string `msgpack:"langEnglish" json:"langEnglish"`
	CountryNative  string `msgpack:"countryNative" json:"countryNative"`
	CountryEnglish string `msgpack:"countryEnglish" json:"countryEnglish"`
	Direction      string `msgpack:"direction" json:"direction"`
	HrefLang       string `msgpack:"hrefLang" json:"hrefLang"`
}

type reqValidateEmail struct {
	Email        string `json:"email"`
	GameWorldURL string `json:"gameWorldUrl"`
}

type reqRegister struct {
	Username            string `json:"username"`
	Email               string `json:"email"`
	GameWorld           `json:"gameWorld"`
	RegistrationKey     string      `json:"registrationKey"`
	TermsAndConditions  bool        `json:"termsAndConditions"`
	SubscribeNewsletter bool        `json:"subscribeNewsletter"`
	Host                string      `json:"host"`
	Inviter             interface{} `json:"inviter"`
	AdInvite            struct {
	} `json:"adInvite"`
	LanguageGroup      `json:"languageGroup"`
	GRecaptchaResponse string `json:"gRecaptchaResponse"`
}

type parGameWorldAndLangugeGroup struct {
	A struct {
		List []GameWorld `msgpack:"list"`
	} `msgpack:"gameWorlds"`
	LanguageGroup `msgpack:"group"`
}
