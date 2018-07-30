const (
	// Dictionary names.
	active = "active"
	dequed = "dequed"
)

// Queue represents a named queue.
type Queue string

// Task represents a task in a queue.
type Task struct {
	ID    uint64 `json:"id"`    // Task's globally unique ID assigned by taskq.
	Queue Queue  `json:"queue"` // Task's queue.
	Body  []byte `json:"body"`  // Task's client data.
}

func (h EnQHandler) Rcv(msg beehive.Msg, ctx beehive.RcvContext) error {
	// TODO: assign a unique ID.
	enq := msg.Data().(Enque)

	dict := ctx.Dict(active)
	key := string(enq.Task.Queue)

	var tasks []Task
	if val, err := dict.Get(key); err == nil {
		tasks = val.([]Task)
	}

	tasks = append(tasks, enq.Task)
	return dict.Put(key, tasks)
}
