package hello_test

import (
	hello "messing-around"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := hello.Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

/*

- Created go mod with `go mod init messing-around`
- This generated a go.mod file:

`
module messing-around

go 1.13
`

Then when I either ran the test above, it automatically pulled in the dependencies that I needed:

The go.mod file now looks like this:

`
module messing-around

go 1.13

require rsc.io/quote v1.5.2
`

A go.sum file has also been created:

`
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c h1:qgOY6WgZOaTkIIMiVjBQcw93ERBE4m30iBm00nkL0i8=
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3Y=
rsc.io/quote v1.5.2/go.mod h1:LzX7hefJvL54yjefDEDHNONDjII0t9xZLPXsUe+TKr0=
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=
rsc.io/sampler v1.3.0/go.mod h1:T1hPZKmBbMNahiBKFy5HrXp6adAjACjK9JXDnKaTXpA=
`

Interestingly enough, I can also now access hello.Hello() in the test above. before I couldn't. So it seems like go modules manages that too.

Next, the command go list -m all lists the current module and all its dependencies:

`
messing-around
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
`


...TODO Up to 4:09 in adding dependencies...


*/