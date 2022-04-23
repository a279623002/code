#! /bin/bash

run() {
  project="$1"
  cd src/$project
  go mod tidy
  if [ -f $project ]; then
    rm $project
  fi

  go build -o $project
  chmod +x $project

  sudo supervisorctl restart $srv:*
  cd ../../

}

if [ $1 == "all" ]; then
  for srv in $(ls src); do
    if [[ ${srv:0-4} == "-srv" ]]; then
#      if [[ ${srv} != "api-srv" ]]; then
        echo "run $srv"
        run $srv
#      fi
    fi
  done
else
  for srv in "$@"; do
    if [[ "${srv:0-4}" != "-srv" ]]; then
      srv="${srv}-srv"
    fi
    echo "run $srv"
    run ${srv}
  done
fi
