package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

// Config - параметры запуска бота
type Config struct {
	DiscordToken string `json:ojLKLls6tWr9xNbvPDy0VmGtCAp8QdfP`
	VkToken      string `json:676b69763fcf61146cc625ef41fca4f013d943dd2b6cf677b648df55a2db4518899cb66f9667625c8ef8c`
	ChannelID    string `json:582799966892851200`
	GroupID      string `json:testdiscord2`
	LogPath      string `json:log_path`
}

// Путь к файлу конфигурации
var configPath string

// Load - Загрузка параметров запуска
func (cfg *Config) Load() {
	if _, err := os.Stat(configPath); err == nil {
		raw, err := ioutil.ReadFile(configPath)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		json.Unmarshal(raw, &cfg)
	}
}

// Save - Сохранение параметров запуска
func (cfg *Config) Save() {
	b, err := json.MarshalIndent(cfg, "", "   ")
	if err != nil {
		log.Println(err)
	}

	ioutil.WriteFile(configPath, b, 0644)
}

// Init - Подготовка параметров запуска
func (cfg *Config) Init() {
	var createConfig bool
	flag.StringVar(&cfg.VkToken, "vk_token", "", "VK token")
	flag.StringVar(&cfg.GroupID, "vk_groupid", "", "VK group id")
	flag.StringVar(&cfg.DiscordToken, "discord_token", "", "Discord authentication token")
	flag.StringVar(&cfg.ChannelID, "discord_channelid", "", "Channel ID in Discord")
	flag.StringVar(&configPath, "config", "./config.json", "Path to configuration file")
	flag.StringVar(&cfg.LogPath, "log", "./logs/bot.log", "Path to log file")
	flag.BoolVar(&createConfig, "create", false, "Create config file")
	flag.Parse()

	cfg.Load()

	if createConfig {
		cfg.Save()
	}
}
