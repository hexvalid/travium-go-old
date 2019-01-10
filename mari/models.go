package mari

type reqValidateEmail struct {
	Email        string `json:"email"`
	GameWorldURL string `json:"gameWorldUrl"`
}

type resGameWorldList struct {
	A struct {
		List []GameWorld `msgpack:"list"`
	} `msgpack:"gameWorlds"`
}
