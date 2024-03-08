package api

type Error struct {
	ErrorMsg string `json:"error"`
}

func ErrorMessage(m string) Error {
	return Error{m}
}