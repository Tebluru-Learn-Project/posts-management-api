package config

type AppConfig struct {
	AppName string
	AppPort string
	AppEnv string
}

func LoadAppConfig() AppConfig {
	return AppConfig{
		AppName: GetEnv("APP_NAME", "golang-api"),
		AppPort: GetEnv("APP_PORT", "8080"),
		AppEnv: GetEnv("APP_ENV", "development"),
	}
}