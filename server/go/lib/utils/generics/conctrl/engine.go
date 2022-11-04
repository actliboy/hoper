package conctrl

import (
	"context"
	"github.com/actliboy/hoper/server/go/lib/utils/gen"
	"github.com/actliboy/hoper/server/go/lib/utils/generics/slices"
	"log"
	"sync"
	"time"
)

type Engine[KEY comparable, T, W any] struct {
	*BaseEngine[KEY, T, W]
	done        sync.Map
	TasksChan   chan []*Task[KEY, T]
	kindHandler []*KindHandler[KEY, T]
}

type KindHandler[KEY comparable, T any] struct {
	Skip bool
	*time.Ticker
	// TODO 指定Kind的Handler
	HandleFun TaskFunc[KEY, T]
}

func NewEngine[KEY comparable, T, W any](workerCount uint) *Engine[KEY, T, W] {
	return &Engine[KEY, T, W]{
		BaseEngine: NewBaseEngine[KEY, T, W](workerCount),
	}
}

func (e *Engine[KEY, T, W]) SkipKind(kinds ...Kind) *Engine[KEY, T, W] {
	length := slices.Max(kinds) + 1
	if e.kindHandler == nil {
		e.kindHandler = make([]*KindHandler[KEY, T], length)
	}
	if int(length) > len(e.kindHandler) {
		e.kindHandler = append(e.kindHandler, make([]*KindHandler[KEY, T], int(length)-len(e.kindHandler))...)
	}
	for _, kind := range kinds {
		if e.kindHandler[kind] == nil {
			e.kindHandler[kind] = &KindHandler[KEY, T]{Skip: true}
		} else {
			e.kindHandler[kind].Skip = true
		}

	}
	return e
}
func (e *Engine[KEY, T, W]) StopAfter(interval time.Duration) *Engine[KEY, T, W] {
	time.AfterFunc(interval, e.Cancel)
	return e
}

func (e *Engine[KEY, T, W]) Timer(kind Kind, interval time.Duration) *Engine[KEY, T, W] {
	if e.kindHandler == nil {
		e.kindHandler = make([]*KindHandler[KEY, T], int(kind)+1)
	}
	if int(kind)+1 > len(e.kindHandler) {
		e.kindHandler = append(e.kindHandler, make([]*KindHandler[KEY, T], int(kind)+1-len(e.kindHandler))...)
	}
	if e.kindHandler[kind] == nil {
		e.kindHandler[kind] = &KindHandler[KEY, T]{Ticker: time.NewTicker(interval)}
	} else {
		e.kindHandler[kind].Ticker = time.NewTicker(interval)
	}
	return e
}

func (e *Engine[KEY, T, W]) Run(tasks ...*Task[KEY, T]) {
	baseTasks := make([]*BaseTask[KEY, T], 0, len(tasks))
	for _, task := range tasks {
		baseTasks = append(baseTasks, e.NewTask(task))
	}
	e.BaseEngine.Run(baseTasks...)
}

func (e *Engine[KEY, T, W]) NewTask(task *Task[KEY, T]) *BaseTask[KEY, T] {

	if task == nil {
		return nil
	}

	task.Id = gen.GenOrderID()

	var kindHandler *KindHandler[KEY, T]
	if e.kindHandler != nil && int(task.Kind) < len(e.kindHandler) {
		kindHandler = e.kindHandler[task.Kind]
	}

	if kindHandler != nil && kindHandler.Skip {
		return nil
	}

	zeroKey := *new(KEY)

	if task.Key != zeroKey {
		if _, ok := e.done.Load(task.Key); ok {
			return nil
		}
	}
	return &BaseTask[KEY, T]{
		BaseTaskMeta: task.BaseTaskMeta,
		BaseTaskFunc: func(ctx context.Context) {
			if kindHandler != nil && kindHandler.Ticker != nil {
				<-kindHandler.Ticker.C
			}
			tasks, err := task.TaskFunc(ctx)
			if err != nil {
				task.ErrTimes++
				log.Println("执行失败", err)
				log.Println("重新执行,key :", task.Key)
				if task.ErrTimes < 5 {
					e.AsyncAddTask(task.Priority+1, task)
				}
				return
			}
			if task.Key != zeroKey {
				e.done.Store(task.Key, struct{}{})
			}
			if len(tasks) > 0 {
				e.AsyncAddTask(task.Priority+1, tasks...)
			}
			return
		},
	}
}

func (e *Engine[KEY, T, W]) AddTasks(generation int, tasks ...*Task[KEY, T]) {
	for _, req := range tasks {
		if req != nil {
			req.Priority = generation
			e.BaseEngine.AddTask(e.NewTask(req))
		}
	}
}

func (e *Engine[KEY, T, W]) AsyncAddTask(generation int, tasks ...*Task[KEY, T]) {
	go func() {
		for _, task := range tasks {
			if task != nil {
				task.Priority = generation
				e.BaseEngine.AddTask(e.NewTask(task))
			}
		}
	}()
}

func (e *Engine[KEY, T, W]) AddFixedTask(workerId int, task *Task[KEY, T]) {
	if workerId > len(e.fixedWorker)-1 {
		return
	}
	ch := e.fixedWorker[workerId]
	baseTask := &BaseTask[KEY, T]{
		BaseTaskMeta: task.BaseTaskMeta,
		Props:        task.Props,
	}
	baseTask.BaseTaskFunc = func(ctx context.Context) {
		_, err := task.TaskFunc(ctx)
		if err != nil {
			e.wg.Add(1)
			go func() {
				ch <- baseTask
			}()
		}
	}
	e.wg.Add(1)
	go func() {
		ch <- baseTask
	}()
}
