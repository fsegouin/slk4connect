#! /bin/bash

# Push to dokku-alt
# App will be deployed here: http://slk4connect.segouin.me
git remote add dokku dokku@segouin.me:slk4connect
git push dokku master
