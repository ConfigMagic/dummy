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

1. Запустите `dummy-admin`:
    ```bash
    dummy-admin
    ```

    Вы увидите примерно такой вывод:
    ```
    ... длинное описание ...

    Usage:
      dummy-admin [command]

    Available Commands:
      completion  Generate shell completion script
      help        Help about any command
      new_server  Create the dummy's server
      push        Publish configuration to the server
      users       Manage users

    Flags:
      -h, --help   help for dummy-admin

    Use "dummy-admin [command] --help" for more информации о команде.
    ```

2. Для запуска сервера используйте команду:
    ```bash
    sudo dummy-admin new_server
    ```

    В выводе будет:
    ```
    Server started on port 50051
    ```

## Быстрый старт: запуск сервера через dummy

Теперь вы можете поднять сервер с помощью утилиты `dummy` и минимального конфига. Это удобно для локального запуска сервера dummy в docker-контейнере.

1. Перейдите в корень репозитория:
    ```bash
    cd /path/to/your/dummy
    ```
2. Запустите сервер через dummy:
    ```bash
    dummy up -c examples/server.yaml
    ```

Это поднимет сервер в контейнере с проброшенным портом 8080. Конфиг для запуска находится в `examples/server.yaml`.

## Быстрый старт: минимальный вариант

Для запуска серверной части достаточно указать только нужные переменные окружения:

```bash
export DUMMY_SERVER_URL="http://localhost:8080"
export LOCAL_FILES="./local_data"
# Запуск
 dummy up -c examples/server.yaml
```

Всё остальное (пути, шаблоны, порты, docker-compose) dummy подхватит из дефолтных настроек внутри репозитория.

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