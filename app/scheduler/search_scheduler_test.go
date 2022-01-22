package scheduler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHtmlWithLink(t *testing.T) {
	html := "<a>Hello</a>"
	result, err := parseHtml([]byte(html))
	if err != nil {
		t.Errorf("Error parsing html: %d", err)
	} else {
		assert.Equal(t, "1", result.TotalLink)
	}
}

func TestParseHtmlWithDivAd(t *testing.T) {
	html := "<div><div></div><div id='tads'><div></div></div></div>"
	result, err := parseHtml([]byte(html))
	if err != nil {
		t.Errorf("Error parsing html: %d", err)
	} else {
		assert.Equal(t, "1", result.TotalAdword)
	}
}

func TestConvertingLineToArray(t *testing.T) {
	lines, err := readLines("../../user_agents.txt")
	if err != nil {
		t.Errorf("Error reading lines from file: %d", err)
	}
	assert.Equal(t, true, len(lines) != 0)
}
