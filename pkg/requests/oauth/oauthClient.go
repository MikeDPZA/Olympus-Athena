package oauth

const (
	AuthorizationCodeGrantType = "authorization_code"
	ImplicitGrantType          = "implicit"
	PasswordGrantType          = "password"
	ClientCredentialsGrantType = "client_credentials"
	RefreshTokenGrantType      = "refresh_token"
)

const (
	CodeResponseType  = "code"
	TokenResponseType = "token"
)

const (
	NoneTokenEndpointAuthMethod              = "none"
	ClientSecretPostTokenEndpointAuthMethod  = "client_secret_post"
	ClientSecretBasicTokenEndpointAuthMethod = "client_secret_basic"
)

type ClientRegistrationRequest struct {
	ClientName string   `json:"client_name"`
	ClientUri  string   `json:"client_uri"`
	Contacts   []string `json:"contacts"`

	// authorization_code (default), implicit, password, client_credentials, refresh_token
	GrantTypes   []string `json:"grant_types"`
	LogoUri      string   `json:"logo_uri"`
	Jwks         string   `json:"jwks"`
	JwksUri      string   `json:"jwks_uri"`
	RedirectUris []string `json:"redirect_uris"`

	PolicyUri string `json:"policy_uri"`

	// code (default), token
	ResponseType string `json:"response_type"`

	Scope           string `json:"scope"`
	SoftwareId      string `json:"software_id"`
	SoftwareVersion string `json:"software_version"`

	// none - public client
	// client_secret_post - uses http post params
	// client_secret_basic - uses http basic
	TokenEndpointAuthMethod string `json:"token_endpoint_auth_method"`
	TosUri                  string `json:"tos_uri"`
}
