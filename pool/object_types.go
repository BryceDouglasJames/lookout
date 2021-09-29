package pool

import (
	"sync"
)

//FUNCTIONS
/*
what are the properities of the instance?
		context of the request
			whos the user
			what do they want done
			handle errors

		-pool duplex
			how is this started? how will it end?
			where did it come from?

		log every event
*/

type PoolObject interface {
	getID() int
}

type Driver_Pool struct {
	Waiting  []PoolObject
	Active   []PoolObject
	Capacity int
	Mulock   *sync.Mutex
}
