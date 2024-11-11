// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc2
// source: weather.proto

package weather

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WeatherData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coord      *WeatherData_Coord              `protobuf:"bytes,1,opt,name=coord,proto3" json:"coord,omitempty"`
	Weather    []*WeatherData_WeatherCondition `protobuf:"bytes,2,rep,name=weather,proto3" json:"weather,omitempty"`
	Base       string                          `protobuf:"bytes,3,opt,name=base,proto3" json:"base,omitempty"`
	Main       *WeatherData_MainData           `protobuf:"bytes,4,opt,name=main,proto3" json:"main,omitempty"`
	Visibility int32                           `protobuf:"varint,5,opt,name=visibility,proto3" json:"visibility,omitempty"`
	Wind       *WeatherData_Wind               `protobuf:"bytes,6,opt,name=wind,proto3" json:"wind,omitempty"`
	Clouds     *WeatherData_Clouds             `protobuf:"bytes,7,opt,name=clouds,proto3" json:"clouds,omitempty"`
	Dt         int64                           `protobuf:"varint,8,opt,name=dt,proto3" json:"dt,omitempty"`
	Sys        *WeatherData_Sys                `protobuf:"bytes,9,opt,name=sys,proto3" json:"sys,omitempty"`
	Timezone   int32                           `protobuf:"varint,10,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Id         int32                           `protobuf:"varint,11,opt,name=id,proto3" json:"id,omitempty"`
	Name       string                          `protobuf:"bytes,12,opt,name=name,proto3" json:"name,omitempty"`
	Cod        int32                           `protobuf:"varint,13,opt,name=cod,proto3" json:"cod,omitempty"`
}

func (x *WeatherData) Reset() {
	*x = WeatherData{}
	mi := &file_weather_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WeatherData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherData) ProtoMessage() {}

func (x *WeatherData) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherData.ProtoReflect.Descriptor instead.
func (*WeatherData) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{0}
}

func (x *WeatherData) GetCoord() *WeatherData_Coord {
	if x != nil {
		return x.Coord
	}
	return nil
}

func (x *WeatherData) GetWeather() []*WeatherData_WeatherCondition {
	if x != nil {
		return x.Weather
	}
	return nil
}

func (x *WeatherData) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *WeatherData) GetMain() *WeatherData_MainData {
	if x != nil {
		return x.Main
	}
	return nil
}

func (x *WeatherData) GetVisibility() int32 {
	if x != nil {
		return x.Visibility
	}
	return 0
}

func (x *WeatherData) GetWind() *WeatherData_Wind {
	if x != nil {
		return x.Wind
	}
	return nil
}

func (x *WeatherData) GetClouds() *WeatherData_Clouds {
	if x != nil {
		return x.Clouds
	}
	return nil
}

func (x *WeatherData) GetDt() int64 {
	if x != nil {
		return x.Dt
	}
	return 0
}

func (x *WeatherData) GetSys() *WeatherData_Sys {
	if x != nil {
		return x.Sys
	}
	return nil
}

func (x *WeatherData) GetTimezone() int32 {
	if x != nil {
		return x.Timezone
	}
	return 0
}

func (x *WeatherData) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WeatherData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WeatherData) GetCod() int32 {
	if x != nil {
		return x.Cod
	}
	return 0
}

type CityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	City string `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *CityRequest) Reset() {
	*x = CityRequest{}
	mi := &file_weather_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityRequest) ProtoMessage() {}

func (x *CityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityRequest.ProtoReflect.Descriptor instead.
func (*CityRequest) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{1}
}

func (x *CityRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type WeatherData_Coord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lon float32 `protobuf:"fixed32,1,opt,name=lon,proto3" json:"lon,omitempty"`
	Lat float32 `protobuf:"fixed32,2,opt,name=lat,proto3" json:"lat,omitempty"`
}

func (x *WeatherData_Coord) Reset() {
	*x = WeatherData_Coord{}
	mi := &file_weather_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WeatherData_Coord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherData_Coord) ProtoMessage() {}

func (x *WeatherData_Coord) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherData_Coord.ProtoReflect.Descriptor instead.
func (*WeatherData_Coord) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{0, 0}
}

func (x *WeatherData_Coord) GetLon() float32 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *WeatherData_Coord) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

type WeatherData_WeatherCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Main        string `protobuf:"bytes,2,opt,name=main,proto3" json:"main,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Icon        string `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
}

func (x *WeatherData_WeatherCondition) Reset() {
	*x = WeatherData_WeatherCondition{}
	mi := &file_weather_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WeatherData_WeatherCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherData_WeatherCondition) ProtoMessage() {}

func (x *WeatherData_WeatherCondition) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherData_WeatherCondition.ProtoReflect.Descriptor instead.
func (*WeatherData_WeatherCondition) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{0, 1}
}

