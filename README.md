# practice

# create a development.env
cp env.sample development.env

replace with your own local configuration
loadenv()
{
  echo "Loading $1"
  for i in $(cat $1); do
    export $i
  done
}

export -f loadenv

#create db and migrate
bundle exec rake db:create
bundle exec rake db:migrate

