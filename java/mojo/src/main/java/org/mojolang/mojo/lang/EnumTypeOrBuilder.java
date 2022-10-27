// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/type.proto

package org.mojolang.mojo.lang;

public interface EnumTypeOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.EnumType)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   *&#47; position of first character belonging to the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position start_position = 1;</code>
   */
  boolean hasStartPosition();
  /**
   * <pre>
   *&#47; position of first character belonging to the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position start_position = 1;</code>
   */
  org.mojolang.mojo.lang.Position getStartPosition();
  /**
   * <pre>
   *&#47; position of first character belonging to the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position start_position = 1;</code>
   */
  org.mojolang.mojo.lang.PositionOrBuilder getStartPositionOrBuilder();

  /**
   * <pre>
   *&#47; position of first character immediately after the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position end_position = 2;</code>
   */
  boolean hasEndPosition();
  /**
   * <pre>
   *&#47; position of first character immediately after the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position end_position = 2;</code>
   */
  org.mojolang.mojo.lang.Position getEndPosition();
  /**
   * <pre>
   *&#47; position of first character immediately after the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position end_position = 2;</code>
   */
  org.mojolang.mojo.lang.PositionOrBuilder getEndPositionOrBuilder();

  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>string package = 6;</code>
   */
  java.lang.String getPackage();
  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>string package = 6;</code>
   */
  com.google.protobuf.ByteString
      getPackageBytes();

  /**
   * <pre>
   * </pre>
   *
   * <code>repeated .mojo.lang.ValueDecl enumerators = 10;</code>
   */
  java.util.List<org.mojolang.mojo.lang.ValueDecl> 
      getEnumeratorsList();
  /**
   * <pre>
   * </pre>
   *
   * <code>repeated .mojo.lang.ValueDecl enumerators = 10;</code>
   */
  org.mojolang.mojo.lang.ValueDecl getEnumerators(int index);
  /**
   * <pre>
   * </pre>
   *
   * <code>repeated .mojo.lang.ValueDecl enumerators = 10;</code>
   */
  int getEnumeratorsCount();
  /**
   * <pre>
   * </pre>
   *
   * <code>repeated .mojo.lang.ValueDecl enumerators = 10;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.ValueDeclOrBuilder> 
      getEnumeratorsOrBuilderList();
  /**
   * <pre>
   * </pre>
   *
   * <code>repeated .mojo.lang.ValueDecl enumerators = 10;</code>
   */
  org.mojolang.mojo.lang.ValueDeclOrBuilder getEnumeratorsOrBuilder(
      int index);

  /**
   * <pre>
   * </pre>
   *
   * <code>.mojo.lang.NominalType underlying_type = 11;</code>
   */
  boolean hasUnderlyingType();
  /**
   * <pre>
   * </pre>
   *
   * <code>.mojo.lang.NominalType underlying_type = 11;</code>
   */
  org.mojolang.mojo.lang.NominalType getUnderlyingType();
  /**
   * <pre>
   * </pre>
   *
   * <code>.mojo.lang.NominalType underlying_type = 11;</code>
   */
  org.mojolang.mojo.lang.NominalTypeOrBuilder getUnderlyingTypeOrBuilder();
}