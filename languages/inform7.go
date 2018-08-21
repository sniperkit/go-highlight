/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"inform7", "i7"}, `{"aliases":["i7"],"case_insensitive":true,"keywords":{"keyword":"thing room person man woman animal container supporter backdrop door scenery open closed locked inside gender is are say understand kind of rule"},"contains":[{"className":"string","begin":"\"","end":"\"","relevance":0,"contains":[{"className":"subst","begin":"\\[","end":"\\]"}]},{"className":"section","begin":"^(Volume|Book|Part|Chapter|Section|Table)\\b","end":"$"},{"begin":"^(Check|Carry out|Report|Instead of|To|Rule|When|Before|After)\\b","end":":","contains":[{"begin":"\\(This","end":"\\)"}]},{"className":"comment","begin":"\\[","end":"\\]","contains":[{"Ref":["contains","3"]}]}]}`)
}
