package criocli

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

const markdownDocTemplate = `
[//]: <> "This file has been autogenerated, please do not edit"

% {{ .App.Name }} 8

# NAME

{{ .App.Name }}{{ if .App.Usage }} - {{ .App.Usage }}{{ end }}

# SYNOPSIS

{{ .App.Name }}
{{ if .SynopsisArgs }}
` + "```" + `
{{ range $v := .SynopsisArgs }}{{ $v }}{{ end }}` + "```" + `
{{ end }}{{ if .App.UsageText }}
# DESCRIPTION

{{ .App.UsageText }}
{{ end }}
**Usage**:

` + "```" + `
{{ .App.Name }} [GLOBAL OPTIONS] command [COMMAND OPTIONS] [ARGUMENTS...]
` + "```" + `
{{ if .GlobalArgs }}
# GLOBAL OPTIONS
{{ range $v := .GlobalArgs }}
{{ $v }}{{ end }}
{{ end }}{{ if .Commands }}
# COMMANDS
{{ range $v := .Commands }}
{{ $v }}{{ end }}{{ end }}
## FILES

**crio.conf** (/etc/crio/crio.conf)
  cri-o configuration file for all of the available command-line options for
  the crio(8) program, but in a TOML format that can be more easily modified
  and versioned.

**policy.json** (/etc/containers/policy.json)
  Signature verification policy files are used to specify policy, e.g. trusted
  keys, applicable when deciding whether to accept an image, or individual
  signatures of that image, as valid.

**registries.conf** (/etc/containers/registries.conf)
  Registry configuration file specifies registries which are consulted when
  completing image names that do not include a registry or domain portion.

**storage.conf** (/etc/containers/storage.conf)
  Storage configuration file specifies all of the available container storage
  options for tools using shared container storage.

# SEE ALSO

crio.conf(5), oci-hooks(5), policy.json(5), registries.conf(5), storage.conf(5)`

func man() *cli.Command {
	return &cli.Command{
		Name:  "man",
		Usage: "Generate the man page documentation.",
		Action: func(c *cli.Context) error {
			cli.MarkdownDocTemplate = markdownDocTemplate
			res, err := c.App.ToMan()
			if err != nil {
				return err
			}
			fmt.Print(res)
			return nil
		},
	}
}

func markdown() *cli.Command {
	return &cli.Command{
		Name:    "markdown",
		Aliases: []string{"md"},
		Usage:   "Generate the markdown documentation.",
		Action: func(c *cli.Context) error {
			cli.MarkdownDocTemplate = markdownDocTemplate
			res, err := c.App.ToMarkdown()
			if err != nil {
				return err
			}
			fmt.Print(strings.TrimSpace(res))
			return nil
		},
	}
}
