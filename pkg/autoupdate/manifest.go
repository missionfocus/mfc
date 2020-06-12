package autoupdate

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/blang/semver"
	"github.com/inconshreveable/go-update"
	"gopkg.in/yaml.v2"
)

type manifest struct {
	Version string `yaml:"version"`
	URL     struct {
		Linux   string `yaml:"linux"`
		Darwin  string `yaml:"darwin"`
		Windows string `yaml:"windows"`
	} `yaml:"url"`
}

func (m *manifest) URLForPlatform(platform string) (string, error) {
	var url string

	switch runtime.GOOS {
	case "linux":
		url = m.URL.Linux
	case "darwin":
		url = m.URL.Darwin
	case "windows":
		url = m.URL.Windows
	default:
		return "", fmt.Errorf("unsupported platform %s", runtime.GOOS)
	}

	return url, nil
}

type ManifestUpdater struct {
	ManifestURL string
	HTTPClient  *http.Client
}

func (u *ManifestUpdater) Check(currentTagName string) (string, error) {
	curVer, err := semver.Parse(normalizeSemver(currentTagName))
	if err != nil {
		return "", fmt.Errorf("parsing current version: %w", err)
	}

	manifest, err := u.getManifest()
	if err != nil {
		return "", err
	}

	nextVer, err := semver.Parse(normalizeSemver(manifest.Version))
	if err != nil {
		return "", fmt.Errorf("parsing next version: %w", err)
	}

	if nextVer.GT(curVer) {
		return nextVer.String(), nil
	}

	return "", nil
}

func (u *ManifestUpdater) Update() error {
	manifest, err := u.getManifest()
	if err != nil {
		return err
	}

	url, err := manifest.URLForPlatform(runtime.GOOS)
	if err != nil {
		return err
	}

	res, err := u.HTTPClient.Get(url)
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

func (u *ManifestUpdater) getManifest() (*manifest, error) {
	res, err := u.HTTPClient.Get(u.ManifestURL)
	if err != nil {
		return nil, err
	}

	m := &manifest{}
	if err := yaml.NewDecoder(res.Body).Decode(m); err != nil {
		return nil, err
	}

	return m, err
}
