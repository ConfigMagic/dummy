# dummy

## Установка и запуск сервера

### Зависимости
Перед началом убедитесь, что у вас установлен Go и настроено окружение для сборки Go-проектов.

### Шаг 1: Установка бинарников

Выполните следующие команды для сборки и установки бинарников `dummy` и `dummy-admin`:

1. Перейдите в директорию `deploy`:
    ```bash
    cd deploy
    ```

2. Установите необходимые инструменты:
    ```bash
    sudo make install
    ```

    После этого вы увидите:
    ```
    ✅ Installed dummy and dummy-admin to /usr/local/bin.
    ```

### Шаг 2: Запуск сервера через `dummy-admin`

После установки вы можете использовать инструмент `dummy-admin` для управления сервером. Для запуска сервера:

1. Перейдите в корневую директорию проекта:
    ```bash
    cd /path/to/your/dummy
    ```

2. Обновите зависимости Go:
    ```bash
    cd server
    go mod tidy
    cd ..
    ```

3. Запустите сервер:
    ```bash
    export MONGO_URI="mongodb://localhost:27017"
    export MONGO_DB="dummy"
    export MONGO_CONFIGS_COLLECTION="configs"
    sudo dummy-admin new_server
    ```

    В выводе будет:
    ```
    Server started on port 50051
    ```

### Возможные проблемы и их решения

1. **Ошибка "no such file or directory"**:
   Если вы видите ошибку вида:
   ```
   stat ./server/main/main.go: no such file or directory
   ```
   Это означает, что вы находитесь не в той директории. Убедитесь, что вы запускаете команду из корневой директории проекта.

2. **Ошибки с пакетами Go**:
   Если вы видите ошибки вида:
   ```
   package config_saver/internal/config is not in std
   no required module provides package github.com/gin-gonic/gin
   ```
   Выполните следующие шаги:
   ```bash
   cd server
   go mod tidy
   cd ..
   ```
   Затем попробуйте запустить сервер снова.

3. **Проблемы с правами доступа**:
   Если возникают проблемы с правами доступа, убедитесь, что:
   - Вы используете `sudo` для команд, требующих повышенных привилегий
   - У вас есть права на запись в директорию проекта

## Как работает шаблон для сервера

Для запуска серверной части dummy используется шаблон docker-compose, который лежит в `examples/server/docker-compose.tmpl`.

- Админ может редактировать этот шаблон под свои нужды.
- Разработчик указывает только переменные окружения в своём server.yaml (например, DUMMY_SERVER_URL, LOCAL_FILES).
- dummy автоматически подставляет значения из env в шаблон и генерирует итоговый docker-compose.yaml.

**Пример шаблона:**

```yaml
version: '3'
services:
  server:
    build:
      context: ../../server
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - DUMMY_SERVER_URL={{ .DUMMY_SERVER_URL }}
      - LOCAL_FILES={{ .LOCAL_FILES }}
    command: ["/app/server"]
    volumes:
      - ../../server:/app
```

**Путь к шаблону:**
- `examples/server/docker-compose.tmpl`

**Путь к пользовательскому конфигу:**
- `examples/server.yaml`

## Переменные окружения

### DUMMY_SERVER_URL

Вы можете задать адрес сервера для всех CLI-команд через переменную окружения `DUMMY_SERVER_URL`. Если не указано, по умолчанию используется `http://localhost:8080`.

Пример использования:

```bash
DUMMY_SERVER_URL="http://myserver:8080" dummy-admin push myconfig.yaml
DUMMY_SERVER_URL="http://myserver:8080" dummy sync myconfig
```

Если `DUMMY_SERVER_URL` не задана, команды используют `http://localhost:8080` по умолчанию.

### Примечания

1. **Логи сервера**: Логи сервера хранятся в `/var/log/dummy-admin/server.log`.
2. **Формат логов**: Логи пишутся в формате **JSON** — это удобно для мониторинга и анализа.
3. **Запуск сервера**: Сервер можно запустить в любой момент через `dummy-admin new_server`. Для остановки сервера завершите процесс вручную или реализуйте graceful shutdown.

Инструмент для автоматизации разработки локальных окружений.

## Примеры для разработчиков: запуск окружений через dummy

### Docker Compose (dev-docker)

1. Пример конфига: `examples/dev-docker.yaml`
2. Пример runner.yaml и шаблона: `examples/dev-docker/`

Запуск:
```bash
dummy up -c examples/dev-docker.yaml
dummy down -c examples/dev-docker.yaml
```

### Bash (dev-bash)

1. Пример конфига: `examples/dev-bash.yaml`
2. Пример runner.yaml и шаблонов: `examples/dev-bash/`

Запуск:
```bash
dummy up -c examples/dev-bash.yaml
dummy down -c examples/dev-bash.yaml
```

### Сервер использует следующие переменные окружения:

- `SERVER_PORT` - порт для запуска сервера (по умолчанию 50051)
- `MONGO_URI` - URI для подключения к MongoDB (по умолчанию mongodb://localhost:27017)
- `MONGO_DB` - имя базы данных MongoDB (по умолчанию dummy)
- `MONGO_CONFIGS_COLLECTION` - имя коллекции для конфигураций (по умолчанию configs)

### Логи

Логи сервера сохраняются в следующих местах:
- Файл: `/var/log/dummy-admin/server.log`
- Консоль: вывод в stdout

Формат логов: JSON для файла, читаемый текст для консоли.

## Новый способ запуска сервера

Теперь вы можете использовать docker-compose для запуска сервера.

Запуск:
```bash
dummy up -c examples/server.yaml
```