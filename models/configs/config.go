/**
 * @Author: lzw5399
 * @Date: 2020/7/30 13:30
 * @Desc: config models
 */
package configs

type Config struct {
	Db Db `yaml:"db"`
	Auth Auth `yaml:"auth"`
	Locate string `yaml:"locate"`
}

type Db struct {
	Hostname string `yaml:"hostname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Auth struct {
	Secret string `yaml:"secret"`
	Issuer string `yaml:"issuer"`
	Audience string `yaml:"audience"`
	ExpireMinutes int64 `yaml:"expireMinutes"`
}