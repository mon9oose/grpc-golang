#!/bin/bash

protoc -I=. --go_out=. --go-gprc_out=. ./protos/chat.proto