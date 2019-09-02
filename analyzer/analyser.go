package analyzer

import (
	"io"
	"math"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	pageInit       = 126.4
	tableInit      = 21.8
	rowInit        = 16.8
	rowHeight      = 20.8
	singleColWidth = 24
	doubleColWidth = 48
)

var (
	wordRegex = regexp.MustCompile(`[^\s]+`)
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func rowsCount(text string, maxWidth int) int {
	count := 1
	lineLen := 0
	for _, word := range wordRegex.FindAllStringSubmatch(string(text), -1) {
		wordLen := len([]rune(word[0]))
		// Old length + Space + Word Length
		newLineLen := lineLen + 1 + wordLen
		if newLineLen > maxWidth {
			// On new line will be only last word
			lineLen = wordLen
			count++
		} else {
			lineLen = newLineLen
		}
	}
	return count
}

func CountHeight(data io.Reader) (float64, error) {
	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		return 0, err
	}

	pageHeight := pageInit
	var leftTableHeight float64

	// Tables with different days of week
	doc.Find("div.col-md-6.hidden-xs tbody").Each(func(tableNum int, tbody *goquery.Selection) {
		tableHeight := tableInit

		// Rows in one table
		tbody.Find("tr").Each(func(trNum int, tr *goquery.Selection) {
			var hasLeft, hasRight bool
			var leftText, rightText string

			// Columns in each row
			tr.Find("td").Each(func(tdNum int, td *goquery.Selection) {
				// 0 - time, 1 - left subject, 2 - right subject
				switch tdNum {
				case 1:
					hasLeft = true
					leftText = td.Text()
				case 2:
					hasRight = true
					rightText = td.Text()
				}
			})

			if hasRight {
				// Two single columns
				count := max(rowsCount(leftText, singleColWidth), rowsCount(rightText, singleColWidth))
				tableHeight += rowInit + rowHeight*float64(count)
			} else if hasLeft {
				// One double column
				count := rowsCount(leftText, doubleColWidth)
				tableHeight += rowInit + rowHeight*float64(count)
			}
		})

		// Even numbers - left tables, odd numbers - right tables
		if tableNum%2 == 0 {
			leftTableHeight = tableHeight
		} else {
			// Height of tables row is max of them
			pageHeight += math.Max(leftTableHeight, tableHeight)
			leftTableHeight = 0
		}
	})

	return pageHeight, nil
}

func CountHeightFromString(data string) (float64, error) {
	return CountHeight(strings.NewReader(data))
}
