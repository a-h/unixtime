package main

import (
	"fmt"
	"os"
	"time"
)

var local bool

func display(unit string) {
	t := time.Now().UTC()
	if local {
		t = t.Local()
	}
	switch unit {
	case "ns":
		fmt.Println(t.UnixNano())
	case "ms":
		fmt.Println(t.UnixMilli())
	case "us":
		fmt.Println(t.UnixMicro())
	case "s":
		fmt.Println(t.Unix())
	}
}

func parse(unit string) {
	d, err := time.ParseDuration(os.Args[len(os.Args)-1] + unit)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	t := time.Unix(0, d.Nanoseconds()).UTC()
	if local {
		t = t.Local()
	}
	fmt.Println(t.Format(time.RFC3339Nano))
}

func printUsageAndQuit() {
	fmt.Print(`Parse and print Unix timestamps.

Usage: unixtime <ns|us|ms|s> [-l local] [timestamp]

Print time in seconds:
  unixtime s

Parse time in seconds:
  unixtime s 12345
`)
	os.Exit(1)
}

func main() {
	var args []string
	for _, arg := range os.Args {
		if arg == "-local" || arg == "--local" || arg == "-l" {
			local = true
			continue
		}
		args = append(args, arg)
	}
	if len(args) <= 1 {
		printUsageAndQuit()
	}
	unit := args[1]
	if !(unit == "ns" || unit == "ms" || unit == "us" || unit == "s") {
		fmt.Printf("Unknown unit: '%s'\n\n", unit)
		printUsageAndQuit()
	}
	switch len(args) {
	case 2: // No more args. Just display time.
		display(unit)
		return
	case 3: // Has a positional arg, parse it.
		parse(unit)
		return
	}
	printUsageAndQuit()
}
