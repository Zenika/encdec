all: 
	cd ../src
	CGO_ENABLED=0 go build . -o ./encdec.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.21.3 amd64
	sudo apk update --no-cache

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/encdec.exe $(DESTDIR)$(PREFIX)/bin/encdec
	strip $(DESTDIR)$(PREFIX)/bin/encdec
