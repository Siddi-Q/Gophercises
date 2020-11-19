module example.com/cmd

go 1.15

replace example.com/db => ../db

require (
	example.com/db v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.1.1
)
