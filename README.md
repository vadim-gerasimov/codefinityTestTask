# Test Task Codefinity

## Task 2 (Container Implementation in Golang)

This Golang program demonstrates a basic implementation of creating and running a container-like environment from scratch. The code provides a simplified version of Docker's functionality, highlighting the key steps involved in setting up a container.

```golang
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// Function to create and run a container
func createAndRunContainer(imagePath, containerName, containerRootDir string) {
	// Create container's root filesystem
	createRootFS(containerRootDir)

	// Populate the container's root filesystem with necessary files and directories
	populateRootFS(containerRootDir, imagePath)

	// Set up mount points for /proc, /sys, /dev, etc.
	mountFilesystems(containerRootDir)

	// Change root to the container's root directory
	chroot(containerRootDir)

	// Set up network configurations (if required)
	setupNetwork()

	// Create and configure cgroups for resource isolation
	createCgroups(containerName)

	// Set CPU affinity for the container
	setCPUSets(containerName)

	// Execute the container's init process
	executeInitProcess(containerRootDir)
}

func createRootFS(containerRootDir string) {
	// Create the directory for the container's root filesystem
	if err := os.MkdirAll(containerRootDir, 0755); err != nil {
		fmt.Println("Error creating container's root filesystem directory:", err)
		os.Exit(1)
	}
}

func populateRootFS(containerRootDir, imagePath string) {
	// Implement the logic to populate the container's root filesystem
	// with necessary files and directories from the specified image path.
}

func mountFilesystems(containerRootDir string) {
	// Implement the logic to mount /proc, /sys, /dev, etc. from the host
	// into the container's root filesystem at appropriate mount points.
}

func chroot(containerRootDir string) {
	// Implement the logic to change the container's root directory to the
	// specified directory using the `syscall.Chroot` function.
}

func setupNetwork() {
	// Implement the logic to set up network configurations inside the container.
}

func createCgroups(containerName string) {
	// Implement the logic to create and configure cgroups for the container
	// using the `cgroup` package or other relevant APIs.
}

func setCPUSets(containerName string) {
	// Implement the logic to set CPU affinity for the container using
	// the `cpuset` package or other relevant APIs.
}

func executeInitProcess(containerRootDir string) {
	// Implement the logic to execute the container's init process (e.g., /sbin/init)
	// using the `syscall.Exec` function.
	initProcessPath := "/sbin/init"
	initProcessArgs := []string{initProcessPath}
	if err := syscall.Exec(initProcessPath, initProcessArgs, os.Environ()); err != nil {
		fmt.Println("Error executing container's init process:", err)
		os.Exit(1)
	}
}

// Entry point of the program
func main() {
	imagePath := "/path/to/rootfs/image"
	containerName := "my_container"
	containerRootDir := "/path/to/container/root"

	// Create and run the container
	createAndRunContainer(imagePath, containerName, containerRootDir)
}
```