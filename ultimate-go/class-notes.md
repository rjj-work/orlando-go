# Ultimate Go 2019.02.09-10 w/Bill Kennedy

https://github.com/ardanlabs/gotraining/blob/master/topics/courses/go/README.md

Rules:
1. you may NOT got from pointer to value sematics
1. Mutex lock/unlock in same routine and displayed on same page

## Engineering Trade-offs
Go: resource + runtime
Rust: no run-time

## Costs of Latency
- First Level: Networking
-- Networking
-- OS calls

- Second Level: Internal
-- gargage collection
-- sync
-- (heap) allocation

- Third Level: Data Access
-- patterns of access

- Forth Level: Algorithm Inefficiency


# GC
Have GCPCT knob to adjust the GAP

- Go 1.11 (or earlier)
-- Tri-color; marking; concurretn
-- Stop The World to setup Write Barrier

- GAP is the available space
-- target: same as in use

- Limite GC by
-- remove non-productive allocation
--- reduce the amount of memory per allocation
-- lower the number of allocations
-- not use pointer into collection, use index

Ralph J. Jackson
rjj.work@gmail.com
407.319.1390


## 2019.02.010 Day 2:
Polymorphism means a piece of code changes it behavior depending on the concrete data it is working on.

## Sofware Design

Done: 70% code coverage.  100% happy path
- part 2 of `done`: Are we de-coupled?

1. Developers don't know what `done` means
1. Need developers that know how to break down a problem.
- How to break a problem into smaller things?
1. Write code in layers
- usually 3 layers is sufficient
-- primitive layer: solves 1 problem very well
-- lower level API: may be unexported, but sites on-top of the primitive layer
-- higher level API: management or "makes thigns easy to do"

Douglas Crawford:
- typing systems cause you problems

### scheduling opportunities
- go routine start
- GC
- system calls
-- Net Poller: winOS, *nix, macOS all have excellent networking asyn support
--- IOCP, epoll, K-events

- blocking calls: sync / orches

### Go Routines
This is parenting.
1. you have to understand your workload
1. How to implemention sync and orchestration
- sync: share state
- orch: data passing between parties

- mutex: synch
- changels, wait groups: orchestration


#### Race detector
- on macOS use `-cpu 24`

### Channels
Channels provide `signaling`
Channels support orchestration

Q1: Is the signal I'm sending need guarantee it was received.
- guarantee is receive *happens* before the send completes.

