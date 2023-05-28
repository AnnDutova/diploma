#!/bin/bash
sleep 10;
/usr/bin/mc config host add minio http://minio-service:9000 minioadmin minioadmin;
/usr/bin/mc mb minio/backups;
/usr/bin/mc policy set public minio/backups;
exit 0;