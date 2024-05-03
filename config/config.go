package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigInstance Config

var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.0.0"

type Config struct {
	PushPlusToken  string `yaml:"pushplus_token"`
	RefreshToken   string `yaml:"refresh_token"`
	BilibiliCookie string `yaml:"bilibili_cookie"`
	KKCookie       string `yaml:"kk_cookie"`_UP_A4A_11_=wb965177c9c84da384f1f3eac1bf8cdd; _UP_D_=pc; __pus=e3a9f973072212b6c49dcf601859e4ebAAT1/lKKJ0KnXhk9hmmdeePbPWPWx+5qBT0akQS8labwfa2XnMrwSEJIJaPAlHowwGHZPhSBQRi4QwQMHHi+PSo+; __kp=af0f0aa0-0910-11ef-ae38-f3e2d29e3ac7; __kps=AAQugX1apwiRNz/Cg+r3s8PN; __ktd=R2mWNClAFTsskuCL9rB9jg==; __uid=AAQugX1apwiRNz/Cg+r3s8PN; __puus=73a3fc3828b038439b21b4d70ede0d1aAARh/j5LNF2TRXFvVl27WmMKnHQoli8CoQfmgPvHD5gUV5NB213wY3JZZoAELDkrnpqls8Yk5i4MVcVHlINklNHR7599/BcundF9IAtJ5YyENaV25XvRSsUEKwbkUFQqDQ9J5i49bmhCPgxb/Fa6Bq2h7KkkTjZaLM9IQlYQTcM5K9JR17smg0+anierlmJI27eGonHdfzFdeUasY8N/rt8F
	JdCookie       string `yaml:"jd_cookie"`
}

func init() {
	LoadConfig()
}

func LoadConfig() {
	confFIle, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err.Error())
	}
	config := Config{}
	yaml.Unmarshal(confFIle, &config)
	ConfigInstance = config
}
