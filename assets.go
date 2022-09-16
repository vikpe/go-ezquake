package ezquake

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type AssetManager struct {
	baseDirPath string
}

func NewAssetManager(baseDirpath string) *AssetManager {
	return &AssetManager{
		baseDirPath: baseDirpath,
	}
}

func (a *AssetManager) AbsPath(path ...string) string {
	paths := append([]string{a.BaseDir()}, path...)
	return filepath.Join(paths...)
}

func (a *AssetManager) BaseDir() string {
	return a.baseDirPath
}

func (a *AssetManager) DemosDir() string {
	return a.AbsPath(filepath.Join("qw", "demos"))
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
	return filepath.Join(a.MapsDir(), fmt.Sprintf("%s.bsp", name))
}

func (a *AssetManager) MapsDir() string {
	return a.AbsPath(filepath.Join("qw", "maps"))
}
