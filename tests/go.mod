module github.com/a631807682/sqlcommenter-gorm/tests

go 1.20

replace github.com/a631807682/sqlcommenter-gorm => ../

require (
	github.com/a631807682/sqlcommenter-gorm v0.0.0-00010101000000-000000000000
	github.com/google/sqlcommenter/go/core v0.1.2
	github.com/google/sqlcommenter/go/net/http v0.1.0
	gorm.io/driver/mysql v1.4.7
	gorm.io/gorm v1.24.6
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	go.opentelemetry.io/otel v1.11.1 // indirect
	go.opentelemetry.io/otel/trace v1.11.1 // indirect
)
