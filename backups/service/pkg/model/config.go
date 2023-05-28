package model

type Config struct {
	Db struct {
		DbName     string `yaml:"dbName"`
		DbPort     string `yaml:"dbPort"`
		DbHost     string `yaml:"dbHost"`
		DbUsername string `yaml:"dbUsername"`
		DbPassword string `yaml:"dbPassword"`
		Path       string `yaml:"path"`
	} `yaml:"DB"`
	Minio struct {
		MinioEndpoint  string `yaml:"minioEndpoint"`
		MinioPort      string `yaml:"minioPort"`
		MinioAccessKey string `yaml:"minioAccessKey"`
		MinioSecretKey string `yaml:"minioSecretKey"`
		MinioSSL       string `yaml:"minioSSL"`
		MinioBucket    string `yaml:"minioBucket"`
	} `yaml:"Minio"`
	Server struct {
		Port string `yaml:"Port"`
	} `yaml:"Server"`
}
