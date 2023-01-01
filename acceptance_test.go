package bibx_test

import (
	"strings"
	"testing"

	"github.com/blazejsewera/bibx"
	"github.com/stretchr/testify/assert"
)

var testStringStandard = "test string 1"
var standardBibMarkdownContent = `
# Some markdown

This is a test paragraph.
` + "```bibtex\n" + testStringStandard + "\n```" + `
some other text
`

var testStringDecoy = "test string 2"
var bibMarkdownContentWithDecoy = standardBibMarkdownContent + `
This is not a code fence,
because it doesn't start from a new line / whitespace: ` + "```bibtex" + `
However, a code fence will be below:
` + "```bibtex\n" + testStringDecoy + "\n```" + `
some other text`

var testStringIndent = "test string 3"
var bibMarkdownContentWithIndent = standardBibMarkdownContent + `
This will be another bibtex, this time with whitespace:
  ` + "```bibtex\n  " + testStringIndent + "\n  ```"

func TestStandardExtract(t *testing.T) {
	// given
	input := strings.NewReader(standardBibMarkdownContent)
	expected := []string{testStringStandard}

	// when
	actual := bibx.Extract(input)

	// then
	assert.ElementsMatch(t, expected, actual)
}

func DisableTestDecoyExtract(t *testing.T) {
	// given
	input := strings.NewReader(bibMarkdownContentWithDecoy)
	expected := []string{testStringStandard, testStringDecoy}

	// when
	actual := bibx.Extract(input)

	// then
	assert.ElementsMatch(t, expected, actual)
}

func DisableTestIndentExtract(t *testing.T) {
	// given
	input := strings.NewReader(bibMarkdownContentWithIndent)
	expected := []string{testStringStandard, testStringIndent}

	// when
	actual := bibx.Extract(input)

	// then
	assert.ElementsMatch(t, expected, actual)
}

func TestMerge(t *testing.T) {
	// given
	input := []string{testStringStandard, testStringDecoy, testStringIndent}
	expected := testStringStandard + "\n" +
		testStringDecoy + "\n" +
		testStringIndent + "\n"

	// when
	actual := bibx.Merge(input)

	// then
	assert.Equal(t, expected, actual)
}
