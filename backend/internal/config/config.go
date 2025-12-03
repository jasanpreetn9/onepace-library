package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type MetadataConfig struct {
	EpisodesURL string `yaml:"episodesUrl"`
	ArcsURL     string `yaml:"arcsUrl"`
}

type Config struct {
	Port            string `yaml:"port"`
	LibraryPath     string `yaml:"libraryPath"`
	DownloadPath    string `yaml:"downloadPath"`
	LibraryJSONPath string `yaml:"libraryJsonPath"`

	Metadata MetadataConfig `yaml:"metadata"`

	QBittorrent QBittorrentConfig `yaml:"qbittorrent"`
}

type QBittorrentConfig struct {
	Enabled  bool   `yaml:"enabled"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func Load(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	cfg := &Config{}
	if err := yaml.Unmarshal(b, cfg); err != nil {
		return nil, fmt.Errorf("yaml unmarshal: %w", err)
	}

	applyDefaults(cfg)
	applyEnvOverrides(cfg)

	if err := validate(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func applyDefaults(cfg *Config) {
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	if cfg.LibraryPath == "" {
		cfg.LibraryPath = "./media"
	}
	if cfg.DownloadPath == "" {
		cfg.DownloadPath = "./downloads"
	}
	if cfg.LibraryJSONPath == "" {
		cfg.LibraryJSONPath = "./data/library.json"
	}
}

func applyEnvOverrides(cfg *Config) {
	if port := os.Getenv("OP_PORT"); port != "" {
		cfg.Port = port
	}
	if path := os.Getenv("OP_LIBRARY"); path != "" {
		cfg.LibraryPath = path
	}
	if path := os.Getenv("OP_DOWNLOADS"); path != "" {
		cfg.DownloadPath = path
	}
	if url := os.Getenv("OP_METADATA_EPISODES_URL"); url != "" {
		cfg.Metadata.EpisodesURL = url
	}
	if url := os.Getenv("OP_METADATA_ARCS_URL"); url != "" {
		cfg.Metadata.ArcsURL = url
	}

	if host := os.Getenv("OP_QB_HOST"); host != "" {
		cfg.QBittorrent.Host = host
	}
	if user := os.Getenv("OP_QB_USER"); user != "" {
		cfg.QBittorrent.Username = user
	}
	if pass := os.Getenv("OP_QB_PASS"); pass != "" {
		cfg.QBittorrent.Password = pass
	}
}

func validate(cfg *Config) error {
	if cfg.Metadata.EpisodesURL == "" {
		return errors.New("metadataUrl must be set in config.yml")
	}

	return nil
}
