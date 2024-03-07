Set of random scripts.

Execute:

```bash
# install dependencies
go mod tidy

# Usage:
# check hashing block transaction
make hashing

# distribute "transactions" between nodes to parse
make dmin=5 dmax=15 distr

# connect to grpc node
make connect
```