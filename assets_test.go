package ezquake_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vikpe/go-ezquake"
)

func GetAssetManager() *ezquake.AssetManager {
	quakeDirAbsPath, _ := filepath.Abs("./test_files/quake")
	return ezquake.NewAssetManager(quakeDirAbsPath)
}

func TestAssetManager_AbsPath(t *testing.T) {
	manager := GetAssetManager()
	expect, _ := filepath.Abs("./test_files/quake/qw/textures/foo.png")
	assert.Equal(t, expect, manager.AbsPath("qw/textures/foo.png"))
}

func TestAssetManager_BaseDir(t *testing.T) {
	manager := GetAssetManager()
	expect, _ := filepath.Abs("./test_files/quake")
	assert.Equal(t, expect, manager.BaseDir())
}

func TestAssetManager_DemosDir(t *testing.T) {
	manager := GetAssetManager()
	expect, _ := filepath.Abs("./test_files/quake/qw/demos")
	assert.Equal(t, expect, manager.DemosDir())
}

func TestAssetManager_DownloadMap(t *testing.T) {
	manager := GetAssetManager()

	t.Run("map not found on maps.quakeworld.nu", func(t *testing.T) {
		err := manager.DownloadMap("foobar")
		mapAbsPath, _ := filepath.Abs("./test_files/quake/qw/maps/foobar.bsp")
		assert.NoFileExists(t, mapAbsPath)
		assert.ErrorContains(t, err, "NOT FOUND: https://maps.quakeworld.nu/all/foobar.bsp")
	})

	t.Run("map already exist in qw/maps", func(t *testing.T) {
		err := manager.DownloadMap("dm2")
		mapAbsPath, _ := filepath.Abs("./test_files/quake/qw/maps/dm2.bsp")
		assert.FileExists(t, mapAbsPath)
		assert.Nil(t, err)
	})

	t.Run("new map", func(t *testing.T) {
		err := manager.DownloadMap("fragtwn1")
		mapAbsPath, _ := filepath.Abs("./test_files/quake/qw/maps/fragtwn1.bsp")
		assert.FileExists(t, mapAbsPath)
		assert.Nil(t, err)

		os.Remove(manager.MapPath("fragtwn1"))
	})
}

func TestAssetManager_HasMap(t *testing.T) {
	manager := GetAssetManager()
	assert.True(t, manager.HasMap("dm2"))
	assert.True(t, manager.HasMap("dm2.bsp"))
	assert.False(t, manager.HasMap("dm6"))
}

func TestAssetManager_MapPath(t *testing.T) {
	manager := GetAssetManager()
	expect, _ := filepath.Abs("./test_files/quake/qw/maps/dm6.bsp")
	assert.Equal(t, expect, manager.MapPath("dm6"))
}

func TestAssetManager_MapsDir(t *testing.T) {
	manager := GetAssetManager()
	expect, _ := filepath.Abs("./test_files/quake/qw/maps")
	assert.Equal(t, expect, manager.MapsDir())
}

func TestMapFilename(t *testing.T) {
	assert.Equal(t, "dm6.bsp", ezquake.MapFilename("dm6"))
	assert.Equal(t, "dm6.bsp", ezquake.MapFilename("dm6.bsp"))
}

func TestMapUrl(t *testing.T) {
	assert.Equal(t, "https://maps.quakeworld.nu/all/dm6.bsp", ezquake.MapUrl("dm6"))
	assert.Equal(t, "https://maps.quakeworld.nu/all/dm6.bsp", ezquake.MapUrl("dm6.bsp"))
}
