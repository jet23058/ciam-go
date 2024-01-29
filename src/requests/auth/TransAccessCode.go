package requests_atuh

import "github.com/google/uuid"

type TransAccessCode struct {
	AccessCode uuid.UUID `json:"access_code" example:"abcde-fgh-ijk-lmn"`
}
