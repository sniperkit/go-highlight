/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"apache", "apacheconf"}, `{"aliases":["apacheconf"],"case_insensitive":true,"contains":[{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"section","begin":"</?","end":">"},{"className":"attribute","begin":"\\w+","relevance":0,"keywords":{"nomarkup":"order deny allow setenv rewriterule rewriteengine rewritecond documentroot sethandler errordocument loadmodule options header listen serverroot servername"},"starts":{"end":"$","relevance":0,"keywords":{"literal":"on off all"},"contains":[{"className":"meta","begin":"\\s\\[","end":"\\]$"},{"className":"variable","begin":"[\\$%]\\{","end":"\\}","contains":[{"Ref":["contains","2","starts","contains","1"]},{"className":"number","begin":"[\\$%]\\d+"}]},{"className":"number","begin":"[\\$%]\\d+"},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]}]}}],"illegal":"\\S"}`)
}
