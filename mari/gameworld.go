package mari

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"time"
)

type GameWorld struct {
	UUID                    string `msgpack:"uuid"`
	Shortcut                string `msgpack:"shortcut"`
	Name                    string `msgpack:"name"`
	URL                     string `msgpack:"url"`
	Status                  int    `msgpack:"status"`
	RegistrationKeyRequired bool   `msgpack:"registrationKeyRequired"`
	Start                   int64  `msgpack:"start"`
}

func (gs *GameWorld) IsRegistrationOpen() bool {
	return gs.Status == 1
}

func (gs *GameWorld) StartTime() time.Time {
	return time.Unix(gs.Start, 0)
}

func GetGameWorlds(baseUrl string) (gameWorlds []GameWorld, err error) {
	req := newRequest(methodGET, baseUrl, "", nil, nil)
	res, err := execRequest(req, nil)
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

	var gwl resGameWorldList
	err = msgpack.Unmarshal(windowDataE4, &gwl)
	if err != nil {
		return
	}
	gameWorlds = gwl.A.List
	return
}
