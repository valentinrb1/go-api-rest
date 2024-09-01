.PHONY: all build install uninstall clean

BINARY_NAME=lab6
SERVICE_NAME=lab6.service
CONF_NAME=lab6.conf
GO=go

all: check_deps build

build:
	mkdir -p bin

	if [ ! -f go.mod ]; then \
		$(GO) mod init github.com/valentinrb1/go-api-rest.git; \
	fi

	$(GO) mod tidy
	$(GO) mod download

	$(GO) build -o bin/$(BINARY_NAME)

install: check_deps build
	cp bin/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

	cp config/$(SERVICE_NAME) /etc/systemd/system/$(SERVICE_NAME)

	cp config/$(CONF_NAME) /etc/nginx/conf.d/$(CONF_NAME)

	systemctl restart nginx
	systemctl enable $(SERVICE_NAME)
	systemctl start $(SERVICE_NAME)

uninstall:
	systemctl stop $(SERVICE_NAME)
	systemctl disable $(SERVICE_NAME)

	rm -f /etc/systemd/system/$(SERVICE_NAME)

	rm -f /usr/local/bin/$(BINARY_NAME)

clean:
	rm -rf build
	rm -rf bin

check_deps:
	@command -v $(GO) >/dev/null 2>&1 || { echo "Go is not installed. Please install Go and try again."; exit 1; }

	@command -v systemctl >/dev/null 2>&1 || { echo "systemd is not installed. Please install systemd and try again."; exit 1; }

	@command -v ssh >/dev/null 2>&1 || { echo >&2 "SSH is not installed. Please install SSH and try again."; exit 1; }
