package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"practice/lib/db"
	"practice/lib/errors"
	"practice/models"
)

type MerchantPartnerConfigRepository interface {
	Upsert(ctx context.Context, config models.MerchantPartnerConfig) error
	GetConfigByMerchantID(ctx context.Context, merchantID uuid.UUID) (*models.MerchantPartnerConfig, error)
	GetConfigByClientID(ctx context.Context, clientID uuid.UUID) (*models.MerchantPartnerConfig, error)
}

type merchantPartnerConfigRepository struct {
	db db.DB
}

func NewMerchantPartnerConfigRepository() MerchantPartnerConfigRepository {
	return &merchantPartnerConfigRepository{db: db.Get()}
}

func (r *merchantPartnerConfigRepository) Upsert(ctx context.Context, config models.MerchantPartnerConfig) error {

	err := r.db.Get().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "merchant_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"app_configs", "updated_at"}),
	}).Create(&config).Error
	if err != nil {
		return errors.Wrapf(ErrUnexpected, "failed to upsert merchant partner config for merchant_id %s %s", config.MerchantID, err.Error())
	}

	return nil
}

func (r *merchantPartnerConfigRepository) getConfigBy(ctx context.Context, condition string, args ...interface{}) (*models.MerchantPartnerConfig, error) {

	config := models.MerchantPartnerConfig{}
	query := r.db.Get().Where(condition, args...).First(&config)
	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(ErrRecordNotFound, "merchant partner config not found for %s %s", condition, args)
		}
		return nil, errors.Wrapf(ErrUnexpected, "failed to get merchant partner config for %s %s", condition, err.Error())
	}

	return &config, nil
}

func (r *merchantPartnerConfigRepository) GetConfigByMerchantID(ctx context.Context, merchantID uuid.UUID) (*models.MerchantPartnerConfig, error) {
	return r.getConfigBy(ctx, "merchant_id = ?", merchantID)
}

func (r *merchantPartnerConfigRepository) GetConfigByClientID(ctx context.Context, clientID uuid.UUID) (*models.MerchantPartnerConfig, error) {
	return r.getConfigBy(ctx, "client_id = ?", clientID)
}
