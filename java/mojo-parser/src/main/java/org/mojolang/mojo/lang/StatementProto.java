// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/statement.proto

package org.mojolang.mojo.lang;

public final class StatementProto {
  private StatementProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_Statement_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_Statement_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\031mojo/lang/statement.proto\022\tmojo.lang\032\033" +
      "mojo/lang/declaration.proto\032\032mojo/lang/e" +
      "xpression.proto\"t\n\tStatement\022-\n\013declarat" +
      "ion\030\n \001(\0132\026.mojo.lang.DeclarationH\000\022+\n\ne" +
      "xpression\030\013 \001(\0132\025.mojo.lang.ExpressionH\000" +
      "B\013\n\tstatementB[\n\026org.mojolang.mojo.langB" +
      "\016StatementProtoP\001Z/github.com/mojo-lang/" +
      "lang/golang/mojo/lang;langb\006proto3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.lang.DeclarationProto.getDescriptor(),
          org.mojolang.mojo.lang.ExpressionProto.getDescriptor(),
        }, assigner);
    internal_static_mojo_lang_Statement_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_lang_Statement_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_Statement_descriptor,
        new java.lang.String[] { "Declaration", "Expression", "Statement", });
    org.mojolang.mojo.lang.DeclarationProto.getDescriptor();
    org.mojolang.mojo.lang.ExpressionProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}