/**
 * @Author: lzw5399
 * @Date: 2020/7/30 13:30
 * @Desc: config model
 */
package config

type Config struct {
	Db     Db     `yaml:"db"`
	Auth   Auth   `yaml:"auth"`
	System System `yaml:"system"`
}

type Db struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	InitialDb   string `yaml:"initialDb"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	MaxIdleConn int    `yaml:"max-idle-conn"`
	MaxOpenConn int    `yaml:"max-open-conn"`
	LogMode     bool   `yaml:"logmode"`
}

type Auth struct {
	Secret        string `yaml:"secret"`
	Issuer        string `yaml:"issuer"`
	Audience      string `yaml:"audience"`
	ExpireMinutes int    `yaml:"expireMinutes"`
}

type System struct {
	Locate string `yaml:"locate"`
}
