package infrastructure

type config struct {
	ServiceName string `envconfig:"SERVICE_NAME"`
	ServerPort  int    `envconfig:"SERVER_PORT"`
	Credential  struct {
		Path string `envconfig:"CREDENTIAL_PATH"`
	}
	Quota struct {
		InitPerMonth int `envconfig:"QUOTA_INIT_PER_MONTH"`
	}
	DataBase struct {
		TimeOutSecond int    `envconfig:"MONGO_DB_TIMEOUT_SECOND"`
		Url           string `envconfig:"MONGO_DB_URL"`
		Name          string `envconfig:"MONGO_DB_NAME"`
		Collection    struct {
			Quota    string
			LogError string
			LogHTTP  string
		}
	}
}
