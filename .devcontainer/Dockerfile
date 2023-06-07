
FROM golang:1.20.0-bullseye

ARG USER_NAME=vscode

RUN apt-get update \
 && apt-get install -y sudo zsh

RUN adduser --disabled-password --gecos '' $USER_NAME
RUN adduser $USER_NAME sudo
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

USER $USER_NAME

WORKDIR /home/$USER_NAME

RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.5/zsh-in-docker.sh)" -- \
    -t https://github.com/denysdovhan/spaceship-prompt \
    -p https://github.com/zsh-users/zsh-autosuggestions \
    -p https://github.com/zsh-users/zsh-syntax-highlighting \
    -p https://github.com/zsh-users/zsh-completions \
    -a "SPACESHIP_TIME_SHOW=false" \
    -a "SPACESHIP_TIME_COLOR=8" \
    -a "SPACESHIP_USER_SHOW=always" \
    -a "SPACESHIP_HOST_SHOW=always" \
    -a "SPACESHIP_EXEC_TIME_COLOR=8" \
    -a "SPACESHIP_EXEC_TIME_ELAPSED=2" \
    -a "SPACESHIP_EXIT_CODE_SHOW=true" \
    -a "SPACESHIP_EXIT_CODE_PREFIX=[" \
    -a "SPACESHIP_EXIT_CODE_SUFFIX=]" \
    -a "SPACESHIP_EXIT_CODE_SYMBOL=✘ " \
    -a "SPACESHIP_KUBECTL_SHOW=true" \
    -a "SPACESHIP_KUBECTL_COLOR=blue" \
    -a "SPACESHIP_KUBECTL_SYMBOL=" \
    -a "SPACESHIP_KUBECTL_VERSION_SHOW=false" \
    -a "SPACESHIP_KUBECTL_CONTEXT_SHOW=true" \
    -a "SPACESHIP_KUBECTL_CONTEXT_SHOW_NAMESPACE=true" \
    -a "SPACESHIP_PACKAGE_SHOW=false" \
    -a "SPACESHIP_NODE_SHOW=false" \
    -a "SPACESHIP_PHP_SHOW=false" \
    -a "SPACESHIP_DOCKER_SHOW=false" \
    -a "SPACESHIP_ASYNC_SHOW=false" \
    -a "SPACESHIP_GIT_ASYNC=false" \
    -a "SPACESHIP_GIT_BRANCH_ASYNC=false" 

CMD /bin/zsh