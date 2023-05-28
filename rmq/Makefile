rmq-run:
	kubectl apply -f configmap.yaml
	kubectl apply -f rabbitmq-config.yaml
	kubectl apply -f sts.yaml
	kubectl apply -f service.yaml

rmq-delete:
	kubectl delete -f configmap.yaml
	kubectl delete -f rabbitmq-config.yaml
	kubectl delete -f sts.yaml
	kubectl delete -f service.yaml