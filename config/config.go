package config

import "log"

type Config struct {
	Server struct {
	}
	Services struct {
		About struct {
			DB struct {
				URL      string
				UserName string
				Password string
				Protocol string
				DBName   string
			}
			MaxQuerySize int `json:"max_query_size"`
		}
		Calendar struct {
		}
	}
}

//TODO:添加config文件初始化，读取失败panic

func InitConfig() Config {
	exist := haveConfig()
	if !exist {
		initNewConfigFile()
		log.Fatal("file doesn't exist, initializing new config file. please fill up configs then restart the server.")
	}

	config, ok := readFromFiles()
	if !ok {
		log.Fatal("Invalid Config")
	}
	return config
}

func readFromFiles() (Config, bool) {
	return Config{}, false
}

func initNewConfigFile() {

}

func haveConfig() bool {
	return false
}
