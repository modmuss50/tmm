package main

import (
	"fmt"
	"github.com/modmuss50/CAV2"
	"io/ioutil"
	"strconv"
	"strings"
)

func update() {
	fmt.Println("Searching for mod updates")

	files, err := ioutil.ReadDir(runDir)
	if err != nil {
		fmt.Println("Failed to find modHashes")
		panic(err)
	}

	//Map of files to hashes
	modHashes := make(map[string]int)

	for _, f := range files {
		//Only check jar files, we might want to look into other files in the future
		if !strings.HasSuffix(f.Name(), ".jar") {
			continue
		}
		hash, err := cav2.GetFileHash(f.Name())
		if err != nil {
			fmt.Println("Failed to get hash for " + f.Name())
			panic(err)
		}
		modHashes[f.Name()] = hash
	}

	hashes := make([]int, 0)
	for _, h := range modHashes {
		hashes = append(hashes, h)
	}
	matches, err := cav2.GetHashMatches(hashes)
	if err != nil {
		fmt.Println("Failed to find modHashes on curse")
		panic(err)
	}

	fmt.Println("Found " + strconv.Itoa(len(matches.ExactMatches)) + " mods")

	//Map of all the files to curse matches
	mods := make(map[string]cav2.FingerprintMatch)
	for _, match := range matches.ExactMatches {
		for modFile, hash := range modHashes {
			if int64(hash) == match.File.PackageFingerprint {
				mods[modFile] = match
			}
		}
	}

	fmt.Println("Searching for updates")

	//Load all the addon data for the installed mods
	addonIDs := make([]int, 0)
	for _, match := range mods {
		addonIDs = append(addonIDs, match.AddonID)
	}
	addons, err := cav2.GetAddons(addonIDs)
	if err != nil {
		fmt.Println("Failed to load addon data")
		panic(err)
	}

	//Map the files to addons
	addonMap := make(map[string]cav2.Addon)
	for _, addon := range addons {
		for modFile, match := range mods {
			if addon.ID == match.AddonID {
				addonMap[modFile] = addon
			}
		}
	}

	//Find all updates
	updates := make(map[string]cav2.AddonLatestFile)
	for mod, addon := range addonMap {
		for _, latestFile := range addon.LatestFiles {
			//TODO check the puesdo version (alpha, beta, release)
			//TODO check the versions in a better way here
			if latestFile.GameVersion[0] == mods[mod].File.GameVersion[0] {
				if modHashes[mod] != latestFile.PackageFingerprint {
					updates[mod] = latestFile
					break
				}
			}
		}
	}

	fmt.Println("Found " + strconv.Itoa(len(updates)) + " updates:")
	for mod, update := range updates {
		handleUpdate(mod, update)
	}

}

func handleUpdate(file string, newFile cav2.AddonLatestFile) {
	fmt.Println(file + " -> " + newFile.FileNameOnDisk)
	//TODO write this
}

func doUpdate(file string, newFile cav2.AddonLatestFile) {
	//TODO actually update the file
}
