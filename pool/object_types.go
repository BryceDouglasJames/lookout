package pool

import (
	"context"
	"flag"
	"fmt"
	"io"
	"sync"
	"time"

	worker "github.com/brycedouglasjames/lookout/worker_dispatch"
)

type Instance_Object interface {
	GetID() int
	Current_Time() time.Time
	Creation_Time() time.Time
	Destruction_Time() time.Time
}

type Instance_Node struct {
	ctx  context.Context
	pool Driver_Pool
	Instance_Object
}

func (n *Instance_Node) toString() string {
	return fmt.Sprintf("%v", n.pool.ActiveQueue)
}

type Instance_Graph struct {
	nodes []*Instance_Node
	edges map[context.Context][]*Instance_Node
	lock  sync.RWMutex
}

//TODO add user asscoiation with each creation
type spider struct {
	Name           string
	Reader         io.Reader
	Writer         io.Writer
	Master_Context context.Context
	Tasks          []*worker.Job_Type
	Flags          []*flag.Flag
	Instance_Graph
	//...
	//metadata
}

type Driver_Pool struct {
	Capacity    int
	WaitQueue   []Instance_Object
	ActiveQueue []Instance_Object
	Mulock      *sync.Mutex
}
