/*
Sniperkit-Bot
- Status: analyzed
*/

package highlight

import (
	"bytes"
	"fmt"
	"io"
)

func highlightTest(lang string, code []byte) ([]byte, error) {
	h, err := makeAndHighlight(lang, code)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	h.render(&buf, func(w io.Writer, class string, body []byte) {
		if len(class) > 0 {
			fmt.Fprintf(w, "<%s>%s</%s>", class, body, class)
		} else {
			w.Write(body)
		}
	})
	return buf.Bytes(), nil
}
