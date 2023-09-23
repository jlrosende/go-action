FROM golang:1.21-bookworm

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# ********************************************************
# * Anything else you want to do like clean up goes here *
# ********************************************************

# Install libraries and other tools

RUN sh -c "$(wget -O- https://raw.githubusercontent.com/nektos/act/master/install.sh)"

RUN cd /usr/local && sh -c "$(wget -O- https://dl.dagger.io/dagger/install.sh)"

# [Optional] Set the default user. Omit if you want to keep the default as root.
USER $USERNAME

# Uses "Spaceship" theme with some customization. Uses some bundled plugins and installs some more from github
RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.5/zsh-in-docker.sh)" -- \
    -t https://github.com/denysdovhan/spaceship-prompt \
    -p git \
    -p ssh-agent \
    -p golang \
    -p https://github.com/zsh-users/zsh-autosuggestions \
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
