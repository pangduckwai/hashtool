package main

import (
	"fmt"
	"strconv"
	"strings"
)

const bUFFER = 1048576

var aLGORITHM = [...]string{"md5", "sha1", "sha256"}

type Config struct {
	Command   uint8  // default is to do hashing
	Input     string // nil - stdin
	Buffer    int    // buffer size
	Algorithm int    // 0 - md5, 1 - sha1, 2 - sha256
	Verbose   bool
}

func usage() string {
	return "Usage:\n hashtool [version | help]\n" +
		"   {-a ALGR | --algorithm=ALGR}\n" +
		"   {-i FILE | --in=FILE}\n" +
		"   {-b SIZE | --buffer=SIZE}\n" +
		"   {-v | --verbose}"
}

func help() string {
	return fmt.Sprintf("Usage: hashtool {commands} {options}\n"+
		" * commands:\n"+
		"    version - display current version of 'hashtool'\n"+
		"    help    - display this message\n"+
		" * options:\n"+
		"    -a ALGR, --algorithm=ALGR\n"+
		"       hashing algorithm to use, supports %v\n"+
		"    -i FILE, --in=FILE\n"+
		"       name of the input file, omitting means input from stdin\n"+
		"    -b SIZE, --buffer=SIZE\n"+
		"       size of the read buffer (SIZE default: %vKB)\n"+
		"    -v, --verbose\n"+
		"       display detail operation messages during processing", aLGORITHM, bUFFER/1024)
}

func algrm(algr string) (int, error) {
	for i := 0; i < len(aLGORITHM); i++ {
		if aLGORITHM[i] == algr {
			return i, nil
		}
	}
	return -1, &Err{1, fmt.Sprintf("Unsupported algorithm %v", algr)}
}

func parse(args []string) (cfg *Config, err error) {
	cfg = &Config{
		Command:   0,
		Buffer:    bUFFER,
		Algorithm: 2,
		Verbose:   false,
	}

	var val int
loop:
	for i := 1; i < len(args); i++ {
		switch {
		case args[i] == "help":
			cfg.Command = 1
			break loop
		case args[i] == "version":
			cfg.Command = 2
			break loop
		case args[i] == "-v" || args[i] == "--verbose":
			cfg.Verbose = true
		case args[i] == "-a":
			i++
			if i >= len(args) {
				return nil, &Err{2, "Missing algorithm argument"}
			} else {
				cfg.Algorithm, err = algrm(args[i])
				if err != nil {
					return nil, err
				}
			}
		case strings.HasPrefix(args[i], "--algorithm="):
			if len(args[i]) <= 12 {
				return nil, &Err{2, "Missing algorithm"}
			} else {
				cfg.Algorithm, err = algrm(args[i][12:])
				if err != nil {
					return nil, err
				}
			}
		case args[i] == "-i":
			i++
			if i >= len(args) {
				return nil, &Err{3, "Missing input filename argument"}
			} else {
				cfg.Input = args[i]
			}
		case strings.HasPrefix(args[i], "--in="):
			if len(args[i]) <= 5 {
				return nil, &Err{3, "Missing input filename"}
			} else {
				cfg.Input = args[i][5:]
			}
		case args[i] == "-b":
			i++
			if i >= len(args) {
				return nil, &Err{4, "Missing buffer size argument"}
			} else {
				val, err = strconv.Atoi(args[i])
				if err == nil {
					cfg.Buffer = val
				}
			}
		case strings.HasPrefix(args[i], "--buffer="):
			if len(args[i]) <= 9 {
				return nil, &Err{4, "Missing buffer size"}
			} else {
				val, err = strconv.Atoi(args[i][9:])
				if err == nil {
					cfg.Buffer = val
				}
			}
		default:
			return nil, &Err{0, fmt.Sprintf("Invalid argument '%v'", args[i])}
		}
	}

	return
}
