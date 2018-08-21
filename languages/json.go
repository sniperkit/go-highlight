/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"json"}, `{"contains":[{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0},{"begin":"{","end":"}","contains":[{"className":"attr","begin":"\"","end":"\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}],"illegal":"\\n"},{"end":",","endsWithParent":true,"excludeEnd":true,"contains":[{"Ref":["contains"],"IsArray":true}],"keywords":{"literal":"true false null"},"begin":":"}],"illegal":"\\S"},{"begin":"\\[","end":"\\]","contains":[{"end":",","endsWithParent":true,"excludeEnd":true,"contains":[{"Ref":["contains"],"IsArray":true}],"keywords":{"literal":"true false null"}}],"illegal":"\\S"}],"keywords":{"literal":"true false null"},"illegal":"\\S"}`)
}
