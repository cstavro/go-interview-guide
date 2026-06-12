# Problem: Worker Pool with Graceful Shutdown

Build a worker pool that processes jobs from a channel. It must support Start, Submit, and Stop (graceful). On Stop, it should finish all in-flight jobs but reject new submissions. Use a context.Context for cancellation.
