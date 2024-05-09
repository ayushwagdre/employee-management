class CreateEmployees < ActiveRecord::Migration[6.0]
  def change
    safety_assured do
      create_table :employees, id: :uuid, default: 'uuid_generate_v4()' do |t|
        t.string :name
        t.string :position
        t.float :salary
        t.boolean :active, default: true
        t.timestamps default: -> { 'CURRENT_TIMESTAMP' }

        # Add the code column with auto-generated value
        t.string :code, null: false
      end

      # Generate employee codes using a sequence
      execute <<-SQL
        CREATE SEQUENCE employee_codes
          START WITH 1000
          INCREMENT BY 1
          NO MINVALUE
          NO MAXVALUE
          CACHE 1;
      SQL

      # Set the default value for the code column
      execute <<-SQL
        ALTER TABLE employees
        ALTER COLUMN code
        SET DEFAULT CONCAT('EMP', nextval('employee_codes'));
      SQL
    end
  end
end
