package filesystem

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestManager_Local(t *testing.T) {
	// init
	m := NewManager(&Config{
		Default: "local",
		Disk: map[string]interface{}{
			"local": &LocalDriveConfig{
				Root: "temp",
			},
		},
	})

	// Exists
	assert.True(t, m.Disk("local").Exists("base.txt"))
	assert.False(t, m.Disk("local").Exists("none.txt"))

	// Get
	baseContent, _ := m.Disk("local").Get("base.txt")
	assert.Equal(t, "base", baseContent)

	// Put
	writeBody := time.Now().String()
	m.Disk("local").Put("write.txt", writeBody)
	writeContent, _ := m.Disk("local").Get("write.txt")
	assert.Equal(t, writeBody, writeContent)

	// Put None File
	m.Disk("local").Put("none-file.txt", "none")
	noneContent, _ := m.Disk("local").Get("none-file.txt")
	assert.Equal(t, "none", noneContent)

	// Put None Path
	m.Disk("local").Put("none-path/none-file.txt", "none")
	noneContent2, noneErr2 := m.Disk("local").Get("none-path/none-file.txt")
	assert.Equal(t, "", noneContent2) // todo: auto craete path
	assert.Error(t, noneErr2)

	// Size
	size, _ := m.Disk("local").Size("base.txt")
	assert.Equal(t, int64(4), size)
	size2, size2Err := m.Disk("local").Size("none.txt")
	assert.Equal(t, int64(0), size2)
	assert.Error(t, size2Err)

	// Delete
	m.Disk("local").Delete("none-file.txt")
	assert.False(t, m.Disk("local").Exists("none-file.txt"))

	// Files
	files, _ := m.Disk("local").Files("files")
	assert.Equal(t, []string{"files/.txt", "files/1.txt", "files/2.txt"}, files)

	// AllFiles
	allFiles, _ := m.Disk("local").AllFiles("files")
	assert.Equal(t, []string{"files/.txt", "files/1.txt", "files/2.txt", "files/dir1/.txt", "files/dir1/1.txt"}, allFiles)
}

func TestManager_S3(t *testing.T) {
	// init
	// m := NewManager(&Config{
	// 	Default: "s3",
	// 	Disk: map[string]interface{}{
	// 		"s3": &S3DriveConfig{
	// 			AccessKey: "",
	// 			SecretKey: "",
	// 			Bucket:    "",
	// 			Region:    "",
	// 		},
	// 	},
	// })
}
