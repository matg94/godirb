package main

import (
	"flag"

	"github.com/matg94/godirb/config"
)

func ParseFlags() config.AppFlags {
	url := flag.String("url", "", "URL to query")
	profile := flag.String("p", "default", "profile name to use, searches for file in ~/.godirb")
	limiter := flag.Float64("limiter", -1, "Maximum requests per second allowed")
	threads := flag.Int("threads", -1, "Number of threads to use, default 10")
	wordlist := flag.String("words", "", "Path to wordlist file to use")
	cookie := flag.String("cookie", "", "Cookie string to use")
	jsonPipe := flag.Bool("pipe", false, "Results will be pipeable in json")
	outFile := flag.String("out", "", "Path to file to store json results")
	flag.Parse()

	return config.AppFlags{
		URL:      *url,
		Profile:  *profile,
		Limiter:  *limiter,
		Threads:  *threads,
		Wordlist: *wordlist,
		Cookie:   *cookie,
		JsonPipe: *jsonPipe,
		OutFile:  *outFile,
	}

}
