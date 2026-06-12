# Design: Job Queue with Exactly-Once Semantics

Build a job queue system where each job must be processed exactly once. Jobs can fail and be retried.

Discuss: idempotency, deduplication, at-least-once delivery, and the impossibility of exactly-once in distributed systems. Provide a design that achieves "effectively exactly-once" in practice.
