package ezquake

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cavaliergopher/grab/v3"
)

type AssetManager struct {
	baseDir string
}

func NewAssetManager(baseDirAbsPath string) *AssetManager {
	return &AssetManager{
		baseDir: baseDirAbsPath,
	}
}

func (a *AssetManager) AbsPath(path ...string) string {
	paths := append([]string{a.BaseDir()}, path...)
	return filepath.Join(paths...)
}

func (a *AssetManager) BaseDir() string {
	return a.baseDir
}

func (a *AssetManager) DemosDir() string {
	return a.AbsPath(filepath.Join("qw", "demos"))
}

func (a *AssetManager) DownloadMap(name string) error {
	if a.HasMap(name) {
		return nil
	}

	return a.Download(MapUrl(name), a.MapsDir())
}

func (a *AssetManager) Download(url string, destDir string) error {
	resp, err := grab.Get(destDir, url)

	if http.StatusNotFound == resp.HTTPResponse.StatusCode {
		return errors.New(fmt.Sprintf("NOT FOUND: %s", url))
	}

	return err
}

func (a *AssetManager) HasFileOrDir(path ...string) bool {
	if _, err := os.Stat(filepath.Join(path...)); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func (a *AssetManager) HasMap(name string) bool {
	return a.HasFileOrDir(a.MapPath(name))
}

func (a *AssetManager) MapPath(name string) string {
	return filepath.Join(a.MapsDir(), MapFilename(name))
}

func (a *AssetManager) MapsDir() string {
	return a.AbsPath(filepath.Join("qw", "maps"))
}

func MapFilename(name string) string {
	return fmt.Sprintf("%s.bsp", strings.TrimSuffix(name, ".bsp"))
}

func MapUrl(name string) string {
	return fmt.Sprintf("https://maps.quakeworld.nu/all/%s", MapFilename(name))
}
