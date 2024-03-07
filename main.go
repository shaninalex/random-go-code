package main

import (
	"different_tasks/tasks/distr"
	"different_tasks/tasks/hashing"
	"flag"
	"log"
)

func main() {

	hash := flag.Bool("hash", false, "Hash function")
	distribute := flag.Bool("distribute", false, "Distribute items between nodes")
	flag.Parse()

	if *hash {
		hashing.ExecuteHash()
	}

	if *distribute {
		distr.ExecDistribute()
	}

	if !*hash && !*distribute {
		log.Println("task was not selected")
	}
}
