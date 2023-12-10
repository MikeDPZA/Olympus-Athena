package responses

import "github.com/google/uuid"

type PutOAuthClientResponse struct {
	ClientId     uuid.UUID `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
}
