package config
 
type Config struct {
	DB *DBConfig
}
 
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}
 
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "ujhn57abqqlauq4w",
			Password: "4D3xwIKjyfILPn0tEB5V",
			Name:     "headset",
			Charset:  "utf8",
		},
	}
}