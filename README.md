# Golang Fiber Donation App 'DonateON'

# Приложение для сбора донатов от зрителей на трансляциях, написанное на быстром фреймворке Fiber. Реализовано 3 микросервиса, которые собраны в один монолит.
# Первый сервис - Регистрация стримера. Регистрационные данные сохраняются в БД PostgreSQL. При регистрации генерируется и подписывается JWT-Токен для обращения по API и присваивается клиенту. Клиент может получить свой токен по запросу /api/v1/token/new
POST запрос на регистрацию, с указанием полей login, email, password /api/v1/register
![Alt text](prew/register.png?raw=true "Register")
# Второй сервис реализует прием донатов стримеру от зрителя. В полях зритель указывает нужные данные о себе, стримере, вводит сумму доната, сообщение. Ссылка: /api/v1/donate
POST запрос на отправку доната, с указанием полей loginwhodonate, logintodonate, message, summary /api/v1/donate
![Alt text](prew/donate.png?raw=true "Donate")

Пример неверного POST запроса на логин стримера, который не зарегистрирован
![Alt text](prew/donatenotvalidlogin.png?raw=true "DonateNotValid")
# Третий сервис реализует обработку запросов от стримеров через API.

GET запрос на получение нового JWT-Токена для авторизации /api/v1/token/new
![Alt text](prew/new:token.png?raw=true "NewToken")

GET запрос на получение списка донатов по логину, с использованием JWT-Токена в Header /api/v1/donates/:login(например Vasily)
![Alt text](prew/donatesvalue.png?raw=true "Donates")

Пример GET запроса на получение списка донатов по логину, без использования JWT-Токена в Header 
![Alt text](prew/withouttoken.png?raw=true "WithoutToken")

Пример GET запроса на получение списка донатов по логину, c использованием JWT-Токена в Header, но неверного, либо просроченного(срок годности 24 часа) 
![Alt text](prew/notvalidtoken.png?raw=true "NotValidToken")

Также есть методы для Админа на просмотр записей всех пользователей и донатов. Работают только по безлимитному админ-токену в Header по адресу:
/api/v1/admin/users - для просмотра всех зарегистрированных пользователей
/api/v1/admin/donates - для просмотра всех донатов
/api/v1/admin/delete/donate/:id - для удаления доната по ID

Старт - `docker-compose up` потом - `go run main.go`

Если проблема с PostgreSQL, то нужно переместить create_db.sql на машину командой - `psql -U postgres postgres -h 127.0.0.1 < create_db.sql`
Либо в Docker руками перекинуть в папку и проинициализировать командой - `psql -U postgres postgres < create_db.sql`