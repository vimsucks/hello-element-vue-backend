package config

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

type config struct {
	Host          string
	Port          int
	SQLiteFile    string
}

type flg struct {
	Help bool
	ConfigFile string
}

var Conf config
var Flag flg


func InitFlagAndConfig() {
	flag.BoolVar(&Flag.Help, "h", false, "打印帮助")
	flag.StringVar(&Flag.ConfigFile, "c", "./config.toml", "配置文件路径")
	flag.Parse()
	if Flag.Help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	if len(Flag.ConfigFile) != 0 {
		var bytes []byte
		var err error
		if bytes, err = ioutil.ReadFile(Flag.ConfigFile); err != nil {
			fmt.Println(err)
			flag.PrintDefaults()
			os.Exit(1)
		}
		if _, err = toml.Decode(string(bytes), &Conf); err != nil {
			fmt.Println(err)
			flag.PrintDefaults()
			os.Exit(1)
		}
	}
}
