module main

go 1.15

replace example.com/cmd => ./cmd

require (
	example.com/cmd v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.1.1 // indirect
)
