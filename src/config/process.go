package config

import "os"

var Mode = os.Getenv("MODE")

var Environment = os.Getenv("Environment")