func (x *WeatherData_WeatherCondition) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WeatherData_WeatherCondition) GetMain() string {
	if x != nil {
		return x.Main
	}
	return ""
}

func (x *WeatherData_WeatherCondition) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *WeatherData_WeatherCondition) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

type WeatherData_MainData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temp      float32 `protobuf:"fixed32,1,opt,name=temp,proto3" json:"temp,omitempty"`
	FeelsLike float32 `protobuf:"fixed32,2,opt,name=feels_like,json=feelsLike,proto3" json:"feels_like,omitempty"`
	TempMin   float32 `protobuf:"fixed32,3,opt,name=temp_min,json=tempMin,proto3" json:"temp_min,omitempty"`
	TempMax   float32 `protobuf:"fixed32,4,opt,name=temp_max,json=tempMax,proto3" json:"temp_max,omitempty"`
	Pressure  int32   `protobuf:"varint,5,opt,name=pressure,proto3" json:"pressure,omitempty"`
	Humidity  int32   `protobuf:"varint,6,opt,name=humidity,proto3" json:"humidity,omitempty"`
	SeaLevel  int32   `protobuf:"varint,7,opt,name=sea_level,json=seaLevel,proto3" json:"sea_level,omitempty"`
	GrndLevel int32   `protobuf:"varint,8,opt,name=grnd_level,json=grndLevel,proto3" json:"grnd_level,omitempty"`
}

func (x *WeatherData_MainData) Reset() {
	*x = WeatherData_MainData{}
	mi := &file_weather_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WeatherData_MainData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherData_MainData) ProtoMessage() {}

func (x *WeatherData_MainData) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherData_MainData.ProtoReflect.Descriptor instead.
func (*WeatherData_MainData) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{0, 2}
}

func (x *WeatherData_MainData) GetTemp() float32 {
	if x != nil {
		return x.Temp
	}
	return 0
}

func (x *WeatherData_MainData) GetFeelsLike() float32 {
	if x != nil {
		return x.FeelsLike
	}
	return 0
}

func (x *WeatherData_MainData) GetTempMin() float32 {
	if x != nil {
		return x.TempMin
	}
	return 0
}

func (x *WeatherData_MainData) GetTempMax() float32 {
	if x != nil {
		return x.TempMax
	}
	return 0
}

func (x *WeatherData_MainData) GetPressure() int32 {
	if x != nil {
		return x.Pressure
	}
	return 0
}

func (x *WeatherData_MainData) GetHumidity() int32 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *WeatherData_MainData) GetSeaLevel() int32 {
	if x != nil {
		return x.SeaLevel
	}
	return 0
}

func (x *WeatherData_MainData) GetGrndLevel() int32 {
	if x != nil {
		return x.GrndLevel
	}
	return 0
}

type WeatherData_Wind struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Speed float32 `protobuf:"fixed32,1,opt,name=speed,proto3" json:"speed,omitempty"`
	Deg   int32   `protobuf:"varint,2,opt,name=deg,proto3" json:"deg,omitempty"`
	Gust  float32 `protobuf:"fixed32,3,opt,name=gust,proto3" json:"gust,omitempty"` // Added gust field
}

func (x *WeatherData_Wind) Reset() {
	*x = WeatherData_Wind{}
	mi := &file_weather_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WeatherData_Wind) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherData_Wind) ProtoMessage() {}

func (x *WeatherData_Wind) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherData_Wind.ProtoReflect.Descriptor instead.
func (*WeatherData_Wind) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{0, 3}
}

func (x *WeatherData_Wind) GetSpeed() float32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *WeatherData_Wind) GetDeg() int32 {
	if x != nil {
		return x.Deg
	}
	return 0
}

func (x *WeatherData_Wind) GetGust() float32 {
	if x != nil {
		return x.Gust
	}
	return 0
}

type WeatherData_Clouds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	All int32 `protobuf:"varint,1,opt,name=all,proto3" json:"all,omitempty"`
}

func (x *WeatherData_Clouds) Reset() {
	*x = WeatherData_Clouds{}
	mi := &file_weather_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WeatherData_Clouds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherData_Clouds) ProtoMessage() {}

func (x *WeatherData_Clouds) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherData_Clouds.ProtoReflect.Descriptor instead.
func (*WeatherData_Clouds) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{0, 4}
}

func (x *WeatherData_Clouds) GetAll() int32 {
	if x != nil {
		return x.All
	}
	return 0
}

type WeatherData_Sys struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Id      int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Country string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Sunrise int64  `protobuf:"varint,4,opt,name=sunrise,proto3" json:"sunrise,omitempty"`
	Sunset  int64  `protobuf:"varint,5,opt,name=sunset,proto3" json:"sunset,omitempty"`
}

func (x *WeatherData_Sys) Reset() {
	*x = WeatherData_Sys{}
	mi := &file_weather_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WeatherData_Sys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherData_Sys) ProtoMessage() {}

