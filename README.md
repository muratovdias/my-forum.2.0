## forum

### Objectives

This project consists in creating a web forum that allows :

- register users
- communication between users.
- associating categories to posts.
- liking and disliking posts and comments.
- filtering posts.

#### Authentication

In this segment the client must be able to `register` as a new user on the forum, by inputting their credentials. You also have to create a `login session` to access the forum and be able to add posts and comments.

I use cookies to allow each user to have only one opened session. Each of this sessions contains an expiration date, used of UUID.
The password encrypts before store.

#### Communication

In order for users to communicate between each other, they have able to create posts and comments.

- Only registered users will be able to create posts and comments.
- When registered users are creating a post they can associate one or more categories to it.
  - The implementation and choice of the categories is up to you.
- The posts and comments should be visible to all users (registered or not).
- Non-registered users will only be able to see posts and comments.

#### Likes and Dislikes

Only registered users able to like or dislike posts and comments.

The number of likes and dislikes visible by all users (registered or not).

#### Filter

Filter allow users to filter the displayed posts by :

- categories
- created posts
- liked posts

You can look at filtering by categories.

Note that the last two are only available for registered users and refer to the logged in user.

#### Docker

```````````console
docker build -t forum .
docker run -p 8888:8888 forum
```````````
### Used packages

- All [standard Go](https://golang.org/pkg/) packages are allowed.
- [sqlite3](https://github.com/mattn/go-sqlite3)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [UUID](https://github.com/gofrs/uuid)

This project will help you learn about:

- The basics of web :
  - HTML
  - HTTP
  - Sessions and cookies
- Using and [setting up Docker](https://docs.docker.com/get-started/)
  - Containerizing an application
  - Compatibility/Dependency
  - Creating images
- SQL language
  - Manipulation of databases
- The basics of encryption