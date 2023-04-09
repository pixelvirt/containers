# Install Info
$ docker compose up

# Mongo DB Insert
Exec into the mongodb container.

Then create one account and one admin user for the install.

```
$ mongosh
```

```
$ use alertagility;
```

```
db.accounts.insertOne{ "_id" : ObjectId("63327ed92efc60e45f5dccd4"), "uuid" : "47297812-6f41-47b0-857f-27291b2ff794", "organization" : "Admin", "subdomain" : "automation", "status" : "active", "deleted" : 0, "createdat" : ISODate("2022-01-01T13:33:55.715Z"), "updatedat" : ISODate("2022-01-01T13:33:55.715Z"), "primaryuser" : { "fullname" : "Admin", "firstname" : "Admin", "lastname" : "Super", "email" : "info@pixelvirt.com", "phone" : "", "deleted" : 0, "username" : "info@pixelvirt.com", "password" : "$2a$10$m1TdKWtwBi8qvBGiqMbzWula0DYpD6F0ljKkDgF0LjVwkskSb4kEq", "role_id" : 1, "status_id" : 1, "csscolor" : "", "subdomain" : "", "admin" : true, "primary_user" : true, "api_key" : "", "reset_token" : "", "work_days" : {  }, "work_schedule" : { "monday" : { "shiftstarttime" : 0, "shiftendtime" : 0 }, "tuesday" : { "shiftstarttime" : 0, "shiftendtime" : 0 }, "wednesday" : { "shiftstarttime" : 0, "shiftendtime" : 0 }, "thursday" : { "shiftstarttime" : 0, "shiftendtime" : 0 }, "friday" : { "shiftstarttime" : 0, "shiftendtime" : 0 }, "saturday" : { "shiftstarttime" : 0, "shiftendtime" : 0 }, "sunday" : { "shiftstarttime" : 0, "shiftendtime" : 0 } }, "notify_methods" : [ ], "notification_option" : "", "tour_shown" : false, "bot_ip" : "", "bot_port" : "" }, "billingplan" : "Trial", "billingcustomerid" : "" }
```

```
 db.users.insertOne({ "_id" : ObjectId("63ff8da59e4f6ff035600ddf"), "uuid" : "94066bf6-0172-4102-9c94-328608f82f42", "group_uuid" : "", "fullname" : "Admin Super", "firstname" : "Admin", "lastname" : "Super", "email" : "info@pixelvirt.com", "phone" : "", "deleted" : 0, "username" : "admin", "password" : "$2a$10$GsYknLdMSd6VB.aV/7jdkuVg7YKI56Snr.fuLOU8k6fkASXjRHzIa", "role_id" : 2, "status_id" : 1, "csscolor" : "", "subdomain" : "automation", "admin" : true, "primary_user" : false, "api_key" : "1cda891e-5c9c-4137-b9e6-e13b0078b242", "reset_token" : "", "account_uuid" : "47297812-6f41-47b0-857f-27291b2ff794", "work_days" : {  }, "work_schedule" : { "monday" : { "shiftstarttime" : 0, "shiftendtime" : 0 }, "tuesday" : { "shiftstarttime" : 0, "shiftendtime" : 0 }, "wednesday" : { "shiftstarttime" : 0, "shiftendtime" : 0 }, "thursday" : { "shiftstarttime" : 0, "shiftendtime" : 0 }, "friday" : { "shiftstarttime" : 0, "shiftendtime" : 0 }, "saturday" : { "shiftstarttime" : 0, "shiftendtime" : 0 }, "sunday" : { "shiftstarttime" : 0, "shiftendtime" : 0 } }, "notify_methods" : [ ], "notification_option" : "", "tour_shown" : false, "call_language" : "", "bot_ip" : "", "bot_port" : "", "bot_auth_token" : "" })
```
