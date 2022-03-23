package loadbalance

import (
	"strings"
	"sync"
	"sync/atomic"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
)

// Picker struct.
type Picker struct {
	mu        sync.Mutex
	leader    balancer.SubConn
	followers []balancer.SubConn
	current   uint64
}

var _ base.PickerBuilder = (*Picker)(nil)

// Build builds a picker.
// Pickers use the builder pattern just like the resolvers,
// gRPC passes a map of subconnections with information about those
// subconnections to Build() to set up the picker.
func (p *Picker) Build(buildInfo base.PickerBuildInfo) balancer.Picker {
	p.mu.Lock()
	defer p.mu.Unlock()
	var followers []balancer.SubConn
	for sc, scInfo := range buildInfo.ReadySCs {
		isLeader := scInfo.Address.Attributes.Value("is_leader").(bool)
		if isLeader {
			p.leader = sc
			continue
		}
		followers = append(followers, sc)
	}
	p.followers = followers
	return p
}

var _ balancer.Picker = (*Picker)(nil)

// Pick gRPC gives Pick a balancer.PickInfo containing the RPC's name and context
// to help the picker know what subconnection to pick.
func (p *Picker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	var result balancer.PickResult
	if strings.Contains(info.FullMethodName, "Produce") || len(p.followers) == 0 {
		result.SubConn = p.leader
	} else if strings.Contains(info.FullMethodName, "Consume") {
		result.SubConn = p.nextFollower()
	}
	if result.SubConn == nil {
		return result, balancer.ErrNoSubConnAvailable
	}

	return result, nil
}

func (p *Picker) nextFollower() balancer.SubConn {
	cur := atomic.AddUint64(&p.current, uint64(1))
	len := uint64(len(p.followers))
	idx := int(cur % len)
	return p.followers[idx]
}

func init() {
	base.NewBalancerBuilder(Name, &Picker{}, base.Config{})
}
