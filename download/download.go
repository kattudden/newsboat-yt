package download

import (
	"errors"
	"fmt"
	"kattudden/newsboat-yt/config"
	"kattudden/newsboat-yt/utils"
	"os/exec"
)

func YoutubeVideo(url string) error {
	conf, _ := config.New()

	err := utils.EnsureDirectory(conf.DownloadPath)
	if err != nil {
		return errors.New("failed to create download directory!")
	}

	downloadBinary := "yt-dlp"

	output := fmt.Sprintf("%s/%%(title)s.%%(ext)s", conf.DownloadPath)
	args := []string{"-S", "filesize~100M", "-f", "'bv+(ba[format_note*=original]/ba)'", "-o", output, url}
	cmd := exec.Command(downloadBinary, args...)

	fmt.Println(cmd.String())

	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return errors.New("failed to download video!")
	}

	return nil
}
