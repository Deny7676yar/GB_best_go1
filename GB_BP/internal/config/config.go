package config

//Config - структура для конфигурации
type Config struct {
	MaxDepth   int
	MaxResults int
	MaxErrors  int
	URL        string
	Timeout    int //in seconds
}
