#!/bin/bash

: ${BASE_URL:=https://127.0.0.1}
: ${USERNAME_CLI:=admin@example.com}
: ${PASSWORD_CLI:=cloudbreak}
: ${CLI_TEST_FILES:=spec/integration/*.rb}

source /usr/share/rvm/scripts/rvm

if [[ -z "$(echo $TARGET_CBD_VERSION)" ]]; then
	export TARGET_CBD_VERSION=$(curl -sk $BASE_URL/cb/info | grep -oP "(?<=\"version\":\")[^\"]*")
  if [[ -z "$(echo $TARGET_CBD_VERSION)" ]]; then
	    export TARGET_CBD_VERSION=MOCK
	fi
fi

echo "CBD Version: "$TARGET_CBD_VERSION
echo "CBD URL: "$BASE_URL
echo "CBD Username: "$USERNAME_CLI
echo "CBD Password: "$PASSWORD_CLI
echo "CLI Tests: "$CLI_TEST_FILES

curl --verbose --show-error --location --insecure -C - https://s3-us-west-2.amazonaws.com/cb-cli/cb-cli_"${TARGET_CBD_VERSION}"_$(uname)_x86_64.tgz | tar -xvz --directory /usr/local/bin

cb configure --server $BASE_URL --username $USERNAME_CLI --password $PASSWORD_CLI

mkdir -p tmp/aruba
rspec -f RspecJunitFormatter -o test-result.xml -f h $CLI_TEST_FILES | tee test-result.html | ruby -n spec/common/integration_formatter.rb