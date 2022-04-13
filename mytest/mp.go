package mytest

// type T1 struct {
// 	M1		map[string]int
// 	M2		map[int]int
// }

type BaseArgs struct {
	Cid int64
	Seq int64
}
type HeartbeatArgs struct {
	BaseArgs
	NodeId		int
	Addr		string
	Groups		map[int]*GroupInfo
}

type GroupInfo struct {
	Id			int
	ConfNum		int
	IsLeader	bool
	Status		GroupStatus
	Shards		map[int]ShardInfo
	Size		int64
	Peer		int
}

type GroupStatus int

const (
	GroupJoined  = iota
	GroupServing
	GroupLeaving
	GroupRemoving
	GroupRemoved
)

type ShardInfo struct {
	Id			int
	Gid			int
	Status		ShardStatus
	Size		int64
	Capacity	uint64
	RangeStart	string
	RangeEnd	string
}

type ShardStatus int

const (
	SERVING  	ShardStatus = iota
	INVALID
	PULLING
	ERASING
	WAITING
)