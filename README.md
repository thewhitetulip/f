# f

f is a very very basic http server written in Go, intended to be a replacement for `python -m http.server`, I found that python is a tad bit slower than f.

# How to install

`go get github.com/thewhitetulip/f`

it will download and install f to your `$GOBIN`

If you want to do so manually you can clone the repo and run `go install` on it.

## How to use
f takes one argument, which is the port number where you want to run the server. It uses http.Dir method of net/http package.
`f 9090`
this will start a server on port 9090

## Contributing

If you want to contribute, please raise an issue. 

## Use

f can be used wherever a web server is required.
