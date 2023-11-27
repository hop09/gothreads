package gothreads

import (
        "sync"
)

type ThreadPoolExecutor struct {
        workers   int
        taskQueue chan func()
        wg        sync.WaitGroup
}

func NewThreadPoolExecutor(workers int, maxQueueSize int) *ThreadPoolExecutor {
        return &ThreadPoolExecutor{
                workers:   workers,
                taskQueue: make(chan func(), maxQueueSize),
        }
}

func (e *ThreadPoolExecutor) Start() {
        for i := 0; i < e.workers; i++ {
                go func() {
                        for task := range e.taskQueue {
                                task()
                        }
                }()
        }
}

func (e *ThreadPoolExecutor) Submit(task func()) {
        e.wg.Add(1)
        e.taskQueue <- func() {
                task()
                e.wg.Done()
        }
}

func (e *ThreadPoolExecutor) Stop() {
        close(e.taskQueue)
        e.wg.Wait()
}
