/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"fsharp", "fs"}, `{"aliases":["fs"],"keywords":"abstract and as assert base begin class default delegate do done downcast downto elif else end exception extern false finally for fun function global if in inherit inline interface internal lazy let match member module mutable namespace new null of open or override private public rec return sig static struct then to true try type upcast use val void when while with yield","illegal":"\\/\\*","contains":[{"className":"keyword","begin":"\\b(yield|return|let|do)!"},{"className":"string","begin":"@\"","end":"\"","contains":[{"begin":"\"\""}]},{"className":"string","begin":"\"\"\"","end":"\"\"\""},{"className":"comment","begin":"\\(\\*","end":"\\*\\)","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"class","beginKeywords":"type","end":"\\(|=|$","excludeEnd":true,"contains":[{"className":"title","begin":"[a-zA-Z_]\\w*","relevance":0},{"begin":"<","end":">","contains":[{"className":"title","begin":"'[a-zA-Z0-9_]+","relevance":0}]}]},{"className":"meta","begin":"\\[<","end":">\\]","relevance":10},{"className":"symbol","begin":"\\B('[A-Za-z])\\b","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":null,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0}]}`)
}
