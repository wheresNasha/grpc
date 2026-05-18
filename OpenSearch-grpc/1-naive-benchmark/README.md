# Naive Benchmark: REST vs gRPC (CLI-Based)

## Overview

This benchmark compares REST and gRPC search latency using simple command-line tooling:

- REST via `curl`
- gRPC via `grpcurl`

The objective is to establish a baseline comparison under minimal optimization conditions.
This phase intentionally includes tooling and process overhead, making it useful for understanding initial developer-facing performance rather than production-grade transport efficiency.

---

## Benchmark Script

Full reproducible commands are available in:

`naive-benchmark.ps1`

---

## Methodology

### Phase 1: Identical Dataset Preparation

A shared benchmark dataset was created to ensure both REST and gRPC queries operated against the same OpenSearch index and documents.

### Included:
- Benchmark index creation
- Sample document insertion
- Controlled search target

---

### Phase 2: REST Baseline Measurement

REST benchmarking was performed using:

- `curl`
- PowerShell `Measure-Command`
- repeated warm queries

### Purpose:
To establish:
- cold-start latency
- warm average latency
- CLI overhead baseline

---

### Phase 3: gRPC Baseline Measurement

gRPC benchmarking was performed using:

- `grpcurl`
- protobuf reflection
- equivalent search query
- repeated warm runs

### Purpose:
To compare:
- equivalent query execution
- CLI process cost
- protocol behavior under naive tooling conditions

---

## Key Observations

### REST
- Cold requests showed significantly higher startup latency
- Warm repeated requests stabilized substantially

### gRPC
- Functional parity with REST was achieved
- CLI-based gRPC showed higher average latency in this setup
- grpcurl overhead materially impacted results

---

## Observed Environment-Specific Results

| Protocol | Approximate Avg Latency |
|----------|--------------------------|
| REST     | ~50–60 ms                |
| gRPC     | ~80–90 ms                |

---

> **Disclaimer:**  
> Benchmark values are environment-specific and may vary depending on:
>
> - hardware
> - system load
> - OpenSearch build version
> - JVM warmup state
> - network conditions
> - grpcurl version
> - local machine performance
>
> These results should be interpreted as experimental observations, not universal performance guarantees.

---

## Interpretation

### Why REST appeared faster here:
- `curl` incurs lower process overhead
- simpler CLI execution path
- no protobuf reflection layer
- no grpcurl startup penalties

---

### Why gRPC appeared slower:
- grpcurl process spawning
- reflection/service discovery overhead
- CLI serialization/deserialization cost
- not representative of persistent service-to-service gRPC

---

## Important Conclusion

This benchmark primarily measures:

- Tooling overhead
- Process startup cost
- First-level execution path

### It does NOT fully represent:
- Persistent gRPC efficiency
- Real-world microservice performance
- Long-lived connections
- HTTP/2 multiplexing advantages
- optimized protobuf client implementations

---

## Primary Takeaway

> CLI-based benchmarks can distort protocol comparisons.

Naive measurements are valuable for:
- developer setup
- tooling validation
- initial experimentation

But deeper benchmarking is required to fairly isolate:

- protocol efficiency
- serialization cost
- scalability behavior

---

## Supporting Output

Include benchmark screenshots throughout this document:

```md
![Sample Output](../sampleOutput/picture1.png)
