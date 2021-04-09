


### Installing
Add the repo. 
 yum-config-manager --add-repo https://raw.githubusercontent.com/glennswest/esxiredfish/main/esxiredfish.repo 
Cleanup any cache to make sure you get latest
  yum clean all
Install the service. 
 yum install esxiredfish

### Building
A build.sh script is included.
You will need golang installed, as well as the rpm building tool.

