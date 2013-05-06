package ready

import "fmt"

var (
	major = 0
	minor = 0
	tiny  = 1
	pre   = "-dev"
	str   = fmt.Sprintf("%d.%d.%d%s", major, minor, tiny, pre)
)

type version struct {
	Major  int
	Minjor int
	Tiny   int
	String string
}

var Version = version{
	Major:  major,
	Minjor: minor,
	Tiny:   tiny,
	String: str,
}
