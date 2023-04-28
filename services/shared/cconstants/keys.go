package cconstants

// shared context keys
type keys struct {
	USER_ID string
	USER_ROLE string
}

var KEYS = keys{
	USER_ID: "userId",
	USER_ROLE: "userRole",
}

const (
	UserID = "userId"
	UserRole = "userRole"
)

type env struct {
	APP_ENV string
	DEVELOPMENT string
	PRODUCTION string
}

var ENV = env{
	APP_ENV: "APP_ENV",
	DEVELOPMENT: "DEV",
	PRODUCTION: "PROD",
}