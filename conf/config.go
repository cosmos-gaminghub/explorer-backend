package conf

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos-gaminghub/explorer-backend/logger"
)

const (
	KeyDbAddr      = "DB_ADDR"
	KeyDATABASE    = "DB_DATABASE"
	KeyDbUser      = "DB_USER"
	KeyDbPwd       = "DB_PASSWORD"
	KeyDbPoolLimit = "DB_POOL_LIMIT"

	KeyServerPort    = "PORT"
	KeyAddrHubLcd    = "ADDR_NODE_SERVER"
	KeyAddrHubNode   = "ADDR_HUB_RPC"
	KeyAddrFaucet    = "FAUCET_URL"
	KeyChainId       = "CHAIN_ID"
	KeyApiVersion    = "API_VERSION"
	KeyMaxDrawCnt    = "MAX_DRAW_CNT"
	KeyShowFaucet    = "SHOW_FAUCET"
	KeyCurEnv        = "CUR_ENV"
	KeyInitialSupply = "INITIAL_SUPPLY"

	KeyPrefixAccAddr  = "PrefixAccAddr"
	KeyPrefixAccPub   = "PrefixAccPub"
	KeyPrefixValAddr  = "PrefixValAddr"
	KeyPrefixValPub   = "PrefixValPub"
	KeyPrefixConsAddr = "PrefixConsAddr"
	KeyPrefixConsPub  = "PrefixConsPub"

	KeyCronTimeAssetGateways       = "CronTimeAssetGateways"
	KeyCronTimeAssetTokens         = "CronTimeAssetTokens"
	KeyCronTimeGovParams           = "CronTimeGovParams"
	KeyCronTimeTxNumByDay          = "CronTimeTxNumByDay"
	KeyCronTimeControlTask         = "CronTimeControlTask"
	KeyCronTimeHeartBeat           = "CronTimeHeartBeat"
	KeyCronTimeValidators          = "CronTimeValidators"
	KeyCronTimeAccountRewards      = "CronTimeAccountRewards"
	KeyCronTimeValidatorIcons      = "CronTimeValidatorIcons"
	KeyCronTimeProposalVoters      = "CronTimeProposalVoters"
	KeyCronTimeValidatorStaticInfo = "CronTimeValidatorStaticInfo"

	KeyCronTimeFormatStaticDay   = "CronTimeFormatStaticDay"
	KeyCronTimeFormatStaticMonth = "CronTimeFormatStaticMonth"
	KeyCronTimeStaticDataDay     = "CronTimeStaticDataDay"
	KeyCronTimeStaticDataMonth   = "CronTimeStaticDataMonth"
	KeyNetreqLimitMax            = "NetreqLimitMax"
	KeyCaculateDebug             = "CaculateDebug"
	KeyCaculateStartDate         = "CaculateStartDate" //yyyy-mm-ddThh:mm:ss
	KeyCaculateEndDate           = "CaculateEndDate"   //yyyy-mm-ddThh:mm:ss
	KeyCaculateDate              = "CaculateDate"      //yyyy-mm-dd
	KeyFoundationDelegatorAddr   = "FoundationDelegatorAddr"

	EnvironmentDevelop = "dev"
	EnvironmentLocal   = "local"
	EnvironmentQa      = "qa"
	EnvironmentStage   = "stage"
	EnvironmentProd    = "prod"

	InitialSupply      = "2000000000" //IRIS
	DefaultEnvironment = EnvironmentDevelop
)

var (
	config        Config
	defaultConfig = map[string]map[string]string{}
	//IniSupply     string
)

