package model

type Message struct {
	Message string
}

func NewMessageResponse(msg string) *Message {
	message := &Message{
		Message: msg,
	}
	return message
}

type SingInRequest struct {
	Email    string
	Password string
}

type TokenResponse struct {
	Token string
}

func NewTokenResponse(token string) *TokenResponse {
	tokenResponse := &TokenResponse{
		Token: token,
	}
	return tokenResponse
}
