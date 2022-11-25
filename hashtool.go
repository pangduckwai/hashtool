package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io"
	"log"
	"math"
	"os"
)

func main() {
	cfg, err := parse(os.Args)
	if err != nil {
		if errr, ok := err.(*Err); !ok || errr.Code > 0 {
			log.Fatal(err)
		}
		log.Fatalf("%v\n%v\n%v\n", err, app(), usage())
	}

	switch cfg.Command {
	case 0:
		validate(cfg)
		hashing(cfg)
	case 1:
		fmt.Printf("%v\n%v\n", app(), help())
	case 2:
		fmt.Println(app())
	}

	if err != nil {
		log.Fatal(err)
	}
}

func hashing(cfg *Config) error {
	if cfg.Verbose {
		fmt.Printf("Hashing data from %v using %v (buffer size: %v)...\n", display(cfg), aLGORITHM[cfg.Algorithm], cfg.Buffer)
	}

	var err error
	inp := os.Stdin
	if cfg.Input != "" {
		inp, err = os.Open(cfg.Input)
		if err != nil {
			return err
		}
	}
	rdr := bufio.NewReaderSize(inp, cfg.Buffer)

	var hsh hash.Hash
	switch cfg.Algorithm {
	case 0:
		hsh = md5.New()
	case 1:
		hsh = sha1.New()
	case 2:
		hsh = sha256.New()
	}

	buf := make([]byte, 0, cfg.Buffer)
	for idx := 0; ; idx++ {
		cnt, err := rdr.Read(buf[:cap(buf)])
		if cfg.Verbose {
			verbose(idx, cnt, cfg)
		}

		// As described in the doc, process read data first if n > 0 before
		// handling error, which could have been EOF
		if cnt > 0 {
			if cfg.Input == "" && buf[:cnt][cnt-1] == '\n' {
				err = io.EOF
			}

			hsh.Write(buf[:cnt])
		}

		if err != nil {
			if err == io.EOF {
				break // Done
			} else {
				return err
			}
		}
	}

	fmt.Println(hex.EncodeToString(hsh.Sum(nil)))

	if cfg.Verbose {
		fmt.Println("Hashing finished")
	}

	return nil
}

func Version() string {
	return "0.1.0"
}

func app() string {
	return fmt.Sprintf("Hashing tool (version %v)", Version())
}

func validate(cfg *Config) {
	if cfg.Input != "" {
		if _, err := os.Stat(cfg.Input); errors.Is(err, os.ErrNotExist) {
			log.Fatalf("Input file '%v' does not exist\n", cfg.Input)
		} else if err != nil {
			log.Fatal(err)
		}
	}
}

func display(cfg *Config) string {
	inp := "stdin"
	if cfg.Input != "" {
		inp = cfg.Input
	}

	return inp
}

func verbose(idx, cnt int, cfg *Config) {
	digits := int(math.Log10(float64(cfg.Buffer))) + 1
	format := fmt.Sprintf("%%%dv", digits)

	plr := "s"
	if cnt < 2 {
		plr = " "
	}

	fmt.Printf("%4v - read "+format+"/%v byte%v\n", idx, cnt, cfg.Buffer, plr)
}

type Err struct {
	Code uint8
	Msg  string
}

func (e *Err) Error() string {
	return fmt.Sprintf("%v", e.Msg)
}
