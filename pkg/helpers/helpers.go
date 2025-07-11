package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/satnamSandhu2001/stackjet/pkg"
	"github.com/satnamSandhu2001/stackjet/pkg/colors"
)

// Bool, Int, Int16, String, Uint are helpers to create a pointer to a values
func Bool(v bool) *bool       { return &v }
func Int(v int) *int          { return &v }
func Int16(v int16) *int16    { return &v }
func String(v string) *string { return &v }
func Uint(v uint) *uint       { return &v }

// GenerateStackDirPath generates a directory path for a stack based on its repo-url in default base dir
func GenerateStackDirPath(repoUrl string) string {
	// Remove .git suffix if present
	repoUrl = strings.TrimSuffix(repoUrl, ".git")

	// Extract repo name from URL
	parts := strings.Split(repoUrl, "/")
	repoName := parts[len(parts)-1]

	// Slugify
	slug := strings.TrimSpace(strings.ToLower(
		regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(repoName, "-"),
	))

	// Trim to 30 characters max
	if len(slug) > 30 {
		slug = slug[:30]
	}
	slug = strings.Trim(slug, "-") // clean up trailing dash

	uid := uuid.New().String()[:3]
	basePath := pkg.Config().DEFAULT_STACKS_BASE_DIR
	finalPath := filepath.Join(basePath, slug+"__"+uid)
	return finalPath
}

// accepts only: [npm | yarn | pnpm] start or [npm | yarn | pnpm] run <script>
func ValidateNodeStartCommand(command string) error {
	command = strings.TrimSpace(command)

	// disallow chaining, pipes
	if strings.ContainsAny(command, "&|;") {
		return errors.New("chaining or piping is not allowed in start command")
	}

	// allow only npm|yarn|pnpm start / run script-name
	validPattern := regexp.MustCompile(`^(npm|yarn|pnpm)\s+(start|run\s+[a-zA-Z0-9:_-]+)$`)
	if !validPattern.MatchString(command) {
		return errors.New("start command must be 'npm|yarn|pnpm start' or 'run <script>'")
	}

	return nil

}

// ask for user input string via command line
func AskForString(reader *bufio.Reader, label string, defaultVal string) string {
	fmt.Print(colors.PrimaryBold("? ") + colors.Bold(label))
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return defaultVal
	}
	return input
}

// ask for user input int via command line
func AskForInt(reader *bufio.Reader, label string, defaultVal int) int {
	fmt.Print(colors.PrimaryBold("? ") + colors.Bold(label))
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return defaultVal
	}
	if val, err := strconv.Atoi(input); err == nil {
		return val
	}
	return defaultVal
}

// ask for user input bool via command line
func AskForBool(reader *bufio.Reader, label string, defaultVal bool) bool {
	fmt.Printf(colors.PrimaryBold("? ")+"%s [%s/%s]: ", colors.Bold(label), "y", "n")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "" {
		return defaultVal
	}
	return input == "y" || input == "yes"
}
func AskForOptionNumber(reader *bufio.Reader, label string, defaultChoice int, options []string) string {
	fmt.Println(colors.PrimaryBold("? ") + colors.Bold(label))
	for i, option := range options {
		fmt.Printf("   [%d] %s\n", i+1, option)
	}
	fmt.Printf("   Enter Please enter a choice [Default choice(%d)]: ", defaultChoice)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err == nil && choice > 0 && choice <= len(options) {
		return options[choice-1]
	}
	return options[defaultChoice-1]
}
