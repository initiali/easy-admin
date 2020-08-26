package utils

import "gopkg.in/ini.v1"
import "os"

var (
AppMode string
AppAddress string

DbHost string
DbPort string
DbUser string
DbPass string
DbName string
DbDriver string

)


func init() {
    cfg, err := ini.Load("config.ini")
    if err != nil {
        os.Exit(1)
    }
    initServer(cfg)
    initDatabase(cfg)
}

func initServer(cfg *ini.File) error {
	AppMode = cfg.Section("App").Key("level").MustString("debug")
	AppAddress = cfg.Section("App").Key("listen").MustString("127.0.0.1:8080")
    	return nil
}

func initDatabase(cfg *ini.File) error {
    DbHost = cfg.Section("Database").Key("host").MustString("127.0.0.1")
    DbPort = cfg.Section("Database").Key("port").MustString("3306")
    DbUser = cfg.Section("Database").Key("user").MustString("root")
    DbPass = cfg.Section("Database").Key("password").MustString("password")
    DbName = cfg.Section("Database").Key("dbname").MustString("easy-admin")
    DbDriver = cfg.Section("Database").Key("driver").MustString("mysql")
    return nil
}
