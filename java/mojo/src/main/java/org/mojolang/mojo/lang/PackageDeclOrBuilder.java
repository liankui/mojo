// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/package_decl.proto

package org.mojolang.mojo.lang;

public interface PackageDeclOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.PackageDecl)
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
   *&#47; reserved field no. 3 for string document
   * </pre>
   *
   * <code>.mojo.lang.Document document = 4;</code>
   */
  boolean hasDocument();
  /**
   * <pre>
   *&#47; reserved field no. 3 for string document
   * </pre>
   *
   * <code>.mojo.lang.Document document = 4;</code>
   */
  org.mojolang.mojo.lang.Document getDocument();
  /**
   * <pre>
   *&#47; reserved field no. 3 for string document
   * </pre>
   *
   * <code>.mojo.lang.Document document = 4;</code>
   */
  org.mojolang.mojo.lang.DocumentOrBuilder getDocumentOrBuilder();

  /**
   * <code>string name = 10;</code>
   */
  java.lang.String getName();
  /**
   * <code>string name = 10;</code>
   */
  com.google.protobuf.ByteString
      getNameBytes();

  /**
   * <code>.mojo.lang.ObjectLiteralExpr package = 11;</code>
   */
  boolean hasPackage();
  /**
   * <code>.mojo.lang.ObjectLiteralExpr package = 11;</code>
   */
  org.mojolang.mojo.lang.ObjectLiteralExpr getPackage();
  /**
   * <code>.mojo.lang.ObjectLiteralExpr package = 11;</code>
   */
  org.mojolang.mojo.lang.ObjectLiteralExprOrBuilder getPackageOrBuilder();
}