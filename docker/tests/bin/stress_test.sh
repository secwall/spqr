#!/bin/bash
set -ex

spqr-stress --host=spqr_router_1_1 -p 10 --dbname=db1 --usename=user1 --sslmode=disable
