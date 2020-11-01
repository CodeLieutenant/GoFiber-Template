package io

import (
	"io"
	"sync"
)

type parallelMultiWriter struct {
	w    []io.Writer
	wLen int
	wg   *sync.WaitGroup
}

func NewParallelMultiWriter(w ...io.Writer) io.Writer {
	return parallelMultiWriter{
		w:    w,
		wLen: len(w),
		wg:   &sync.WaitGroup{},
	}
}

func (w parallelMultiWriter) Write(bytes []byte) (count int, err error) {
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

	err, more := <-errCh
	if !more && err == nil {
		count = len(bytes)
	}

	return count, err
}
