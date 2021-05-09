package thirdWeek

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type App struct {
	ctx context.Context
	cancel func()
	srvs  []srv
}

func NewApp(ctx context.Context) *App {
	newCtx, cancel := context.WithCancel(ctx)
	return &App{
		ctx: newCtx,
		cancel: cancel,
	}
}

func (app *App) Run() error {
	errGroup, ctx := errgroup.WithContext(app.ctx)
	for _, srv := range app.srvs {
		srv := srv
		errGroup.Go(func() error {
			<-ctx.Done()
			return srv.Stop()
		})
		errGroup.Go(func () error {
			return srv.Start()
		})
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	errGroup.Go(func () error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				app.Stop()
			}
		}
	})
	if err := errGroup.Wait(); err != nil {
		return err
	}
	return nil
}

func (app *App) Stop() {
	if app.cancel != nil {
		app.cancel()
	}
}