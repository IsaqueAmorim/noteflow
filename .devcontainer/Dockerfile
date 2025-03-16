# Use the official Go image as the base image
FROM golang:1.24.1-bullseye

# Set environment variables
ENV DEBIAN_FRONTEND=noninteractive
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOTEST =go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o 

# Set up workspace
WORKDIR /app

COPY .. .

# Install dependencies
RUN apt-get update && apt-get install -y \
    git \
    curl \
    wget \
    zsh \
    build-essential \
    locales \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Set up locale
#RUN locale-gen en_US.UTF-8 && update-locale LANG=en_US.UTF-8

# Install pre-commit
RUN curl -fsSL https://raw.githubusercontent.com/pre-commit/pre-commit.github.io/main/install-local.py | python3 - || \
    wget -qO- https://raw.githubusercontent.com/pre-commit/pre-commit.github.io/main/install-local.py | python3 -


# Set Zsh as the default shell
RUN chsh -s $(which zsh)

# Install GO Tools
RUN go install golang.org/x/tools/gopls@latest && go install github.com/cweill/gotests/gotests@latest && go install github.com/fatih/gomodifytags@latest && go install github.com/josharian/impl@latest && go install github.com/haya14busa/goplay/cmd/goplay@latest && go install github.com/go-delve/delve/cmd/dlv@latest && go install honnef.co/go/tools/cmd/staticcheck@latest

# Expose default Go debugging port
EXPOSE 2345

# Start Zsh shell by default
CMD ["zsh"]