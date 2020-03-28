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

- Lots of really useful information here - https://blog.golang.org/using-go-modules
- And also here - https://github.com/golang/go/wiki/Modules#why-does-go-mod-tidy-record-indirect-and-test-dependencies-in-my-gomod

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

Next, the command `go list -m all` lists the current module and all its dependencies:

`
messing-around
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
`

- Both go.mod and go.sum should be checked into version control
- https://github.com/golang/go/wiki/Modules#is-gosum-a-lock-file-why-does-gosum-include-information-for-module-versions-i-am-no-longer-using

"
Is 'go.sum' a lock file? Why does 'go.sum' include information for module versions I am no longer using?

No, go.sum is not a lock file. The go.mod files in a build provide enough information for 100% reproducible builds.

For validation purposes, go.sum contains the expected cryptographic checksums of the content of specific module versions.
See the FAQ below for more details on go.sum (including why you typically should check in go.sum) as well as
the "Module downloading and verification" section in the tip documentation.

In part because go.sum is not a lock file, it retains cryptographic checksums for module versions even after you stop
using a module or particular module version. This allows validation of the checksums if you later resume using
something, which provides additional safety.

In addition, your module's go.sum records checksums for all direct and indirect dependencies used in a build (and
hence your go.sum will frequently have more modules listed than your go.mod).
"

UPGRADING DEPENDENCIES (when using go.mod file):

- https://github.com/golang/go/wiki/Modules#how-to-upgrade-and-downgrade-dependencies
- https://blog.golang.org/using-go-modules

- `go get` will fetch the latest version of the requested package
	- e.g. go get example.com/package

- You can specify which version you want using the @ symbol:
	- e.g. go get rsc.io/sampler@v1.3.1


WHEN I USE GO MODULES AND I GET NEW DEPENDENCIES, WHERE DO THEY GO?

- Brilliant article - https://medium.com/@adiach3nko/package-management-with-go-modules-the-pragmatic-guide-c831b4eaaf31
- Basically they go into $GOPATH/pkg/mod directory

STRUCTURE OF GO WORKSPACE:

	- bin = where executables get saved to
	- src = your source code. But now that go modules is here, you don't have to have your source code for every application in here.
	- pkg = basically if you're using go modules, this is where your dependencies get saved

*/