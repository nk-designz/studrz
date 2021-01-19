#!/bin/sh

echo -e "Using API: $API_PROTO://$API_HOST";

OLD_HOST="http:\/\/localhost:42069";
NEW_HOST="$API_PROTO:\/\/$API_HOST";

find /usr/share/nginx/html -type f -print0 | xargs -0 sed -i "s/$OLD_HOST/$NEW_HOST/g";
nginx -g "daemon off;"