#!/bin/bash

# Start MongoDB in the background
mongod --fork --logpath /var/log/mongodb.log --wiredTigerCacheSizeGB 1 --bind_ip_all

# Wait for MongoDB to start
sleep 5

# Create the application user
mongo admin -u $MONGO_INITDB_ROOT_USERNAME -p $MONGO_INITDB_ROOT_PASSWORD --eval "
db = db.getSiblingDB('$MONGO_INITDB_DATABASE');
db.createUser({user: '$MONGO_APP_USERNAME', pwd: '$MONGO_APP_PASSWORD', roles: [{role: 'readWrite', db: '$MONGO_INITDB_DATABASE'}]});
"

# Shutdown MongoDB
mongod --shutdown

# Start MongoDB normally
mongod --wiredTigerCacheSizeGB 1 --bind_ip_all
