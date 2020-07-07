package jsonlines

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/sourcegraph/sourcegraph/enterprise/cmd/precise-code-intel-worker/internal/correlation/datastructures"
	"github.com/sourcegraph/sourcegraph/enterprise/cmd/precise-code-intel-worker/internal/correlation/lsif"
)

var unmarshaller = jsoniter.ConfigFastest

func unmarshalElement(line []byte) (_ lsif.Element, err error) {
	var payload struct {
		ID    ID     `json:"id"`
		Type  string `json:"type"`
		Label string `json:"label"`
	}
	if err := unmarshaller.Unmarshal(line, &payload); err != nil {
		return lsif.Element{}, err
	}

	element := lsif.Element{
		ID:    int(payload.ID),
		Type:  payload.Type,
		Label: payload.Label,
	}

	if payload.Type == "edge" {
		element.Payload, err = unmarshalEdge(line)
	} else if payload.Type == "vertex" {
		if unmarshaler, ok := vertexUnmarshalers[payload.Label]; ok {
			element.Payload, err = unmarshaler(line)
		}
	}

	return element, err
}

func unmarshalEdge(line []byte) (interface{}, error) {
	var payload struct {
		OutV     ID   `json:"outV"`
		InV      ID   `json:"inV"`
		InVs     []ID `json:"inVs"`
		Document ID   `json:"document"`
	}
	if err := unmarshaller.Unmarshal(line, &payload); err != nil {
		return lsif.Edge{}, err
	}

	var inVs []int
	for _, inV := range payload.InVs {
		inVs = append(inVs, int(inV))
	}

	return lsif.Edge{
		OutV:     int(payload.OutV),
		InV:      int(payload.InV),
		InVs:     inVs,
		Document: int(payload.Document),
	}, nil
}

var vertexUnmarshalers = map[string]func(line []byte) (interface{}, error){
	"metaData":           unmarshalMetaData,
	"document":           unmarshalDocument,
	"range":              unmarshalRange,
	"hoverResult":        unmarshalHover,
	"moniker":            unmarshalMoniker,
	"packageInformation": unmarshalPackageInformation,
	"diagnosticResult":   unmarshalDiagnosticResult,
}

func unmarshalMetaData(line []byte) (interface{}, error) {
	var payload struct {
		Version     string `json:"version"`
		ProjectRoot string `json:"projectRoot"`
	}
	if err := unmarshaller.Unmarshal(line, &payload); err != nil {
		return nil, err
	}

	return lsif.MetaData{
		Version:     payload.Version,
		ProjectRoot: payload.ProjectRoot,
	}, nil
}

func unmarshalDocument(line []byte) (interface{}, error) {
	var payload struct {
		URI string `json:"uri"`
	}
	if err := unmarshaller.Unmarshal(line, &payload); err != nil {
		return nil, err
	}

	return lsif.Document{
		URI:         payload.URI,
		Contains:    datastructures.NewIDSet(),
		Diagnostics: datastructures.NewIDSet(),
	}, nil
}

func unmarshalRange(line []byte) (interface{}, error) {
	type _position struct {
		Line      int `json:"line"`
		Character int `json:"character"`
	}
	var payload struct {
		Start _position `json:"start"`
		End   _position `json:"end"`
	}
	if err := unmarshaller.Unmarshal(line, &payload); err != nil {
		return nil, err
	}

	return lsif.Range{
		StartLine:      payload.Start.Line,
		StartCharacter: payload.Start.Character,
		EndLine:        payload.End.Line,
		EndCharacter:   payload.End.Character,
		MonikerIDs:     datastructures.NewIDSet(),
	}, nil
}

func unmarshalHover(line []byte) (interface{}, error) {
	type _hoverResult struct {
		Contents json.RawMessage `json:"contents"`
	}
	var payload struct {
		Result _hoverResult `json:"result"`
	}
	if err := unmarshaller.Unmarshal(line, &payload); err != nil {
		return nil, err
	}

	var target []json.RawMessage
	if err := unmarshaller.Unmarshal(payload.Result.Contents, &target); err != nil {
		v, err := unmarshalHoverPart(payload.Result.Contents)
		if err != nil {
			return nil, err
		}

		return string(v), nil
	}

	var parts [][]byte
	for _, t := range target {
		part, err := unmarshalHoverPart(t)
		if err != nil {
			return "", err
		}

		parts = append(parts, part)
	}

	return string(bytes.Join(parts, []byte("\n\n---\n\n"))), nil
}

