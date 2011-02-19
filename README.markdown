Description
===========

A simple application to dump multicast DNS packages on the local lan

A note on Darwin
----------------

If you are trying to run this code on OS X, you will need to patch your Go installation before hand

    patch -p0 $GOROOT/src/pkg/net/sock.go < sock.patch 

then rebuild your Go installation

    cd $GOROOT/src && ./all.bash

Installation
============

_mdns-test_ uses Miek G's godns library http://miek.nl/projects/godns/ which needs to be installed before building

    git clone git://github.com/davecheney/mdns-test.git
    cd mdns-test
    goinstall .
    make
    ./mdns-test

