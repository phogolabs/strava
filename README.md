# Strava

A command line tool based on
[go-getter](https://github.com/hashicorp/go-getter/) from hashicorp. That allows
downloading file by providing a configuration file.  


## Getting Started

You can download all source into the given vendor directories by running:

```bash
$ strava vendor
```

An example of `strava.yaml`:

```yaml
---
vendor:
- name: sdk/google/api
  source:
   - https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto
   - https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/resource.proto
   - https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/field_behavior.proto
   - https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto
   - https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/httpbody.proto
   - https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/client.proto

- name: sdk/protoc-gen-openapiv2/options
  source:
  - https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/master/protoc-gen-openapiv2/options/annotations.proto
  - https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/master/protoc-gen-openapiv2/options/openapiv2.proto
```
