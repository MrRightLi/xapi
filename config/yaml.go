package config

type Yaml struct {
	Mysql struct {
		User     string `ymal:"user"`
		Host     string `ymal:"host"`
		Password string `ymal:"password"`
		Port     string `ymal:"port"`
		Name     string `ymal:"name"`
	}
	Cache struct {
		Enable bool     `ymal:"enable"`
		List   []string `ymal:"list,flow"`
	}
}

// Yaml1 struct of yaml
type Yaml1 struct {
	SQLConf   Mysql `yaml:"mysql"`
	CacheConf Cache `yaml:"cache"`
}

// Yaml2 struct of yaml
type Yaml2 struct {
	Mysql `yaml:"mysql,inline"`
	Cache `yaml:"cache,inline"`
}

// Mysql struct of mysql conf
type Mysql struct {
	User     string `yaml:"user"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

// Cache struct of cache conf
type Cache struct {
	Enable bool     `yaml:"enable"`
	List   []string `yaml:"list,flow"`
}
