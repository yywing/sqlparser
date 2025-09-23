// Code generated from FlinkSqlParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // FlinkSqlParser

import "github.com/antlr4-go/antlr/v4"

// BaseFlinkSqlParserListener is a complete listener for a parse tree produced by FlinkSqlParser.
type BaseFlinkSqlParserListener struct{}

var _ FlinkSqlParserListener = &BaseFlinkSqlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFlinkSqlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFlinkSqlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFlinkSqlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFlinkSqlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSqlStatements is called when production sqlStatements is entered.
func (s *BaseFlinkSqlParserListener) EnterSqlStatements(ctx *SqlStatementsContext) {}

// ExitSqlStatements is called when production sqlStatements is exited.
func (s *BaseFlinkSqlParserListener) ExitSqlStatements(ctx *SqlStatementsContext) {}

// EnterSingleStatement is called when production singleStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterSingleStatement(ctx *SingleStatementContext) {}

// ExitSingleStatement is called when production singleStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitSingleStatement(ctx *SingleStatementContext) {}

// EnterSqlStatement is called when production sqlStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterSqlStatement(ctx *SqlStatementContext) {}

// ExitSqlStatement is called when production sqlStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitSqlStatement(ctx *SqlStatementContext) {}

// EnterEmptyStatement is called when production emptyStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterEmptyStatement(ctx *EmptyStatementContext) {}

// ExitEmptyStatement is called when production emptyStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitEmptyStatement(ctx *EmptyStatementContext) {}

// EnterDdlStatement is called when production ddlStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterDdlStatement(ctx *DdlStatementContext) {}

// ExitDdlStatement is called when production ddlStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitDdlStatement(ctx *DdlStatementContext) {}

// EnterDmlStatement is called when production dmlStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterDmlStatement(ctx *DmlStatementContext) {}

// ExitDmlStatement is called when production dmlStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitDmlStatement(ctx *DmlStatementContext) {}

// EnterDescribeStatement is called when production describeStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterDescribeStatement(ctx *DescribeStatementContext) {}

// ExitDescribeStatement is called when production describeStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitDescribeStatement(ctx *DescribeStatementContext) {}

// EnterExplainStatement is called when production explainStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterExplainStatement(ctx *ExplainStatementContext) {}

// ExitExplainStatement is called when production explainStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitExplainStatement(ctx *ExplainStatementContext) {}

// EnterExplainDetails is called when production explainDetails is entered.
func (s *BaseFlinkSqlParserListener) EnterExplainDetails(ctx *ExplainDetailsContext) {}

// ExitExplainDetails is called when production explainDetails is exited.
func (s *BaseFlinkSqlParserListener) ExitExplainDetails(ctx *ExplainDetailsContext) {}

// EnterExplainDetail is called when production explainDetail is entered.
func (s *BaseFlinkSqlParserListener) EnterExplainDetail(ctx *ExplainDetailContext) {}

// ExitExplainDetail is called when production explainDetail is exited.
func (s *BaseFlinkSqlParserListener) ExitExplainDetail(ctx *ExplainDetailContext) {}

// EnterUseStatement is called when production useStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterUseStatement(ctx *UseStatementContext) {}

// ExitUseStatement is called when production useStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitUseStatement(ctx *UseStatementContext) {}

// EnterUseModuleStatement is called when production useModuleStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterUseModuleStatement(ctx *UseModuleStatementContext) {}

// ExitUseModuleStatement is called when production useModuleStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitUseModuleStatement(ctx *UseModuleStatementContext) {}

// EnterShowStatement is called when production showStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterShowStatement(ctx *ShowStatementContext) {}

// ExitShowStatement is called when production showStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitShowStatement(ctx *ShowStatementContext) {}

// EnterLoadStatement is called when production loadStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterLoadStatement(ctx *LoadStatementContext) {}

// ExitLoadStatement is called when production loadStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitLoadStatement(ctx *LoadStatementContext) {}

// EnterUnloadStatement is called when production unloadStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterUnloadStatement(ctx *UnloadStatementContext) {}

// ExitUnloadStatement is called when production unloadStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitUnloadStatement(ctx *UnloadStatementContext) {}

// EnterSetStatement is called when production setStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterSetStatement(ctx *SetStatementContext) {}

// ExitSetStatement is called when production setStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitSetStatement(ctx *SetStatementContext) {}

// EnterResetStatement is called when production resetStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterResetStatement(ctx *ResetStatementContext) {}

// ExitResetStatement is called when production resetStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitResetStatement(ctx *ResetStatementContext) {}

// EnterJarStatement is called when production jarStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterJarStatement(ctx *JarStatementContext) {}

// ExitJarStatement is called when production jarStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitJarStatement(ctx *JarStatementContext) {}

// EnterDtAddStatement is called when production dtAddStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterDtAddStatement(ctx *DtAddStatementContext) {}

// ExitDtAddStatement is called when production dtAddStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitDtAddStatement(ctx *DtAddStatementContext) {}

// EnterDtFilePath is called when production dtFilePath is entered.
func (s *BaseFlinkSqlParserListener) EnterDtFilePath(ctx *DtFilePathContext) {}

// ExitDtFilePath is called when production dtFilePath is exited.
func (s *BaseFlinkSqlParserListener) ExitDtFilePath(ctx *DtFilePathContext) {}

// EnterCreateTable is called when production createTable is entered.
func (s *BaseFlinkSqlParserListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production createTable is exited.
func (s *BaseFlinkSqlParserListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterSimpleCreateTable is called when production simpleCreateTable is entered.
func (s *BaseFlinkSqlParserListener) EnterSimpleCreateTable(ctx *SimpleCreateTableContext) {}

// ExitSimpleCreateTable is called when production simpleCreateTable is exited.
func (s *BaseFlinkSqlParserListener) ExitSimpleCreateTable(ctx *SimpleCreateTableContext) {}

// EnterCreateTableAsSelect is called when production createTableAsSelect is entered.
func (s *BaseFlinkSqlParserListener) EnterCreateTableAsSelect(ctx *CreateTableAsSelectContext) {}

// ExitCreateTableAsSelect is called when production createTableAsSelect is exited.
func (s *BaseFlinkSqlParserListener) ExitCreateTableAsSelect(ctx *CreateTableAsSelectContext) {}

// EnterColumnOptionDefinition is called when production columnOptionDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterColumnOptionDefinition(ctx *ColumnOptionDefinitionContext) {
}

// ExitColumnOptionDefinition is called when production columnOptionDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitColumnOptionDefinition(ctx *ColumnOptionDefinitionContext) {}

// EnterPhysicalColumnDefinition is called when production physicalColumnDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterPhysicalColumnDefinition(ctx *PhysicalColumnDefinitionContext) {
}

// ExitPhysicalColumnDefinition is called when production physicalColumnDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitPhysicalColumnDefinition(ctx *PhysicalColumnDefinitionContext) {
}

// EnterColumnNameCreate is called when production columnNameCreate is entered.
func (s *BaseFlinkSqlParserListener) EnterColumnNameCreate(ctx *ColumnNameCreateContext) {}

// ExitColumnNameCreate is called when production columnNameCreate is exited.
func (s *BaseFlinkSqlParserListener) ExitColumnNameCreate(ctx *ColumnNameCreateContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseFlinkSqlParserListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseFlinkSqlParserListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterColumnNamePath is called when production columnNamePath is entered.
func (s *BaseFlinkSqlParserListener) EnterColumnNamePath(ctx *ColumnNamePathContext) {}

// ExitColumnNamePath is called when production columnNamePath is exited.
func (s *BaseFlinkSqlParserListener) ExitColumnNamePath(ctx *ColumnNamePathContext) {}

// EnterColumnNameList is called when production columnNameList is entered.
func (s *BaseFlinkSqlParserListener) EnterColumnNameList(ctx *ColumnNameListContext) {}

// ExitColumnNameList is called when production columnNameList is exited.
func (s *BaseFlinkSqlParserListener) ExitColumnNameList(ctx *ColumnNameListContext) {}

// EnterColumnType is called when production columnType is entered.
func (s *BaseFlinkSqlParserListener) EnterColumnType(ctx *ColumnTypeContext) {}

// ExitColumnType is called when production columnType is exited.
func (s *BaseFlinkSqlParserListener) ExitColumnType(ctx *ColumnTypeContext) {}

// EnterLengthOneDimension is called when production lengthOneDimension is entered.
func (s *BaseFlinkSqlParserListener) EnterLengthOneDimension(ctx *LengthOneDimensionContext) {}

// ExitLengthOneDimension is called when production lengthOneDimension is exited.
func (s *BaseFlinkSqlParserListener) ExitLengthOneDimension(ctx *LengthOneDimensionContext) {}

// EnterLengthTwoOptionalDimension is called when production lengthTwoOptionalDimension is entered.
func (s *BaseFlinkSqlParserListener) EnterLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) {
}

