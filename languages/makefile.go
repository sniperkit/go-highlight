/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"makefile", "mk", "mak"}, `{"aliases":["mk","mak"],"contains":[{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"begin":"^\\w+\\s*\\W*=","returnBegin":true,"relevance":0,"starts":{"end":"\\s*\\W*=","excludeEnd":true,"starts":{"end":"$","relevance":0,"contains":[{"className":"variable","begin":"\\$\\(","end":"\\)","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]}]}}},{"className":"section","begin":"^[\\w]+:\\s*$"},{"className":"meta","begin":"^\\.PHONY:","end":"$","keywords":{"meta-keyword":".PHONY"},"lexemes":"[\\.\\w]+"},{"begin":"^\\t+","end":"$","relevance":0,"contains":[{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"variable","begin":"\\$\\(","end":"\\)","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]}]}]}`)
}
