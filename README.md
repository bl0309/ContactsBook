## ContactsBook


Base URL = http://localhost:8080

Contact Book

Add Contact
fields: id, name, email, phone, address
List Contact 2.1 Pagination 2.2 Filters ( name, phone)
Get Contact by id
Bulk Delete Contact
Update Specific Contact
Delete Specific Contact
Endpoints documents

GET /contacts Get all contacts /contacts?page=1&size=20 /contacts?page=1&size=20&name="ram"&phone="909090"
POST /contacts Create new contact data = { id : 1, name: "ram", email :"gopal", phone :"9898"}
DELETE /contacts
GET /contacts/{id} Get detail of contact with id
PUT /contacts/{id} UPdate detail of contact with id
DELETE /contacts/{id} Delete contact with id
Users

/// Contact with users associated /auth/login /auth/register /users?page=1&size=20 /users/{id} /users/{id}/contacts?page=1&size=20 /users/{id}/contacts/{id}

GET /contacts Get all contacts /contacts?page=1&size=20 /contacts?page=1&size=20&name="ram"&phone="909090"
POST /contacts Create new contact data = { id : 1, name: "ram", email :"gopal", phone :"9898", user_id:1}
DELETE /contacts
GET /contacts/{id} Get detail of contact with id
PUT /contacts/{id} UPdate detail of contact with id
DELETE /contacts/{id} Delete contact with id