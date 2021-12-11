package config

//type ConfigPathFile string

type Config struct {
	PathFile string
	Timeout  int
}

//func NewConfigFilePath(c *Config) *Config{
//	return &Config{
//		PathFile: ConfigPathFile(c.PathFile),
//	}
//}
//
//func (c Config) CSVFILE() ConfigPathFile{
//	return c.PathFile
//}
