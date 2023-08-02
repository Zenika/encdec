#!/usr/bin/env bash

PKGDIR=encdec-`cat current_pkg_release`_amd64

mkdir -p ${PKGDIR}/opt/bin ${PKGDIR}/DEBIAN
mv control ${PKGDIR}/DEBIAN/
mv preinst ${PKGDIR}/DEBIAN/

echo "Building binary from source"
cd ../src
go build -o ../__debian/${PKGDIR}/opt/bin/encdec .
strip ../__debian/${PKGDIR}/opt/bin/encdec
sudo chown 0:0 ../__debian/${PKGDIR}/opt/bin/encdec

echo "Binary built. Now packaging..."
cd ../__debian/
dpkg-deb -b ${PKGDIR}
