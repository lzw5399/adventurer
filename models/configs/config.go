/**
 * @Author: lzw5399
 * @Date: 2020/7/30 13:30
 * @Desc: config models
 */
package configs

type Config struct {
	Db  Db       `yaml:"db"`
	No  No       `yaml:"no"`
	Str string   `yaml:"str"`
	Arr []string `yaml:"arr"`
}

type Db struct {
	Hostname string `yaml:"hostname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type No struct {
	A int    `yaml:"a"`
	B bool   `yaml:"b"`
	C string `yaml:"c"`
}