# Gin Blog System

5/19: tag management complete
- create tag 
```bash 
curl -X POST http://127.0.0.1:8000/api/v1/tags -F 'name=Go' -F created_by=ebbi 
```
- update tag 
```bash
curl -X PUT http://127.0.0.1:8000/api/v1/tags/1 -F name=cpp -F modified_by=ebbi
```
- delete tag 
```bash 
curl -X DELETE http://127.0.0.1:8000/api/v1/tags/1
```
- get tag list 
```bash
curl -X GET 'http://127.0.0.1:8000/api/v1/tags?page=1&page_size=2'
```