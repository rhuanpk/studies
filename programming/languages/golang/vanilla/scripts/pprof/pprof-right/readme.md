## Running

Simply `go run ./` and go to <http://localhost:6060/debug/pprof/>.

- While program is running run this to get CPU profile:
```sh
go tool pprof 'http://localhost:6060/debug/pprof/profile?seconds=5'
```

- In the end of program put some code to stop it then run this to get all other profiles (while the program is running):
```sh
for type in heap goroutine allocs block cmdline goroutine heap mutex threadcreate; do
	curl -fsSLo "pprof.$type.pb.gz" "http://localhost:6060/debug/pprof/$type"
done
```

- After get profiles run this to preview them:
```sh
go tool pprof -http=:6060 ~/pprof/profile.pb.gz
```

Create this function in some exposed (or not) package in your project:
```go
func PprofHeap() {
	if err := exec.Command(
		"go", "tool", "pprof", "http://localhost:6060/debug/pprof/heap",
	).Run(); err != nil {
		log.Println("pprof error:", err)
	}
}
```
Then call this function on desired places in your code to get **heap** profile dynamically.
