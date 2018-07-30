const (
	// Dictionary names.
	active = "active"
	dequed = "dequed"
)

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
