#!/bin/bash

session=$(curl -c - http://localhost:8080/api/login \
	 --include \
	 --header "Content-Type: application/x-www-form-urlencoded" \
	 --request "POST" \
	 --data "username=admin&password=password")
for i in $(eval echo {1..$1})
do
	message="message="
	message+=$(fortune)
	curl --cookie <(echo "$session") http://localhost:8080/api/admin/create-post \
		--include \
		--header "Content-Type: application/x-www-form-urlencoded" \
		--request "POST" \
		--data "$message"
done
