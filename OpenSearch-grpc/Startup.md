# STARTUP.md

# OpenSearch gRPC Benchmark Setup (Windows)

> **Note:** This setup was performed on a Windows machine.  
> Commands may vary slightly for macOS/Linux.

---

## Step 1: Clone OpenSearch Repository

If you do not already have the repository:

```bash
git clone https://github.com/opensearch-project/OpenSearch.git
cd OpenSearch

If already cloned:

```bash
cd OpenSearch
git pull origin main
```

---

## Step 2: Initial Source Run

Start OpenSearch using Gradle:

```bash
.\gradlew run
```

![Sample Output](sampleOutput/picture1.png)
---

## Step 3: Verify REST Server

In another terminal window:

```bash
curl localhost:9200
```

### Expected Output:

You should receive a JSON response confirming OpenSearch is running.

```md
![REST Verification](images/rest-verification.png)
```

---

## Step 4: Check gRPC Modules

```bash
dir modules
```

### Verify:

Ensure gRPC-related modules are present.

```md
![gRPC Modules](images/grpc-modules.png)
```

---

## Step 5: Enable gRPC Configuration

Official documentation followed:

[https://docs.opensearch.org/latest/api-reference/grpc-apis/index/#how-to-use-grpc-apis](https://docs.opensearch.org/latest/api-reference/grpc-apis/index/#how-to-use-grpc-apis)

### Config Path Used:

```bash
build\testclusters\runTask-0\distro\3.7.0-ARCHIVE\config\opensearch.yml
```

### Verify gRPC Config:

```bash
findstr "grpc" build\testclusters\runTask-0\distro\3.7.0-ARCHIVE\config\opensearch.yml
```

### Expected Output:

```yaml
aux.transport.types: [transport-grpc]
aux.transport.transport-grpc.port: 9400-9500
```

---

## Step 6: Verify gRPC Port Binding

```bash
netstat -ano | findstr 9400
```

### Observation:

Initially, the plugin may be installed but the gRPC service may not fully expose the port.

---

## Step 7: Recommended Approach (Windows Distro Build)

For stable gRPC functionality, source-run complexity was skipped in favor of packaged distro build.

Cancel Gradle run, then from repo root:

```bash
.\gradlew :distribution:archives:windows-zip:assemble
```

---

## Step 8: Navigate to Built Distribution

```bash
cd distribution\archives\windows-zip\build\install\opensearch-3.7.0-SNAPSHOT
```

---

## Step 9: Start OpenSearch Packaged Distribution

```bash
bin\opensearch.bat
```

---

## Step 10: Confirm gRPC Server Startup

### Expected Terminal Output:

```txt
Started gRPC server on port PortsRange{portRange='9400-9500'}
publish_address {127.0.0.1:9400}, bound_addresses {0.0.0.0:9400}
```

### Port Verification:

```bash
netstat -ano | findstr 9400
```

### Expected:

```txt
TCP    0.0.0.0:9400    LISTENING
```

```md
![gRPC Running](images/grpc-running.png)
```

---

## Final Active Endpoints

### REST

```txt
http://localhost:9200
```

### Native Transport

```txt
localhost:9300
```

### gRPC

```txt
localhost:9400
```

---

## Step 11: Final REST Validation

```bash
curl http://localhost:9200
```

### Expected:

A full OpenSearch JSON server response confirming successful setup.

```md
![Final REST Check](images/final-rest-check.png)
```

---

# Setup Complete

Your environment is now ready for:

* Naive Benchmark
* Persistent Benchmark
* Concurrent Benchmark
