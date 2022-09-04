package app

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// GetBackupFiles retrieves backup files
func GetBackupFiles() ([]os.FileInfo, error) {
	backupFolder := viper.GetString("backup.folder")

	files, err := ioutil.ReadDir(backupFolder)
	if err != nil {
		return nil, err
	}

	var backupFiles []os.FileInfo
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "$PASSWORD_MANAGER_NAME") && strings.HasSuffix(file.Name(), ".bak") {
			backupFiles = append(backupFiles, file)
		}
	}

	return backupFiles, nil
}
