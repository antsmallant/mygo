#!/bin/bash

protoc --go_out=./ --go_opt=module=github.com/antsmallant/testgo ./raw_pbs/*.proto