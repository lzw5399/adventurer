/**
 * @Author: lzw5399
 * @Date: 2020/7/30 13:30
 * @Desc: config model
 */
package config

type Config struct {
	System System `yaml:"system"`
	Db     Db     `yaml:"db"`
	Auth   Auth   `yaml:"auth"`
	Log    Log    `yaml:"log"`
}

type System struct {
	Locate string `yaml:"locate"`
}

type Db struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	InitialDb   string `yaml:"initialDb"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	MaxIdleConn int    `yaml:"max-idle-conn"`
	MaxOpenConn int    `yaml:"max-open-conn"`
	Logmode     bool   `yaml:"logmode"`
}

type Auth struct {
	Secret        string `yaml:"secret"`
	Issuer        string `yaml:"issuer"`
	Audience      string `yaml:"audience"`
	ExpireMinutes int    `yaml:"expireMinutes"`
}

type Log struct {
	Prefix  string `yaml:"prefix"`
	LogFile bool   `yaml:"log-file"`
	Stdout  string `yaml:"stdout"`
	File    string `yaml:"file"`
}