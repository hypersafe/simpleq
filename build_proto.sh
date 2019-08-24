#!/bin/sh
protoc -I ${PWD}/proto/ message_service.proto --go_out=plugins=grpc:services