// Code generated by ent, DO NOT EDIT.

package procedure

import (
	"entgo.io/ent/dialect/sql"
	"github.pitagora/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Procedure {
	return predicate.Procedure(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Procedure {
	return predicate.Procedure(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Procedure {
	return predicate.Procedure(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Procedure {
	return predicate.Procedure(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Procedure {
	return predicate.Procedure(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Procedure {
	return predicate.Procedure(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Procedure {
	return predicate.Procedure(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Procedure {
	return predicate.Procedure(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Procedure {
	return predicate.Procedure(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldEQ(FieldDescription, v))
}

// Metadata applies equality check predicate on the "metadata" field. It's identical to MetadataEQ.
func Metadata(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldEQ(FieldMetadata, v))
}

// Manifest applies equality check predicate on the "manifest" field. It's identical to ManifestEQ.
func Manifest(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldEQ(FieldManifest, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Procedure {
	return predicate.Procedure(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Procedure {
	return predicate.Procedure(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Procedure {
	return predicate.Procedure(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Procedure {
	return predicate.Procedure(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldContainsFold(FieldDescription, v))
}

// MetadataEQ applies the EQ predicate on the "metadata" field.
func MetadataEQ(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldEQ(FieldMetadata, v))
}

// MetadataNEQ applies the NEQ predicate on the "metadata" field.
func MetadataNEQ(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldNEQ(FieldMetadata, v))
}

// MetadataIn applies the In predicate on the "metadata" field.
func MetadataIn(vs ...string) predicate.Procedure {
	return predicate.Procedure(sql.FieldIn(FieldMetadata, vs...))
}

// MetadataNotIn applies the NotIn predicate on the "metadata" field.
func MetadataNotIn(vs ...string) predicate.Procedure {
	return predicate.Procedure(sql.FieldNotIn(FieldMetadata, vs...))
}

// MetadataGT applies the GT predicate on the "metadata" field.
func MetadataGT(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldGT(FieldMetadata, v))
}

// MetadataGTE applies the GTE predicate on the "metadata" field.
func MetadataGTE(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldGTE(FieldMetadata, v))
}

// MetadataLT applies the LT predicate on the "metadata" field.
func MetadataLT(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldLT(FieldMetadata, v))
}

// MetadataLTE applies the LTE predicate on the "metadata" field.
func MetadataLTE(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldLTE(FieldMetadata, v))
}

// MetadataContains applies the Contains predicate on the "metadata" field.
func MetadataContains(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldContains(FieldMetadata, v))
}

// MetadataHasPrefix applies the HasPrefix predicate on the "metadata" field.
func MetadataHasPrefix(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldHasPrefix(FieldMetadata, v))
}

// MetadataHasSuffix applies the HasSuffix predicate on the "metadata" field.
func MetadataHasSuffix(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldHasSuffix(FieldMetadata, v))
}

// MetadataEqualFold applies the EqualFold predicate on the "metadata" field.
func MetadataEqualFold(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldEqualFold(FieldMetadata, v))
}

// MetadataContainsFold applies the ContainsFold predicate on the "metadata" field.
func MetadataContainsFold(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldContainsFold(FieldMetadata, v))
}

// ManifestEQ applies the EQ predicate on the "manifest" field.
func ManifestEQ(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldEQ(FieldManifest, v))
}

// ManifestNEQ applies the NEQ predicate on the "manifest" field.
func ManifestNEQ(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldNEQ(FieldManifest, v))
}

// ManifestIn applies the In predicate on the "manifest" field.
func ManifestIn(vs ...string) predicate.Procedure {
	return predicate.Procedure(sql.FieldIn(FieldManifest, vs...))
}

// ManifestNotIn applies the NotIn predicate on the "manifest" field.
func ManifestNotIn(vs ...string) predicate.Procedure {
	return predicate.Procedure(sql.FieldNotIn(FieldManifest, vs...))
}

// ManifestGT applies the GT predicate on the "manifest" field.
func ManifestGT(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldGT(FieldManifest, v))
}

// ManifestGTE applies the GTE predicate on the "manifest" field.
func ManifestGTE(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldGTE(FieldManifest, v))
}

// ManifestLT applies the LT predicate on the "manifest" field.
func ManifestLT(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldLT(FieldManifest, v))
}

// ManifestLTE applies the LTE predicate on the "manifest" field.
func ManifestLTE(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldLTE(FieldManifest, v))
}

// ManifestContains applies the Contains predicate on the "manifest" field.
func ManifestContains(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldContains(FieldManifest, v))
}

// ManifestHasPrefix applies the HasPrefix predicate on the "manifest" field.
func ManifestHasPrefix(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldHasPrefix(FieldManifest, v))
}

// ManifestHasSuffix applies the HasSuffix predicate on the "manifest" field.
func ManifestHasSuffix(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldHasSuffix(FieldManifest, v))
}

// ManifestEqualFold applies the EqualFold predicate on the "manifest" field.
func ManifestEqualFold(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldEqualFold(FieldManifest, v))
}

// ManifestContainsFold applies the ContainsFold predicate on the "manifest" field.
func ManifestContainsFold(v string) predicate.Procedure {
	return predicate.Procedure(sql.FieldContainsFold(FieldManifest, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Procedure) predicate.Procedure {
	return predicate.Procedure(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Procedure) predicate.Procedure {
	return predicate.Procedure(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Procedure) predicate.Procedure {
	return predicate.Procedure(sql.NotPredicates(p))
}
