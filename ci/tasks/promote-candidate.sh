#!/usr/bin/env bash

set -e

: ${ALICLOUD_ACCESS_KEY_ID:?}
: ${ALICLOUD_ACCESS_KEY_SECRET:?}

source /etc/profile.d/chruby.sh
chruby 2.1.2

# Creates an integer version number from the semantic version format
# May be changed when we decide to fully use semantic versions for releases
integer_version=`cut -d "." -f1 release-version-semver/number`
echo $integer_version > promoted/integer_version

# debug
cp -r bosh-cpi-src promoted/repo

dev_release=$(echo $PWD/bosh-cpi-release/*.tgz)

pushd promoted/repo
  echo creating config/private.yml with blobstore secrets
  cat > config/private.yml << EOF
---
blobstore:
  s3:
    access_key_id: $ALICLOUD_ACCESS_KEY_ID
    secret_access_key: $ALICLOUD_ACCESS_KEY_SECRET
EOF

  echo "using bosh CLI version..."
  bosh version

  echo "finalizing CPI release..."
  bosh finalize release ${dev_release} --version $integer_version

  rm config/private.yml

  git diff | cat
  git add .

  git config --global user.email demon.wy@alibaba-inc.com
  git config --global user.name demonwy
  git commit -m "New final release v $integer_version"
popd