package services

type OAuthService interface {
	UpsertClient() error
}
