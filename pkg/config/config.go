package config

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
	"time"

	"github.com/10n1s-backend/pkg/parser"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

const (
	configTagName = "config"
	configType    = "yaml"
	projectPath   = "/src/10n1s"
)

var supportedExts = []string{"yaml", "yml"}
var errConfigNotLoaded = errors.New("configuration is not loaded")

func ReadConfigFile(defaultCfg, filename string, filepathList []string, addProjectPath, warn bool) *viper.Viper {
	conf := viper.New()
	conf.SetConfigType(configType)
	conf.SetConfigName(filename)

	// load default configurations
	if err := conf.ReadConfig(bytes.NewReader([]byte(defaultCfg))); err != nil {
		panic(err.Error() + "cannot read default configurations")
	}

	// copy filepath list
	newFilepathList := make([]string, len(filepathList))
	copy(newFilepathList, filepathList)

	// append a project path to filepath list if needed
	projectPath := getProjectPath()
	if addProjectPath && len(projectPath) > 0 {
		newFilepathList = append(newFilepathList, projectPath)
	}

	// merge a config file into default configurations
	mergeErr := errConfigNotLoaded

configLoop:
	for _, filepath := range newFilepathList {
		for _, ext := range supportedExts {
			configFile := path.Join(filepath, fmt.Sprintf("%s.%s", filename, ext))
			conf.SetConfigFile(configFile)

			mergeErr = conf.MergeInConfig()
			if errors.Is(mergeErr, os.ErrNotExist) {
				// config file does not exist
				continue
			}

			if mergeErr != nil {
				// config file existed, but failed to parse it
				panic(fmt.Sprintf("failed to parse %s: %s", configFile, mergeErr))
			}

			// succeeded to parse a config file
			fmt.Printf("%s configuration file has been loaded.\n", configFile)
			break configLoop
		}
	}

	// print some warning messages if failed to merge configs
	if mergeErr != nil && warn {
		fmt.Printf("Run with default configs, because it failed to load a user configuration file: %s\n", mergeErr)
		fmt.Printf(" - config filename: %s\n", filename)
		fmt.Printf(" - supported extensions: %s\n", supportedExts)
		fmt.Printf(" - config path list: %s\n", newFilepathList)
	}

	return conf
}

func UnmarshalConfig(cfg *viper.Viper, rawVal interface{}) error {
	// try to unmarshal with ErrorUnused = true
	err := cfg.Unmarshal(rawVal, getConfigDecoderOption(true))
	if err == nil {
		return nil
	}

	// warning message
	fmt.Printf("failed to unmarshal configurations: %s\n", err.Error())
	fmt.Println("retry to unmarshal configurations with 'ErrorUnused = false'")

	// try to unmarshal with ErrorUnused = false (this is for backward compatibility with past configurations)
	err = cfg.Unmarshal(rawVal, getConfigDecoderOption(false))
	if err != nil {
		return err
	}

	fmt.Println("succeeded to unmarshal configurations with 'ErrorUnused = false'")
	return nil
}

func getConfigDecoderOption(errorUnused bool) viper.DecoderConfigOption {
	return func(decoder *mapstructure.DecoderConfig) {
		decoder.TagName = configTagName
		decoder.ErrorUnused = errorUnused
		decoder.WeaklyTypedInput = true
		decoder.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			stringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		)
	}
}

// stringToTimeDurationHookFunc returns a DecodeHookFunc that converts strings to time.Duration.
func stringToTimeDurationHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(time.Duration(5)) {
			return data, nil
		}

		// Convert it by parsing
		return parser.ParseDuration(data.(string))
	}
}

// for finding a config file under a project directory (for test code)
func getProjectPath() string {
	if goPath, ok := os.LookupEnv("GOPATH"); ok {
		return path.Join(goPath, projectPath)
	}
	return ""
}
