#!/bin/bash

sh get_structs.sh
python3 gen_structs.py '*_Request.xml' 'wd'
python3 gen_structs.py '*_Response.xml' ''
