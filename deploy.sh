#bin/bash

cd /home/ubuntu/dev/

if [ ! -d "pickside-backend-v2" ]; then
    git clone https://github.com/PickSide/pickside-backend-v2.git
fi

git config --global --add safe.directory /home/ubuntu/dev/pickside-backend-v2

cd pickside-backend-v2

git checkout release

git pull origin release

docker-compose down
docker-compose pull
docker-compose up -d