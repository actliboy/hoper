module github.com/liov/hoper/v2

go 1.13

require (
	bou.ke/monkey v1.0.2
	cloud.google.com/go v0.72.0 // indirect
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.0.2
	github.com/99designs/gqlgen v0.11.3
	github.com/Shopify/sarama v1.27.2
	github.com/alta/protopatch v0.3.3 // indirect
	github.com/armon/go-metrics v0.3.8 // indirect
	github.com/aws/aws-sdk-go v1.38.3
	github.com/boombuler/barcode v1.0.0
	github.com/cespare/xxhash v1.1.0
	github.com/cockroachdb/errors v1.8.1 // indirect
	github.com/cockroachdb/pebble v0.0.0-20201228155439-c3ef93f9a9ed
	github.com/dgraph-io/badger/v3 v3.0.0-20210309075542-2245c18dfd1f // indirect
	github.com/dgraph-io/ristretto v0.0.4-0.20210311064603-e4f298c8aa88
	github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-gonic/gin v1.7.1
	github.com/go-oauth2/oauth2/v4 v4.2.0
	github.com/go-openapi/loads v0.19.5
	github.com/go-openapi/runtime v0.19.23
	github.com/go-openapi/spec v0.19.8
	github.com/go-openapi/swag v0.19.9
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-redis/redis/v8 v8.4.10
	github.com/gofiber/fiber/v2 v2.3.2
	github.com/gogo/protobuf v1.3.2
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e
	github.com/golang/protobuf v1.5.1
	github.com/google/flatbuffers v1.12.1-0.20210406112345-261cf3b20473 // indirect
	github.com/google/uuid v1.1.2
	github.com/gorilla/sessions v1.2.1
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.1.0
	github.com/json-iterator/go v1.1.10
	github.com/lni/dragonboat/v3 v3.3.1
	github.com/markbates/goth v1.63.0
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	github.com/microcosm-cc/bluemonday v1.0.4
	github.com/modern-go/reflect2 v1.0.1
	github.com/mozillazg/go-pinyin v0.18.0 // indirect
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/nsqio/go-nsq v1.0.7
	github.com/olivere/elastic v6.2.23+incompatible
	github.com/olivere/elastic/v7 v7.0.24 // indirect
	github.com/pelletier/go-toml v1.8.1
	github.com/prometheus/client_golang v1.9.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/russross/blackfriday v2.0.0+incompatible
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tealeg/xlsx/v3 v3.2.1
	github.com/ugorji/go/codec v1.1.7
	github.com/valyala/fasthttp v1.18.0
	go.opencensus.io v0.23.0
	go.uber.org/atomic v1.7.0
	go.uber.org/multierr v1.7.0
	go.uber.org/zap v1.16.1-0.20210329175301-c23abee72d19
	golang.org/x/exp v0.0.0-20201229011636-eab1b5eb1a03 // indirect
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
	golang.org/x/oauth2 v0.0.0-20201208152858-08078c50e5b5
	golang.org/x/tools v0.1.0
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20210106152847-07624b53cd92
	google.golang.org/grpc v1.36.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gorm.io/driver/mysql v1.0.3
	gorm.io/driver/postgres v1.0.0
	gorm.io/driver/sqlite v1.1.3
	gorm.io/gorm v1.21.8
	gorm.io/plugin/prometheus v0.0.0-20210112035011-ae3013937adc
)

replace (
	github.com/cenkalti/backoff => github.com/cenkalti/backoff/v4 v4.1.0
	github.com/coreos/etcd/client => go.etcd.io/etcd/client/v3 v3.5.0-alpha.0
	github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
//google.golang.org/grpc => google.golang.org/grpc v1.29.0
)
