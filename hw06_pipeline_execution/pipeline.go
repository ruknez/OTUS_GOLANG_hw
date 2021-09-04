package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make(In)
	bichannel := make(Bi)
	go func() {
		for {
			select {
			case <-done:
				close(bichannel)
				return
			case val, ok := <-in:
				if ok {
					bichannel <- val
				} else {
					close(bichannel)
					return
				}
			}
		}
	}()

	for i, stage := range stages {
		if i == 0 {
			out = stage(bichannel)
		} else {
			out = stage(out)
		}
	}

	return out
}
