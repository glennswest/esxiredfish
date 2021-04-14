


### Installing 
#### Add the repo.  
  yum-config-manager --add-repo https://raw.githubusercontent.com/glennswest/esxiredfish/main/esxiredfish.repo  
#### Cleanup any cache to make sure you get latest
  yum clean all
#### Install the service.  
 yum install esxiredfish 
#### Enable the firewall 
    firewall-cmd --zone=public --add-service=http
    firewall-cmd --zone=public --permanent --add-service=http
    firewall-cmd --reload
#### Edit /etc/esxiredfish.yaml and put in your esxi server ip and user 
server: 
  host: 192.168.1.150 
  user: root 
#### Start the server
systemctl enable esxiredfish
systemctl start esxiredfish
systemctl status esxiredfish




### Building 
A build.sh script is included. 
You will need golang installed, as well as the rpm building tool. 

