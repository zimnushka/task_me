package app

type DBConectionParams struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Url      string `json:"url"`
	Db       string `json:"db"`
}

type AppConfig struct {
	Debug    bool              `json:"debug"`
	DBParams DBConectionParams `json:"DBParams"`
}

var config AppConfig

func SetConfig(newConfig AppConfig) {
	config = newConfig
}

func GetConfig() AppConfig {
	return config
}

var DebugConfig = AppConfig{
	Debug: true,
	DBParams: DBConectionParams{
		User:     "root",
		Password: "43WYOH5l8W1I",
		Url:      "mariadb:3306",
		Db:       "taskMe",
	}}

var ReleaseConfig = AppConfig{
	Debug: true,
	DBParams: DBConectionParams{
		User:     "root",
		Password: "43WYOH5l8W1I",
		Url:      "localhost:3306",
		Db:       "taskMe",
	}}
