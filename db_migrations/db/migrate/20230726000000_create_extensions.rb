# frozen_string_literal: true

# Migration to enable PostgreSQL extensions for generating UUIDs and using cryptographic functions.
class CreateExtensions < ActiveRecord::Migration[6.1]
  def change
    enable_extension 'pgcrypto'
    enable_extension 'uuid-ossp'
  end
end
