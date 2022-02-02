package config

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Services struct {
		DB struct {
			URL      string `json:"URL"`
			UserName string `json:"UserName"`
			Password string `json:"Password"`
			Protocol string `json:"Protocol"`
			DBName   string `json:"DBName"`
		}
	}
}

//TODO:添加config文件初始化，读取失败panic

func InitConfig() Config {
	config, ok := readFromFiles()
	log.Println("initalizing")
	if !ok {
		log.Println("Invalid Config")
		os.Exit(1)
	}
	return config
}

func readFromFiles() (Config, bool) {
	var config Config
	path := "./files/"
	log.Println(path)
	if !pathExists(path) {
		log.Println("exist")
		if err := os.MkdirAll(path, 0777); err != nil {
			log.Println("making path: ", err)
		}
	}
	if !pathExists(path + "config.json") {
		makeBlankConfig(path + "config.json")
		log.Println("made blank config")
		log.Println("generated new file, please fill in the config.")
		os.Exit(1)
		return Config{}, false
	}
	f, err := os.Open(path + "config.json")
	defer f.Close()
	if err != nil {

		log.Println("open file error")
		os.Exit(0)
	}
	raw, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("reading from file: ", err)
		os.Exit(1)
	}
	if err := json.Unmarshal(raw, &config); err != nil {
		log.Println("unmarshal from file: ", err)
		os.Exit(1)
	}
	return config, true
}

func makeBlankConfig(path string) {
	f, err := os.Create(path)
	defer f.Close()

	if err != nil {
		return
	}
	writer := bufio.NewWriter(f)
	var c Config
	c.Services.DB.URL = ""
	c.Services.DB.DBName = ""
	c.Services.DB.Protocol = ""
	c.Services.DB.Password = ""
	c.Services.DB.UserName = ""
	raw, err := json.Marshal(c)
	if err != nil {
		log.Fatal("writing config:", err)
	}
	_, err = writer.Write(raw)
	if err != nil {
		log.Fatal("writing config:", err)
	}
	writer.Flush()
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
