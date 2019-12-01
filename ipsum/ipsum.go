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

package ipsum

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
)

func getWord() string {
	return dictionary[rand.Intn(len(dictionary))]
}

// LoremIpsum is a struct that holds lorem ipsum generator
type LoremIpsum struct {
	startLorem bool
	limiter    limiter
	w          io.Writer
}

// New return new LoremIpsum
func New(w io.Writer, s bool, o Option, v int) *LoremIpsum {
	return &LoremIpsum{
		w:          w,
		startLorem: s,
		limiter:    limiter{option: o, value: v},
	}
}

// Generate generates lorem ipsum text
// which is written to io.Writer
func (l *LoremIpsum) Generate() error {
	if l.startLorem {
		if err := l.startWithLorem(); err != nil {
			return err
		}
		if l.limiter.limitReached() {
			l.limiter.addParagraph()
			if err := l.print("\n"); err != nil {
				return err
			}
			return nil
		}
	}

	capitalize := true
	for !l.limiter.limitReached() {
		l.printWord(capitalize)

		var err error
		var isParagraph bool
		capitalize, isParagraph, err = l.printPunctuation()
		if err != nil {
			return fmt.Errorf("failed to print sign: %w", err)
		}
		if isParagraph {
			continue
		}

		if !l.limiter.limitReached() {
			if err := l.print(" "); err != nil {
				return err
			}
		}
	}

	if l.limiter.option != paragraphs {
		l.limiter.addParagraph()
		if err := l.print(".\n"); err != nil {
			return err
		}
	}

	return nil
}

func (l *LoremIpsum) printWord(capitalized bool) error {
	l.limiter.addWord()
	w := getWord()

	if capitalized {
		if err := l.printf("%s", strings.Title(w)); err != nil {
			return err
		}

		return nil
	}

	if err := l.print(w); err != nil {
		return err
	}

	return nil
}

func (l *LoremIpsum) printPunctuation() (beginning bool, isParagraph bool, err error) {
	if rand.Float32() >= 0.90 {
		if err := l.print(","); err != nil {
			return false, false, err
		}
		return
	}

	if rand.Float32() >= 0.90 {
		if err := l.print("."); err != nil {
			return false, false, err
		}
		return true, false, nil
	}

	if rand.Float32() >= 0.98 {
		l.limiter.addParagraph()
		if err := l.print(".\n"); err != nil {
			return false, false, err
		}
		return true, true, nil
	}

	return
}

func (l *LoremIpsum) print(a string) error {
	l.limiter.addBytes(len(a))
	_, err := fmt.Fprint(l.w, a)
	return err
}

func (l *LoremIpsum) printf(f string, a string) error {
	l.limiter.addBytes(len(a))
	_, err := fmt.Fprintf(l.w, f, a)
	return err
}

// Stats prints statistics
func (l *LoremIpsum) Stats() *Stats {
	return &Stats{
		WordCount:      l.limiter.wordCount,
		ByteCount:      l.limiter.byteCount,
		ParagraphCount: l.limiter.paragraphCount,
	}
}

var loremStart []string = []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit"}

func (l *LoremIpsum) startWithLorem() error {
	for i, word := range loremStart {
		l.limiter.addWord()
		if err := l.print(word); err != nil {
			return err
		}

		if l.limiter.limitReached() {
			if err := l.print("."); err != nil {
				return err
			}
			return nil
		}

		if l.limiter.wordCount == 5 {
			if err := l.print(","); err != nil {
				return err
			}
		}

		if i < len(loremStart)-1 {
			if err := l.print(" "); err != nil {
				return err
			}
		}

		if i == len(loremStart)-1 {
			if err := l.print(". "); err != nil {
				return err
			}
		}
	}

	return nil
}
