// gommon is the commandline util for generator
package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/generator"
	dlog "github.com/dyweb/gommon/log"
	"github.com/dyweb/gommon/log/handlers/cli"
	"github.com/dyweb/gommon/noodle"
	"github.com/dyweb/gommon/util/logutil"
)

var log, logReg = dlog.NewApplicationLoggerAndRegistry("gommon")

var verbose = false
var (
	version   string
	commit    string
	buildTime string
	buildUser string
	goVersion = runtime.Version()
)

func main() {
	// TODO: most code here are copied from go.ice's cli package, dependency management might break if we import go.ice which also import gommon
	rootCmd := &cobra.Command{
		Use:   "gommon",
		Short: "gommon helpers",
		Long:  "Generate go files for gommon",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if verbose {
				dlog.SetLevel(logReg, dlog.DebugLevel)
				dlog.EnableSource(logReg)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}
	// global flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	// ver
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "print version",
		Long:  "Print current version " + version,
		Run: func(cmd *cobra.Command, args []string) {
			if verbose {
				fmt.Printf("version: %s\n", version)
				fmt.Printf("commit: %s\n", commit)
				fmt.Printf("build time: %s\n", buildTime)
				fmt.Printf("build user: %s\n", buildUser)
				fmt.Printf("go version: %s\n", goVersion)
			} else {
				fmt.Println(version)
			}
		},
	}
	// sub commands
	rootCmd.AddCommand(
		versionCmd,
		genCmd(),
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func genCmd() *cobra.Command {
	gen := &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen"},
		Short:   "generate code based on gommon.yml",
		Run: func(cmd *cobra.Command, args []string) {
			root := "."
			if err := generator.Generate(root); err != nil {
				log.Error(err)
				os.Exit(1)
			}
		},
	}
	root := ""
	name := ""
	pkg := ""
	output := ""
	// TODO: might put noodle as its own top level command?
	noodleCmd := &cobra.Command{
		Use: "noodle",
		Short: "bundle static assets in one directory as a single go file, " +
			"it does not support bundle multiple file into one directory",
		Example: "gommon generate noodle --root assets --output gen/noodle.go --pkg gen --name Assets",
		Run: func(cmd *cobra.Command, args []string) {
			checks := map[string]string{
				"root":   root,
				"name":   name,
				"pkg":    pkg,
				"output": output,
			}
			merr := errors.NewMultiErr()
			for k, v := range checks {
				if v == "" {
					merr.Append(errors.Errorf("%s is required", k))
				}
			}
			if merr.HasError() {
				log.Fatal(merr.Error())
				return
			}
			cfg := noodle.EmbedConfig{
				Src:     root,
				Name:    name,
				Dst:     output,
				Package: pkg,
			}
			if err := noodle.GenerateEmbedFile(cfg); err != nil {
				log.Fatal(err)
				return
			}
			log.Infof("generated %s from %s with package %s and name %s", output, root, pkg, name)
		},
	}
	noodleCmd.Flags().StringVar(&root, "root", "", "path of assets folder")
	noodleCmd.Flags().StringVar(&name, "name", "Asset", "name of generate ")
	noodleCmd.Flags().StringVar(&pkg, "pkg", "gen", "go package of generated file")
	noodleCmd.Flags().StringVar(&output, "output", "noodle.go", "path for generated file")
	gen.AddCommand(noodleCmd)
	return gen
}

func init() {
	logReg.AddRegistry(logutil.Registry())
	dlog.SetHandler(logReg, cli.New(os.Stderr, true))
}
