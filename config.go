package goelector

// Config type with input to elector
type Config struct {
	LeaseDuration int    `json:"lease_duration"`
	RenewDeadline int    `json:"renew_deadline"`
	RetryPeriod   int    `json:"retry_period"`
	Lock          string `json:"lock"`
	Namespace     string `json:"namespace"`
}

// GetDefaultConfig returns defaulted config
func GetDefaultConfig() *Config {
	return &Config{
		LeaseDuration: 15,
		RenewDeadline: 10,
		RetryPeriod:   2,
		Lock:          "elector-lock",
		Namespace:     "default",
	}
}
