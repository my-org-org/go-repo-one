package dependency

import (
	"fmt"
	"github.com/my-org-for-test/go-repo-one/internal"
)

func PrintDependencyMessage() {
	fmt.Println("go-repo-one:  Rel v4.0.0. OLD NAME")
    internal.InternalFunction()
}
