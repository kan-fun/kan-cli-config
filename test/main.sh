#!/bin/bash

set -e

./kan-cli-config init --access-key 123 --secret-key 456
diff ~/.kanrc.yml ./test/init.yml