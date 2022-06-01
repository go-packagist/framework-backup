package filesystem

import (
	"fmt"
	"time"
)

type Manager struct {
	config *Config
	disks  map[string]Drive
}

type Config struct {
	Default string
	Disk    map[string]interface{}
}

type Drive interface {
	Exists(path string) bool
	Get(path string) (string, error)
	Put(path, contents string) error
	Prepend(path, contents string) error
	Append(path, contents string) error
	Delete(path string) error
	Copy(from, to string) error
	Move(from, to string) error
	Rename(from, to string) error
	Size(path string) (int64, error)
	LastModified(path string) (time.Time, error)
	Files(directory string) ([]string, error)
	AllFiles(directory string) ([]string, error)
}

// NewManager creates a new filesystem manager
//
// 	fs := filesystem.NewManager(&filesystem.Config{
// 		Default: "local",
// 		Disk: map[string]interface{}{
// 			"local": &filesystem.LocalDriveConfig{
// 				Root: "temp",
// 			},
// 		},
// 	})
//
// 	fmt.Println(fs.Disk("local").Exists("test.txt"))
//
func NewManager(config *Config) *Manager {
	return &Manager{
		config: config,
	}
}

// Drive returns a filesystem drive; alias Disk
// fs.Drive("local")
func (m *Manager) Drive(name string) Drive {
	return m.Disk(name)
}

// Disk returns a filesystem drive
// fs.Disk("local")
func (m *Manager) Disk(name string) Drive {
	if disk, ok := m.disks[name]; ok {
		return disk
	}

	return m.resolve(name)
}

// resolve returns a filesystem drive
func (m *Manager) resolve(name string) Drive {
	config, err := m.getConfig(name)

	if err != nil {
		panic(err)
	}

	switch c := config.(type) {
	case *LocalDriveConfig:
		return m.CreateLocalDrive(c)
	default:
		panic(fmt.Sprintf("Unknown drive type: %s", name))
	}
}

// getConfig returns a filesystem drive configuration
func (m *Manager) getConfig(name string) (interface{}, error) {
	if config, ok := m.config.Disk[name]; ok {
		return config, nil
	}

	return nil, fmt.Errorf("config [%s] not found", name)
}

// CreateLocalDrive creates a local filesystem drive
func (m *Manager) CreateLocalDrive(config *LocalDriveConfig) Drive {
	return NewLocalDrive(config)
}
