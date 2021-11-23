# chat_backend

Chat backend which serves REST APIs

# Run server

    1. make build //build the go modules

    2. make run //run the go server

Once the server starts running, go to the postman, give the api endpoints in postman.

# Postman

    1. Create a new user api endpoint,

        localhost:8080/users -- use post method for this api

        //response

        {
            "ID": 7,
            "created_at": "2021-11-23T09:25:44.590575+05:30",
            "updated_at": "2021-11-23T09:25:44.590575+05:30",
            "name": "testuser7"
        }

    2. Create a new chatroom,

        localhost:8080/chatrooms -- post

        //response

        {
            "ID": 4,
            "created_at": "2021-11-23T09:29:03.308708+05:30",
            "updated_at": "2021-11-23T09:29:03.308708+05:30",
            "room_name": "testchatroom4",
            "users": [
                {
                    "ID": 1,
                    "created_at": "2021-11-22T20:38:45.7013+05:30",
                    "updated_at": "2021-11-22T20:38:45.7013+05:30",
                    "name": "testuser1"
                },
                {
                    "ID": 2,
                    "created_at": "2021-11-22T20:38:59.696595+05:30",
                    "updated_at": "2021-11-22T20:38:59.696595+05:30",
                    "name": "testuser2"
                }
            ]

        }

    3. List all the chatrooms,

        localhost:8080/chatrooms -- get method

        //response

                [
            {
                "ID": 1,
                "created_at": "2021-11-22T20:44:57.772457+05:30",
                "updated_at": "2021-11-22T20:44:57.772457+05:30",
                "room_name": "testchatroom1",
                "users": [
                    {
                        "ID": 1,
                        "created_at": "2021-11-22T20:38:45.7013+05:30",
                        "updated_at": "2021-11-22T20:38:45.7013+05:30",
                        "name": "testuser1"
                    },
                    {
                        "ID": 2,
                        "created_at": "2021-11-22T20:38:59.696595+05:30",
                        "updated_at": "2021-11-22T20:38:59.696595+05:30",
                        "name": "testuser2"
                    },
                    {
                        "ID": 4,
                        "created_at": "2021-11-22T20:39:17.665711+05:30",
                        "updated_at": "2021-11-22T20:39:17.665711+05:30",
                        "name": "testuser4"
                    }
                ]
            },
            {
                "ID": 4,
                "created_at": "2021-11-23T09:29:03.308708+05:30",
                "updated_at": "2021-11-23T09:29:03.308708+05:30",
                "room_name": "testchatroom4",
                "users": [
                    {
                        "ID": 1,
                        "created_at": "2021-11-22T20:38:45.7013+05:30",
                        "updated_at": "2021-11-22T20:38:45.7013+05:30",
                        "name": "testuser1"
                    },
                    {
                        "ID": 2,
                        "created_at": "2021-11-22T20:38:59.696595+05:30",
                        "updated_at": "2021-11-22T20:38:59.696595+05:30",
                        "name": "testuser2"
                    }
                ]
            }
        ]


    4. Get a single record chatroom,

        localhost:8080/chatroom/1 -- get method

        // response

            {
                "ID": 1,
                "created_at": "2021-11-22T20:44:57.772457+05:30",
                "updated_at": "2021-11-22T20:44:57.772457+05:30",
                "room_name": "testchatroom1",
                "users": [
                    {
                        "ID": 1,
                        "created_at": "2021-11-22T20:38:45.7013+05:30",
                        "updated_at": "2021-11-22T20:38:45.7013+05:30",
                        "name": "testuser1"
                    },
                    {
                        "ID": 2,
                        "created_at": "2021-11-22T20:38:59.696595+05:30",
                        "updated_at": "2021-11-22T20:38:59.696595+05:30",
                        "name": "testuser2"
                    },
                    {
                        "ID": 4,
                        "created_at": "2021-11-22T20:39:17.665711+05:30",
                        "updated_at": "2021-11-22T20:39:17.665711+05:30",
                        "name": "testuser4"
                    }
                ]
            }

    5. Update a chatroom record,

        localhost:8080/chatroom/4 -- put method

        //response

        {
            "ID": 4,
            "created_at": "2021-11-23T09:29:03.308708+05:30",
            "updated_at": "2021-11-23T09:35:13.765902+05:30",
            "room_name": "roamchatroom",
            "users": [
                {
                    "ID": 1,
                    "created_at": "2021-11-22T20:38:45.7013+05:30",
                    "updated_at": "2021-11-22T20:38:45.7013+05:30",
                    "name": "testuser1"
                },
                {
                    "ID": 2,
                    "created_at": "2021-11-22T20:38:59.696595+05:30",
                    "updated_at": "2021-11-22T20:38:59.696595+05:30",
                    "name": "testuser2"
                }
            ]
        }
