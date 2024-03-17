package core

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

func GetOpenpackLocalStore() (string, error) {
	if //goland:noinspection GoBoolExpressions
	runtime.GOOS == "linux" {
		// Prefer $XDG_DATA_HOME over $XDG_CACHE_HOME
		dataHome := os.Getenv("XDG_DATA_HOME")
		if dataHome != "" {
			return filepath.Join(dataHome, "openpack"), nil
		}
	}
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userConfigDir, "openpack"), nil
}

func GetOpenpackLocalCache() (string, error) {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userCacheDir, "openpack"), nil
}

func GetOpenpackInstallBinPath() (string, error) {
	localStore, err := GetOpenpackLocalStore()
	if err != nil {
		return "", err
	}
	return filepath.Join(localStore, "bin"), nil
}

func GetOpenpackInstallBinFile() (string, error) {
	binPath, err := GetOpenpackInstallBinPath()
	if err != nil {
		return "", err
	}
	var exeName string
	if //goland:noinspection GoBoolExpressions
	runtime.GOOS == "windows" {
		exeName = "openpack.exe"
	} else {
		exeName = "openpack"
	}
	return filepath.Join(binPath, exeName), nil
}

func GetOpenpackCache() (string, error) {
	configuredCache := viper.GetString("cache.directory")
	if configuredCache != "" {
		return configuredCache, nil
	}
	localStore, err := GetOpenpackLocalCache()
	if err != nil {
		return "", err
	}
	return filepath.Join(localStore, "cache"), nil
}
