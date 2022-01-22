package utility

import (
	"os"
)

func GetEnv(Key, Default string) *string {
	if data := os.Getenv(Key); data != "" {
		return &data
	}
	return &Default
}
