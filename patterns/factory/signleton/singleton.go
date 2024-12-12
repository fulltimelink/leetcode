package signleton

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
var once sync.Once

type single struct{}

var singleInstance *single

func getInstance() *single {
	if nil == singleInstance {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &single{}
			fmt.Println("Single instance created")
		} else {
			fmt.Println("Single instance alredy created")
		}
	} else {
		fmt.Println("Single instance alredy created")
	}
	return singleInstance
}

func getInstanceSyncOnce() *single {
	if nil == singleInstance {
		once.Do(func() {
			singleInstance = &single{}
			fmt.Println("Single instance created")
		})
	} else {
		fmt.Println("Single instance alredy created")
	}
	return singleInstance
}
