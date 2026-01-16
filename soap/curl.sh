#!/bin/bash
# Copyright 2025 Nametag Inc.
#
# All information contained herein is the property of Nametag Inc.. The
# intellectual and technical concepts contained herein are proprietary, trade
# secrets, and/or confidential to Nametag, Inc. and may be covered by U.S.
# and Foreign Patents, patents in process, and are protected by trade secret or
# copyright law. Reproduction or distribution, in whole or in part, is
# forbidden except by express written permission of Nametag, Inc.

AUTH_TOKEN=$(cat ../.oauth2_token_cache.json | jq -r '.access_token')
SOAP_URL=https://impl-services1.wd12.myworkday.com/ccx/service/nametag_dpt1/Human_Resources/v46.0
curl -X POST --header "Authorization: Bearer ${AUTH_TOKEN}" --header 'Accept: text/xml; charset=utf-8' --data '@./spec/request.xml' "${SOAP_URL}"
