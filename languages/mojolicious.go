/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"mojolicious"}, `{"subLanguage":["xml"],"contains":[{"className":"meta","begin":"^__(END|DATA)__$"},{"begin":"^\\s*%{1,2}={0,2}","end":"$","subLanguage":["perl"]},{"begin":"<%{1,2}={0,2}","end":"={0,1}%>","subLanguage":["perl"],"excludeBegin":true,"excludeEnd":true}]}`)
}
