awk -F '= ' '/CURRENT_BUILD/{$2=$2+1";"}1' OFS='= ' projectbuild
go build esxiredfish.go
export theversion=1.`./incrementbuild.sh`.0-1
export thedir=`pwd | sed 's:/*$::'`
echo ${thedir}
go-bin-rpm generate --file rpm.json --version ${theversion} -o esxiredfish-${theversion}.rpm -b ${thedir}/buildarea --arch x86_64
mv esxiredfish-${theversion}.rpm rpms
(cd rpms;ls -t | tail -n +2 | xargs rm -- )
rm -r -f rpms/repodata
createrepo rpms
yum clean all
