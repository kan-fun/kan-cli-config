# wget http://gosspublic.alicdn.com/ossutil/1.6.10/ossutil64
# chmod 755 ossutil64
# ./ossutil64 config -e oss-cn-beijing.aliyuncs.com -i $ACCESSKEY -k $SECRETKEY
# ./ossutil64 rm -rf oss://kan-bin/linux --include "kan-config_*"
# ./ossutil64 cp -f ./kan-cli-config oss://kan-bin/linux/kan-config_0.0.1

echo $GITHUB_REF