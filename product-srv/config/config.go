package config

type Mysql struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}
type Redis struct {
	Addr     string
	Password string
	Database int
}
type AddConfig struct {
	Mysql Mysql
	Redis Redis
}
