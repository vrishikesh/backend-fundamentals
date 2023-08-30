package main

import (
	pubsub "github.com/vrishikesh/backend-fundamentals/pub_sub"
)

func main() {
	// req_res.ReqRes()
	// sync_async.Sync()
	// sync_async.Async()
	// push.Websockets()
	// poll.ShortPolling()
	// poll.LongPolling()
	// sse.SSE()
	// pubsub.Publish()
	pubsub.Consume()
}
