package main

import (
	connectrpcnode "different_tasks/tasks/connectRpcNode"
	"different_tasks/tasks/distr"
	"different_tasks/tasks/hashing"
	"flag"
)

func main() {

	hash := flag.Bool("hash", false, "Execute hash function")

	distribute := flag.Bool("distribute", false, "Execute distribute items between nodes")
	distributeMin := flag.Int("dmin", 0, "Minimal amout of hashes set")
	distributeMax := flag.Int("dmax", 0, "Maximum amout of hashes set")

	connectGrpc := flag.Bool("connect_grpc", false, "Connecto to node with grpc")

	flag.Parse()

	if *hash {
		hashing.ExecuteHash()
	}

	if *distribute {
		distr.ExecDistribute(distributeMin, distributeMax)
	}

	if *connectGrpc {
		connectrpcnode.ExecuteConnect()
	}
}
