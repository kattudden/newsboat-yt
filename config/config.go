package config

import (
	"errors"
	"fmt"
	"kattudden/newsboat-yt/utils"
	"path/filepath"
)


type Config struct {
    DatabaseDirectory string
    DatabaseFilename string
    DatabasePath string
    DownloadPath string
}


func New() (*Config, error) {
    databaseFileName := "db.sqlite"
    databaseDirectory := ".cache/newsboat-yt"
    downloadDirectory := "Videos/newsboat-yt"

    userHomeDir, err := utils.GetCurrentUserHomeDir()
    if err != nil {
        fmt.Println(err)
            return nil, errors.New("failed to generate config!") 
    }

    return &Config {
        DatabaseDirectory: databaseDirectory,
        DatabaseFilename: databaseFileName,
        DatabasePath: filepath.Join(userHomeDir, databaseDirectory, databaseFileName),
        DownloadPath: filepath.Join(userHomeDir, downloadDirectory),
    }, nil
}
