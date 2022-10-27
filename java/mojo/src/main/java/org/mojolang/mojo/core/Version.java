// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/version.proto

package org.mojolang.mojo.core;

/**
 * <pre>
 *&#47; Semantic Versioning 2.0.0
 * /
 * / Given a version number MAJOR.MINOR.PATCH, increment the:
 * /
 * / MAJOR version when you make incompatible API changes,
 * / MINOR version when you add functionality in a backwards compatible manner, and
 * / PATCH version when you make backwards compatible bug fixes.
 * / Additional labels for pre-release and build metadata are available as extensions to the MAJOR.MINOR.PATCH format.
 * /
 * </pre>
 *
 * Protobuf type {@code mojo.core.Version}
 */
public  final class Version extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.core.Version)
    VersionOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Version.newBuilder() to construct.
  private Version(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Version() {
    major_ = 0;
    minor_ = 0;
    patch_ = 0;
    preRelease_ = "";
    build_ = "";
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private Version(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    int mutable_bitField0_ = 0;
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 8: {

            major_ = input.readInt32();
            break;
          }
          case 16: {

            minor_ = input.readInt32();
            break;
          }
          case 24: {

            patch_ = input.readInt32();
            break;
          }
          case 34: {
            java.lang.String s = input.readStringRequireUtf8();

            preRelease_ = s;
            break;
          }
          case 42: {
            java.lang.String s = input.readStringRequireUtf8();

            build_ = s;
            break;
          }
          default: {
            if (!parseUnknownFieldProto3(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.mojolang.mojo.core.VersionProto.internal_static_mojo_core_Version_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.core.VersionProto.internal_static_mojo_core_Version_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.core.Version.class, org.mojolang.mojo.core.Version.Builder.class);
  }

  public static final int MAJOR_FIELD_NUMBER = 1;
  private int major_;
  /**
   * <pre>
   *&#47; major version
   * </pre>
   *
   * <code>int32 major = 1;</code>
   */
  public int getMajor() {
    return major_;
  }

  public static final int MINOR_FIELD_NUMBER = 2;
  private int minor_;
  /**
   * <pre>
   *&#47; minor version
   * </pre>
   *
   * <code>int32 minor = 2;</code>
   */
  public int getMinor() {
    return minor_;
  }

  public static final int PATCH_FIELD_NUMBER = 3;
  private int patch_;
  /**
   * <pre>
   *&#47; patch version
   * </pre>
   *
   * <code>int32 patch = 3;</code>
   */
  public int getPatch() {
    return patch_;
  }

  public static final int PRE_RELEASE_FIELD_NUMBER = 4;
  private volatile java.lang.Object preRelease_;
  /**
   * <pre>
   *&#47; pre-release version
   * </pre>
   *
   * <code>string pre_release = 4;</code>
   */
  public java.lang.String getPreRelease() {
    java.lang.Object ref = preRelease_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      preRelease_ = s;
      return s;
    }
  }
  /**
   * <pre>
   *&#47; pre-release version
   * </pre>
   *
   * <code>string pre_release = 4;</code>
   */
  public com.google.protobuf.ByteString
      getPreReleaseBytes() {
    java.lang.Object ref = preRelease_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      preRelease_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int BUILD_FIELD_NUMBER = 5;
  private volatile java.lang.Object build_;
  /**
   * <pre>
   *&#47; build metadata
   * </pre>
   *
   * <code>string build = 5;</code>
   */
  public java.lang.String getBuild() {
    java.lang.Object ref = build_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      build_ = s;
      return s;
    }
  }
  /**
   * <pre>
   *&#47; build metadata
   * </pre>
   *
   * <code>string build = 5;</code>
   */
  public com.google.protobuf.ByteString
      getBuildBytes() {
    java.lang.Object ref = build_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      build_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (major_ != 0) {
      output.writeInt32(1, major_);
    }
    if (minor_ != 0) {
      output.writeInt32(2, minor_);
    }
    if (patch_ != 0) {
      output.writeInt32(3, patch_);
    }
    if (!getPreReleaseBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 4, preRelease_);
    }
    if (!getBuildBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 5, build_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (major_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(1, major_);
    }
    if (minor_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(2, minor_);
    }
    if (patch_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(3, patch_);
    }
    if (!getPreReleaseBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(4, preRelease_);
    }
    if (!getBuildBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(5, build_);
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof org.mojolang.mojo.core.Version)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.core.Version other = (org.mojolang.mojo.core.Version) obj;

    boolean result = true;
    result = result && (getMajor()
        == other.getMajor());
    result = result && (getMinor()
        == other.getMinor());
    result = result && (getPatch()
        == other.getPatch());
    result = result && getPreRelease()
        .equals(other.getPreRelease());
    result = result && getBuild()
        .equals(other.getBuild());
    result = result && unknownFields.equals(other.unknownFields);
    return result;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + MAJOR_FIELD_NUMBER;
    hash = (53 * hash) + getMajor();
    hash = (37 * hash) + MINOR_FIELD_NUMBER;
    hash = (53 * hash) + getMinor();
    hash = (37 * hash) + PATCH_FIELD_NUMBER;
    hash = (53 * hash) + getPatch();
    hash = (37 * hash) + PRE_RELEASE_FIELD_NUMBER;
    hash = (53 * hash) + getPreRelease().hashCode();
    hash = (37 * hash) + BUILD_FIELD_NUMBER;
    hash = (53 * hash) + getBuild().hashCode();
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.core.Version parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Version parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Version parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Version parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Version parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Version parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Version parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Version parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Version parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Version parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Version parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Version parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(org.mojolang.mojo.core.Version prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * <pre>
   *&#47; Semantic Versioning 2.0.0
   * /
   * / Given a version number MAJOR.MINOR.PATCH, increment the:
   * /
   * / MAJOR version when you make incompatible API changes,
   * / MINOR version when you add functionality in a backwards compatible manner, and
   * / PATCH version when you make backwards compatible bug fixes.
   * / Additional labels for pre-release and build metadata are available as extensions to the MAJOR.MINOR.PATCH format.
   * /
   * </pre>
   *
   * Protobuf type {@code mojo.core.Version}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.core.Version)
      org.mojolang.mojo.core.VersionOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.core.VersionProto.internal_static_mojo_core_Version_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.core.VersionProto.internal_static_mojo_core_Version_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.core.Version.class, org.mojolang.mojo.core.Version.Builder.class);
    }

    // Construct using org.mojolang.mojo.core.Version.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      major_ = 0;

      minor_ = 0;

      patch_ = 0;

      preRelease_ = "";

      build_ = "";

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.core.VersionProto.internal_static_mojo_core_Version_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Version getDefaultInstanceForType() {
      return org.mojolang.mojo.core.Version.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Version build() {
      org.mojolang.mojo.core.Version result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Version buildPartial() {
      org.mojolang.mojo.core.Version result = new org.mojolang.mojo.core.Version(this);
      result.major_ = major_;
      result.minor_ = minor_;
      result.patch_ = patch_;
      result.preRelease_ = preRelease_;
      result.build_ = build_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return (Builder) super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return (Builder) super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return (Builder) super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return (Builder) super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return (Builder) super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return (Builder) super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof org.mojolang.mojo.core.Version) {
        return mergeFrom((org.mojolang.mojo.core.Version)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.core.Version other) {
      if (other == org.mojolang.mojo.core.Version.getDefaultInstance()) return this;
      if (other.getMajor() != 0) {
        setMajor(other.getMajor());
      }
      if (other.getMinor() != 0) {
        setMinor(other.getMinor());
      }
      if (other.getPatch() != 0) {
        setPatch(other.getPatch());
      }
      if (!other.getPreRelease().isEmpty()) {
        preRelease_ = other.preRelease_;
        onChanged();
      }
      if (!other.getBuild().isEmpty()) {
        build_ = other.build_;
        onChanged();
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      org.mojolang.mojo.core.Version parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.mojolang.mojo.core.Version) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private int major_ ;
    /**
     * <pre>
     *&#47; major version
     * </pre>
     *
     * <code>int32 major = 1;</code>
     */
    public int getMajor() {
      return major_;
    }
    /**
     * <pre>
     *&#47; major version
     * </pre>
     *
     * <code>int32 major = 1;</code>
     */
    public Builder setMajor(int value) {
      
      major_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     *&#47; major version
     * </pre>
     *
     * <code>int32 major = 1;</code>
     */
    public Builder clearMajor() {
      
      major_ = 0;
      onChanged();
      return this;
    }

    private int minor_ ;
    /**
     * <pre>
     *&#47; minor version
     * </pre>
     *
     * <code>int32 minor = 2;</code>
     */
    public int getMinor() {
      return minor_;
    }
    /**
     * <pre>
     *&#47; minor version
     * </pre>
     *
     * <code>int32 minor = 2;</code>
     */
    public Builder setMinor(int value) {
      
      minor_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     *&#47; minor version
     * </pre>
     *
     * <code>int32 minor = 2;</code>
     */
    public Builder clearMinor() {
      
      minor_ = 0;
      onChanged();
      return this;
    }

    private int patch_ ;
    /**
     * <pre>
     *&#47; patch version
     * </pre>
     *
     * <code>int32 patch = 3;</code>
     */
    public int getPatch() {
      return patch_;
    }
    /**
     * <pre>
     *&#47; patch version
     * </pre>
     *
     * <code>int32 patch = 3;</code>
     */
    public Builder setPatch(int value) {
      
      patch_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     *&#47; patch version
     * </pre>
     *
     * <code>int32 patch = 3;</code>
     */
    public Builder clearPatch() {
      
      patch_ = 0;
      onChanged();
      return this;
    }

    private java.lang.Object preRelease_ = "";
    /**
     * <pre>
     *&#47; pre-release version
     * </pre>
     *
     * <code>string pre_release = 4;</code>
     */
    public java.lang.String getPreRelease() {
      java.lang.Object ref = preRelease_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        preRelease_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <pre>
     *&#47; pre-release version
     * </pre>
     *
     * <code>string pre_release = 4;</code>
     */
    public com.google.protobuf.ByteString
        getPreReleaseBytes() {
      java.lang.Object ref = preRelease_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        preRelease_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     *&#47; pre-release version
     * </pre>
     *
     * <code>string pre_release = 4;</code>
     */
    public Builder setPreRelease(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      preRelease_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     *&#47; pre-release version
     * </pre>
     *
     * <code>string pre_release = 4;</code>
     */
    public Builder clearPreRelease() {
      
      preRelease_ = getDefaultInstance().getPreRelease();
      onChanged();
      return this;
    }
    /**
     * <pre>
     *&#47; pre-release version
     * </pre>
     *
     * <code>string pre_release = 4;</code>
     */
    public Builder setPreReleaseBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      preRelease_ = value;
      onChanged();
      return this;
    }

    private java.lang.Object build_ = "";
    /**
     * <pre>
     *&#47; build metadata
     * </pre>
     *
     * <code>string build = 5;</code>
     */
    public java.lang.String getBuild() {
      java.lang.Object ref = build_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        build_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <pre>
     *&#47; build metadata
     * </pre>
     *
     * <code>string build = 5;</code>
     */
    public com.google.protobuf.ByteString
        getBuildBytes() {
      java.lang.Object ref = build_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        build_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     *&#47; build metadata
     * </pre>
     *
     * <code>string build = 5;</code>
     */
    public Builder setBuild(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      build_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     *&#47; build metadata
     * </pre>
     *
     * <code>string build = 5;</code>
     */
    public Builder clearBuild() {
      
      build_ = getDefaultInstance().getBuild();
      onChanged();
      return this;
    }
    /**
     * <pre>
     *&#47; build metadata
     * </pre>
     *
     * <code>string build = 5;</code>
     */
    public Builder setBuildBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      build_ = value;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFieldsProto3(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:mojo.core.Version)
  }

  // @@protoc_insertion_point(class_scope:mojo.core.Version)
  private static final org.mojolang.mojo.core.Version DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.core.Version();
  }

  public static org.mojolang.mojo.core.Version getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Version>
      PARSER = new com.google.protobuf.AbstractParser<Version>() {
    @java.lang.Override
    public Version parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new Version(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<Version> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Version> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.core.Version getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}
