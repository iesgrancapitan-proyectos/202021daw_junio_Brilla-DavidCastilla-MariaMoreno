// +build !prod

package config

const (
	db_host = "database"
	db_port = "8529"
	DB_USER = ""
	DB_PASS = ""
	DB_URL  = "http://" + db_host + ":" + db_port
)
