export theversion=1.`./incrementbuild.sh`.0-1
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build esxiredfish.go -ldflags "-X main.sha1ver=$theversion}"
export theversion=1.`./incrementbuild.sh`.0-1
export thedir=`pwd | sed 's:/*$::'`
echo ${thedir}
go-bin-rpm generate --file rpm.json --version ${theversion} -o esxiredfish-${theversion}.rpm -b ${thedir}/buildarea --arch x86_64
mv esxiredfish-${theversion}.rpm rpms
(cd rpms;ls -t | tail -n +2 | xargs rm -- )
rm -r -f rpms/repodata
createrepo rpms
#yum clean all
git add *
git commit -m "Update after build"
git push
yum update -y esxiredfish

