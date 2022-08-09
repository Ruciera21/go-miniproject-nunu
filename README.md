# Manga API

---

**API for manga**

- `GET` /manga
- `GET` /manga/{id}
- `POST` /manga
- `PUT` /manga/{id}
- `DELETE` /manga/{id}

**Register & Login**

- `POST` /auth/register

```
{
    "name": "ayam",
    "email": "ayam@gmail.com",
    "password": "ayamgoreng"
}

```

**Register Success**

```
{
    "message": "User registered successfully"
}

```

**Register Failed**

```
{
    "message": "User already exists"
}

```

---

- `POST` /auth/login

```
{
    "email": "ayam@gmail.com",
    "password": "ayamgoreng"
}

```

**Login Success**

```
{
    "access_token": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImJ1ZGl3In0.XNdpE8HmQxKlRG2GBF2IZ0h8UbTtiH91siS5OaFHyEM"
    }
}

```
