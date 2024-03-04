package dependency

import (
	"fmt"
	"github.com/my-org-for-test/go-repo-one/internal"
)

func PrintDependencyMessage() {
	fmt.Println("Rel v1.0.0")
    internal.InternalFunction()
}
