<div align="center">

<img src="assets/logo.png" alt="voxigo" width="200" />

**A WebRTC-native, audio-first conversational-AI framework for Go.**

[![CI](https://github.com/nuxflix/voxigo/actions/workflows/ci.yml/badge.svg)](https://github.com/nuxflix/voxigo/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/nuxflix/voxigo.svg)](https://pkg.go.dev/github.com/nuxflix/voxigo)
[![Go Report Card](https://goreportcard.com/badge/github.com/nuxflix/voxigo)](https://goreportcard.com/report/github.com/nuxflix/voxigo)
![Go version](https://img.shields.io/github/go-mod/go-version/nuxflix/voxigo)
[![Release](https://img.shields.io/github/v/release/nuxflix/voxigo?sort=semver)](https://github.com/nuxflix/voxigo/releases)
[![License: BSD-2-Clause](https://img.shields.io/badge/license-BSD--2--Clause-blue.svg)](LICENSE)

</div>

---

voxigo builds real-time voice agents: audio comes in over WebRTC, flows through a
pipeline of processors (transcription → reasoning → speech), and audio goes back
out — with [RTVI](https://docs.pipecat.ai/client/introduction) on the data
channel so existing RTVI clients interoperate.

> **Status:** early work in progress. APIs are unstable and will change.

## Install

```sh
go get github.com/nuxflix/voxigo
```

## Examples

```sh
# Echo bot — speak into the browser, hear yourself back over WebRTC.
go run ./examples/echo                 # then open http://localhost:8080

# Voice bot — Deepgram (STT) → Anthropic (LLM) → ElevenLabs (TTS).
export DEEPGRAM_API_KEY=...
export ANTHROPIC_API_KEY=...
export ELEVENLABS_API_KEY=...
go run ./examples/voicebot             # then open http://localhost:8080
```

## License & attribution

voxigo is a Go port of [Pipecat](https://github.com/pipecat-ai/pipecat),
distributed under the same **BSD 2-Clause License**. The upstream copyright —
*Copyright (c) 2024–2026, Daily* — is preserved verbatim in [`LICENSE`](LICENSE);
see [`NOTICE`](NOTICE) for details. voxigo is an independent project, not
affiliated with or endorsed by Daily.
