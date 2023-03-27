package tests

import (
	"context"
	"strings"
	"testing"

	sqlcommenter "github.com/a631807682/sqlcommenter-gorm"
	"github.com/google/sqlcommenter/go/core"
	"github.com/google/sqlcommenter/go/net/http"
	"gorm.io/gorm"
)

func TestSQLCommenter(t *testing.T) {
	const framework = "gorm"
	const route = "GET--/mytest/:id"
	const handlerName = "mockHandler"
	ctx1 := core.ContextInject(context.TODO(), http.NewHTTPRequestTags(framework, route, handlerName))
	sess := DB.Scopes(sqlcommenter.ContextInject(ctx1)).Session(&gorm.Session{})
	user := User{Name: "foo"}
	err := sess.Create(&user).Error
	if err != nil {
		t.Fatalf("create failed %v", err)
	}

	var user1 User
	err = sess.First(&user1).Error
	if err != nil {
		t.Fatalf("find failed %v", err)
	}

	ctx2 := core.ContextInject(context.TODO(), http.NewHTTPRequestTags("mock", "mock", "mock"))
	sess2 := DB.Scopes(sqlcommenter.ContextInject(ctx2)).Session(&gorm.Session{})
	user2 := User{Name: "bar"}
	err = sess2.Create(&user2).Error
	if err != nil {
		t.Fatalf("create failed %v", err)
	}

	sql := sess2.ToSQL(func(tx *gorm.DB) *gorm.DB {
		var user3 User
		return tx.First(&user3)
	})

	comments := `/* action='mock',db_driver='mysql',framework='mock',route='mock' */`
	if !strings.Contains(sql, comments) {
		t.Fatalf("not contains comments:\n%s\nsql:\n%s", comments, sql)
	}
}
