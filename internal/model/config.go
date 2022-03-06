package model

type ServerSettings struct {
	ListenAddress string `yaml:"listen_address"`
}

func (s *ServerSettings) SetDefaults() {
	if len(s.ListenAddress) <= 0 {
		s.ListenAddress = ":8080"
	}
}

type Config struct {
	ServerSettings ServerSettings `yaml:"server_settings"`
}

func (c *Config) SetDefaults() {
	c.ServerSettings.SetDefaults()
}