func init() {
	logger.Info("==================================load config start==================================")
	loadDefault()
	// IniSupply = getEnv(KeyInitialSupply, DefaultEnvironment)
	addrs := strings.Split(getEnv(KeyDbAddr, DefaultEnvironment), ",")
	fmt.Print(addrs)
	db := dbConf{
		Addrs:     addrs,
		Database:  getEnv(KeyDATABASE, DefaultEnvironment),
		UserName:  getEnv(KeyDbUser, DefaultEnvironment),
		Password:  getEnv(KeyDbPwd, DefaultEnvironment),
		PoolLimit: getEnvInt(KeyDbPoolLimit, DefaultEnvironment),
	}
	config.Db = db

	rand.Seed(time.Now().Unix())

	hubcf := hubConf{
		Prefix: bech32Prefix{
			AccAddr:  getEnv(KeyPrefixAccAddr, DefaultEnvironment),
			AccPub:   getEnv(KeyPrefixAccPub, DefaultEnvironment),
			ValAddr:  getEnv(KeyPrefixValAddr, DefaultEnvironment),
			ValPub:   getEnv(KeyPrefixValPub, DefaultEnvironment),
			ConsAddr: getEnv(KeyPrefixConsAddr, DefaultEnvironment),
			ConsPub:  getEnv(KeyPrefixConsPub, DefaultEnvironment),
		},
		LcdUrl:  getEnv(KeyAddrHubLcd, DefaultEnvironment),
		NodeUrl: getEnv(KeyAddrHubNode, DefaultEnvironment),
		ChainId: getEnv(KeyChainId, DefaultEnvironment),
	}
	config.Hub = hubcf

	logger.Info("==================================load config end==================================")
}

func loadDefault() {
	defaultConfig[EnvironmentDevelop] = map[string]string{
		KeyDbAddr:         "127.0.0.1:27017",
		KeyDATABASE:       "mydb",
		KeyDbUser:         "",
		KeyDbPwd:          "",
		KeyDbPoolLimit:    "4096",
		KeyServerPort:     "8080",
		KeyAddrHubLcd:     "http://108.61.162.170:1317",
		KeyAddrHubNode:    "http://108.61.162.170:26657",
		KeyAddrFaucet:     "http://127.0.0.1:30200",
		KeyChainId:        "bifrost-2",
		KeyApiVersion:     "v0.6.5",
		KeyMaxDrawCnt:     "10",
		KeyPrefixAccAddr:  "cosmos",
		KeyPrefixAccPub:   "cosmospub",
		KeyPrefixValAddr:  "cosmosvaloper",
		KeyPrefixValPub:   "cosmosvaloperpub",
		KeyPrefixConsAddr: "cosmosvalcons",
		KeyPrefixConsPub:  "cosmosvalconspub",
		KeyShowFaucet:     "1",
		KeyCurEnv:         "dev",
		KeyInitialSupply:  InitialSupply,
	}
}

func Get() Config {
	return config
}

type Config struct {
	Db     dbConf
	Server serverConf
	Hub    hubConf
}

type dbConf struct {
	Addrs     []string
	Database  string
	UserName  string
	Password  string
	PoolLimit int
}

type serverConf struct {
	InstanceNo string
	ServerPort int
	FaucetUrl  string
	ApiVersion string
	MaxDrawCnt int
	ShowFaucet string
	CurEnv     string
}

type hubConf struct {
	Prefix  bech32Prefix
	LcdUrl  string
	NodeUrl string
	ChainId string
}

type bech32Prefix struct {
	AccAddr  string
	AccPub   string
	ValAddr  string
	ValPub   string
	ConsAddr string
	ConsPub  string
}

func getEnv(key string, environment string) string {
	var value string
	if v, ok := os.LookupEnv(key); ok {
		value = v
	} else {
		if DefaultEnvironment == EnvironmentStage || DefaultEnvironment == EnvironmentProd {
			logger.Panic("config is not able to use default config", logger.String("Environment", DefaultEnvironment))
		}
		value = defaultConfig[environment][key]
	}
	// if value == "" && (key != KeyCaculateStartDate && key != KeyCaculateEndDate) {
	// 	logger.Panic("config must be not empty", logger.String("key", key))
	// }
	// if key == KeyDbUser || key == KeyDbPwd {
	// 	logger.Info("config", logger.Bool(key+" is empty", value == ""))
	// } else {
	// 	logger.Info("config", logger.String(key, value))
	// }

	return value
}

func getEnvInt(key string, environment string) int {
	var value string
	if v, ok := os.LookupEnv(key); ok {
		value = v
	} else {
		if DefaultEnvironment == EnvironmentStage || DefaultEnvironment == EnvironmentProd {
			logger.Panic("config is not able to use default config", logger.String("Environment", DefaultEnvironment))
		}
		value = defaultConfig[environment][key]
	}

	i, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		logger.Panic("config must be not empty", logger.String("key", key))
	}
	logger.Info("config", logger.Int64(key, i))
	return int(i)
}
