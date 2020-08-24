# CLI Runner -> CLI Streamer

## Tasks

- Create CLIStreammer which will stream string [Given Seconds] delay using goroutine or any other async way.
    - It should take argument as (`"Title,Message 1,Message 2,Stream Delay,Run Times\nCLI Invoke1,First Msg1,Second Msg2,2,10"`) this format is known as CSV (comma separated values -> Search for it)
      - As per the arguments, it will print `Message1` outputs `CLI Invoke1->First Msg1\n` sleeps for `2` seconds then prints `Message2` `CLI Invoke1->Second Msg2\n` and then repeats the process to message one and repeat until it runs for `10 times`.
  - Create another `CLIRunner` which will run `CLIStreamer` based `arguments`
      - arguments (`"Run,Title,Message 1,Message 2,Stream Delay,Run Times\n2,CLI Invoke1,First Msg1,Second Msg2,2,500\n2,CLI Invoke2,First Msg~1,Second Msg~2,2,500"`) this should run CLIStreamer 2 processcess based on Run argument to 2 in the `CSV` values. For each row it will`run` 2 `process` of `CLIStreamer` parallely with the respective arguments given.
      - Takes `std outs` from the `CLIStreamer` prints to console and also writes it to file (using `mutex`). During the process one needs to lock using `mutex`.

## Modules Initialize

- To initialize go modules: `go mod init github.com/nf/go-lang-testing`
- To install zap: `go get go.uber.org/zap`

## Links

- [go - Anonymous function on struct - Stack Overflow](https://stackoverflow.com/questions/45866477/anonymous-function-on-struct)
- Abstraction : https://play.golang.org/p/WDB8BqVyJS
