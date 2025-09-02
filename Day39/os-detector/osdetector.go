package main

import (
	"fmt"
	"io/ioutil"
	"os-detector/utils"
	"runtime"
	"strings"
)

// RunOSDetector определяет и выводит информацию об ОС
func RunOSDetector() {
	osInfo := getOSInfo()
	message := fmt.Sprintf("Эта утилита, написанная на языке программирования Go, запущена на операционной системе %s", osInfo)
	utils.PrintCyrillic(message)
}

// getOSInfo возвращает информацию об операционной системе с иконкой
func getOSInfo() string {
	os := runtime.GOOS

	switch os {
	case "windows":
		return "💻 " + getWindowsVersion()
	case "linux":
		return getLinuxDistroWithIcon()
	case "darwin":
		return "🍎 macOS"
	default:
		return os
	}
}

// getWindowsVersion возвращает версию Windows
func getWindowsVersion() string {
	// Для Windows используем более точное определение версии
	if runtime.GOOS == "windows" {
		return getWindowsDetails()
	}
	return "Windows"
}

func getLinuxDistroWithIcon() string {
	distro := getLinuxDistro()

	// Добавляем иконки для популярных дистрибутивов
	distroLower := strings.ToLower(distro)
	switch {
	case strings.Contains(distroLower, "ubuntu"):
		// Определяем версию Ubuntu
		if strings.Contains(distroLower, "24.04") || strings.Contains(distroLower, "noble") {
			return "🦁 " + distro + " (Noble Numbat)"
		} else if strings.Contains(distroLower, "22.04") || strings.Contains(distroLower, "jammy") {
			return "🦁 " + distro + " (Jammy Jellyfish)"
		}
		return "🦁 " + distro
	case strings.Contains(distroLower, "debian"):
		return "🌀 " + distro
	case strings.Contains(distroLower, "centos"):
		return "🔴 " + distro
	case strings.Contains(distroLower, "fedora"):
		return "🔵 " + distro
	case strings.Contains(distroLower, "arch"):
		return "💠 " + distro
	case strings.Contains(distroLower, "mint"):
		return "🍃 " + distro
	case strings.Contains(distroLower, "kali"):
		return "🐉 " + distro
	case strings.Contains(distroLower, "alpine"):
		return "🏔️ " + distro
	case strings.Contains(distroLower, "opensuse"):
		return "🦎 " + distro
	}

	return "🐧 " + distro // Иконка пингвина по умолчанию для Linux
}

// getLinuxDistro возвращает информацию о дистрибутиве Linux
func getLinuxDistro() string {
	// Пробуем разные файлы с информацией о дистрибутиве
	files := []string{
		"/etc/os-release",
		"/usr/lib/os-release",
		"/etc/lsb-release",
		"/etc/redhat-release",
		"/etc/debian_version",
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			continue
		}

		// Анализируем содержимое файла
		lines := strings.Split(string(content), "\n")
		var name, version, prettyName, versionCodename string

		for _, line := range lines {
			if strings.HasPrefix(line, "NAME=") {
				name = strings.Trim(strings.TrimPrefix(line, "NAME="), "\"")
			}
			if strings.HasPrefix(line, "VERSION_ID=") {
				version = strings.Trim(strings.TrimPrefix(line, "VERSION_ID="), "\"")
			}
			if strings.HasPrefix(line, "PRETTY_NAME=") {
				prettyName = strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), "\"")
			}
			if strings.HasPrefix(line, "VERSION_CODENAME=") {
				versionCodename = strings.Trim(strings.TrimPrefix(line, "VERSION_CODENAME="), "\"")
			}
			if strings.HasPrefix(line, "DISTRIB_DESCRIPTION=") {
				prettyName = strings.Trim(strings.TrimPrefix(line, "DISTRIB_DESCRIPTION="), "\"")
			}
		}

		// Если нашли PRETTY_NAME, используем его
		if prettyName != "" {
			// Добавляем кодовое имя для Ubuntu
			if versionCodename != "" && strings.Contains(strings.ToLower(prettyName), "ubuntu") {
				return fmt.Sprintf("%s (%s)", prettyName, strings.Title(versionCodename))
			}
			return prettyName
		}

		// Формируем имя из компонентов
		if name != "" {
			if version != "" {
				return fmt.Sprintf("%s %s", name, version)
			}
			return name
		}

		// Пытаемся определить дистрибутив по содержимому файла
		contentStr := string(content)
		switch {
		case strings.Contains(contentStr, "Ubuntu"):
			return "Ubuntu"
		case strings.Contains(contentStr, "Debian"):
			return "Debian"
		case strings.Contains(contentStr, "CentOS"):
			return "CentOS"
		case strings.Contains(contentStr, "Fedora"):
			return "Fedora"
		case strings.Contains(contentStr, "Arch"):
			return "Arch Linux"
		}
	}

	return "Linux (неизвестный дистрибутив)"
}
