package api

import (
	"net/http"
	"fmt"
	"io"
	"os"
)

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}


func	InstallMlx(path string) (error) {
	if (path == "") {return fmt.Errorf("no Path")}
	err := downloadFile(path + "/MLX.zip", "https://cdn.ochouati.me/content/MLX.zip")
	if (err != nil) {return fmt.Errorf("error while wownloading MLX")}
	return nil
}
