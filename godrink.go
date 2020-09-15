package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
	"time"
)

const CONFIG_NAME = ".godrink_conf"

func main() {
	// args[0] / first always hold the program name
	if len(os.Args) <= 1 {
		checkThirsty()
		return
	}

	action := os.Args[1]

	switch action {
	case "not_thirsty":
		resetThirsty()

	case "setup":
		// There is no isset().
		if len(os.Args) >= 3 {
			interval := os.Args[2]
			success := setup(interval)

			if success == true {
				fmt.Println("Setup successfull")
			} else {
				fmt.Println("Something went wrong, please try again later.")
			}

			return
		}

		// Stop execution and display an error.
		log.Fatal("Mising interval argument")
	}
}

// Create a .conf file and save the interval string.
func setup(interval string) bool {
	var file *os.File
	var err error

	user, err := user.Current()
	config_path := user.HomeDir + "/" + CONFIG_NAME

	if _, err := os.Stat(config_path); err == nil {
		file, err = os.OpenFile(config_path, os.O_TRUNC|os.O_WRONLY, 0644)

	} else if os.IsNotExist(err) {
		file, err = os.Create(config_path)
	}

	if err != nil {
		log.Fatal("Can't create/open the config file.")
	}

	file.WriteString(fmt.Sprintf("%s", interval))
	err = file.Close()

	if err != nil {
		log.Fatal("Can't write to file.")
	}

	return true
}

// Get user directory.
func get_user_dir() string {
	user, err := user.Current()

	if err != nil {
		log.Fatal("Can't locate user directory")
	}

	return user.HomeDir
}

// Check the timestamp and interval.
func checkThirsty() {
	// Get last timestamp
	var file *os.File
	var err error
	var saved_timestamp int64
	var file_contents []byte
	var interval int64

	user_dir := get_user_dir()
	config := user_dir + "/.godrink"

	if _, err := os.Stat(config); err == nil {
		file_contents, err = ioutil.ReadFile(config)

	} else if os.IsNotExist(err) {
		file, err = os.Create(config)
		file.WriteString(fmt.Sprintf("%d", time.Now().Unix()))
		return
	}

	if err != nil {
		return
	}

	// calculate time diff
	saved_timestamp, err = strconv.ParseInt(string(file_contents), 10, 64)
	current_timestamp := time.Now().Unix()
	timestamp_diff := current_timestamp - saved_timestamp

	// Get saved interval
	config_path := user_dir + "/" + CONFIG_NAME
	interval_contents, err := ioutil.ReadFile(config_path)
	interval_string := string(interval_contents)
	interval, err = strconv.ParseInt(interval_string, 10, 64)

	if timestamp_diff > interval {
		fmt.Println("💧You're thirsty!")
		return
	}

	return
}

// Reset the timestamp
func resetThirsty() {
	var file *os.File
	var err error

	config := get_user_dir() + "/.godrink"

	if _, err := os.Stat(config); err == nil {
		file, err = os.OpenFile(config, os.O_TRUNC|os.O_WRONLY, 0644)
		file.WriteString(fmt.Sprintf("%d", time.Now().Unix()))
		return
	}

	if err != nil {
		log.Fatal("Can't reset thirsty, please try again later.")
	}

	return
}
