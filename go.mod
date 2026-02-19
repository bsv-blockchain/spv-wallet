module github.com/bsv-blockchain/spv-wallet

go 1.24.3

// NOTE: The following replace directives are essential for maintaining the cohesion and functionality of this project.
// We are using the packages github.com/bsv-blockchain/spv-wallet/models and github.com/bsv-blockchain/spv-wallet/engine directly
// to facilitate the seamless integration of features across various components of our application.
// Removing these replaces could disrupt the interdependency between modules and hinder our ability to build cohesive features
// that often require modifications across multiple packages. Please refrain from removing these directives.
replace github.com/bsv-blockchain/spv-wallet/models => ./models

require (
	github.com/99designs/gqlgen v0.17.86
	github.com/bitcoinschema/go-map v0.2.2
	github.com/bsv-blockchain/go-paymail v0.25.1
	github.com/bsv-blockchain/go-sdk v1.2.17
	github.com/bsv-blockchain/spv-wallet/models v1.0.1
	github.com/coocood/freecache v1.2.4
	github.com/fergusstrange/embedded-postgres v1.33.0
	github.com/getkin/kin-openapi v0.133.0
	github.com/gin-contrib/pprof v1.5.3
	github.com/gin-gonic/gin v1.11.0
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-redis/redis_rate/v9 v9.1.2
	github.com/go-viper/mapstructure/v2 v2.5.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/uuid v1.6.0
	github.com/iancoleman/strcase v0.3.0
	github.com/jarcoal/httpmock v1.4.1
	github.com/joomcode/errorx v1.2.0
	github.com/mrz1836/go-cache v1.0.7
	github.com/mrz1836/go-cachestore v1.0.5
	github.com/mrz1836/go-logger v1.0.3
	github.com/mrz1836/go-sanitize v1.5.5
	github.com/mrz1836/go-validate v1.0.2
	github.com/oapi-codegen/runtime v1.1.2
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.23.2
	github.com/rafaeljusto/redigomock v2.4.0+incompatible
	github.com/robfig/cron/v3 v3.0.1
	github.com/rs/zerolog v1.34.0
	github.com/samber/lo v1.52.0
	github.com/spf13/viper v1.21.0
	github.com/stretchr/testify v1.11.1
	github.com/swaggo/swag v1.16.6
	github.com/testcontainers/testcontainers-go v0.40.0
	github.com/testcontainers/testcontainers-go/modules/postgres v0.40.0
	github.com/vmihailenco/taskq/v3 v3.2.9
	go.elastic.co/ecszerolog v0.2.0
	gorm.io/datatypes v1.2.7
	gorm.io/driver/postgres v1.6.0
	gorm.io/driver/sqlite v1.6.0
	gorm.io/gorm v1.31.1
	gorm.io/plugin/dbresolver v1.6.2
)

