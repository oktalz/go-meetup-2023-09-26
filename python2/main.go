package main

import (
	"bufio"
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// START OMIT
	python := `
		
def sumFirstN(x):
  total = sum(range(1, x + 1))
  print(f"The sum of the first {x} numbers is: {total}")
  
# Call the function with x = 10
sumFirstN(10)	

	`
	//END OMIT
	_ = python
	err := pythonRun(python)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func pythonRun(code string) error {
	os.WriteFile("example.py", []byte(code), 0o644)

	ctx := context.Background()
	cmd := exec.CommandContext(ctx, "python", "./example.py")
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	cmd.Dir = dir
	cmd.Env = os.Environ()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) //nolint:forbidigo
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	scanner = bufio.NewScanner(stdErr)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) //nolint:forbidigo
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
