# Design: Leader Election

You have a service that must run a background task on exactly one node at a time. Design a leader election mechanism.

Discuss: split-brain prevention, network partition handling, graceful handoff, and liveness detection. Provide a Go sketch using a distributed lock (e.g., Redis Redlock, etcd, or Consul).
