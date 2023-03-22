build:
	mkdir -p build
	go build -o ./build/inn ./cmd/inn.go

install:
	@echo -e '\033[1;32mInstalling the program...\033[0m'
	useradd --shell /bin/bash --system --home-dir "/var/lib/inn" inn
	install -d "/var/lib/inn"
	chown -R "inn:inn" "/var/lib/inn"
	install -d "/etc/inn"
	chown -R "inn:inn" "/etc/inn"
	install -Dm 755 ./build/inn -t "/usr/bin"
	install -Dm 644 ./config.yml.default -t "/etc/inn"

uninstall:
	@echo -e '\033[1;32mInstalling the program...\033[0m'
	-deluser inn
	-rmdir /var/lib/inn
	-rmdir /etc/inn
	-rm /usr/bin/inn
