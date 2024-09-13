package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env         string     `yaml:"env" env-default:"local"` //если мы парсим ямл файл, то у него будет такое название,по умолчанию будет локал, если пустой
	StoragePath string     `yaml:"storage_path" env-reguired:"./data"`
	GRPC        GRPCConfig `yaml:"grpc"`
}
type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config { //Must означает, что функция просто упадет и не вызовет ошибку и соотв вобще все упадет
	path := fetchConfigPath() //получили путь для конфиг файла
	if path == "" {
		panic("config file path is empty")
	}
	return MustLoadPath(path)
}

func MustLoadPath(configPath string) *Config { // это вызывают тесты (в видео MustLoadByPath называется)
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) { //почему тут пакет OS?
		panic("config file does not exist: " + configPath)
	}

	var cfg Config
	//configPath -- путь до файла, мы заходим и считываем оттуда инфу согласно тегам и записывае в cfg
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}

//Запустить конфиг файл можно через указание переменной окружения (CONFIG_PATH = ./что-то там), или можно написать sso -- config=./path...
