package handler

type failAnswer struct {
	Message string `json:"message"`
}

type successRegister struct {
	Id int `json:"id"`
}

type successLogin struct {
	Token string `json:"token"`
}
