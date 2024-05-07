package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func main() {
	dosyadi := os.Args[1] + ".go"
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error :NewCLientWithOpt: %v", err)
	}
	defer cli.Close()
	contImage := "docker.io/library/golang"

	read, err := cli.ImagePull(ctx, contImage, image.PullOptions{})
	if err != nil {
		log.Fatalf("Error :ImagePull: %v", err)
	}
	defer read.Close()

	io.Copy(os.Stdout, read)

	path, err := FindFile(dosyadi)
	if err != nil {
		log.Fatalf("Error :FindFile: %v", err)
	}
	fmt.Println(path)
	createResp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: contImage,
		Cmd:   []string{"go", "run", "/app" + path},
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: path,
				Target: "/app" + path,
			},
		},
	},
		nil, nil, "")
	if err != nil {
		log.Fatalf("Error :ContainerCreate: %v", err)
	}

	if err := cli.ContainerStart(ctx, createResp.ID, container.StartOptions{}); err != nil {
		log.Fatalf("Error :ContainerStart: %v", err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, createResp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			log.Fatalf("Error :ContainerWait: %v", err)
		}
	case <-statusCh:
	}

	output, err := cli.ContainerLogs(ctx, createResp.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		log.Fatalf("Error :ContainerLogs: %v", err)
	}
	stdcopy.StdCopy(os.Stdout, os.Stderr, output)
	defer output.Close()

	if _, err := io.Copy(os.Stdout, output); err != nil {
		log.Fatalf("Error :Stdout: %v", err)
	}

}

func FindFile(dosyadi string) (string, error) {
	var realPath string

	err := filepath.Walk("/home/melkor/Desktop/gofiles", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == dosyadi {
			realPath = path
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	if realPath == "" {
		return "", fmt.Errorf("dosya bulunamadı: %s", dosyadi)
	}

	return realPath, nil
}
