# Myths About Protobuf

Usually, people strongly associate Protobuf with gRPC, and many assume that if you use Protobuf, you are using gRPC. That is not always true.

gRPC requires Protobuf to function properly.  
If you want to use gRPC, Protobuf is mandatory.

However, Protobuf can be used independently without gRPC.

Protobuf works as a serialization format over both:

- REST (HTTP)
- gRPC (HTTP/2)

The difference is in the transport protocol, not the serialization format.

- Protobuf and JSON = Serialization formats
- REST and gRPC = Transport/protocol layers

## Why Protobuf Instead of JSON?

The main advantage of Protobuf over JSON is compact payload size.

JSON is human-readable and widely supported, but it is verbose:

- Repeated field names (the payload has field names)
- More structural characters (quotes, braces)
- Thus increasing the payload size

Larger payloads mean:
- More bytes to send over the transport layer
- More network bandwidth usage
- More Resource consumption like RAM and CPU
- More bytes to be send -> thereby increases latency 

### Why Not Just Compress JSON?

Compression also consumes resources:

- CPU cycles for compression
- CPU cycles for decompression
- Additional RAM usage

Nothing is free.
This is where JSON inefficiencies become more noticeable at scale.

---

## 1. Protobuf over gRPC (Most Common)

- gRPC uses HTTP/2
- Protobuf is the default message format
- Strongly typed APIs from `.proto` files
- Supports streaming, deadlines, and multiplexing

### Example

```proto
service UserService {
  rpc GetUser(GetUserRequest) returns (User);
}
```
Client calls a method like a local function call (it doesn't look like a Remote Procedure Call, Thats the beauty!!! ):

```txt
userService.getUser(id)
```

### Under the hood

- HTTP/2
- Binary Protobuf payload
- RPC semantics

### Good for

- Internal microservices
- Low latency/High throughput systems

---

## 2. Protobuf over REST (HTTP APIs)

You can use Protobuf payloads with normal REST endpoints.

### Example Request

```http
POST /users
Content-Type: application/x-protobuf
Accept: application/x-protobuf
```

Body = serialized Protobuf bytes.

### Example Server

```java
@PostMapping(
  value = "/users",
  consumes = "application/x-protobuf",
  produces = "application/x-protobuf"
)
public UserResponse createUser(
    @RequestBody UserRequest request
) {
    ...
}
```

### You still have

- HTTP verbs (GET, POST, etc.)
- URLs/resources
- REST semantics

But the payload is binary Protobuf instead of JSON.

### Good for

- REST APIs needing smaller payloads
- Backward compatibility with REST systems
- Mobile or high-performance APIs

### Tradeoffs

- Harder to debug as protobuf format is not human readable than JSON
- Less browser-friendly
- Many API tools assume JSON

---

## Quick Comparison

| Feature | REST + JSON | REST + Protobuf | gRPC |
|--------|-------------|-----------------|------|
| Transport | HTTP/1.1 or HTTP/2 | HTTP | HTTP/2 |
| Payload | JSON | Protobuf | Protobuf |
| Human readable | Yes | No | No |
| Streaming | Limited | Limited | Native |
| Code generation | Optional | Optional | Strong |
| Browser support | Excellent | Okay | Limited (needs gRPC-Web) |
| Performance | Medium | High | Very High |

---

## Common Architecture

- External/public APIs → REST + JSON
- Internal service-to-service → gRPC + Protobuf

---

## Important Note

gRPC can expose HTTP/JSON endpoints through transcoding using API gateways.
This allows one `.proto` contract to support both: REST and gRPC

---

## Real-World Examples

### LinkedIn
**LinkedIn Integrates Protocol Buffers With Rest.li for Improved Microservices Performance**
- Blog: [LinkedIn Engineering Blog](https://www.linkedin.com/blog/engineering/infrastructure/linkedin-integrates-protocol-buffers-with-rest-li-for-improved-m)
- Video: [How LinkedIn improved latency by 60%](https://youtu.be/DgNncAnhkIY?si=yWLVmc0ZRGrHu6Xt)
---
### Atlassian (Jira Cloud)
**Using Protobuf to Make Jira Cloud Faster**
- Blog: [Atlassian Engineering Blog](https://www.atlassian.com/blog/atlassian-engineering/using-protobuf-to-make-jira-cloud-faster)
- Video: [Saved 55% cost, 75% CPU, at 33x speed by moving from JSON to Protobuf](https://youtu.be/roNgG4QVjTU?si=Ol7f1V-L0cUF2Ghu)
