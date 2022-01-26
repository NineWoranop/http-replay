package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

var usage = fmt.Sprintf(`Usage: %s [-path=]|[-scrape-interval=10s]|[-total-dat-file=1]]`, os.Args[0])
var maxTotalFiles uint = 999999

type CmdArguments struct {
	ListenAddress  string
	MetricsPath    string
	ScrapeInternal time.Duration
	Path           string
	TotalFiles     uint
	AutoRepeat     bool
}

func ReadArguments() (args *CmdArguments, err error) {
	listenAddressPtr := flag.String("web.listen-address", ":9191", "Address to listen on for web interface and telemetry")
	metricsPathPtr := flag.String("web.telemetry-path", "/metrics", "Address to listen on for web interface and telemetry")
	scrapeInternalPtr := flag.Duration("scrape-interval", time.Second*10, "Scrape interval to fetch metrics and write dat file")
	pathPtr := flag.String("path", "./", "Path for read dat file")
	totalFilesPtr := flag.Uint("total-dat-file", 1, "Number of dat files to write")
	autoRepeatPtr := flag.Bool("autorepeat", true, "Auto repeat for replay from start")

	flag.Parse()

	if pathPtr == nil {
		return nil, fmt.Errorf("need a path.\n%s", usage)
	}
	if *totalFilesPtr > maxTotalFiles {
		fmt.Printf("total-dat-file is more than 999999. It will reduce to be 999999")
	}

	return &CmdArguments{*listenAddressPtr, *metricsPathPtr, *scrapeInternalPtr, *pathPtr, *totalFilesPtr, *autoRepeatPtr}, nil
}

func (args CmdArguments) String() string {
	jsonText, _ := json.Marshal(args)
	return string(jsonText)
}

func (args CmdArguments) Sleep() bool {
	fmt.Printf("sleep for %s\n", args.ScrapeInternal)
	time.Sleep(args.ScrapeInternal)
	return true
}