// ExitLengthTwoOptionalDimension is called when production lengthTwoOptionalDimension is exited.
func (s *BaseFlinkSqlParserListener) ExitLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) {
}

// EnterLengthTwoStringDimension is called when production lengthTwoStringDimension is entered.
func (s *BaseFlinkSqlParserListener) EnterLengthTwoStringDimension(ctx *LengthTwoStringDimensionContext) {
}

// ExitLengthTwoStringDimension is called when production lengthTwoStringDimension is exited.
func (s *BaseFlinkSqlParserListener) ExitLengthTwoStringDimension(ctx *LengthTwoStringDimensionContext) {
}

// EnterLengthOneTypeDimension is called when production lengthOneTypeDimension is entered.
func (s *BaseFlinkSqlParserListener) EnterLengthOneTypeDimension(ctx *LengthOneTypeDimensionContext) {
}

// ExitLengthOneTypeDimension is called when production lengthOneTypeDimension is exited.
func (s *BaseFlinkSqlParserListener) ExitLengthOneTypeDimension(ctx *LengthOneTypeDimensionContext) {}

// EnterMapTypeDimension is called when production mapTypeDimension is entered.
func (s *BaseFlinkSqlParserListener) EnterMapTypeDimension(ctx *MapTypeDimensionContext) {}

// ExitMapTypeDimension is called when production mapTypeDimension is exited.
func (s *BaseFlinkSqlParserListener) ExitMapTypeDimension(ctx *MapTypeDimensionContext) {}

// EnterRowTypeDimension is called when production rowTypeDimension is entered.
func (s *BaseFlinkSqlParserListener) EnterRowTypeDimension(ctx *RowTypeDimensionContext) {}

// ExitRowTypeDimension is called when production rowTypeDimension is exited.
func (s *BaseFlinkSqlParserListener) ExitRowTypeDimension(ctx *RowTypeDimensionContext) {}

// EnterColumnConstraint is called when production columnConstraint is entered.
func (s *BaseFlinkSqlParserListener) EnterColumnConstraint(ctx *ColumnConstraintContext) {}

// ExitColumnConstraint is called when production columnConstraint is exited.
func (s *BaseFlinkSqlParserListener) ExitColumnConstraint(ctx *ColumnConstraintContext) {}

// EnterMetadataColumnDefinition is called when production metadataColumnDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterMetadataColumnDefinition(ctx *MetadataColumnDefinitionContext) {
}

// ExitMetadataColumnDefinition is called when production metadataColumnDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitMetadataColumnDefinition(ctx *MetadataColumnDefinitionContext) {
}

// EnterMetadataKey is called when production metadataKey is entered.
func (s *BaseFlinkSqlParserListener) EnterMetadataKey(ctx *MetadataKeyContext) {}

// ExitMetadataKey is called when production metadataKey is exited.
func (s *BaseFlinkSqlParserListener) ExitMetadataKey(ctx *MetadataKeyContext) {}

// EnterComputedColumnDefinition is called when production computedColumnDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterComputedColumnDefinition(ctx *ComputedColumnDefinitionContext) {
}

// ExitComputedColumnDefinition is called when production computedColumnDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitComputedColumnDefinition(ctx *ComputedColumnDefinitionContext) {
}

// EnterComputedColumnExpression is called when production computedColumnExpression is entered.
func (s *BaseFlinkSqlParserListener) EnterComputedColumnExpression(ctx *ComputedColumnExpressionContext) {
}

// ExitComputedColumnExpression is called when production computedColumnExpression is exited.
func (s *BaseFlinkSqlParserListener) ExitComputedColumnExpression(ctx *ComputedColumnExpressionContext) {
}

// EnterWatermarkDefinition is called when production watermarkDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterWatermarkDefinition(ctx *WatermarkDefinitionContext) {}

// ExitWatermarkDefinition is called when production watermarkDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitWatermarkDefinition(ctx *WatermarkDefinitionContext) {}

// EnterTableConstraint is called when production tableConstraint is entered.
func (s *BaseFlinkSqlParserListener) EnterTableConstraint(ctx *TableConstraintContext) {}

// ExitTableConstraint is called when production tableConstraint is exited.
func (s *BaseFlinkSqlParserListener) ExitTableConstraint(ctx *TableConstraintContext) {}

// EnterConstraintName is called when production constraintName is entered.
func (s *BaseFlinkSqlParserListener) EnterConstraintName(ctx *ConstraintNameContext) {}

// ExitConstraintName is called when production constraintName is exited.
func (s *BaseFlinkSqlParserListener) ExitConstraintName(ctx *ConstraintNameContext) {}

// EnterSelfDefinitionClause is called when production selfDefinitionClause is entered.
func (s *BaseFlinkSqlParserListener) EnterSelfDefinitionClause(ctx *SelfDefinitionClauseContext) {}

// ExitSelfDefinitionClause is called when production selfDefinitionClause is exited.
func (s *BaseFlinkSqlParserListener) ExitSelfDefinitionClause(ctx *SelfDefinitionClauseContext) {}

// EnterPartitionDefinition is called when production partitionDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterPartitionDefinition(ctx *PartitionDefinitionContext) {}

// ExitPartitionDefinition is called when production partitionDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitPartitionDefinition(ctx *PartitionDefinitionContext) {}

// EnterTransformList is called when production transformList is entered.
func (s *BaseFlinkSqlParserListener) EnterTransformList(ctx *TransformListContext) {}

// ExitTransformList is called when production transformList is exited.
func (s *BaseFlinkSqlParserListener) ExitTransformList(ctx *TransformListContext) {}

// EnterIdentityTransform is called when production identityTransform is entered.
func (s *BaseFlinkSqlParserListener) EnterIdentityTransform(ctx *IdentityTransformContext) {}

// ExitIdentityTransform is called when production identityTransform is exited.
func (s *BaseFlinkSqlParserListener) ExitIdentityTransform(ctx *IdentityTransformContext) {}

// EnterApplyTransform is called when production applyTransform is entered.
func (s *BaseFlinkSqlParserListener) EnterApplyTransform(ctx *ApplyTransformContext) {}

// ExitApplyTransform is called when production applyTransform is exited.
func (s *BaseFlinkSqlParserListener) ExitApplyTransform(ctx *ApplyTransformContext) {}

// EnterTransformArgument is called when production transformArgument is entered.
func (s *BaseFlinkSqlParserListener) EnterTransformArgument(ctx *TransformArgumentContext) {}

// ExitTransformArgument is called when production transformArgument is exited.
func (s *BaseFlinkSqlParserListener) ExitTransformArgument(ctx *TransformArgumentContext) {}

// EnterLikeDefinition is called when production likeDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterLikeDefinition(ctx *LikeDefinitionContext) {}

// ExitLikeDefinition is called when production likeDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitLikeDefinition(ctx *LikeDefinitionContext) {}

// EnterLikeOption is called when production likeOption is entered.
func (s *BaseFlinkSqlParserListener) EnterLikeOption(ctx *LikeOptionContext) {}

// ExitLikeOption is called when production likeOption is exited.
func (s *BaseFlinkSqlParserListener) ExitLikeOption(ctx *LikeOptionContext) {}

// EnterCreateCatalog is called when production createCatalog is entered.
func (s *BaseFlinkSqlParserListener) EnterCreateCatalog(ctx *CreateCatalogContext) {}

// ExitCreateCatalog is called when production createCatalog is exited.
func (s *BaseFlinkSqlParserListener) ExitCreateCatalog(ctx *CreateCatalogContext) {}

// EnterCreateDatabase is called when production createDatabase is entered.
func (s *BaseFlinkSqlParserListener) EnterCreateDatabase(ctx *CreateDatabaseContext) {}

// ExitCreateDatabase is called when production createDatabase is exited.
func (s *BaseFlinkSqlParserListener) ExitCreateDatabase(ctx *CreateDatabaseContext) {}

// EnterCreateView is called when production createView is entered.
func (s *BaseFlinkSqlParserListener) EnterCreateView(ctx *CreateViewContext) {}

// ExitCreateView is called when production createView is exited.
func (s *BaseFlinkSqlParserListener) ExitCreateView(ctx *CreateViewContext) {}

// EnterCreateFunction is called when production createFunction is entered.
func (s *BaseFlinkSqlParserListener) EnterCreateFunction(ctx *CreateFunctionContext) {}

