curl -X GET http://localhost:8000/person?id=0
curl -X GET http://localhost:8000/person?id=1
curl -X GET http://localhost:8000/person?id=123
curl -d '{"id":123, "name":"name123", "email":"email123@email.com", "gender":"female", "pregnant":"pregnant_yes", "child":"child_yes", "childgender":"child_male", "childage":"child_5", "country":"country123"}' -H "Content-Type: application/json" -X POST http://localhost:8000/addperson
curl -X GET http://localhost:8000/person?id=123
