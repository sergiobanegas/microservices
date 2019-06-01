## Microservices (WIP)

This project contains multiple microservices that communicate between each other using Consul and GRPC.

### Services:

|  Service | Description  |  Endpoints |  Language |
|---|---|---|---|
|  Auth | Authentication | POST /login | GO |
| Product  | CRUD operations with products | - GET /products: get list of products <br/>- GET /products/{id}: get specific product | GO  |
| Cart  | The user can add and remove products to the cart | POST /cart: add product to the cart<br/> PUT /cart: modify cart product<br/> PUT /cart/delete: delete cart product <br/>DELETE /cart: clear cart | GO  |
| Checkout  |  Finish the purchasing process | POST /checkout | GO |
| Payment  | Generate a transaction id  | Not exposed | Java |

### Requirements:
- GO
- Java
- Maven
- Google Protocol Buffers
- Consul running on port 8500
- Redis server(port 6379)
- MySQL server(port 3306) with user=admin, password=root and a database called 'microservices'

### How to run
```bash
cd services/{service_name}
make proto
make install
make run -B
```

### TODO:
- Add Open API specification
- Add filters to product search
- Add load balancer
- Add api gateway
- Add unit tests, integration tests and E2E tests
- Dockerize the application
- Add configuration server
- Add authorization filter


#### Application made by Sergio Banegas Cortijo
