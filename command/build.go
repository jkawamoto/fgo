package command

import "github.com/urfave/cli"

// BuildOpt defines options for cmdInit.
type BuildOpt struct {
	// Directory to store package files
	Dest string
	// Directory to store brew file
	Brew string

	Version string
}

type BrewParam struct {
	Version     string
	FileName64  string
	FileName386 string
	Hash64      string
	Hash386     string
}

func CmdBuild(c *cli.Context) error {

	opt := BuildOpt{
		Dest:    c.String("dest"),
		Brew:    c.String("brew"),
		Version: c.Args().First(),
	}

	// These codes are not nessesary but urfave/cli doesn't work.
	if opt.Dest == "" {
		opt.Dest = "pkg"
	}
	if opt.Brew == "" {
		opt.Brew = "brew"
	}

	if err := cmdBuild(&opt); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil

}

func cmdBuild(opt *BuildOpt) error {

	// Build and upload via make.

	return nil
}

// func make(tag string) {
//
// 	var cmd *exec.Cmd
// 	if tag != "" {
// 		cmd = exec.Command("make", "build", "release")
// 	} else {
//
// 	}
//
// }
