package types

import (
	"encoding/json"
	"fmt"
	"marketplace/config"
	"net/http"
	"strings"
	"unicode"
)

func checkLogin(login string, cfg config.ServiceConfig) bool {
	if len(login) < cfg.MinLoginLength || len(login) > cfg.MaxLoginLength {
		return false
	}

	for _, c := range login {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '_' {
			return false
		}
	}

	return true
}

func checkPassword(password string, cfg config.ServiceConfig) bool {
	return strings.ContainsAny(password, "!@#$%^&*?/") && 
		len(password) >= cfg.MinPasswordLength &&
		len(password) <= cfg.MaxPasswordLength
}

// Requests ----------------------------------------------------------------------

type PostRegisterRequest struct {
	Login 		string		`json:"login"`
	Password	string		`json:"password"`
}

func CreatePostRegisterRequest(r *http.Request, cfg config.ServiceConfig) (*PostRegisterRequest, error) {
	const op = "CreatePostRegisterRequest"

	var req PostRegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if !checkLogin(req.Login, cfg) {
		return nil, fmt.Errorf("%s: %w", op, ErrBadLoginFormat)
	} else if !checkPassword(req.Password, cfg) {
		return nil, fmt.Errorf("%s: %w", op, ErrBadPasswordFormat)
	}

	return &req, nil
}

type PostLoginRequest struct {
	Login 		string		`json:"login"`
	Password	string		`json:"password"`
}

func CreatePostLoginRequest(r *http.Request) (*PostLoginRequest, error) {
	const op = "CreatePostLoginRequest"

	var req PostLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &req, nil
}

// Responses ---------------------------------------------------------------------

type PostRegisterResponse struct {
	UserId 	string	`json:"user_id"`
}

type PostLoginResponse struct {
	Token	string	`json:"token"`
}
