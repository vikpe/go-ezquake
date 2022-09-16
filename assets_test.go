package ezquake_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vikpe/go-ezquake"
)

func GetAssetManager() *ezquake.AssetManager {
	quakeDirAbsPath, _ := filepath.Abs("./test_files/quake")
	return ezquake.NewAssetManager(quakeDirAbsPath)
}

func TestAssetManager_HasMap(t *testing.T) {
	manager := GetAssetManager()
	assert.True(t, manager.HasMap("foo"))
	assert.False(t, manager.HasMap("dm6"))
}

func TestAssetManager_MapPath(t *testing.T) {
	manager := GetAssetManager()
	expect, _ := filepath.Abs("./test_files/quake/qw/maps/dm6.bsp")
	assert.Equal(t, expect, manager.MapPath("dm6"))
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

func TestAssetManager_MapsDir(t *testing.T) {
	manager := GetAssetManager()
	expect, _ := filepath.Abs("./test_files/quake/qw/maps")
	assert.Equal(t, expect, manager.MapsDir())
}
