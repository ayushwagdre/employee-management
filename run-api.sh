#!/bin/bash


echo "Starting DB migrate."

cd db_migrations

#create db and migrate
bundle exec rake db:create
bundle exec rake db:migrate

if [[ $? != 0 ]]; then
 echo "Failed to migrate. Exiting."
 exit 1
fi
echo "DB migrate done!"
cd -

exec ./api
