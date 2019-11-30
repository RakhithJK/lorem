// MIT License
//
// Copyright (c) 2019 Sergey Kibish
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/skibish/lorem/ipsum"
)

var (
	buildVersion    string
	buildCommitHash string
)

//go:generate go run gen.go
func main() {
	log.SetFlags(0)              // remove all print flags from logger (create simple output, without time)
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	fs := flag.NewFlagSet("lorem", flag.ExitOnError)
	var (
		startWithLorem = fs.Bool("ipsum", false, "Start with \"Lorem ipsum dolor sit amet...\"")
		printStats     = fs.Bool("stats", false, "Print statistics")
		gType          = fs.String("type", "p", "What to generate: p (paragraphs); w (words); b (bytes)")
		number         = fs.Int("number", 5, "How many <type> to generate")
		showVersion    = fs.Bool("v", false, "Show version and exit")
	)
	err := fs.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("ERROR: Failed to get configuration - %v", err)
	}

	if *showVersion {
		fmt.Printf("Version: %s\nCommitHash: %s\n", buildVersion, buildCommitHash)
		return
	}

	if *number <= 0 {
		log.Fatal("ERROR: \"number\" should be greater than 0")
	}

	loremIpsum := ipsum.New(os.Stdout, *startWithLorem, ipsum.Option(*gType), *number)
	err = loremIpsum.Generate()

	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	if *printStats {
		stats := loremIpsum.Stats()
		log.Print("\n")
		log.Printf("STATS:\n")
		log.Printf("Words      %d\n", stats.WordCount)
		log.Printf("Bytes      %d\n", stats.ByteCount)
		log.Printf("Paragraphs %d\n", stats.ParagraphCount)
	}
}
