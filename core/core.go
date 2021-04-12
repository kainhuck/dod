package core

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func runCmd(command string,view bool, args ...string){
	docker, err := exec.LookPath(command)
	if err != nil {
		log.Fatalf("can't find `%s` in path: %v", command, err)
	}

	cmd := exec.Command(docker, args...)
	if view {
		cmd.Stdout = os.Stdout
	}
	err = cmd.Run()
	if err != nil {
		log.Fatalf("run `%s` failed: %v",cmd.String(), err)
	}
	return
}

func runCmdOut(command string, args ...string)(output string){
	docker, err := exec.LookPath(command)
	if err != nil {
		log.Fatalf("can't find `%s` in path: %v", command, err)
	}

	cmd := exec.Command(docker, args...)
	bts, err := cmd.Output()
	if err != nil {
		log.Fatalf("run `%s` failed: %v",cmd.String(), err)
	}
	return string(bts)
}

func splitSpace(str string) []string {
	result := make([]string, 0)
	firstSplit := strings.Split(str, " ")
	for _, each := range firstSplit{
		if len(strings.TrimSpace(each)) > 0{
			result = append(result, each)
		}
	}
	return result
}

func listImages()(imageIds []string){
	output := runCmdOut("docker", "image", "ls", "-a")
	//REPOSITORY         TAG               IMAGE ID       CREATED          SIZE
	//imagename          latest            f59452732a0d   22 minutes ago   22.2MB
	lineS := strings.Split(output, "\n")
	lines := lineS[1:len(lineS)-1]
	for _, line := range lines{
		if len(line) == 0{
			break
		}

		imageIds = append(imageIds, splitSpace(line)[2])
	}

	return
}

func listContainers()(containerIds []string){
	output := runCmdOut("docker", "container", "ls", "-a")
	//CONTAINER ID   IMAGE       COMMAND       CREATED          STATUS                      PORTS     NAMES
	//a7194a46f443   imagename   "/bin/e-o…"   29 minutes ago   Exited (2) 28 minutes ago             opcua
	lineS := strings.Split(output, "\n")
	lines := lineS[1:len(lineS)-1]
	for _, line := range lines{
		if len(line) == 0{
			break
		}

		containerIds = append(containerIds, splitSpace(line)[0])
	}

	return
}

func DeleteImage(imageId string, view bool){
	runCmd("docker", view,"image", "rm", imageId)
}

func DeleteContainer(containerId string, view bool){
	runCmd("docker", view, "container", "rm", "-f", containerId)
}

func DeleteAllImage(view bool){
	imageIds := listImages()
	for _, id := range imageIds{
		DeleteImage(id, view)
	}
	fmt.Printf("Delete All(%d) Image Success!\n", len(imageIds))
}

func DeleteAllContainer(view bool){
	containerIds := listContainers()
	for _, id := range containerIds{
		DeleteContainer(id, view)
	}
	fmt.Printf("Delete All(%d) Container Success!\n", len(containerIds))
}

func DeleteAll(view bool){
	DeleteAllContainer(view)
	DeleteAllImage(view)
}