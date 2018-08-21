/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"clojure-repl"}, `{"contains":[{"className":"meta","begin":"^([\\w.-]+|\\s*#_)=>","starts":{"end":"$","subLanguage":["clojure"]}}]}`)
}
