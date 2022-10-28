# HTTP API Микросервис для работы с балансом пользователей
____
## Запуск
Чтобы запустить необходимо ввести команды
```
docker build -t avito-test-backend .
docker-compose up --build avito-test-backend
```
При неудачном подключении, необходимо ещё раз ввести
```
docker-compose up avito-test-backend
```
Для подключения swagger надо прописать  
```
swag init -g cmd/main.go
```
## Примеры запросов и ответов
1) Узнать баланс пользователя  
запрос:  

http://localhost:8080/balance/1  
ответ:  
{100}  
2) Перевести пользователю деньги на счёт/снять со счёта пользователя деньги (Deposit/Withdraw)  
запрос:  

http://localhost:8080/deposit  

и JSON  

необходимо указать от кого пришли деньги и кому, для учёта всех операций  
```
{
  "amount": 12,
  "fromuserid": 2,
  "orderid": 2,
  "serviceid": 2,
  "touserid": 1
}
```
ответ:  

Данные о пользователе, его id, его счёт и его дополнительный счёт  
```
{
  "id": 1,
  "amount": 10012,
  "bookamount": 0
}
```
3) Забронировать/разбронировать деньги на счету пользователя (Book/UnBook)  
запрос:  

http://localhost:8080/book  

и JSON  
```
{
  "amount": 100,
  "id": 1
}
```
ответ:  

Данные о пользователе, его id, его счёт и его дополнительный счёт  
```
{
  "id": 1,
  "amount": 9912,
  "bookamount": 100
}
```
4) Отчёт о суммах ВЫРУЧКИ для каждой услуги  
Запрос:  

http://localhost:8080/report/10/2022  

Ответ:  

ссылка на CSV файл  

...\\avito-test-backend\\reports\\reportmonth10year2022.csv"  

5) Списое транзакций для пользователя  
Запрос:  

возможно отсортировать по дате "date", по суммам в транзакциях "amount" или не сортировать ""  

http://localhost:8080/transactions/1/date  

Ответ:  

Массив в формате JSON  
```
[
  {
    "id": 1,
    "fromuserid": 1,
    "touserid": 2,
    "serviceid": 1,
    "orderid": 1,
    "amount": 100,
    "date": "2022-10-26T00:00:00Z",
    "description": "deposit"
  },
  {
    "id": 2,
    "fromuserid": 1,
    "touserid": 3,
    "serviceid": 1,
    "orderid": 1,
    "amount": 100,
    "date": "2022-10-26T00:00:00Z",
    "description": "deposit"
  }
]
```
