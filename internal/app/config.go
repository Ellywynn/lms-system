package app

type config struct {
	Port string
}

func NewAppConfig() *config {
	return &config{
		Port: ":3000",
	}
}
