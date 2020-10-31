package io

import (
	"context"
	"io"
	goio "io"
	"sync"
)

type parallelMultiWriter struct {
	ctx    context.Context
	cancel context.CancelFunc
	w      []goio.Writer
	wLen   int
	wg     *sync.WaitGroup
}

func NewParallelMultiWriter(w ...goio.Writer) io.Writer {
	ctx, cancel := context.WithCancel(context.Background())
	return parallelMultiWriter{
		ctx:    ctx,
		cancel: cancel,
		w:      w,
		wLen:   len(w),
		wg:     &sync.WaitGroup{},
	}
}

func (w parallelMultiWriter) Write(bytes []byte) (int, error) {
	errCh := make(chan error, w.wLen)
	defer close(errCh)
	w.wg.Add(w.wLen)
	for _, x := range w.w {
		go func(wg *sync.WaitGroup, x io.Writer) {
			defer wg.Done()
			_, err := x.Write(bytes)
			if err != nil {
				errCh <- err
			}
		}(w.wg, x)
	}

	go func() {
		w.wg.Wait()
		errCh <- nil
		close(errCh)
	}()

	count := 0
	select {
	case err, more := <-errCh:
		if !more && err == nil {
			count = len(bytes)
		}

		return count, err
	}

}
