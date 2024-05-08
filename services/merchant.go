package services

//go:generate mockgen -source=./merchant.go -destination=./mock_services/mock_merchant.go -package=mock_services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"practice/lib/cache"
	"practice/models"
	repository "practice/repositories"
)

type MerchantPartnerAppsConfig struct {
	PartnerAppsConfig interface{}
}

type PartnerAppsConfig struct {
	MerchantID    uuid.UUID
	ClientID      uuid.UUID
	PartnerConfig interface{}
}

type MerchantService interface {
	UpsertPartnerAppsConfig(context.Context, PartnerAppsConfig) error
	GetPartnerAppsConfigByMerchantID(context.Context, uuid.UUID) (MerchantPartnerAppsConfig, error)
}

type merchantService struct {
	cache                     cache.Cache
	merchantPartnerConfigRepo repository.MerchantPartnerConfigRepository
}

const (
	merchantConfigCachePrefix          = "merchant:config"
	merchantConfigCacheTimeoutDuration = 15 * time.Minute
)

func NewMerchantService() MerchantService {
	return &merchantService{
		cache:                     cache.NewRedisCache(),
		merchantPartnerConfigRepo: repository.NewMerchantPartnerConfigRepository(),
	}
}

func (m *merchantService) UpsertPartnerAppsConfig(ctx context.Context, updateConfigParams PartnerAppsConfig) error {

	marshalPartnerConfig, err := json.Marshal(updateConfigParams.PartnerConfig)
	if err != nil {
		return err
	}
	params := models.MerchantPartnerConfig{
		MerchantID: updateConfigParams.MerchantID,
		ClientID:   updateConfigParams.ClientID,
		AppConfigs: marshalPartnerConfig,
	}

	err = m.merchantPartnerConfigRepo.Upsert(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (m *merchantService) GetPartnerAppsConfigByMerchantID(ctx context.Context, merchantID uuid.UUID) (MerchantPartnerAppsConfig, error) {
	config, err := m.merchantPartnerConfigRepo.GetConfigByMerchantID(ctx, merchantID)
	if err != nil {
		return MerchantPartnerAppsConfig{}, err
	}

	return MerchantPartnerAppsConfig{
		PartnerAppsConfig: config.AppConfigs,
	}, nil
}
