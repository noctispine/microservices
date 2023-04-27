package cconstants





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