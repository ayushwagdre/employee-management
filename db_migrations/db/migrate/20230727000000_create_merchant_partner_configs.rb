# frozen_string_literal: true

# Migration to create the `merchant_partner_configs` table for storing configuration
# data related to merchant partners.
class CreateMerchantPartnerConfigs < ActiveRecord::Migration[6.1]
  def change
    create_table :merchant_partner_configs, id: :uuid, default: 'uuid_generate_v4()' do |t|
      t.uuid :merchant_id
      t.uuid :client_id
      t.jsonb :app_configs, default: {}
      t.timestamps default: -> { 'CURRENT_TIMESTAMP' }
    end

    add_index :merchant_partner_configs, %i[merchant_id client_id], unique: true
  end
end
