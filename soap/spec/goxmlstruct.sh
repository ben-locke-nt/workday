#!/bin/bash

~/go/bin/goxmlstruct --package-name human_resources \
    --use-pointers-for-optional-fields \
    --named-types \
    --named-root \
    --ignore-errors \
     --top-level-attributes \
    ./Get_Workers_Request.xml > get_workers_request_model.go

~/go/bin/goxmlstruct --package-name human_resources \
    --use-pointers-for-optional-fields \
    --named-types \
    --named-root \
    --ignore-errors \
     --top-level-attributes \
    ./Get_Workers_Response.xml > get_workers_response_model.go
