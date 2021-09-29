package pool

import (
	"errors"
	"sync"
)

func Initialize_Driver_Pool(driver_objects []PoolObject) (*Driver_Pool, error) {
	if len(driver_objects) <= 0 {
		return nil, errors.New("there are no objects available to initialize pool")
	}

	active := make([]PoolObject, 0)
	pool := &Driver_Pool{
		Waiting:  driver_objects,
		Active:   active,
		Capacity: 10,
		Mulock:   new(sync.Mutex),
	}

	return pool, nil
}
