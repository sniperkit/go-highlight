package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("lisp", `{"illegal":"/\\S/","contains":[0,1,2,3,4,5,6,7,8]}`)
}