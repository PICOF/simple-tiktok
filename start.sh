#!/usr/bin/env bash

my_module=('comment' 'favorite' 'feed' 'message' 'publish' 'relation' 'user')

case "$1" in
  build)
    echo "Start building..."
    cd ./cmd/api || exit
    go build
    echo "Server build successfully"
    for i in "${!my_module[@]}";
    do
      cd ../"${my_module[$i]}" || exit
      ./build.sh
      echo "Module ${my_module[$i]} build successfully"
    done
    echo "Build successfully!";;
  start)
    echo "Start running..."
    cd ./cmd/api || exit
    ./api
    echo "Server start successfully"
    for i in "${!my_module[@]}";
    do
      cd ../"${my_module[$i]}" || exit
      ./output/bootstrap.sh
      echo "Module ${my_module[$i]} start successfully"
    done
    echo "Finish.";;
esac