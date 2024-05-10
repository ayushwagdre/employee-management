class CreateEmployees < ActiveRecord::Migration[6.0]
  def change
    safety_assured do
       # Generate employee codes using a sequence
      execute <<-SQL
        CREATE SEQUENCE employee_codes
          START WITH 1000
          INCREMENT BY 1
          NO MINVALUE
          NO MAXVALUE
          CACHE 1;
      SQL
      create_table :employees, id: :uuid, default: 'uuid_generate_v4()' do |t|
        t.string :name
        t.string :position
        t.float :salary
        t.boolean :active, default: true
        t.string :code
        t.timestamps default: -> { 'CURRENT_TIMESTAMP' }
      end


      # Set the default value for the code column using the sequence
      execute <<-SQL
        ALTER TABLE employees
        ALTER COLUMN code
        SET DEFAULT CONCAT('EMP', LPAD(nextval('employee_codes')::text, 4, '0'));
      SQL
    end
  end
end
