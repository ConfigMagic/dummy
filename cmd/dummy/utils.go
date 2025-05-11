package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runDockerCompose(args ...string) {
	cmd := exec.Command("docker-compose", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Ошибка: %v\n%s\n", err, out)
		os.Exit(1)
	}
	fmt.Printf("✅ Команда выполнена: docker-compose %v\n", args)
}
