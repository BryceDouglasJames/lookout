package pool

import (
	"errors"
	"sync"
)

//personalize a pool for each instance owner

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
			How will we grab and store data?
			how can we represent each relationship?

		-conduits
			IO Reader/writer to serioalize code
			forwarder to send to next node

		-router
			augment a modem and accept incoming signals
			interfaces wwith forwarder
			THE ROOT NODE ONLY NEEDS ACCESS TO INIT ONTOLOGY

		log every event

*/

func Initialize_Driver_Pool(_object []Instance_Object) (*Driver_Pool, error) {
	if len(_object) <= 0 {
		return nil, errors.New("there are no objects available to initialize pool")
	}

	active := make([]Instance_Object, 0)
	pool := &Driver_Pool{
		WaitQueue:   []Instance_Object{},
		ActiveQueue: active,
		Capacity:    10,
		Mulock:      new(sync.Mutex),
	}

	return pool, nil

}
