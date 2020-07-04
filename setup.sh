#!/bin/sh

sudo apt-get update
sudo apt-get install software-properties-common
sudo add-apt-repository universe
sudo add-apt-repository ppa:certbot/certbot
sudo apt-get update

sudo apt-get install certbot python3-certbot-nginx -y

sudo apt-get install wget screen git -y

bash <(curl -L -s https://install.direct/go.sh)

git clone https://github.com/Herts/ray-ui-cli.git

cd ray-ui-cli
cp conf/v2ray.tpl /etc/v2ray/config.json
bash update.sh




