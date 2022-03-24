package rules

import (
	C "github.com/Dreamacro/clash/constant"
	"github.com/gmccue/go-ipset"
)

type IPSET struct {
	set         string
	adapter     string
	noResolveIP bool
}

func (g *IPSET) RuleType() C.RuleType {
	return C.IPSET
}

func (g *IPSET) Match(metadata *C.Metadata) bool {
	ip := metadata.DstIP
	if ip == nil {
		return false
	}
	ipset, _ := ipset.New()
	err := ipset.Test(g.set, ip.String())
	return err == nil
}

func (g *IPSET) Adapter() string {
	return g.adapter
}

func (g *IPSET) Payload() string {
	return g.set
}

func (g *IPSET) ShouldResolveIP() bool {
	return !g.noResolveIP
}

func (g *IPSET) ShouldFindProcess() bool {
	return false
}

func NewIPSET(set string, adapter string, noResolveIP bool) *IPSET {
	ipset := &IPSET{
		set:         set,
		adapter:     adapter,
		noResolveIP: noResolveIP,
	}

	return ipset
}