func unmarshalHoverPart(raw json.RawMessage) ([]byte, error) {
	var strPayload string
	if err := unmarshaller.Unmarshal(raw, &strPayload); err == nil {
		return bytes.TrimSpace([]byte(strPayload)), nil
	}

	var objPayload struct {
		Language string `json:"language"`
		Value    string `json:"value"`
	}
	if err := unmarshaller.Unmarshal(raw, &objPayload); err != nil {
		return nil, errors.New("unrecognized hover format")
	}

	if len(objPayload.Language) > 0 {
		v := make([]byte, 0, 8+len(objPayload.Language)+len(objPayload.Value))
		v = append(v, "```"...)
		v = append(v, []byte(objPayload.Language)...)
		v = append(v, '\n')
		v = append(v, []byte(objPayload.Value)...)
		v = append(v, "\n```"...)

		return v, nil
	}

	return bytes.TrimSpace([]byte(objPayload.Value)), nil
}

func unmarshalMoniker(line []byte) (interface{}, error) {
	var payload struct {
		Kind       string `json:"kind"`
		Scheme     string `json:"scheme"`
		Identifier string `json:"identifier"`
	}
	if err := unmarshaller.Unmarshal(line, &payload); err != nil {
		return nil, err
	}

	if payload.Kind == "" {
		payload.Kind = "local"
	}

	return lsif.Moniker{
		Kind:       payload.Kind,
		Scheme:     payload.Scheme,
		Identifier: payload.Identifier,
	}, nil
}

func unmarshalPackageInformation(line []byte) (interface{}, error) {
	var payload struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
	if err := unmarshaller.Unmarshal(line, &payload); err != nil {
		return nil, err
	}

	return lsif.PackageInformation{
		Name:    payload.Name,
		Version: payload.Version,
	}, nil
}

func unmarshalDiagnosticResult(line []byte) (interface{}, error) {
	type _position struct {
		Line      int `json:"line"`
		Character int `json:"character"`
	}
	type _range struct {
		Start _position `json:"start"`
		End   _position `json:"end"`
	}
	type _result struct {
		Severity int         `json:"severity"`
		Code     StringOrInt `json:"code"`
		Message  string      `json:"message"`
		Source   string      `json:"source"`
		Range    _range      `json:"range"`
	}
	var payload struct {
		Results []_result `json:"result"`
	}
	if err := unmarshaller.Unmarshal(line, &payload); err != nil {
		return nil, err
	}

	var diagnostics []lsif.Diagnostic
	for _, result := range payload.Results {
		diagnostics = append(diagnostics, lsif.Diagnostic{
			Severity:       result.Severity,
			Code:           string(result.Code),
			Message:        result.Message,
			Source:         result.Source,
			StartLine:      result.Range.Start.Line,
			StartCharacter: result.Range.Start.Character,
			EndLine:        result.Range.End.Line,
			EndCharacter:   result.Range.End.Character,
		})
	}

	return lsif.DiagnosticResult{Result: diagnostics}, nil
}

// TODO(efritz) - make non-global
var DefaultInterner = NewInterner()

type ID int

func (id *ID) UnmarshalJSON(raw []byte) error {
	v, err := DefaultInterner.Intern(raw)
	if err != nil {
		return err
	}

	*id = ID(v)
	return nil
}

type StringOrInt string

func (id *StringOrInt) UnmarshalJSON(raw []byte) error {
	if raw[0] == '"' {
		var v string
		if err := unmarshaller.Unmarshal(raw, &v); err != nil {
			return err
		}

		*id = StringOrInt(v)
		return nil
	}

	var v int64
	if err := unmarshaller.Unmarshal(raw, &v); err != nil {
		return err
	}

	*id = StringOrInt(strconv.FormatInt(v, 10))
	return nil
}
