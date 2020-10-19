package routers

//Config struct
type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

// Cfg returnet variable
var Cfg Config
