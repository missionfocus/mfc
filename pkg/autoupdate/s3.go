package autoupdate

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"runtime"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/blang/semver"
	"github.com/inconshreveable/go-update"
)

type s3Updater struct {
	sess   *session.Session
	bucket string
	prefix string
}

func NewS3Updater(sess *session.Session, bucket string, prefix string) Updater {
	return &s3Updater{sess, bucket, prefix}
}

func (u *s3Updater) Check(currentTagName string) (string, error) {
	curVer, err := semver.Parse(normalizeSemver(currentTagName))
	if err != nil {
		return "", fmt.Errorf("parsing current version: %w", err)
	}

	releases, err := u.getReleases()
	if err != nil {
		return "", fmt.Errorf("getting releases: %w", err)
	}
	for _, rel := range releases {
		if rel.Version.GT(curVer) {
			return rel.Version.String(), nil
		}
	}

	return "", nil
}

func (u *s3Updater) Update() error {
	releases, err := u.getReleases()
	if err != nil {
		return fmt.Errorf("getting releases: %w", err)
	}
	if len(releases) == 0 {
		return errors.New("no releases found")
	}

	var next *s3Release
	for _, rel := range releases {
		if rel.Platform == runtime.GOOS && (next == nil || rel.Version.GT(next.Version)) {
			next = rel
		}
	}
	if next == nil {
		return errors.New("no compatible release found")
	}

	downloader := s3manager.NewDownloader(u.sess)
	buf := &aws.WriteAtBuffer{}
	_, err = downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(u.bucket),
		Key:    aws.String(next.Key),
	})
	if err != nil {
		return fmt.Errorf("downloading release: %w", err)
	}

	if err := update.Apply(bytes.NewReader(buf.Bytes()), update.Options{}); err != nil {
		if rerr := update.RollbackError(err); rerr != nil {
			return fmt.Errorf("rolling back from failed update: %w", rerr)
		}
		return fmt.Errorf("updating: %w", err)
	}

	return nil
}

func (u *s3Updater) getReleases() ([]*s3Release, error) {
	svc := s3.New(u.sess)
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(u.bucket),
		Prefix: aws.String(u.prefix),
	})
	if err != nil {
		return nil, fmt.Errorf("listing releases: %w", err)
	}

	releases := make([]*s3Release, 0)
	for _, obj := range resp.Contents {
		r := &s3Release{}
		if r.Parse(*obj.Key) == nil {
			releases = append(releases, r)
		}
	}

	return releases, nil
}

type s3Release struct {
	Version  semver.Version
	Key      string
	Platform string
}

var s3ReleaseRe = regexp.MustCompile(`mfc_(v\d+\.\d+\.\d+)_(darwin|linux|windows)(?:\.exe)?$`)

// Parse parses a key to an S3 release. Format example: mfc_v2.0.0_darwin.tar.gz
func (r *s3Release) Parse(s string) error {
	match := s3ReleaseRe.FindStringSubmatch(s)
	if len(match) != 3 {
		return fmt.Errorf("failed to parse %s as S3 release", s)
	}

	r.Key = s
	r.Platform = match[2]

	v, err := semver.Parse(normalizeSemver(match[1]))
	if err != nil {
		return fmt.Errorf("failed to parse semantic version: %w", err)
	}
	r.Version = v

	return nil
}
