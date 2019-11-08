package autoupdate

import (
	"encoding/json"
	"fmt"
	"github.com/blang/semver"
	"github.com/inconshreveable/go-update"
	"net/http"
	"net/url"
	"path"
	"runtime"
	"strings"
)

type GitLabUpdater struct {
	ProjectURL *url.URL
	HTTPClient http.Client
}

type GitLabRelease struct {
	TagName string `json:"tag_name"`
	Assets  GitLabAssets
}

type GitLabAssets struct {
	Links []GitLabAssetLink `json:"links"`
}

type GitLabAssetLink struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (u *GitLabUpdater) Check(currentTagName string) (string, error) {
	releases, err := u.getReleases()
	if err != nil {
		return "", err
	}

	curVer, err := semver.Parse(normalizeSemver(currentTagName))
	if err != nil {
		return "", err
	}
	latestVer, err := semver.Parse(normalizeSemver(releases[0].TagName))
	if err != nil {
		return "", err
	}

	if latestVer.GT(curVer) {
		return latestVer.String(), nil
	}

	return "", nil
}

func (u *GitLabUpdater) Update() error {
	releases, err := u.getReleases()
	if err != nil {
		return err
	}

	name, err := assetNameForPlatform(runtime.GOOS)
	if err != nil {
		return err
	}

	latestURL := ""
	for _, link := range releases[0].Assets.Links {
		if link.Name == name {
			latestURL = link.URL
		}
	}
	if latestURL == "" {
		return fmt.Errorf("asset not found for platform %s", name)
	}

	res, err := u.HTTPClient.Get(latestURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := update.Apply(res.Body, update.Options{}); err != nil {
		if rerr := update.RollbackError(err); rerr != nil {
			return fmt.Errorf("failed to rollback from failed update: %v", rerr)
		}
		return err
	}

	return nil
}

func assetNameForPlatform(platform string) (string, error) {
	switch platform {
	case "darwin":
		return "macOS", nil
	case "linux":
		return "Linux", nil
	case "windows":
		return "Windows", nil
	default:
		return "", fmt.Errorf("unsupported platform %s", platform)
	}
}

func (u *GitLabUpdater) getReleases() ([]GitLabRelease, error) {
	r := *u.ProjectURL
	r.Path = path.Join(r.Path, "releases")
	res, err := u.HTTPClient.Get(r.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	releases := make([]GitLabRelease, 0)
	if err := json.NewDecoder(res.Body).Decode(&releases); err != nil {
		return nil, err
	}
	return releases, nil
}

func normalizeSemver(ver string) string {
	return strings.Trim(ver, "v")
}
