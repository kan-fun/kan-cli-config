#!/bin/bash

set -e
mv ./kan-cli-config ./kan-config

./kan-config init --access-key 123 --secret-key 456
diff ~/.kanrc.yml ./test/init.yml