# URL Response Hasher

Simple CLI application that accepts URLs, fetch responses and prints the response hash in CLI. Also one can define the number of parallel network calls the system can handle, it is defaulted to 10.
---
[Go Lang](https://go.dev/)
---

### Setup ###

1. Move cwd to the repository path, \
   `cd url_hasher/`

1. Build the application \
   `go build url_hasher.go`

1. One can run the application in either ways, by providing parallel flag or allowing system to set it by default way.\
   `./url_hasher http://www.adjust.com http://google.com` \
   `./url_hasher -parallel 2 http://www.adjust.com http://google.com facebook.com`