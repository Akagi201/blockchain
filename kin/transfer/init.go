package main

import (
	"runtime"
	"strings"
	"time"

	"github.com/Akagi201/utilgo/conflag"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

var opts struct {
	Conf       string `long:"conf" description:"config file"`
	IssuerSeed string `long:"issuer-seed" description:"issuer seed"`
	ToAddr     string `long:"to" description:"to public key"`
	Amount     string `long:"amount" description:"amount to transfer"`
	Memo       string `long:"memo" description:"memo"`
	LogLevel   string `long:"log-level" default:"info" description:"Adjust the log level. Valid options are: error, warn, info, debug"`
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func init() {
	parser := flags.NewParser(&opts, flags.Default|flags.IgnoreUnknown)

	parser.Parse()

	if opts.Conf != "" {
		conflag.LongHyphen = true
		conflag.BoolValue = false
		args, err := conflag.ArgsFrom(opts.Conf)
		if err != nil {
			panic(err)
		}

		parser.ParseArgs(args)
	}

	log.Infof("opts: %+v", opts)
}

func init() {
	level, err := log.ParseLevel(strings.ToLower(opts.LogLevel))
	if err != nil {
		log.Fatalf("log level error: %v", err)
	}

	log.SetLevel(level)

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
}
