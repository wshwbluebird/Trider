#!/usr/bin/env bash
now_dir=`pwd`
govendor add +external
govendor update +external
