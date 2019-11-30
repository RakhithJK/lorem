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

// Option to limit generation by
type Option string

const (
	paragraphs Option = "p"
	words      Option = "w"
	bytes      Option = "b"
)

// Stats is used to collect statistics
type Stats struct {
	WordCount      int
	ByteCount      int
	ParagraphCount int
}

type limiter struct {
	option         Option
	value          int
	wordCount      int
	byteCount      int
	paragraphCount int
}

func (l *limiter) addWord() {
	l.wordCount++
}

func (l *limiter) addParagraph() {
	l.paragraphCount++
}

func (l *limiter) addBytes(size int) {
	l.byteCount += size
}

func (l *limiter) limitReached() bool {
	switch l.option {
	case words:
		return l.wordCount >= l.value
	case bytes:
		return l.byteCount >= l.value
	case paragraphs:
		return l.paragraphCount >= l.value
	}
	return false
}
