
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


