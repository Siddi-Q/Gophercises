module main

go 1.15

replace example.com/story => ../../story

replace example.com/storyhandler => ../storyhandler

require (
	example.com/story v0.0.0-00010101000000-000000000000
	example.com/storyhandler v0.0.0-00010101000000-000000000000
)
