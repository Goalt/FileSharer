FROM mcr.microsoft.com/vscode/devcontainers/go:1.17

USER vscode
RUN go get github.com/cweill/gotests/gotests
RUN go get github.com/google/wire/cmd/wire
RUN go get github.com/golang/mock/gomock
RUN go get github.com/golang/mock/mockgen

WORKDIR /workspace

COPY . .

EXPOSE 8080