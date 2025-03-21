package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func validateDate(dateStr string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	return time.Parse(layout, dateStr)
}

func printHelp() {
	fmt.Println("Usage: git-cheat <subcommand> [options]")
	fmt.Println("\nSubcommands:")
	fmt.Println("  create   - Create a commit with a message and timestamp")
	fmt.Println("  redate   - Change the date of an existing commit")
	fmt.Println("  reauthor - Change the author of an existing commit")
	fmt.Println("\nUse 'git-cheat <subcommand> -h' for more details.")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
	}

	subcommand := os.Args[1]

	switch subcommand {
	case "create":
		createCmd := flag.NewFlagSet("create", flag.ExitOnError)
		message := createCmd.String("m", "", "Commit message (required)")
		date := createCmd.String("d", "", "Date in YYYY-MM-DD HH:MM:SS format (required)")

		createCmd.Usage = func() {
			fmt.Println("Create a commit with a message and timestamp")
			fmt.Println("Usage: git-cheat create -m <message> -d <date>")
			fmt.Println("\nOptions:")
			createCmd.PrintDefaults()
		}

		createCmd.Parse(os.Args[2:])
		if *message == "" || *date == "" {
			createCmd.Usage()
			os.Exit(1)
		}

		if _, err := validateDate(*date); err != nil {
			fmt.Println("Error: Invalid date format. Use YYYY-MM-DD HH:MM:SS")
			os.Exit(1)
		}

		fmt.Printf("Creating commit with message: %s at %s\n", *message, *date)

	case "redate":

		redateCmd := flag.NewFlagSet("redate", flag.ExitOnError)
		commitHash := redateCmd.String("c", "", "Commit hash (required)")
		date := redateCmd.String("d", "", "New date in YYYY-MM-DD HH:MM:SS format (required)")

		redateCmd.Usage = func() {
			fmt.Println("Change the date of an existing commit")
			fmt.Println("Usage: git-cheat redate -c <commit-hash> -d <date>")
			fmt.Println("\nOptions:")
			redateCmd.PrintDefaults()
		}

		redateCmd.Parse(os.Args[2:])
		if *commitHash == "" || *date == "" {
			redateCmd.Usage()
			os.Exit(1)
		}

		if _, err := validateDate(*date); err != nil {
			fmt.Println("Error: Invalid date format. Use YYYY-MM-DD HH:MM:SS")
			os.Exit(1)
		}

		fmt.Printf("Updating commit %s with new date %s\n", *commitHash, *date)

	case "reauthor":
		reauthorCmd := flag.NewFlagSet("reauthor", flag.ExitOnError)
		commitHash := reauthorCmd.String("c", "", "Commit hash (required)")
		name := reauthorCmd.String("n", "", "New author name (required)")
		email := reauthorCmd.String("e", "", "New author email (required)")

		reauthorCmd.Usage = func() {
			fmt.Println("Change the author of an existing commit")
			fmt.Println("Usage: git-cheat reauthor -c <commit-hash> -n <name> -e <email>")
			fmt.Println("\nOptions:")
			reauthorCmd.PrintDefaults()
		}

		reauthorCmd.Parse(os.Args[2:])
		if *commitHash == "" || *name == "" || *email == "" {
			reauthorCmd.Usage()
			os.Exit(1)
		}

		fmt.Printf("Updating commit %s with new author: %s <%s>\n", *commitHash, *name, *email)

	case "-h", "--help", "help":
		printHelp()

	default:
		fmt.Println("Error: Unknown subcommand:", subcommand)
		printHelp()
	}
}
