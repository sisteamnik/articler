package date

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"time"
)

/*func Extract(in []byte) (time.Time, error) {
	return time.Now(), nil
}*/

func ExtractFromSelection(sel *goquery.Selection) (time.Time, error) {
	if strDatetime, ok := sel.Attr("datetime"); ok {
		return time.Parse(time.RFC3339, strDatetime)
	}
	return time.Now(), fmt.Errorf("date not found")
}
