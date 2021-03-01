package cli

import "github.com/urfave/cli"

var (
	commands = []cli.Command{
		{
			Name:      "ls",
			Usage:     "List installed versions",
			UsageText: "g ls",
			Action:    list,
			Aliases: []string{"list", "lls"},
		},
		{
			Name:      "ls-remote",
			Usage:     "List remote versions available for install",
			UsageText: "g ls-remote [stable|archived|unstable]",
			Action:    listRemote,
			Aliases: []string{"lsr", "rls"},
		},
		{
			Name:      "use",
			Usage:     "Switch to specified version",
			UsageText: "g use <version>",
			Action:    use,
			Aliases: []string{"switch"},
		},
		{
			Name:      "install",
			Usage:     "Download and install a version",
			UsageText: "g install <version>",
			Action:    install,
			Aliases: []string{"i"},
		},
		{
			Name:      "uninstall",
			Usage:     "Uninstall a version",
			UsageText: "g uninstall <version>",
			Action:    uninstall,
			Aliases: []string{"un"},
		},
		{
			Name:      "clean",
			Usage:     "Remove files from the package download directory",
			UsageText: "g clean",
			Action:    clean,
		},
		{
			Name:      "download",
			Usage:     "Download version(s)",
			UsageText: "g dl <version1>,<version2>",
			Action:    download,
			Aliases: []string{"d", "dl"},
		},
	}
)
