package helpers

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"runtime"
	"slices"
	"strings"

	"github.com/kglaus/stac-client/internal/stac/models"
)

func SaveAsset(assetObject models.AssetObject) {
	dir, assetRootPath := OperatingSystem()

	path := os.DirFS(dir)
	entries, _ := fs.Glob(path, "stac_downloads")
	if !slices.Contains(entries, "stac_downloads") {
		// TODO Errorhandling via popup
		if err := os.Mkdir(assetRootPath, os.ModePerm); err != nil {
			fmt.Println("error creating folder stac_downloads")
		}
	}

	// TODO Errorhandling via popup
	resp, err := http.Get(assetObject.Href)
	if err != nil {
		fmt.Println("error downloading asset")
	}
	defer resp.Body.Close()

	assetUrlSplit := strings.Split(assetObject.Href, "/")
	assetTitle := assetUrlSplit[len(assetUrlSplit)-1]
	save, err := os.Create(assetRootPath + assetTitle)
	if err != nil {
		fmt.Println("error creating zip file")
	}
	defer save.Close()

	_, err = io.Copy(save, resp.Body)
	if err != nil {
		fmt.Println("error saving zip file")
	}
}

func OperatingSystem() (string, string) {
	var dir string
	var assetRootPath string
	switch runtime.GOOS {
	case "windows":
		dir = os.Getenv("USERPROFILE") + "\\Documents"
		assetRootPath = dir + "\\stac_downloads\\"
	case "linux":
		// TODO Implement linux paths
		dir = "~"
	default:
		// TODO errorhandling
		fmt.Println("Unknown Operating System")
	}
	return dir, assetRootPath
}
