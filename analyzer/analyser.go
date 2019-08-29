package analyzer

import (
	"io"
	"math"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	nilString      = "NIL"
	pageInit       = 126.4
	tableInit      = 21.8
	rowInit        = 16.8
	rowHeight      = 20.8
	singleColWidth = 26
	doubleColWidth = 53
	// singleColWidth = 18.9
	// doubleColWidth = 35.72
)

var (
	wordRegex = regexp.MustCompile(`[^\s]+`)
)

func rowsCount(text string, maxWidth int) float64 {
	res := 1
	length := 0
	for _, word := range wordRegex.FindAllStringSubmatch(string(text), -1) {
		wordLen := len([]rune(word[0]))
		if length+1+wordLen > maxWidth {
			res++
			length = wordLen
		} else {
			length += 1 + wordLen
		}
	}
	return float64(res)
}

func CountHeight(data io.Reader) (float64, error) {
	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		return 0, err
	}

	pageHeight := pageInit
	var leftTableHeight, rightTableHeight float64
	doc.Find("div.col-md-6.hidden-xs tbody").Each(func(tableNum int, tbody *goquery.Selection) {
		tableHeight := tableInit
		tbody.Find("tr").Each(func(trNum int, tr *goquery.Selection) {
			var trHeight float64
			leftText := nilString
			rightText := nilString
			tr.Find("td").Each(func(tdNum int, td *goquery.Selection) {
				if tdNum == 1 {
					leftText = td.Text()
				}
				if tdNum == 2 {
					rightText = td.Text()
				}
			})
			if rightText != nilString {
				// Single columns
				trHeight = math.Max(rowsCount(leftText, singleColWidth), rowsCount(rightText, singleColWidth))
				trHeight = rowInit + rowHeight*trHeight
			} else if leftText != nilString {
				// Double column
				trHeight = rowInit + rowHeight*rowsCount(leftText, doubleColWidth)
			} else {
				// Empty column
				trHeight = 0
			}
			tableHeight += trHeight
		})
		if tableNum%2 == 0 {
			leftTableHeight = tableHeight
		} else {
			rightTableHeight = tableHeight
			pageHeight += math.Max(leftTableHeight, rightTableHeight)
		}
	})

	return pageHeight, nil
}

func CountHeightFromString(data string) (float64, error) {
	return CountHeight(strings.NewReader(data))
}
