package dependency

import (
	"fmt"
	"github.com/my-org-for-test/go-repo-one/internal"
)

func PrintDependencyMessage() {
	fmt.Println("go-repo-one:  Rel v3.0.0")
    internal.InternalFunction()
}
