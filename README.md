# sqlcommenter-gorm

gorm plugin for https://github.com/google/sqlcommenter

[![go report card](https://goreportcard.com/badge/github.com/a631807682/sqlcommenter-gorm "go report card")](https://goreportcard.com/report/github.com/a631807682/sqlcommenter-gorm)
[![test status](https://github.com/a631807682/sqlcommenter-gorm/workflows/tests/badge.svg?branch=main "test status")](https://github.com/a631807682/sqlcommenter-gorm/actions)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/a631807682/sqlcommenter-gorm)
![visitor badge](https://visitor-badge.glitch.me/badge?page_id=a631807682.sqlcommenter-gorm)

> sqlcommenter is a suite of middlewares/plugins that enable your ORMs to augment SQL statements before execution, with comments containing information about the code that caused its execution. This helps in easily correlating slow performance with source code and giving insights into backend database performance. In short it provides some observability into the state of your client-side applications and their impact on the database’s server-side.

## Usage

```go
    ...
    db.Use(sqlcommentergorm.Default())
    ...
    sess := DB.Scopes(sqlcommentergorm.ContextInject(ctx)).Session(&gorm.Session{})
    sess.Create(&user)
    // inject the following comments
    /* action='mockHandler',db_driver='mysql',framework='gorm',route='GET--%2Fmytest%2F%3Aid' */
```
