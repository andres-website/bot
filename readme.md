
# Инициализируем модуль
go mod init github.com/andres-website/bot

# Запуск
go run .\cmd\bot\main.go
make run

# Устанавливает модули из гитхаба
go mod tidy

# Собрать бинарник
go build ./cmd/bot/main.go

# Windows soft
GNU make
https://www.gnu.org/software/make/

Gpg4win
https://www.gpg4win.org/

// gpg.program
// https://www.gnupg.org/

# Git
git pull
git push

## Выпуск тэгов
git tag testing_calc
git push --tags

# Библиотеки
Переменные окружения
github.com/joho/godotenv

Telegram bot
https://github.com/go-telegram-bot-api/telegram-bot-api
