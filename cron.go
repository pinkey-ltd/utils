package utils

import "github.com/robfig/cron/v3"

// NewCorn .
func NewCorn() *cron.Cron {
	return cron.New()
}
