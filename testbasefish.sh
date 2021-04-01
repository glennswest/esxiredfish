export token="xyzzy"
export bmc="127.0.0.1:8080"
curl -k -H "X-Auth-Token: $token" -X GET http://${bmc}/redfish/v1 

