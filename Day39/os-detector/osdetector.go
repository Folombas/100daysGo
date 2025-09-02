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

// getWindowsVersion возвращает точную версию Windows
func getWindowsVersion() string {
	// Для Windows используем более точное определение версии
	version := runtime.GOOS
	if os := runtime.GOOS; os == "windows" {
		// Получаем более детальную информацию о версии Windows
		version = getWindowsDetails()
	}
	return version
}

// getLinuxDistroWithIcon возвращает информацию о дистрибутиве Linux с иконкой
func getLinuxDistroWithIcon() string {
	distro := getLinuxDistro()
	
	// Добавляем иконки для популярных дистрибутивов
	if strings.Contains(strings.ToLower(distro), "ubuntu") {
		return "🐧 " + distro
	} else if strings.Contains(strings.ToLower(distro), "debian") {
		return "🌀 " + distro
	} else if strings.Contains(strings.ToLower(distro), "centos") {
		return "🔴 " + distro
	} else if strings.Contains(strings.ToLower(distro), "fedora") {
		return "🔵 " + distro
	} else if strings.Contains(strings.ToLower(distro), "arch") {
		return "💠 " + distro
	} else if strings.Contains(strings.ToLower(distro), "mint") {
		return "🍃 " + distro
	} else if strings.Contains(strings.ToLower(distro), "kali") {
		return "🐉 " + distro
	}
	
	return "🐧 " + distro // Иконка пингвина по умолчанию для Linux
}

// getLinuxDistro возвращает информацию о дистрибутиве Linux
func getLinuxDistro() string {
	// Пытаемся прочитать информацию о дистрибутиве
	content, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		return "Linux (неизвестный дистрибутив)"
	}

	// Парсим содержимое файла
	lines := strings.Split(string(content), "\n")
	var name, version string

	for _, line := range lines {
		if strings.HasPrefix(line, "NAME=") {
			name = strings.Trim(strings.TrimPrefix(line, "NAME="), "\"")
		}
		if strings.HasPrefix(line, "VERSION_ID=") {
			version = strings.Trim(strings.TrimPrefix(line, "VERSION_ID="), "\"")
		}
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			// Если есть PRETTY_NAME, используем его
			prettyName := strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), "\"")
			return prettyName
		}
	}

	if name != "" && version != "" {
		return fmt.Sprintf("%s %s", name, version)
	} else if name != "" {
		return fmt.Sprintf("%s", name)
	}

	return "Linux (неизвестный дистрибутив)"
}