{ pkgs ? import (builtins.fetchGit {
         # Descriptive name to make the store path easier to identify                
         name = "dev-go";                                                 
         url = "https://github.com/NixOS/nixpkgs";                       
         ref = "refs/heads/nixpkgs-unstable";                     
         rev = "7cf5ccf1cdb2ba5f08f0ac29fc3d04b0b59a07e4"; 
}) {} }:

with pkgs;

mkShell {
  buildInputs = [
    clang-tools
    gitlint
    gnupg
    go_1_19
    go-tools
    go-mockery
    gogetdoc
    golangci-lint
    goreleaser
    gosec
    gotools
    gofumpt
    golint
    mysql80
    openapi-generator-cli
    postgresql
    pre-commit
    protobuf
    protoc-gen-go
    protoc-gen-go-grpc
    ngrok
  ];

  shellHook =
    ''
      # Setup the binaries installed via `go install` to be accessible globally.
      export PATH="$(go env GOPATH)/bin:$PATH"

      # Install pre-commit hooks.
      pre-commit install

      # Install Go binaries.
      which protoc-gen-grpc-gateway || go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
      which protoc-gen-openapiv2 || go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
      which enumer || go install github.com/dmarkham/enumer@v1.5.3
      which gocritic || go install github.com/go-critic/go-critic/cmd/gocritic@latest
      which goreturns || go install github.com/sqs/goreturns@latest
      which swag || go get -u github.com/swaggo/swag
      which mockgen || go install github.com/golang/mock/mockgen@v1.6.0
      
      # Add the repo shared gitconfig
      git config --local include.path ../.gitconfig

      # Add the ngrok account
      ngrok config add-authtoken 2SUGILEFSKRGKXVuxBGTTcxdjrp_3aqq7Mw1Tx5sTQ7vhHMnN
      
      # Clear the terminal screen.
      clear
    '';
}
