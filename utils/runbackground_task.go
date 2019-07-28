package utils

import (
	"context"
	"sync"
)

func RunBackgroundTask(bgContext context.Context, bgTask func(context.Context)) func() {
	ctx, cancel := context.WithCancel(bgContext)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		bgTask(ctx)
		wg.Done()
	}()

	bgTaskCanceller := func() {
		cancel()
		wg.Wait()
	}

	return bgTaskCanceller
}
