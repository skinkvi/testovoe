# testovoe

копия файла .env с моими настройками 
`POSTGRES_CONN=postgres://postgres:postgres@db:5432/testovoe_db
POSTGRES_JDBC_URL=jdbc:postgresql://db:5432/testovoe_db
POSTGRES_USERNAME=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_HOST=db
POSTGRES_PORT=5432
POSTGRES_DATABASE=testovoe_db

SERVER_ADDRESS=0.0.0.0:8080

LOG_LEVEL=info
`

так же я не успел написать тесты но протестировал все запросы в ручную все работает


Проект запускается командой `docker-compose up` делать это из папки `\zadanie-6105>` 
так же что бы проверить работоспособность предоставляю свои запросы из postmen, тестировал там, при желании можно через curl 

`POST
NewTender
http://localhost:8080/api/tenders/new
Body
raw (json)
json
{
    "name": "New Tender",
    "description": "This is a new tender",
    "creator_id": 1,
    "organization_id": 1
}`


`GET
getALLtenders
http://localhost:8080/api/tenders`


`POST
createUser
http://localhost:8080/api/users/new
Body
raw (json)
json
{
    "username": "user1",
    "first_name": "Dima",
    "last_name": "Varauyn"
}`

`POST
newOrganization
http://localhost:8080/api/organizations/new
Body
raw (json)
json
{
    "name": "Organization Name",
    "description": "Organization Description",
    "type": "Organization Type"
}`
`
GET
getTendersByUsername
http://localhost:8080/api/tenders/my?username=user1
Query Params
username
user1`

`PATCH
editTender
http://localhost:8080/api/tenders/1/edit
Body
raw (json)
json
{
    "name": "Updated Tender",
    "description": "This is an updated tender",
    "service_type": "Updated Service Type",
    "organization_id": 1,
    "creator_id": 1,
    "version": 2
}`


`POST
bidNew
http://localhost:8080/api/bids/new
Body
raw (json)
json
{
    "name": "New Bid",
    "description": "This is a new bid",
    "status": "active",
    "tender_id": 1,
    "organization_id": 1,
    "creator_username": "user1",
    "version": 1
}`

`GET
getBid
http://localhost:8080/api/bids/my?username=user1
Query Params
username
user1`

`
GET
getBidByID
http://localhost:8080/api/bids/tender/1/list`


`PATCH
editBid
http://localhost:8080/api/bids/1/edit
Body
raw (json)
json
{
    "name": "Updated Bid",
    "description": "This is an updated bid",
    "status": "active",
    "tender_id": 1,
    "organization_id": 1,
    "creator_username": "user1",
    "version": 2
}`


`PUT
rollabackBid
http://localhost:8080/api/bids/1/rollback/1`

`POST
newDecision
http://localhost:8080/api/decisions/new
Body
raw (json)
json
{
    "bid_id": 1,
    "organization_id": 1,
    "decision": "Approved"
}`

`GET
bidApprove
http://localhost:8080/api/bids/approve/1`

`POST
newReviews
http://localhost:8080/api/reviews/new
Body
raw (json)
json
{
    "bid_id": 1,
    "organization_id": 1,
    "review": "This is a review"
}
GET
getReview
http://localhost:8080/api/bids/1/reviews
`




