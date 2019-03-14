package version

import (
	"fmt"
	"os"
	"path"
)

func init() {
	fmt.Printf("%s gitCommit:%s \n", path.Base(os.Args[0]))
}
