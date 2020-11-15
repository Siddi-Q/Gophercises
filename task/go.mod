module main

go 1.15

replace example.com/cmd => ./cmd

replace example.com/db => ./db

require (
	example.com/cmd v0.0.0-00010101000000-000000000000
	example.com/db v0.0.0-00010101000000-000000000000
	github.com/boltdb/bolt v1.3.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.1.1 // indirect
)
