package tools

import "sync"

func RunTasks(wg *sync.WaitGroup, tasks []func(), numWorkers int) {
	workerCh := make(chan struct{}, numWorkers)
	completionCh := make(chan struct{})

	for _, task := range tasks {
		wg.Add(1)
		go func(taskFunc func()) {
			defer wg.Done()
			workerCh <- struct{}{}
			taskFunc()
			<-workerCh
			completionCh <- struct{}{}
		}(task)
	}

	go func() {
		wg.Wait()
		close(completionCh)
	}()

	for range tasks {
		<-completionCh
	}
}
