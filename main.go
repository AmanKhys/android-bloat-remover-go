package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Open the file for reading
	file, err := os.Open("bloat.txt")
	if err != nil {
		println()
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text() // Get the current line
		cmd := exec.Command("adb","shell","pm", "uninstall", "-k", "--user", "0", line)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		// Start the command
		err := cmd.Start()
		if err != nil {
			fmt.Println("Error starting command:", err)
			return
		}

		// Wait for the command to finish
		err = cmd.Wait()
		if err != nil {
			println(line,"NOT DELETED")
		} else {
			fmt.Println("deleted : ", line)
		}	
		println()
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Scan Error:", err)
	}
}
