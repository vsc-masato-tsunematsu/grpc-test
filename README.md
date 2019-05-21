# gRPC test

## Prerequisites

* [Git](https://git-scm.com/)
  * [Git for Windows](https://git-scm.com/download/win)
  * [Git for Mac](https://git-scm.com/download/mac)
* [Docker](https://www.docker.com/)
  * [Docker Desktop for Windows (Requires Microsoft Hyper-V)](https://docs.docker.com/docker-for-windows/install/)
  * [Docker Toolbox on Windows](https://docs.docker.com/toolbox/toolbox_install_windows/)
  * [Docker Desktop for Mac](https://docs.docker.com/docker-for-mac/install/)

## Getting started
1. **Clone a copy of the repo & change to the directory** 
    ```
    git clone https://github.com/vsc-masato-tsunematsu/grpc-test
    cd grpc-test
    ```

2. **Generate the gRPC code**
    ```bash
    docker-compose run --rm protoc
    ```

3. **Run the server**
    ```bash
    docker-compose up -d server
    ```

4. **Run the client**
    ```bash
    docker-compose run --rm client hoge
    ```

5. **Down the server**
    ```bash
    docker-compose down
    ```
