package conf

// 组合全部配置模型
type Config struct {
	Server       Server       `mapstructure:"server"`
	Mysql        Mysql        `mapstructure:"mysql"`
	Redis        Redis        `mapstructure:"Redis"`
	Es			 Es           `mapstructure:"Es"`
	Jwt			Jwt           `mapstructure:"Jwt"`
	
}

// 服务启动端口号配置
type Server struct {
	Post string `mapstructure:"post"`
}

// MySQL数据源配置
type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Url      string `mapstructure:"url"`
}
// Redis 配置
type Redis struct{
	RedisDb      string `mapstructure:"RedisDb"`
	RedisAddr    string `mapstructure:"RedisAddr"`
	RedisPw		 string `mapstructure:"RedisPw"`
	RedisDbName  string `mapstructure:"RedisName"`
}
type Es struct{
	EsHost	string `mapstructure:"EsHost"`
	EsPort	string `mapstructure:"EsPort"`
	EsIndex	string `mapstructure:"EsIndex"`
}



// 文件上传相关路径配置
type Upload struct {
	SavePath  string `mapstructure:"savePath"`
	AccessUrl string `mapstructure:"accessUrl"`
}

// 用户认证配置
type Jwt struct {
	SigningKey string `mapstructure:"signingKey"`
}

// 微信小程序相关配置
type Code2Session struct {
	Code      string `mapstructure:"code"`
	AppId     string `mapstructure:"appId"`
	AppSecret string `mapstructure:"appSecret"`
}
