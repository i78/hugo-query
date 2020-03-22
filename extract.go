package main

import (
	"bufio"
	"encoding/json"
	"github.com/alecthomas/kong"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type ExtractCommand struct {
	ContentFolder string   `flag name:"content-folder" help:"Your hugo content folder"  type:"path"`
	ContentType   string   `flag name:"content-type" help:"The content type to search for"`
	Field         []string `flag help:"Name of the frontmatter field to extract from"`
	Pretty        bool     `help:"Pretty-Print JSON (true/false)"`
}

var typeExtractor = regexp.MustCompile("(.*) = \"(.*?)\"")

const separator = "+++"

func ExtractFrontmatterFields(path string) (fields map[string]string) {
	fields = make(map[string]string)
	file, err := os.Open(path)
	separatorOccourcencesRemaining := 2

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() && separatorOccourcencesRemaining > 0 {
		if scanner.Text() == separator {
			separatorOccourcencesRemaining--
		}

		tokens := typeExtractor.FindStringSubmatch(scanner.Text())

		if len(tokens) > 1 {
			fields[tokens[1]] = tokens[2]
		}
	}

	return
}

func (ex *ExtractCommand) HandleExtractCommand(ctx *kong.Context) error {
	results := make(map[string]interface{})

	err := filepath.Walk(ex.ContentFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.HasSuffix(path, ".md") {

				fields := ExtractFrontmatterFields(path)

				if fields["type"] != ex.ContentType {
					return nil
				}

				if _, exists := fields["title"]; !exists {
					return nil
				}
				name := fields["title"]

				documentValues := make(map[string]interface{})

				for _, fieldName := range ex.Field {
					if fieldContent, exists := fields[fieldName]; exists {
						documentValues[fieldName] = fieldContent
					}
				}

				results[name] = documentValues
			}
			return nil
		})

	if err != nil {
		return err
	}

	if ex.Pretty {
		b, _ := json.MarshalIndent(results, "", "  ")
		ctx.Stdout.Write(b)
	} else {
		b, _ := json.Marshal(results)
		ctx.Stdout.Write(b)
	}

	return nil
}
