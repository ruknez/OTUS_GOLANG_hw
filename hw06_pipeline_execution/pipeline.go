package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := tmpStage(in, done)
	for _, stage := range stages {
		if stage == nil {
			continue
		}
		out = stage(tmpStage(out, done))
	}

	return tmpStage(out, done)
}

func tmpStage(in In, done In) Out {
	bichannel := make(Bi)
	go func() {
		defer func() {
			close(bichannel)
			for range in {
			}
		}()
		for {
			select {
			case <-done:
				return
			case val, ok := <-in:
				if !ok {
					return
				}
				bichannel <- val
			}
		}
	}()
	return bichannel
}
