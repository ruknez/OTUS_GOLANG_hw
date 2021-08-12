package hw05parallelexecution

import (
	"context"
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 || n <= 0 {
		return ErrErrorsLimitExceeded
	}

	var (
		wg          = sync.WaitGroup{}
		tasksChanal = make(chan Task, n)
		errorChanal = make(chan error, n)
		ctx, cancel = context.WithCancel(context.Background())
	)
	defer close(errorChanal)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(ctx context.Context) {
			defer wg.Done()
			for doIt := range tasksChanal {
				select {
				case <-ctx.Done():
					return
				default:
					if err := doIt(); err != nil {
						errorChanal <- err
					}
				}
			}
		}(ctx)
	}

	go func() {
		errorCounter := 0
		for range errorChanal {
			errorCounter++
			if errorCounter >= m {
				cancel()
			}
		}
	}()

	err := func() error {
		defer close(tasksChanal)
		for _, oneTask := range tasks {
			select {
			case <-ctx.Done():
				return ErrErrorsLimitExceeded
			default:
				tasksChanal <- oneTask
			}
		}
		return nil
	}()

	wg.Wait()

	return err
}
