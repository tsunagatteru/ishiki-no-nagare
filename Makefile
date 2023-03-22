install:
        @echo -e '\033[1;32mInstalling the program...\033[0m'
        useradd --shell /bin/bash --system --home-dir "/var/lib/inn" inn
        install -d "/var/lib/inn"
        chown -R "inn:inn" "/var/lib/inn"
        install -d "/etc/inn"
        chown -R "inn:inn" "/etc/inn"
        install -Dm 755 ./cmd/inn -t "/usr/bin"
        install -Dm 644 ./config.yml.default -T "/etc/inn/config.yml"
uninstall:
        @echo -e '\033[1;32mInstalling the program...\033[0m'
		deluser inn
		rm -rf /var/lib/inn
		rm -rf /etc/inn
		rm /usr/bin/inn
