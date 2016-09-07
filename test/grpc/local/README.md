# Files in this folder


## nginx.conf

This file is used for testing grpc and transcoding in local environment.

 - Download your service account secret file here as "service_account.json"

 - At top folder run:

   script/run_nginx -c $PWD/test/grpc/local/nginx.conf

 - Start backend grpc test as:

   bazel-bin/test/grpc/grpc-test-server 0.0.0.0:8081


## Service Config

The `service.json` service config file is generated from
`grpc-test.proto` and `grpc-test.yaml` using
[API Compiler](https://github.com/googleapis/api-compiler).

 - At top folder run:

   protoc test/grpc/grpc-test.proto -Itest/grpc -Ithird_party/config/ -Ithird_party/googleapis/ --include_imports --descriptor_set_out=out.descriptors

 - Follow the [API Compiler](https://github.com/googleapis/api-compiler)
   instructions to build it.

 - Copy endpoints.yaml to that api-compiler folder.

 - At api-compiler folder, run:

   SRC=SOURCE_CODE_TOP_LEVEL_FOLDER

   ./run.sh --configs $SRC/test/grpc/grpc-test.yaml --configs endpoints.yaml --descriptor $SRC/out.descriptors --json_out service.json


 - Copy the service.json to this local folder.

 - Push it to service-managment.

   gcloud alpha service-management deploy service.json