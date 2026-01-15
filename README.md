# Event-App-API-Go
A simple API using Go with JWT authentication.


# create up down scripts cmd
migrate create -ext sql -dir .\cmd\migrate\migrations -seq create_users_table
## cmd to run migration
go run .\cmd\migrate\main.go up
## for CGO_ENABLED=0 error and later 2nd error cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%
Run go env -w CGO_ENABLED=1
Download and install tdm64-gcc-5.1.0-2.exe with default options (to add to PATH env variable), and restart your terminal or IDE session.
(https://sourceforge.net/projects/tdm-gcc/files/TDM-GCC%20Installer/tdm64-gcc-5.1.0-2.exe/download)