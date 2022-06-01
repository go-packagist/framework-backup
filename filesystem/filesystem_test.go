package filesystem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileSystem_Exists(t *testing.T) {
	assert.True(t, Exists("./temp/base.txt"))
	assert.False(t, Exists("./temp/none.txt"))
}

func TestFileSystem_Stat(t *testing.T) {
	info, _ := Info("./temp/base.txt")
	fileinfo, _ := FileInfo("./temp/base.txt")
	stat, _ := Stat("./temp/base.txt")

	assert.Equal(t, info, fileinfo)
	assert.Equal(t, info, stat)

	assert.Equal(t, info.Name(), "base.txt")
	assert.Equal(t, info.Size(), int64(4))
}

func TestFileSystem_Size(t *testing.T) {
	size, _ := Size("./temp/base.txt")

	assert.Equal(t, int64(4), size)
}

func TestFileSystem_IsDirOrFile(t *testing.T) {
	isDir1 := IsDir("./temp/base.txt")
	isDir2 := IsDir("./temp/")

	assert.False(t, isDir1)
	assert.True(t, isDir2)

	isFile1 := IsFile("./temp/base.txt")
	isFile2 := IsFile("./temp/")

	assert.True(t, isFile1)
	assert.False(t, isFile2)
}

func TestFileSystem_ReadDir(t *testing.T) {
	files, _ := ReadDir("./temp/")

	assert.Equal(t, 3, len(files))
	assert.Equal(t, "base.txt", files[0].Name())
}

func TestFileSystem_ReadFile(t *testing.T) {
	content1, err1 := ReadFile("./temp/base.txt")
	content2, err2 := ReadFile("./temp/none.txt")

	assert.Equal(t, "base", content1)
	assert.Equal(t, "", content2)
	assert.Equal(t, nil, err1)
	assert.Error(t, err2)
}

func TestFileSystem_WriteFile(t *testing.T) {
	body := "tests1"

	err := WriteFile("./temp/write.txt", body)
	assert.Equal(t, nil, err)

	content, _ := ReadFile("./temp/write.txt")
	assert.Equal(t, body, content)
}
