package postgresql

import (
	"fmt"
	"go_pg_es_sync/internals/types"
	"strings"
)

type Fields types.Fields

func (fields *Fields) asJsonBuildObjectQuery(table string, additional map[string]string) string {
	parsedFields := fields.getParsedFields(table, additional)
	var rawFields []string
	for alias, raw := range parsedFields {
		rawFields = append(rawFields, fmt.Sprintf("'%s',%s", alias, raw))
	}

	return "JSONB_BUILD_OBJECT(" + strings.Join(rawFields, ",") + ")"
}

func (fields *Fields) getParsedFields(table string, additional map[string]string) map[string]string {
	raw := map[string]string{}
	if additional != nil {
		raw = additional
	}

	for _, field := range fields.Simple {
		raw[field.Alias] = fmt.Sprintf(`"%s"."%s"`, table, field.Field)
	}
	for _, field := range fields.Scripted {
		raw[field.Alias] = strings.ReplaceAll(field.Script, "{{table}}", table)
	}

	return raw
}
