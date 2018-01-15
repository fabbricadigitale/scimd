package server

// AuthenticationScheme ...
type AuthenticationScheme int

//go:generate stringer -type AuthenticationScheme $GOFILE

const (
	// OAuth identifies ...
	OAuth AuthenticationScheme = iota
	// OAuth identifies ...

	// OAuth2 identifies ...
	OAuth2

	// OAuthBearerToken identifies ...
	OAuthBearerToken

	// HTTPBasic identifies ...
	HTTPBasic

	// HTTPDigest identifies ...
	HTTPDigest
)
