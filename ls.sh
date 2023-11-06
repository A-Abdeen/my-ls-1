echo Run both my-ls and the system command ls with no arguments.
ls
echo Run both my-ls and the system command ls with the arguments: "<file name>".
ls main.go
echo Run both my-ls and the system command ls with the arguments: "<directory name>".
ls pkg
echo Run both my-ls and the system command ls with the flag: "-l".
ls -l
echo Run both my-ls and the system command ls with the arguments: "-l <directory name>".
ls -l pkg
echo Run both my-ls and the system command ls with the flag: "-R", in a directory with folders in it.
ls -R
echo Run both my-ls and the system command ls with the flag: "-a".
ls -a
echo Run both my-ls and the system command ls with the flag: "-r".
ls -r
echo Run both my-ls and the system command ls with the flag: "-t".
ls -t
echo Run both my-ls and the system command ls with the flag: "-la".
ls -la
echo Run both my-ls and the system command ls with the arguments: "-l -t <directory name>".
ls -l -t pkg
echo Run both my-ls and the system command ls with the arguments: "-lRr <directory name>", in which the directory chosen contains folders.
ls -lRr pkg
echo Run both my-ls and the system command ls with the arguments: "-l <directory name> -a <file name>"
ls -l pkg -a mysplit.go
echo Run both my-ls and the system command ls with the arguments: "-lR <directory name>///<sub directory name>/// <directory name>/<sub directory name>/"
ls -lR pkg///test/// pkg/test/
echo Create directory with - name and run both my-ls and the system command ls with the arguments: 
ls -
echo Create file and link for this file and run both my-ls-1 and the system command ls with the arguments: "-l <symlink file>/"
ls -l fl/
echo Create file and link for this file and run both my-ls-1 and the system command ls with the arguments: "-l <symlink file>"
ls -l fl
echo Create directory that contains files and link for this directory and run both my-ls-1 and the system command ls with the arguments: "-l <symlink dir>/"
ls -l pk9/
echo Create directory that contains files and link for this directory and run both my-ls-1 and the system command ls with the arguments: "-l <symlink dir>"
ls -l pk9