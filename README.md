# GAME Explorer Backend
Blockchain explorer for the [CosmosSDK](https://github.com/cosmos/cosmos-sdk) based Blockchain(ex [gaia](https://github.com/cosmos/gaia)). This backend server can be applied to any cosmosSDK based blockchain.

**Prerequisites**
* go1.16.0+
* mongoDB 4.4.0+
* docker (for local usage)

## Get Started
- set up mongoDB

- build source
Just run explorer binary after building source.
```bash
git clone https://github.com/cosmos-gaminghub/explorer-backend.git
cd explorer-backend
cp .env.example .env
make all
./build/explorer
```

or apply service to mangae explorer process.

### Use Service
```sh
### service
tee /etc/systemd/system/explorer-backend.service > /dev/null <<EOF
[Unit]
Description=exploder-backend

[Service]
Type=simple
User=root
Group=root

WorkingDirectory=/root/explorer-backend
ExecStart=/root/explorer-backend/build/explorer

Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable explorer-backend
systemctl start explorer-backend
```

### Start in local

```bash
docker run -p 27017:27017 --name dev-mongo mongo
./build/explorer
```

## Setting environment variables

Env Variables| Description | Default
------------------ | ---------------------------- | --------
DB_ADDR            | mongo server's address       | 127.0.0.1
DB_DATABASE        | database's name              | test
DB_USER            | database's username          |
DB_PASSWORD        | database's password          |
DB_POOL_LOMIT      | database max connection num  | 4096
PORT               | explorer server's port       | 8080
ADDR_NODE_SERVER   | node lcd URI                 | http://198.13.33.206:1317

CoinGecko Env Variables | Description  | Default
----------------------- | -----------  | ------
COINGECKO_API_ENDPOINT  | API endpoint | https://api.coingecko.com/api
COINGECKO_API_VERSION   | API version  | v3
COINGECKO_CURRENCY      | API currency | usd


## System Architechture
In the backend, explorer binary watch cosmosSDK based chain node and insert sync data in mongoDB. Also cron binary update coingecko price data every 20min. explorer-graphql response API data to frontend. explorer-fronted consists of nuxt.js framework.
