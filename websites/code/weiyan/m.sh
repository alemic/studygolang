#!/bin/sh 
#===================== 
#ni da ye
#http://pingliwang.com:9090/
#===================== 
while : 
do 
echo "Current DIR is " $PWD 
stillRunning=$(ps -ef |grep "$PWD/bin/studygolang.exe" |grep -v "grep") 
if [ "$stillRunning" ] ; then 
echo "my service was already started by another way" 
echo "Kill it and then startup by this shell, other wise this shell will loop out this message annoyingly" 
kill -9 $pidof $PWD/bin/studygolang.exe
else 
echo "my service was not started" 
echo "Starting service ..." 
$PWD/bin/studygolang.exe
echo "my service was exited!" 
fi 
sleep 1
done 