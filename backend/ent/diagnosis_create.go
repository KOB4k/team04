// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/B6001186/Contagions/ent/diagnosis"
	"github.com/B6001186/Contagions/ent/disease"
	"github.com/B6001186/Contagions/ent/employee"
	"github.com/B6001186/Contagions/ent/patient"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DiagnosisCreate is the builder for creating a Diagnosis entity.
type DiagnosisCreate struct {
	config
	mutation *DiagnosisMutation
	hooks    []Hook
}

// SetDiagnosticMessages sets the DiagnosticMessages field.
func (dc *DiagnosisCreate) SetDiagnosticMessages(s string) *DiagnosisCreate {
	dc.mutation.SetDiagnosticMessages(s)
	return dc
}

// SetSurveillancePeriod sets the SurveillancePeriod field.
func (dc *DiagnosisCreate) SetSurveillancePeriod(s string) *DiagnosisCreate {
	dc.mutation.SetSurveillancePeriod(s)
	return dc
}

// SetDiagnosisDate sets the DiagnosisDate field.
func (dc *DiagnosisCreate) SetDiagnosisDate(s string) *DiagnosisCreate {
	dc.mutation.SetDiagnosisDate(s)
	return dc
}

// AddDiseaseIDs adds the disease edge to Disease by ids.
func (dc *DiagnosisCreate) AddDiseaseIDs(ids ...int) *DiagnosisCreate {
	dc.mutation.AddDiseaseIDs(ids...)
	return dc
}

// AddDisease adds the disease edges to Disease.
func (dc *DiagnosisCreate) AddDisease(d ...*Disease) *DiagnosisCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDiseaseIDs(ids...)
}

// SetPatientID sets the patient edge to Patient by id.
func (dc *DiagnosisCreate) SetPatientID(id int) *DiagnosisCreate {
	dc.mutation.SetPatientID(id)
	return dc
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (dc *DiagnosisCreate) SetNillablePatientID(id *int) *DiagnosisCreate {
	if id != nil {
		dc = dc.SetPatientID(*id)
	}
	return dc
}

// SetPatient sets the patient edge to Patient.
func (dc *DiagnosisCreate) SetPatient(p *Patient) *DiagnosisCreate {
	return dc.SetPatientID(p.ID)
}

// SetEmployeeID sets the employee edge to Employee by id.
func (dc *DiagnosisCreate) SetEmployeeID(id int) *DiagnosisCreate {
	dc.mutation.SetEmployeeID(id)
	return dc
}

// SetNillableEmployeeID sets the employee edge to Employee by id if the given value is not nil.
func (dc *DiagnosisCreate) SetNillableEmployeeID(id *int) *DiagnosisCreate {
	if id != nil {
		dc = dc.SetEmployeeID(*id)
	}
	return dc
}

// SetEmployee sets the employee edge to Employee.
func (dc *DiagnosisCreate) SetEmployee(e *Employee) *DiagnosisCreate {
	return dc.SetEmployeeID(e.ID)
}

// Mutation returns the DiagnosisMutation object of the builder.
func (dc *DiagnosisCreate) Mutation() *DiagnosisMutation {
	return dc.mutation
}

// Save creates the Diagnosis in the database.
func (dc *DiagnosisCreate) Save(ctx context.Context) (*Diagnosis, error) {
	if _, ok := dc.mutation.DiagnosticMessages(); !ok {
		return nil, &ValidationError{Name: "DiagnosticMessages", err: errors.New("ent: missing required field \"DiagnosticMessages\"")}
	}
	if _, ok := dc.mutation.SurveillancePeriod(); !ok {
		return nil, &ValidationError{Name: "SurveillancePeriod", err: errors.New("ent: missing required field \"SurveillancePeriod\"")}
	}
	if _, ok := dc.mutation.DiagnosisDate(); !ok {
		return nil, &ValidationError{Name: "DiagnosisDate", err: errors.New("ent: missing required field \"DiagnosisDate\"")}
	}
	var (
		err  error
		node *Diagnosis
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiagnosisMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DiagnosisCreate) SaveX(ctx context.Context) *Diagnosis {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DiagnosisCreate) sqlSave(ctx context.Context) (*Diagnosis, error) {
	d, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}

func (dc *DiagnosisCreate) createSpec() (*Diagnosis, *sqlgraph.CreateSpec) {
	var (
		d     = &Diagnosis{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: diagnosis.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: diagnosis.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.DiagnosticMessages(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diagnosis.FieldDiagnosticMessages,
		})
		d.DiagnosticMessages = value
	}
	if value, ok := dc.mutation.SurveillancePeriod(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diagnosis.FieldSurveillancePeriod,
		})
		d.SurveillancePeriod = value
	}
	if value, ok := dc.mutation.DiagnosisDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diagnosis.FieldDiagnosisDate,
		})
		d.DiagnosisDate = value
	}
	if nodes := dc.mutation.DiseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   diagnosis.DiseaseTable,
			Columns: []string{diagnosis.DiseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disease.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   diagnosis.PatientTable,
			Columns: []string{diagnosis.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   diagnosis.EmployeeTable,
			Columns: []string{diagnosis.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return d, _spec
}