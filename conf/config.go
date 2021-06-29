package conf

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/cosmos-gaminghub/explorer-backend/logger"
	"github.com/joho/godotenv"
)

const (
	KeyDbAddr      = "DB_ADDR"
	KeyDATABASE    = "DB_DATABASE"
	KeyDbUser      = "DB_USER"
	KeyDbPwd       = "DB_PASSWORD"
	KeyDbPoolLimit = "DB_POOL_LIMIT"

	KeyAddrHubLcd = "ADDR_NODE_SERVER"

	KeyCoinMarketApi       = "COINMARKET_API_KEY"
	KeyCoinMarketEndPoint  = "COINMARKET_API_ENDPOINT"
	KeyCoinMarketApiVerion = "COINMARKET_API_VERSION"

	KeyPrefixAccAddr  = "cosmos"
	KeyPrefixAccPub   = "cosmospub"
	KeyPrefixValAddr  = "cosmosvaloper"
	KeyPrefixValPub   = "cosmosvaloperpub"
	KeyPrefixConsAddr = "cosmosvalcons"
	KeyPrefixConsPub  = "cosmosvalconspub"

	EnvironmentDevelop = ".env"
	DefaultEnvironment = EnvironmentDevelop
)

var (
	config Config
)

func init() {
	logger.Info("==================================load config start==================================")
	addrs := strings.Split(getEnv(KeyDbAddr, DefaultEnvironment), ",")
	db := dbConf{
		Addrs:     addrs,
		Database:  getEnv(KeyDATABASE, DefaultEnvironment),
		UserName:  getEnv(KeyDbUser, DefaultEnvironment),
		Password:  getEnv(KeyDbPwd, DefaultEnvironment),
		PoolLimit: getEnvInt(KeyDbPoolLimit, DefaultEnvironment),
	}
	config.Db = db

	coinMarket := CoinMarket{
		EndPoint: getEnv(KeyCoinMarketEndPoint, DefaultEnvironment),
		Version:  getEnv(KeyCoinMarketApiVerion, DefaultEnvironment),
		Apikey:   getEnv(KeyCoinMarketApi, DefaultEnvironment),
	}
	config.CoinMarket = coinMarket

	hubcf := hubConf{
		LcdUrl: getEnv(KeyAddrHubLcd, DefaultEnvironment),
	}
	config.Hub = hubcf
	logger.Info("==================================load config end==================================")
}

func Get() Config {
	return config
}

type Config struct {
	Db         dbConf
	CoinMarket CoinMarket
	Hub        hubConf
}

type hubConf struct {
	LcdUrl string
}

type CoinMarket struct {
	EndPoint string
	Version  string
	Apikey   string
}
type dbConf struct {
	Addrs     []string
	Database  string
	UserName  string
	Password  string
	PoolLimit int
}

func getEnv(key string, environment string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func getEnvInt(key string, environment string) int {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Fatalf("Error convert %s to string", key)
	}
	return value
}
