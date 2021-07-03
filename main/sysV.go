package main

var systemVConfig = `#! /bin/sh
#
#       /etc/rc.d/init.d/{{.Name}}
#
#       Starts {{.Name}} as a daemon
#
# chkconfig: 2345 87 17
# description: Starts and stops a single {{.Name}} instance on this system
### BEGIN INIT INFO
# Provides: {{.Name}} 
# Required-Start: $network $named
# Required-Stop: $network $named
# Default-Start: 2 3 4 5
# Default-Stop: 0 1 6
# Short-Description: This service manages the {{.Description}}.
# Description: {{.Description}}
### END INIT INFO
#
# Source function library.
#
if [ -f /etc/rc.d/init.d/functions ]; then
    . /etc/rc.d/init.d/functions
fi
exec="{{.Path}}"
servname="{{.Description}}"
proc="{{.Name}}"
pidfile="/var/run/$proc.pid"
lockfile="/var/lock/subsys/$proc"
stdoutlog="/var/log/$proc.log"
stderrlog="/var/log/$proc.err"
[ -d $(dirname $lockfile) ] || mkdir -p $(dirname $lockfile)
[ -e /etc/sysconfig/$proc ] && . /etc/sysconfig/$proc
start() {
    [ -x $exec ] || exit 5
    if [ -f $pidfile ]; then
        if ! [ -d "/proc/$(cat $pidfile)" ]; then
            rm $pidfile
            if [ -f $lockfile ]; then
                rm $lockfile
            fi
        fi
    fi
    if ! [ -f $pidfile ]; then
        printf "Starting $servname:\t"
        echo "$(date)" >> $stdoutlog
        $exec {{.Args}} >> $stdoutlog 2>> $stderrlog &
        echo $! > $pidfile
        touch $lockfile
        success
        echo
    else
        # failure
        echo
        printf "$pidfile still exists...\n"
        exit 7
    fi
}
stop() {
    echo -n $"Stopping $servname: "
    killproc -p $pidfile $proc
    retval=$?
    echo
    [ $retval -eq 0 ] && rm -f $lockfile
    return $retval
}
restart() {
    stop
    start
}
rh_status() {
    status -p $pidfile $proc
}
rh_status_q() {
    rh_status >/dev/null 2>&1
}
case "$1" in
    start)
        rh_status_q && exit 0
        $1
        ;;
    stop)
        rh_status_q || exit 0
        $1
        ;;
    restart)
        $1
        ;;
    status)
        rh_status
        ;;
    *)
        echo $"Usage: $0 {start|stop|status|restart}"
        exit 2
esac
exit $?
`
