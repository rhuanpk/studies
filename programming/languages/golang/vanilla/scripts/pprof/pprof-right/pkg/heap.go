package pkg

import (
	"log"
	"os/exec"
)

func PprofHeap() {
	if err := exec.Command(
		"go", "tool", "pprof", "http://localhost:6060/debug/pprof/heap",
	).Run(); err != nil {
		log.Println("pprof error:", err)
	}
}
