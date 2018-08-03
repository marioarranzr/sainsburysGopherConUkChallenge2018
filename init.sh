#!/bin/bash

set -eo pipefail

go install

go build

./sainsburysGopherConUkChallenge2018 --inputString "<><"