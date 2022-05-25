module github.com/go-packagist/cache

go 1.16

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/stretchr/testify v1.7.1
)

replace (
	github.com/go-packagist/foundation => ../foundation/
)