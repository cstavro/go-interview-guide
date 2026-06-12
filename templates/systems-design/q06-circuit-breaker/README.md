# Design: Circuit Breaker

Implement a circuit breaker pattern for calling a downstream service. The breaker should open after N failures in a window, half-open after a timeout, and close after a successful probe.

Provide the state machine, the metrics collection, and the Go code for a generic breaker.
