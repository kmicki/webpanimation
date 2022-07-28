package webpanimation

/*
#cgo CFLAGS: -Wno-pointer-sign -DHAVE_CONFIG_H
#cgo !windows LDFLAGS: -lm

#include <stdlib.h>

#include "webp_encode.h"
#include "webp_mux.h"
#include "webp_demux.h"

WebPData* GenerateWebPData(uint8_t* indata, size_t n)
{
	uint8_t *data = (uint8_t*) malloc(n*sizeof(uint8_t));
	memcpy(data,indata,n*sizeof(uint8_t));
	struct WebPData webPDataInter = { data, n };
	struct WebPData *webPData = malloc(sizeof(WebPData));
	memcpy(webPData, &webPDataInter, sizeof(WebPData));

	return webPData;
}

void DeleteWebPData(WebPData* webPData)
{
	free((uint8_t*) webPData->bytes);
	free(webPData);
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

type WebPMuxError int

const (
	WebpMuxAbiVersion     = 0x0108
	WebpEncoderAbiVersion = 0x020f
	WebpDemuxAbiVersion   = 0x0107
)

const (
	WebpMuxOk              = WebPMuxError(C.WEBP_MUX_OK)
	WebpMuxNotFound        = WebPMuxError(C.WEBP_MUX_NOT_FOUND)
	WebpMuxInvalidArgument = WebPMuxError(C.WEBP_MUX_INVALID_ARGUMENT)
	WebpMuxBadData         = WebPMuxError(C.WEBP_MUX_BAD_DATA)
	WebpMuxMemoryError     = WebPMuxError(C.WEBP_MUX_MEMORY_ERROR)
	WebpMuxNotEnoughData   = WebPMuxError(C.WEBP_MUX_NOT_ENOUGH_DATA)
)

type WebPPicture C.WebPPicture
type WebPAnimEncoder C.WebPAnimEncoder
type WebPAnimDecoder C.WebPAnimDecoder
type WebPAnimInfo C.WebPAnimInfo
type WebPAnimEncoderOptions C.WebPAnimEncoderOptions
type WebPAnimDecoderOptions C.WebPAnimDecoderOptions
type WebPData C.WebPData
type WebPMux C.WebPMux
type WebPMuxAnimParams C.WebPMuxAnimParams
type webPConfig struct {
	webpConfig *C.WebPConfig
}

type WebPConfig interface {
	getRawPointer() *C.WebPConfig
	SetLossless(v int)
	GetLossless() int
	SetMethod(v int)
	SetImageHint(v int)
	SetTargetSize(v int)
	SetTargetPSNR(v float32)
	SetSegments(v int)
	SetSnsStrength(v int)
	SetFilterStrength(v int)
	SetFilterSharpness(v int)
	SetAutofilter(v int)
	SetAlphaCompression(v int)
	SetAlphaFiltering(v int)
	SetPass(v int)
	SetShowCompressed(v int)
	SetPreprocessing(v int)
	SetPartitions(v int)
	SetPartitionLimit(v int)
	SetEmulateJpegSize(v int)
	SetThreadLevel(v int)
	SetLowMemory(v int)
	SetNearLossless(v int)
	SetExact(v int)
	SetUseDeltaPalette(v int)
	SetUseSharpYuv(v int)
	SetAlphaQuality(v int)
	SetFilterType(v int)
	SetQuality(v float32)
}

func WebPDataClear(webPData *WebPData) {
	C.WebPDataClear((*C.WebPData)(unsafe.Pointer(webPData)))
}

func WebPMuxDelete(webPMux *WebPMux) {
	C.WebPMuxDelete((*C.WebPMux)(unsafe.Pointer(webPMux)))
}

func WebPPictureFree(webPPicture *WebPPicture) {
	C.WebPPictureFree((*C.WebPPicture)(unsafe.Pointer(webPPicture)))
}

func WebPAnimEncoderDelete(webPAnimEncoder *WebPAnimEncoder) {
	C.WebPAnimEncoderDelete((*C.WebPAnimEncoder)(unsafe.Pointer(webPAnimEncoder)))
}

func (wmap *WebPMuxAnimParams) SetBgcolor(v uint32) {
	((*C.WebPMuxAnimParams)(wmap)).bgcolor = (C.uint32_t)(v)
}

func (wmap *WebPMuxAnimParams) SetLoopCount(v int) {
	(*C.WebPMuxAnimParams)(wmap).loop_count = (C.int)(v)
}

func WebPPictureInit(webPPicture *WebPPicture) int {
	return int(C.WebPPictureInit((*C.WebPPicture)(unsafe.Pointer(webPPicture))))
}

func (wpp *WebPPicture) SetWidth(v int) {
	((*C.WebPPicture)(wpp)).width = (C.int)(v)
}

func (wpp *WebPPicture) SetHeight(v int) {
	((*C.WebPPicture)(wpp)).height = (C.int)(v)
}

func (wpp WebPPicture) GetWidth() int {
	return int(((C.WebPPicture)(wpp)).width)
}

func (wpp WebPPicture) GetHeight() int {
	return int(((C.WebPPicture)(wpp)).height)
}

func (wpp *WebPPicture) SetUseArgb(v int) {
	((*C.WebPPicture)(wpp)).use_argb = (C.int)(v)
}

func (wpd WebPData) GetBytes() []byte {
	return C.GoBytes(unsafe.Pointer(((C.WebPData)(wpd)).bytes), (C.int)(((C.WebPData)(wpd)).size))
}

func GenerateWebPData(data []byte, size uint) *WebPData {
	return (*WebPData)(C.GenerateWebPData(
		(*C.uint8_t)(unsafe.Pointer(&data[0])),
		(C.size_t)(size),
	))
}

func WebPDataInit(webPData *WebPData) {
	C.WebPDataInit((*C.WebPData)(unsafe.Pointer(webPData)))
}

// NewWebpConfig create webpconfig instance
func NewWebpConfig() WebPConfig {
	webpcfg := &webPConfig{}
	webpcfg.webpConfig = &C.WebPConfig{}
	WebPConfigInitInternal(webpcfg)
	return webpcfg
}

func WebPConfigInitInternal(config WebPConfig) int {
	return int(C.WebPConfigInitInternal(
		config.getRawPointer(),
		(C.WebPPreset)(0),
		(C.float)(75.0),
		(C.int)(WebpEncoderAbiVersion),
	))
}

func (webpCfg *webPConfig) getRawPointer() *C.WebPConfig {
	return webpCfg.webpConfig
}

func (webpCfg *webPConfig) SetLossless(v int) {
	webpCfg.webpConfig.lossless = (C.int)(v)
}

func (webpCfg *webPConfig) GetLossless() int {
	return int(webpCfg.webpConfig.lossless)
}

func (webpCfg *webPConfig) SetMethod(v int) {
	webpCfg.webpConfig.method = (C.int)(v)
}

func (webpCfg *webPConfig) SetImageHint(v int) {
	webpCfg.webpConfig.image_hint = (C.WebPImageHint)(v)
}

func (webpCfg *webPConfig) SetTargetSize(v int) {
	webpCfg.webpConfig.target_size = (C.int)(v)
}

func (webpCfg *webPConfig) SetTargetPSNR(v float32) {
	webpCfg.webpConfig.target_PSNR = (C.float)(v)
}

func (webpCfg *webPConfig) SetSegments(v int) {
	webpCfg.webpConfig.segments = (C.int)(v)
}

func (webpCfg *webPConfig) SetSnsStrength(v int) {
	webpCfg.webpConfig.sns_strength = (C.int)(v)
}

func (webpCfg *webPConfig) SetFilterStrength(v int) {
	webpCfg.webpConfig.filter_strength = (C.int)(v)
}

func (webpCfg *webPConfig) SetFilterSharpness(v int) {
	webpCfg.webpConfig.filter_sharpness = (C.int)(v)
}

func (webpCfg *webPConfig) SetAutofilter(v int) {
	webpCfg.webpConfig.autofilter = (C.int)(v)
}

func (webpCfg *webPConfig) SetAlphaCompression(v int) {
	webpCfg.webpConfig.alpha_compression = (C.int)(v)
}

func (webpCfg *webPConfig) SetAlphaFiltering(v int) {
	webpCfg.webpConfig.alpha_filtering = (C.int)(v)
}

func (webpCfg *webPConfig) SetPass(v int) {
	webpCfg.webpConfig.pass = (C.int)(v)
}

func (webpCfg *webPConfig) SetShowCompressed(v int) {
	webpCfg.webpConfig.show_compressed = (C.int)(v)
}

func (webpCfg *webPConfig) SetPreprocessing(v int) {
	webpCfg.webpConfig.preprocessing = (C.int)(v)
}

func (webpCfg *webPConfig) SetPartitions(v int) {
	webpCfg.webpConfig.partitions = (C.int)(v)
}

func (webpCfg *webPConfig) SetPartitionLimit(v int) {
	webpCfg.webpConfig.partition_limit = (C.int)(v)
}

func (webpCfg *webPConfig) SetEmulateJpegSize(v int) {
	webpCfg.webpConfig.emulate_jpeg_size = (C.int)(v)
}

func (webpCfg *webPConfig) SetThreadLevel(v int) {
	webpCfg.webpConfig.thread_level = (C.int)(v)
}

func (webpCfg *webPConfig) SetLowMemory(v int) {
	webpCfg.webpConfig.low_memory = (C.int)(v)
}

func (webpCfg *webPConfig) SetNearLossless(v int) {
	webpCfg.webpConfig.near_lossless = (C.int)(v)
}

func (webpCfg *webPConfig) SetExact(v int) {
	webpCfg.webpConfig.exact = (C.int)(v)
}

func (webpCfg *webPConfig) SetUseDeltaPalette(v int) {
	webpCfg.webpConfig.use_delta_palette = (C.int)(v)
}

func (webpCfg *webPConfig) SetUseSharpYuv(v int) {
	webpCfg.webpConfig.use_sharp_yuv = (C.int)(v)
}

func (webpCfg *webPConfig) SetAlphaQuality(v int) {
	webpCfg.webpConfig.alpha_quality = (C.int)(v)
}

func (webpCfg *webPConfig) SetFilterType(v int) {
	webpCfg.webpConfig.filter_type = (C.int)(v)
}

func (webpCfg *webPConfig) SetQuality(v float32) {
	webpCfg.webpConfig.quality = (C.float)(v)
}

func (encOptions *WebPAnimEncoderOptions) GetAnimParams() WebPMuxAnimParams {
	return WebPMuxAnimParams(((*C.WebPAnimEncoderOptions)(encOptions)).anim_params)
}

func (encOptions *WebPAnimEncoderOptions) SetAnimParams(v WebPMuxAnimParams) {
	((*C.WebPAnimEncoderOptions)(encOptions)).anim_params = (C.WebPMuxAnimParams)(v)
}

func (encOptions *WebPAnimEncoderOptions) SetMinimizeSize(v int) {
	((*C.WebPAnimEncoderOptions)(encOptions)).minimize_size = (C.int)(v)
}

func (encOptions *WebPAnimEncoderOptions) SetKmin(v int) {
	((*C.WebPAnimEncoderOptions)(encOptions)).kmin = (C.int)(v)
}

func (encOptions *WebPAnimEncoderOptions) SetKmax(v int) {
	((*C.WebPAnimEncoderOptions)(encOptions)).kmax = (C.int)(v)
}

func (encOptions *WebPAnimEncoderOptions) SetAllowMixed(v int) {
	((*C.WebPAnimEncoderOptions)(encOptions)).allow_mixed = (C.int)(v)
}

func (encOptions *WebPAnimEncoderOptions) SetVerbose(v int) {
	((*C.WebPAnimEncoderOptions)(encOptions)).verbose = (C.int)(v)
}

func WebPAnimEncoderOptionsInitInternal(webPAnimEncoderOptions *WebPAnimEncoderOptions) int {
	return int(C.WebPAnimEncoderOptionsInitInternal(
		(*C.WebPAnimEncoderOptions)(unsafe.Pointer(webPAnimEncoderOptions)),
		(C.int)(WebpMuxAbiVersion),
	))
}

func WebPPictureImportRGBA(data []byte, stride int, webPPicture *WebPPicture) error {
	res := int(C.WebPPictureImportRGBA(
		(*C.WebPPicture)(unsafe.Pointer(webPPicture)),
		(*C.uint8_t)(unsafe.Pointer(&data[0])),
		(C.int)(stride),
	))
	if res == 0 {
		return errors.New("error: WebPPictureImportBGRA")
	}
	return nil
}

func WebPAnimEncoderNewInternal(width, height int, webPAnimEncoderOptions *WebPAnimEncoderOptions) *WebPAnimEncoder {
	return (*WebPAnimEncoder)(C.WebPAnimEncoderNewInternal(
		(C.int)(width),
		(C.int)(height),
		(*C.WebPAnimEncoderOptions)(unsafe.Pointer(webPAnimEncoderOptions)),
		(C.int)(WebpMuxAbiVersion),
	))
}

func WebPAnimEncoderAdd(webPAnimEncoder *WebPAnimEncoder, webPPicture *WebPPicture, timestamp int, webpcfg WebPConfig) int {
	return int(C.WebPAnimEncoderAdd(
		(*C.WebPAnimEncoder)(unsafe.Pointer(webPAnimEncoder)),
		(*C.WebPPicture)(unsafe.Pointer(webPPicture)),
		(C.int)(timestamp),
		webpcfg.getRawPointer(),
	))
}

func WebPAnimEncoderAssemble(webPAnimEncoder *WebPAnimEncoder, webPData *WebPData) int {
	return int(C.WebPAnimEncoderAssemble(
		(*C.WebPAnimEncoder)(unsafe.Pointer(webPAnimEncoder)),
		(*C.WebPData)(unsafe.Pointer(webPData)),
	))
}

func WebPMuxCreateInternal(webPData *WebPData, copyData int) *WebPMux {
	return (*WebPMux)(C.WebPMuxCreateInternal(
		(*C.WebPData)(unsafe.Pointer(webPData)),
		(C.int)(copyData),
		(C.int)(WebpMuxAbiVersion),
	))
}

func WebPMuxSetAnimationParams(webPMux *WebPMux, webPMuxAnimParams *WebPMuxAnimParams) WebPMuxError {
	return (WebPMuxError)(C.WebPMuxSetAnimationParams(
		(*C.WebPMux)(unsafe.Pointer(webPMux)),
		(*C.WebPMuxAnimParams)(unsafe.Pointer(webPMuxAnimParams)),
	))
}

func WebPMuxGetAnimationParams(webPMux *WebPMux, webPMuxAnimParams *WebPMuxAnimParams) WebPMuxError {
	return (WebPMuxError)(C.WebPMuxGetAnimationParams(
		(*C.WebPMux)(unsafe.Pointer(webPMux)),
		(*C.WebPMuxAnimParams)(unsafe.Pointer(webPMuxAnimParams)),
	))
}

func WebPMuxAssemble(webPMux *WebPMux, webPData *WebPData) WebPMuxError {
	return (WebPMuxError)(C.WebPMuxAssemble(
		(*C.WebPMux)(unsafe.Pointer(webPMux)),
		(*C.WebPData)(unsafe.Pointer(webPData)),
	))
}

// Decoder

func DeleteWebPData(webPData *WebPData) {
	C.DeleteWebPData((*C.WebPData)(unsafe.Pointer(webPData)))
}

func WebPAnimDecoderDelete(webPAnimDecoder *WebPAnimDecoder) {
	C.WebPAnimDecoderDelete(
		(*C.WebPAnimDecoder)(unsafe.Pointer(webPAnimDecoder)),
	)
}

func WebPAnimDecoderNew(webPData *WebPData) *WebPAnimDecoder {
	return (*WebPAnimDecoder)(C.WebPAnimDecoderNew(
		(*C.WebPData)(unsafe.Pointer(webPData)),
		nil,
	))
}

func WebPAnimDecoderGetNext(webPAnimDecoder *WebPAnimDecoder, buffer *([]byte), timestamp *int, size int) bool {

	var tstamp C.int
	var buf *C.uchar

	result := int(C.WebPAnimDecoderGetNext(
		(*C.WebPAnimDecoder)(unsafe.Pointer(webPAnimDecoder)),
		(**C.uchar)(unsafe.Pointer(&buf)),
		(*C.int)(unsafe.Pointer(&tstamp)),
	)) != 0

	if !result {
		return false
	}

	*buffer = C.GoBytes(unsafe.Pointer(buf), C.int(size))
	*timestamp = int(tstamp)

	return true
}

func WebPAnimDecoderGetInfo(webPAnimDecoder *WebPAnimDecoder, webPAnimInfo *WebPAnimInfo) bool {
	return int(C.WebPAnimDecoderGetInfo(
		(*C.WebPAnimDecoder)(unsafe.Pointer(webPAnimDecoder)),
		(*C.WebPAnimInfo)(unsafe.Pointer(webPAnimInfo)),
	)) != 0
}

func WebPAnimDecoderReset(webPAnimDecoder *WebPAnimDecoder) {
	C.WebPAnimDecoderReset(
		(*C.WebPAnimDecoder)(unsafe.Pointer(webPAnimDecoder)),
	)
}

func (wpi WebPAnimInfo) GetWidth() uint32 {
	return uint32(((C.WebPAnimInfo)(wpi)).canvas_width)
}

func (wpi WebPAnimInfo) GetHeight() uint32 {
	return uint32(((C.WebPAnimInfo)(wpi)).canvas_height)
}

func (wpi WebPAnimInfo) GetLoopCount() uint32 {
	return uint32(((C.WebPAnimInfo)(wpi)).loop_count)
}

func (wpi WebPAnimInfo) GetBgColor() uint32 {
	return uint32(((C.WebPAnimInfo)(wpi)).bgcolor)
}

func (wpi WebPAnimInfo) GetFrameCnt() uint32 {
	return uint32(((C.WebPAnimInfo)(wpi)).frame_count)
}
