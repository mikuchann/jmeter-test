#!/bin/sh
echo "hello world"
mkdir /test
touch /test/hello.txt
mkdir -p /test/test1
touch /test/test1/world.txt
echo "hello wrold12" > /test/hello.txt
echo "hello wraaaaaaaaold12" > /test/hello.txt
while true
do
  echo "`hostname`: `date`" >> /var/hehe.txt
  sleep 1
  ls -l /var/hehe.txt
done