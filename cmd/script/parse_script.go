package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func extractRaycastMetadatas(content string) map[string]string {
	r := regexp.MustCompile("@raycast.([A-Za-z0-9]+)\\s([\\S ]+)")
	groups := r.FindAllStringSubmatch(content, -1)

	metadataMap := make(map[string]string)
	for _, group := range groups {
		metadataMap[group[0]] = group[1]
	}

	return metadataMap
}

func main() {
	content, _ := ioutil.ReadFile("/workspace/raycast-linux/scripts/test.sh")

	metadatas := extractRaycastMetadatas(string(content))

	fmt.Println(metadatas)
}
