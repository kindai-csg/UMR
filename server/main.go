package main

import (
	"github.com/kindaidensan/UMR/infrastructure"
)

func main() {
	infrastructure.Router.Run(":8040")
}
