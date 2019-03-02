package mari

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kataras/golog"
)

func (a *Account) Register(baseURL string, gameWorld GameWorld, languageGroup LanguageGroup,
	registrationKey string, subscribeNewsletter bool) (err error) {

	a.GameWorldUUID = gameWorld.UUID
	golog.Info("Kayıt için ana sayfa yükleniyor...")
	req := newRequest(methodGET, baseURL, "", nil, a)
	res, err := execRequest(req, false, a)
	if err != nil {
		golog.Errorf("Kayıt için ana sayfa yüklenemedi: %s\n", err.Error())
		return
	}

	golog.Info("E-posta kontrol ediliyor...")
	postBody, _ := json.Marshal(&reqValidateEmail{a.Email, gameWorld.URL})
	req = newRequest(methodPOST, baseURL+urlTailValidateEmail, baseURL, bytes.NewBuffer(postBody), a)
	req.Header.Add(headerAccept, headerVAcceptAll)
	req.Header.Add(headerContentType, headerVContentTypeJson)
	res, err = execRequest(req, true, a)
	//todo: check: {"success":true}
	fmt.Println(res)
	//
	//6LeG9noUAAAAAPq5kQg-QmlxmmsrBXOX5jCQShUJ
	//6LeyUoIUAAAAAGkF-wYb0Od-SIggnBAbSSZeHpYQ
	captcha := "asdasd-aQPEeGAB9j07_jxJMjGyzoIxtiKbz9oT0t3uugQzvx8T5I0WGsnqhgqqXKPpJEd9egvFIvWVOsWplFDIposNkkOG0wuEDDZRGbqfSiUkbgg1GF93iisxZogd9P_d4R_nFiK2aLnVk5QyoaTg06WKs9dExhqj67WE6XdHMUot2E8yOLkdXW5UbKoO-KYU08MhSWBPbGMXge75lhIZqfmf3OVCJHiOaBCh42mYQi3Vz_WlYoMR1U6T7crd-tIjpA0iUzbpNindSuo3gWF0Js"
	golog.Info("Kayıt yapılıyor...")
	postBody, _ = json.Marshal(&reqRegister{
		Username:            a.Username,
		Email:               a.Email,
		GameWorld:           gameWorld,
		RegistrationKey:     registrationKey,
		TermsAndConditions:  true,
		SubscribeNewsletter: subscribeNewsletter,
		Host:                hostUrl,
		Inviter:             nil,
		AdInvite:            struct{}{},
		LanguageGroup:       languageGroup,
		GRecaptchaResponse:  captcha,
	})
	req = newRequest(methodPOST, baseURL+urlTailRegister, baseURL, bytes.NewBuffer(postBody), a)
	req.Header.Set(headerAccept, headerVAcceptAll)
	req.Header.Add(headerContentType, headerVContentTypeJson)
	res, err = execRequest(req, true, a)
	fmt.Println(res, err)
	return
}
