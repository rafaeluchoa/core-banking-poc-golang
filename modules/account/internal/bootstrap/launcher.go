package bootstrap

import (
	"log"
	"sync"
)

type App interface {
	Run(done chan error)
}

type Launcher struct {
	wg   *sync.WaitGroup
	apps []*App
}

func NewLauncher() *Launcher {

	l := &Launcher{
		wg:   new(sync.WaitGroup),
		apps: make([]*App, 0),
	}

	return l
}

func (s *Launcher) Run(app App) {
	s.apps = append(s.apps, &app)
	s.wg.Add(1)
	done := make(chan error)

	go func() {
		defer s.wg.Done()
		app.Run(done)
	}()

	err := <-done
	if err != nil {
		log.Panicln(err, "Stoping")
	}
}

func (s *Launcher) Wait() {
	s.wg.Wait()
}
