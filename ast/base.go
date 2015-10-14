// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package ast

import "github.com/pingcap/tidb/util/types"

// node is the struct implements node interface except for Accept method.
// Node implementations should embed it in.
type node struct {
	txt string
}

// SetText implements Node interface.
func (n *node) SetText(text string) {
	n.txt = text
}

// Text implements Node interface.
func (n *node) Text() string {
	return n.txt
}

// stmtNode implements StmtNode interface.
// Statement implementations should embed it in.
type stmtNode struct {
	node
}

// statement implements StmtNode interface.
func (sn *stmtNode) statement() {}

// ddlNode implements DDLNode interface.
// DDL implementations should embed it in.
type ddlNode struct {
	stmtNode
}

// ddlStatement implements DDLNode interface.
func (dn *ddlNode) ddlStatement() {}

// dmlNode is the struct implements DMLNode interface.
// DML implementations should embed it in.
type dmlNode struct {
	stmtNode
}

// dmlStatement implements DMLNode interface.
func (dn *dmlNode) dmlStatement() {}

// expressionNode is the struct implements Expression interface.
// Expression implementations should embed it in.
type exprNode struct {
	node
	tp *types.FieldType
}

// IsStatic implements Expression interface.
func (en *exprNode) IsStatic() bool {
	return false
}

// SetType implements Expression interface.
func (en *exprNode) SetType(tp *types.FieldType) {
	en.tp = tp
}

// GetType implements Expression interface.
func (en *exprNode) GetType() *types.FieldType {
	return en.tp
}

type funcNode struct {
	exprNode
}

// FunctionExpression implements FounctionNode interface.
func (fn *funcNode) functionExpression() {}
