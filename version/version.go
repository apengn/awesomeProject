package version

import (
	"fmt"
	"os"
	"path"
)

var (
	gitCommit string
	buildTime string
)

func init() {
	fmt.Printf("%s gitCommit:%s  buildTime:%s \n", path.Base(os.Args[0]), gitCommit, buildTime)
}
