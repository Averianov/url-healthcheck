# url-healthcheck

Программный комплекс состоящий из утилиты-диспетчера проверяющего доступность url согласно полученного списка и grpc интерфейса предоставляющего данные согласно интерфеса описанного в proto файле.

Изменения json файла с данными проверяемых url возможно без остановки диспетчера.

Для применения изменений основных параметров (через переменные окружения) требуется перезапуск процессов.

С учетом того, что в роли api была выбрана технология grpc, то названия типов проверок и их статусов были приведены к требованиям интерфеса proto файла.

```
  enum CheckType {
    CHECK_TYPE_UNSPECIFIED = 0;
    CHECK_TYPE_STATUS_CODE = 1;
    CHECK_TYPE_TEXT = 2;
  }
  enum CheckStatus {
    CHECK_STATUS_UNSPECIFIED = 0;
    CHECK_STATUS_OK = 1;
    CHECK_STATUS_FAIL = 2;
  }
```

### Общие требования для развертывания

подготовить .env файл, либо в запускаемой системе установить переменные окружения согласно образца

> DB_HOST=localhost

> DB_PORT=3306

> DB_SCHEMA=work

> DB_USER=checker

> DB_PASSWORD=checker

> DB_DROP=false

> GRPC_PORT=443

> HCK_DURATION=10

Скачать репозиторий и перейти в него
```
gh repo clone Averianov/url-healthcheck
cd ./url-healthckeck
git checkout dev
git pull
```

Перевести проект в режим GO111MODULE
```
go mod init
```

### Локальное развертывание

Для локального запуска предполагается, что база данных развернута, создана требуемая схема и параметры доступа к ней известны

произвести сборку исполняемых вайлов
```
make compile
```
для запуска исполняемых вайлов выполнить команды
```
./urlapi > logapi.log 2>&1 &
./urldisp > logdisp.log 2>&1 &
```

### Развертывание приложений в контейнере

произвести развертывание контейнеров
```
make local-up
```
остановка контейнеров
```
make local-down
```

### Образец запроса результатов проверок через api (grpc) средствами GRPCURL

```
grpcurl -plaintext localhost:443 com.url.healthcheck.Info/Checks
```
