export token="xyzzy"
export bmc="127.0.0.1:8080"
curl -k -H "X-Auth-Token: $token" -X POST http://${bmc}/redfish/v1/Systems/master-0.bm.lo/Actions/ComputerSystem.Reset -d '{"ResetType": "ForceOff"}'
curl -k -H "X-Auth-Token: $token" -X POST http://${bmc}/redfish/v1/Systems/1/Actions/ComputerSystem.Reset -d '{"ResetType": "ForceOff"}'

