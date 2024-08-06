package utils

import (
	"errors"
	"os"
	"os/user"
	"runtime"
)


func GetCurrentUserHomeDir() (homeDir string, err error){
    usr, err := user.Current()
    if err != nil {
        return "", errors.New("failed to get user homedir.") 
    }

    homeDir = usr.HomeDir 
    return homeDir, nil
}


func GetCurrentUserName() (userName string, err error) {
    usr, err := user.Current()
    if err != nil {
        return "", errors.New("failed to get username.")
    }

    userName = usr.Username
    return userName, nil
}

func EnsureDirectory(path string) error {
    err := os.MkdirAll(path, os.ModePerm)
    if err != nil {
        return errors.New("failed to create folder.")
    }
    return nil
}


func GetOSFamily() (os string) {
    os = runtime.GOOS
    return
}
