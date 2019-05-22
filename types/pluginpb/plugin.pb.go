// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/compiler/plugin.proto

package pluginpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoregistry "google.golang.org/protobuf/reflect/protoregistry"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	sync "sync"
)

const (
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 0)
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(0 - protoimpl.MinVersion)
)

// The version number of protocol compiler.
type Version struct {
	Major *int32 `protobuf:"varint,1,opt,name=major" json:"major,omitempty"`
	Minor *int32 `protobuf:"varint,2,opt,name=minor" json:"minor,omitempty"`
	Patch *int32 `protobuf:"varint,3,opt,name=patch" json:"patch,omitempty"`
	// A suffix for alpha, beta or rc release, e.g., "alpha-1", "rc2". It should
	// be empty for mainline stable releases.
	Suffix               *string                 `protobuf:"bytes,4,opt,name=suffix" json:"suffix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *Version) Reset() {
	*x = Version{}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	return file_google_protobuf_compiler_plugin_proto_msgTypes[0].MessageOf(x)
}

func (m *Version) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_compiler_plugin_proto_msgTypes[0].Methods()
}

// Deprecated: Use Version.ProtoReflect.Type instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_google_protobuf_compiler_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetMajor() int32 {
	if x != nil && x.Major != nil {
		return *x.Major
	}
	return 0
}

func (x *Version) GetMinor() int32 {
	if x != nil && x.Minor != nil {
		return *x.Minor
	}
	return 0
}

func (x *Version) GetPatch() int32 {
	if x != nil && x.Patch != nil {
		return *x.Patch
	}
	return 0
}

func (x *Version) GetSuffix() string {
	if x != nil && x.Suffix != nil {
		return *x.Suffix
	}
	return ""
}

// An encoded CodeGeneratorRequest is written to the plugin's stdin.
type CodeGeneratorRequest struct {
	// The .proto files that were explicitly listed on the command-line.  The
	// code generator should generate code only for these files.  Each file's
	// descriptor will be included in proto_file, below.
	FileToGenerate []string `protobuf:"bytes,1,rep,name=file_to_generate,json=fileToGenerate" json:"file_to_generate,omitempty"`
	// The generator parameter passed on the command-line.
	Parameter *string `protobuf:"bytes,2,opt,name=parameter" json:"parameter,omitempty"`
	// FileDescriptorProtos for all files in files_to_generate and everything
	// they import.  The files will appear in topological order, so each file
	// appears before any file that imports it.
	//
	// protoc guarantees that all proto_files will be written after
	// the fields above, even though this is not technically guaranteed by the
	// protobuf wire format.  This theoretically could allow a plugin to stream
	// in the FileDescriptorProtos and handle them one by one rather than read
	// the entire set into memory at once.  However, as of this writing, this
	// is not similarly optimized on protoc's end -- it will store all fields in
	// memory at once before sending them to the plugin.
	//
	// Type names of fields and extensions in the FileDescriptorProto are always
	// fully qualified.
	ProtoFile []*descriptorpb.FileDescriptorProto `protobuf:"bytes,15,rep,name=proto_file,json=protoFile" json:"proto_file,omitempty"`
	// The version number of protocol compiler.
	CompilerVersion      *Version                `protobuf:"bytes,3,opt,name=compiler_version,json=compilerVersion" json:"compiler_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *CodeGeneratorRequest) Reset() {
	*x = CodeGeneratorRequest{}
}

func (x *CodeGeneratorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeGeneratorRequest) ProtoMessage() {}

func (x *CodeGeneratorRequest) ProtoReflect() protoreflect.Message {
	return file_google_protobuf_compiler_plugin_proto_msgTypes[1].MessageOf(x)
}

func (m *CodeGeneratorRequest) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_compiler_plugin_proto_msgTypes[1].Methods()
}

// Deprecated: Use CodeGeneratorRequest.ProtoReflect.Type instead.
func (*CodeGeneratorRequest) Descriptor() ([]byte, []int) {
	return file_google_protobuf_compiler_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *CodeGeneratorRequest) GetFileToGenerate() []string {
	if x != nil {
		return x.FileToGenerate
	}
	return nil
}

func (x *CodeGeneratorRequest) GetParameter() string {
	if x != nil && x.Parameter != nil {
		return *x.Parameter
	}
	return ""
}