// ExitCreateFunction is called when production createFunction is exited.
func (s *BaseFlinkSqlParserListener) ExitCreateFunction(ctx *CreateFunctionContext) {}

// EnterUsingClause is called when production usingClause is entered.
func (s *BaseFlinkSqlParserListener) EnterUsingClause(ctx *UsingClauseContext) {}

// ExitUsingClause is called when production usingClause is exited.
func (s *BaseFlinkSqlParserListener) ExitUsingClause(ctx *UsingClauseContext) {}

// EnterJarFileName is called when production jarFileName is entered.
func (s *BaseFlinkSqlParserListener) EnterJarFileName(ctx *JarFileNameContext) {}

// ExitJarFileName is called when production jarFileName is exited.
func (s *BaseFlinkSqlParserListener) ExitJarFileName(ctx *JarFileNameContext) {}

// EnterAlterTable is called when production alterTable is entered.
func (s *BaseFlinkSqlParserListener) EnterAlterTable(ctx *AlterTableContext) {}

// ExitAlterTable is called when production alterTable is exited.
func (s *BaseFlinkSqlParserListener) ExitAlterTable(ctx *AlterTableContext) {}

// EnterRenameDefinition is called when production renameDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterRenameDefinition(ctx *RenameDefinitionContext) {}

// ExitRenameDefinition is called when production renameDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitRenameDefinition(ctx *RenameDefinitionContext) {}

// EnterSetKeyValueDefinition is called when production setKeyValueDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterSetKeyValueDefinition(ctx *SetKeyValueDefinitionContext) {}

// ExitSetKeyValueDefinition is called when production setKeyValueDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitSetKeyValueDefinition(ctx *SetKeyValueDefinitionContext) {}

// EnterAddConstraint is called when production addConstraint is entered.
func (s *BaseFlinkSqlParserListener) EnterAddConstraint(ctx *AddConstraintContext) {}

// ExitAddConstraint is called when production addConstraint is exited.
func (s *BaseFlinkSqlParserListener) ExitAddConstraint(ctx *AddConstraintContext) {}

// EnterDropConstraint is called when production dropConstraint is entered.
func (s *BaseFlinkSqlParserListener) EnterDropConstraint(ctx *DropConstraintContext) {}

// ExitDropConstraint is called when production dropConstraint is exited.
func (s *BaseFlinkSqlParserListener) ExitDropConstraint(ctx *DropConstraintContext) {}

// EnterAddUnique is called when production addUnique is entered.
func (s *BaseFlinkSqlParserListener) EnterAddUnique(ctx *AddUniqueContext) {}

// ExitAddUnique is called when production addUnique is exited.
func (s *BaseFlinkSqlParserListener) ExitAddUnique(ctx *AddUniqueContext) {}

// EnterNotForced is called when production notForced is entered.
func (s *BaseFlinkSqlParserListener) EnterNotForced(ctx *NotForcedContext) {}

// ExitNotForced is called when production notForced is exited.
func (s *BaseFlinkSqlParserListener) ExitNotForced(ctx *NotForcedContext) {}

// EnterAlterView is called when production alterView is entered.
func (s *BaseFlinkSqlParserListener) EnterAlterView(ctx *AlterViewContext) {}

// ExitAlterView is called when production alterView is exited.
func (s *BaseFlinkSqlParserListener) ExitAlterView(ctx *AlterViewContext) {}

// EnterAlterDatabase is called when production alterDatabase is entered.
func (s *BaseFlinkSqlParserListener) EnterAlterDatabase(ctx *AlterDatabaseContext) {}

// ExitAlterDatabase is called when production alterDatabase is exited.
func (s *BaseFlinkSqlParserListener) ExitAlterDatabase(ctx *AlterDatabaseContext) {}

// EnterAlterFunction is called when production alterFunction is entered.
func (s *BaseFlinkSqlParserListener) EnterAlterFunction(ctx *AlterFunctionContext) {}

// ExitAlterFunction is called when production alterFunction is exited.
func (s *BaseFlinkSqlParserListener) ExitAlterFunction(ctx *AlterFunctionContext) {}

// EnterDropCatalog is called when production dropCatalog is entered.
func (s *BaseFlinkSqlParserListener) EnterDropCatalog(ctx *DropCatalogContext) {}

// ExitDropCatalog is called when production dropCatalog is exited.
func (s *BaseFlinkSqlParserListener) ExitDropCatalog(ctx *DropCatalogContext) {}

// EnterDropTable is called when production dropTable is entered.
func (s *BaseFlinkSqlParserListener) EnterDropTable(ctx *DropTableContext) {}

// ExitDropTable is called when production dropTable is exited.
func (s *BaseFlinkSqlParserListener) ExitDropTable(ctx *DropTableContext) {}

// EnterDropDatabase is called when production dropDatabase is entered.
func (s *BaseFlinkSqlParserListener) EnterDropDatabase(ctx *DropDatabaseContext) {}

// ExitDropDatabase is called when production dropDatabase is exited.
func (s *BaseFlinkSqlParserListener) ExitDropDatabase(ctx *DropDatabaseContext) {}

// EnterDropView is called when production dropView is entered.
func (s *BaseFlinkSqlParserListener) EnterDropView(ctx *DropViewContext) {}

// ExitDropView is called when production dropView is exited.
func (s *BaseFlinkSqlParserListener) ExitDropView(ctx *DropViewContext) {}

// EnterDropFunction is called when production dropFunction is entered.
func (s *BaseFlinkSqlParserListener) EnterDropFunction(ctx *DropFunctionContext) {}

// ExitDropFunction is called when production dropFunction is exited.
func (s *BaseFlinkSqlParserListener) ExitDropFunction(ctx *DropFunctionContext) {}

// EnterInsertStatement is called when production insertStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterInsertStatement(ctx *InsertStatementContext) {}

// ExitInsertStatement is called when production insertStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitInsertStatement(ctx *InsertStatementContext) {}

// EnterInsertSimpleStatement is called when production insertSimpleStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterInsertSimpleStatement(ctx *InsertSimpleStatementContext) {}

// ExitInsertSimpleStatement is called when production insertSimpleStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitInsertSimpleStatement(ctx *InsertSimpleStatementContext) {}

// EnterInsertPartitionDefinition is called when production insertPartitionDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterInsertPartitionDefinition(ctx *InsertPartitionDefinitionContext) {
}

// ExitInsertPartitionDefinition is called when production insertPartitionDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitInsertPartitionDefinition(ctx *InsertPartitionDefinitionContext) {
}

// EnterValuesDefinition is called when production valuesDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterValuesDefinition(ctx *ValuesDefinitionContext) {}

// ExitValuesDefinition is called when production valuesDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitValuesDefinition(ctx *ValuesDefinitionContext) {}

// EnterValuesRowDefinition is called when production valuesRowDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterValuesRowDefinition(ctx *ValuesRowDefinitionContext) {}

// ExitValuesRowDefinition is called when production valuesRowDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitValuesRowDefinition(ctx *ValuesRowDefinitionContext) {}

// EnterInsertMulStatementCompatibility is called when production insertMulStatementCompatibility is entered.
func (s *BaseFlinkSqlParserListener) EnterInsertMulStatementCompatibility(ctx *InsertMulStatementCompatibilityContext) {
}

// ExitInsertMulStatementCompatibility is called when production insertMulStatementCompatibility is exited.
func (s *BaseFlinkSqlParserListener) ExitInsertMulStatementCompatibility(ctx *InsertMulStatementCompatibilityContext) {
}

// EnterInsertMulStatement is called when production insertMulStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterInsertMulStatement(ctx *InsertMulStatementContext) {}

// ExitInsertMulStatement is called when production insertMulStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitInsertMulStatement(ctx *InsertMulStatementContext) {}

// EnterQueryStatement is called when production queryStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterQueryStatement(ctx *QueryStatementContext) {}

// ExitQueryStatement is called when production queryStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitQueryStatement(ctx *QueryStatementContext) {}

// EnterValuesClause is called when production valuesClause is entered.
func (s *BaseFlinkSqlParserListener) EnterValuesClause(ctx *ValuesClauseContext) {}

// ExitValuesClause is called when production valuesClause is exited.
func (s *BaseFlinkSqlParserListener) ExitValuesClause(ctx *ValuesClauseContext) {}

