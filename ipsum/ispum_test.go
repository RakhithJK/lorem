package ipsum

import (
	b "bytes"
	"math/rand"
	"strings"
	"testing"
)

type testCase struct {
	option     Option
	value      int
	startLorem bool
}

func TestGeneratation(t *testing.T) {
	t.Run("one word", func(t *testing.T) {
		expected := 1
		str, stats := runGenerator(t, testCase{words, expected, false})

		wCount := len(strings.Fields(str))
		if stats.WordCount != wCount {
			t.Errorf("Expected count of words to be %d, got: %d", expected, wCount)
		}
	})

	t.Run("three words", func(t *testing.T) {
		expected := 3
		str, stats := runGenerator(t, testCase{words, expected, false})

		wCount := len(strings.Fields(str))
		if stats.WordCount != wCount {
			t.Errorf("Expected count of words to be %d, got: %d", expected, wCount)
		}
	})

	t.Run("start with lorem", func(t *testing.T) {
		expected := 8
		str, stats := runGenerator(t, testCase{words, expected, true})
		if strings.Contains(str, strings.Join(loremStart, " ")) {
			t.Errorf("Expected to start with lorem: %s", str)
		}

		wCount := len(strings.Fields(str))
		if stats.WordCount != wCount {
			t.Errorf("Expected count of words to be %d, got: %d", expected, wCount)
		}

		realParagraphCount := strings.Count(str, "\n")
		realParagraphCount++
		expectedParagraphs := 1
		if stats.ParagraphCount != realParagraphCount {
			t.Errorf("Expected count of paragraphs to be %d, got: %d", expectedParagraphs, stats.ParagraphCount)
		}
	})

	t.Run("start with lorem and one word ahead", func(t *testing.T) {
		expected := 9
		str, stats := runGenerator(t, testCase{words, expected, true})
		if strings.Contains(str, strings.Join(loremStart, " ")) {
			t.Errorf("Expected to start with lorem: %s", str)
		}

		wCount := len(strings.Fields(str))
		if stats.WordCount != wCount {
			t.Errorf("Expected count of words to be %d, got: %d", expected, wCount)
		}

		realParagraphCount := strings.Count(str, "\n")
		realParagraphCount++
		expectedParagraphs := 1
		if stats.ParagraphCount != realParagraphCount {
			t.Errorf("Expected count of paragraphs to be %d, got: %d", expectedParagraphs, stats.ParagraphCount)
		}
	})

	t.Run("generate 42 words", func(t *testing.T) {
		expectedWords := 42
		str, stats := runGenerator(t, testCase{words, expectedWords, false})

		wCount := len(strings.Fields(str))
		if stats.WordCount != wCount {
			t.Errorf("Expected count of words to be %d, got: %d", expectedWords, wCount)
		}

		expectedParagraphs := 2
		if stats.ParagraphCount != expectedParagraphs {
			t.Errorf("Expected count of paragraphs to be %d, got: %d", expectedParagraphs, stats.ParagraphCount)
		}

		expectedBytes := 455
		if stats.ByteCount != expectedBytes {
			t.Errorf("Expected count of bytes to be %d, got: %d", expectedBytes, stats.ByteCount)
		}
	})

	t.Run("generate 5 paragraphs", func(t *testing.T) {
		expectedParagraphs := 5

		str, stats := runGenerator(t, testCase{paragraphs, expectedParagraphs, false})
		realParag := strings.Count(str, "\n")
		realParag++

		if stats.ParagraphCount != realParag {
			t.Errorf("Expected count of paragraphs to be %d, got: %d", expectedParagraphs, realParag)
		}
	})

	t.Run("generate >=404 bytes", func(t *testing.T) {
		expectedBytes := 404

		str, stats := runGenerator(t, testCase{bytes, expectedBytes, false})

		if stats.ByteCount < expectedBytes {
			t.Errorf("Expected count of bytes to be %d, got: %d", expectedBytes, len(str))
		}
	})
}

func runGenerator(t *testing.T, tc testCase) (string, *Stats) {
	rand.Seed(42)
	buf := new(b.Buffer)
	l := New(buf, tc.startLorem, tc.option, tc.value)
	if err := l.Generate(); err != nil {
		t.Fatalf("generation failed: %v", err)
	}
	return buf.String(), l.Stats()
}
