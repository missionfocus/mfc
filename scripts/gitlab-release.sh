#!/usr/bin/env sh

curl -X POST "https://git.missionfocus.com/api/v4/projects/${CI_PROJECT_ID}/releases" \
    -H "Content-Type: application/json" -H "PRIVATE-TOKEN: ${MARIO_PAT}" \
    -d @- << EOF
{
  "name": "mfc ${CI_COMMIT_TAG}",
  "tag_name": "${CI_COMMIT_TAG}",
  "ref": "${CI_COMMIT_REF_NAME}",
  "description": "Mission Focus Vault CLI.",
  "assets": {
    "links": [
      {
        "name": "Linux",
        "url": "https://git.missionfocus.com/ours/code/tools/mfc/-/jobs/${CI_JOB_ID}/artifacts/raw/out/binaries/x86_64-linux/mfc"
      },
      {
        "name": "Linux (Checksum)",
        "url": "https://git.missionfocus.com/ours/code/tools/mfc/-/jobs/${CI_JOB_ID}/artifacts/raw/out/binaries/x86_64-linux/mfc.checksum"
      },
      {
        "name": "macOS",
        "url": "https://git.missionfocus.com/ours/code/tools/mfc/-/jobs/${CI_JOB_ID}/artifacts/raw/out/binaries/x86_64-darwin/mfc"
      },
      {
        "name": "macOS (Checksum)",
        "url": "https://git.missionfocus.com/ours/code/tools/mfc/-/jobs/${CI_JOB_ID}/artifacts/raw/out/binaries/x86_64-darwin/mfc.checksum"
      },
      {
        "name": "Windows",
        "url": "https://git.missionfocus.com/ours/code/tools/mfc/-/jobs/${CI_JOB_ID}/artifacts/raw/out/binaries/x86_64-windows/mfc.exe"
      },
      {
        "name": "Windows (Checksum)",
        "url": "https://git.missionfocus.com/ours/code/tools/mfc/-/jobs/${CI_JOB_ID}/artifacts/raw/out/binaries/x86_64-darwin/mfc.exe.checksum"
      }
    ]
  }
}
EOF
