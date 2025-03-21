package main

import (
	"fmt"
	"os"
	"os/exec"
)

func createCommit(message, date string) error {
	if err := runGitCommand("add", "."); err != nil {
		return fmt.Errorf("error staging changes: %w", err)
	}

	env := []string{
		fmt.Sprintf("GIT_AUTHOR_DATE=%s", date),
		fmt.Sprintf("GIT_COMMITTER_DATE=%s", date),
	}
	if err := runGitCommandWithEnv(env, "commit", "-m", message); err != nil {
		return fmt.Errorf("error creating commit: %w", err)
	}

	return nil
}

func redateCommit(commitHash, date string) error {
	filter := fmt.Sprintf(`
        if [ $GIT_COMMIT = "%s" ]
        then
            export GIT_AUTHOR_DATE="%s"
            export GIT_COMMITTER_DATE="%s"
        fi
    `, commitHash, date, date)

	if err := runGitCommand("filter-branch", "--env-filter", filter, "HEAD"); err != nil {
		return fmt.Errorf("error redating commit: %w", err)
	}

	return nil
}

func reauthorCommit(commitHash, authorName, authorEmail string) error {
	filter := fmt.Sprintf(`
        if [ $GIT_COMMIT = "%s" ]
        then
            export GIT_AUTHOR_NAME="%s"
            export GIT_AUTHOR_EMAIL="%s"
            export GIT_COMMITTER_NAME="%s"
            export GIT_COMMITTER_EMAIL="%s"
        fi
    `, commitHash, authorName, authorEmail, authorName, authorEmail)

	if err := runGitCommand("filter-branch", "--env-filter", filter, "HEAD"); err != nil {
		fmt.Println("Error reauthoring commit:", err)
		return fmt.Errorf("error reauthoring commit: %w", err)
	}

	return nil
}

func runGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func runGitCommandWithEnv(env []string, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Env = append(os.Environ(), env...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
