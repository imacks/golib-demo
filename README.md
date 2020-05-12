golib-demo
==========
This repo demonstrates how to build Golang packages as DLL/SO dynamic library.

Ref. https://github.com/golang/go/wiki/cgo


Windows DLL
-----------
Using example in src/goutil

Run script below in Linux env docker: cr.lizoc.com/dockerguys/golang:1.14

```bash
apk-install git mingw-w64-gcc
mkdir /go/src/testlab && cd /go/src/testlab
git clone https://github.com/imacks/golib-demo.git ./
cd src/goutil
GOOS=windows CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1 \
    go build  -ldflags="-s -w" -buildmode=c-shared -o /go/bin/libgoutil.dll ./main.go
ls /go/bin/libgoutil.*
```

Copy out to host and then FTP to Windows box:

```bash
docker cp testgo:/go/bin/ ./
```

Test dll works on Windows x64, using PowerShell/.NET:

```powershell
$libsrc = @'
   [DllImport("C:\\libgoutil.dll", EntryPoint="Greet")]
   public static extern void Greet();

   [DllImport("C:\\libgoutil.dll", EntryPoint="Add")]
   public static extern int Add(int a, int b);

   [DllImport("C:\\libgoutil.dll", EntryPoint="Minus")]
   public statc extern int Minus(int a, int b);
'@
$testlib = Add-Type -MemberDefinition $libsrc -Name 'TestLib' -PassThru -Namespace System.Runtime.InteropServices
$testlib::Greet()
$testlib::Add(5, 3)
$testlib::Minus(5, 3)
```

Linux .SO
---------
Building on Alpine still doesn't work. Use golang on debian buster (golang:1.14):

```bash
mkdir /go/src/testlab && cd /go/src/testlab
git clone https://github.com/imacks/golib-demo.git ./
cd src/goutil
# must prefix with 'lib'
GOOS=linux CGO_ENABLED=1 go build  -ldflags="-s -w" -buildmode=c-shared -o /go/bin/libgoutil.so ./main.go
ls /go/bin/libgoutil.*

cd ../cgtest
gcc -I/go/bin -o /go/bin/cgtest main.c -L/go/bin -lgoutil -Wl,-R '-Wl,$ORIGIN'
# test
/go/bin/cgtest
```
