# gRPC vs REST OpenSearch Benchmark Experiment

## Overview

This repository contains controlled benchmarks comparing REST and gRPC for OpenSearch-style search workloads under different execution models.

The goal is to isolate where performance differences actually come from:

- tooling overhead  
- connection setup cost  
- protocol efficiency  
- concurrency scaling  

Rather than claiming one approach is “faster”, this experiment breaks performance into measurable layers.

---

## What this experiment is not

- Not a framework comparison  
- Not a language benchmark  
- Not a production deployment guide  

It is a controlled performance decomposition of REST vs gRPC request handling.

---

## Benchmark Design

The experiment is divided into three layers to progressively eliminate noise.

---

## 1. Naive Benchmark (Tooling + Request Overhead)

### What it measures
- One-off request execution  
- No connection reuse (REST)  
- grpcurl CLI overhead (gRPC)  

CLI-based gRPC is not representative of actual gRPC performance due to:

- reflection overhead  
- process spawning  
- serialization via external tool  

---

## 2. Persistent Connection Benchmark (Protocol Efficiency)

### What it measures
- Reused HTTP connection (REST)  
- Reused gRPC channel (gRPC)  
- Same request executed repeatedly  

### Why it matters

This removes tooling overhead and isolates:

- HTTP/JSON vs Protobuf efficiency  
- network + serialization cost  
- connection reuse benefits  

---

## 3. Concurrent Load Benchmark (Production Behavior)

Real systems are not single-request systems. They are concurrency systems.

### What it measures
- 50+ parallel requests  
- sustained load simulation  
- connection pooling behavior  

Under load:

- REST incurs higher CPU + parsing overhead per request  
- gRPC benefits from multiplexed connections and binary framing  
- scalability differences become more visible than raw latency differences  

---

## How to Run

Please follow `STARTUP.md` to set up the prerequisites for this experiment.

---

## Structure

- `1-naive-benchmark/` → CLI-based request comparisons  
- `2-persistent-benchmark/` → reused connection benchmarks  
- `3-concurrent-benchmark/` → load simulation tests  
- `shared/` → proto, clients, server setup  
