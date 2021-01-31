package cmd

import (
	"context"
	"io"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/crowdstrike/kubectl-falcon/version"
)

type globalOptions struct {
	debug          bool          // Enable debug output
	commandTimeout time.Duration // Timeout for the command execution
}

func (opts *globalOptions) before(cmd *cobra.Command) error {
	if opts.debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	return nil
}

func NewCmdFalcon(streams genericclioptions.IOStreams) *cobra.Command {
	opts := globalOptions{}

	rootCommand := &cobra.Command{
		Use:          "falcon",
		Short:        "kubectl falcon plug-in",
		Long:         "Various operations with CrowdStrike Falcon sensor on the cluster",
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.before(cmd)
		},
		Version: version.Version,
	}

	// Override default `--version` global flag to enable `-v` shorthand
	var dummyVersion bool
	rootCommand.Flags().BoolVarP(&dummyVersion, "version", "v", false, "prints version")
	rootCommand.PersistentFlags().DurationVar(&opts.commandTimeout, "command-timeout", 0, "timeout for the command execution")
	rootCommand.AddCommand(
		imageRefresh(&opts),
	)

	return rootCommand
}

// commandAction unwraps cobra.Command from RunE interface to avoid misuse
func commandAction(handler func(args []string, stdout io.Writer) error) func(cmd *cobra.Command, args []string) error {
	return func(c *cobra.Command, args []string) error {
		return handler(args, c.OutOrStdout())
	}
}

func (opts *globalOptions) commandTimeoutContext() (context.Context, context.CancelFunc) {
	ctx := context.Background()
	var cancel context.CancelFunc = func() {}
	if opts.commandTimeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, opts.commandTimeout)
	}
	return ctx, cancel
}
