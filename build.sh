version=v1.5
repo=page-ss

rm -rf dist/
# 打包mac版本程序
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ./
mkdir -p dist/${version}/${repo}_mac_${version}
mv ${repo} dist/${version}/${repo}_mac_${version}
cp ./script/run.sh dist/${version}/${repo}_mac_${version}
#cp -rf config dist/${version}/${repo}_mac_${version}
cp docs/help.md dist/${version}/${repo}_mac_${version}


# 打包Linux 版本程序
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./
mkdir -p dist/${version}/${repo}_linux_${version}
mv ${repo} dist/${version}/${repo}_linux_${version}
cp ./script/run.sh dist/${version}/${repo}_linux_${version}
#cp -rf config dist/${version}/${repo}_linux_${version}
cp docs/help.md dist/${version}/${repo}_linux_${version}


# 打包windows版本程序
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ./
mkdir -p dist/${version}/${repo}_win_${version}
mv ${repo}.exe dist/${version}/${repo}_win_${version}
cp ./script/run.bat dist/${version}/${repo}_win_${version}
#cp -rf config dist/${version}/${repo}_win_${version}
cp docs/help.md dist/${version}/${repo}_win_${version}

cd dist
tar czvf ${repo}-${version}.tar.gz ${version}