PWD = $(shell pwd)

master: configmap
	kubectl apply -f deployment/sts/sts-master.yaml
	kubectl apply -f deployment/service/service-master.yaml

delete-master:
	kubectl delete statefulset.apps/postgres-db
	kubectl delete service/db-service

slave:
	kubectl apply -f deployment/sts/sts-replica.yaml
	kubectl apply -f deployment/service/service-replica.yaml

delete-slave:
	kubectl delete statefulset.apps/postgres-db-slave
	kubectl delete service/db-service-replica

configmap:
	kubectl create configmap init-config --from-file=db-init-scripts/common/ --from-file=db-init-scripts/single/
	kubectl apply -f deployment/config/configmap.yaml
	kubectl apply -f deployment/config/master-config.yaml
	kubectl apply -f deployment/config/slave-config.yaml

delete-conf:
	kubectl delete configmap init-config
	kubectl delete configmap  db-config
	kubectl delete configmap master-config
	kubectl delete configmap slave-config

patroni-run:
	docker build -t patroni -f patroni/Dockerfile.patroni .
	kubectl create configmap init-config --from-file=db-init-scripts/common
	kubectl apply -f patroni/service-acc.yaml
	kubectl apply -f patroni/secret.yaml
	kubectl apply -f patroni/sts.yaml
	kubectl apply -f patroni/cluster-role.yaml
	kubectl apply -f patroni/cluster-role-bind.yaml
	kubectl apply -f patroni/role.yaml
	kubectl apply -f patroni/role-bind.yaml
	kubectl apply -f patroni/service.yaml
	kubectl apply -f patroni/service-repl.yaml

patroni-delete:
	kubectl delete configmap init-config
	kubectl delete secret/patroni-secret
	kubectl delete statefulset.apps/patroni-psql
	kubectl delete service/patroni-psql
	kubectl delete service/patroni-psql-repl
	kubectl delete role patroni-psql
	kubectl delete rolebinding patroni-psql

consul-run:
	docker build -t consul consul/
	kubectl apply -f consul/service.yaml
	kubectl apply -f consul/sts.yaml

consul-delete:
	kubectl delete statefulset.apps/consul
	kubectl delete service/consul


