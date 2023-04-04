pkgname=dirf2.0
pkgver=1.0.0
pkgrel=1
pkgdesc="A simple tool for directory listing"
arch=('x86_64')
url="https://github.com/rishavmngo/dirf2.0"
license=('MIT')
depends=('go')
source=("git+$url")
md5sums=('SKIP')

build() {
cd $pkgname
  go build -o dirf2 
}

package() {
cd $pkgname
  install -Dm755 dirf2 "$pkgdir/usr/bin/dirf2"
}
