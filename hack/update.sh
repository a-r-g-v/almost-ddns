#!/bin/bash

ipaddr=$1

# clone nsd-zonefiles
git clone git@github.com:argvc/nsd-zonefiles.git ../tmp/nsd-zonefiles

# update zone-files
cd ../tmp/nsd-zonefiles/util
./update.sh $ipaddr

# commit and push
git add -A
git commit -m "Update srv001 by system-a"
git push origin master # call circle-ci for deploy

cd ../../
rm -Rf ./nsd-zonefiles


