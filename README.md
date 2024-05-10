
# create a development.env
cp env.sample development.env

# replace with your own local configuration
loadenv()
{
  echo "Loading $1"
  for i in $(cat $1); do
    export $i
  done
}

export -f loadenv

migration is taken by active record migration
# create database
bundle exec rake db:create
# create tables
bundle exec rake db:migrate

Table structure
```
Table employees {
  id uuid [primary key, default: 'uuid_generate_v4()']
  name varchar
  position varchar
  salary float
  active boolean [default: true]
  code varchar [default: "'EMP' || lpad(nextval('employee_codes')::text, 4, '0')"]
  created_at timestamp [default: `CURRENT_TIMESTAMP`]
}
```
code is automatically with prefix EMP followed by 4 digit code in sequntial order(EMP1001)

Rest api is created with following operations
Create
Update
Get
GetAll
Delete is not created we can update the active column to false
