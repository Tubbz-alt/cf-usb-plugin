package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudfoundry/cli/cf/configuration/config_helpers"
)

type ConfigInterface interface {
	GetTarget() (string, error)
	SetTarget(string) error
}

func GetTarget() (target string, err error) {
	jsonConf, err := ioutil.ReadFile(getUsbConfigFile())
	if err != nil {
		return "", err
	}

	config := NewData()

	err = config.ReadConfigData(jsonConf)
	if err != nil {
		return "", err
	}

	return config.MgmtTarget, nil
}

func SetTarget(target string) (err error) {
	if !strings.Contains(target, "http") {
		target = fmt.Sprintf("http://%[1]s", target)
	}

	file, err := os.OpenFile(getUsbConfigFile(), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	defer file.Close()

	config := NewData()
	config.MgmtTarget = target

	output, err := config.SetConfigData()
	if err != nil {
		return err
	}

	_, err = file.Write(output)
	if err != nil {
		return err
	}

	return nil
}

func getUsbConfigFile() string {
	return filepath.Join(filepath.Dir(config_helpers.DefaultFilePath()), "usb-config.json")
}
