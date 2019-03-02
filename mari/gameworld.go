package mari

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"github.com/kataras/golog"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
)

func GetGameWorlds(baseUrl string) (gameWorlds []GameWorld, languageGroup LanguageGroup, err error) {
	golog.Info("Server listesi ve dil grubu alınıyor...")
	req := newRequest(methodGET, baseUrl, "", nil, nil)
	res, err := execRequest(req, true, nil)
	if err != nil {
		return
	}
	windowDataE0 := regexWindowData.FindStringSubmatch(res)
	windowDataE1, err := base64.StdEncoding.DecodeString(windowDataE0[1])
	if err != nil {
		return
	}
	windowDataE2 := bytes.NewReader(windowDataE1)
	windowDataE3, err := gzip.NewReader(windowDataE2)
	if err != nil {
		return
	}
	windowDataE4, err := ioutil.ReadAll(windowDataE3)
	if err != nil {
		return
	}

	var par parGameWorldAndLangugeGroup
	err = msgpack.Unmarshal(windowDataE4, &par)
	if err != nil {
		return
	}
	gameWorlds = par.A.List
	languageGroup = par.LanguageGroup
	return
}
