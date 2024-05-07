package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/gofiber/fiber/v2"
)

type TestCase struct {
	Input          int `json:"input"`
	ExpectedOutput int `json:"expected_output"`
}

type Problem struct {
	Description string     `json:"description"`
	TestCases   []TestCase `json:"test_cases"`
}

type Data struct {
	Problems []Problem `json:"problems"`
}

type RunRequest struct {
	Code string `json:"Gocode"`
	Fonk string `json:"Fonksiyon adi"`
}

type RunResponse struct {
	Stdout string `json:"Sonuc"`
}

func main() {
	app := fiber.New()

	app.Post("/rungo", handleRunGo)

	log.Fatal(app.Listen(":8080"))
}

func handleRunGo(c *fiber.Ctx) error {
	var request RunRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	file, err := os.Open("testcases.json")
	if err != nil {
		fmt.Println("Dosya açma hatası:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer file.Close()

	var data Data
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("JSON verisini decode etme hatası:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var response string

	for _, problem := range data.Problems {
		if problem.Description == request.Fonk {
			fmt.Println("Problem:", problem.Description)
			allCorrect := true
			for _, testCase := range problem.TestCases {
				input := testCase.Input
				expectedOutput := testCase.ExpectedOutput

				actualOutput := runDocker(request, c, input)
				clearStringActualOutput := strings.Replace(actualOutput.Error(), "\n", "", -1)

				i, err := strconv.Atoi(clearStringActualOutput)
				if i != expectedOutput {
					allCorrect = false
				}
				if err != nil {
					fmt.Println(err)
				}
			}
			if allCorrect {
				response = "Tebrikler, tüm test vakalarını geçtiniz."
			} else {
				response = "Bazı test vakalarını geçemediniz."
			}
		} else {
			response = "Problem bulunamadı."
		}
	}

	resp := RunResponse{
		Stdout: response,
	}
	return c.JSON(resp)
}

func runDocker(request RunRequest, c *fiber.Ctx, input int) error {
	tmpDir := "/home/melkor" // MELKOR YERINE PCDEKI KULLANICI ADI YAZILMALI
	tmpFile := filepath.Join(tmpDir, "tmpfile.go")
	if err := os.WriteFile(tmpFile, []byte(stringEkle(request.Code, request.Fonk, input)), 0644); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Geçici dosyaya yazma hatası: %v", err))
	}

	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Docker istemcisi oluşturma hatası: %v", err))
	}
	defer cli.Close()

	contImage := "docker.io/library/golang"

	read, err := cli.ImagePull(ctx, contImage, image.PullOptions{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Docker imajı çekme hatası: %v", err))
	}
	defer read.Close()

	io.Copy(os.Stdout, read)

	path := "/home/melkor/tmpfile.go" // MELKOR YERINE PCDEKI KULLANICI ADI YAZILMALI
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
	}, nil, nil, "")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Docker konteyneri oluşturma hatası: %v", err))
	}

	if err := cli.ContainerStart(ctx, createResp.ID, container.StartOptions{}); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Docker konteyneri başlatma hatası: %v", err))
	}

	statusCh, errCh := cli.ContainerWait(ctx, createResp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Docker konteyneri bekleniyor hatası: %v", err))
		}
	case <-statusCh:
	}

	output, err := cli.ContainerLogs(ctx, createResp.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Docker konteyneri loglarını alma hatası: %v", err))
	}
	defer output.Close()

	var cikti bytes.Buffer
	var hata bytes.Buffer
	_, err = stdcopy.StdCopy(&cikti, &hata, output)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Docker konteyneri loglarını kopyalama hatası: %v", err))
	}

	ciktiStr := cikti.String()

	fmt.Println("Çıktı:", ciktiStr)
	return errors.New(ciktiStr)
}

func stringEkle(kod string, fonkad string, kontrolInput int) string {
	data, err := os.ReadFile("template.txt")
	if err != nil {
		fmt.Println("Dosya okunurken hata oluştu:", err)
		return ""
	}

	text := string(data)
	text = text + kod
	text = strings.Replace(text, "tempfonk", fonkad, -1)
	text = strings.Replace(text, "tempvar", strconv.Itoa(kontrolInput), -1)
	fmt.Println(text)
	return text
}
