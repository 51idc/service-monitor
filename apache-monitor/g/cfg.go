package g

import (
	"encoding/json"
	log "github.com/cihub/seelog"
	"os"
	"sync"

	"github.com/toolkits/file"
)

type ApacheConfig struct {
	Enabled bool   `json:"enabled"`
	Staturl string `json:"staturl"`
}

type SmartAPIConfig struct {
	Enabled bool   `json:"enabled"`
	Url     string `json:"url"`
}

type TransferConfig struct {
	Enabled  bool   `json:"enabled"`
	Addrs    []string `json:"addrs"`
	Interval int    `json:"interval"`
	Timeout  int    `json:"timeout"`
}

type HttpConfig struct {
	Enabled bool   `json:"enabled"`
	Listen  string `json:"listen"`
}

type GlobalConfig struct {
	Debug    bool            `json:"debug"`
	Hostname string          `json:"hostname"`
	Apache   *ApacheConfig   `json:"apache"`
	SmartAPI *SmartAPIConfig `json:"smartAPI`
	Transfer *TransferConfig `json:"transfer"`
	Http     *HttpConfig     `json:"http"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	lock       = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func Hostname() (string, error) {
	hostname := Config().Hostname
	if hostname != "" {
		return hostname, nil
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Info("ERROR: os.Hostname() fail", err)
	}
	return hostname, err
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Error("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Error("config file:", cfg, "is not existent. maybe you need `mv cfg.example.json cfg.json`")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Error("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Error("parse config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c

	log.Info("read config file:", cfg, "successfully")

}
