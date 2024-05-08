package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type MerchantPartnerConfig struct {
	MerchantID uuid.UUID
	ClientID   uuid.UUID
	AppConfigs datatypes.JSON
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
