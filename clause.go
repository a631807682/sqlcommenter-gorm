package sqlcommentergorm

import (
	"gorm.io/gorm/clause"
)

type SQLCommenter struct {
	Prefix  string
	Content string
	Suffix  string
}

// Name clause name
func (*SQLCommenter) Name() string {
	return "SQL_COMMENTER"
}

// Build
func (c *SQLCommenter) Build(builder clause.Builder) {
	builder.WriteString(c.Prefix)
	builder.WriteString(c.Content)
	builder.WriteString(c.Suffix)
}

// MergeClause merge With clauses
func (c *SQLCommenter) MergeClause(clause *clause.Clause) {
	clause.Name = ""
	clause.Expression = c
}
