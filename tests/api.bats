#!/usr/bin/env bats

@test "query bookmarks should be empty" {
  result="$(curl -s -X GET -H "${TEST_USER_1}" ${API_URI}/v1/bookmark/)"
  [ "$(echo $result | jq -c)" == '{"Bookmarks":[],"Remarks":[],"Clicks":[]}' ]
}

@test "add bookmark X" {
  result="$(curl -s -X POST -H "${TEST_USER_1}" ${API_URI}/v1/bookmark/remark/ -d 'url=https://github.com')"
  [ "$(echo $result | jq -c)" == '{"message":"ok"}' ]
}

@test "query bookmarks should have one bookmark with remark count 1" {
  result="$(curl -s -X GET -H "${TEST_USER_1}" ${API_URI}/v1/bookmark/)"
  [ "$(echo $result | jq -r .Bookmarks[0].Url)" == "https://github.com" ]
  [ "$(echo $result | jq -r .Bookmarks[0].Title)" != "https://github.com" ]
  [ "$(echo $result | jq -r .Bookmarks[0].RemarkCount)" == "1" ]
  [ "$(echo $result | jq '.Remarks | length')" == "1" ]
  [ "$(echo $result | jq -r .Bookmarks[0].ClickCount)" == "0" ]
  [ "$(echo $result | jq '.Clicks | length')" == "0" ]
}

@test "add bookmark X again" {
  result="$(curl -s -X POST -H "${TEST_USER_1}" ${API_URI}/v1/bookmark/remark/ -d 'url=https://github.com')"
  [ "$(echo $result | jq -c)" == '{"message":"ok"}' ]
}

@test "query bookmarks should have one bookmark with remark count 2" {
  result="$(curl -s -X GET -H "${TEST_USER_1}" ${API_URI}/v1/bookmark/)"
  [ "$(echo $result | jq -r .Bookmarks[0].Url)" == "https://github.com" ]
  [ "$(echo $result | jq -r .Bookmarks[0].RemarkCount)" == "2" ]
  [ "$(echo $result | jq '.Remarks | length')" == "2" ]
  [ "$(echo $result | jq -r .Bookmarks[0].ClickCount)" == "0" ]
  [ "$(echo $result | jq '.Clicks | length')" == "0" ]
}

@test "click bookmark X" {
  result="$(curl -s -X POST -H "${TEST_USER_1}" ${API_URI}/v1/bookmark/click/ -d id=$(curl -s -X GET -H "${TEST_USER_1}" ${API_URI}/v1/bookmark/ | jq -r .Bookmarks[0].ID))"
  [ "$(echo $result | jq -c)" == '{"message":"ok"}' ]
}

@test "query bookmarks should have one bookmark with remark count 2 and click count 1" {
  result="$(curl -s -X GET -H "${TEST_USER_1}" ${API_URI}/v1/bookmark/)"
  [ "$(echo $result | jq -r .Bookmarks[0].Url)" == "https://github.com" ]
  [ "$(echo $result | jq -r .Bookmarks[0].Title)" != "https://github.com" ]
  [ "$(echo $result | jq -r .Bookmarks[0].RemarkCount)" == "2" ]
  [ "$(echo $result | jq '.Remarks | length')" == "2" ]
  [ "$(echo $result | jq -r .Bookmarks[0].ClickCount)" == "1" ]
  [ "$(echo $result | jq '.Clicks | length')" == "1" ]
}

@test "click bookmark X again" {
  result="$(curl -s -X POST -H "${TEST_USER_1}" ${API_URI}/v1/bookmark/click/ -d id=$(curl -s -X GET -H "${TEST_USER_1}" ${API_URI}/v1/bookmark/ | jq -r .Bookmarks[0].ID))"
  [ "$(echo $result | jq -c)" == '{"message":"ok"}' ]
}

@test "query bookmarks should have one bookmark with remark count 2 and click count 2" {
  result="$(curl -s -X GET -H "${TEST_USER_1}" ${API_URI}/v1/bookmark/)"
  [ "$(echo $result | jq -r .Bookmarks[0].Url)" == "https://github.com" ]
  [ "$(echo $result | jq -r .Bookmarks[0].RemarkCount)" == "2" ]
  [ "$(echo $result | jq '.Remarks | length')" == "2" ]
  [ "$(echo $result | jq -r .Bookmarks[0].ClickCount)" == "2" ]
  [ "$(echo $result | jq '.Clicks | length')" == "2" ]
}

@test "add bookmark Y" {
  result="$(curl -s -X POST -H "${TEST_USER_1}" ${API_URI}/v1/bookmark/remark/ -d url=${API_URI}/swagger/doc.json)"
  [ "$(echo $result | jq -c)" == '{"message":"ok"}' ]
}

@test "query bookmarks should have bookmark Y with remark count 1" {
  result="$(curl -s -X GET -H "${TEST_USER_1}" ${API_URI}/v1/bookmark/)"
  [ "$(echo $result | jq -r .Bookmarks[1].Url)" == "${API_URI}/swagger/doc.json" ]
  [ "$(echo $result | jq -r .Bookmarks[1].Title)" == "${API_URI}/swagger/doc.json" ]
  [ "$(echo $result | jq -r .Bookmarks[1].RemarkCount)" == "1" ]
  [ "$(echo $result | jq '.Remarks | length')" == "3" ]
  [ "$(echo $result | jq -r .Bookmarks[0].ClickCount)" == "2" ]
  [ "$(echo $result | jq '.Clicks | length')" == "2" ]
}