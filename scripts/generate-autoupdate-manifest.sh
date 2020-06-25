#!/usr/bin/env sh
# Generates and uploads a manifest that points to the latest mfc version in S3.
set -e

CI_COMMIT_TAG=${CI_COMMIT_TAG:=0.0.0}
BUCKET=${BUCKET:=public.missionfocus.com}
PREFIX=${PREFIX:=mfc}
S3_WEBSITE_ENDPOINT=http://${BUCKET}.s3-website-us-east-1.amazonaws.com

mkdir -p ./scripts/files
cat << EOF > ./scripts/files/manifest.yaml
version: ${CI_COMMIT_TAG}
url:
  linux: ${S3_WEBSITE_ENDPOINT}/$PREFIX/mfc_${CI_COMMIT_TAG}_linux
  darwin: ${S3_WEBSITE_ENDPOINT}/$PREFIX/mfc_${CI_COMMIT_TAG}_darwin
  windows: ${S3_WEBSITE_ENDPOINT}/$PREFIX/mfc_${CI_COMMIT_TAG}_windows.exe
EOF

aws s3 cp ./scripts/files/manifest.yaml s3://$BUCKET/$PREFIX/manifest.yaml
rm ./scripts/files/manifest.yaml