func (x *CodeGeneratorRequest) GetProtoFile() []*descriptorpb.FileDescriptorProto {
	if x != nil {
		return x.ProtoFile
	}
	return nil
}

func (x *CodeGeneratorRequest) GetCompilerVersion() *Version {
	if x != nil {
		return x.CompilerVersion
	}
	return nil
}

// The plugin writes an encoded CodeGeneratorResponse to stdout.
type CodeGeneratorResponse struct {
	// Error message.  If non-empty, code generation failed.  The plugin process
	// should exit with status code zero even if it reports an error in this way.
	//
	// This should be used to indicate errors in .proto files which prevent the
	// code generator from generating correct code.  Errors which indicate a
	// problem in protoc itself -- such as the input CodeGeneratorRequest being
	// unparseable -- should be reported by writing a message to stderr and
	// exiting with a non-zero status code.
	Error                *string                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	File                 []*CodeGeneratorResponse_File `protobuf:"bytes,15,rep,name=file" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields       `json:"-"`
	XXX_sizecache        protoimpl.SizeCache           `json:"-"`
}

func (x *CodeGeneratorResponse) Reset() {
	*x = CodeGeneratorResponse{}
}

func (x *CodeGeneratorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeGeneratorResponse) ProtoMessage() {}

func (x *CodeGeneratorResponse) ProtoReflect() protoreflect.Message {
	return file_google_protobuf_compiler_plugin_proto_msgTypes[2].MessageOf(x)
}

func (m *CodeGeneratorResponse) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_compiler_plugin_proto_msgTypes[2].Methods()
}

// Deprecated: Use CodeGeneratorResponse.ProtoReflect.Type instead.
func (*CodeGeneratorResponse) Descriptor() ([]byte, []int) {
	return file_google_protobuf_compiler_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *CodeGeneratorResponse) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

func (x *CodeGeneratorResponse) GetFile() []*CodeGeneratorResponse_File {
	if x != nil {
		return x.File
	}
	return nil
}

// Represents a single generated file.
type CodeGeneratorResponse_File struct {
	// The file name, relative to the output directory.  The name must not
	// contain "." or ".." components and must be relative, not be absolute (so,
	// the file cannot lie outside the output directory).  "/" must be used as
	// the path separator, not "\".
	//
	// If the name is omitted, the content will be appended to the previous
	// file.  This allows the generator to break large files into small chunks,
	// and allows the generated text to be streamed back to protoc so that large
	// files need not reside completely in memory at one time.  Note that as of
	// this writing protoc does not optimize for this -- it will read the entire
	// CodeGeneratorResponse before writing files to disk.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// If non-empty, indicates that the named file should already exist, and the
	// content here is to be inserted into that file at a defined insertion
	// point.  This feature allows a code generator to extend the output
	// produced by another code generator.  The original generator may provide
	// insertion points by placing special annotations in the file that look
	// like:
	//   @@protoc_insertion_point(NAME)
	// The annotation can have arbitrary text before and after it on the line,
	// which allows it to be placed in a comment.  NAME should be replaced with
	// an identifier naming the point -- this is what other generators will use
	// as the insertion_point.  Code inserted at this point will be placed
	// immediately above the line containing the insertion point (thus multiple
	// insertions to the same point will come out in the order they were added).
	// The double-@ is intended to make it unlikely that the generated code
	// could contain things that look like insertion points by accident.
	//
	// For example, the C++ code generator places the following line in the
	// .pb.h files that it generates:
	//   // @@protoc_insertion_point(namespace_scope)
	// This line appears within the scope of the file's package namespace, but
	// outside of any particular class.  Another plugin can then specify the
	// insertion_point "namespace_scope" to generate additional classes or
	// other declarations that should be placed in this scope.
	//
	// Note that if the line containing the insertion point begins with
	// whitespace, the same whitespace will be added to every line of the
	// inserted text.  This is useful for languages like Python, where
	// indentation matters.  In these languages, the insertion point comment
	// should be indented the same amount as any inserted code will need to be
	// in order to work correctly in that context.
	//
	// The code generator that generates the initial file and the one which
	// inserts into it must both run as part of a single invocation of protoc.
	// Code generators are executed in the order in which they appear on the
	// command line.
	//
	// If |insertion_point| is present, |name| must also be present.
	InsertionPoint *string `protobuf:"bytes,2,opt,name=insertion_point,json=insertionPoint" json:"insertion_point,omitempty"`
	// The file contents.
	Content              *string                 `protobuf:"bytes,15,opt,name=content" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *CodeGeneratorResponse_File) Reset() {
	*x = CodeGeneratorResponse_File{}
}

