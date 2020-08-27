### GraphQL Playground

Connect to http://localhost:8080

### Authentication : JWT

You need to set the Http request headers `Authorization`: `{JWT_token}`

## Usage

### Sign Up

```graphql
mutation {
  signUp(
    email: "test@test.com"
    password: "12345678"
    firstName: "graphql"
    lastName: "go"
  ) {
    ok
    error
    user {
      id
      email
      firstName
      lastName
      bio
      avatar
      createdAt
      updatedAt
    }
  }
}
```

### Sign In

```graphql
mutation {
  signIn(email: "test@test.com", password: "12345678") {
    ok
    error
    token
  }
}
```

### Change a Password

```graphql
mutation {
  changePassword(password: "87654321") {
    ok
    error
    user {
      id
      email
      firstName
      lastName
      bio
      avatar
      createdAt
      updatedAt
    }
  }
}
```

### Change a Profile

```graphql
mutation {
  changeProfile(bio: "Go developer", avatar: "go-developer.png") {
    ok
    error
    user {
      id
      email
      firstName
      lastName
      bio
      avatar
      createdAt
      updatedAt
    }
  }
}


{"query":"mutation {changeProfile(bio:\"Go developer\" avatar: \"go-developer.png\") {ok error user {id email firstName lastName bio avatar createdAt updatedAt}}}"}
```

### Get my profile

```graphql
query {
  getMyProfile {
    ok
    error
    user {
      id
      email
      firstName
      lastName
      bio
      avatar
      createdAt
      updatedAt
    }
  }
}
```
