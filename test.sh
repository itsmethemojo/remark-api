#!/bin/bash

if [ ! -d bats-core ]; then
  git clone -b v1.8.2 https://github.com/bats-core/bats-core.git
fi

# clean up data for tests

curl -s  -X DELETE localhost:8080/api/v1/bookmark/ 

echo -e "\n\nall Data cleaned \n"

TEST_USER_HEADER_PREFIX="Authorization: TEST_TOKEN_" \
API_URI=http://localhost:8080 \
bats-core/bin/bats tests/api.bats