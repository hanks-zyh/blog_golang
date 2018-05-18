package conf

type Server struct {
	Port string
}

type Mysql struct {
	Address  string
	Port     string
	DbName   string
	Username string
	Password string
}

type Config struct {
	Mode   string
	Server Server
	Mysql  Mysql
}

var product = Config{
	Mode: "release",
	Server: Server{
		Port: ":9345",
	},
	Mysql: Mysql{
		Address:  "127.0.0.1",
		Port:     "3306",
		DbName:   "blog_go",
		Username: "hanks",
		Password: "123456",
	},
}

var dev = Config{
	Mode: "debug",
	Server: Server{
		Port: ":9345",
	},
	Mysql: Mysql{
		Address:  "127.0.0.1",
		Port:     "3306",
		DbName:   "blog_go",
		Username: "hanks",
		Password: "123456",
	},
}

var AppConfig *Config

func init() {
	AppConfig = &product
}
