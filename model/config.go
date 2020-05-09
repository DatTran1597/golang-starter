package model

type Configuration struct {
	ServiceSettings ServiceSettings `yaml:"service_settings"`
	SQLSettings     SQLSettings     `yaml:"sql_settings"`
	CacheSettings   CacheSetting    `yaml:"cache_settings"`
	SearchSettings  SearchSetting   `yaml:"search_settings"`
}

type ServiceSettings struct {
	Port          string `yaml:"port"`
	ListenAddress string `yaml:"listen_address"`
}

type SQLSettings struct {
	DriverName string `yaml:"driver_name"`
	URI        string `yaml:"uri"`
	Timeout    int    `yaml:"timeout"`
}

type CacheSetting struct {
	URI      string `yaml:"uri"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
	Timeout  string `yaml:"timeout"`
}

type SearchSetting struct {
	ConnectionURL string `yaml:"uri"`
	UserName      string `yaml:"username"`
	Password      string `yaml:"password"`
	Sniff         bool   `yaml:"sniff"`
}
