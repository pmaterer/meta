package config

type Config struct {
	ServerListenAddress string `default:"localhost"`
	ServerListenPort    int64  `default:"9999"`
	DatabaseName        string `required:"true" split_words:"true"`
	DatabaseUser        string `required:"true" split_words:"true"`
	DatabasePassword    string `required:"true" split_words:"true"`
	DatabasePort        int64  `default:"5432"`
	DatabaseHost        string `default:"localhost"`
	DatabaseSSLMode     string `default:"disable"`
}
