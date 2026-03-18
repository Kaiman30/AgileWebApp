golang Agile Webapp

# Запустить PostgreSQL

`make env-up`

# Запустить проброс порта

`make env-port-forward`

# Проверить, что контейнер жив

`docker ps | grep agilewebapp-env-postgres`

# Должен быть статус "Up"

# Проверить, что порт слушается

`ss -tlnp | grep 5432`

# Должно быть: 127.0.0.1:5432 LISTEN

# Запустить приложение

`make agilewebapp-run`
