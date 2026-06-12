# Design: Distributed Rate Limiter

Your service runs across 10 nodes. You need to enforce a global rate limit: 1000 requests per minute per user across all nodes.

Design the system. Consider: synchronization strategy, consistency model, failure modes, and performance.
