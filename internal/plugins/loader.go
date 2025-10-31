package plugins

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var pluginDirs = []string{
	filepath.Join(os.Getenv("HOME"), ".local", "share", "hycli", "plugins"),
}

func makePluginCmd(path string) *cobra.Command {
	name := filepath.Base(path)
	name = strings.TrimSuffix(name, filepath.Ext(name))
	name = strings.TrimPrefix(name, "hycli-")

	cmd := &cobra.Command{
		Use:   name,
		Short: "External plugin: " + name,

		DisableFlagParsing: true,
		RunE: func(c *cobra.Command, args []string) error {

			execCmd := exec.Command(path, args...)
			execCmd.Stdin = os.Stdin
			execCmd.Stdout = os.Stdout
			execCmd.Stderr = os.Stderr

			if err := execCmd.Run(); err != nil {
				return err
			}
			return nil
		},
	}

	cmd.ValidArgsFunction = func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {

		completeArgs := []string{"__complete"}
		completeArgs = append(completeArgs, args...)

		if toComplete != "" {
			completeArgs = append(completeArgs, toComplete)
		} else {

			completeArgs = append(completeArgs, "")
		}

		out, err := runPluginCapture(path, completeArgs)
		if err != nil {

			return nil, cobra.ShellCompDirectiveError
		}

		lines := splitLines(out)
		directive := 0

		if len(lines) > 0 && len(lines[len(lines)-1]) > 0 && lines[len(lines)-1][0] == ':' {
			last := lines[len(lines)-1]

			if v, parseErr := parseDirective(last[1:]); parseErr == nil {
				directive = v
				lines = lines[:len(lines)-1]
			}
		}

		return lines, cobra.ShellCompDirective(directive)
	}

	cmd.SetUsageFunc(func(cmd *cobra.Command) error { return errors.New("") })

	return cmd
}

func Discover() []*cobra.Command {
	found := map[string]bool{}
	var cmds []*cobra.Command

	for _, dir := range pluginDirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, f := range entries {
			if f.IsDir() {
				continue
			}
			name := f.Name()
			trim := strings.TrimSuffix(name, filepath.Ext(name))
			trim = strings.TrimPrefix(trim, "hycli-")
			if found[trim] {
				continue
			}
			path := filepath.Join(dir, f.Name())

			if !isExecutable(path) {
				continue
			}
			found[trim] = true
			cmds = append(cmds, makePluginCmd(path))
		}
	}

	pathDirs := strings.SplitSeq(os.Getenv("PATH"), string(os.PathListSeparator))
	for pd := range pathDirs {
		matches, _ := filepath.Glob(filepath.Join(pd, "hycli-*"))
		for _, m := range matches {
			name := filepath.Base(m)
			trim := strings.TrimSuffix(name, filepath.Ext(name))
			trim = strings.TrimPrefix(trim, "hycli-")
			if found[trim] {
				continue
			}
			if !isExecutable(m) {
				continue
			}
			found[trim] = true
			cmds = append(cmds, makePluginCmd(m))
		}
	}

	return cmds
}

func runPluginCapture(path string, args []string) (string, error) {
	cmd := exec.Command(path, args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return stdout.String(), nil
}

func splitLines(s string) []string {
	scanner := bufio.NewScanner(strings.NewReader(s))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func parseDirective(s string) (int, error) {
	var v int
	_, err := fmt.Sscanf(s, "%d", &v)
	return v, err
}

func isExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	if runtime.GOOS == "windows" {
		ext := strings.ToLower(filepath.Ext(path))
		return ext == ".exe" || ext == ".bat" || ext == ".cmd" || ext == ".ps1"
	}
	return info.Mode()&0111 != 0
}
