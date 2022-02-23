package config

type Database struct {
	URI  string
	Name string
}

type JWT struct {
	SecretKey string
}

// ENV .env struct
type ENV struct {
	// App port
	AppPort string

	// Database
	Database Database

	// JWT
	JWT JWT
}
