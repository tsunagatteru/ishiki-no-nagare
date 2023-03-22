install:
	@echo -e '\033[1;32mInstalling the program...\033[0m'
	useradd --shell /bin/bash --system --home-dir "/var/lib/inn" inn
	mkdir "/var/lib/inn"
	chwon -R "inn:inn" "/var/lib/inn"
	mkdir "/etc/inn"
	chwon -R "inn:inn" "/etc/inn"
	install -Dm 755 .cmd/inn -t "/usr/bin"
	install -Dm 644 ./config.yml.default -T "/etc/inn/config.yml"
