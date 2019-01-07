package model

import (
	"time"
)

type GameServer struct {
	UUID                    string `msgpack:"uuid"`
	Shortcut                string `msgpack:"shortcut"`
	Name                    string `msgpack:"name"`
	URL                     string `msgpack:"url"`
	Status                  int    `msgpack:"status"`
	RegistrationKeyRequired bool   `msgpack:"registrationKeyRequired"`
	Start                   int64  `msgpack:"start"`
}

func (gs *GameServer) IsRegistrationOpen() bool {
	return gs.Status == 1
}

func (gs *GameServer) StartTime() time.Time {
	return time.Unix(gs.Start, 0)
}
