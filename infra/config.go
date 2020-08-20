package infra

const (
	AppName = "restapi-example"
	AppPort = 3000
)

type Config struct {
	Mysql MysqlConfig
}

type MysqlConfig struct {
	DBName   string
	TCP      string
	Port     string
	User     string
	Password string
}

func NewConfig() *Config {
	return &Config{
		Mysql: MysqlConfig{
			DBName:   "mysql",
			TCP:      "tcp(test-mysql.cru040ksu4gu.sa-east-1.rds.amazonaws.com:3306)",
			User:     "oblue",
			Password: "oblue123",
		},
	}
}