require (
	dario.cat/mergo v1.0.2 // indirect
	filippo.io/edwards25519 v1.1.1 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20250102033503-faa5f7b0171c // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/aws/aws-sdk-go v1.55.5 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bitcoinschema/go-bpu v0.2.2 // indirect
	github.com/bsm/redislock v0.9.4 // indirect
	github.com/bytedance/gopkg v0.1.3 // indirect
	github.com/bytedance/sonic v1.15.0 // indirect
	github.com/bytedance/sonic/loader v0.5.0 // indirect
	github.com/capnm/sysinfo v0.0.0-20130621111458-5909a53897f3 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/containerd/errdefs v1.0.0 // indirect
	github.com/containerd/errdefs/pkg v0.3.0 // indirect
	github.com/containerd/log v0.1.0 // indirect
	github.com/containerd/platforms v0.2.1 // indirect
	github.com/cpuguy83/dockercfg v0.3.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgryski/go-farm v0.0.0-20240924180020-3414d57e47da // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/docker v28.5.2+incompatible // indirect
	github.com/docker/go-connections v0.6.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/ebitengine/purego v0.9.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.13 // indirect
	github.com/gin-contrib/sse v1.1.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.22.4 // indirect
	github.com/go-openapi/jsonreference v0.21.4 // indirect
	github.com/go-openapi/spec v0.22.3 // indirect
	github.com/go-openapi/swag/conv v0.25.4 // indirect
	github.com/go-openapi/swag/jsonname v0.25.4 // indirect
	github.com/go-openapi/swag/jsonutils v0.25.4 // indirect
	github.com/go-openapi/swag/loading v0.25.4 // indirect
	github.com/go-openapi/swag/stringutils v0.25.4 // indirect
	github.com/go-openapi/swag/typeutils v0.25.4 // indirect
	github.com/go-openapi/swag/yamlutils v0.25.4 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.30.1 // indirect
	github.com/go-resty/resty/v2 v2.17.1
	github.com/go-sql-driver/mysql v1.9.3 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/goccy/go-yaml v1.19.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.8.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.18.3 // indirect
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/lib/pq v1.11.1 // indirect
	github.com/lufia/plan9stats v0.0.0-20251013123823-9fd1530e3ec3 // indirect
	github.com/magiconair/properties v1.8.10 // indirect
	github.com/mailru/easyjson v0.9.1 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v1.14.33
	github.com/miekg/dns v1.1.72 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/moby/go-archive v0.2.0 // indirect
	github.com/moby/patternmatcher v0.6.0 // indirect
	github.com/moby/sys/sequential v0.6.0 // indirect
	github.com/moby/sys/user v0.4.0 // indirect
	github.com/moby/sys/userns v0.1.0 // indirect
	github.com/moby/term v0.5.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/morikuni/aec v1.1.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/newrelic/go-agent/v3 v3.42.0 // indirect
	github.com/oasdiff/yaml v0.0.0-20250309154309-f31be36b4037 // indirect
	github.com/oasdiff/yaml3 v0.0.0-20250309153720-d2182401db90 // indirect
	github.com/onsi/gomega v1.19.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.1 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/perimeterx/marshmallow v1.1.5 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20240221224432-82ca36839d55 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.67.5 // indirect
	github.com/prometheus/procfs v0.19.2 // indirect
	github.com/quic-go/qpack v0.6.0 // indirect
	github.com/quic-go/quic-go v0.59.0 // indirect
	github.com/sagikazarmark/locafero v0.12.0 // indirect
	github.com/shirou/gopsutil/v4 v4.26.1 // indirect
	github.com/sirupsen/logrus v1.9.4 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	github.com/spf13/cast v1.10.0 // indirect
	github.com/spf13/pflag v1.0.10
	github.com/stretchr/objx v0.5.3 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/swaggo/files v1.0.1
	github.com/swaggo/gin-swagger v1.6.1
	github.com/tklauser/go-sysconf v0.3.16 // indirect
	github.com/tklauser/numcpus v0.11.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.3.1 // indirect
	github.com/vektah/gqlparser/v2 v2.5.31 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/woodsbury/decimal128 v1.4.0 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.65.0 // indirect
	go.opentelemetry.io/otel v1.40.0 // indirect
	go.opentelemetry.io/otel/metric v1.40.0 // indirect
	go.opentelemetry.io/otel/trace v1.40.0 // indirect
	go.yaml.in/yaml/v2 v2.4.3 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/arch v0.24.0 // indirect
	golang.org/x/crypto v0.47.0
	golang.org/x/mod v0.32.0 // indirect
	golang.org/x/net v0.49.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.40.0 // indirect
	golang.org/x/text v0.33.0 // indirect
	golang.org/x/tools v0.41.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260203192932-546029d2fa20 // indirect
	google.golang.org/grpc v1.78.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/mysql v1.6.0 // indirect
)

// Issue with redislock package
replace github.com/bsm/redislock => github.com/bsm/redislock v0.7.2

// Issue with using wrong version of Redigo
replace github.com/gomodule/redigo => github.com/gomodule/redigo v1.9.3

// Issue with godotenv v1.6.0 pre-release
replace github.com/joho/godotenv => github.com/joho/godotenv v1.5.1

// CVE-2026-24051: Path hijacking vulnerability in otel/sdk v1.20.0-1.39.0
// Vulnerability: Executes ioreg command with untrusted search path on macOS
// Fixed in v1.40.0. Force upgrade of indirect test dependency.
// Reference: https://nvd.nist.gov/vuln/detail/CVE-2026-24051
replace go.opentelemetry.io/otel/sdk => go.opentelemetry.io/otel/sdk v1.40.0
