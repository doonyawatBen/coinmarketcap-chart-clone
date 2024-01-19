package infrastructure

import "sync"

func ShutdownServices() {
	var wg sync.WaitGroup

	for _, close := range ServiceCloses {
		wg.Add(1)

		go func(close func()) {
			defer wg.Done()
			close()
		}(close)
	}

	wg.Wait()
}
