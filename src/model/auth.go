package model

type AuthCredential struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Authorization struct {
	Token string `json:"token"`
}
