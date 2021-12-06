package config

import "os"

var ProjectId = os.Getenv("GOOGLE_CLOUD_PROJECT")
