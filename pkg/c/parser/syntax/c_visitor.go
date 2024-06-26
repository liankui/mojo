// Code generated from C.g4 by ANTLR 4.13.0. DO NOT EDIT.

package syntax // C
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CParser.
type CVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CParser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by CParser#includeDirective.
	VisitIncludeDirective(ctx *IncludeDirectiveContext) interface{}

	// Visit a parse tree produced by CParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by CParser#genericSelection.
	VisitGenericSelection(ctx *GenericSelectionContext) interface{}

	// Visit a parse tree produced by CParser#genericAssocList.
	VisitGenericAssocList(ctx *GenericAssocListContext) interface{}

	// Visit a parse tree produced by CParser#genericAssociation.
	VisitGenericAssociation(ctx *GenericAssociationContext) interface{}

	// Visit a parse tree produced by CParser#postfixExpression.
	VisitPostfixExpression(ctx *PostfixExpressionContext) interface{}

	// Visit a parse tree produced by CParser#argumentExpressionList.
	VisitArgumentExpressionList(ctx *ArgumentExpressionListContext) interface{}

	// Visit a parse tree produced by CParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by CParser#unaryOperator.
	VisitUnaryOperator(ctx *UnaryOperatorContext) interface{}

	// Visit a parse tree produced by CParser#castExpression.
	VisitCastExpression(ctx *CastExpressionContext) interface{}

	// Visit a parse tree produced by CParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by CParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by CParser#shiftExpression.
	VisitShiftExpression(ctx *ShiftExpressionContext) interface{}

	// Visit a parse tree produced by CParser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by CParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by CParser#andExpression.
	VisitAndExpression(ctx *AndExpressionContext) interface{}

	// Visit a parse tree produced by CParser#exclusiveOrExpression.
	VisitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) interface{}

	// Visit a parse tree produced by CParser#inclusiveOrExpression.
	VisitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) interface{}

	// Visit a parse tree produced by CParser#logicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by CParser#logicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by CParser#conditionalExpression.
	VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{}

	// Visit a parse tree produced by CParser#assignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by CParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by CParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by CParser#constantExpression.
	VisitConstantExpression(ctx *ConstantExpressionContext) interface{}

	// Visit a parse tree produced by CParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by CParser#declarationSpecifiers.
	VisitDeclarationSpecifiers(ctx *DeclarationSpecifiersContext) interface{}

	// Visit a parse tree produced by CParser#declarationSpecifiers2.
	VisitDeclarationSpecifiers2(ctx *DeclarationSpecifiers2Context) interface{}

	// Visit a parse tree produced by CParser#declarationSpecifier.
	VisitDeclarationSpecifier(ctx *DeclarationSpecifierContext) interface{}

	// Visit a parse tree produced by CParser#initDeclaratorList.
	VisitInitDeclaratorList(ctx *InitDeclaratorListContext) interface{}

	// Visit a parse tree produced by CParser#initDeclarator.
	VisitInitDeclarator(ctx *InitDeclaratorContext) interface{}

	// Visit a parse tree produced by CParser#storageClassSpecifier.
	VisitStorageClassSpecifier(ctx *StorageClassSpecifierContext) interface{}

	// Visit a parse tree produced by CParser#typeSpecifier.
	VisitTypeSpecifier(ctx *TypeSpecifierContext) interface{}

	// Visit a parse tree produced by CParser#structOrUnionSpecifier.
	VisitStructOrUnionSpecifier(ctx *StructOrUnionSpecifierContext) interface{}

	// Visit a parse tree produced by CParser#structOrUnion.
	VisitStructOrUnion(ctx *StructOrUnionContext) interface{}

	// Visit a parse tree produced by CParser#structDeclarationList.
	VisitStructDeclarationList(ctx *StructDeclarationListContext) interface{}

	// Visit a parse tree produced by CParser#structDeclaration.
	VisitStructDeclaration(ctx *StructDeclarationContext) interface{}

	// Visit a parse tree produced by CParser#specifierQualifierList.
	VisitSpecifierQualifierList(ctx *SpecifierQualifierListContext) interface{}

	// Visit a parse tree produced by CParser#structDeclaratorList.
	VisitStructDeclaratorList(ctx *StructDeclaratorListContext) interface{}

	// Visit a parse tree produced by CParser#structDeclarator.
	VisitStructDeclarator(ctx *StructDeclaratorContext) interface{}

	// Visit a parse tree produced by CParser#enumSpecifier.
	VisitEnumSpecifier(ctx *EnumSpecifierContext) interface{}

	// Visit a parse tree produced by CParser#enumeratorList.
	VisitEnumeratorList(ctx *EnumeratorListContext) interface{}

	// Visit a parse tree produced by CParser#enumerator.
	VisitEnumerator(ctx *EnumeratorContext) interface{}

	// Visit a parse tree produced by CParser#enumerationConstant.
	VisitEnumerationConstant(ctx *EnumerationConstantContext) interface{}

	// Visit a parse tree produced by CParser#atomicTypeSpecifier.
	VisitAtomicTypeSpecifier(ctx *AtomicTypeSpecifierContext) interface{}

	// Visit a parse tree produced by CParser#typeQualifier.
	VisitTypeQualifier(ctx *TypeQualifierContext) interface{}

	// Visit a parse tree produced by CParser#functionSpecifier.
	VisitFunctionSpecifier(ctx *FunctionSpecifierContext) interface{}

	// Visit a parse tree produced by CParser#alignmentSpecifier.
	VisitAlignmentSpecifier(ctx *AlignmentSpecifierContext) interface{}

	// Visit a parse tree produced by CParser#declarator.
	VisitDeclarator(ctx *DeclaratorContext) interface{}

	// Visit a parse tree produced by CParser#directDeclarator.
	VisitDirectDeclarator(ctx *DirectDeclaratorContext) interface{}

	// Visit a parse tree produced by CParser#vcSpecificModifer.
	VisitVcSpecificModifer(ctx *VcSpecificModiferContext) interface{}

	// Visit a parse tree produced by CParser#gccDeclaratorExtension.
	VisitGccDeclaratorExtension(ctx *GccDeclaratorExtensionContext) interface{}

	// Visit a parse tree produced by CParser#gccAttributeSpecifier.
	VisitGccAttributeSpecifier(ctx *GccAttributeSpecifierContext) interface{}

	// Visit a parse tree produced by CParser#gccAttributeList.
	VisitGccAttributeList(ctx *GccAttributeListContext) interface{}

	// Visit a parse tree produced by CParser#gccAttribute.
	VisitGccAttribute(ctx *GccAttributeContext) interface{}

	// Visit a parse tree produced by CParser#nestedParenthesesBlock.
	VisitNestedParenthesesBlock(ctx *NestedParenthesesBlockContext) interface{}

	// Visit a parse tree produced by CParser#pointer.
	VisitPointer(ctx *PointerContext) interface{}

	// Visit a parse tree produced by CParser#typeQualifierList.
	VisitTypeQualifierList(ctx *TypeQualifierListContext) interface{}

	// Visit a parse tree produced by CParser#parameterTypeList.
	VisitParameterTypeList(ctx *ParameterTypeListContext) interface{}

	// Visit a parse tree produced by CParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by CParser#parameterDeclaration.
	VisitParameterDeclaration(ctx *ParameterDeclarationContext) interface{}

	// Visit a parse tree produced by CParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by CParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by CParser#abstractDeclarator.
	VisitAbstractDeclarator(ctx *AbstractDeclaratorContext) interface{}

	// Visit a parse tree produced by CParser#directAbstractDeclarator.
	VisitDirectAbstractDeclarator(ctx *DirectAbstractDeclaratorContext) interface{}

	// Visit a parse tree produced by CParser#typedefName.
	VisitTypedefName(ctx *TypedefNameContext) interface{}

	// Visit a parse tree produced by CParser#initializer.
	VisitInitializer(ctx *InitializerContext) interface{}

	// Visit a parse tree produced by CParser#initializerList.
	VisitInitializerList(ctx *InitializerListContext) interface{}

	// Visit a parse tree produced by CParser#designation.
	VisitDesignation(ctx *DesignationContext) interface{}

	// Visit a parse tree produced by CParser#designatorList.
	VisitDesignatorList(ctx *DesignatorListContext) interface{}

	// Visit a parse tree produced by CParser#designator.
	VisitDesignator(ctx *DesignatorContext) interface{}

	// Visit a parse tree produced by CParser#staticAssertDeclaration.
	VisitStaticAssertDeclaration(ctx *StaticAssertDeclarationContext) interface{}

	// Visit a parse tree produced by CParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by CParser#labeledStatement.
	VisitLabeledStatement(ctx *LabeledStatementContext) interface{}

	// Visit a parse tree produced by CParser#compoundStatement.
	VisitCompoundStatement(ctx *CompoundStatementContext) interface{}

	// Visit a parse tree produced by CParser#blockItemList.
	VisitBlockItemList(ctx *BlockItemListContext) interface{}

	// Visit a parse tree produced by CParser#blockItem.
	VisitBlockItem(ctx *BlockItemContext) interface{}

	// Visit a parse tree produced by CParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by CParser#selectionStatement.
	VisitSelectionStatement(ctx *SelectionStatementContext) interface{}

	// Visit a parse tree produced by CParser#iterationStatement.
	VisitIterationStatement(ctx *IterationStatementContext) interface{}

	// Visit a parse tree produced by CParser#forCondition.
	VisitForCondition(ctx *ForConditionContext) interface{}

	// Visit a parse tree produced by CParser#forDeclaration.
	VisitForDeclaration(ctx *ForDeclarationContext) interface{}

	// Visit a parse tree produced by CParser#forExpression.
	VisitForExpression(ctx *ForExpressionContext) interface{}

	// Visit a parse tree produced by CParser#jumpStatement.
	VisitJumpStatement(ctx *JumpStatementContext) interface{}

	// Visit a parse tree produced by CParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by CParser#translationUnit.
	VisitTranslationUnit(ctx *TranslationUnitContext) interface{}

	// Visit a parse tree produced by CParser#externalDeclaration.
	VisitExternalDeclaration(ctx *ExternalDeclarationContext) interface{}

	// Visit a parse tree produced by CParser#functionDefinition.
	VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{}

	// Visit a parse tree produced by CParser#declarationList.
	VisitDeclarationList(ctx *DeclarationListContext) interface{}
}
