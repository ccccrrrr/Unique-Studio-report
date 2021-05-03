# typeorm

## create connection

```typescript
import "reflect-metadata";
import { createConnection } from "typeorm";
import { Photo } from "./entity/Photo";

createConnection({
    type: "mysql",
    host: "localhost",
    port: 3306,
    username: "root",
    password: "admin",
    database: "test",
    entities: [Photo],
    synchronize: true,
    logging: false,
})
    .then((connection) => {
        // 这里可以写实体操作相关的代码
    })
    .catch((error) => console.log(error));

```

use table

```typescript
let photoRepository = connection.getRepository(Photo);
```



## create

```
await photoRepository.save(photo);
```



## delete

```
let photoToRemove = await photoRepository.findOne(1);
await photoRepository.remove(photoToRemove);
```



## retrieve

```typescript
let savedPhotos = await photoRepository.find();
let savedPhotos = await photoRepository.findOne(1);
        let meAndBearsPhoto = await photoRepository.findOne({ name: "Me and Bears" });
```



## update



## 自定义orm(typescript)

driver : four basic operations of database

connection : open mysql

decoration(entity, column)

### open sql

```typescript
createConnection({
    type: "mysql",
    host: "localhost",
    port: 3306,
    username: "root",
    password: "admin",
    database: "test",
    entities: [Photo],
    synchronize: true,
    logging: false,
})
    .then((connection) => {
        // 这里可以写实体操作相关的代码
    })
    .catch((error) => console.log(error));
```





