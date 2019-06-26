package main

import (
  "fmt"
  "strings"
  "log"
  "os"
  "os/exec"
  "github.com/urfave/cli"
)

var app = cli.NewApp()

func Info() {
	app.Name = "Gradle git version calculator"
	app.Usage = "Script for generating semantic version based on the version in gradle.properties and git branch"
	app.Version = "0.0.1"
}

func Commands() {
  app.Commands = []cli.Command{
    {
      Name:    "version",
      Aliases: []string{"v", "--version"},
      Usage:   "Calculate semantic version",
      Action: func(c *cli.Context) {
        branch, _ := ResolveGitBranch()
        version, _ := ResolveGradleVersion()

        semanticVersion := CalculateSemanticVersion(branch, version)
        fmt.Printf("%s", semanticVersion)
      },
    },
  }
}

// Calculates the semantic version depending on the gradle version (taken from the gradle.properties) and git branch
//
// Branch         gradle.version    semantic version
// develop        1.0.0-SNAPSHOT    1.0.0-SNAPSHOT
// some-branch    1.0.0-SNAPSHOT    1.0.0-some-branch-SNAPSHOT
// feature/x      1.0.0-SNAPSHOT    1.0.0-feature-x-SNAPSHOT
// defect/x       1.0.0-SNAPSHOT    1.0.0-feature-x-SNAPSHOT
// release/1.x    1.0.0             1.0.0-rc{BUILD}
// hotfix/1.x     1.0.0             1.0.0-rc{BUILD}
// master         1.0.0             1.0.0
func CalculateSemanticVersion(branch, version string) (string) {
  var versionRoot, versionExt, semanticVersion string

  // convert brenches like feature/XYZ to feature-xyz
  branch = strings.ToLower(strings.Replace(branch, "/", "-", -1))

  if strings.Contains(version, "-") {
    versionRoot = strings.TrimSpace(strings.Split(version, "-")[0])
    versionExt = strings.TrimSpace(strings.Split(version, "-")[1])
  } else {
    versionRoot = version
    versionExt = ""
  }

  // Calculate semantic version
  if branch == "develop" || branch == "master" {
    semanticVersion = versionRoot
  } else if strings.HasPrefix(branch, "release") || strings.HasPrefix(branch, "hotfix") {
    buildNumber := GetEnvVariable("BUILD_NUMBER", "0")
    semanticVersion = fmt.Sprintf("%v-rc%v", versionRoot, buildNumber)
  } else {
    semanticVersion = fmt.Sprintf("%v-%v", versionRoot, branch)
  }

  if len(versionExt) > 0 {
    semanticVersion = fmt.Sprintf("%v-%v", semanticVersion, versionExt)
  }

  return strings.TrimSpace(semanticVersion)
}

func ResolveGradleVersion() (string, error) {
  out, err := exec.Command("bash", "-c", "cat gradle.properties | grep version | cut -d'=' -f2").Output()
  return strings.TrimSpace(string(out[:])), err
}

func ResolveGitBranch() (string, error) {
  // Determine the git branch from env if running on CI, otherwise from git
  if IsEnvVariableDefined("BUILD_NUMBER") { // magic jenkins variable
    if IsEnvVariableDefined("CHANGE_ID") { // Are we building a PR?
      return GetEnvVariable("CHANGE_BRANCH", "unknown"), nil
    } else { // Not a PR
      return GetEnvVariable("BRANCH_NAME", "unknown"), nil
    }
  }
  // Not a CI build
  out, err := exec.Command("bash", "-c", "git rev-parse --abbrev-ref HEAD 2> /dev/null  || echo 'unknown'").Output()
  return strings.TrimSpace(string(out[:])), err
}

func IsEnvVariableDefined(key string) (bool) {
  _, exists := os.LookupEnv(key)
  return exists
}

func GetEnvVariable(key, fallback string) (string) {
  value, exists := os.LookupEnv(key)
  if !exists {
    value = fallback
  }
  return value
}

func main() {
  Info()
  Commands()

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}