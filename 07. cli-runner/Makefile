run: 
	# Should runner with the given arguments to run the streamer

run-runner: 
	# does the same as run
	make run

test-runner:
	make build-streamer
	go run cli-runner.go -run "Run,Title,Message 1,Message 2,Stream Delay,Run Times\n2,CLI Invoke1,First Msg1,Second Mod Msg2,2,2\n2,CLI Invoke2,First Msg~1,Second Msg~2,2,2"

build-streamer: 
	go build cli-streamer.go

run-CLIRunner:
	go run .\cli-runner.go -run "Run,Title,Message 1,Message 2,Stream Delay,Run Times\n2,CLI Invoke1,First Msg1,Second Msg2,2,500\n2,CLI Invoke2,First Msg~1,Second Msg~2,2,500"
