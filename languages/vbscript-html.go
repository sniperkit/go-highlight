/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"vbscript-html"}, `{"subLanguage":["xml"],"contains":[{"begin":"<%","end":"%>","subLanguage":["vbscript"]}]}`)
}
