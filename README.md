# Gin Blog System (Updating)

## Overview
A blog system based on [Gin](https://github.com/gin-gonic/gin).

## Features Completed
- Article release and modification
- Article tags


## TODO List
- [x] Upload image and files
- [ ] API Access Control: Implement authentication to secure the API endpoints.
- [ ] Middleware Implementation: Integrate common middleware for enhanced logging and error handling.
- [ ] Distributed Tracing


### Tag Management
1. **Create Tag**
    ```bash
    curl -X POST http://127.0.0.1:8000/api/v1/tags -F 'name=Go' -F created_by=ebbi
    ```

2. **Update Tag**
    ```bash
    curl -X PUT http://127.0.0.1:8000/api/v1/tags/1 -F name=cpp -F modified_by=ebbi
    ```

3. **Delete Tag**
    ```bash
    curl -X DELETE http://127.0.0.1:8000/api/v1/tags/1
    ```

4. **Get Tag List**
    ```bash
    curl -X GET 'http://127.0.0.1:8000/api/v1/tags?page=1&page_size=2'
    ```

### Article Management
1. **Create Article**
    ```bash
    curl -X POST http://127.0.0.1:8000/api/v1/articles -F 'title=Go Tutorial' -F 'desc=go test 1 desc' -F 'content=go test 1 content' -F created_by=ebbi
    ```

2. **Add Tag to Article**
    ```bash
    curl -X POST http://127.0.0.1:8000/api/v1/articles/:id/tags -F tag_id=3 -F created_by=ebbi
    ```

3. **Remove Tag from Article**
    ```bash
    curl -X DELETE http://127.0.0.1:8000/api/v1/articles/:id/tags -F tag_id=3
    ```

4. **Get Article List by Tag**
    ```bash
    curl -X GET 'http://127.0.0.1:8000/api/v1/articles?page=1&page_size=2' -F tag_id=3
    ```

5. **Get Article by ID**
    ```bash
    curl -X GET http://127.0.0.1:8000/api/v1/articles/:id
    ```

6. **Update Article**
    ```bash
    curl -X PUT http://127.0.0.1:8000/api/v1/articles/:id -F 'title=Go Tutorial 1' -F modified_by=ebbi
    ```

7. **Delete Article**
    ```bash
    curl -X DELETE http://127.0.0.1:8000/api/v1/articles/:id
    ```

### Update Image

```bash 
curl -X POST http://127.0.0.1:8000/upload/file -F file=@{file_path} -F type=1
```