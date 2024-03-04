package dependency

import (
	"fmt"
	"github.com/my-org-org/go-repo-one/internal"
)

func PrintDependencyMessage() {
	fmt.Println("go-repo-one:  Rel v5.0.0. NEW NAME")
    internal.InternalFunction()
}
