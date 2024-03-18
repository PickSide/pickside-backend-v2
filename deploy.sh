#bin/bash

REPO_DIR="/home/ubuntu/dev/pickside-backend-v2"

cd /home/ubuntu/dev/

if [ ! -d "$REPO_DIR" ]; then
    git clone https://github.com/PickSide/pickside-backend-v2.git
fi

cd $REPO_DIR

git config --global --add safe.directory $REPO_DIR

git checkout release || exit 1
git pull origin release || exit 1

docker-compose down || exit 1
docker-compose pull || exit 1
docker-compose up -d || exit 1