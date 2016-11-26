package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"zephir", "zep"}, `{"aliases":["zep"],"case_insensitive":true,"keywords":"and include_once list abstract global private echo interface as static endswitch array null if endwhile or const for endforeach self var let while isset public protected exit foreach throw elseif include __FILE__ empty require_once do xor return parent clone use __CLASS__ __LINE__ else break print eval new catch __METHOD__ case exception default die require __FUNCTION__ enddeclare final try switch continue endfor endif declare unset true false trait goto instanceof insteadof __DIR__ __NAMESPACE__ yield finally int uint long ulong char uchar double float bool boolean stringlikely unlikely","contains":[{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"__halt_compiler.+?;","end":".^","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"endsWithParent":true,"keywords":"__halt_compiler","lexemes":"[a-zA-Z_]\\w*"},{"className":"string","begin":"<<<['\"]?\\w+['\"]?$","end":"^\\w+;","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"begin":"(::|->)+[a-zA-Z_\\x7f-\\xff][a-zA-Z0-9_\\x7f-\\xff]*"},{"className":"function","beginKeywords":"function","end":"[;{]","excludeEnd":true,"illegal":"\\$|\\[|%","contains":[{"className":"title","begin":"[a-zA-Z_]\\w*","relevance":0},{"className":"params","begin":"\\(","end":"\\)","contains":[{"Ref":["contains","6","contains","1"]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"string","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}],"variants":[{"begin":"b\"","end":"\""},{"begin":"b'","end":"'"},{"className":"string","begin":"'","end":"'","illegal":null,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":null,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]}]},{"variants":[{"className":"number","begin":"\\b(0b[01]+)","relevance":0},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0}]}]}]},{"className":"class","beginKeywords":"class interface","end":"{","excludeEnd":true,"illegal":"[:\\(\\$\"]","contains":[{"beginKeywords":"extends implements"},{"className":"title","begin":"[a-zA-Z_]\\w*","relevance":0}]},{"beginKeywords":"namespace","end":";","illegal":"[\\.']","contains":[{"className":"title","begin":"[a-zA-Z_]\\w*","relevance":0}]},{"beginKeywords":"use","end":";","contains":[{"className":"title","begin":"[a-zA-Z_]\\w*","relevance":0}]},{"begin":"=>"},{"className":"string","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}],"variants":[{"begin":"b\"","end":"\""},{"begin":"b'","end":"'"},{"className":"string","begin":"'","end":"'","illegal":null,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":null,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]}]},{"variants":[{"className":"number","begin":"\\b(0b[01]+)","relevance":0},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0}]}]}`)
}