package main

import (
	"github.com/smith-30/qiita-adv-calendar/domain/model"
	"github.com/smith-30/qiita-adv-calendar/domain/service"
	"github.com/smith-30/qiita-adv-calendar/helper/env"
)

var (
	name  = "go"
	count = 4
)

func main() {
	env.LoadEnv()
	cap := 25 * count

	// ready grid.
	ag := service.NewAggregater(cap)
	gridUpdateCh := ag.UpdateGrid(cap)

	// fetch grids each calendar.
	cs := model.NewCalendars(name, count)
	cs.FetchGrids(gridUpdateCh)

	// wait to send grid.
	cs.Wait()
	close(gridUpdateCh)
	// wait aggregate.
	ag.Wait()
}
