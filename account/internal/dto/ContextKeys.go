package dto

type ContextKey string

const (
	CKEY_UserId    ContextKey = "USER_ID"
	CKEY_IpAddress ContextKey = "IP_ADDRESS"
	CKEY_UserAgent ContextKey = "USER_AGENT"
)
