#!/bin/bash

if [ ! -d bats-core ]; then
  git clone -b v1.8.2 https://github.com/bats-core/bats-core.git
fi

# clean up data for tests

curl -X DELETE localhost:8080/v1/bookmark/ 

echo -e "\n\nall Data cleaned \n"

TEST_USER_1="Authorization: TEST_TOKEN_1" API_URI=http://localhost:8080 bats-core/bin/bats tests/api.bats