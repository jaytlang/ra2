#!/bin/sh

# Helper script to run a web server showing
# how far along I am or am not in making this
# crazy recitation assignment thing

$SRVDIR = /mnt/c/Users/Jay/Documents/quarantine

go build
./ra2
cp data/out.csv $SRVDIR
cp data/stats.txt $SRVDIR

cat > $SRVDIR/status.txt <<EOF 

****FINAL CODE REVIEW****

Expect readiness by 9PM Eastern
out.csv temporarily frozen for manual
review, this might be the final assignment

=============

-> initial assignments: DONE
-> modularized import/export: DONE
-> statistics strategy: DONE
-> FIX BROKEN TUTORIAL CAPACITY: DONE
-> cap recitation capacity via rtunions: DONE
-> better teaming algorithm: DONE

-> better configuration: NOT YET
-> better documentation: NOT YET
-> final source code release: NOT YET

EOF


