package src

import (
	"github.com/google/wire"
	"github.com/wild46cat/golib/utils"
	"sync"
)

func main() {
	wire.Build()
	utils.SToI64("aa")
}

func holdTheWorld() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
