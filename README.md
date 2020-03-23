# v2ray Server-Side Management Tool

## Features

1. User management (add, update, delete)

2. Statistics of data consumed of users.

3. Add server name in nginx

4. Apply tls cert using certbot

## Limitations

1. Only support debian/ubuntu

2. Only available for nginx + ws + (tls)

## Steps

1. Install v2ray using ``bash <(curl -L -s https://install.direct/go.sh)``

2. Install nginx, certbot

For certbot: https://certbot.eff.org/

3. Run this program as root, if system control operations are needed
