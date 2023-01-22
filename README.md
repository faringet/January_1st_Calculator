# January_1st_Calculator
С помощью этого веб-приложения можно узнать сколько дней осталось или прошло до 1-го января введенного года.
___
# Route
Маршрут имеет вид - `http://localhost:3000/when/введенный параметр`

Напрмиер при вводе:

• `http://localhost:3000/when/2030`, output - `Days left: 2535`

• `http://localhost:3000/when/2007`, output - `Days gone: 5865`


___
# Header Checker
Checker проверяет наличие HTTP-заголовка. При заголовке `X-PING` со значением `ping`, в ответе сервиса будет заголовок ответа `X-PONG` со значением `pong`:
![](https://github.com/faringet/January_1st_Calculator/blob/master/Header%20Checker.png)

___
### На данный момент:
- сервис принимает от юзера по маршруту параметр года и выдает рассчеты
- обработка запросов с помощю [**Gin Web Framework**](https://gin-gonic.com/docs/)
- логирование [**Logrus'ом**](https://github.com/sirupsen/logrus) 
- запуск приложения через Makefile (`run` - запускает приложение, `build` - компилирует приложение)
