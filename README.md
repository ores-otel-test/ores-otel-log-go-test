# ores-otel-log-go-test

Exact-head conformance harness for **go**.

This repository tests both `ores-otel/ores.otel.log` and `ORESoftware/next-loggers.ts` using explicit commit SHAs.
The required native command is recorded in `conformance.json`: `go test -race ./...`.
