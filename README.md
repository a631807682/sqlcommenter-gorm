# sqlcommenter-gorm

gorm plugin for https://github.com/google/sqlcommenter

> sqlcommenter is a suite of middlewares/plugins that enable your ORMs to augment SQL statements before execution, with comments containing information about the code that caused its execution. This helps in easily correlating slow performance with source code and giving insights into backend database performance. In short it provides some observability into the state of your client-side applications and their impact on the database’s server-side.

## Usage

```go
    ...
    db.Use(sqlcommentergorm.Default())
    ...
    sess := DB.Scopes(sqlcommentergorm.ContextInject(ctx)).Session(&gorm.Session{})
    sess.Create(&user)
    // INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`birthday`,`active`) VALUES ('2023-03-27 16:03:14.276','2023-03-27 16:03:14.276',NULL,'foo',18,NULL,false) /* action='mockHandler',db_driver='mysql',framework='gorm',route='GET--%2Fmytest%2F%3Aid' */
```
