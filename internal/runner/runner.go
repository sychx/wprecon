package runner

import "sync"

func NewRunner(run_func func()) {
	var wg sync.WaitGroup
	
	for i := 0; i < 35; i++ {
		wg.Add(1)
	
		go func () {
			run_func()

			defer wg.Done()
		}()
	}

	wg.Wait()
}

