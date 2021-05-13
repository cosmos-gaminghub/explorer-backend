# Nibiru Explorer Backend
Blockchain explorer for the Cosmos Gaming Hub Node ([nibiru](https://github.com/cosmos-gaminghub/nibiru))

**Prerequisites**
* go1.16.0+
* docker (for local usage)

## Get Started

```bash
git clone https://github.com/cosmos-gaminghub/explorer-backend.git
make all
./build/explorer

```

### Start in local

```bash
docker run -p 27017:27017 --name dev-mongo mongo
./build/explorer
```

## Setting environment variables

```
    DB_ADDR          : mongo server's address
    DB_DATABASE      : database's name
    DB_USER          : database's username
    DB_PASSWORD      : database's password
    DB_POOL_LIMIT    : database max connection num

    PORT             : explorer server's port
    ADDR_NODE_SERVER : node lcd address
    ADDR_HUB_RPC     : node rpc address
    FAUCET_URL       : faucet address
    CHAIN_ID         : chain-id
    API_VERSION      : explorer api version
    MAX_DRAW_CNT     : Maximum number of collections
    SHOW_FAUCET      : switch of show faucet
    INITIAL_SUPPLY   : initial supplay of token
    CUR_ENV          : current environment(dev/qa/testnet/mainnet)
    CronTimeAssetGateways: time interval of update asset gateways
    CronTimeAssetTokens: time interval of update asset tokens
    CronTimeGovParams: time interval of update gov params
    CronTimeTxNumByDay: time interval of update tx num by day
    CronTimeControlTask: time interval of monitor task execute
    CronTimeAccountRewards: time interval of update account rewards
    CronTimeValidators: time interval of update validators
    CronTimeValidatorIcons: time interval of update validator icons
    CronTimeProposalVoters: time interval of update voter info of proposal
    CronTimeValidatorStaticInfo: time interval of cronjob to update validator static info include uptime, selfBond, delegatorNum
    CronTimeFormatStaticDay: define time format of cronjob execute by every day
    CronTimeFormatStaticMonth: define time format of cronjob execute by eveny month
    CronTimeStaticDataDay: time interval of cronjob to snapshot delegator and validator rewards info
    CronTimeStaticDataMonth: time interval of cronjob to caculate delegator and validator rewards info
    CronTimeHeartBeat: time interval of heart beat in cron task
    NetreqLimitMax: max network request to lcd node


    //cosmos gaming hub v0.2.0 add
    PrefixAccAddr    : nibiru
    PrefixAccPub     : nibirupub
    PrefixValAddr    : nibiruval
    PrefixValPub     : nibiruvalpub
    PrefixConsAddr   : nibirucons
    PrefixConsPub    : nibiruconspub

```
