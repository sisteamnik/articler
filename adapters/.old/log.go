package adapters

import (
	"github.com/fatih/color"
	"log"
)

func logErr(e error, file string, line int) {
	log.Println(color.RedString("[%s] %s:%d", e, file, line))
}
