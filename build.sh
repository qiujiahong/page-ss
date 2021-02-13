version=v1.5_$(date +%Y%m%d)
repo=page-ss

rm -rf dist/
# 打包mac版本程序
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ./
mkdir -p dist/${version}/${repo}_mac
mv ${repo} dist/${version}/${repo}_mac
cp ./script/run.sh dist/${version}/${repo}_mac
#cp -rf config dist/${version}/${repo}_mac
cp docs/help.md dist/${version}/${repo}_mac


# 打包Linux 版本程序
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./
mkdir -p dist/${version}/${repo}_linux
mv ${repo} dist/${version}/${repo}_linux
cp ./script/run.sh dist/${version}/${repo}_linux
#cp -rf config dist/${version}/${repo}_linux
cp docs/help.md dist/${version}/${repo}_linux


# 打包windows版本程序
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ./
mkdir -p dist/${version}/${repo}_win
mv ${repo}.exe dist/${version}/${repo}_win
cp ./script/run.bat dist/${version}/${repo}_win
#cp -rf config dist/${version}/${repo}_win
cp docs/help.md dist/${version}/${repo}_win

# 拷贝

cp -rf dist/${version}  dist/latest

cd dist
tar czvf ${repo}-${version}.tar.gz ${version}