func (x *WeatherData_Sys) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherData_Sys.ProtoReflect.Descriptor instead.
func (*WeatherData_Sys) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{0, 5}
}

func (x *WeatherData_Sys) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *WeatherData_Sys) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WeatherData_Sys) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *WeatherData_Sys) GetSunrise() int64 {
	if x != nil {
		return x.Sunrise
	}
	return 0
}

func (x *WeatherData_Sys) GetSunset() int64 {
	if x != nil {
		return x.Sunset
	}
	return 0
}

var File_weather_proto protoreflect.FileDescriptor

var file_weather_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x22, 0xb5, 0x08, 0x0a, 0x0b, 0x57, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x05, 0x63, 0x6f, 0x6f, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x52, 0x05, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x12, 0x3f, 0x0a, 0x07, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x2d, 0x0a, 0x04, 0x77, 0x69, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x77, 0x69, 0x6e,
	0x64, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x52, 0x06,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x64, 0x74, 0x12, 0x2a, 0x0a, 0x03, 0x73, 0x79, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x57, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x79, 0x73, 0x52, 0x03, 0x73,
	0x79, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6f, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x63, 0x6f, 0x64, 0x1a, 0x2b, 0x0a, 0x05, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61,
	0x74, 0x1a, 0x6c, 0x0a, 0x10, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x1a,
	0xe7, 0x01, 0x0a, 0x08, 0x4d, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x74, 0x65, 0x6d, 0x70,
	0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65, 0x65, 0x6c, 0x73, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x66, 0x65, 0x65, 0x6c, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x07, 0x74, 0x65, 0x6d, 0x70, 0x4d, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x65,
	0x6d, 0x70, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x74, 0x65,
	0x6d, 0x70, 0x4d, 0x61, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x65, 0x61, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x73, 0x65, 0x61, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72,
	0x6e, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x67, 0x72, 0x6e, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x1a, 0x42, 0x0a, 0x04, 0x57, 0x69, 0x6e,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x65, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x67, 0x75, 0x73, 0x74, 0x1a, 0x1a, 0x0a,
	0x06, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x6c, 0x6c, 0x1a, 0x75, 0x0a, 0x03, 0x53, 0x79, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x6e, 0x73,
	0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x75, 0x6e, 0x73, 0x65, 0x74,
	0x22, 0x21, 0x0a, 0x0b, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x32, 0x4a, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x43, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x77, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x42,
	0x13, 0x5a, 0x11, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x79, 0x2f, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weather_proto_rawDescOnce sync.Once
	file_weather_proto_rawDescData = file_weather_proto_rawDesc
)

func file_weather_proto_rawDescGZIP() []byte {
	file_weather_proto_rawDescOnce.Do(func() {
		file_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_weather_proto_rawDescData)
	})
	return file_weather_proto_rawDescData
}

var file_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_weather_proto_goTypes = []any{
	(*WeatherData)(nil),                  // 0: weather.WeatherData
	(*CityRequest)(nil),                  // 1: weather.CityRequest
	(*WeatherData_Coord)(nil),            // 2: weather.WeatherData.Coord
	(*WeatherData_WeatherCondition)(nil), // 3: weather.WeatherData.WeatherCondition
	(*WeatherData_MainData)(nil),         // 4: weather.WeatherData.MainData
	(*WeatherData_Wind)(nil),             // 5: weather.WeatherData.Wind
	(*WeatherData_Clouds)(nil),           // 6: weather.WeatherData.Clouds
	(*WeatherData_Sys)(nil),              // 7: weather.WeatherData.Sys
}
var file_weather_proto_depIdxs = []int32{
	2, // 0: weather.WeatherData.coord:type_name -> weather.WeatherData.Coord
	3, // 1: weather.WeatherData.weather:type_name -> weather.WeatherData.WeatherCondition
	4, // 2: weather.WeatherData.main:type_name -> weather.WeatherData.MainData
	5, // 3: weather.WeatherData.wind:type_name -> weather.WeatherData.Wind
	6, // 4: weather.WeatherData.clouds:type_name -> weather.WeatherData.Clouds
	7, // 5: weather.WeatherData.sys:type_name -> weather.WeatherData.Sys
	1, // 6: weather.WeatherService.GetWeather:input_type -> weather.CityRequest
	0, // 7: weather.WeatherService.GetWeather:output_type -> weather.WeatherData
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_weather_proto_init() }
func file_weather_proto_init() {
	if File_weather_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weather_proto_goTypes,
		DependencyIndexes: file_weather_proto_depIdxs,
		MessageInfos:      file_weather_proto_msgTypes,
	}.Build()
	File_weather_proto = out.File
	file_weather_proto_rawDesc = nil
	file_weather_proto_goTypes = nil
	file_weather_proto_depIdxs = nil
}