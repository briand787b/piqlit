#!/bin/bash

# Need to have a docker-entrypoint.sh to get env vars into production Docker container
# without recompiling for every different environment

#!/bin/sh

# Replace env vars in JavaScript files
echo "Replacing env vars in JS"
for file in /usr/share/nginx/html/js/app.*.js;
do
  echo "Processing $file ...";

  # Use the existing JS file as template
  if [ ! -f $file.tmpl.js ]; then
    cp $file $file.tmpl.js
  fi

  envsubst '$VUE_APP_BACKEND_HOST' < $file.tmpl.js > $file
done

echo "Starting Nginx"
# MIGHT NEED TO REMOVE 'EXEC' SINCE IT WASN'T IN ORIGINAL SCRIPT
exec nginx -g 'daemon off;'