module github.com/bsv-blockchain/spv-wallet/api/manualtests

go 1.24.3

replace github.com/bsv-blockchain/spv-wallet/models => ../../models //nolint:gomoddirectives // local development

// Issue with godotenv v1.6.0 pre-release
replace github.com/joho/godotenv => github.com/joho/godotenv v1.5.1 //nolint:gomoddirectives // version override

require (
	github.com/bsv-blockchain/go-sdk v1.2.14
	github.com/bsv-blockchain/spv-wallet-go-client v1.1.0
	github.com/go-viper/mapstructure/v2 v2.4.0
	github.com/joomcode/errorx v1.2.0
	github.com/oapi-codegen/runtime v1.1.2
	github.com/rs/zerolog v1.34.0
	github.com/samber/lo v1.52.0
	github.com/spf13/viper v1.21.0
	github.com/stretchr/testify v1.11.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/boombuler/barcode v1.1.0 // indirect
	github.com/bsv-blockchain/spv-wallet/models v1.0.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-resty/resty/v2 v2.17.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/pquerna/otp v1.5.0 // indirect
	github.com/sagikazarmark/locafero v0.12.0 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	github.com/spf13/cast v1.10.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/crypto v0.46.0 // indirect
	golang.org/x/net v0.48.0 // indirect
	golang.org/x/sys v0.40.0 // indirect
	golang.org/x/text v0.33.0 // indirect
)
