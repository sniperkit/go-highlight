package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"prolog", }, `{"contains":[{"begin":"[a-z][A-Za-z0-9_]*","relevance":0},{"className":"symbol","variants":[{"begin":"[A-Z][a-zA-Z0-9_]*"},{"begin":"_[A-Za-z0-9_]*"}],"relevance":0},{"begin":"\\(","end":"\\)","relevance":0,"contains":[{"begin":"[a-z][A-Za-z0-9_]*","relevance":0},{"className":"symbol","variants":[{"begin":"[A-Z][a-zA-Z0-9_]*"},{"begin":"_[A-Za-z0-9_]*"}],"relevance":0},{"Ref":["contains","2"]},{"begin":":-"},{"begin":"\\[","end":"\\]","contains":[{"Ref":["contains","2","contains"],"IsArray":true}]},{"className":"comment","begin":"%","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"`+"`"+`","end":"`+"`"+`","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"0\\'(\\\\\\'|.)"},{"className":"string","begin":"0\\'\\\\s"},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0}]},{"begin":":-"},{"begin":"\\[","end":"\\]","contains":[{"Ref":["contains","2","contains"],"IsArray":true}]},{"className":"comment","begin":"%","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"`+"`"+`","end":"`+"`"+`","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"0\\'(\\\\\\'|.)"},{"className":"string","begin":"0\\'\\\\s"},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0},{"begin":"\\.$"}]}`)
}