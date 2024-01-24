#!/bin/bash
tar -xvf go1.21.6.linux-amd64.tar.gz
cp -rf go /usr/bin/

echo '
export GOPATH=$HOME/work
export PATH=$PATH:/usr/bin/go/bin:$GOPATH/bin
' >> ~/.profile

echo '











run the following command

source ~/.profile

'
