package ssh

import (
	"errors"
	"fmt"

	"github.com/debu99/cicd-runner/common"
	"github.com/debu99/cicd-runner/executors"
	"github.com/debu99/cicd-runner/helpers/ssh"
)

type executor struct {
	executors.AbstractExecutor
	sshCommand ssh.Client
}

func (s *executor) Prepare(options common.ExecutorPrepareOptions) error {
	err := s.AbstractExecutor.Prepare(options)
	if err != nil {
		return fmt.Errorf("prearing AbstractExecutor: %w", err)
	}

	s.Println("Using SSH executor...")
	if s.BuildShell.PassFile {
		return errors.New("SSH doesn't support shells that require script file")
	}

	if s.Config.SSH == nil {
		return errors.New("missing SSH configuration")
	}

	s.Debugln("Starting SSH command...")

	// Create SSH command
	s.sshCommand = ssh.Client{
		Config: *s.Config.SSH,
		Stdout: s.Trace,
		Stderr: s.Trace,
	}

	s.Debugln("Connecting to SSH server...")
	err = s.sshCommand.Connect()
	if err != nil {
		return fmt.Errorf("ssh command Connect() error: %w", err)
	}

	return nil
}

func (s *executor) Run(cmd common.ExecutorCommand) error {
	err := s.sshCommand.Run(cmd.Context, ssh.Command{
		Environment: s.BuildShell.Environment,
		Command:     s.BuildShell.GetCommandWithArguments(),
		Stdin:       cmd.Script,
	})
	if exitError, ok := err.(*ssh.ExitError); ok {
		exitCode := exitError.ExitCode()
		err = &common.BuildError{Inner: err, ExitCode: exitCode}
	}
	return err
}

func (s *executor) Cleanup() {
	s.sshCommand.Cleanup()
	s.AbstractExecutor.Cleanup()
}

func init() {
	options := executors.ExecutorOptions{
		DefaultCustomBuildsDirEnabled: false,
		DefaultBuildsDir:              "builds",
		DefaultCacheDir:               "cache",
		SharedBuildsDir:               true,
		Shell: common.ShellScriptInfo{
			Shell:         "bash",
			Type:          common.LoginShell,
			RunnerCommand: "cicd-runner",
		},
		ShowHostname: true,
	}

	creator := func() common.Executor {
		return &executor{
			AbstractExecutor: executors.AbstractExecutor{
				ExecutorOptions: options,
			},
		}
	}

	featuresUpdater := func(features *common.FeaturesInfo) {
		features.Variables = true
		features.Shared = true
	}

	common.RegisterExecutorProvider("ssh", executors.DefaultExecutorProvider{
		Creator:          creator,
		FeaturesUpdater:  featuresUpdater,
		DefaultShellName: options.Shell.Shell,
	})
}
