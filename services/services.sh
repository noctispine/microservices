serviceDirs=("users" "auth")

if [ $# -ne 1 ]
  then
    echo "Write 'up' or 'down' as argument"
    exit 1
fi

if [ $1 == "up" ]; then
    for serviceDir in ${serviceDirs[@]}; do
        docker compose -f ./${serviceDir}/dev.docker-compose.yml up -d
    done
    echo "\n---------\nCONTAINERS ARE READY\n---------"
    exit 0
fi

if [ $1 == "down" ]; then
    for ((i=${#serviceDirs[@]} -1; i>=0; i--)); do
        docker compose -f ./${serviceDirs[i]}/dev.docker-compose.yml down
    done
    echo "\n---------\nCONTAINERS HAVE BEEN STOPPED THEN REMOVED\n---------"
    exit 0
fi

echo "Write 'up' or 'down' as argument"
exit 1
