package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"

	"github.com/meteocima/magda_rd2wrf/radar"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("usage: magda_rd2wrf <inputdir> <outfilename> YYYYMMDDHHNN")
	}

	var instant time.Time
	var reader io.Reader
	var outfile *os.File

	var err error

	if instant, err = time.Parse("200601021504", os.Args[3]); err != nil {
		log.Fatal(err)
	}

	flags := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	m := os.FileMode(0644)
	if outfile, err = os.OpenFile(os.Args[2], flags, m); err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	if reader, err = radar.Convert(os.Args[1], instant); err != nil {
		log.Fatal(err)
	}

	outfileBuff := bufio.NewWriter(outfile)

	if _, err = io.Copy(outfileBuff, reader); err != nil {
		log.Fatal(err)
	}
}
