package utils

import "os/exec"

func OpenBrowser(url string) error {
	url, error := exec.LookPath("xdg-open")
}
