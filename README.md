# DiskInfo

[English](#english) | [Русский](#russian)

# <a name="english"></a>English

A command-line utility written in Go that provides detailed disk usage information for specified directories.

## Features

- Display total disk space, used space, and free space
- Show disk usage percentage
- Recursively calculate size of subdirectories
- Support for custom path input
- Error handling for non-existent or inaccessible paths

## Installation

```bash
git clone https://github.com/yourusername/DiskInfo.git
cd DiskInfo
go build 
```

## Usage

Run without arguments (defaults to /home directory):
```bash
./DiskInfo 
```

Specify a custom path:
```bash
./DiskInfo /path/to/directory 
```

## Example Output

```
Disk usage of /home:
Total: 500 GB
Used: 250 GB (50.00%)
Free: 250 GB

Directory: /home/user/Documents
Size: 15.50 GB

Directory: /home/user/Downloads
Size: 25.75 GB
```

## Requirements

- Go 1.11 or higher
- Linux/Unix-based operating system



# <a name="russian"></a>Русский

Консольная утилита, написанная на Go, предоставляющая подробную информацию об использовании дискового пространства для указанных директорий.

## Возможности

- Отображение общего, используемого и свободного места на диске
- Показ процента использования диска
- Рекурсивный подсчет размера поддиректорий
- Поддержка пользовательских путей
- Обработка ошибок для несуществующих или недоступных путей

## Установка

```bash
git clone https://github.com/yourusername/DiskInfo.git
cd DiskInfo
go build
```

## Использование

Запуск без аргументов (по умолчанию для директории /home):
```bash
./DiskInfo
```

Указание пользовательского пути:
```bash
./DiskInfo /путь/к/директории
```

## Пример вывода

```
Использование диска для /home:
Всего: 500 ГБ
Использовано: 250 ГБ (50.00%)
Свободно: 250 ГБ
## License

MIT License
Директория: /home/user/Documents
Размер: 15.50 ГБ

Директория: /home/user/Downloads
Размер: 25.75 ГБ
```

## Требования

- Go 1.11 или выше
- Операционная система на базе Linux/Unix

