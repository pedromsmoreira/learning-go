package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// We’ll parse this example URL, which includes a scheme, authentication info,
	// host, port, path, query params, and query fragment.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Accessing the scheme is straightforward.
	fmt.Println("Scheme: ", u.Scheme)

	// User contains all authentication info; call Username and Password on this for individual values.
	fmt.Println("User: ", u.User)
	fmt.Println("Username: ", u.User.Username())
	p, _ := u.User.Password()
	fmt.Println("Password: ", p)

	// The Host contains both the hostname and the port, if present.
	// Use SplitHostPort to extract them.
	fmt.Println("Host: ", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("Splitted Host: ", host)
	fmt.Println("Splitted Port: ", port)

	// Here we extract the path and the fragment after the #.
	fmt.Println("Path: ", u.Path)
	fmt.Println("Fragment: ", u.Fragment)

	// To get query params in a string of k=v format, use RawQuery.
	// You can also parse query params into a map.
	// The parsed query param maps are from strings to slices of strings,
	// so index into [0] if you only want the first value.
	fmt.Println("Raw Query: ", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("Map: ", m)
	fmt.Println("Key Value: ", m["k"][0])
}
