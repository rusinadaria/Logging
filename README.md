# Logging

golang 1.22.5  
go-chi  
postgresql

В качестве логгера использовала slog - встроенный пакет, вышедший в версии 1.21. Более популярный, чем страндартная библиотека log и обладающий достаточной функциональностью, чтобы не прибегать к использованию сторонних пакетов (Zerolog, Zap, Logrus, Log15, Logr и т.п.).

![alt text](<img/main_screen.png>)  
Info в функции main

![db screen](<img/db_screen.png>)  
Debug и Error в функции  ConnectDatabase

![applog screen](<img/app_log_screen.png>)  
Файл app.log
