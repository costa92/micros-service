package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"k8s.io/kubernetes/cmd/genutils"
)

func main() {
	// 使用 os.Args 而不是 "flags"，因为 "flags" 会搞乱 man 页面！
	path := "docs/"
	if len(os.Args) == 2 {
		path = os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [output directory]\n", os.Args[0])
		os.Exit(1)
	}

	outDir, err := genutils.OutDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get output directory: %v\n", err)
		os.Exit(1)
	}

	onexctl := &cobra.Command{
		Use: "",
	}

	// 设置 onexctl 使用的环境变量，以确保输出一致，
	// 无论我们在哪里运行。
	os.Setenv("HOME", "/home/username")

	doc.GenMarkdownTree(onexctl, outDir)
}
