create-kafka-topic:
	docker exec -it e67baab4d667 kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic obu_reads

list-kafka-topics:
	docker exec -it e67baab4d667 kafka-topics.sh --list --bootstrap-server localhost:9092