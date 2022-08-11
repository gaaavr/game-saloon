package handler

type failAnswer struct {
	Message string `json:"message"`
}

type successRegister struct {
	Id int `json:"id"`
}
