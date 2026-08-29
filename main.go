package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	_ "embed"
)

//go:embed LICENSE
var license []byte

//go:embed README.md
var readme []byte

func help() {
	fmt.Fprintf(os.Stderr, "%s\n", readme)
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func args(out http.ResponseWriter, _ *http.Request) {
	for i, s := range os.Args {
		fmt.Fprintf(out, "ARGV[%d]: \"%s\"\n", i, s)
	}
}

func doc(out http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(out, "%s", readme)
}

func env(out http.ResponseWriter, _ *http.Request) {
	// May contain garbage characters that smell like binary.
	// Send BOM to prevent downloading instead of displaying.
	fmt.Fprintf(out, "\xef\xbb\xbf")

	for _, e := range os.Environ() {
		fmt.Fprintf(out, "%s\n", e)
	}
}

func headers(out http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(out, "%v: %v\n", name, h)
		}
	}
}

func process(out http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(out, "pid: %d\n", os.Getpid())
	fmt.Fprintf(out, "parent: %d\n", os.Getppid())
}

func quit(_ http.ResponseWriter, _ *http.Request) {
	os.Exit(0)
}

func rights(out http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(out, "%s", license)
}

func main() {
	http.HandleFunc("/", doc)
	http.HandleFunc("/args", args)
	http.HandleFunc("/env", env)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/process", process)
	http.HandleFunc("/quit", quit)
	http.HandleFunc("/rights", rights)

	addr := flag.String("bind", "localhost", "Listener address")
	port := flag.Uint("port", 9080, "Listener port")
	flag.Usage = help
	flag.Parse()

	sock := fmt.Sprintf("%s:%d", *addr, *port)
	http.ListenAndServe(sock, nil)
}
