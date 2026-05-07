package dto

import (
	"net/http"
	"net/url"
)

const SESSION_COOKIE_NAME = "neblab-session-token"
const SESSION_EXPIRATION_DURATION = 24 * 60 * 60

func NewSessionCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:     SESSION_COOKIE_NAME,
		Value:    url.QueryEscape(token),
		Path:     "/",
		Domain:   "",
		MaxAge:   SESSION_EXPIRATION_DURATION,
		Secure:   false, //todo(Prod) : should be true on production
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}

func NewSessionDeletionToken() *http.Cookie {
	return &http.Cookie{
		Name:   SESSION_COOKIE_NAME,
		Value:  "",
		Path:   "/",
		Domain: "",
		MaxAge: -1,
	}
}
