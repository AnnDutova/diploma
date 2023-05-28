cronjob:
	kubectl apply -f cronjob.yaml

minio-run:
	docker build -t minio -f minio/Dockerfile.minio .
	kubectl apply -f minio/config.yaml
	kubectl apply -f minio/secret.yaml
	kubectl apply -f minio/sts.yaml
	kubectl apply -f minio/service.yaml
	kubectl apply -f minio-client/client-job.yaml

store:
	kubectl apply -f storage/pv.yaml
	kubectl apply -f storage/pvc.yaml

backup-run:
	docker build -t backup -f service/Dockerfile.backup .
	kubectl apply -f service/deployment.yaml
	kubectl apply -f service/service.yaml

cron-run:
	docker build -t kubectl_curl -f minio-client/Dockerfile.kubectl .
	kubectl apply -f minio-client/cronjob-role.yaml
	kubectl apply -f minio-client/backup-cronjob.yaml


