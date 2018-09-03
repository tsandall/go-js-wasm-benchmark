package main

import (
	"syscall/js"
	"time"
)

func addOne(args []js.Value) {
	t0 := time.Now()
	args[1].Invoke(args[0].Int()+1, int64(time.Since(t0)))
}

func main() {
	quit := make(chan struct{})
	js.Global().Set("addOne", js.NewCallback(addOne))
	<-quit
}
