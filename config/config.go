package config

import "os"

const (
	NEW       = "new"
	CANCELLED = "cancelled"
	FINISHED  = "finished"
	PROD      = "production"
)

func IsProduction() bool {
	return os.Getenv("ENVIRONMENT") == PROD
}
