package loadbalance_test

import (
	"testing"

	// "google.golang.org/grpc"
	// "google.golang.org/grpc/attributes"
	"google.golang.org/grpc/balancer"
	// "google.golang.org/grpc/balancer/base"
	// "google.golang.org/grpc/resolver"

	"prolog/internal/loadbalance"

	"github.com/stretchr/testify/require"
)

func TestPickerNoSubConnAvailable(t *testing.T) {
	picker := &loadbalance.Picker{}
	methods := []string{
		"/log.vX.Log/Produce",
		"/log.vX.Log/Consume",
	}
	for _, method := range methods {
		info := balancer.PickInfo{
			FullMethodName: method,
		}
		result, err := picker.Pick(info)
		require.Equal(t, balancer.ErrNoSubConnAvailable, err)
		require.Nil(t, result.SubConn)
	}
}

// func TestPickerProducesToLeader(t *testing.T){
// 	picker, subConns := setupTest()
// 	info := balancer.PickInfo{
// 		FullMethodName: "/log.vX.Log/Produce",
// 	}
// 	for _ = range 5 {
// 		gotPick, err := picker.Pick(info)
// 		require.NoError(t, err)
// 		require.Equal(t, subConns[0], gotPick.SubConn)
// 	}
// }

// func TestPickerConsumesFromFollowers(t *testing.T){
// 	picker, subConns := setupTest()
// 	info := balancer.PickInfo{
// 		FullMethodName: "/log.vX.Log/Consume",
// 	}
// 	for i := range 5 {
// 		pick, err := picker.Pick(info)
// 		require.NoError(t, err)
// 		require.Equal(t, subConns[i%2+1], pick.SubConn)
// 	}
// }

// func setupTest() (*loadbalance.Picker, []*subConn) {
// 	var subConns []*subConn
// 	buildInfo := base.PickerBuildInfo{
// 		ReadySCs: make(map[balancer.SubConn]base.SubConnInfo),
// 	}
// 	for i := range 3 {
// 		sc := &grpc.SubConn{}
// 		addr := resolver.Address{
// 			Attributes: attributes.New("is_leader", i == 0),
// 		}
// 		// 0th sub con is the leader
// 		sc.UpdateAddresses([]resolver.Address{addr})
// 		buildInfo.ReadySCs[sc] = base.SubConnInfo{Address: addr}
// 		subConns = append(subConns, sc)
// 	}
// 	picker := &loadbalance.Picker{}
// 	picker.Build(buildInfo)
// 	return picker, subConns
// }
