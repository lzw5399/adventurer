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
	Hostname string `yaml:"hostname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
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