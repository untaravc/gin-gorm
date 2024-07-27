package app_config

import "os"

var PORT = ":8000"
var STATIC_ROUTE = "/public"
var STATIC_DIR = "./public"

func InitAppConfig() {
	portEnv := os.Getenv("APP_PORT")
	staticRouteEnv := os.Getenv("STATIC_ROUTE")
	staticDirEnv := os.Getenv("STATIC_DIR")

	if portEnv != "" {
		PORT = portEnv
	}

	if staticRouteEnv != "" {
		STATIC_ROUTE = staticRouteEnv
	}

	if staticDirEnv != "" {
		STATIC_DIR = staticDirEnv
	}
}
