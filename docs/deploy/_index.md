---
title: Deploying
weight: 5
---

# Deploying

The default build is cgo-free, so a bot ships as a single static binary. What it
still needs at run time is the ONNX Runtime, for VAD and end-of-turn detection,
and optionally RNNoise.

- **[Deploy with Docker](docker.md)**: the `nuxflix/voxigo-build` build base and the
  distroless `nuxflix/voxigo` runtime base (which bundles the native libraries),
  plus a copyable two-stage Dockerfile.

One task per session, and a task dies with its connection: build the pipeline
inside your `/offer` handler and cancel on close. Scaling is therefore a matter of
how many concurrent sessions one process holds. See the
[benchmarks](https://github.com/nuxflix/voxigo-benchmarks) for measured numbers.
