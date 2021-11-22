package database

type _ssh struct {
	// 地址
	Addr string `json:"addr" yaml:"addr" xml:"addr" toml:"addr" validate:"required,hostname_port"`
	// 用户名
	Username string `json:"username" yaml:"username" xml:"username" toml:"username"`
	// 密码
	Password string `json:"password" yaml:"password" xml:"password" toml:"password" validate:"required_without=Keyfile"`
	// 私钥文件地址
	Keyfile string `json:"keyfile" yaml:"keyfile" xml:"keyfile" toml:"keyfile" validate:"required_without=Password"`
}
