module gochess

require go.uber.org/fx v1.17.1

require adapter v1.0.0

require repository v1.0.0

require config v1.0.0

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	go.uber.org/atomic v1.6.0 // indirect
	go.uber.org/dig v1.14.0 // indirect
	go.uber.org/multierr v1.5.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/sys v0.0.0-20210903071746-97244b99971b // indirect
)

replace domain v1.0.0 => ./domain

replace config v1.0.0 => ./config

replace adapter v1.0.0 => ./adapter

replace repository v1.0.0 => ./domain/repository

go 1.18
