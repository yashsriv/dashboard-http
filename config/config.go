package config

import (
	"fmt"

	"github.com/markbates/pop"
	"github.com/olebedev/config"
	"github.com/ybbus/jsonrpc"
)

// HTTPPort on which server should listen
var HTTPPort int

// CookieSecret for hashing cookies
var CookieSecret string

// DatabaseConnection for all database tasks
var DatabaseConnection *pop.Connection

// FacebookAccessToken is the secret app access token
var FacebookAccessToken string

// Timetable rpc client
var Timetable *jsonrpc.RPCClient

// Share
var Share *jsonrpc.RPCClient

// InitConfig for setting up config
func InitConfig() {

	cfg, err := config.ParseYamlFile("./config.yml")
	if err != nil {
		panic(err)
	}
	cfg.EnvPrefix("DASHBOARD")

	// Can be set using DASHBOARD_HTTP_PORT environment variable
	HTTPPort, err = cfg.Int("http.port")
	if err != nil {
		panic(err)
	}

	// Can be set using DASHBOARD_SECRET_VALUE environment variable
	CookieSecret, err = cfg.String("secret.value")
	if err != nil {
		panic(err)
	}

	// Can be set using DASHBOARD_DATABASE environment variable
	database, err := cfg.String("database")
	if err != nil {
		panic(err)
	}

	DatabaseConnection, err = pop.Connect(database)
	if err != nil {
		panic(err)
	}

	// Can be set using DASHBOARD_FACEBOOK_APPID environment variable
	facebookAppID, err := cfg.String("facebook.appid")
	if err != nil {
		panic(err)
	}

	// Can be set using DASHBOARD_FACEBOOK_SECRET environment variable
	facebookAppSecret, err := cfg.String("facebook.secret")
	if err != nil {
		panic(err)
	}

	FacebookAccessToken = fmt.Sprintf("%s|%s", facebookAppID, facebookAppSecret)

	shareHost, err := cfg.String("share.host")
	if err != nil {
		panic(err)
	}

	// Can be set using DASHBOARD_SHARE_PORT environment variable

	sharePort, err := cfg.String("share.port")
	if err != nil {
		panic(err)
	}

	Share = jsonrpc.NewRPCClient(fmt.Sprintf("http://%s:%s/jrpc", shareHost, sharePort))
}
