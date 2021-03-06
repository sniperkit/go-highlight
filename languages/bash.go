/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"bash", "sh", "zsh"}, `{"aliases":["sh","zsh"],"lexemes":"-?[a-z\\._]+","keywords":{"keyword":"if then else elif fi for while in do done case esac function","literal":"true false","built_in":"break cd continue eval exec exit export getopts hash pwd readonly return shift test times trap umask unset alias bind builtin caller command declare echo enable help let local logout mapfile printf read readarray source type typeset ulimit unalias set shopt autoload bg bindkey bye cap chdir clone comparguments compcall compctl compdescribe compfiles compgroups compquote comptags comptry compvalues dirs disable disown echotc echoti emulate fc fg float functions getcap getln history integer jobs kill limit log noglob popd print pushd pushln rehash sched setcap setopt stat suspend ttyctl unfunction unhash unlimit unsetopt vared wait whence where which zcompile zformat zftp zle zmodload zparseopts zprof zpty zregexparse zsocket zstyle ztcp","_":"-ne -eq -lt -gt -f -d -e -s -l -a"},"contains":[{"className":"meta","begin":"^#![^\\n]+sh\\s*$","relevance":10},{"className":"function","begin":"\\w[\\w\\d_]*\\s*\\(\\s*\\)\\s*\\{","returnBegin":true,"contains":[{"className":"title","begin":"\\w[\\w\\d_]*","relevance":0}],"relevance":0},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"string","begin":"\"","end":"\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"variable","variants":[{"begin":"\\$[\\w\\d#@][\\w\\d_]*"},{"begin":"\\$\\{(.*?)}"}]},{"className":"variable","begin":"\\$\\(","end":"\\)","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]}]},{"className":"string","begin":"'","end":"'"},{"className":"variable","variants":[{"begin":"\\$[\\w\\d#@][\\w\\d_]*"},{"begin":"\\$\\{(.*?)}"}]}]}`)
}
