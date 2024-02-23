# Chat Service

Chat Service includes 3 docker containers:
- chat-server - controls chats.
- postgres - permanent storage of data.
- migrator - performs migration in database using goose package.

## Deploy

To deploy Chat Service:
```
# make docker-deploy ENV=<environment>
```
*ENV is used then as a config name. Possible ENV values are now `stage` and `prod` as these configs are now in the repository.*

To stop Chat Service:
```
# make docker-stop ENV=<environment>
```
