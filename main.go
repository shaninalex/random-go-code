package main

import (
	"different_tasks/tasks/distr"
	"different_tasks/tasks/hashing"
	"flag"
	"log"
)

func main() {

	hash := flag.Bool("hash", false, "Execute hash function")
	hashStr := flag.String("hash_str", "", "Hash string")

	distribute := flag.Bool("distribute", false, "Execute distribute items between nodes")
	distributeMin := flag.Int("dmin", 0, "Minimal amout of hashes set")
	distributeMax := flag.Int("dmax", 0, "Maximum amout of hashes set")

	flag.Parse()

	if *hash {
		log.Println(*hashStr)
		hashing.ExecuteHash()
	}

	if *distribute {
		distr.ExecDistribute(distributeMin, distributeMax)
	}

	if !*hash && !*distribute {
		log.Println("task was not selected")
	}
}
