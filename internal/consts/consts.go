package consts

const (
	// Constant keys

	// API_URL is a key used to lookup the api url (Used by: data-collector)
	API_URL = "API_URL"
	// API_PORT is a key used to lookup the api port (Used by: data-collector)
	API_PORT = "API_PORT"
	// CERT_DIR is a key used to lookup the certificate directory (Used by: api/docker)
	CERT_DIR = "CERT_DIR"
	// CLIENT_CERT is a key used to lookup the data-collector client certificate (Used by: data-collector/api/docker)
	CLIENT_CERT = "CLIENT_CERT"
	// API_CERT is a key used to lookup the api certificate file (Used by: api)
	API_CERT = "API_CERT"
	// API_KEY is a key used to lookup the api certificate key (Used by: api)
	API_KEY = "API_KEY"
	// DB_URI is a key used to lookup the database URI (Used by: api)
	DB_URI = "DB_URI"
	// DB_NAME is a key used to lookup the database name (Used by: api)
	DB_NAME = "DB_NAME"
	// DB_USER is a key used to lookup the database user (Used by: api)
	DB_USER = "DB_USER"
	// DB_PASS is a key used to lookup the database pass (Used by: api)
	DB_PASS = "DB_PASS"
	// LOG_FILE is a key used to lookup the location of the log file (Used by: api/data-collector)
	LOG_FILE = "LOG_FILE"
	// HEALTH_DELAY is a key used to lookup the minimum delay since last health reported (in minutes) to show server as offline (Used by: api)
	HEALTH_DELAY = "HEALTH_DELAY"
	// DATA_WEBSOCKET_DELAY is a key used to lookup the delay (in seconds) between sending more data via websockets (Used by: api)
	DATA_WEBSOCKET_DELAY = "DATA_WEBSOCKET_DELAY"
	// Constant filenames

	// AGENT_STORE_FILENAME is the file location of agent filestore (Used by: data-collector)
	AGENT_STORE_FILENAME = "config.json"

	// Constant collection names

	// COLLECTION_HEALTH is the collection name used for the health collection (Used by: api)
	COLLECTION_HEALTH = "health"
	// COLLECTION_HOST is the collection name used for the host collection (Used by: api)
	COLLECTION_HOST = "host"
)
