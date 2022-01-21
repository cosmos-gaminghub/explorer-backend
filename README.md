# GAME Explorer Backend
Blockchain exporter for the [CosmosSDK](https://github.com/cosmos/cosmos-sdk) based Blockchain(ex [gaia](https://github.com/cosmos/gaia)). This backend server can be applied to any cosmosSDK based blockchain to sync as mongoDB collections.

**Prerequisites**
* go1.17.0+
* mongoDB 4.4.0+
* docker (for local usage)

## Get Started
- set up mongoDB

- build source

Just run exporter binary after building source.

```bash
git clone https://github.com/cosmos-gaminghub/explorer-backend.git
cd explorer-backend
cp .env.example .env
make all
./build/exporter
./build/cron
```

or apply service to mangae exporter process.

### Use Service
```sh
### exporter service
tee /etc/systemd/system/explorer-backend.service > /dev/null <<EOF
[Unit]
Description=exploder-backend

[Service]
Type=simple
User=root
Group=root

WorkingDirectory=/root/explorer-backend
ExecStart=/root/explorer-backend/build/exporter

Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable explorer-backend
systemctl start explorer-backend


### cron service
tee /etc/systemd/system/explorer-cron.service > /dev/null <<EOF
[Unit]
Description=exploder-cron

[Service]
Type=simple
User=root
Group=root

WorkingDirectory=/root/explorer-backend
ExecStart=/root/explorer-backend/build/cron

Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable explorer-cron
systemctl start explorer-cron
```

### Start in local

```bash
docker run -p 27017:27017 --name dev-mongo mongo
./build/exporter
```

## Setting environment variables

Env Variables| Description | Default
------------------ | ---------------------------- | --------
DB_ADDR            | mongo server's address       | 127.0.0.1
DB_DATABASE        | database's name              | test
DB_USER            | database's username          |
DB_PASSWORD        | database's password          |
DB_POOL_LOMIT      | database max connection num  | 4096
ADDR_NODE_SERVER   | node lcd URI                 | http://198.13.33.206:1317

CoinGecko Env Variables | Description  | Default
----------------------- | -----------  | ------
COINGECKO_API_ENDPOINT  | API endpoint | https://api.coingecko.com/api
COINGECKO_API_VERSION   | API version  | v3
COINGECKO_CURRENCY      | API currency | usd


## System Architechture
In the backend, exporter binary watches cosmosSDK based chain node(especially blocks and proposals) and insert synced data into mongoDB. Also cron binary updates coingecko price data every 20min.

**ref**:
- [explorer-graphql](https://github.com/cosmos-gaminghub/explorer-graphql) responses API data to the frontend.
- [explorer-fronted](https://github.com/cosmos-gaminghub/explorer-frontend) consists of [nuxt.js](https://nuxtjs.org/) framework.
