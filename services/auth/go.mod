module github.com/capstone-project-bunker/backend/services/auth

go 1.19

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.10.0
	github.com/golang-jwt/jwt/v4 v4.4.3
	github.com/google/uuid v1.3.0
	github.com/jackc/pgerrcode v0.0.0-20220416144525-469b46aa5efa
	github.com/joho/godotenv v1.4.0
	github.com/lib/pq v1.10.7
	go.uber.org/zap v1.24.0
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.28.1
	shared v0.0.0-00010101000000-000000000000
)

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace shared => ../shared
