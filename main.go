package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	modelTags := []string{}
	modelFile, _ := os.Open("zek/model.txt")
	defer modelFile.Close()
	scanner := bufio.NewScanner(modelFile)
	for scanner.Scan() {
		modelTags = append(modelTags, strings.TrimSpace(scanner.Text()))
	}

	tagFile, _ := os.Open("zek/tags.txt")
	defer tagFile.Close()
	scanner = bufio.NewScanner(tagFile)
	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		if contains(modelTags, t) == false {
			fmt.Println(t)
		}
	}
}

func gen() {
	obj := regexp.MustCompile("{")
	f, _ := os.Open("zek/sample-set-2-ead-files.txt")
	tags := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		if obj.MatchString(txt) == true {
			cleaned := strings.ReplaceAll(txt, " []struct {", "")
			cleaned = strings.ReplaceAll(cleaned, "struct {", "")
			tag := strings.TrimSpace(cleaned)
			if contains(tags, tag) != true {
				tags = append(tags, tag)
			}
			sort.Strings(tags)
		}
	}

	outfile, _ := os.Create("zek/tags.txt")
	defer outfile.Close()
	writer := bufio.NewWriter(outfile)
	for _, tag := range tags {
		writer.WriteString(strings.ToLower(tag) + "\n")
		writer.Flush()
	}
}

func contains(tags []string, s string) bool {
	for _,tag := range tags {
		if s == tag {
			return true
		}
	}
	return false
}
