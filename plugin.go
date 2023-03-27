package sqlcommentergorm

import (
	"context"
	"errors"
	"runtime/debug"

	"github.com/google/sqlcommenter/go/core"
	"gorm.io/gorm"
)

type cacheKey string

const optionsKey cacheKey = "gorm:sqlcommenter:options"

var ErrMissSQLCommenterOptions = errors.New("missing sqlcommenter options")

// SQLCommenterPlugin gorm plugin
type SQLCommenterPlugin struct {
	Options core.CommenterOptions
}

// Name gorm plugin name
func (*SQLCommenterPlugin) Name() string {
	return "gorm:sqlcommenter"
}

// Initialize gorm plugin initialize
func (p *SQLCommenterPlugin) Initialize(db *gorm.DB) error {
	// driver name
	p.Options.Tags.DriverName = db.Dialector.Name()
	// clause ext
	callback := db.Callback()
	clauseName := (&SQLCommenter{}).Name()
	callback.Create().Clauses = append(callback.Create().Clauses, clauseName)
	callback.Query().Clauses = append(callback.Query().Clauses, clauseName)
	callback.Update().Clauses = append(callback.Update().Clauses, clauseName)
	callback.Delete().Clauses = append(callback.Delete().Clauses, clauseName)
	callback.Row().Clauses = append(callback.Row().Clauses, clauseName)
	callback.Raw().Clauses = append(callback.Raw().Clauses, clauseName)
	// save options
	db.Statement.Context = context.WithValue(db.Statement.Context, optionsKey, p.Options)
	return nil
}

type Config struct {
	Application string

	DisableRoute       bool
	DisableFramework   bool
	DisableController  bool
	DisableAction      bool
	DisableTraceparent bool
}

func New(config *Config) *SQLCommenterPlugin {
	return &SQLCommenterPlugin{
		Options: core.CommenterOptions{
			Config: core.CommenterConfig{
				EnableDBDriver:    true,
				EnableRoute:       !config.DisableRoute,
				EnableFramework:   !config.DisableFramework,
				EnableController:  !config.DisableController,
				EnableAction:      !config.DisableAction,
				EnableTraceparent: !config.DisableTraceparent,
				EnableApplication: config.Application != "",
			},
			Tags: core.StaticTags{
				Application: config.Application,
			},
		},
	}
}

func Default() *SQLCommenterPlugin {
	var application string
	bi, ok := debug.ReadBuildInfo()
	if ok {
		application = bi.Path
	}
	return &SQLCommenterPlugin{
		Options: core.CommenterOptions{
			Config: core.CommenterConfig{
				EnableDBDriver:    true,
				EnableRoute:       true,
				EnableFramework:   true,
				EnableController:  true,
				EnableAction:      true,
				EnableTraceparent: true,
				EnableApplication: application != "",
			},
			Tags: core.StaticTags{
				Application: application,
			},
		},
	}
}
