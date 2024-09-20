# Logging

golang 1.22.5
go-chi
postgresql

В качестве логгера использовала slog - встроенный пакет, вышедший в версии 1.21. Более популярный, чем страндартная библиотека лог и обладающий достаточной функциональностью, чтобы не прибегать к использованию сторонних пакетов (Zerolog, Zap, Logrus, Log15, Logr и т.п.).

![main screen](image-2.png)/
Info в функции main

![db screen](image-1.png)/
Debug и Error в функции  ConnectDatabase

![applog screen](image.png)/
Файл app.log
