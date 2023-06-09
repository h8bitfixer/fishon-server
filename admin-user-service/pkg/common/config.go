package config

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

var (
	_, b, _, _ = runtime.Caller(0)
	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../../..")
)

var Config config

type config struct {
	Environment string `yaml:"environment"`

	JwtAuthKeyUserID string `yaml:"jwt-auth-key-user-id"`
	JwtAuthKeyPhone  string `yaml:"jwt-auth-key-phone"`

	AdminUserServiceIP   string `yaml:"admin-user-service-ip"`
	AdminUserServicePort string `yaml:"admin-user-service-port"`

	ServerIP       string `yaml:"serverip"`
	IsSkipDatabase bool   `yaml:"is_skip_database"`

	RpcRegisterIP string `yaml:"rpcRegisterIP"`
	ListenIP      string `yaml:"listenIP"`

	ServerVersion string `yaml:"serverversion"`
	Api           struct {
		GinPort  []int  `yaml:"openImApiPort"`
		ListenIP string `yaml:"listenIP"`
	}
	CmsApi struct {
		GinPort  []int  `yaml:"openImCmsApiPort"`
		ListenIP string `yaml:"listenIP"`
	}

	Mysql struct {
		DBAddress      []string `yaml:"dbMysqlAddress"`
		DBUserName     string   `yaml:"dbMysqlUserName"`
		DBPassword     string   `yaml:"dbMysqlPassword"`
		DBDatabaseName string   `yaml:"dbMysqlDatabaseName"`
		DBTableName    string   `yaml:"DBTableName"`
		DBMsgTableNum  int      `yaml:"dbMsgTableNum"`
		DBMaxOpenConns int      `yaml:"dbMaxOpenConns"`
		DBMaxIdleConns int      `yaml:"dbMaxIdleConns"`
		DBMaxLifeTime  int      `yaml:"dbMaxLifeTime"`
	}
	Mongo struct {
		DBUri               string `yaml:"dbUri"`
		DBAddress           string `yaml:"dbAddress"`
		DBDirect            bool   `yaml:"dbDirect"`
		DBTimeout           int    `yaml:"dbTimeout"`
		DBDatabase          string `yaml:"dbDatabase"`
		DBSource            string `yaml:"dbSource"`
		DBUserName          string `yaml:"dbUserName"`
		DBPassword          string `yaml:"dbPassword"`
		DBMaxPoolSize       int    `yaml:"dbMaxPoolSize"`
		DBRetainChatRecords int    `yaml:"dbRetainChatRecords"`
	}
	Redis struct {
		DBAddress     []string `yaml:"dbAddress"`
		DBMaxIdle     int      `yaml:"dbMaxIdle"`
		DBMaxActive   int      `yaml:"dbMaxActive"`
		DBIdleTimeout int      `yaml:"dbIdleTimeout"`
		DBPassWord    string   `yaml:"dbPassWord"`
		EnableCluster bool     `yaml:"enableCluster"`
	}

	Etcd struct {
		EtcdSchema string   `yaml:"etcdSchema"`
		EtcdAddr   []string `yaml:"etcdAddr"`
	}
	Secret string `yaml:"secret"`
	//
}

func init() {
	cfgName := os.Getenv("CONFIG_NAME")
	if len(cfgName) == 0 {
		dir, err := os.Getwd()
		if err != nil {
			return
		}

		// Construct the file path relative to the current directory
		cfgName = path.Join(dir, "/pkg/common/config.yaml")
	}

	bytes, err := ioutil.ReadFile(cfgName)
	if err == nil {
		if err = yaml.Unmarshal(bytes, &Config); err != nil {
			panic(err.Error())
		}
	}
	//if err != nil {
	//	panic(err.Error())
	//}
	//if err = yaml.Unmarshal(bytes, &Config); err != nil {
	//	panic(err.Error())
	//}
}