// EnterWithClause is called when production withClause is entered.
func (s *BaseFlinkSqlParserListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BaseFlinkSqlParserListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterWithItem is called when production withItem is entered.
func (s *BaseFlinkSqlParserListener) EnterWithItem(ctx *WithItemContext) {}

// ExitWithItem is called when production withItem is exited.
func (s *BaseFlinkSqlParserListener) ExitWithItem(ctx *WithItemContext) {}

// EnterWithItemName is called when production withItemName is entered.
func (s *BaseFlinkSqlParserListener) EnterWithItemName(ctx *WithItemNameContext) {}

// ExitWithItemName is called when production withItemName is exited.
func (s *BaseFlinkSqlParserListener) ExitWithItemName(ctx *WithItemNameContext) {}

// EnterSelectStatement is called when production selectStatement is entered.
func (s *BaseFlinkSqlParserListener) EnterSelectStatement(ctx *SelectStatementContext) {}

// ExitSelectStatement is called when production selectStatement is exited.
func (s *BaseFlinkSqlParserListener) ExitSelectStatement(ctx *SelectStatementContext) {}

// EnterSelectClause is called when production selectClause is entered.
func (s *BaseFlinkSqlParserListener) EnterSelectClause(ctx *SelectClauseContext) {}

// ExitSelectClause is called when production selectClause is exited.
func (s *BaseFlinkSqlParserListener) ExitSelectClause(ctx *SelectClauseContext) {}

// EnterProjectItemDefinition is called when production projectItemDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterProjectItemDefinition(ctx *ProjectItemDefinitionContext) {}

// ExitProjectItemDefinition is called when production projectItemDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitProjectItemDefinition(ctx *ProjectItemDefinitionContext) {}

// EnterOverWindowItem is called when production overWindowItem is entered.
func (s *BaseFlinkSqlParserListener) EnterOverWindowItem(ctx *OverWindowItemContext) {}

// ExitOverWindowItem is called when production overWindowItem is exited.
func (s *BaseFlinkSqlParserListener) ExitOverWindowItem(ctx *OverWindowItemContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseFlinkSqlParserListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseFlinkSqlParserListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterTableExpression is called when production tableExpression is entered.
func (s *BaseFlinkSqlParserListener) EnterTableExpression(ctx *TableExpressionContext) {}

// ExitTableExpression is called when production tableExpression is exited.
func (s *BaseFlinkSqlParserListener) ExitTableExpression(ctx *TableExpressionContext) {}

// EnterTableReference is called when production tableReference is entered.
func (s *BaseFlinkSqlParserListener) EnterTableReference(ctx *TableReferenceContext) {}

// ExitTableReference is called when production tableReference is exited.
func (s *BaseFlinkSqlParserListener) ExitTableReference(ctx *TableReferenceContext) {}

// EnterTablePrimary is called when production tablePrimary is entered.
func (s *BaseFlinkSqlParserListener) EnterTablePrimary(ctx *TablePrimaryContext) {}

// ExitTablePrimary is called when production tablePrimary is exited.
func (s *BaseFlinkSqlParserListener) ExitTablePrimary(ctx *TablePrimaryContext) {}

// EnterSystemTimePeriod is called when production systemTimePeriod is entered.
func (s *BaseFlinkSqlParserListener) EnterSystemTimePeriod(ctx *SystemTimePeriodContext) {}

// ExitSystemTimePeriod is called when production systemTimePeriod is exited.
func (s *BaseFlinkSqlParserListener) ExitSystemTimePeriod(ctx *SystemTimePeriodContext) {}

// EnterDateTimeExpression is called when production dateTimeExpression is entered.
func (s *BaseFlinkSqlParserListener) EnterDateTimeExpression(ctx *DateTimeExpressionContext) {}

// ExitDateTimeExpression is called when production dateTimeExpression is exited.
func (s *BaseFlinkSqlParserListener) ExitDateTimeExpression(ctx *DateTimeExpressionContext) {}

// EnterInlineDataValueClause is called when production inlineDataValueClause is entered.
func (s *BaseFlinkSqlParserListener) EnterInlineDataValueClause(ctx *InlineDataValueClauseContext) {}

// ExitInlineDataValueClause is called when production inlineDataValueClause is exited.
func (s *BaseFlinkSqlParserListener) ExitInlineDataValueClause(ctx *InlineDataValueClauseContext) {}

// EnterWindowTVFClause is called when production windowTVFClause is entered.
func (s *BaseFlinkSqlParserListener) EnterWindowTVFClause(ctx *WindowTVFClauseContext) {}

// ExitWindowTVFClause is called when production windowTVFClause is exited.
func (s *BaseFlinkSqlParserListener) ExitWindowTVFClause(ctx *WindowTVFClauseContext) {}

// EnterWindowTVFExpression is called when production windowTVFExpression is entered.
func (s *BaseFlinkSqlParserListener) EnterWindowTVFExpression(ctx *WindowTVFExpressionContext) {}

// ExitWindowTVFExpression is called when production windowTVFExpression is exited.
func (s *BaseFlinkSqlParserListener) ExitWindowTVFExpression(ctx *WindowTVFExpressionContext) {}

// EnterWindowTVFName is called when production windowTVFName is entered.
func (s *BaseFlinkSqlParserListener) EnterWindowTVFName(ctx *WindowTVFNameContext) {}

// ExitWindowTVFName is called when production windowTVFName is exited.
func (s *BaseFlinkSqlParserListener) ExitWindowTVFName(ctx *WindowTVFNameContext) {}

// EnterWindowTVFParam is called when production windowTVFParam is entered.
func (s *BaseFlinkSqlParserListener) EnterWindowTVFParam(ctx *WindowTVFParamContext) {}

// ExitWindowTVFParam is called when production windowTVFParam is exited.
func (s *BaseFlinkSqlParserListener) ExitWindowTVFParam(ctx *WindowTVFParamContext) {}

// EnterTimeIntervalParamName is called when production timeIntervalParamName is entered.
func (s *BaseFlinkSqlParserListener) EnterTimeIntervalParamName(ctx *TimeIntervalParamNameContext) {}

// ExitTimeIntervalParamName is called when production timeIntervalParamName is exited.
func (s *BaseFlinkSqlParserListener) ExitTimeIntervalParamName(ctx *TimeIntervalParamNameContext) {}

// EnterColumnDescriptor is called when production columnDescriptor is entered.
func (s *BaseFlinkSqlParserListener) EnterColumnDescriptor(ctx *ColumnDescriptorContext) {}

// ExitColumnDescriptor is called when production columnDescriptor is exited.
func (s *BaseFlinkSqlParserListener) ExitColumnDescriptor(ctx *ColumnDescriptorContext) {}

// EnterJoinCondition is called when production joinCondition is entered.
func (s *BaseFlinkSqlParserListener) EnterJoinCondition(ctx *JoinConditionContext) {}

// ExitJoinCondition is called when production joinCondition is exited.
func (s *BaseFlinkSqlParserListener) ExitJoinCondition(ctx *JoinConditionContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseFlinkSqlParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseFlinkSqlParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseFlinkSqlParserListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseFlinkSqlParserListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterGroupItemDefinition is called when production groupItemDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterGroupItemDefinition(ctx *GroupItemDefinitionContext) {}

// ExitGroupItemDefinition is called when production groupItemDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitGroupItemDefinition(ctx *GroupItemDefinitionContext) {}

// EnterGroupingSets is called when production groupingSets is entered.
func (s *BaseFlinkSqlParserListener) EnterGroupingSets(ctx *GroupingSetsContext) {}

// ExitGroupingSets is called when production groupingSets is exited.
func (s *BaseFlinkSqlParserListener) ExitGroupingSets(ctx *GroupingSetsContext) {}

// EnterGroupingSetsNotationName is called when production groupingSetsNotationName is entered.
func (s *BaseFlinkSqlParserListener) EnterGroupingSetsNotationName(ctx *GroupingSetsNotationNameContext) {
}

// ExitGroupingSetsNotationName is called when production groupingSetsNotationName is exited.
func (s *BaseFlinkSqlParserListener) ExitGroupingSetsNotationName(ctx *GroupingSetsNotationNameContext) {
}

// EnterGroupWindowFunction is called when production groupWindowFunction is entered.
func (s *BaseFlinkSqlParserListener) EnterGroupWindowFunction(ctx *GroupWindowFunctionContext) {}

// ExitGroupWindowFunction is called when production groupWindowFunction is exited.
func (s *BaseFlinkSqlParserListener) ExitGroupWindowFunction(ctx *GroupWindowFunctionContext) {}

// EnterGroupWindowFunctionName is called when production groupWindowFunctionName is entered.
func (s *BaseFlinkSqlParserListener) EnterGroupWindowFunctionName(ctx *GroupWindowFunctionNameContext) {
}

// ExitGroupWindowFunctionName is called when production groupWindowFunctionName is exited.
func (s *BaseFlinkSqlParserListener) ExitGroupWindowFunctionName(ctx *GroupWindowFunctionNameContext) {
}

// EnterTimeAttrColumn is called when production timeAttrColumn is entered.
func (s *BaseFlinkSqlParserListener) EnterTimeAttrColumn(ctx *TimeAttrColumnContext) {}

// ExitTimeAttrColumn is called when production timeAttrColumn is exited.
func (s *BaseFlinkSqlParserListener) ExitTimeAttrColumn(ctx *TimeAttrColumnContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseFlinkSqlParserListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseFlinkSqlParserListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterWindowClause is called when production windowClause is entered.
func (s *BaseFlinkSqlParserListener) EnterWindowClause(ctx *WindowClauseContext) {}

// ExitWindowClause is called when production windowClause is exited.
func (s *BaseFlinkSqlParserListener) ExitWindowClause(ctx *WindowClauseContext) {}

// EnterNamedWindow is called when production namedWindow is entered.
func (s *BaseFlinkSqlParserListener) EnterNamedWindow(ctx *NamedWindowContext) {}

// ExitNamedWindow is called when production namedWindow is exited.
func (s *BaseFlinkSqlParserListener) ExitNamedWindow(ctx *NamedWindowContext) {}

// EnterWindowSpec is called when production windowSpec is entered.
func (s *BaseFlinkSqlParserListener) EnterWindowSpec(ctx *WindowSpecContext) {}

// ExitWindowSpec is called when production windowSpec is exited.
func (s *BaseFlinkSqlParserListener) ExitWindowSpec(ctx *WindowSpecContext) {}

// EnterMatchRecognizeClause is called when production matchRecognizeClause is entered.
func (s *BaseFlinkSqlParserListener) EnterMatchRecognizeClause(ctx *MatchRecognizeClauseContext) {}

// ExitMatchRecognizeClause is called when production matchRecognizeClause is exited.
func (s *BaseFlinkSqlParserListener) ExitMatchRecognizeClause(ctx *MatchRecognizeClauseContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseFlinkSqlParserListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseFlinkSqlParserListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterOrderItemDefinition is called when production orderItemDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterOrderItemDefinition(ctx *OrderItemDefinitionContext) {}

// ExitOrderItemDefinition is called when production orderItemDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitOrderItemDefinition(ctx *OrderItemDefinitionContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseFlinkSqlParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseFlinkSqlParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterPartitionByClause is called when production partitionByClause is entered.
func (s *BaseFlinkSqlParserListener) EnterPartitionByClause(ctx *PartitionByClauseContext) {}

// ExitPartitionByClause is called when production partitionByClause is exited.
func (s *BaseFlinkSqlParserListener) ExitPartitionByClause(ctx *PartitionByClauseContext) {}

// EnterQuantifiers is called when production quantifiers is entered.
func (s *BaseFlinkSqlParserListener) EnterQuantifiers(ctx *QuantifiersContext) {}

// ExitQuantifiers is called when production quantifiers is exited.
func (s *BaseFlinkSqlParserListener) ExitQuantifiers(ctx *QuantifiersContext) {}

// EnterMeasuresClause is called when production measuresClause is entered.
func (s *BaseFlinkSqlParserListener) EnterMeasuresClause(ctx *MeasuresClauseContext) {}

// ExitMeasuresClause is called when production measuresClause is exited.
func (s *BaseFlinkSqlParserListener) ExitMeasuresClause(ctx *MeasuresClauseContext) {}

// EnterPatternDefinition is called when production patternDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterPatternDefinition(ctx *PatternDefinitionContext) {}

// ExitPatternDefinition is called when production patternDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitPatternDefinition(ctx *PatternDefinitionContext) {}

// EnterPatternVariable is called when production patternVariable is entered.
func (s *BaseFlinkSqlParserListener) EnterPatternVariable(ctx *PatternVariableContext) {}

// ExitPatternVariable is called when production patternVariable is exited.
func (s *BaseFlinkSqlParserListener) ExitPatternVariable(ctx *PatternVariableContext) {}

// EnterOutputMode is called when production outputMode is entered.
func (s *BaseFlinkSqlParserListener) EnterOutputMode(ctx *OutputModeContext) {}

// ExitOutputMode is called when production outputMode is exited.
func (s *BaseFlinkSqlParserListener) ExitOutputMode(ctx *OutputModeContext) {}

// EnterAfterMatchStrategy is called when production afterMatchStrategy is entered.
func (s *BaseFlinkSqlParserListener) EnterAfterMatchStrategy(ctx *AfterMatchStrategyContext) {}

// ExitAfterMatchStrategy is called when production afterMatchStrategy is exited.
func (s *BaseFlinkSqlParserListener) ExitAfterMatchStrategy(ctx *AfterMatchStrategyContext) {}

// EnterPatternVariablesDefinition is called when production patternVariablesDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterPatternVariablesDefinition(ctx *PatternVariablesDefinitionContext) {
}

// ExitPatternVariablesDefinition is called when production patternVariablesDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitPatternVariablesDefinition(ctx *PatternVariablesDefinitionContext) {
}

// EnterWindowFrame is called when production windowFrame is entered.
func (s *BaseFlinkSqlParserListener) EnterWindowFrame(ctx *WindowFrameContext) {}

// ExitWindowFrame is called when production windowFrame is exited.
func (s *BaseFlinkSqlParserListener) ExitWindowFrame(ctx *WindowFrameContext) {}

// EnterFrameBound is called when production frameBound is entered.
func (s *BaseFlinkSqlParserListener) EnterFrameBound(ctx *FrameBoundContext) {}

// ExitFrameBound is called when production frameBound is exited.
func (s *BaseFlinkSqlParserListener) ExitFrameBound(ctx *FrameBoundContext) {}

// EnterWithinClause is called when production withinClause is entered.
func (s *BaseFlinkSqlParserListener) EnterWithinClause(ctx *WithinClauseContext) {}

// ExitWithinClause is called when production withinClause is exited.
func (s *BaseFlinkSqlParserListener) ExitWithinClause(ctx *WithinClauseContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFlinkSqlParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFlinkSqlParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLogicalNot is called when production logicalNot is entered.
func (s *BaseFlinkSqlParserListener) EnterLogicalNot(ctx *LogicalNotContext) {}

// ExitLogicalNot is called when production logicalNot is exited.
func (s *BaseFlinkSqlParserListener) ExitLogicalNot(ctx *LogicalNotContext) {}

// EnterPredicated is called when production predicated is entered.
func (s *BaseFlinkSqlParserListener) EnterPredicated(ctx *PredicatedContext) {}

// ExitPredicated is called when production predicated is exited.
func (s *BaseFlinkSqlParserListener) ExitPredicated(ctx *PredicatedContext) {}

// EnterExists is called when production exists is entered.
func (s *BaseFlinkSqlParserListener) EnterExists(ctx *ExistsContext) {}

// ExitExists is called when production exists is exited.
func (s *BaseFlinkSqlParserListener) ExitExists(ctx *ExistsContext) {}

// EnterLogicalNested is called when production logicalNested is entered.
func (s *BaseFlinkSqlParserListener) EnterLogicalNested(ctx *LogicalNestedContext) {}

// ExitLogicalNested is called when production logicalNested is exited.
func (s *BaseFlinkSqlParserListener) ExitLogicalNested(ctx *LogicalNestedContext) {}

// EnterLogicalBinary is called when production logicalBinary is entered.
func (s *BaseFlinkSqlParserListener) EnterLogicalBinary(ctx *LogicalBinaryContext) {}

// ExitLogicalBinary is called when production logicalBinary is exited.
func (s *BaseFlinkSqlParserListener) ExitLogicalBinary(ctx *LogicalBinaryContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseFlinkSqlParserListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseFlinkSqlParserListener) ExitPredicate(ctx *PredicateContext) {}

// EnterLikePredicate is called when production likePredicate is entered.
func (s *BaseFlinkSqlParserListener) EnterLikePredicate(ctx *LikePredicateContext) {}

// ExitLikePredicate is called when production likePredicate is exited.
func (s *BaseFlinkSqlParserListener) ExitLikePredicate(ctx *LikePredicateContext) {}

// EnterValueExpressionDefault is called when production valueExpressionDefault is entered.
func (s *BaseFlinkSqlParserListener) EnterValueExpressionDefault(ctx *ValueExpressionDefaultContext) {
}

// ExitValueExpressionDefault is called when production valueExpressionDefault is exited.
func (s *BaseFlinkSqlParserListener) ExitValueExpressionDefault(ctx *ValueExpressionDefaultContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseFlinkSqlParserListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseFlinkSqlParserListener) ExitComparison(ctx *ComparisonContext) {}

// EnterArithmeticBinary is called when production arithmeticBinary is entered.
func (s *BaseFlinkSqlParserListener) EnterArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// ExitArithmeticBinary is called when production arithmeticBinary is exited.
func (s *BaseFlinkSqlParserListener) ExitArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// EnterArithmeticUnary is called when production arithmeticUnary is entered.
func (s *BaseFlinkSqlParserListener) EnterArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// ExitArithmeticUnary is called when production arithmeticUnary is exited.
func (s *BaseFlinkSqlParserListener) ExitArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// EnterFunctionCallExpression is called when production functionCallExpression is entered.
func (s *BaseFlinkSqlParserListener) EnterFunctionCallExpression(ctx *FunctionCallExpressionContext) {
}

// ExitFunctionCallExpression is called when production functionCallExpression is exited.
func (s *BaseFlinkSqlParserListener) ExitFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// EnterDereference is called when production dereference is entered.
func (s *BaseFlinkSqlParserListener) EnterDereference(ctx *DereferenceContext) {}

// ExitDereference is called when production dereference is exited.
func (s *BaseFlinkSqlParserListener) ExitDereference(ctx *DereferenceContext) {}

// EnterSimpleCase is called when production simpleCase is entered.
func (s *BaseFlinkSqlParserListener) EnterSimpleCase(ctx *SimpleCaseContext) {}

// ExitSimpleCase is called when production simpleCase is exited.
func (s *BaseFlinkSqlParserListener) ExitSimpleCase(ctx *SimpleCaseContext) {}

// EnterColumnReference is called when production columnReference is entered.
func (s *BaseFlinkSqlParserListener) EnterColumnReference(ctx *ColumnReferenceContext) {}

// ExitColumnReference is called when production columnReference is exited.
func (s *BaseFlinkSqlParserListener) ExitColumnReference(ctx *ColumnReferenceContext) {}

// EnterLast is called when production last is entered.
func (s *BaseFlinkSqlParserListener) EnterLast(ctx *LastContext) {}

// ExitLast is called when production last is exited.
func (s *BaseFlinkSqlParserListener) ExitLast(ctx *LastContext) {}

// EnterStar is called when production star is entered.
func (s *BaseFlinkSqlParserListener) EnterStar(ctx *StarContext) {}

// ExitStar is called when production star is exited.
func (s *BaseFlinkSqlParserListener) ExitStar(ctx *StarContext) {}

// EnterSubscript is called when production subscript is entered.
func (s *BaseFlinkSqlParserListener) EnterSubscript(ctx *SubscriptContext) {}

// ExitSubscript is called when production subscript is exited.
func (s *BaseFlinkSqlParserListener) ExitSubscript(ctx *SubscriptContext) {}

// EnterSubqueryExpression is called when production subqueryExpression is entered.
func (s *BaseFlinkSqlParserListener) EnterSubqueryExpression(ctx *SubqueryExpressionContext) {}

// ExitSubqueryExpression is called when production subqueryExpression is exited.
func (s *BaseFlinkSqlParserListener) ExitSubqueryExpression(ctx *SubqueryExpressionContext) {}

// EnterCast is called when production cast is entered.
func (s *BaseFlinkSqlParserListener) EnterCast(ctx *CastContext) {}

// ExitCast is called when production cast is exited.
func (s *BaseFlinkSqlParserListener) ExitCast(ctx *CastContext) {}

// EnterConstantDefault is called when production constantDefault is entered.
func (s *BaseFlinkSqlParserListener) EnterConstantDefault(ctx *ConstantDefaultContext) {}

// ExitConstantDefault is called when production constantDefault is exited.
func (s *BaseFlinkSqlParserListener) ExitConstantDefault(ctx *ConstantDefaultContext) {}

// EnterParenthesizedExpression is called when production parenthesizedExpression is entered.
func (s *BaseFlinkSqlParserListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {
}

// ExitParenthesizedExpression is called when production parenthesizedExpression is exited.
func (s *BaseFlinkSqlParserListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {
}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseFlinkSqlParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseFlinkSqlParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterSearchedCase is called when production searchedCase is entered.
func (s *BaseFlinkSqlParserListener) EnterSearchedCase(ctx *SearchedCaseContext) {}

// ExitSearchedCase is called when production searchedCase is exited.
func (s *BaseFlinkSqlParserListener) ExitSearchedCase(ctx *SearchedCaseContext) {}

// EnterPosition is called when production position is entered.
func (s *BaseFlinkSqlParserListener) EnterPosition(ctx *PositionContext) {}

// ExitPosition is called when production position is exited.
func (s *BaseFlinkSqlParserListener) ExitPosition(ctx *PositionContext) {}

// EnterFirst is called when production first is entered.
func (s *BaseFlinkSqlParserListener) EnterFirst(ctx *FirstContext) {}

// ExitFirst is called when production first is exited.
func (s *BaseFlinkSqlParserListener) ExitFirst(ctx *FirstContext) {}

// EnterFunctionNameCreate is called when production functionNameCreate is entered.
func (s *BaseFlinkSqlParserListener) EnterFunctionNameCreate(ctx *FunctionNameCreateContext) {}

// ExitFunctionNameCreate is called when production functionNameCreate is exited.
func (s *BaseFlinkSqlParserListener) ExitFunctionNameCreate(ctx *FunctionNameCreateContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BaseFlinkSqlParserListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BaseFlinkSqlParserListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterFunctionNameAndParams is called when production functionNameAndParams is entered.
func (s *BaseFlinkSqlParserListener) EnterFunctionNameAndParams(ctx *FunctionNameAndParamsContext) {}

// ExitFunctionNameAndParams is called when production functionNameAndParams is exited.
func (s *BaseFlinkSqlParserListener) ExitFunctionNameAndParams(ctx *FunctionNameAndParamsContext) {}

// EnterFunctionNameWithParams is called when production functionNameWithParams is entered.
func (s *BaseFlinkSqlParserListener) EnterFunctionNameWithParams(ctx *FunctionNameWithParamsContext) {
}

// ExitFunctionNameWithParams is called when production functionNameWithParams is exited.
func (s *BaseFlinkSqlParserListener) ExitFunctionNameWithParams(ctx *FunctionNameWithParamsContext) {}

// EnterFunctionParam is called when production functionParam is entered.
func (s *BaseFlinkSqlParserListener) EnterFunctionParam(ctx *FunctionParamContext) {}

// ExitFunctionParam is called when production functionParam is exited.
func (s *BaseFlinkSqlParserListener) ExitFunctionParam(ctx *FunctionParamContext) {}

// EnterDereferenceDefinition is called when production dereferenceDefinition is entered.
func (s *BaseFlinkSqlParserListener) EnterDereferenceDefinition(ctx *DereferenceDefinitionContext) {}

// ExitDereferenceDefinition is called when production dereferenceDefinition is exited.
func (s *BaseFlinkSqlParserListener) ExitDereferenceDefinition(ctx *DereferenceDefinitionContext) {}

// EnterCorrelationName is called when production correlationName is entered.
func (s *BaseFlinkSqlParserListener) EnterCorrelationName(ctx *CorrelationNameContext) {}

// ExitCorrelationName is called when production correlationName is exited.
func (s *BaseFlinkSqlParserListener) ExitCorrelationName(ctx *CorrelationNameContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseFlinkSqlParserListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseFlinkSqlParserListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterTimeIntervalExpression is called when production timeIntervalExpression is entered.
func (s *BaseFlinkSqlParserListener) EnterTimeIntervalExpression(ctx *TimeIntervalExpressionContext) {
}

// ExitTimeIntervalExpression is called when production timeIntervalExpression is exited.
func (s *BaseFlinkSqlParserListener) ExitTimeIntervalExpression(ctx *TimeIntervalExpressionContext) {}

// EnterErrorCapturingMultiUnitsInterval is called when production errorCapturingMultiUnitsInterval is entered.
func (s *BaseFlinkSqlParserListener) EnterErrorCapturingMultiUnitsInterval(ctx *ErrorCapturingMultiUnitsIntervalContext) {
}

// ExitErrorCapturingMultiUnitsInterval is called when production errorCapturingMultiUnitsInterval is exited.
func (s *BaseFlinkSqlParserListener) ExitErrorCapturingMultiUnitsInterval(ctx *ErrorCapturingMultiUnitsIntervalContext) {
}

// EnterMultiUnitsInterval is called when production multiUnitsInterval is entered.
func (s *BaseFlinkSqlParserListener) EnterMultiUnitsInterval(ctx *MultiUnitsIntervalContext) {}

// ExitMultiUnitsInterval is called when production multiUnitsInterval is exited.
func (s *BaseFlinkSqlParserListener) ExitMultiUnitsInterval(ctx *MultiUnitsIntervalContext) {}

// EnterErrorCapturingUnitToUnitInterval is called when production errorCapturingUnitToUnitInterval is entered.
func (s *BaseFlinkSqlParserListener) EnterErrorCapturingUnitToUnitInterval(ctx *ErrorCapturingUnitToUnitIntervalContext) {
}

// ExitErrorCapturingUnitToUnitInterval is called when production errorCapturingUnitToUnitInterval is exited.
func (s *BaseFlinkSqlParserListener) ExitErrorCapturingUnitToUnitInterval(ctx *ErrorCapturingUnitToUnitIntervalContext) {
}

// EnterUnitToUnitInterval is called when production unitToUnitInterval is entered.
func (s *BaseFlinkSqlParserListener) EnterUnitToUnitInterval(ctx *UnitToUnitIntervalContext) {}

// ExitUnitToUnitInterval is called when production unitToUnitInterval is exited.
func (s *BaseFlinkSqlParserListener) ExitUnitToUnitInterval(ctx *UnitToUnitIntervalContext) {}

// EnterIntervalValue is called when production intervalValue is entered.
func (s *BaseFlinkSqlParserListener) EnterIntervalValue(ctx *IntervalValueContext) {}

// ExitIntervalValue is called when production intervalValue is exited.
func (s *BaseFlinkSqlParserListener) ExitIntervalValue(ctx *IntervalValueContext) {}

// EnterTableAlias is called when production tableAlias is entered.
func (s *BaseFlinkSqlParserListener) EnterTableAlias(ctx *TableAliasContext) {}

// ExitTableAlias is called when production tableAlias is exited.
func (s *BaseFlinkSqlParserListener) ExitTableAlias(ctx *TableAliasContext) {}

// EnterErrorCapturingIdentifier is called when production errorCapturingIdentifier is entered.
func (s *BaseFlinkSqlParserListener) EnterErrorCapturingIdentifier(ctx *ErrorCapturingIdentifierContext) {
}

// ExitErrorCapturingIdentifier is called when production errorCapturingIdentifier is exited.
func (s *BaseFlinkSqlParserListener) ExitErrorCapturingIdentifier(ctx *ErrorCapturingIdentifierContext) {
}

// EnterErrorIdent is called when production errorIdent is entered.
func (s *BaseFlinkSqlParserListener) EnterErrorIdent(ctx *ErrorIdentContext) {}

// ExitErrorIdent is called when production errorIdent is exited.
func (s *BaseFlinkSqlParserListener) ExitErrorIdent(ctx *ErrorIdentContext) {}

// EnterRealIdent is called when production realIdent is entered.
func (s *BaseFlinkSqlParserListener) EnterRealIdent(ctx *RealIdentContext) {}

// ExitRealIdent is called when production realIdent is exited.
func (s *BaseFlinkSqlParserListener) ExitRealIdent(ctx *RealIdentContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseFlinkSqlParserListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseFlinkSqlParserListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterIdentifierSeq is called when production identifierSeq is entered.
func (s *BaseFlinkSqlParserListener) EnterIdentifierSeq(ctx *IdentifierSeqContext) {}

// ExitIdentifierSeq is called when production identifierSeq is exited.
func (s *BaseFlinkSqlParserListener) ExitIdentifierSeq(ctx *IdentifierSeqContext) {}

// EnterUnquotedIdentifierAlternative is called when production unquotedIdentifierAlternative is entered.
func (s *BaseFlinkSqlParserListener) EnterUnquotedIdentifierAlternative(ctx *UnquotedIdentifierAlternativeContext) {
}

// ExitUnquotedIdentifierAlternative is called when production unquotedIdentifierAlternative is exited.
func (s *BaseFlinkSqlParserListener) ExitUnquotedIdentifierAlternative(ctx *UnquotedIdentifierAlternativeContext) {
}

// EnterQuotedIdentifierAlternative is called when production quotedIdentifierAlternative is entered.
func (s *BaseFlinkSqlParserListener) EnterQuotedIdentifierAlternative(ctx *QuotedIdentifierAlternativeContext) {
}

// ExitQuotedIdentifierAlternative is called when production quotedIdentifierAlternative is exited.
func (s *BaseFlinkSqlParserListener) ExitQuotedIdentifierAlternative(ctx *QuotedIdentifierAlternativeContext) {
}

// EnterNonReservedKeywordsAlternative is called when production nonReservedKeywordsAlternative is entered.
func (s *BaseFlinkSqlParserListener) EnterNonReservedKeywordsAlternative(ctx *NonReservedKeywordsAlternativeContext) {
}

// ExitNonReservedKeywordsAlternative is called when production nonReservedKeywordsAlternative is exited.
func (s *BaseFlinkSqlParserListener) ExitNonReservedKeywordsAlternative(ctx *NonReservedKeywordsAlternativeContext) {
}

// EnterUnquotedIdentifier is called when production unquotedIdentifier is entered.
func (s *BaseFlinkSqlParserListener) EnterUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// ExitUnquotedIdentifier is called when production unquotedIdentifier is exited.
func (s *BaseFlinkSqlParserListener) ExitUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// EnterQuotedIdentifier is called when production quotedIdentifier is entered.
func (s *BaseFlinkSqlParserListener) EnterQuotedIdentifier(ctx *QuotedIdentifierContext) {}

// ExitQuotedIdentifier is called when production quotedIdentifier is exited.
func (s *BaseFlinkSqlParserListener) ExitQuotedIdentifier(ctx *QuotedIdentifierContext) {}

// EnterWhenClause is called when production whenClause is entered.
func (s *BaseFlinkSqlParserListener) EnterWhenClause(ctx *WhenClauseContext) {}

// ExitWhenClause is called when production whenClause is exited.
func (s *BaseFlinkSqlParserListener) ExitWhenClause(ctx *WhenClauseContext) {}

// EnterCatalogPath is called when production catalogPath is entered.
func (s *BaseFlinkSqlParserListener) EnterCatalogPath(ctx *CatalogPathContext) {}

// ExitCatalogPath is called when production catalogPath is exited.
func (s *BaseFlinkSqlParserListener) ExitCatalogPath(ctx *CatalogPathContext) {}

// EnterCatalogPathCreate is called when production catalogPathCreate is entered.
func (s *BaseFlinkSqlParserListener) EnterCatalogPathCreate(ctx *CatalogPathCreateContext) {}

// ExitCatalogPathCreate is called when production catalogPathCreate is exited.
func (s *BaseFlinkSqlParserListener) ExitCatalogPathCreate(ctx *CatalogPathCreateContext) {}

// EnterDatabasePath is called when production databasePath is entered.
func (s *BaseFlinkSqlParserListener) EnterDatabasePath(ctx *DatabasePathContext) {}

// ExitDatabasePath is called when production databasePath is exited.
func (s *BaseFlinkSqlParserListener) ExitDatabasePath(ctx *DatabasePathContext) {}

// EnterDatabasePathCreate is called when production databasePathCreate is entered.
func (s *BaseFlinkSqlParserListener) EnterDatabasePathCreate(ctx *DatabasePathCreateContext) {}

// ExitDatabasePathCreate is called when production databasePathCreate is exited.
func (s *BaseFlinkSqlParserListener) ExitDatabasePathCreate(ctx *DatabasePathCreateContext) {}

// EnterTablePathCreate is called when production tablePathCreate is entered.
func (s *BaseFlinkSqlParserListener) EnterTablePathCreate(ctx *TablePathCreateContext) {}

// ExitTablePathCreate is called when production tablePathCreate is exited.
func (s *BaseFlinkSqlParserListener) ExitTablePathCreate(ctx *TablePathCreateContext) {}

// EnterTablePath is called when production tablePath is entered.
func (s *BaseFlinkSqlParserListener) EnterTablePath(ctx *TablePathContext) {}

// ExitTablePath is called when production tablePath is exited.
func (s *BaseFlinkSqlParserListener) ExitTablePath(ctx *TablePathContext) {}

// EnterViewPath is called when production viewPath is entered.
func (s *BaseFlinkSqlParserListener) EnterViewPath(ctx *ViewPathContext) {}

// ExitViewPath is called when production viewPath is exited.
func (s *BaseFlinkSqlParserListener) ExitViewPath(ctx *ViewPathContext) {}

// EnterViewPathCreate is called when production viewPathCreate is entered.
func (s *BaseFlinkSqlParserListener) EnterViewPathCreate(ctx *ViewPathCreateContext) {}

// ExitViewPathCreate is called when production viewPathCreate is exited.
func (s *BaseFlinkSqlParserListener) ExitViewPathCreate(ctx *ViewPathCreateContext) {}

// EnterUid is called when production uid is entered.
func (s *BaseFlinkSqlParserListener) EnterUid(ctx *UidContext) {}

// ExitUid is called when production uid is exited.
func (s *BaseFlinkSqlParserListener) ExitUid(ctx *UidContext) {}

// EnterWithOption is called when production withOption is entered.
func (s *BaseFlinkSqlParserListener) EnterWithOption(ctx *WithOptionContext) {}

// ExitWithOption is called when production withOption is exited.
func (s *BaseFlinkSqlParserListener) ExitWithOption(ctx *WithOptionContext) {}

// EnterIfNotExists is called when production ifNotExists is entered.
func (s *BaseFlinkSqlParserListener) EnterIfNotExists(ctx *IfNotExistsContext) {}

// ExitIfNotExists is called when production ifNotExists is exited.
func (s *BaseFlinkSqlParserListener) ExitIfNotExists(ctx *IfNotExistsContext) {}

// EnterIfExists is called when production ifExists is entered.
func (s *BaseFlinkSqlParserListener) EnterIfExists(ctx *IfExistsContext) {}

// ExitIfExists is called when production ifExists is exited.
func (s *BaseFlinkSqlParserListener) ExitIfExists(ctx *IfExistsContext) {}

// EnterTablePropertyList is called when production tablePropertyList is entered.
func (s *BaseFlinkSqlParserListener) EnterTablePropertyList(ctx *TablePropertyListContext) {}

// ExitTablePropertyList is called when production tablePropertyList is exited.
func (s *BaseFlinkSqlParserListener) ExitTablePropertyList(ctx *TablePropertyListContext) {}

// EnterTableProperty is called when production tableProperty is entered.
func (s *BaseFlinkSqlParserListener) EnterTableProperty(ctx *TablePropertyContext) {}

// ExitTableProperty is called when production tableProperty is exited.
func (s *BaseFlinkSqlParserListener) ExitTableProperty(ctx *TablePropertyContext) {}

// EnterTablePropertyKey is called when production tablePropertyKey is entered.
func (s *BaseFlinkSqlParserListener) EnterTablePropertyKey(ctx *TablePropertyKeyContext) {}

// ExitTablePropertyKey is called when production tablePropertyKey is exited.
func (s *BaseFlinkSqlParserListener) ExitTablePropertyKey(ctx *TablePropertyKeyContext) {}

// EnterTablePropertyValue is called when production tablePropertyValue is entered.
func (s *BaseFlinkSqlParserListener) EnterTablePropertyValue(ctx *TablePropertyValueContext) {}

// ExitTablePropertyValue is called when production tablePropertyValue is exited.
func (s *BaseFlinkSqlParserListener) ExitTablePropertyValue(ctx *TablePropertyValueContext) {}

// EnterLogicalOperator is called when production logicalOperator is entered.
func (s *BaseFlinkSqlParserListener) EnterLogicalOperator(ctx *LogicalOperatorContext) {}

// ExitLogicalOperator is called when production logicalOperator is exited.
func (s *BaseFlinkSqlParserListener) ExitLogicalOperator(ctx *LogicalOperatorContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseFlinkSqlParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseFlinkSqlParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterBitOperator is called when production bitOperator is entered.
func (s *BaseFlinkSqlParserListener) EnterBitOperator(ctx *BitOperatorContext) {}

// ExitBitOperator is called when production bitOperator is exited.
func (s *BaseFlinkSqlParserListener) ExitBitOperator(ctx *BitOperatorContext) {}

// EnterMathOperator is called when production mathOperator is entered.
func (s *BaseFlinkSqlParserListener) EnterMathOperator(ctx *MathOperatorContext) {}

// ExitMathOperator is called when production mathOperator is exited.
func (s *BaseFlinkSqlParserListener) ExitMathOperator(ctx *MathOperatorContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseFlinkSqlParserListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseFlinkSqlParserListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseFlinkSqlParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseFlinkSqlParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterTimePointLiteral is called when production timePointLiteral is entered.
func (s *BaseFlinkSqlParserListener) EnterTimePointLiteral(ctx *TimePointLiteralContext) {}

// ExitTimePointLiteral is called when production timePointLiteral is exited.
func (s *BaseFlinkSqlParserListener) ExitTimePointLiteral(ctx *TimePointLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseFlinkSqlParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseFlinkSqlParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseFlinkSqlParserListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseFlinkSqlParserListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseFlinkSqlParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseFlinkSqlParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterSetQuantifier is called when production setQuantifier is entered.
func (s *BaseFlinkSqlParserListener) EnterSetQuantifier(ctx *SetQuantifierContext) {}

// ExitSetQuantifier is called when production setQuantifier is exited.
func (s *BaseFlinkSqlParserListener) ExitSetQuantifier(ctx *SetQuantifierContext) {}

// EnterTimePointUnit is called when production timePointUnit is entered.
func (s *BaseFlinkSqlParserListener) EnterTimePointUnit(ctx *TimePointUnitContext) {}

// ExitTimePointUnit is called when production timePointUnit is exited.
func (s *BaseFlinkSqlParserListener) ExitTimePointUnit(ctx *TimePointUnitContext) {}

// EnterTimeIntervalUnit is called when production timeIntervalUnit is entered.
func (s *BaseFlinkSqlParserListener) EnterTimeIntervalUnit(ctx *TimeIntervalUnitContext) {}

// ExitTimeIntervalUnit is called when production timeIntervalUnit is exited.
func (s *BaseFlinkSqlParserListener) ExitTimeIntervalUnit(ctx *TimeIntervalUnitContext) {}

// EnterReservedKeywordsUsedAsFuncParam is called when production reservedKeywordsUsedAsFuncParam is entered.
func (s *BaseFlinkSqlParserListener) EnterReservedKeywordsUsedAsFuncParam(ctx *ReservedKeywordsUsedAsFuncParamContext) {
}

// ExitReservedKeywordsUsedAsFuncParam is called when production reservedKeywordsUsedAsFuncParam is exited.
func (s *BaseFlinkSqlParserListener) ExitReservedKeywordsUsedAsFuncParam(ctx *ReservedKeywordsUsedAsFuncParamContext) {
}

// EnterReservedKeywordsNoParamsUsedAsFuncName is called when production reservedKeywordsNoParamsUsedAsFuncName is entered.
func (s *BaseFlinkSqlParserListener) EnterReservedKeywordsNoParamsUsedAsFuncName(ctx *ReservedKeywordsNoParamsUsedAsFuncNameContext) {
}

// ExitReservedKeywordsNoParamsUsedAsFuncName is called when production reservedKeywordsNoParamsUsedAsFuncName is exited.
func (s *BaseFlinkSqlParserListener) ExitReservedKeywordsNoParamsUsedAsFuncName(ctx *ReservedKeywordsNoParamsUsedAsFuncNameContext) {
}

// EnterReservedKeywordsFollowParamsUsedAsFuncName is called when production reservedKeywordsFollowParamsUsedAsFuncName is entered.
func (s *BaseFlinkSqlParserListener) EnterReservedKeywordsFollowParamsUsedAsFuncName(ctx *ReservedKeywordsFollowParamsUsedAsFuncNameContext) {
}

// ExitReservedKeywordsFollowParamsUsedAsFuncName is called when production reservedKeywordsFollowParamsUsedAsFuncName is exited.
func (s *BaseFlinkSqlParserListener) ExitReservedKeywordsFollowParamsUsedAsFuncName(ctx *ReservedKeywordsFollowParamsUsedAsFuncNameContext) {
}

// EnterReservedKeywordsUsedAsFuncName is called when production reservedKeywordsUsedAsFuncName is entered.
func (s *BaseFlinkSqlParserListener) EnterReservedKeywordsUsedAsFuncName(ctx *ReservedKeywordsUsedAsFuncNameContext) {
}

// ExitReservedKeywordsUsedAsFuncName is called when production reservedKeywordsUsedAsFuncName is exited.
func (s *BaseFlinkSqlParserListener) ExitReservedKeywordsUsedAsFuncName(ctx *ReservedKeywordsUsedAsFuncNameContext) {
}

// EnterNonReservedKeywords is called when production nonReservedKeywords is entered.
func (s *BaseFlinkSqlParserListener) EnterNonReservedKeywords(ctx *NonReservedKeywordsContext) {}

// ExitNonReservedKeywords is called when production nonReservedKeywords is exited.
func (s *BaseFlinkSqlParserListener) ExitNonReservedKeywords(ctx *NonReservedKeywordsContext) {}
