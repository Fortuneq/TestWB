# TestWB RESTful server with NATS-streaming
## DEV
This server uses Nats-streaming and PostgreSQL.
```
./go test -race -v main.go
```

#### Структура
- **В папке modules хранятся все модели для json**
- **В папке db всё , что касается миграциями** 
- **В папке cash всё, что отвечает за кэширование In memory** 

### WRK test benchmark

## Запуск теста из 80 горутин на 5 секунд 
```
./go-wrk -c 80 -d 5  http://localhost:8080/json
```

Output:  
```
Running 10s test @ http://192.168.1.118:8080/json  
  80 goroutine(s) running concurrently  
   142470 requests in 4.949028953s, 19.57MB read  
     Requests/sec:		28787.47  
     Transfer/sec:		3.95MB  
     Avg Req Time:		0.0347ms  
     Fastest Request:	0.0340ms  
     Slowest Request:	0.0421ms  
     Number of Errors:	0  
```
