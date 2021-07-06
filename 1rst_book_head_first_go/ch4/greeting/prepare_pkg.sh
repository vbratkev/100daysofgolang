ln -s "$(pwd)/greeting/" $GOPATH/src/greeting
go mod init greeting


# Example
#
# root -> go mod init xxx
#  |- main.go -> package "main"
#  |- tools
#       |- helper.go -> package "tools"
# and you want to import tools package from main.go you need to import this "xxx/tools"
