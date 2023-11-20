package responses

import "time"

type HealthResponse struct {
	Success bool      `json:"success"`
	Time    time.Time `json:"time"`
}
