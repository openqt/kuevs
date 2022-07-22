# plan

## Init a project
1. Create a project in github.com
2. Add main.go
3. go mod init
4. Add Makeilfe

## Watch k8s events
1. Import client-go
2. Write events watch
```json
{
  "Type": "ADDED",
  "Object": {
    "metadata": {
      "name": "node-exporter-66ct8.16fb7f27f5fe5722",
      "namespace": "kubesphere-monitoring-system",
      "uid": "7d4a5742-dd60-4ca8-9f0d-31795ff3205f",
      "resourceVersion": "23507808",
      "creationTimestamp": "2022-06-24T07:51:37Z",
      "managedFields": [
        {
          "manager": "k3s",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2022-06-24T07:51:37Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:count": {},
            "f:firstTimestamp": {},
            "f:involvedObject": {},
            "f:lastTimestamp": {},
            "f:message": {},
            "f:reason": {},
            "f:source": {
              "f:component": {},
              "f:host": {}
            },
            "f:type": {}
          }
        }
      ]
    },
    "involvedObject": {
      "kind": "Pod",
      "namespace": "kubesphere-monitoring-system",
      "name": "node-exporter-66ct8",
      "uid": "e707165d-da17-45bd-83a0-05d12e18a27f",
      "apiVersion": "v1",
      "resourceVersion": "12289573",
      "fieldPath": "spec.containers{kube-rbac-proxy}"
    },
    "reason": "BackOff",
    "message": "Back-off restarting failed container",
    "source": {
      "component": "kubelet",
      "host": "orch.novalocal"
    },
    "firstTimestamp": "2022-06-24T07:51:37Z",
    "lastTimestamp": "2022-07-21T08:36:39Z",
    "count": 178758,
    "type": "Warning",
    "eventTime": null,
    "reportingComponent": "",
    "reportingInstance": ""
  }
}
```

## save to sqlite
1. Design table structure

namespace | timestamp | type | reason | object | source | message | count | name | uid | invol_kind | invol_uid | last_stamp | operation

uid | name | namespace | object | reason | message | source | created_at | updated_at | count | type

```json
{
  "Type": "ADDED",
  "Object": {
    "metadata": {
      "name": "node-exporter-66ct8.16fb7f27f5fe5722",
      "namespace": "kubesphere-monitoring-system",
      "uid": "7d4a5742-dd60-4ca8-9f0d-31795ff3205f"
    },
    "involvedObject": {
      "kind": "Pod",
      "name": "node-exporter-66ct8"
    },
    "reason": "BackOff",
    "message": "Back-off restarting failed container",
    "source": {
      "component": "kubelet",
      "host": "orch.novalocal"
    },
    "firstTimestamp": "2022-06-24T07:51:37Z",
    "lastTimestamp": "2022-07-21T08:36:39Z",
    "count": 178758,
    "type": "Warning",
  }
}
```

2. Write save code
