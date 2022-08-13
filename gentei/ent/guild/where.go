// Code generated by ent, DO NOT EDIT.

package guild

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/member-gentei/member-gentei/gentei/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// IconHash applies equality check predicate on the "icon_hash" field. It's identical to IconHashEQ.
func IconHash(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIconHash), v))
	})
}

// AuditChannel applies equality check predicate on the "audit_channel" field. It's identical to AuditChannelEQ.
func AuditChannel(v uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuditChannel), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Guild {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Guild {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// IconHashEQ applies the EQ predicate on the "icon_hash" field.
func IconHashEQ(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIconHash), v))
	})
}

// IconHashNEQ applies the NEQ predicate on the "icon_hash" field.
func IconHashNEQ(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIconHash), v))
	})
}

// IconHashIn applies the In predicate on the "icon_hash" field.
func IconHashIn(vs ...string) predicate.Guild {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIconHash), v...))
	})
}

// IconHashNotIn applies the NotIn predicate on the "icon_hash" field.
func IconHashNotIn(vs ...string) predicate.Guild {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIconHash), v...))
	})
}

// IconHashGT applies the GT predicate on the "icon_hash" field.
func IconHashGT(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIconHash), v))
	})
}

// IconHashGTE applies the GTE predicate on the "icon_hash" field.
func IconHashGTE(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIconHash), v))
	})
}

// IconHashLT applies the LT predicate on the "icon_hash" field.
func IconHashLT(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIconHash), v))
	})
}

// IconHashLTE applies the LTE predicate on the "icon_hash" field.
func IconHashLTE(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIconHash), v))
	})
}

// IconHashContains applies the Contains predicate on the "icon_hash" field.
func IconHashContains(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIconHash), v))
	})
}

// IconHashHasPrefix applies the HasPrefix predicate on the "icon_hash" field.
func IconHashHasPrefix(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIconHash), v))
	})
}

// IconHashHasSuffix applies the HasSuffix predicate on the "icon_hash" field.
func IconHashHasSuffix(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIconHash), v))
	})
}

// IconHashIsNil applies the IsNil predicate on the "icon_hash" field.
func IconHashIsNil() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIconHash)))
	})
}

// IconHashNotNil applies the NotNil predicate on the "icon_hash" field.
func IconHashNotNil() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIconHash)))
	})
}

// IconHashEqualFold applies the EqualFold predicate on the "icon_hash" field.
func IconHashEqualFold(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIconHash), v))
	})
}

// IconHashContainsFold applies the ContainsFold predicate on the "icon_hash" field.
func IconHashContainsFold(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIconHash), v))
	})
}

// AuditChannelEQ applies the EQ predicate on the "audit_channel" field.
func AuditChannelEQ(v uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuditChannel), v))
	})
}

// AuditChannelNEQ applies the NEQ predicate on the "audit_channel" field.
func AuditChannelNEQ(v uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAuditChannel), v))
	})
}

// AuditChannelIn applies the In predicate on the "audit_channel" field.
func AuditChannelIn(vs ...uint64) predicate.Guild {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAuditChannel), v...))
	})
}

// AuditChannelNotIn applies the NotIn predicate on the "audit_channel" field.
func AuditChannelNotIn(vs ...uint64) predicate.Guild {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAuditChannel), v...))
	})
}

// AuditChannelGT applies the GT predicate on the "audit_channel" field.
func AuditChannelGT(v uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAuditChannel), v))
	})
}

// AuditChannelGTE applies the GTE predicate on the "audit_channel" field.
func AuditChannelGTE(v uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAuditChannel), v))
	})
}

// AuditChannelLT applies the LT predicate on the "audit_channel" field.
func AuditChannelLT(v uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAuditChannel), v))
	})
}

// AuditChannelLTE applies the LTE predicate on the "audit_channel" field.
func AuditChannelLTE(v uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAuditChannel), v))
	})
}

// AuditChannelIsNil applies the IsNil predicate on the "audit_channel" field.
func AuditChannelIsNil() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAuditChannel)))
	})
}

// AuditChannelNotNil applies the NotNil predicate on the "audit_channel" field.
func AuditChannelNotNil() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAuditChannel)))
	})
}

// LanguageEQ applies the EQ predicate on the "language" field.
func LanguageEQ(v Language) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLanguage), v))
	})
}

// LanguageNEQ applies the NEQ predicate on the "language" field.
func LanguageNEQ(v Language) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLanguage), v))
	})
}

// LanguageIn applies the In predicate on the "language" field.
func LanguageIn(vs ...Language) predicate.Guild {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLanguage), v...))
	})
}

// LanguageNotIn applies the NotIn predicate on the "language" field.
func LanguageNotIn(vs ...Language) predicate.Guild {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLanguage), v...))
	})
}

// ModeratorSnowflakesIsNil applies the IsNil predicate on the "moderator_snowflakes" field.
func ModeratorSnowflakesIsNil() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldModeratorSnowflakes)))
	})
}

// ModeratorSnowflakesNotNil applies the NotNil predicate on the "moderator_snowflakes" field.
func ModeratorSnowflakesNotNil() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldModeratorSnowflakes)))
	})
}

// SettingsIsNil applies the IsNil predicate on the "settings" field.
func SettingsIsNil() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSettings)))
	})
}

// SettingsNotNil applies the NotNil predicate on the "settings" field.
func SettingsNotNil() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSettings)))
	})
}

// HasMembers applies the HasEdge predicate on the "members" edge.
func HasMembers() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MembersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, MembersTable, MembersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMembersWith applies the HasEdge predicate on the "members" edge with a given conditions (other predicates).
func HasMembersWith(preds ...predicate.User) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MembersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, MembersTable, MembersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAdmins applies the HasEdge predicate on the "admins" edge.
func HasAdmins() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AdminsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AdminsTable, AdminsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAdminsWith applies the HasEdge predicate on the "admins" edge with a given conditions (other predicates).
func HasAdminsWith(preds ...predicate.User) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AdminsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AdminsTable, AdminsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoles applies the HasEdge predicate on the "roles" edge.
func HasRoles() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RolesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RolesTable, RolesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolesWith applies the HasEdge predicate on the "roles" edge with a given conditions (other predicates).
func HasRolesWith(preds ...predicate.GuildRole) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RolesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RolesTable, RolesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasYoutubeTalents applies the HasEdge predicate on the "youtube_talents" edge.
func HasYoutubeTalents() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(YoutubeTalentsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, YoutubeTalentsTable, YoutubeTalentsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasYoutubeTalentsWith applies the HasEdge predicate on the "youtube_talents" edge with a given conditions (other predicates).
func HasYoutubeTalentsWith(preds ...predicate.YouTubeTalent) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(YoutubeTalentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, YoutubeTalentsTable, YoutubeTalentsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Guild) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Guild) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Guild) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		p(s.Not())
	})
}
