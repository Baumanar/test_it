# Mower take home test

## Implementation choice 

Each `mower` moves simultaneously on the field in a separate routine.
To avoid collision, a `field` structure keeps information about all 
the mowers positions with a map. Before moving, a `mower` checks by using the `field` if its new position is not
the same position as the one of another `mower`. If the movement would cause it
 to run into another `mower`, the instruction is silently discarded.
 
This implementation is relatively simple and makes sure mowers are processed simultaneously. 
- One concern would be the fact that `mowers` have to acquire a lock of the `field` to check and update their position.
 If there were a thousand of mowers on the same field, this would cause execution speed problems. However I thought 
 it was unlikely to be the case for mowers, and one solution would be to use a worker pool. 
 If collisions management was not necessary, we would just have to remove the `field` and not bother updating 
 the mowers positions in it.
- A second concern is that as all mowers run concurrently, if the mowers are likely to run into each other,
  it can be that for two runs of the same program the results differ because one mower took the conflicting position
  first making the runs non-deterministic. As the specification did not specify it, I chose this design.
 
 
 
## Run instruction

- to build the project, run: `make build`
- to run it: `./mowers -f=fixtures/example.txt`,  with `-f` flag pointing to the instructions file
- to run it directly: `go run cmd/mowers/main.go -f=fixtures/example.txt`
- to run tests: `make test`