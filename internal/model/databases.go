package model

type DATABASES struct {
	MySQL struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Charset  string `yaml:"charset"`
	} `yaml:"mysql"`

	REDIS struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		DB       int    `yaml:"db"`
		Password string `yaml:"password"`
		Timeout  int    `yaml:"timeout"`
		PoolSize int    `yaml:"pool_size"`
	} `yaml:"redis"`
}
