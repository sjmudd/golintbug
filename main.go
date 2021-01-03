package main

// GoLintBug is really an int
type GoLintBug int

const (
	intConst1                 = 1             // int constant
	intConst2                 = 2             // int constant
	stringConst               = "stringConst" // string constant
	GoLintBugConst1 GoLintBug = iota          // GoLintBugConst1 constant
	GoLintBugConst2                           // GoLintBugConst2 constant
	GoLintBugConst3                           // GoLintBugConst3 constant
)

func main() {
}
