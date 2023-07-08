package config

import (
	"encoding/json"
	"flag"
	"fmt"

	"os"
	"path"
	"runtime"
	"sync"

	"github.com/10n1s-backend/cmd/auth"
	"github.com/10n1s-backend/cmd/game"
	"github.com/10n1s-backend/cmd/group"
	"github.com/10n1s-backend/cmd/rank"
	"github.com/10n1s-backend/cmd/repository"
	"github.com/10n1s-backend/cmd/route"
	"github.com/10n1s-backend/cmd/user"
	"github.com/10n1s-backend/pkg/config"
)

const (
	configName     = "tennis"
	testConfigName = "tennis-test"
)

var localPath = "." //set as environment later

var (
	tennisConfig = &TennisConfig{}
	once         sync.Once
)

type TennisConfig struct {
	RouteConfig      route.Config      `config:"route"`
	RepositoryConfig repository.Config `config:"repository"`
	AuthConfig       auth.Config       `config:"auth"`
	GameConfig       game.Config       `config:"game"`
	GroupConfig      group.Config      `config:"group"`
	RankConfig       rank.Config       `config:"rank"`
	UserConfig       user.Config       `config:"user"`
}

func Get(configFilePath string) *TennisConfig {
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
