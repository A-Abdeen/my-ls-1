echo Run both my-ls and the system command ls with no arguments.
go run .
echo Run both my-ls and the system command ls with the arguments: "<file name>".
go run . main.go
echo Run both my-ls and the system command ls with the arguments: "<directory name>".
go run . pkg
echo Run both my-ls and the system command ls with the flag: "-l".
go run . -l
echo Run both my-ls and the system command ls with the arguments: "-l <directory name>".
go run . -l pkg
echo Run both my-ls and the system command ls with the flag: "-R", in a directory with folders in it.
go run . -R
echo Run both my-ls and the system command ls with the flag: "-a".
go run . -a
echo Run both my-ls and the system command ls with the flag: "-r".
go run . -r
echo Run both my-ls and the system command ls with the flag: "-t".
go run . -t
echo Run both my-ls and the system command ls with the flag: "-la".
go run . -la
echo Run both my-ls and the system command ls with the arguments: "-l -t <directory name>".
go run . -l -t pkg
echo Run both my-ls and the system command ls with the arguments: "-lRr <directory name>", in which the directory chosen contains folders.
go run . -lRr pkg
echo Run both my-ls and the system command ls with the arguments: "-l <directory name> -a <file name>"
go run . -l pkg -a mysplit.go
echo Run both my-ls and the system command ls with the arguments: "-lR <directory name>///<sub directory name>/// <directory name>/<sub directory name>/"
go run . -lR pkg///test/// pkg/test/
echo Create directory with - name and run both my-ls and the system command ls with the arguments: 
go run . -
echo Create file and link for this file and run both my-ls-1 and the system command ls with the arguments: "-l <symlink file>/"
go run . -l fl/
echo Create file and link for this file and run both my-ls-1 and the system command ls with the arguments: "-l <symlink file>"
go run . -l fl
echo Create directory that contains files and link for this directory and run both my-ls-1 and the system command ls with the arguments: "-l <symlink dir>/"
go run . -l pk9/
echo Create directory that contains files and link for this directory and run both my-ls-1 and the system command ls with the arguments: "-l <symlink dir>"
go run . -l pk9