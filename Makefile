# Usage:
# make hashing
hashing:
	go run . --hash

# Usage:
# make dmin=5 dmax=15 distr
distr:
	go run .  --dmin=$(dmin) --dmax=$(dmax) --distribute

# Usage:
# make connect
connect:
	go run . --connect_grpc

# Usage:
# make task_compare
task_compare:
	go run . --task_compare