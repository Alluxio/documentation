/*
 * The Alluxio Open Foundation licenses this work under the Apache License, version 2.0
 * (the "License"). You may not use this work except in compliance with the License, which is
 * available at www.apache.org/licenses/LICENSE-2.0
 *
 * This software is distributed on an "AS IS" basis, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied, as more fully set forth in the License.
 *
 * See the NOTICE file distributed with this work for information regarding copyright ownership.
 */

package main

import (
	"fmt"
	"github.com/palantir/stacktrace"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

const (
	cnPath    = "cn"
	menuEnYml = "_data/menu-en.yml"
	menuCnYml = "_data/menu-cn.yml"
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatalln(err)
	}
	log.Println(".gitbook.yaml and Summary.md created successfully")
}

func run(args []string) error {
	if len(args) < 2 {
		return stacktrace.NewError("expected run this scripts with the docs folder. For example generate-config.sh docs-edge/en ")
	}
	docsRootPath := args[1]
	repoRoot, err := FindRepoRoot()
	if err != nil {
		return stacktrace.Propagate(err, "error finding repo root")
	}
	docsAbsPath := filepath.Join(repoRoot, docsRootPath)
	/*
		if err := generateGitbookYaml(docsRootPath, repoRoot); err != nil {
			return stacktrace.Propagate(err, "error creating .gitbook.yaml")
		}
	*/
	if err := convertRelativeLink(docsAbsPath); err != nil {
		return stacktrace.Propagate(err, "error converting relative link")
	}
	// summary
	/*
		menuYml := menuEnYml
		if strings.Contains(docsRootPath, cnPath) {
			menuYml = menuCnYml
		}
		if err := convertMenuToSummary(docsAbsPath, menuYml); err != nil {
			return nil
		}
	*/
	return nil
}
func FindRepoRoot() (string, error) {
	// navigate 5 parent directories to reach repo root,
	// assuming this go file is located in <repoRoot>/scripts/src/alluxio.org/generate_docs/main.go
	const repoRootDepth = 5
	_, repoRoot, _, ok := runtime.Caller(0)
	if !ok {
		return "", stacktrace.NewError("error getting call stack")
	}
	for i := 0; i < repoRootDepth; i++ {
		repoRoot = filepath.Dir(repoRoot)
	}
	log.Printf("Repository root at directory: %v", repoRoot)
	return repoRoot, nil
}

type Structure struct {
	Readme  string `yaml:"readme"`
	Summary string `yaml:"summary"`
}

type Config struct {
	Root      string    `yaml:"root"`
	Structure Structure `yaml:"structure"`
}

func generateGitbookYaml(docsRootPath, repoRoot string) error {
	config := Config{
		Root: fmt.Sprintf("./%v/", docsRootPath),
		Structure: Structure{
			Readme:  "./overview/Overview.md",
			Summary: "./Summary.md",
		},
	}
	fmt.Println(config.Root)
	if strings.Contains(docsRootPath, "ai") {
		config.Structure.Readme = "./Overview.md"
	}
	data, err := yaml.Marshal(&config)
	if err != nil {
		return stacktrace.Propagate(err, "error converting .gitbook.yaml content: %v", err)
	}
	file, err := os.Create(filepath.Join(repoRoot, ".gitbook.yaml"))
	if err != nil {
		return stacktrace.Propagate(err, "error creating .gitbook.yaml: %v", err)
	}
	defer file.Close()
	if _, err := file.Write(data); err != nil {
		return stacktrace.Propagate(err, "error writing .gitbook.yaml: %v", err)
	}
	return nil
}

type MenuFile struct {
	Title       string     `yaml:"title"`
	URL         string     `yaml:"url"`
	ButtonTitle string     `yaml:"buttonTitle"`
	Subfiles    []MenuFile `yaml:"subfiles"`
}

