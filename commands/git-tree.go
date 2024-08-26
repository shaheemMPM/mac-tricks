package commands

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"golang.design/x/clipboard"
)

func DisplayGitAwareTree() {
	if !isGitRepository() {
		fmt.Println("Error: Not a git repository")
		return
	}

	files, err := getGitTrackedFiles()
	if err != nil {
		fmt.Println("Error getting git tracked files:", err)
		return
	}

	root := &Node{name: ".", isDir: true}
	for _, file := range files {
		addToTree(root, strings.Split(file, string(os.PathSeparator)))
	}

	var buf bytes.Buffer
	printTree(root, "", true, &buf)

	// Print to console
	fmt.Print(buf.String())

	// Copy to clipboard
	err = clipboard.Init()
	if err != nil {
		fmt.Println("Error initializing clipboard:", err)
		return
	}
	clipboard.Write(clipboard.FmtText, buf.Bytes())
	fmt.Println("\nTree structure copied to clipboard.")
}

type Node struct {
	name     string
	isDir    bool
	children []*Node
}

func isGitRepository() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()
	return err == nil
}

func getGitTrackedFiles() ([]string, error) {
	cmd := exec.Command("git", "ls-files")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var files []string
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		files = append(files, scanner.Text())
	}

	return files, scanner.Err()
}

func addToTree(root *Node, path []string) {
	current := root
	for i, part := range path {
		found := false
		for _, child := range current.children {
			if child.name == part {
				current = child
				found = true
				break
			}
		}
		if !found {
			newNode := &Node{name: part, isDir: i < len(path)-1}
			current.children = append(current.children, newNode)
			current = newNode
		}
	}
}

func printTree(node *Node, prefix string, isLast bool, buf *bytes.Buffer) {
	if node.name != "." {
		buf.WriteString(prefix)
		if isLast {
			buf.WriteString("└── ")
			prefix += "    "
		} else {
			buf.WriteString("├── ")
			prefix += "│   "
		}
		buf.WriteString(node.name + "\n")
	}

	for i, child := range node.children {
		printTree(child, prefix, i == len(node.children)-1, buf)
	}
}