func (x *CodeGeneratorResponse_File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeGeneratorResponse_File) ProtoMessage() {}

func (x *CodeGeneratorResponse_File) ProtoReflect() protoreflect.Message {
	return file_google_protobuf_compiler_plugin_proto_msgTypes[3].MessageOf(x)
}

func (m *CodeGeneratorResponse_File) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_compiler_plugin_proto_msgTypes[3].Methods()
}

// Deprecated: Use CodeGeneratorResponse_File.ProtoReflect.Type instead.
func (*CodeGeneratorResponse_File) Descriptor() ([]byte, []int) {
	return file_google_protobuf_compiler_plugin_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CodeGeneratorResponse_File) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CodeGeneratorResponse_File) GetInsertionPoint() string {
	if x != nil && x.InsertionPoint != nil {
		return *x.InsertionPoint
	}
	return ""
}

func (x *CodeGeneratorResponse_File) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

var File_google_protobuf_compiler_plugin_proto protoreflect.FileDescriptor

var file_google_protobuf_compiler_plugin_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x72, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d,
	0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x22, 0xf1, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x64,
	0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x69, 0x6c,
	0x65, 0x54, 0x6f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x4c,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x65, 0x72, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xd6, 0x01, 0x0a,
	0x15, 0x43, 0x6f, 0x64, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x48, 0x0a, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x5d, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x57, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x42, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x5a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x70, 0x62,
}

var (
	file_google_protobuf_compiler_plugin_proto_rawDescOnce sync.Once
	file_google_protobuf_compiler_plugin_proto_rawDescData = file_google_protobuf_compiler_plugin_proto_rawDesc
)

func file_google_protobuf_compiler_plugin_proto_rawDescGZIP() []byte {
	file_google_protobuf_compiler_plugin_proto_rawDescOnce.Do(func() {
		file_google_protobuf_compiler_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_protobuf_compiler_plugin_proto_rawDescData)
	})
	return file_google_protobuf_compiler_plugin_proto_rawDescData
}

var file_google_protobuf_compiler_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_protobuf_compiler_plugin_proto_goTypes = []interface{}{
	(*Version)(nil),                          // 0: google.protobuf.compiler.Version
	(*CodeGeneratorRequest)(nil),             // 1: google.protobuf.compiler.CodeGeneratorRequest
	(*CodeGeneratorResponse)(nil),            // 2: google.protobuf.compiler.CodeGeneratorResponse
	(*CodeGeneratorResponse_File)(nil),       // 3: google.protobuf.compiler.CodeGeneratorResponse.File
	(*descriptorpb.FileDescriptorProto)(nil), // 4: google.protobuf.FileDescriptorProto
}
var file_google_protobuf_compiler_plugin_proto_depIdxs = []int32{
	4, // google.protobuf.compiler.CodeGeneratorRequest.proto_file:type_name -> google.protobuf.FileDescriptorProto
	0, // google.protobuf.compiler.CodeGeneratorRequest.compiler_version:type_name -> google.protobuf.compiler.Version
	3, // google.protobuf.compiler.CodeGeneratorResponse.file:type_name -> google.protobuf.compiler.CodeGeneratorResponse.File
}

func init() { file_google_protobuf_compiler_plugin_proto_init() }
func file_google_protobuf_compiler_plugin_proto_init() {
	if File_google_protobuf_compiler_plugin_proto != nil {
		return
	}
	File_google_protobuf_compiler_plugin_proto = protoimpl.FileBuilder{
		RawDescriptor:      file_google_protobuf_compiler_plugin_proto_rawDesc,
		GoTypes:            file_google_protobuf_compiler_plugin_proto_goTypes,
		DependencyIndexes:  file_google_protobuf_compiler_plugin_proto_depIdxs,
		MessageOutputTypes: file_google_protobuf_compiler_plugin_proto_msgTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	file_google_protobuf_compiler_plugin_proto_rawDesc = nil
	file_google_protobuf_compiler_plugin_proto_goTypes = nil
	file_google_protobuf_compiler_plugin_proto_depIdxs = nil
}
