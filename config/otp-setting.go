package config

import "time"

var OtpLifeTime time.Duration

func init() {
	OtpLifeTime = 15 * time.Minute
}
