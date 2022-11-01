# Avito_internship_Go

# Задание
Необходимо реализовать микросервис для работы с балансом пользователей (зачисление средств, списание средств, перевод средств от пользователя к пользователю, а также метод получения баланса пользователя). Сервис должен предоставлять HTTP API и принимать/отдавать запросы/ответы в формате JSON.

## Запуск проекта
После загрузки проекта а своё устроиство, нужно перейти в главную деррикторию проекта (Avito_internship_Go) и ввести пару команд  
Сначала скачаем использующиеся библиотеки  
go get github.com/lib/pq   
go get github.com/julienschmidt/httprouter  
Далее нужно запустить скрипт создания бд (scripts/database)
Я делал это через pgadmin
И запускаем проект(из рабочей дериктории)  
go run .\cmd\api\.  
чтобы остановить нажмите "cntrl" + "c" в командной строке

## Структура бд

Account(Содержит id, доступные и зарезервированные средства)  
Report(информация об уже предоставленных услугах(для бугалтерии))  
Service(id, название и цена услуги)  
Reserv(хранит информацию о зарезервированном сервисе)  
Operation(операции помиму услуг: пополнение,снятие, перевод)  
Transaction_history(хранит информацию о перевода зачислениях снятиях)  

## Запросы
Я использовал postman

#### Service
Создание сервиса (post)
http://localhost:4000/service/create/:name/:price  
Получение сервиса (get)
http://localhost:4000/service/get/:id  

#### Reserv 

Получение записи о резервации денег (get)  
http://localhost:4000/reserv/get/:id  

#### Account  

Создание пустого аккаунта(cash = reserv_cash = 0) (post)  
http://localhost:4000/account/create  
Получение аккаунта (get)  
http://localhost:4000/account/getId/:id  
Пополнение средсв аккаунта (put)  
http://localhost:4000/account/add/:id/:account_cash  
Снятие средсв аккаунта (put)  
http://localhost:4000/account/withdrawal/:id/:account_cash  
Перевод средсв с одного на другой аккаунт (put)  
http://localhost:4000/account/transfer/:id/:ToId/:account_cash  

Резервирование средсв для преобретение сервиса (put)  
http://localhost:4000/account/reserv/:id_account/:id_service  
Средвства из cash записываются в reserved_cash(если средств хватает), также делается запись в reserv о резервировании  

#### Transaction 

Получение информации о операция пользоателя (переводы/пополнения/снятия) (get)  
http://localhost:4000/transaction/get/:id  

#### Report
Создает отчет по reserv, средства с reserved_cash  снимаются, запись о резервации из reserv удаляется, создaётся запись в report (put)  
http://localhost:4000/report/create/:id  
Получение информации о отчёте (get)  
http://localhost:4000/report/getID/:id  

#### Доп задание 1
Получение отчета по году и месецу , считается сумарная выручка за конкретный сервис(get)    
http://localhost:4000/report/get/:year/:month  
Используется для чтения файла , путь к файлу возвращается при создании репорта по месяцу и году  
http://localhost:4000/file/:filename  

#### Доп задание 2
Получение информации о транзакциях по user_id(пополнение/снятие/перевод)(get)  
page = номер страницы  
page_size = размеру  
sort = date/-date  sum/-sum(убывание)  
http://localhost:4000/history/transactions/:id?page&page_size&sort

Получение информации о покупки сервисов по user_id(get)  
Получение информации из report по user_id(get)  
page = номеру страницы  
page_size = размеру  
sort = date/-date  sum/-sum(убывание)  
http://localhost:4000/history/user/:id?page&page_size&sort
