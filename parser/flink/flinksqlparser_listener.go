// Code generated from FlinkSqlParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // FlinkSqlParser

import "github.com/antlr4-go/antlr/v4"

// FlinkSqlParserListener is a complete listener for a parse tree produced by FlinkSqlParser.
type FlinkSqlParserListener interface {
	antlr.ParseTreeListener

	// EnterSqlStatements is called when entering the sqlStatements production.
	EnterSqlStatements(c *SqlStatementsContext)

	// EnterSingleStatement is called when entering the singleStatement production.
	EnterSingleStatement(c *SingleStatementContext)

	// EnterSqlStatement is called when entering the sqlStatement production.
	EnterSqlStatement(c *SqlStatementContext)

	// EnterEmptyStatement is called when entering the emptyStatement production.
	EnterEmptyStatement(c *EmptyStatementContext)

	// EnterDdlStatement is called when entering the ddlStatement production.
	EnterDdlStatement(c *DdlStatementContext)

	// EnterDmlStatement is called when entering the dmlStatement production.
	EnterDmlStatement(c *DmlStatementContext)

	// EnterDescribeStatement is called when entering the describeStatement production.
	EnterDescribeStatement(c *DescribeStatementContext)

	// EnterExplainStatement is called when entering the explainStatement production.
	EnterExplainStatement(c *ExplainStatementContext)

	// EnterExplainDetails is called when entering the explainDetails production.
	EnterExplainDetails(c *ExplainDetailsContext)

	// EnterExplainDetail is called when entering the explainDetail production.
	EnterExplainDetail(c *ExplainDetailContext)

	// EnterUseStatement is called when entering the useStatement production.
	EnterUseStatement(c *UseStatementContext)

	// EnterUseModuleStatement is called when entering the useModuleStatement production.
	EnterUseModuleStatement(c *UseModuleStatementContext)

	// EnterShowStatement is called when entering the showStatement production.
	EnterShowStatement(c *ShowStatementContext)

	// EnterLoadStatement is called when entering the loadStatement production.
	EnterLoadStatement(c *LoadStatementContext)

	// EnterUnloadStatement is called when entering the unloadStatement production.
	EnterUnloadStatement(c *UnloadStatementContext)

	// EnterSetStatement is called when entering the setStatement production.
	EnterSetStatement(c *SetStatementContext)

	// EnterResetStatement is called when entering the resetStatement production.
	EnterResetStatement(c *ResetStatementContext)

	// EnterJarStatement is called when entering the jarStatement production.
	EnterJarStatement(c *JarStatementContext)

	// EnterDtAddStatement is called when entering the dtAddStatement production.
	EnterDtAddStatement(c *DtAddStatementContext)

	// EnterDtFilePath is called when entering the dtFilePath production.
	EnterDtFilePath(c *DtFilePathContext)

	// EnterCreateTable is called when entering the createTable production.
	EnterCreateTable(c *CreateTableContext)

	// EnterSimpleCreateTable is called when entering the simpleCreateTable production.
	EnterSimpleCreateTable(c *SimpleCreateTableContext)

	// EnterCreateTableAsSelect is called when entering the createTableAsSelect production.
	EnterCreateTableAsSelect(c *CreateTableAsSelectContext)

	// EnterColumnOptionDefinition is called when entering the columnOptionDefinition production.
	EnterColumnOptionDefinition(c *ColumnOptionDefinitionContext)

	// EnterPhysicalColumnDefinition is called when entering the physicalColumnDefinition production.
	EnterPhysicalColumnDefinition(c *PhysicalColumnDefinitionContext)

	// EnterColumnNameCreate is called when entering the columnNameCreate production.
	EnterColumnNameCreate(c *ColumnNameCreateContext)

	// EnterColumnName is called when entering the columnName production.
	EnterColumnName(c *ColumnNameContext)

	// EnterColumnNamePath is called when entering the columnNamePath production.
	EnterColumnNamePath(c *ColumnNamePathContext)

	// EnterColumnNameList is called when entering the columnNameList production.
	EnterColumnNameList(c *ColumnNameListContext)

	// EnterColumnType is called when entering the columnType production.
	EnterColumnType(c *ColumnTypeContext)

	// EnterLengthOneDimension is called when entering the lengthOneDimension production.
	EnterLengthOneDimension(c *LengthOneDimensionContext)

	// EnterLengthTwoOptionalDimension is called when entering the lengthTwoOptionalDimension production.
	EnterLengthTwoOptionalDimension(c *LengthTwoOptionalDimensionContext)

	// EnterLengthTwoStringDimension is called when entering the lengthTwoStringDimension production.
	EnterLengthTwoStringDimension(c *LengthTwoStringDimensionContext)

	// EnterLengthOneTypeDimension is called when entering the lengthOneTypeDimension production.
	EnterLengthOneTypeDimension(c *LengthOneTypeDimensionContext)

	// EnterMapTypeDimension is called when entering the mapTypeDimension production.
	EnterMapTypeDimension(c *MapTypeDimensionContext)

	// EnterRowTypeDimension is called when entering the rowTypeDimension production.
	EnterRowTypeDimension(c *RowTypeDimensionContext)

	// EnterColumnConstraint is called when entering the columnConstraint production.
	EnterColumnConstraint(c *ColumnConstraintContext)

	// EnterMetadataColumnDefinition is called when entering the metadataColumnDefinition production.
	EnterMetadataColumnDefinition(c *MetadataColumnDefinitionContext)

	// EnterMetadataKey is called when entering the metadataKey production.
	EnterMetadataKey(c *MetadataKeyContext)

	// EnterComputedColumnDefinition is called when entering the computedColumnDefinition production.
	EnterComputedColumnDefinition(c *ComputedColumnDefinitionContext)

	// EnterComputedColumnExpression is called when entering the computedColumnExpression production.
	EnterComputedColumnExpression(c *ComputedColumnExpressionContext)

	// EnterWatermarkDefinition is called when entering the watermarkDefinition production.
	EnterWatermarkDefinition(c *WatermarkDefinitionContext)

	// EnterTableConstraint is called when entering the tableConstraint production.
	EnterTableConstraint(c *TableConstraintContext)

	// EnterConstraintName is called when entering the constraintName production.
	EnterConstraintName(c *ConstraintNameContext)

	// EnterSelfDefinitionClause is called when entering the selfDefinitionClause production.
	EnterSelfDefinitionClause(c *SelfDefinitionClauseContext)

	// EnterPartitionDefinition is called when entering the partitionDefinition production.
	EnterPartitionDefinition(c *PartitionDefinitionContext)

	// EnterTransformList is called when entering the transformList production.
	EnterTransformList(c *TransformListContext)

	// EnterIdentityTransform is called when entering the identityTransform production.
	EnterIdentityTransform(c *IdentityTransformContext)

	// EnterApplyTransform is called when entering the applyTransform production.
	EnterApplyTransform(c *ApplyTransformContext)

	// EnterTransformArgument is called when entering the transformArgument production.
	EnterTransformArgument(c *TransformArgumentContext)

	// EnterLikeDefinition is called when entering the likeDefinition production.
	EnterLikeDefinition(c *LikeDefinitionContext)

	// EnterLikeOption is called when entering the likeOption production.
	EnterLikeOption(c *LikeOptionContext)

	// EnterCreateCatalog is called when entering the createCatalog production.
	EnterCreateCatalog(c *CreateCatalogContext)

	// EnterCreateDatabase is called when entering the createDatabase production.
	EnterCreateDatabase(c *CreateDatabaseContext)

	// EnterCreateView is called when entering the createView production.
	EnterCreateView(c *CreateViewContext)

	// EnterCreateFunction is called when entering the createFunction production.
	EnterCreateFunction(c *CreateFunctionContext)

	// EnterUsingClause is called when entering the usingClause production.
	EnterUsingClause(c *UsingClauseContext)

	// EnterJarFileName is called when entering the jarFileName production.
	EnterJarFileName(c *JarFileNameContext)

	// EnterAlterTable is called when entering the alterTable production.
	EnterAlterTable(c *AlterTableContext)

	// EnterRenameDefinition is called when entering the renameDefinition production.
	EnterRenameDefinition(c *RenameDefinitionContext)

	// EnterSetKeyValueDefinition is called when entering the setKeyValueDefinition production.
	EnterSetKeyValueDefinition(c *SetKeyValueDefinitionContext)

	// EnterAddConstraint is called when entering the addConstraint production.
	EnterAddConstraint(c *AddConstraintContext)

	// EnterDropConstraint is called when entering the dropConstraint production.
	EnterDropConstraint(c *DropConstraintContext)

	// EnterAddUnique is called when entering the addUnique production.
	EnterAddUnique(c *AddUniqueContext)

	// EnterNotForced is called when entering the notForced production.
	EnterNotForced(c *NotForcedContext)

	// EnterAlterView is called when entering the alterView production.
	EnterAlterView(c *AlterViewContext)

	// EnterAlterDatabase is called when entering the alterDatabase production.
	EnterAlterDatabase(c *AlterDatabaseContext)

	// EnterAlterFunction is called when entering the alterFunction production.
	EnterAlterFunction(c *AlterFunctionContext)

	// EnterDropCatalog is called when entering the dropCatalog production.
	EnterDropCatalog(c *DropCatalogContext)

	// EnterDropTable is called when entering the dropTable production.
	EnterDropTable(c *DropTableContext)

	// EnterDropDatabase is called when entering the dropDatabase production.
	EnterDropDatabase(c *DropDatabaseContext)

	// EnterDropView is called when entering the dropView production.
	EnterDropView(c *DropViewContext)

	// EnterDropFunction is called when entering the dropFunction production.
	EnterDropFunction(c *DropFunctionContext)

	// EnterInsertStatement is called when entering the insertStatement production.
	EnterInsertStatement(c *InsertStatementContext)

	// EnterInsertSimpleStatement is called when entering the insertSimpleStatement production.
	EnterInsertSimpleStatement(c *InsertSimpleStatementContext)

	// EnterInsertPartitionDefinition is called when entering the insertPartitionDefinition production.
	EnterInsertPartitionDefinition(c *InsertPartitionDefinitionContext)

	// EnterValuesDefinition is called when entering the valuesDefinition production.
	EnterValuesDefinition(c *ValuesDefinitionContext)

	// EnterValuesRowDefinition is called when entering the valuesRowDefinition production.
	EnterValuesRowDefinition(c *ValuesRowDefinitionContext)

	// EnterInsertMulStatementCompatibility is called when entering the insertMulStatementCompatibility production.
	EnterInsertMulStatementCompatibility(c *InsertMulStatementCompatibilityContext)

	// EnterInsertMulStatement is called when entering the insertMulStatement production.
	EnterInsertMulStatement(c *InsertMulStatementContext)

	// EnterQueryStatement is called when entering the queryStatement production.
	EnterQueryStatement(c *QueryStatementContext)

	// EnterValuesClause is called when entering the valuesClause production.
	EnterValuesClause(c *ValuesClauseContext)

	// EnterWithClause is called when entering the withClause production.
	EnterWithClause(c *WithClauseContext)

	// EnterWithItem is called when entering the withItem production.
	EnterWithItem(c *WithItemContext)

	// EnterWithItemName is called when entering the withItemName production.
	EnterWithItemName(c *WithItemNameContext)

	// EnterSelectStatement is called when entering the selectStatement production.
	EnterSelectStatement(c *SelectStatementContext)

	// EnterSelectClause is called when entering the selectClause production.
	EnterSelectClause(c *SelectClauseContext)

	// EnterProjectItemDefinition is called when entering the projectItemDefinition production.
	EnterProjectItemDefinition(c *ProjectItemDefinitionContext)

	// EnterOverWindowItem is called when entering the overWindowItem production.
	EnterOverWindowItem(c *OverWindowItemContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterTableExpression is called when entering the tableExpression production.
	EnterTableExpression(c *TableExpressionContext)

	// EnterTableReference is called when entering the tableReference production.
	EnterTableReference(c *TableReferenceContext)

	// EnterTablePrimary is called when entering the tablePrimary production.
	EnterTablePrimary(c *TablePrimaryContext)

	// EnterSystemTimePeriod is called when entering the systemTimePeriod production.
	EnterSystemTimePeriod(c *SystemTimePeriodContext)

	// EnterDateTimeExpression is called when entering the dateTimeExpression production.
	EnterDateTimeExpression(c *DateTimeExpressionContext)

	// EnterInlineDataValueClause is called when entering the inlineDataValueClause production.
	EnterInlineDataValueClause(c *InlineDataValueClauseContext)

	// EnterWindowTVFClause is called when entering the windowTVFClause production.
	EnterWindowTVFClause(c *WindowTVFClauseContext)

	// EnterWindowTVFExpression is called when entering the windowTVFExpression production.
	EnterWindowTVFExpression(c *WindowTVFExpressionContext)

	// EnterWindowTVFName is called when entering the windowTVFName production.
	EnterWindowTVFName(c *WindowTVFNameContext)

	// EnterWindowTVFParam is called when entering the windowTVFParam production.
	EnterWindowTVFParam(c *WindowTVFParamContext)

	// EnterTimeIntervalParamName is called when entering the timeIntervalParamName production.
	EnterTimeIntervalParamName(c *TimeIntervalParamNameContext)

	// EnterColumnDescriptor is called when entering the columnDescriptor production.
	EnterColumnDescriptor(c *ColumnDescriptorContext)

	// EnterJoinCondition is called when entering the joinCondition production.
	EnterJoinCondition(c *JoinConditionContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterGroupByClause is called when entering the groupByClause production.
	EnterGroupByClause(c *GroupByClauseContext)

	// EnterGroupItemDefinition is called when entering the groupItemDefinition production.
	EnterGroupItemDefinition(c *GroupItemDefinitionContext)

	// EnterGroupingSets is called when entering the groupingSets production.
	EnterGroupingSets(c *GroupingSetsContext)

	// EnterGroupingSetsNotationName is called when entering the groupingSetsNotationName production.
	EnterGroupingSetsNotationName(c *GroupingSetsNotationNameContext)

	// EnterGroupWindowFunction is called when entering the groupWindowFunction production.
	EnterGroupWindowFunction(c *GroupWindowFunctionContext)

	// EnterGroupWindowFunctionName is called when entering the groupWindowFunctionName production.
	EnterGroupWindowFunctionName(c *GroupWindowFunctionNameContext)

	// EnterTimeAttrColumn is called when entering the timeAttrColumn production.
	EnterTimeAttrColumn(c *TimeAttrColumnContext)

	// EnterHavingClause is called when entering the havingClause production.
	EnterHavingClause(c *HavingClauseContext)

	// EnterWindowClause is called when entering the windowClause production.
	EnterWindowClause(c *WindowClauseContext)

	// EnterNamedWindow is called when entering the namedWindow production.
	EnterNamedWindow(c *NamedWindowContext)

	// EnterWindowSpec is called when entering the windowSpec production.
	EnterWindowSpec(c *WindowSpecContext)

	// EnterMatchRecognizeClause is called when entering the matchRecognizeClause production.
	EnterMatchRecognizeClause(c *MatchRecognizeClauseContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterOrderItemDefinition is called when entering the orderItemDefinition production.
	EnterOrderItemDefinition(c *OrderItemDefinitionContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterPartitionByClause is called when entering the partitionByClause production.
	EnterPartitionByClause(c *PartitionByClauseContext)

	// EnterQuantifiers is called when entering the quantifiers production.
	EnterQuantifiers(c *QuantifiersContext)

	// EnterMeasuresClause is called when entering the measuresClause production.
	EnterMeasuresClause(c *MeasuresClauseContext)

	// EnterPatternDefinition is called when entering the patternDefinition production.
	EnterPatternDefinition(c *PatternDefinitionContext)

	// EnterPatternVariable is called when entering the patternVariable production.
	EnterPatternVariable(c *PatternVariableContext)

	// EnterOutputMode is called when entering the outputMode production.
	EnterOutputMode(c *OutputModeContext)

	// EnterAfterMatchStrategy is called when entering the afterMatchStrategy production.
	EnterAfterMatchStrategy(c *AfterMatchStrategyContext)

	// EnterPatternVariablesDefinition is called when entering the patternVariablesDefinition production.
	EnterPatternVariablesDefinition(c *PatternVariablesDefinitionContext)

	// EnterWindowFrame is called when entering the windowFrame production.
	EnterWindowFrame(c *WindowFrameContext)

	// EnterFrameBound is called when entering the frameBound production.
	EnterFrameBound(c *FrameBoundContext)

	// EnterWithinClause is called when entering the withinClause production.
	EnterWithinClause(c *WithinClauseContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLogicalNot is called when entering the logicalNot production.
	EnterLogicalNot(c *LogicalNotContext)

	// EnterPredicated is called when entering the predicated production.
	EnterPredicated(c *PredicatedContext)

	// EnterExists is called when entering the exists production.
	EnterExists(c *ExistsContext)

	// EnterLogicalNested is called when entering the logicalNested production.
	EnterLogicalNested(c *LogicalNestedContext)

	// EnterLogicalBinary is called when entering the logicalBinary production.
	EnterLogicalBinary(c *LogicalBinaryContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterLikePredicate is called when entering the likePredicate production.
	EnterLikePredicate(c *LikePredicateContext)

	// EnterValueExpressionDefault is called when entering the valueExpressionDefault production.
	EnterValueExpressionDefault(c *ValueExpressionDefaultContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterArithmeticBinary is called when entering the arithmeticBinary production.
	EnterArithmeticBinary(c *ArithmeticBinaryContext)

	// EnterArithmeticUnary is called when entering the arithmeticUnary production.
	EnterArithmeticUnary(c *ArithmeticUnaryContext)

	// EnterFunctionCallExpression is called when entering the functionCallExpression production.
	EnterFunctionCallExpression(c *FunctionCallExpressionContext)

	// EnterDereference is called when entering the dereference production.
	EnterDereference(c *DereferenceContext)

	// EnterSimpleCase is called when entering the simpleCase production.
	EnterSimpleCase(c *SimpleCaseContext)

	// EnterColumnReference is called when entering the columnReference production.
	EnterColumnReference(c *ColumnReferenceContext)

	// EnterLast is called when entering the last production.
	EnterLast(c *LastContext)

	// EnterStar is called when entering the star production.
	EnterStar(c *StarContext)

	// EnterSubscript is called when entering the subscript production.
	EnterSubscript(c *SubscriptContext)

	// EnterSubqueryExpression is called when entering the subqueryExpression production.
	EnterSubqueryExpression(c *SubqueryExpressionContext)

	// EnterCast is called when entering the cast production.
	EnterCast(c *CastContext)

	// EnterConstantDefault is called when entering the constantDefault production.
	EnterConstantDefault(c *ConstantDefaultContext)

	// EnterParenthesizedExpression is called when entering the parenthesizedExpression production.
	EnterParenthesizedExpression(c *ParenthesizedExpressionContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterSearchedCase is called when entering the searchedCase production.
	EnterSearchedCase(c *SearchedCaseContext)

	// EnterPosition is called when entering the position production.
	EnterPosition(c *PositionContext)

	// EnterFirst is called when entering the first production.
	EnterFirst(c *FirstContext)

	// EnterFunctionNameCreate is called when entering the functionNameCreate production.
	EnterFunctionNameCreate(c *FunctionNameCreateContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterFunctionNameAndParams is called when entering the functionNameAndParams production.
	EnterFunctionNameAndParams(c *FunctionNameAndParamsContext)

	// EnterFunctionNameWithParams is called when entering the functionNameWithParams production.
	EnterFunctionNameWithParams(c *FunctionNameWithParamsContext)

	// EnterFunctionParam is called when entering the functionParam production.
	EnterFunctionParam(c *FunctionParamContext)

	// EnterDereferenceDefinition is called when entering the dereferenceDefinition production.
	EnterDereferenceDefinition(c *DereferenceDefinitionContext)

	// EnterCorrelationName is called when entering the correlationName production.
	EnterCorrelationName(c *CorrelationNameContext)

	// EnterQualifiedName is called when entering the qualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// EnterTimeIntervalExpression is called when entering the timeIntervalExpression production.
	EnterTimeIntervalExpression(c *TimeIntervalExpressionContext)

	// EnterErrorCapturingMultiUnitsInterval is called when entering the errorCapturingMultiUnitsInterval production.
	EnterErrorCapturingMultiUnitsInterval(c *ErrorCapturingMultiUnitsIntervalContext)

	// EnterMultiUnitsInterval is called when entering the multiUnitsInterval production.
	EnterMultiUnitsInterval(c *MultiUnitsIntervalContext)

	// EnterErrorCapturingUnitToUnitInterval is called when entering the errorCapturingUnitToUnitInterval production.
	EnterErrorCapturingUnitToUnitInterval(c *ErrorCapturingUnitToUnitIntervalContext)

	// EnterUnitToUnitInterval is called when entering the unitToUnitInterval production.
	EnterUnitToUnitInterval(c *UnitToUnitIntervalContext)

	// EnterIntervalValue is called when entering the intervalValue production.
	EnterIntervalValue(c *IntervalValueContext)

	// EnterTableAlias is called when entering the tableAlias production.
	EnterTableAlias(c *TableAliasContext)

	// EnterErrorCapturingIdentifier is called when entering the errorCapturingIdentifier production.
	EnterErrorCapturingIdentifier(c *ErrorCapturingIdentifierContext)

	// EnterErrorIdent is called when entering the errorIdent production.
	EnterErrorIdent(c *ErrorIdentContext)

	// EnterRealIdent is called when entering the realIdent production.
	EnterRealIdent(c *RealIdentContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterIdentifierSeq is called when entering the identifierSeq production.
	EnterIdentifierSeq(c *IdentifierSeqContext)

	// EnterUnquotedIdentifierAlternative is called when entering the unquotedIdentifierAlternative production.
	EnterUnquotedIdentifierAlternative(c *UnquotedIdentifierAlternativeContext)

	// EnterQuotedIdentifierAlternative is called when entering the quotedIdentifierAlternative production.
	EnterQuotedIdentifierAlternative(c *QuotedIdentifierAlternativeContext)

	// EnterNonReservedKeywordsAlternative is called when entering the nonReservedKeywordsAlternative production.
	EnterNonReservedKeywordsAlternative(c *NonReservedKeywordsAlternativeContext)

	// EnterUnquotedIdentifier is called when entering the unquotedIdentifier production.
	EnterUnquotedIdentifier(c *UnquotedIdentifierContext)

	// EnterQuotedIdentifier is called when entering the quotedIdentifier production.
	EnterQuotedIdentifier(c *QuotedIdentifierContext)

	// EnterWhenClause is called when entering the whenClause production.
	EnterWhenClause(c *WhenClauseContext)

	// EnterCatalogPath is called when entering the catalogPath production.
	EnterCatalogPath(c *CatalogPathContext)

	// EnterCatalogPathCreate is called when entering the catalogPathCreate production.
	EnterCatalogPathCreate(c *CatalogPathCreateContext)

	// EnterDatabasePath is called when entering the databasePath production.
	EnterDatabasePath(c *DatabasePathContext)

	// EnterDatabasePathCreate is called when entering the databasePathCreate production.
	EnterDatabasePathCreate(c *DatabasePathCreateContext)

	// EnterTablePathCreate is called when entering the tablePathCreate production.
	EnterTablePathCreate(c *TablePathCreateContext)

	// EnterTablePath is called when entering the tablePath production.
	EnterTablePath(c *TablePathContext)

	// EnterViewPath is called when entering the viewPath production.
	EnterViewPath(c *ViewPathContext)

	// EnterViewPathCreate is called when entering the viewPathCreate production.
	EnterViewPathCreate(c *ViewPathCreateContext)

	// EnterUid is called when entering the uid production.
	EnterUid(c *UidContext)

	// EnterWithOption is called when entering the withOption production.
	EnterWithOption(c *WithOptionContext)

	// EnterIfNotExists is called when entering the ifNotExists production.
	EnterIfNotExists(c *IfNotExistsContext)

	// EnterIfExists is called when entering the ifExists production.
	EnterIfExists(c *IfExistsContext)

	// EnterTablePropertyList is called when entering the tablePropertyList production.
	EnterTablePropertyList(c *TablePropertyListContext)

	// EnterTableProperty is called when entering the tableProperty production.
	EnterTableProperty(c *TablePropertyContext)

	// EnterTablePropertyKey is called when entering the tablePropertyKey production.
	EnterTablePropertyKey(c *TablePropertyKeyContext)

	// EnterTablePropertyValue is called when entering the tablePropertyValue production.
	EnterTablePropertyValue(c *TablePropertyValueContext)

	// EnterLogicalOperator is called when entering the logicalOperator production.
	EnterLogicalOperator(c *LogicalOperatorContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterBitOperator is called when entering the bitOperator production.
	EnterBitOperator(c *BitOperatorContext)

	// EnterMathOperator is called when entering the mathOperator production.
	EnterMathOperator(c *MathOperatorContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterTimePointLiteral is called when entering the timePointLiteral production.
	EnterTimePointLiteral(c *TimePointLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterDecimalLiteral is called when entering the decimalLiteral production.
	EnterDecimalLiteral(c *DecimalLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterSetQuantifier is called when entering the setQuantifier production.
	EnterSetQuantifier(c *SetQuantifierContext)

	// EnterTimePointUnit is called when entering the timePointUnit production.
	EnterTimePointUnit(c *TimePointUnitContext)

	// EnterTimeIntervalUnit is called when entering the timeIntervalUnit production.
	EnterTimeIntervalUnit(c *TimeIntervalUnitContext)

	// EnterReservedKeywordsUsedAsFuncParam is called when entering the reservedKeywordsUsedAsFuncParam production.
	EnterReservedKeywordsUsedAsFuncParam(c *ReservedKeywordsUsedAsFuncParamContext)

	// EnterReservedKeywordsNoParamsUsedAsFuncName is called when entering the reservedKeywordsNoParamsUsedAsFuncName production.
	EnterReservedKeywordsNoParamsUsedAsFuncName(c *ReservedKeywordsNoParamsUsedAsFuncNameContext)

	// EnterReservedKeywordsFollowParamsUsedAsFuncName is called when entering the reservedKeywordsFollowParamsUsedAsFuncName production.
	EnterReservedKeywordsFollowParamsUsedAsFuncName(c *ReservedKeywordsFollowParamsUsedAsFuncNameContext)

	// EnterReservedKeywordsUsedAsFuncName is called when entering the reservedKeywordsUsedAsFuncName production.
	EnterReservedKeywordsUsedAsFuncName(c *ReservedKeywordsUsedAsFuncNameContext)

	// EnterNonReservedKeywords is called when entering the nonReservedKeywords production.
	EnterNonReservedKeywords(c *NonReservedKeywordsContext)

	// ExitSqlStatements is called when exiting the sqlStatements production.
	ExitSqlStatements(c *SqlStatementsContext)

	// ExitSingleStatement is called when exiting the singleStatement production.
	ExitSingleStatement(c *SingleStatementContext)

	// ExitSqlStatement is called when exiting the sqlStatement production.
	ExitSqlStatement(c *SqlStatementContext)

	// ExitEmptyStatement is called when exiting the emptyStatement production.
	ExitEmptyStatement(c *EmptyStatementContext)

	// ExitDdlStatement is called when exiting the ddlStatement production.
	ExitDdlStatement(c *DdlStatementContext)

	// ExitDmlStatement is called when exiting the dmlStatement production.
	ExitDmlStatement(c *DmlStatementContext)

	// ExitDescribeStatement is called when exiting the describeStatement production.
	ExitDescribeStatement(c *DescribeStatementContext)

	// ExitExplainStatement is called when exiting the explainStatement production.
	ExitExplainStatement(c *ExplainStatementContext)

	// ExitExplainDetails is called when exiting the explainDetails production.
	ExitExplainDetails(c *ExplainDetailsContext)

	// ExitExplainDetail is called when exiting the explainDetail production.
	ExitExplainDetail(c *ExplainDetailContext)

	// ExitUseStatement is called when exiting the useStatement production.
	ExitUseStatement(c *UseStatementContext)

	// ExitUseModuleStatement is called when exiting the useModuleStatement production.
	ExitUseModuleStatement(c *UseModuleStatementContext)

	// ExitShowStatement is called when exiting the showStatement production.
	ExitShowStatement(c *ShowStatementContext)

	// ExitLoadStatement is called when exiting the loadStatement production.
	ExitLoadStatement(c *LoadStatementContext)

	// ExitUnloadStatement is called when exiting the unloadStatement production.
	ExitUnloadStatement(c *UnloadStatementContext)

	// ExitSetStatement is called when exiting the setStatement production.
	ExitSetStatement(c *SetStatementContext)

	// ExitResetStatement is called when exiting the resetStatement production.
	ExitResetStatement(c *ResetStatementContext)

	// ExitJarStatement is called when exiting the jarStatement production.
	ExitJarStatement(c *JarStatementContext)

	// ExitDtAddStatement is called when exiting the dtAddStatement production.
	ExitDtAddStatement(c *DtAddStatementContext)

	// ExitDtFilePath is called when exiting the dtFilePath production.
	ExitDtFilePath(c *DtFilePathContext)

	// ExitCreateTable is called when exiting the createTable production.
	ExitCreateTable(c *CreateTableContext)

	// ExitSimpleCreateTable is called when exiting the simpleCreateTable production.
	ExitSimpleCreateTable(c *SimpleCreateTableContext)

	// ExitCreateTableAsSelect is called when exiting the createTableAsSelect production.
	ExitCreateTableAsSelect(c *CreateTableAsSelectContext)

	// ExitColumnOptionDefinition is called when exiting the columnOptionDefinition production.
	ExitColumnOptionDefinition(c *ColumnOptionDefinitionContext)

	// ExitPhysicalColumnDefinition is called when exiting the physicalColumnDefinition production.
	ExitPhysicalColumnDefinition(c *PhysicalColumnDefinitionContext)

	// ExitColumnNameCreate is called when exiting the columnNameCreate production.
	ExitColumnNameCreate(c *ColumnNameCreateContext)

	// ExitColumnName is called when exiting the columnName production.
	ExitColumnName(c *ColumnNameContext)

	// ExitColumnNamePath is called when exiting the columnNamePath production.
	ExitColumnNamePath(c *ColumnNamePathContext)

	// ExitColumnNameList is called when exiting the columnNameList production.
	ExitColumnNameList(c *ColumnNameListContext)

	// ExitColumnType is called when exiting the columnType production.
	ExitColumnType(c *ColumnTypeContext)

	// ExitLengthOneDimension is called when exiting the lengthOneDimension production.
	ExitLengthOneDimension(c *LengthOneDimensionContext)

	// ExitLengthTwoOptionalDimension is called when exiting the lengthTwoOptionalDimension production.
	ExitLengthTwoOptionalDimension(c *LengthTwoOptionalDimensionContext)

	// ExitLengthTwoStringDimension is called when exiting the lengthTwoStringDimension production.
	ExitLengthTwoStringDimension(c *LengthTwoStringDimensionContext)

	// ExitLengthOneTypeDimension is called when exiting the lengthOneTypeDimension production.
	ExitLengthOneTypeDimension(c *LengthOneTypeDimensionContext)

	// ExitMapTypeDimension is called when exiting the mapTypeDimension production.
	ExitMapTypeDimension(c *MapTypeDimensionContext)

	// ExitRowTypeDimension is called when exiting the rowTypeDimension production.
	ExitRowTypeDimension(c *RowTypeDimensionContext)

	// ExitColumnConstraint is called when exiting the columnConstraint production.
	ExitColumnConstraint(c *ColumnConstraintContext)

	// ExitMetadataColumnDefinition is called when exiting the metadataColumnDefinition production.
	ExitMetadataColumnDefinition(c *MetadataColumnDefinitionContext)

	// ExitMetadataKey is called when exiting the metadataKey production.
	ExitMetadataKey(c *MetadataKeyContext)

	// ExitComputedColumnDefinition is called when exiting the computedColumnDefinition production.
	ExitComputedColumnDefinition(c *ComputedColumnDefinitionContext)

	// ExitComputedColumnExpression is called when exiting the computedColumnExpression production.
	ExitComputedColumnExpression(c *ComputedColumnExpressionContext)

	// ExitWatermarkDefinition is called when exiting the watermarkDefinition production.
	ExitWatermarkDefinition(c *WatermarkDefinitionContext)

	// ExitTableConstraint is called when exiting the tableConstraint production.
	ExitTableConstraint(c *TableConstraintContext)

	// ExitConstraintName is called when exiting the constraintName production.
	ExitConstraintName(c *ConstraintNameContext)

	// ExitSelfDefinitionClause is called when exiting the selfDefinitionClause production.
	ExitSelfDefinitionClause(c *SelfDefinitionClauseContext)

	// ExitPartitionDefinition is called when exiting the partitionDefinition production.
	ExitPartitionDefinition(c *PartitionDefinitionContext)

	// ExitTransformList is called when exiting the transformList production.
	ExitTransformList(c *TransformListContext)

	// ExitIdentityTransform is called when exiting the identityTransform production.
	ExitIdentityTransform(c *IdentityTransformContext)

	// ExitApplyTransform is called when exiting the applyTransform production.
	ExitApplyTransform(c *ApplyTransformContext)

	// ExitTransformArgument is called when exiting the transformArgument production.
	ExitTransformArgument(c *TransformArgumentContext)

	// ExitLikeDefinition is called when exiting the likeDefinition production.
	ExitLikeDefinition(c *LikeDefinitionContext)

	// ExitLikeOption is called when exiting the likeOption production.
	ExitLikeOption(c *LikeOptionContext)

	// ExitCreateCatalog is called when exiting the createCatalog production.
	ExitCreateCatalog(c *CreateCatalogContext)

	// ExitCreateDatabase is called when exiting the createDatabase production.
	ExitCreateDatabase(c *CreateDatabaseContext)

	// ExitCreateView is called when exiting the createView production.
	ExitCreateView(c *CreateViewContext)

	// ExitCreateFunction is called when exiting the createFunction production.
	ExitCreateFunction(c *CreateFunctionContext)

	// ExitUsingClause is called when exiting the usingClause production.
	ExitUsingClause(c *UsingClauseContext)

	// ExitJarFileName is called when exiting the jarFileName production.
	ExitJarFileName(c *JarFileNameContext)

	// ExitAlterTable is called when exiting the alterTable production.
	ExitAlterTable(c *AlterTableContext)

	// ExitRenameDefinition is called when exiting the renameDefinition production.
	ExitRenameDefinition(c *RenameDefinitionContext)

	// ExitSetKeyValueDefinition is called when exiting the setKeyValueDefinition production.
	ExitSetKeyValueDefinition(c *SetKeyValueDefinitionContext)

	// ExitAddConstraint is called when exiting the addConstraint production.
	ExitAddConstraint(c *AddConstraintContext)

	// ExitDropConstraint is called when exiting the dropConstraint production.
	ExitDropConstraint(c *DropConstraintContext)

	// ExitAddUnique is called when exiting the addUnique production.
	ExitAddUnique(c *AddUniqueContext)

	// ExitNotForced is called when exiting the notForced production.
	ExitNotForced(c *NotForcedContext)

	// ExitAlterView is called when exiting the alterView production.
	ExitAlterView(c *AlterViewContext)

	// ExitAlterDatabase is called when exiting the alterDatabase production.
	ExitAlterDatabase(c *AlterDatabaseContext)

	// ExitAlterFunction is called when exiting the alterFunction production.
	ExitAlterFunction(c *AlterFunctionContext)

	// ExitDropCatalog is called when exiting the dropCatalog production.
	ExitDropCatalog(c *DropCatalogContext)

	// ExitDropTable is called when exiting the dropTable production.
	ExitDropTable(c *DropTableContext)

	// ExitDropDatabase is called when exiting the dropDatabase production.
	ExitDropDatabase(c *DropDatabaseContext)

	// ExitDropView is called when exiting the dropView production.
	ExitDropView(c *DropViewContext)

	// ExitDropFunction is called when exiting the dropFunction production.
	ExitDropFunction(c *DropFunctionContext)

	// ExitInsertStatement is called when exiting the insertStatement production.
	ExitInsertStatement(c *InsertStatementContext)

	// ExitInsertSimpleStatement is called when exiting the insertSimpleStatement production.
	ExitInsertSimpleStatement(c *InsertSimpleStatementContext)

	// ExitInsertPartitionDefinition is called when exiting the insertPartitionDefinition production.
	ExitInsertPartitionDefinition(c *InsertPartitionDefinitionContext)

	// ExitValuesDefinition is called when exiting the valuesDefinition production.
	ExitValuesDefinition(c *ValuesDefinitionContext)

	// ExitValuesRowDefinition is called when exiting the valuesRowDefinition production.
	ExitValuesRowDefinition(c *ValuesRowDefinitionContext)

	// ExitInsertMulStatementCompatibility is called when exiting the insertMulStatementCompatibility production.
	ExitInsertMulStatementCompatibility(c *InsertMulStatementCompatibilityContext)

	// ExitInsertMulStatement is called when exiting the insertMulStatement production.
	ExitInsertMulStatement(c *InsertMulStatementContext)

	// ExitQueryStatement is called when exiting the queryStatement production.
	ExitQueryStatement(c *QueryStatementContext)

	// ExitValuesClause is called when exiting the valuesClause production.
	ExitValuesClause(c *ValuesClauseContext)

	// ExitWithClause is called when exiting the withClause production.
	ExitWithClause(c *WithClauseContext)

	// ExitWithItem is called when exiting the withItem production.
	ExitWithItem(c *WithItemContext)

	// ExitWithItemName is called when exiting the withItemName production.
	ExitWithItemName(c *WithItemNameContext)

	// ExitSelectStatement is called when exiting the selectStatement production.
	ExitSelectStatement(c *SelectStatementContext)

	// ExitSelectClause is called when exiting the selectClause production.
	ExitSelectClause(c *SelectClauseContext)

	// ExitProjectItemDefinition is called when exiting the projectItemDefinition production.
	ExitProjectItemDefinition(c *ProjectItemDefinitionContext)

	// ExitOverWindowItem is called when exiting the overWindowItem production.
	ExitOverWindowItem(c *OverWindowItemContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitTableExpression is called when exiting the tableExpression production.
	ExitTableExpression(c *TableExpressionContext)

	// ExitTableReference is called when exiting the tableReference production.
	ExitTableReference(c *TableReferenceContext)

	// ExitTablePrimary is called when exiting the tablePrimary production.
	ExitTablePrimary(c *TablePrimaryContext)

	// ExitSystemTimePeriod is called when exiting the systemTimePeriod production.
	ExitSystemTimePeriod(c *SystemTimePeriodContext)

	// ExitDateTimeExpression is called when exiting the dateTimeExpression production.
	ExitDateTimeExpression(c *DateTimeExpressionContext)

	// ExitInlineDataValueClause is called when exiting the inlineDataValueClause production.
	ExitInlineDataValueClause(c *InlineDataValueClauseContext)

	// ExitWindowTVFClause is called when exiting the windowTVFClause production.
	ExitWindowTVFClause(c *WindowTVFClauseContext)

	// ExitWindowTVFExpression is called when exiting the windowTVFExpression production.
	ExitWindowTVFExpression(c *WindowTVFExpressionContext)

	// ExitWindowTVFName is called when exiting the windowTVFName production.
	ExitWindowTVFName(c *WindowTVFNameContext)

	// ExitWindowTVFParam is called when exiting the windowTVFParam production.
	ExitWindowTVFParam(c *WindowTVFParamContext)

	// ExitTimeIntervalParamName is called when exiting the timeIntervalParamName production.
	ExitTimeIntervalParamName(c *TimeIntervalParamNameContext)

	// ExitColumnDescriptor is called when exiting the columnDescriptor production.
	ExitColumnDescriptor(c *ColumnDescriptorContext)

	// ExitJoinCondition is called when exiting the joinCondition production.
	ExitJoinCondition(c *JoinConditionContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitGroupByClause is called when exiting the groupByClause production.
	ExitGroupByClause(c *GroupByClauseContext)

	// ExitGroupItemDefinition is called when exiting the groupItemDefinition production.
	ExitGroupItemDefinition(c *GroupItemDefinitionContext)

	// ExitGroupingSets is called when exiting the groupingSets production.
	ExitGroupingSets(c *GroupingSetsContext)

	// ExitGroupingSetsNotationName is called when exiting the groupingSetsNotationName production.
	ExitGroupingSetsNotationName(c *GroupingSetsNotationNameContext)

	// ExitGroupWindowFunction is called when exiting the groupWindowFunction production.
	ExitGroupWindowFunction(c *GroupWindowFunctionContext)

	// ExitGroupWindowFunctionName is called when exiting the groupWindowFunctionName production.
	ExitGroupWindowFunctionName(c *GroupWindowFunctionNameContext)

	// ExitTimeAttrColumn is called when exiting the timeAttrColumn production.
	ExitTimeAttrColumn(c *TimeAttrColumnContext)

	// ExitHavingClause is called when exiting the havingClause production.
	ExitHavingClause(c *HavingClauseContext)

	// ExitWindowClause is called when exiting the windowClause production.
	ExitWindowClause(c *WindowClauseContext)

	// ExitNamedWindow is called when exiting the namedWindow production.
	ExitNamedWindow(c *NamedWindowContext)

	// ExitWindowSpec is called when exiting the windowSpec production.
	ExitWindowSpec(c *WindowSpecContext)

	// ExitMatchRecognizeClause is called when exiting the matchRecognizeClause production.
	ExitMatchRecognizeClause(c *MatchRecognizeClauseContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitOrderItemDefinition is called when exiting the orderItemDefinition production.
	ExitOrderItemDefinition(c *OrderItemDefinitionContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitPartitionByClause is called when exiting the partitionByClause production.
	ExitPartitionByClause(c *PartitionByClauseContext)

	// ExitQuantifiers is called when exiting the quantifiers production.
	ExitQuantifiers(c *QuantifiersContext)

	// ExitMeasuresClause is called when exiting the measuresClause production.
	ExitMeasuresClause(c *MeasuresClauseContext)

	// ExitPatternDefinition is called when exiting the patternDefinition production.
	ExitPatternDefinition(c *PatternDefinitionContext)

	// ExitPatternVariable is called when exiting the patternVariable production.
	ExitPatternVariable(c *PatternVariableContext)

	// ExitOutputMode is called when exiting the outputMode production.
	ExitOutputMode(c *OutputModeContext)

	// ExitAfterMatchStrategy is called when exiting the afterMatchStrategy production.
	ExitAfterMatchStrategy(c *AfterMatchStrategyContext)

	// ExitPatternVariablesDefinition is called when exiting the patternVariablesDefinition production.
	ExitPatternVariablesDefinition(c *PatternVariablesDefinitionContext)

	// ExitWindowFrame is called when exiting the windowFrame production.
	ExitWindowFrame(c *WindowFrameContext)

	// ExitFrameBound is called when exiting the frameBound production.
	ExitFrameBound(c *FrameBoundContext)

	// ExitWithinClause is called when exiting the withinClause production.
	ExitWithinClause(c *WithinClauseContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLogicalNot is called when exiting the logicalNot production.
	ExitLogicalNot(c *LogicalNotContext)

	// ExitPredicated is called when exiting the predicated production.
	ExitPredicated(c *PredicatedContext)

	// ExitExists is called when exiting the exists production.
	ExitExists(c *ExistsContext)

	// ExitLogicalNested is called when exiting the logicalNested production.
	ExitLogicalNested(c *LogicalNestedContext)

	// ExitLogicalBinary is called when exiting the logicalBinary production.
	ExitLogicalBinary(c *LogicalBinaryContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitLikePredicate is called when exiting the likePredicate production.
	ExitLikePredicate(c *LikePredicateContext)

	// ExitValueExpressionDefault is called when exiting the valueExpressionDefault production.
	ExitValueExpressionDefault(c *ValueExpressionDefaultContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitArithmeticBinary is called when exiting the arithmeticBinary production.
	ExitArithmeticBinary(c *ArithmeticBinaryContext)

	// ExitArithmeticUnary is called when exiting the arithmeticUnary production.
	ExitArithmeticUnary(c *ArithmeticUnaryContext)

	// ExitFunctionCallExpression is called when exiting the functionCallExpression production.
	ExitFunctionCallExpression(c *FunctionCallExpressionContext)

	// ExitDereference is called when exiting the dereference production.
	ExitDereference(c *DereferenceContext)

	// ExitSimpleCase is called when exiting the simpleCase production.
	ExitSimpleCase(c *SimpleCaseContext)

	// ExitColumnReference is called when exiting the columnReference production.
	ExitColumnReference(c *ColumnReferenceContext)

	// ExitLast is called when exiting the last production.
	ExitLast(c *LastContext)

	// ExitStar is called when exiting the star production.
	ExitStar(c *StarContext)

	// ExitSubscript is called when exiting the subscript production.
	ExitSubscript(c *SubscriptContext)

	// ExitSubqueryExpression is called when exiting the subqueryExpression production.
	ExitSubqueryExpression(c *SubqueryExpressionContext)

	// ExitCast is called when exiting the cast production.
	ExitCast(c *CastContext)

	// ExitConstantDefault is called when exiting the constantDefault production.
	ExitConstantDefault(c *ConstantDefaultContext)

	// ExitParenthesizedExpression is called when exiting the parenthesizedExpression production.
	ExitParenthesizedExpression(c *ParenthesizedExpressionContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitSearchedCase is called when exiting the searchedCase production.
	ExitSearchedCase(c *SearchedCaseContext)

	// ExitPosition is called when exiting the position production.
	ExitPosition(c *PositionContext)

	// ExitFirst is called when exiting the first production.
	ExitFirst(c *FirstContext)

	// ExitFunctionNameCreate is called when exiting the functionNameCreate production.
	ExitFunctionNameCreate(c *FunctionNameCreateContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitFunctionNameAndParams is called when exiting the functionNameAndParams production.
	ExitFunctionNameAndParams(c *FunctionNameAndParamsContext)

	// ExitFunctionNameWithParams is called when exiting the functionNameWithParams production.
	ExitFunctionNameWithParams(c *FunctionNameWithParamsContext)

	// ExitFunctionParam is called when exiting the functionParam production.
	ExitFunctionParam(c *FunctionParamContext)

	// ExitDereferenceDefinition is called when exiting the dereferenceDefinition production.
	ExitDereferenceDefinition(c *DereferenceDefinitionContext)

	// ExitCorrelationName is called when exiting the correlationName production.
	ExitCorrelationName(c *CorrelationNameContext)

	// ExitQualifiedName is called when exiting the qualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)

	// ExitTimeIntervalExpression is called when exiting the timeIntervalExpression production.
	ExitTimeIntervalExpression(c *TimeIntervalExpressionContext)

	// ExitErrorCapturingMultiUnitsInterval is called when exiting the errorCapturingMultiUnitsInterval production.
	ExitErrorCapturingMultiUnitsInterval(c *ErrorCapturingMultiUnitsIntervalContext)

	// ExitMultiUnitsInterval is called when exiting the multiUnitsInterval production.
	ExitMultiUnitsInterval(c *MultiUnitsIntervalContext)

	// ExitErrorCapturingUnitToUnitInterval is called when exiting the errorCapturingUnitToUnitInterval production.
	ExitErrorCapturingUnitToUnitInterval(c *ErrorCapturingUnitToUnitIntervalContext)

	// ExitUnitToUnitInterval is called when exiting the unitToUnitInterval production.
	ExitUnitToUnitInterval(c *UnitToUnitIntervalContext)

	// ExitIntervalValue is called when exiting the intervalValue production.
	ExitIntervalValue(c *IntervalValueContext)

	// ExitTableAlias is called when exiting the tableAlias production.
	ExitTableAlias(c *TableAliasContext)

	// ExitErrorCapturingIdentifier is called when exiting the errorCapturingIdentifier production.
	ExitErrorCapturingIdentifier(c *ErrorCapturingIdentifierContext)

	// ExitErrorIdent is called when exiting the errorIdent production.
	ExitErrorIdent(c *ErrorIdentContext)

	// ExitRealIdent is called when exiting the realIdent production.
	ExitRealIdent(c *RealIdentContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitIdentifierSeq is called when exiting the identifierSeq production.
	ExitIdentifierSeq(c *IdentifierSeqContext)

	// ExitUnquotedIdentifierAlternative is called when exiting the unquotedIdentifierAlternative production.
	ExitUnquotedIdentifierAlternative(c *UnquotedIdentifierAlternativeContext)

	// ExitQuotedIdentifierAlternative is called when exiting the quotedIdentifierAlternative production.
	ExitQuotedIdentifierAlternative(c *QuotedIdentifierAlternativeContext)

	// ExitNonReservedKeywordsAlternative is called when exiting the nonReservedKeywordsAlternative production.
	ExitNonReservedKeywordsAlternative(c *NonReservedKeywordsAlternativeContext)

	// ExitUnquotedIdentifier is called when exiting the unquotedIdentifier production.
	ExitUnquotedIdentifier(c *UnquotedIdentifierContext)

	// ExitQuotedIdentifier is called when exiting the quotedIdentifier production.
	ExitQuotedIdentifier(c *QuotedIdentifierContext)

	// ExitWhenClause is called when exiting the whenClause production.
	ExitWhenClause(c *WhenClauseContext)

	// ExitCatalogPath is called when exiting the catalogPath production.
	ExitCatalogPath(c *CatalogPathContext)

	// ExitCatalogPathCreate is called when exiting the catalogPathCreate production.
	ExitCatalogPathCreate(c *CatalogPathCreateContext)

	// ExitDatabasePath is called when exiting the databasePath production.
	ExitDatabasePath(c *DatabasePathContext)

	// ExitDatabasePathCreate is called when exiting the databasePathCreate production.
	ExitDatabasePathCreate(c *DatabasePathCreateContext)

	// ExitTablePathCreate is called when exiting the tablePathCreate production.
	ExitTablePathCreate(c *TablePathCreateContext)

	// ExitTablePath is called when exiting the tablePath production.
	ExitTablePath(c *TablePathContext)

	// ExitViewPath is called when exiting the viewPath production.
	ExitViewPath(c *ViewPathContext)

	// ExitViewPathCreate is called when exiting the viewPathCreate production.
	ExitViewPathCreate(c *ViewPathCreateContext)

	// ExitUid is called when exiting the uid production.
	ExitUid(c *UidContext)

	// ExitWithOption is called when exiting the withOption production.
	ExitWithOption(c *WithOptionContext)

	// ExitIfNotExists is called when exiting the ifNotExists production.
	ExitIfNotExists(c *IfNotExistsContext)

	// ExitIfExists is called when exiting the ifExists production.
	ExitIfExists(c *IfExistsContext)

	// ExitTablePropertyList is called when exiting the tablePropertyList production.
	ExitTablePropertyList(c *TablePropertyListContext)

	// ExitTableProperty is called when exiting the tableProperty production.
	ExitTableProperty(c *TablePropertyContext)

	// ExitTablePropertyKey is called when exiting the tablePropertyKey production.
	ExitTablePropertyKey(c *TablePropertyKeyContext)

	// ExitTablePropertyValue is called when exiting the tablePropertyValue production.
	ExitTablePropertyValue(c *TablePropertyValueContext)

	// ExitLogicalOperator is called when exiting the logicalOperator production.
	ExitLogicalOperator(c *LogicalOperatorContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitBitOperator is called when exiting the bitOperator production.
	ExitBitOperator(c *BitOperatorContext)

	// ExitMathOperator is called when exiting the mathOperator production.
	ExitMathOperator(c *MathOperatorContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitTimePointLiteral is called when exiting the timePointLiteral production.
	ExitTimePointLiteral(c *TimePointLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitDecimalLiteral is called when exiting the decimalLiteral production.
	ExitDecimalLiteral(c *DecimalLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitSetQuantifier is called when exiting the setQuantifier production.
	ExitSetQuantifier(c *SetQuantifierContext)

	// ExitTimePointUnit is called when exiting the timePointUnit production.
	ExitTimePointUnit(c *TimePointUnitContext)

	// ExitTimeIntervalUnit is called when exiting the timeIntervalUnit production.
	ExitTimeIntervalUnit(c *TimeIntervalUnitContext)

	// ExitReservedKeywordsUsedAsFuncParam is called when exiting the reservedKeywordsUsedAsFuncParam production.
	ExitReservedKeywordsUsedAsFuncParam(c *ReservedKeywordsUsedAsFuncParamContext)

	// ExitReservedKeywordsNoParamsUsedAsFuncName is called when exiting the reservedKeywordsNoParamsUsedAsFuncName production.
	ExitReservedKeywordsNoParamsUsedAsFuncName(c *ReservedKeywordsNoParamsUsedAsFuncNameContext)

	// ExitReservedKeywordsFollowParamsUsedAsFuncName is called when exiting the reservedKeywordsFollowParamsUsedAsFuncName production.
	ExitReservedKeywordsFollowParamsUsedAsFuncName(c *ReservedKeywordsFollowParamsUsedAsFuncNameContext)

	// ExitReservedKeywordsUsedAsFuncName is called when exiting the reservedKeywordsUsedAsFuncName production.
	ExitReservedKeywordsUsedAsFuncName(c *ReservedKeywordsUsedAsFuncNameContext)

	// ExitNonReservedKeywords is called when exiting the nonReservedKeywords production.
	ExitNonReservedKeywords(c *NonReservedKeywordsContext)
}
