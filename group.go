package group

type Actor interface {
	Start() error
	Stop(error)
}

type Group struct {
	actors []Actor
}

func (g *Group) Add(actor Actor) {
	g.actors = append(g.actors, actor)
}

func (g *Group) Run() error {
	n := len(g.actors)
	if n == 0 {
		return nil
	}
	errors := make(chan error, n)
	for _, a := range g.actors {
		go func(a Actor) {
			errors <- a.Start()
		}(a)
	}
	err := <-errors
	for _, a := range g.actors {
		a.Stop(err)
	}
	for i := 1; i < cap(errors); i++ {
		<-errors
	}
	return err
}
