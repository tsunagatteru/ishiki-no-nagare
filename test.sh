#!/bin/bash


for i in 0{1..45}
do
	curl http://localhost:8080/api/create-post \
		--include \
		--header "Content-Type: text/html; charset=utf-8" \
		--request "POST" \
		--data 'test'
done