func convertMenuToSummary(docsAbsPath, menuYml string) error {
	menuPath := filepath.Join(docsAbsPath, "..", menuYml)
	if _, err := os.Stat(menuPath); os.IsNotExist(err) {
		return stacktrace.Propagate(err, "error finding file %v", menuPath)
	}
	yamlFile, err := ioutil.ReadFile(menuPath)
	if err != nil {
		return stacktrace.Propagate(err, "Error reading YAML file: %v", menuPath)
	}
	var menuContent []MenuFile
	if err := yaml.Unmarshal(yamlFile, &menuContent); err != nil {
		return stacktrace.Propagate(err, "error converting yaml to struct for %v", menuPath)
	}
	summaryContent := "# Table of contents\n\n"
	for _, content := range menuContent {
		summaryContent += fmt.Sprintf("* %s\n", content.Title)
		for _, subfile := range content.Subfiles {
			summaryContent += fmt.Sprintf("\t* [%s](%s)\n", subfile.Title, subfile.URL[len("/en/"):len(subfile.URL)-len(".html")]+".md")
		}
	}
	file, err := os.Create(filepath.Join(docsAbsPath, "Summary.md"))
	if err != nil {
		return stacktrace.Propagate(err, "error creating Summary.md")
	}
	defer file.Close()

	_, err = file.WriteString(summaryContent)
	if err != nil {
		return stacktrace.Propagate(err, "Error writing to Summary.md")
	}
	return nil
}
func convertRelativeLink(docsAbsPath string) error {
	trimmedPath := strings.TrimRight(docsAbsPath, "/")
	lastSlashIndex := strings.LastIndex(trimmedPath, "/")
	docsPathWithoutanguage := trimmedPath[:lastSlashIndex+1]
	// pattern to match: [anystring]({{ 'xxx/xxxx/zzz.html' | relativize_url }})
	relativeHtmlPattern := regexp.MustCompile(`\[[^\]]+\]\(\{\{ *'([^']+\.html)' *\| *relativize_url *\}\}\)`)
	link2Pattern := regexp.MustCompile(`\[[^\]]+\]\(\{\{ *'([^']+\.html(?:#[^']*)?)' *\| *relativize_url *\}\}\)`)
	// pattern to match: ---
	//					 title :
	//					 ---
	headerPattern := regexp.MustCompile(`(?s)^---.*?title:\s*(.*?)\s*---\s*`)
	// pattern to match: <img anything {{ 'xxxx.png' | relativize_url }}" anything/>
	img1Pattern := regexp.MustCompile(`<img[^>]*\{\{\s*'([^']+\.png)'\s*\|\s*relativize_url\s*\}\}[^>]*\/>`)
	// pattern to match: ![Alluxio Namespace]({{ '/img/alluxio_namespace.png' | relativize_url }})
	img2Pattern := regexp.MustCompile(`\!\[[^\]]+\]\(\{\{ *'([^']+\.png)' *\| *relativize_url *\}\}\)`)

	err := filepath.Walk(docsAbsPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return stacktrace.Propagate(err, "error")
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			fmt.Println("Processing file:", path)
			if err := processFile(path, docsPathWithoutanguage, relativeHtmlPattern, link2Pattern, headerPattern, img1Pattern, img2Pattern); err != nil {
				return stacktrace.Propagate(err, "error converting relative link")
			}

		}
		return nil
	})
	if err != nil {
		return stacktrace.Propagate(err, "error walking through directory")
	}
	return nil
}

func processFile(filePath, docsPathWithoutLanguage string, linkPattern, link2Pattern, headerPattern, img1Pattern, img2Pattern *regexp.Regexp) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return stacktrace.Propagate(err, "Error reading file %s: %v", filePath, err)
	}
	contentStr := string(content)
	// convert link
	modifiedContent := linkPattern.ReplaceAllStringFunc(contentStr, func(match string) string {
		submatches := linkPattern.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match
		}
		extractedPath := submatches[1]
		fullPath := filepath.Join(docsPathWithoutLanguage, extractedPath)
		var newPath string
		if fileExists(fullPath) {
			newPath = extractedPath
		} else {
			if strings.Contains(extractedPath, "/cn/") {
				newPath = strings.Replace(extractedPath, "/cn/", "../", 1)
			} else {
				newPath = strings.Replace(extractedPath, "/en/", "../", 1)
			}
			newPath = strings.Replace(newPath, ".html", ".md", 1)
		}
		start := strings.Index(match, "[")
		end := strings.Index(match, "]")
		name := ""
		if start != -1 && end != -1 && start < end {
			name = match[start+1 : end]
		}
		return fmt.Sprintf("[%s](%s)", name, newPath)
	})
	modifiedContent = link2Pattern.ReplaceAllStringFunc(contentStr, func(match string) string {
		submatches := link2Pattern.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match
		}
		extractedPath := submatches[1]
		fullPath := filepath.Join(docsPathWithoutLanguage, extractedPath)
		var newPath string
		if fileExists(fullPath) {
			newPath = extractedPath
		} else {
			if strings.Contains(extractedPath, "/cn/") {
				newPath = strings.Replace(extractedPath, "/cn/", "../", 1)
			} else {
				newPath = strings.Replace(extractedPath, "/en/", "../", 1)
			}
			newPath = strings.Replace(newPath, ".html", ".md", 1)
		}
		start := strings.Index(match, "[")
		end := strings.Index(match, "]")
		name := ""
		if start != -1 && end != -1 && start < end {
			name = match[start+1 : end]
		}
		return fmt.Sprintf("[%s](%s)", name, newPath)
	})
	// convert header
	match := headerPattern.FindStringSubmatch(modifiedContent)
	if len(match) > 1 {
		// Prepend the title as a Markdown heading
		modifiedContent = "# " + match[1] + "\n\n" + headerPattern.ReplaceAllString(modifiedContent, "")
	}
	// convert img
	modifiedContent = img1Pattern.ReplaceAllStringFunc(modifiedContent, func(match string) string {
		submatches := img1Pattern.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match
		}
		imgPath := filepath.Join("../.gitbook/assets", filepath.Base(submatches[1]))
		return fmt.Sprintf(`<figure><img src="%s" alt=""><figcaption></figcaption></figure>`, imgPath)
	})
	modifiedContent = img2Pattern.ReplaceAllStringFunc(modifiedContent, func(match string) string {
		submatches := img2Pattern.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match
		}
		imgPath := filepath.Join("../.gitbook/assets", filepath.Base(submatches[1]))
		return fmt.Sprintf(`<figure><img src="%s" alt=""><figcaption></figcaption></figure>`, imgPath)
	})
	err = ioutil.WriteFile(filePath, []byte(modifiedContent), 0644)
	if err != nil {
		return stacktrace.Propagate(err, "Error writing file %s: %v", filePath, err)
	}
	return nil
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
