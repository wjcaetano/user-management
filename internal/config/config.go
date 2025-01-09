package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"

	"strings"
	"time"

	"github.com/magiconair/properties"
)

const (
	localConfigScope       = "resources/config/local.properties"
	applicationConfigScope = "resources/config/application.properties"
	scopeEnv               = "SCOPE"
	appPathEnv             = "APP_PATH"
	localScope             = "local"
)

type (
	HTTPClient struct {
		MaxOpenConns    int           `properties:"max_open_conns,default=10"`
		MaxIdleConns    int           `properties:"max_idle_conns,default=5"`
		ConnMaxLifetime time.Duration `properties:"conn_max_lifetime,default=10m"`
		Addr            string        `properties:"addr,default=:8080"`
	}

	Configuration struct {
		AppPath    string     `properties:"app_path,default="`
		Scope      string     `properties:"scope,default=local"`
		HTTPClient HTTPClient `properties:"http_client"`
		Database   Database   `properties:"database"`
	}

	Database struct {
		User     string
		Password string
		Host     string
		Port     string
		DBName   string
	}
)

func NewConfig() (*Configuration, error) {
	prop, err := loadProperties()
	if err != nil {
		return nil, err
	}

	conf, err := decodeConfig(prop)
	if err != nil {
		return nil, err
	}

	conf.overrideConfigurations()

	return conf, nil
}

func loadProperties() (*properties.Properties, error) {
	if err := checkMandatoryEnvs(); err != nil {
		return nil, err
	}

	if getEnv(scopeEnv, "SCOPE") == localScope {
		prop, err := loadLocalProperties()
		if err != nil {
			return nil, err
		}
		return prop, nil
	}
	return loadServiceProperties()
}

func checkMandatoryEnvs() error {
	mandatoryEnvs := [...]string{appPathEnv, scopeEnv}

	for _, env := range mandatoryEnvs {
		if _, ok := os.LookupEnv(env); !ok {
			return fmt.Errorf("environment variable %s not set", env)
		}
	}
	return nil
}

func (c *Configuration) overrideConfigurations() {
	c.AppPath = getEnv(appPathEnv, "")
	c.Scope = getEnv(scopeEnv, "")
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func loadLocalProperties() (*properties.Properties, error) {
	err := godotenv.Load("variables.env")
	if err != nil {
		return nil, fmt.Errorf("error loading env file: %w", err)
	}

	appPath, err := getProjectPath()
	if err != nil {
		return nil, fmt.Errorf("unable to get project path: %w", err)
	}

	configFile := filepath.Join(appPath, localConfigScope)

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("unable to find local config file: %s", configFile)
	}

	return properties.MustLoadFile(configFile, properties.UTF8), nil
}

func loadServiceProperties() (*properties.Properties, error) {
	inputConfig := os.Getenv("configFileName")
	if inputConfig == "" {
		inputConfig = applicationConfigScope
	}

	appPath, err := getProjectPath()
	if err != nil {
		return nil, fmt.Errorf("unable to get project path: %w", err)
	}

	configFile := filepath.Join(appPath, inputConfig)

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("unable to find config file: %s", configFile)
	}

	prop, _ := properties.LoadFile(configFile, properties.UTF8)

	return prop, nil
}

func decodeConfig(prop *properties.Properties) (*Configuration, error) {
	var cfg Configuration
	if err := prop.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func IsLocalScope() bool {
	return getEnv(scopeEnv, "SCOPE") == localScope
}

func getProjectPath() (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("unable to get working directory: %w", err)
	}

	path := setDefaultPath(workingDir)

	return path, nil
}

func setDefaultPath(workingDir string) string {
	for {
		if isProjectRoot(workingDir) {
			return workingDir
		}

		parentDir := getParentDir(workingDir)
		if parentDir == workingDir {
			break
		}

		workingDir = parentDir
	}
	return ""
}

func isProjectRoot(path string) bool {
	return strings.HasSuffix(path, "user-management") || fileExists(filepath.Join(path, "go.mod"))
}

func getParentDir(path string) string {
	return filepath.Dir(path)
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}
