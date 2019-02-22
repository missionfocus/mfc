#!/usr/bin/env sh

curl -X POST https://git.missionfocus.com/api/v4/projects/${CI_PROJECT_ID}/releases \
    -H "Content-Type: application/json" -H "PRIVATE-TOKEN: ${MARIO_PAT}" \
    -d @- << EOF
{
  "name": "mf-vault ${CI_COMMIT_TAG}",
  "tag_name": "${CI_COMMIT_TAG}",
  "ref": "${CI_COMMIT_REF_NAME}",
  "description": "Mission Focus Vault CLI.",
  "assets": {
    "links": [
      {
        "name": "Linux",
        "url": "https://git.missionfocus.com/open-source/mf-vault/-/jobs/${CI_JOB_ID}/artifacts/raw/out/binaries/x86_64-linux/mf-vault"
      },
      {
        "name": "Linux (Checksum)",
        "url": "https://git.missionfocus.com/open-source/mf-vault/-/jobs/${CI_JOB_ID}/artifacts/raw/out/binaries/x86_64-linux/mf-vault.checksum"
      },
      {
        "name": "macOS",
        "url": "https://git.missionfocus.com/open-source/mf-vault/-/jobs/${CI_JOB_ID}/artifacts/raw/out/binaries/x86_64-darwin/mf-vault"
      },
      {
        "name": "macOS (Checksum)",
        "url": "https://git.missionfocus.com/open-source/mf-vault/-/jobs/${CI_JOB_ID}/artifacts/raw/out/binaries/x86_64-darwin/mf-vault.checksum"
      }
    ]
  }
}
EOF
