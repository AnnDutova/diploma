cadvisor_run:
	kubectl apply -f cadvisor/daemonset.yaml
	kubectl apply -f cadvisor/service.yaml

node-exp_run:
	kubectl apply -f node-exporter/daemonset.yaml
	kubectl apply -f node-exporter/service.yaml

prometheus_run:
	kubectl apply -f prometheus/cluster-role.yaml
	kubectl apply -f prometheus/cluster-role-bind.yaml
	kubectl apply -f prometheus/configmap.yaml
	kubectl apply -f prometheus/deployment.yaml
	kubectl apply -f prometheus/service.yaml
	kubectl apply -f prometheus/secret.yaml
	kubectl apply -f prometheus/ingress.yaml

grafana_run:
	kubectl apply -f grafana/config.yaml
	kubectl apply -f grafana/deployment.yaml
	kubectl apply -f grafana/service.yaml

cadvisor-delete:
	kubectl delete -f cadvisor/daemonset.yaml
	kubectl delete -f cadvisor/service.yaml

node-exp-delete:
	kubectl delete -f node-exporter/daemonset.yaml
	kubectl delete -f node-exporter/service.yaml

prometheus-delete:
	kubectl delete -f prometheus/cluster-role.yaml
	kubectl delete -f prometheus/cluster-role-bind.yaml
	kubectl delete -f prometheus/configmap.yaml
	kubectl delete -f prometheus/deployment.yaml
	kubectl delete -f prometheus/service.yaml
	kubectl delete -f prometheus/secret.yaml
	kubectl delete -f prometheus/ingress.yaml

grafana-delete:
	kubectl delete -f grafana/config.yaml
	kubectl delete -f grafana/deployment.yaml
	kubectl delete -f grafana/service.yaml

monitor-delete: cadvisor-delete node-exp-delete prometheus-delete grafana-delete

monitor: cadvisor_run node-exp_run prometheus_run grafana_run

.PHONY: monitor