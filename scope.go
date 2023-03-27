package sqlcommentergorm

import (
	"context"

	"github.com/google/sqlcommenter/go/core"
	"gorm.io/gorm"
)

const instanceCommentsKey = "gorm:sqlcommenter:instance:comments"

func ContextInject(ctx context.Context) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// cachec comment
		if c, ok := db.Statement.Settings.Load(instanceCommentsKey); ok {
			return db.Clauses(c.(*SQLCommenter))
		}

		options := db.Statement.Context.Value(optionsKey)
		if opts, ok := options.(core.CommenterOptions); options != nil && ok {
			config := opts.Config
			var commentsMap = map[string]string{}
			// Sorted alphabetically
			if config.EnableAction && (ctx.Value(core.Action) != nil) {
				commentsMap[core.Action] = ctx.Value(core.Action).(string)
			}

			// `driver` information should not be coming from framework.
			// So, explicitly adding that here.
			if config.EnableDBDriver {
				commentsMap[core.Driver] = opts.Tags.DriverName
			}

			if config.EnableFramework && (ctx.Value(core.Framework) != nil) {
				commentsMap[core.Framework] = ctx.Value(core.Framework).(string)
			}

			if config.EnableRoute && (ctx.Value(core.Route) != nil) {
				commentsMap[core.Route] = ctx.Value(core.Route).(string)
			}

			if config.EnableTraceparent {
				carrier := core.ExtractTraceparent(ctx)
				if val, ok := carrier["traceparent"]; ok {
					commentsMap[core.Traceparent] = val
				}
			}

			if config.EnableApplication {
				commentsMap[core.Application] = opts.Tags.Application
			}

			c := &SQLCommenter{
				Prefix:  "/* ",
				Content: core.ConvertMapToComment(commentsMap),
				Suffix:  " */",
			}

			db.Statement.Settings.Store(instanceCommentsKey, c)
			return db.Clauses(c)
		}
		db.AddError(ErrMissSQLCommenterOptions)
		return db
	}
}
