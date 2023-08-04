#!/bin/sh
set -e

# started as hassio addon
HASSIO_OPTIONSFILE=/data/options.json

if [ -f ${HASSIO_OPTIONSFILE} ]; then
    CONFIG=$(grep -o '"config_file": "[^"]*' ${HASSIO_OPTIONSFILE} | grep -o '[^"]*$')
    echo "Using config file: ${CONFIG}"
    
    SQLITE_FILE=$(grep -o '"sqlite_file": "[^"]*' ${HASSIO_OPTIONSFILE} | grep -o '[^"]*$')
    
    if [ ! -f "${CONFIG}" ]; then
        echo "Config not found. Please create a config under ${CONFIG}."
        echo "For details see robotui documentation at https://github.com/thommyho/robotui#readme."
    else
        if [ "${SQLITE_FILE}" ]; then
            echo "starting robotui: 'robotui_DATABASE_DSN=${SQLITE_FILE} robotui --config ${CONFIG}'"
            exec env robotui_DATABASE_DSN="${SQLITE_FILE}" robotui --config "${CONFIG}"
        else
            echo "starting robotui: 'robotui --config ${CONFIG}'"
            exec robotui --config "${CONFIG}"
        fi
    fi
else
    if [ "$1" == '"robotui"' ] || expr "$1" : '-*' > /dev/null; then
        exec robotui "$@"
    else
        exec "$@"
    fi
fi
