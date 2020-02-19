DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
VERSION=0.1.0

docker build -t azure-service-operator-dev-container:$VERSION -f $DIR/Dockerfile $DIR/..
