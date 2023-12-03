//go:build event_trace

package engine

import (
	"github.com/ZenLiuCN/fn"

	"log"
)

func debug(args ...any) {
	log.Println(append(args, fn.CallerN(1))...)
}
func debugN(n uint, args ...any) {
	log.Println(append(args, fn.CallerN(n))...)
}
