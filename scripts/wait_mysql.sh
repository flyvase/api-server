until mysqladmin ping -h mysql --silent; do
  echo 'waiting for mysqld to be connectable...'
  sleep 2
done

exec go run src/main.go