module github.com/dreadpiraterobots/siredmond

go 1.25.4

require github.com/urfave/cli/v2 v2.27.7

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.7 // indirect; converts markdown to man pages.
	github.com/russross/blackfriday/v2 v2.1.0 // indirect; go-md2man depends on this to parse text.
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect; Levenshtein distance for cli command mistype suggestions.
)
