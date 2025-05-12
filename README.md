# dummy

## Установка и запуск сервера

### Зависимости
Перед началом убедитесь, что у вас установлен Go и настроено окружение для сборки Go-проектов.

### Установка бинарников

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

## Быстрый старт: универсальный workflow

1. **Админ** готовит шаблон окружения и runner.yaml в папке (например, `examples/dev-docker/`).
2. **Админ** пушит окружение на сервер:
   ```bash
   dummy-admin push examples/dev-docker.yaml
   ```
   В архив попадут:
   - сам YAML (config.yaml)
   - все файлы из одноимённой папки (runner.yaml, шаблоны и т.д.)
3. **Разработчик** синхронизирует окружение:
   ```bash
   dummy sync dev-docker
   ```
   Всё распакуется в `~/.dummy/dev-docker/` (config.yaml, runner.yaml, шаблоны и т.д.)
4. **Разработчик** запускает окружение:
   ```bash
   dummy up -c ~/.dummy/dev-docker/config.yaml
   ```
   или, если работает из examples:
   ```bash
   dummy up -c examples/dev-docker.yaml
   ```

---

## Структура директорий и правила именования окружений

- **Пользовательский конфиг** (например, `examples/server.yaml`) должен лежать в директории `examples/`.
- Для каждого окружения создаётся папка с тем же именем, что и конфиг без расширения (например, `examples/server/`).
- В этой папке должны находиться:
  - `runner.yaml` — основной runner-конфиг для окружения
  - шаблоны (например, `docker-compose.tmpl`)
- Путь к runner.yaml формируется как: `{директория user config}/{имя окружения}/runner.yaml`.
  - Например, для `examples/server.yaml` runner ищет `examples/server/runner.yaml`.
- Все шаблоны для окружения также должны лежать в этой папке (например, `examples/server/docker-compose.tmpl`).

**Пример структуры:**

```
examples/
  server.yaml
  server/
    runner.yaml
    docker-compose.tmpl
  dev-docker.yaml
  dev-docker/
    runner.yaml
    docker-compose.tmpl
  dev-bash.yaml
  dev-bash/
    runner.yaml
    start.sh.tmpl
    stop.sh.tmpl
```

> **Важно:**
> - Имя папки окружения всегда совпадает с именем user config без `.yaml`.
> - Все файлы runner и шаблоны должны лежать внутри этой папки.
> - Если структура не соблюдена, запуск окружения завершится ошибкой поиска runner.yaml или шаблонов.