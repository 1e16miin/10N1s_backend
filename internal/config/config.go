package config

import (
	"encoding/json"
	"flag"
	"fmt"

	"os"
	"path"
	"runtime"
	"sync"

	"github.com/10n1s-backend/internal/room"
	"github.com/10n1s-backend/pkg/cache"
	"github.com/10n1s-backend/pkg/config"
	"github.com/10n1s-backend/pkg/database"
)

const (
	configName     = "tennis"
	testConfigName = "tennis-test"
)

var localPath = "." //set as environment later

var (
	tennisConfig = &Config{}
	once         sync.Once
)

type ServerConfig struct {
	Port string `config:"port"`
}

type Config struct {
	ServerConfig   ServerConfig    `config:"server"`
	DatabaseConfig database.Config `config:"database"`
	CacheConfig    cache.Config    `config:"cache"`
	RoomConfig     room.Config     `config:"room"`
}

func Get(configFilePath string) *Config {
	once.Do(func() {
		load(configFilePath)
	})
	return tennisConfig
}

func load(configPath string) {
	var configPaths []string

	if configPath != "" {
		configPaths = append(configPaths, configPath)
	}

	configPaths = append(configPaths, localPath)

	defaultConfig := defaultConfig
	configName := configName

	// for test code
	if flag.Lookup("test.v") != nil {
		defaultConfig = testDefaultConfig
		configName = testConfigName
		configPaths = append(configPaths, getProjectPath())
	}

	cfg := config.ReadConfigFile(defaultConfig, configName, configPaths, false, true)
	if err := config.UnmarshalConfig(cfg, tennisConfig); err != nil {
		panic(fmt.Errorf("failed to unmarshal configurations : %e", err))
	}

	configJson, err := jsonPretty(tennisConfig)
	if err != nil {
		panic(fmt.Errorf("cannot marshal config to json : %e", err))
	}

	fmt.Printf("10n1s Configurations: %s\n", string(configJson))
}

func jsonPretty(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func getProjectPath() string {
	_, filename, _, _ := runtime.Caller(0)

	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		fmt.Printf("fail to get project path : %e\n", err)
	}

	return dir
}
