package element

import (
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Element holds the build configuration of rlxos package
type Element struct {
	Id      string `yaml:"id"`
	Version string `yaml:"version"`
	About   string `yaml:"about"`
	Release int    `yaml:"release"`

	Variables map[string]string `yaml:"variables"`

	// Depends dependencies are required both during buildtime and runtime
	Depends []string `yaml:"depends"`

	// BuildTime dependencies are required during buildtime only
	BuildTime []string `yaml:"build-time"`

	// Runtime dependencies are those needed only during runtime
	Runtime []string `yaml:"run-time"`

	BuildDir string `yaml:"build-dir"`

	BuildType string `yaml:"build-type"`

	Config struct {
		Source string `yaml:"source"`
		Target string `yaml:"target"`
	} `yaml:"config"`

	Sources []string `yaml:"sources"`
	Environ []string `yaml:"environ"`
	Include []string `yaml:"include"`

	PreScript  string `yaml:"pre-script"`
	Script     string `yaml:"script"`
	PostScript string `yaml:"post-script"`

	Configure string `yaml:"configure"`
	Compile   string `yaml:"compile"`
	Install   string `yaml:"install"`

	Integration string `yaml:"integration"`

	Split []ElementSplit `yaml:"split"`

	NoStrip   bool     `yaml:"no-strip"`
	SkipStrip []string `yaml:"skip-strip"`
}

// ElementSplit holds the information of sub package that can be
// seperated from rlxos Element
type ElementSplit struct {
	// Into defines the suffix name of package
	// for example Element 'gcc' contains sub package 'gcc:lib'
	Into string `yaml:"into"`

	// About provide a basic description about the sub package
	About string `yaml:"about"`

	// Files holds the list of files that need to be seperated from parent Element
	Files []string `yaml:"files"`
}

// Open open the rlxos package element file
func Open(filepath string, environ []string, variables map[string]string) (*Element, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var e Element
	if err := yaml.Unmarshal(data, &e); err != nil {
		return nil, err
	}
	if e.Variables == nil {
		e.Variables = map[string]string{}
	}
	e.Variables["id"] = e.Id
	e.Variables["version"] = e.Version
	// e.Variables["release"] = fmt.Sprint(e.Release)

	if environ != nil {
		e.Environ = append(e.Environ, environ...)
	}

	for key, value := range variables {
		e.Variables[key] = value
	}

	for i := range e.Sources {
		e.Sources[i] = e.resolveVariable(e.Sources[i])
	}

	for i := range e.Environ {
		e.Environ[i] = e.resolveVariable(e.Environ[i])
	}

	e.BuildDir = e.resolveVariable(e.BuildDir)

	e.PreScript = e.resolveVariable(e.PreScript)
	e.Script = e.resolveVariable(e.Script)
	e.PostScript = e.resolveVariable(e.PostScript)

	e.Configure = e.resolveVariable(e.Configure)
	e.Compile = e.resolveVariable(e.Compile)
	e.Install = e.resolveVariable(e.Install)

	return &e, nil
}

func (e *Element) resolveVariable(v string) string {
	for key, value := range e.Variables {
		v = strings.ReplaceAll(v, "%{"+key+"}", value)
	}
	return v
}

type DependencyType int

const (
	DependencyBuildTime DependencyType = iota
	DependencyRunTime
	DependencyAll
)

func (e *Element) AllDepends(dep DependencyType) []string {
	depends := []string{}
	if e.Depends != nil {
		depends = append(depends, e.Depends...)
	}

	if dep == DependencyBuildTime || dep == DependencyAll {
		if e.BuildTime != nil {
			depends = append(depends, e.BuildTime...)
		}
	}

	if dep == DependencyRunTime || dep == DependencyAll {
		if e.Runtime != nil {
			depends = append(depends, e.Runtime...)
		}
	}

	return depends
}
