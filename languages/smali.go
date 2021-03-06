/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"smali", "smali"}, `{"aliases":["smali"],"contains":[{"className":"string","begin":"\"","end":"\"","relevance":0},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":0},{"className":"keyword","variants":[{"begin":"\\s*\\.end\\s[a-zA-Z0-9]*"},{"begin":"^[ ]*\\.[a-zA-Z]*","relevance":0},{"begin":"\\s:[a-zA-Z_0-9]*","relevance":0},{"begin":"\\s(transient|constructor|abstract|final|synthetic|public|private|protected|static|bridge|system)"}]},{"className":"built_in","variants":[{"begin":"\\s(add|and|cmp|cmpg|cmpl|const|div|double|float|goto|if|int|long|move|mul|neg|new|nop|not|or|rem|return|shl|shr|sput|sub|throw|ushr|xor)\\s"},{"begin":"\\s(add|and|cmp|cmpg|cmpl|const|div|double|float|goto|if|int|long|move|mul|neg|new|nop|not|or|rem|return|shl|shr|sput|sub|throw|ushr|xor)((\\-|/)[a-zA-Z0-9]+)+\\s","relevance":10},{"begin":"\\s(aget|aput|array|check|execute|fill|filled|goto/16|goto/32|iget|instance|invoke|iput|monitor|packed|sget|sparse)((\\-|/)[a-zA-Z0-9]+)*\\s","relevance":10}]},{"className":"class","begin":"L[^(;:\n]*;","relevance":0},{"begin":"[vp][0-9]+"}]}`)
}
