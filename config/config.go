package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	gcfg "gopkg.in/gcfg.v1"
)

// MainConfig : secret file struct
type MainConfig struct {
	Server struct {
		BaseURL    string
		PublicPort string
	}

	Cors struct {
		AllowedOrigins     []string
		AllowedHeaders     []string
		AllowCredentials   bool
		Debug              bool
		MaxAge             int
		OptionsPassthrough bool
	}
}

// Environment List
const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

// ReadConfig : function to read secret file
func ReadConfig(cfg interface{}, module string) interface{} {
	env := os.Getenv("BASICGOENV")
	if env == "" {
		env = EnvDevelopment
	}

	fileName := fmt.Sprintf("%s.%s.ini", module, env)
	configFileDir := os.Getenv("CONFIG_FILE_DIR")
	if len(configFileDir) == 0 {
		if env == EnvDevelopment {
			configFileDir = filepath.Join("files", "etc", "basic_golang") // default if CONFIG_FILE_DIR is not set
		} else {
			configFileDir = "/opt/basic_golang/files/etc"
		}
	}
	if configFileDir[len(configFileDir)-1] != os.PathSeparator {
		configFileDir += string(os.PathSeparator)
	}
	filePath := configFileDir + fileName

	if err := gcfg.ReadFileInto(cfg, filePath); err != nil {
		log.Fatal(err)
	}
	return cfg
}

var cfg *MainConfig

func LoadMainConfig() MainConfig {
	if cfg == nil {
		cfg = &MainConfig{}
		ReadConfig(cfg, "main")
	}
	return *cfg
}
