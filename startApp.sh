docker stop $(docker ps -q --filter ancestor=meant4task)

docker run -d -p 8989:8989 meant4task