package conf

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/cosmos-gaminghub/explorer-backend/logger"
	"github.com/joho/godotenv"
)

const (
	KeyDbAddr         = "DB_ADDR"
	KeyDATABASE       = "DB_DATABASE"
	KeyDbUser         = "DB_USER"
	KeyDbPwd          = "DB_PASSWORD"
	KeyDbPoolLimit    = "DB_POOL_LIMIT"
	KeyFastSyn        = "FAST_SYNC"
	KeySyncFromHeight = "SYNC_FROM_HEIGHT"
	KeyAddressPrefix  = "ADDRESS_PREFIX"

	KeyAddrHubLcd = "ADDR_NODE_SERVER"
	KeyHubRPC     = "RPC_ENPOINT"
	KeyCoin       = "DEFAULT_COIN"

	KeyCoingeckoEndPoint  = "COINGECKO_API_ENDPOINT"
	KeyCoingeckoApiVerion = "COINGECKO_API_VERSION"
	KeyCoingeckoCurrency  = "COINGECKO_CURRENCY"

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
		Addrs:          addrs,
		Database:       getEnv(KeyDATABASE, DefaultEnvironment),
		UserName:       getEnv(KeyDbUser, DefaultEnvironment),
		Password:       getEnv(KeyDbPwd, DefaultEnvironment),
		PoolLimit:      getEnvInt(KeyDbPoolLimit, DefaultEnvironment),
		FastSync:       getEnv(KeyFastSyn, DefaultEnvironment),
		SyncFromHeight: getEnvInt(KeySyncFromHeight, DefaultEnvironment),
		AddresPrefix:   getEnv(KeyAddressPrefix, DefaultEnvironment),
	}
	config.Db = db

	coingecko := Coingecko{
		EndPoint: getEnv(KeyCoingeckoEndPoint, DefaultEnvironment),
		Version:  getEnv(KeyCoingeckoApiVerion, DefaultEnvironment),
		Currency: getEnv(KeyCoingeckoCurrency, DefaultEnvironment),
	}
	config.Coingecko = coingecko

	hubcf := hubConf{
		LcdUrl: getEnv(KeyAddrHubLcd, DefaultEnvironment),
		RpcUrl: getEnv(KeyHubRPC, DefaultEnvironment),
		Coin:   getEnv(KeyCoin, DefaultEnvironment),
	}
	config.Hub = hubcf
	logger.Info("==================================load config end==================================")
}

func Get() Config {
	return config
}

type Config struct {
	Db        dbConf
	Coingecko Coingecko
	Hub       hubConf
}

type hubConf struct {
	LcdUrl string
	RpcUrl string
	Coin   string
}

type Coingecko struct {
	EndPoint string
	Version  string
	Currency string
}
type dbConf struct {
	Addrs          []string
	Database       string
	UserName       string
	Password       string
	PoolLimit      int
	FastSync       string
	SyncFromHeight int
	AddresPrefix   string
}

func getEnv(key string, environment string) string {
	_, b, _, _ := runtime.Caller(0)
	err := godotenv.Load(filepath.Join(filepath.Dir(b), "../"+environment))

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func getEnvInt(key string, environment string) int {
	_, b, _, _ := runtime.Caller(0)
	err := godotenv.Load(filepath.Join(filepath.Dir(b), "../"+environment))

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Fatalf("Error convert %s to string", key)
	}
	return value
}
