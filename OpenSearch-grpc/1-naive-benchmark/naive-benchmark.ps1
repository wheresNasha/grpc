# Phase 1: Create benchmark index
curl -X PUT "http://localhost:9200/perf-test"

# Insert sample document
curl -X POST "http://localhost:9200/perf-test/_doc/1" `
-H "Content-Type: application/json" `
-d "{\"title\":\"OpenSearch grpc benchmark\",\"value\":100}"

# Phase 2: Cold REST benchmark
Measure-Command {
    curl "http://localhost:9200/perf-test/_search?q=title:OpenSearch"
}

# Warm REST benchmark (10 runs)
$restTimes = 1..10 | ForEach-Object {
    (Measure-Command {
        curl "http://localhost:9200/perf-test/_search?q=title:OpenSearch" | Out-Null
    }).TotalMilliseconds
}

$restTimes
"Average REST latency: $(([math]::Round(($restTimes | Measure-Object -Average).Average,2))) ms"

# List gRPC services
grpcurl.exe -plaintext localhost:9400 list

# Describe SearchService
grpcurl.exe -plaintext localhost:9400 describe org.opensearch.protobufs.services.SearchService

# Run equivalent gRPC search query
grpcurl.exe `
  -plaintext `
  -d '{\"index\":[\"perf-test\"],\"q\":\"title:OpenSearch\"}' `
  localhost:9400 `
  org.opensearch.protobufs.services.SearchService/Search

# Warm gRPC benchmark (10 runs)
$grpcTimes = 1..10 | ForEach-Object {
    (Measure-Command {
        grpcurl.exe `
          -plaintext `
          -d '{\"index\":[\"perf-test\"],\"q\":\"title:OpenSearch\"}' `
          localhost:9400 `
          org.opensearch.protobufs.services.SearchService/Search | Out-Null
    }).TotalMilliseconds
}

$grpcTimes
"Average gRPC latency: $(([math]::Round(($grpcTimes | Measure-Object -Average).Average,2))) ms"
