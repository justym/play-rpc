Check to run this
``` sh
curl -X POST \
   http://localhost:8000/rpc \
   -H 'cache-control: no-cache' \
   -H 'content-type: application/json' \
   -d '{
   "method": "JSONServer.Call",
   "params": [{
   "ID": "0001"
   }],
   "id": "1"
}'
```
