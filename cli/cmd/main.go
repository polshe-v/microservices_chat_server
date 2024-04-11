package main

import (
	"github.com/polshe-v/microservices_chat_server/cli/cmd/root"
	"github.com/polshe-v/microservices_common/pkg/closer"
)

func main() {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	root.Execute()
}
